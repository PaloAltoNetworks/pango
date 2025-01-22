package adminrole

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
	Suffix = []string{"admin-role"}
)

type Entry struct {
	Name        string
	Description *string
	Role        *Role

	Misc map[string][]generic.Xml
}

type Role struct {
	Device *RoleDevice
	Vsys   *RoleVsys
}
type RoleDevice struct {
	Cli     *string
	Restapi *RoleDeviceRestapi
	Webui   *RoleDeviceWebui
	Xmlapi  *RoleDeviceXmlapi
}
type RoleDeviceRestapi struct {
	Device   *RoleDeviceRestapiDevice
	Network  *RoleDeviceRestapiNetwork
	Objects  *RoleDeviceRestapiObjects
	Policies *RoleDeviceRestapiPolicies
	System   *RoleDeviceRestapiSystem
}
type RoleDeviceRestapiDevice struct {
	EmailServerProfiles    *string
	HttpServerProfiles     *string
	LdapServerProfiles     *string
	LogInterfaceSetting    *string
	SnmpTrapServerProfiles *string
	SyslogServerProfiles   *string
	VirtualSystems         *string
}
type RoleDeviceRestapiNetwork struct {
	AggregateEthernetInterfaces             *string
	BfdNetworkProfiles                      *string
	BgpRoutingProfiles                      *string
	DhcpRelays                              *string
	DhcpServers                             *string
	DnsProxies                              *string
	EthernetInterfaces                      *string
	GlobalprotectClientlessAppGroups        *string
	GlobalprotectClientlessApps             *string
	GlobalprotectGateways                   *string
	GlobalprotectIpsecCryptoNetworkProfiles *string
	GlobalprotectMdmServers                 *string
	GlobalprotectPortals                    *string
	GreTunnels                              *string
	IkeCryptoNetworkProfiles                *string
	IkeGatewayNetworkProfiles               *string
	InterfaceManagementNetworkProfiles      *string
	IpsecCryptoNetworkProfiles              *string
	IpsecTunnels                            *string
	Lldp                                    *string
	LldpNetworkProfiles                     *string
	LogicalRouters                          *string
	LoopbackInterfaces                      *string
	QosInterfaces                           *string
	QosNetworkProfiles                      *string
	SdwanInterfaceProfiles                  *string
	SdwanInterfaces                         *string
	TunnelInterfaces                        *string
	TunnelMonitorNetworkProfiles            *string
	VirtualRouters                          *string
	VirtualWires                            *string
	VlanInterfaces                          *string
	Vlans                                   *string
	ZoneProtectionNetworkProfiles           *string
	Zones                                   *string
}
type RoleDeviceRestapiObjects struct {
	AddressGroups                           *string
	Addresses                               *string
	AntiSpywareSecurityProfiles             *string
	AntivirusSecurityProfiles               *string
	ApplicationFilters                      *string
	ApplicationGroups                       *string
	Applications                            *string
	AuthenticationEnforcements              *string
	CustomDataPatterns                      *string
	CustomSpywareSignatures                 *string
	CustomUrlCategories                     *string
	CustomVulnerabilitySignatures           *string
	DataFilteringSecurityProfiles           *string
	DecryptionProfiles                      *string
	Devices                                 *string
	DosProtectionSecurityProfiles           *string
	DynamicUserGroups                       *string
	ExternalDynamicLists                    *string
	FileBlockingSecurityProfiles            *string
	GlobalprotectHipObjects                 *string
	GlobalprotectHipProfiles                *string
	GtpProtectionSecurityProfiles           *string
	LogForwardingProfiles                   *string
	PacketBrokerProfiles                    *string
	Regions                                 *string
	Schedules                               *string
	SctpProtectionSecurityProfiles          *string
	SdwanErrorCorrectionProfiles            *string
	SdwanPathQualityProfiles                *string
	SdwanSaasQualityProfiles                *string
	SdwanTrafficDistributionProfiles        *string
	SecurityProfileGroups                   *string
	ServiceGroups                           *string
	Services                                *string
	Tags                                    *string
	UrlFilteringSecurityProfiles            *string
	VulnerabilityProtectionSecurityProfiles *string
	WildfireAnalysisSecurityProfiles        *string
}
type RoleDeviceRestapiPolicies struct {
	ApplicationOverrideRules   *string
	AuthenticationRules        *string
	DecryptionRules            *string
	DosRules                   *string
	NatRules                   *string
	NetworkPacketBrokerRules   *string
	PolicyBasedForwardingRules *string
	QosRules                   *string
	SdwanRules                 *string
	SecurityRules              *string
	TunnelInspectionRules      *string
}
type RoleDeviceRestapiSystem struct {
	Configuration *string
}
type RoleDeviceWebui struct {
	Acc        *string
	Commit     *RoleDeviceWebuiCommit
	Dashboard  *string
	Device     *RoleDeviceWebuiDevice
	Global     *RoleDeviceWebuiGlobal
	Monitor    *RoleDeviceWebuiMonitor
	Network    *RoleDeviceWebuiNetwork
	Objects    *RoleDeviceWebuiObjects
	Operations *RoleDeviceWebuiOperations
	Policies   *RoleDeviceWebuiPolicies
	Privacy    *RoleDeviceWebuiPrivacy
	Save       *RoleDeviceWebuiSave
	Tasks      *string
	Validate   *string
}
type RoleDeviceWebuiCommit struct {
	CommitForOtherAdmins *string
	Device               *string
	ObjectLevelChanges   *string
}
type RoleDeviceWebuiDevice struct {
	AccessDomain           *string
	AdminRoles             *string
	Administrators         *string
	AuthenticationProfile  *string
	AuthenticationSequence *string
	BlockPages             *string
	CertificateManagement  *RoleDeviceWebuiDeviceCertificateManagement
	ConfigAudit            *string
	DataRedistribution     *string
	DeviceQuarantine       *string
	DhcpSyslogServer       *string
	DynamicUpdates         *string
	GlobalProtectClient    *string
	HighAvailability       *string
	Licenses               *string
	LocalUserDatabase      *RoleDeviceWebuiDeviceLocalUserDatabase
	LogFwdCard             *string
	LogSettings            *RoleDeviceWebuiDeviceLogSettings
	MasterKey              *string
	Plugins                *string
	PolicyRecommendations  *RoleDeviceWebuiDevicePolicyRecommendations
	ScheduledLogExport     *string
	ServerProfile          *RoleDeviceWebuiDeviceServerProfile
	Setup                  *RoleDeviceWebuiDeviceSetup
	SharedGateways         *string
	Software               *string
	Support                *string
	Troubleshooting        *string
	UserIdentification     *string
	VirtualSystems         *string
	VmInfoSource           *string
}
type RoleDeviceWebuiDeviceCertificateManagement struct {
	CertificateProfile     *string
	Certificates           *string
	OcspResponder          *string
	Scep                   *string
	SshServiceProfile      *string
	SslDecryptionExclusion *string
	SslTlsServiceProfile   *string
}
type RoleDeviceWebuiDeviceLocalUserDatabase struct {
	UserGroups *string
	Users      *string
}
type RoleDeviceWebuiDeviceLogSettings struct {
	CcAlarm       *string
	Config        *string
	Correlation   *string
	Globalprotect *string
	Hipmatch      *string
	Iptag         *string
	ManageLog     *string
	System        *string
	UserId        *string
}
type RoleDeviceWebuiDevicePolicyRecommendations struct {
	Iot  *string
	Saas *string
}
type RoleDeviceWebuiDeviceServerProfile struct {
	Dns      *string
	Email    *string
	Http     *string
	Kerberos *string
	Ldap     *string
	Mfa      *string
	Netflow  *string
	Radius   *string
	SamlIdp  *string
	Scp      *string
	SnmpTrap *string
	Syslog   *string
	Tacplus  *string
}
type RoleDeviceWebuiDeviceSetup struct {
	ContentId  *string
	Hsm        *string
	Interfaces *string
	Management *string
	Operations *string
	Services   *string
	Session    *string
	Telemetry  *string
	Wildfire   *string
}
type RoleDeviceWebuiGlobal struct {
	SystemAlarms *string
}
type RoleDeviceWebuiMonitor struct {
	AppScope                   *string
	ApplicationReports         *string
	AutomatedCorrelationEngine *RoleDeviceWebuiMonitorAutomatedCorrelationEngine
	BlockIpList                *string
	Botnet                     *string
	CustomReports              *RoleDeviceWebuiMonitorCustomReports
	ExternalLogs               *string
	GtpReports                 *string
	Logs                       *RoleDeviceWebuiMonitorLogs
	PacketCapture              *string
	PdfReports                 *RoleDeviceWebuiMonitorPdfReports
	SctpReports                *string
	SessionBrowser             *string
	ThreatReports              *string
	TrafficReports             *string
	UrlFilteringReports        *string
	ViewCustomReports          *string
}
type RoleDeviceWebuiMonitorAutomatedCorrelationEngine struct {
	CorrelatedEvents   *string
	CorrelationObjects *string
}
type RoleDeviceWebuiMonitorCustomReports struct {
	ApplicationStatistics *string
	Auth                  *string
	DataFilteringLog      *string
	DecryptionLog         *string
	DecryptionSummary     *string
	Globalprotect         *string
	GtpLog                *string
	GtpSummary            *string
	Hipmatch              *string
	Iptag                 *string
	SctpLog               *string
	SctpSummary           *string
	ThreatLog             *string
	ThreatSummary         *string
	TrafficLog            *string
	TrafficSummary        *string
	TunnelLog             *string
	TunnelSummary         *string
	UrlLog                *string
	UrlSummary            *string
	Userid                *string
	WildfireLog           *string
}
type RoleDeviceWebuiMonitorLogs struct {
	Alarm          *string
	Authentication *string
	Configuration  *string
	DataFiltering  *string
	Decryption     *string
	Globalprotect  *string
	Gtp            *string
	Hipmatch       *string
	Iptag          *string
	Sctp           *string
	System         *string
	Threat         *string
	Traffic        *string
	Tunnel         *string
	Url            *string
	Userid         *string
	Wildfire       *string
}
type RoleDeviceWebuiMonitorPdfReports struct {
	EmailScheduler             *string
	ManagePdfSummary           *string
	PdfSummaryReports          *string
	ReportGroups               *string
	SaasApplicationUsageReport *string
	UserActivityReport         *string
}
type RoleDeviceWebuiNetwork struct {
	Dhcp                  *string
	DnsProxy              *string
	GlobalProtect         *RoleDeviceWebuiNetworkGlobalProtect
	GreTunnels            *string
	Interfaces            *string
	IpsecTunnels          *string
	Lldp                  *string
	NetworkProfiles       *RoleDeviceWebuiNetworkNetworkProfiles
	Qos                   *string
	Routing               *RoleDeviceWebuiNetworkRouting
	SdwanInterfaceProfile *string
	SecureWebGateway      *string
	VirtualRouters        *string
	VirtualWires          *string
	Vlans                 *string
	Zones                 *string
}
type RoleDeviceWebuiNetworkGlobalProtect struct {
	ClientlessAppGroups *string
	ClientlessApps      *string
	Gateways            *string
	Mdm                 *string
	Portals             *string
}
type RoleDeviceWebuiNetworkNetworkProfiles struct {
	BfdProfile       *string
	GpAppIpsecCrypto *string
	IkeCrypto        *string
	IkeGateways      *string
	InterfaceMgmt    *string
	IpsecCrypto      *string
	LldpProfile      *string
	QosProfile       *string
	TunnelMonitor    *string
	ZoneProtection   *string
}
type RoleDeviceWebuiNetworkRouting struct {
	LogicalRouters  *string
	RoutingProfiles *RoleDeviceWebuiNetworkRoutingRoutingProfiles
}
type RoleDeviceWebuiNetworkRoutingRoutingProfiles struct {
	Bfd       *string
	Bgp       *string
	Filters   *string
	Multicast *string
	Ospf      *string
	Ospfv3    *string
	Ripv2     *string
}
type RoleDeviceWebuiObjects struct {
	AddressGroups         *string
	Addresses             *string
	ApplicationFilters    *string
	ApplicationGroups     *string
	Applications          *string
	Authentication        *string
	CustomObjects         *RoleDeviceWebuiObjectsCustomObjects
	Decryption            *RoleDeviceWebuiObjectsDecryption
	Devices               *string
	DynamicBlockLists     *string
	DynamicUserGroups     *string
	GlobalProtect         *RoleDeviceWebuiObjectsGlobalProtect
	LogForwarding         *string
	PacketBrokerProfile   *string
	Regions               *string
	Schedules             *string
	Sdwan                 *RoleDeviceWebuiObjectsSdwan
	SecurityProfileGroups *string
	SecurityProfiles      *RoleDeviceWebuiObjectsSecurityProfiles
	ServiceGroups         *string
	Services              *string
	Tags                  *string
}
type RoleDeviceWebuiObjectsCustomObjects struct {
	DataPatterns  *string
	Spyware       *string
	UrlCategory   *string
	Vulnerability *string
}
type RoleDeviceWebuiObjectsDecryption struct {
	DecryptionProfile *string
}
type RoleDeviceWebuiObjectsGlobalProtect struct {
	HipObjects  *string
	HipProfiles *string
}
type RoleDeviceWebuiObjectsSdwan struct {
	SdwanDistProfile            *string
	SdwanErrorCorrectionProfile *string
	SdwanProfile                *string
	SdwanSaasQualityProfile     *string
}
type RoleDeviceWebuiObjectsSecurityProfiles struct {
	AntiSpyware             *string
	Antivirus               *string
	DataFiltering           *string
	DosProtection           *string
	FileBlocking            *string
	GtpProtection           *string
	SctpProtection          *string
	UrlFiltering            *string
	VulnerabilityProtection *string
	WildfireAnalysis        *string
}
type RoleDeviceWebuiOperations struct {
	DownloadCoreFiles       *string
	DownloadPcapFiles       *string
	GenerateStatsDumpFile   *string
	GenerateTechSupportFile *string
	Reboot                  *string
}
type RoleDeviceWebuiPolicies struct {
	ApplicationOverrideRulebase *string
	AuthenticationRulebase      *string
	DosRulebase                 *string
	NatRulebase                 *string
	NetworkPacketBrokerRulebase *string
	PbfRulebase                 *string
	QosRulebase                 *string
	RuleHitCountReset           *string
	SdwanRulebase               *string
	SecurityRulebase            *string
	SslDecryptionRulebase       *string
	TunnelInspectRulebase       *string
}
type RoleDeviceWebuiPrivacy struct {
	ShowFullIpAddresses           *string
	ShowUserNamesInLogsAndReports *string
	ViewPcapFiles                 *string
}
type RoleDeviceWebuiSave struct {
	ObjectLevelChanges *string
	PartialSave        *string
	SaveForOtherAdmins *string
}
type RoleDeviceXmlapi struct {
	Commit *string
	Config *string
	Export *string
	Import *string
	Iot    *string
	Log    *string
	Op     *string
	Report *string
	UserId *string
}
type RoleVsys struct {
	Cli     *string
	Restapi *RoleVsysRestapi
	Webui   *RoleVsysWebui
	Xmlapi  *RoleVsysXmlapi
}
type RoleVsysRestapi struct {
	Device   *RoleVsysRestapiDevice
	Network  *RoleVsysRestapiNetwork
	Objects  *RoleVsysRestapiObjects
	Policies *RoleVsysRestapiPolicies
	System   *RoleVsysRestapiSystem
}
type RoleVsysRestapiDevice struct {
	EmailServerProfiles    *string
	HttpServerProfiles     *string
	LdapServerProfiles     *string
	LogInterfaceSetting    *string
	SnmpTrapServerProfiles *string
	SyslogServerProfiles   *string
	VirtualSystems         *string
}
type RoleVsysRestapiNetwork struct {
	GlobalprotectClientlessAppGroups *string
	GlobalprotectClientlessApps      *string
	GlobalprotectGateways            *string
	GlobalprotectMdmServers          *string
	GlobalprotectPortals             *string
	SdwanInterfaceProfiles           *string
	Zones                            *string
}
type RoleVsysRestapiObjects struct {
	AddressGroups                           *string
	Addresses                               *string
	AntiSpywareSecurityProfiles             *string
	AntivirusSecurityProfiles               *string
	ApplicationFilters                      *string
	ApplicationGroups                       *string
	Applications                            *string
	AuthenticationEnforcements              *string
	CustomDataPatterns                      *string
	CustomSpywareSignatures                 *string
	CustomUrlCategories                     *string
	CustomVulnerabilitySignatures           *string
	DataFilteringSecurityProfiles           *string
	DecryptionProfiles                      *string
	Devices                                 *string
	DosProtectionSecurityProfiles           *string
	DynamicUserGroups                       *string
	ExternalDynamicLists                    *string
	FileBlockingSecurityProfiles            *string
	GlobalprotectHipObjects                 *string
	GlobalprotectHipProfiles                *string
	GtpProtectionSecurityProfiles           *string
	LogForwardingProfiles                   *string
	PacketBrokerProfiles                    *string
	Regions                                 *string
	Schedules                               *string
	SctpProtectionSecurityProfiles          *string
	SdwanErrorCorrectionProfiles            *string
	SdwanPathQualityProfiles                *string
	SdwanSaasQualityProfiles                *string
	SdwanTrafficDistributionProfiles        *string
	SecurityProfileGroups                   *string
	ServiceGroups                           *string
	Services                                *string
	Tags                                    *string
	UrlFilteringSecurityProfiles            *string
	VulnerabilityProtectionSecurityProfiles *string
	WildfireAnalysisSecurityProfiles        *string
}
type RoleVsysRestapiPolicies struct {
	ApplicationOverrideRules   *string
	AuthenticationRules        *string
	DecryptionRules            *string
	DosRules                   *string
	NatRules                   *string
	NetworkPacketBrokerRules   *string
	PolicyBasedForwardingRules *string
	QosRules                   *string
	SdwanRules                 *string
	SecurityRules              *string
	TunnelInspectionRules      *string
}
type RoleVsysRestapiSystem struct {
	Configuration *string
}
type RoleVsysWebui struct {
	Acc        *string
	Commit     *RoleVsysWebuiCommit
	Dashboard  *string
	Device     *RoleVsysWebuiDevice
	Monitor    *RoleVsysWebuiMonitor
	Network    *RoleVsysWebuiNetwork
	Objects    *RoleVsysWebuiObjects
	Operations *RoleVsysWebuiOperations
	Policies   *RoleVsysWebuiPolicies
	Privacy    *RoleVsysWebuiPrivacy
	Save       *RoleVsysWebuiSave
	Tasks      *string
	Validate   *string
}
type RoleVsysWebuiCommit struct {
	CommitForOtherAdmins *string
	VirtualSystems       *string
}
type RoleVsysWebuiDevice struct {
	Administrators         *string
	AuthenticationProfile  *string
	AuthenticationSequence *string
	BlockPages             *string
	CertificateManagement  *RoleVsysWebuiDeviceCertificateManagement
	DataRedistribution     *string
	DeviceQuarantine       *string
	DhcpSyslogServer       *string
	LocalUserDatabase      *RoleVsysWebuiDeviceLocalUserDatabase
	LogSettings            *RoleVsysWebuiDeviceLogSettings
	PolicyRecommendations  *RoleVsysWebuiDevicePolicyRecommendations
	ServerProfile          *RoleVsysWebuiDeviceServerProfile
	Setup                  *RoleVsysWebuiDeviceSetup
	Troubleshooting        *string
	UserIdentification     *string
	VmInfoSource           *string
}
type RoleVsysWebuiDeviceCertificateManagement struct {
	CertificateProfile     *string
	Certificates           *string
	OcspResponder          *string
	Scep                   *string
	SshServiceProfile      *string
	SslDecryptionExclusion *string
	SslTlsServiceProfile   *string
}
type RoleVsysWebuiDeviceLocalUserDatabase struct {
	UserGroups *string
	Users      *string
}
type RoleVsysWebuiDeviceLogSettings struct {
	Config        *string
	Correlation   *string
	Globalprotect *string
	Hipmatch      *string
	Iptag         *string
	System        *string
	UserId        *string
}
type RoleVsysWebuiDevicePolicyRecommendations struct {
	Iot  *string
	Saas *string
}
type RoleVsysWebuiDeviceServerProfile struct {
	Dns      *string
	Email    *string
	Http     *string
	Kerberos *string
	Ldap     *string
	Mfa      *string
	Netflow  *string
	Radius   *string
	SamlIdp  *string
	Scp      *string
	SnmpTrap *string
	Syslog   *string
	Tacplus  *string
}
type RoleVsysWebuiDeviceSetup struct {
	ContentId  *string
	Hsm        *string
	Interfaces *string
	Management *string
	Operations *string
	Services   *string
	Session    *string
	Telemetry  *string
	Wildfire   *string
}
type RoleVsysWebuiMonitor struct {
	AppScope                   *string
	AutomatedCorrelationEngine *RoleVsysWebuiMonitorAutomatedCorrelationEngine
	BlockIpList                *string
	CustomReports              *RoleVsysWebuiMonitorCustomReports
	ExternalLogs               *string
	Logs                       *RoleVsysWebuiMonitorLogs
	PdfReports                 *RoleVsysWebuiMonitorPdfReports
	SessionBrowser             *string
	ViewCustomReports          *string
}
type RoleVsysWebuiMonitorAutomatedCorrelationEngine struct {
	CorrelatedEvents   *string
	CorrelationObjects *string
}
type RoleVsysWebuiMonitorCustomReports struct {
	ApplicationStatistics *string
	Auth                  *string
	DataFilteringLog      *string
	DecryptionLog         *string
	DecryptionSummary     *string
	Globalprotect         *string
	GtpLog                *string
	GtpSummary            *string
	Hipmatch              *string
	Iptag                 *string
	SctpLog               *string
	SctpSummary           *string
	ThreatLog             *string
	ThreatSummary         *string
	TrafficLog            *string
	TrafficSummary        *string
	TunnelLog             *string
	TunnelSummary         *string
	UrlLog                *string
	UrlSummary            *string
	Userid                *string
	WildfireLog           *string
}
type RoleVsysWebuiMonitorLogs struct {
	Authentication *string
	DataFiltering  *string
	Decryption     *string
	Globalprotect  *string
	Gtp            *string
	Hipmatch       *string
	Iptag          *string
	Sctp           *string
	Threat         *string
	Traffic        *string
	Tunnel         *string
	Url            *string
	Userid         *string
	Wildfire       *string
}
type RoleVsysWebuiMonitorPdfReports struct {
	EmailScheduler             *string
	ManagePdfSummary           *string
	PdfSummaryReports          *string
	ReportGroups               *string
	SaasApplicationUsageReport *string
	UserActivityReport         *string
}
type RoleVsysWebuiNetwork struct {
	GlobalProtect         *RoleVsysWebuiNetworkGlobalProtect
	SdwanInterfaceProfile *string
	Zones                 *string
}
type RoleVsysWebuiNetworkGlobalProtect struct {
	ClientlessAppGroups *string
	ClientlessApps      *string
	Gateways            *string
	Mdm                 *string
	Portals             *string
}
type RoleVsysWebuiObjects struct {
	AddressGroups         *string
	Addresses             *string
	ApplicationFilters    *string
	ApplicationGroups     *string
	Applications          *string
	Authentication        *string
	CustomObjects         *RoleVsysWebuiObjectsCustomObjects
	Decryption            *RoleVsysWebuiObjectsDecryption
	Devices               *string
	DynamicBlockLists     *string
	DynamicUserGroups     *string
	GlobalProtect         *RoleVsysWebuiObjectsGlobalProtect
	LogForwarding         *string
	PacketBrokerProfile   *string
	Regions               *string
	Schedules             *string
	Sdwan                 *RoleVsysWebuiObjectsSdwan
	SecurityProfileGroups *string
	SecurityProfiles      *RoleVsysWebuiObjectsSecurityProfiles
	ServiceGroups         *string
	Services              *string
	Tags                  *string
}
type RoleVsysWebuiObjectsCustomObjects struct {
	DataPatterns  *string
	Spyware       *string
	UrlCategory   *string
	Vulnerability *string
}
type RoleVsysWebuiObjectsDecryption struct {
	DecryptionProfile *string
}
type RoleVsysWebuiObjectsGlobalProtect struct {
	HipObjects  *string
	HipProfiles *string
}
type RoleVsysWebuiObjectsSdwan struct {
	SdwanDistProfile            *string
	SdwanErrorCorrectionProfile *string
	SdwanProfile                *string
	SdwanSaasQualityProfile     *string
}
type RoleVsysWebuiObjectsSecurityProfiles struct {
	AntiSpyware             *string
	Antivirus               *string
	DataFiltering           *string
	DosProtection           *string
	FileBlocking            *string
	GtpProtection           *string
	SctpProtection          *string
	UrlFiltering            *string
	VulnerabilityProtection *string
	WildfireAnalysis        *string
}
type RoleVsysWebuiOperations struct {
	DownloadCoreFiles       *string
	DownloadPcapFiles       *string
	GenerateStatsDumpFile   *string
	GenerateTechSupportFile *string
	Reboot                  *string
}
type RoleVsysWebuiPolicies struct {
	ApplicationOverrideRulebase *string
	AuthenticationRulebase      *string
	DosRulebase                 *string
	NatRulebase                 *string
	NetworkPacketBrokerRulebase *string
	PbfRulebase                 *string
	QosRulebase                 *string
	RuleHitCountReset           *string
	SdwanRulebase               *string
	SecurityRulebase            *string
	SslDecryptionRulebase       *string
	TunnelInspectRulebase       *string
}
type RoleVsysWebuiPrivacy struct {
	ShowFullIpAddresses           *string
	ShowUserNamesInLogsAndReports *string
	ViewPcapFiles                 *string
}
type RoleVsysWebuiSave struct {
	ObjectLevelChanges *string
	PartialSave        *string
	SaveForOtherAdmins *string
}
type RoleVsysXmlapi struct {
	Commit *string
	Config *string
	Export *string
	Import *string
	Iot    *string
	Log    *string
	Op     *string
	Report *string
	UserId *string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName     xml.Name `xml:"entry"`
	Name        string   `xml:"name,attr"`
	Description *string  `xml:"description,omitempty"`
	Role        *RoleXml `xml:"role,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleXml struct {
	Device *RoleDeviceXml `xml:"device,omitempty"`
	Vsys   *RoleVsysXml   `xml:"vsys,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceXml struct {
	Cli     *string               `xml:"cli,omitempty"`
	Restapi *RoleDeviceRestapiXml `xml:"restapi,omitempty"`
	Webui   *RoleDeviceWebuiXml   `xml:"webui,omitempty"`
	Xmlapi  *RoleDeviceXmlapiXml  `xml:"xmlapi,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceRestapiXml struct {
	Device   *RoleDeviceRestapiDeviceXml   `xml:"device,omitempty"`
	Network  *RoleDeviceRestapiNetworkXml  `xml:"network,omitempty"`
	Objects  *RoleDeviceRestapiObjectsXml  `xml:"objects,omitempty"`
	Policies *RoleDeviceRestapiPoliciesXml `xml:"policies,omitempty"`
	System   *RoleDeviceRestapiSystemXml   `xml:"system,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceRestapiDeviceXml struct {
	EmailServerProfiles    *string `xml:"email-server-profiles,omitempty"`
	HttpServerProfiles     *string `xml:"http-server-profiles,omitempty"`
	LdapServerProfiles     *string `xml:"ldap-server-profiles,omitempty"`
	LogInterfaceSetting    *string `xml:"log-interface-setting,omitempty"`
	SnmpTrapServerProfiles *string `xml:"snmp-trap-server-profiles,omitempty"`
	SyslogServerProfiles   *string `xml:"syslog-server-profiles,omitempty"`
	VirtualSystems         *string `xml:"virtual-systems,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceRestapiNetworkXml struct {
	AggregateEthernetInterfaces             *string `xml:"aggregate-ethernet-interfaces,omitempty"`
	BfdNetworkProfiles                      *string `xml:"bfd-network-profiles,omitempty"`
	BgpRoutingProfiles                      *string `xml:"bgp-routing-profiles,omitempty"`
	DhcpRelays                              *string `xml:"dhcp-relays,omitempty"`
	DhcpServers                             *string `xml:"dhcp-servers,omitempty"`
	DnsProxies                              *string `xml:"dns-proxies,omitempty"`
	EthernetInterfaces                      *string `xml:"ethernet-interfaces,omitempty"`
	GlobalprotectClientlessAppGroups        *string `xml:"globalprotect-clientless-app-groups,omitempty"`
	GlobalprotectClientlessApps             *string `xml:"globalprotect-clientless-apps,omitempty"`
	GlobalprotectGateways                   *string `xml:"globalprotect-gateways,omitempty"`
	GlobalprotectIpsecCryptoNetworkProfiles *string `xml:"globalprotect-ipsec-crypto-network-profiles,omitempty"`
	GlobalprotectMdmServers                 *string `xml:"globalprotect-mdm-servers,omitempty"`
	GlobalprotectPortals                    *string `xml:"globalprotect-portals,omitempty"`
	GreTunnels                              *string `xml:"gre-tunnels,omitempty"`
	IkeCryptoNetworkProfiles                *string `xml:"ike-crypto-network-profiles,omitempty"`
	IkeGatewayNetworkProfiles               *string `xml:"ike-gateway-network-profiles,omitempty"`
	InterfaceManagementNetworkProfiles      *string `xml:"interface-management-network-profiles,omitempty"`
	IpsecCryptoNetworkProfiles              *string `xml:"ipsec-crypto-network-profiles,omitempty"`
	IpsecTunnels                            *string `xml:"ipsec-tunnels,omitempty"`
	Lldp                                    *string `xml:"lldp,omitempty"`
	LldpNetworkProfiles                     *string `xml:"lldp-network-profiles,omitempty"`
	LogicalRouters                          *string `xml:"logical-routers,omitempty"`
	LoopbackInterfaces                      *string `xml:"loopback-interfaces,omitempty"`
	QosInterfaces                           *string `xml:"qos-interfaces,omitempty"`
	QosNetworkProfiles                      *string `xml:"qos-network-profiles,omitempty"`
	SdwanInterfaceProfiles                  *string `xml:"sdwan-interface-profiles,omitempty"`
	SdwanInterfaces                         *string `xml:"sdwan-interfaces,omitempty"`
	TunnelInterfaces                        *string `xml:"tunnel-interfaces,omitempty"`
	TunnelMonitorNetworkProfiles            *string `xml:"tunnel-monitor-network-profiles,omitempty"`
	VirtualRouters                          *string `xml:"virtual-routers,omitempty"`
	VirtualWires                            *string `xml:"virtual-wires,omitempty"`
	VlanInterfaces                          *string `xml:"vlan-interfaces,omitempty"`
	Vlans                                   *string `xml:"vlans,omitempty"`
	ZoneProtectionNetworkProfiles           *string `xml:"zone-protection-network-profiles,omitempty"`
	Zones                                   *string `xml:"zones,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceRestapiObjectsXml struct {
	AddressGroups                           *string `xml:"address-groups,omitempty"`
	Addresses                               *string `xml:"addresses,omitempty"`
	AntiSpywareSecurityProfiles             *string `xml:"anti-spyware-security-profiles,omitempty"`
	AntivirusSecurityProfiles               *string `xml:"antivirus-security-profiles,omitempty"`
	ApplicationFilters                      *string `xml:"application-filters,omitempty"`
	ApplicationGroups                       *string `xml:"application-groups,omitempty"`
	Applications                            *string `xml:"applications,omitempty"`
	AuthenticationEnforcements              *string `xml:"authentication-enforcements,omitempty"`
	CustomDataPatterns                      *string `xml:"custom-data-patterns,omitempty"`
	CustomSpywareSignatures                 *string `xml:"custom-spyware-signatures,omitempty"`
	CustomUrlCategories                     *string `xml:"custom-url-categories,omitempty"`
	CustomVulnerabilitySignatures           *string `xml:"custom-vulnerability-signatures,omitempty"`
	DataFilteringSecurityProfiles           *string `xml:"data-filtering-security-profiles,omitempty"`
	DecryptionProfiles                      *string `xml:"decryption-profiles,omitempty"`
	Devices                                 *string `xml:"devices,omitempty"`
	DosProtectionSecurityProfiles           *string `xml:"dos-protection-security-profiles,omitempty"`
	DynamicUserGroups                       *string `xml:"dynamic-user-groups,omitempty"`
	ExternalDynamicLists                    *string `xml:"external-dynamic-lists,omitempty"`
	FileBlockingSecurityProfiles            *string `xml:"file-blocking-security-profiles,omitempty"`
	GlobalprotectHipObjects                 *string `xml:"globalprotect-hip-objects,omitempty"`
	GlobalprotectHipProfiles                *string `xml:"globalprotect-hip-profiles,omitempty"`
	GtpProtectionSecurityProfiles           *string `xml:"gtp-protection-security-profiles,omitempty"`
	LogForwardingProfiles                   *string `xml:"log-forwarding-profiles,omitempty"`
	PacketBrokerProfiles                    *string `xml:"packet-broker-profiles,omitempty"`
	Regions                                 *string `xml:"regions,omitempty"`
	Schedules                               *string `xml:"schedules,omitempty"`
	SctpProtectionSecurityProfiles          *string `xml:"sctp-protection-security-profiles,omitempty"`
	SdwanErrorCorrectionProfiles            *string `xml:"sdwan-error-correction-profiles,omitempty"`
	SdwanPathQualityProfiles                *string `xml:"sdwan-path-quality-profiles,omitempty"`
	SdwanSaasQualityProfiles                *string `xml:"sdwan-saas-quality-profiles,omitempty"`
	SdwanTrafficDistributionProfiles        *string `xml:"sdwan-traffic-distribution-profiles,omitempty"`
	SecurityProfileGroups                   *string `xml:"security-profile-groups,omitempty"`
	ServiceGroups                           *string `xml:"service-groups,omitempty"`
	Services                                *string `xml:"services,omitempty"`
	Tags                                    *string `xml:"tags,omitempty"`
	UrlFilteringSecurityProfiles            *string `xml:"url-filtering-security-profiles,omitempty"`
	VulnerabilityProtectionSecurityProfiles *string `xml:"vulnerability-protection-security-profiles,omitempty"`
	WildfireAnalysisSecurityProfiles        *string `xml:"wildfire-analysis-security-profiles,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceRestapiPoliciesXml struct {
	ApplicationOverrideRules   *string `xml:"application-override-rules,omitempty"`
	AuthenticationRules        *string `xml:"authentication-rules,omitempty"`
	DecryptionRules            *string `xml:"decryption-rules,omitempty"`
	DosRules                   *string `xml:"dos-rules,omitempty"`
	NatRules                   *string `xml:"nat-rules,omitempty"`
	NetworkPacketBrokerRules   *string `xml:"network-packet-broker-rules,omitempty"`
	PolicyBasedForwardingRules *string `xml:"policy-based-forwarding-rules,omitempty"`
	QosRules                   *string `xml:"qos-rules,omitempty"`
	SdwanRules                 *string `xml:"sdwan-rules,omitempty"`
	SecurityRules              *string `xml:"security-rules,omitempty"`
	TunnelInspectionRules      *string `xml:"tunnel-inspection-rules,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceRestapiSystemXml struct {
	Configuration *string `xml:"configuration,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiXml struct {
	Acc        *string                       `xml:"acc,omitempty"`
	Commit     *RoleDeviceWebuiCommitXml     `xml:"commit,omitempty"`
	Dashboard  *string                       `xml:"dashboard,omitempty"`
	Device     *RoleDeviceWebuiDeviceXml     `xml:"device,omitempty"`
	Global     *RoleDeviceWebuiGlobalXml     `xml:"global,omitempty"`
	Monitor    *RoleDeviceWebuiMonitorXml    `xml:"monitor,omitempty"`
	Network    *RoleDeviceWebuiNetworkXml    `xml:"network,omitempty"`
	Objects    *RoleDeviceWebuiObjectsXml    `xml:"objects,omitempty"`
	Operations *RoleDeviceWebuiOperationsXml `xml:"operations,omitempty"`
	Policies   *RoleDeviceWebuiPoliciesXml   `xml:"policies,omitempty"`
	Privacy    *RoleDeviceWebuiPrivacyXml    `xml:"privacy,omitempty"`
	Save       *RoleDeviceWebuiSaveXml       `xml:"save,omitempty"`
	Tasks      *string                       `xml:"tasks,omitempty"`
	Validate   *string                       `xml:"validate,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiCommitXml struct {
	CommitForOtherAdmins *string `xml:"commit-for-other-admins,omitempty"`
	Device               *string `xml:"device,omitempty"`
	ObjectLevelChanges   *string `xml:"object-level-changes,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiDeviceXml struct {
	AccessDomain           *string                                        `xml:"access-domain,omitempty"`
	AdminRoles             *string                                        `xml:"admin-roles,omitempty"`
	Administrators         *string                                        `xml:"administrators,omitempty"`
	AuthenticationProfile  *string                                        `xml:"authentication-profile,omitempty"`
	AuthenticationSequence *string                                        `xml:"authentication-sequence,omitempty"`
	BlockPages             *string                                        `xml:"block-pages,omitempty"`
	CertificateManagement  *RoleDeviceWebuiDeviceCertificateManagementXml `xml:"certificate-management,omitempty"`
	ConfigAudit            *string                                        `xml:"config-audit,omitempty"`
	DataRedistribution     *string                                        `xml:"data-redistribution,omitempty"`
	DeviceQuarantine       *string                                        `xml:"device-quarantine,omitempty"`
	DhcpSyslogServer       *string                                        `xml:"dhcp-syslog-server,omitempty"`
	DynamicUpdates         *string                                        `xml:"dynamic-updates,omitempty"`
	GlobalProtectClient    *string                                        `xml:"global-protect-client,omitempty"`
	HighAvailability       *string                                        `xml:"high-availability,omitempty"`
	Licenses               *string                                        `xml:"licenses,omitempty"`
	LocalUserDatabase      *RoleDeviceWebuiDeviceLocalUserDatabaseXml     `xml:"local-user-database,omitempty"`
	LogFwdCard             *string                                        `xml:"log-fwd-card,omitempty"`
	LogSettings            *RoleDeviceWebuiDeviceLogSettingsXml           `xml:"log-settings,omitempty"`
	MasterKey              *string                                        `xml:"master-key,omitempty"`
	Plugins                *string                                        `xml:"plugins,omitempty"`
	PolicyRecommendations  *RoleDeviceWebuiDevicePolicyRecommendationsXml `xml:"policy-recommendations,omitempty"`
	ScheduledLogExport     *string                                        `xml:"scheduled-log-export,omitempty"`
	ServerProfile          *RoleDeviceWebuiDeviceServerProfileXml         `xml:"server-profile,omitempty"`
	Setup                  *RoleDeviceWebuiDeviceSetupXml                 `xml:"setup,omitempty"`
	SharedGateways         *string                                        `xml:"shared-gateways,omitempty"`
	Software               *string                                        `xml:"software,omitempty"`
	Support                *string                                        `xml:"support,omitempty"`
	Troubleshooting        *string                                        `xml:"troubleshooting,omitempty"`
	UserIdentification     *string                                        `xml:"user-identification,omitempty"`
	VirtualSystems         *string                                        `xml:"virtual-systems,omitempty"`
	VmInfoSource           *string                                        `xml:"vm-info-source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiDeviceCertificateManagementXml struct {
	CertificateProfile     *string `xml:"certificate-profile,omitempty"`
	Certificates           *string `xml:"certificates,omitempty"`
	OcspResponder          *string `xml:"ocsp-responder,omitempty"`
	Scep                   *string `xml:"scep,omitempty"`
	SshServiceProfile      *string `xml:"ssh-service-profile,omitempty"`
	SslDecryptionExclusion *string `xml:"ssl-decryption-exclusion,omitempty"`
	SslTlsServiceProfile   *string `xml:"ssl-tls-service-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiDeviceLocalUserDatabaseXml struct {
	UserGroups *string `xml:"user-groups,omitempty"`
	Users      *string `xml:"users,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiDeviceLogSettingsXml struct {
	CcAlarm       *string `xml:"cc-alarm,omitempty"`
	Config        *string `xml:"config,omitempty"`
	Correlation   *string `xml:"correlation,omitempty"`
	Globalprotect *string `xml:"globalprotect,omitempty"`
	Hipmatch      *string `xml:"hipmatch,omitempty"`
	Iptag         *string `xml:"iptag,omitempty"`
	ManageLog     *string `xml:"manage-log,omitempty"`
	System        *string `xml:"system,omitempty"`
	UserId        *string `xml:"user-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiDevicePolicyRecommendationsXml struct {
	Iot  *string `xml:"iot,omitempty"`
	Saas *string `xml:"saas,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiDeviceServerProfileXml struct {
	Dns      *string `xml:"dns,omitempty"`
	Email    *string `xml:"email,omitempty"`
	Http     *string `xml:"http,omitempty"`
	Kerberos *string `xml:"kerberos,omitempty"`
	Ldap     *string `xml:"ldap,omitempty"`
	Mfa      *string `xml:"mfa,omitempty"`
	Netflow  *string `xml:"netflow,omitempty"`
	Radius   *string `xml:"radius,omitempty"`
	SamlIdp  *string `xml:"saml_idp,omitempty"`
	Scp      *string `xml:"scp,omitempty"`
	SnmpTrap *string `xml:"snmp-trap,omitempty"`
	Syslog   *string `xml:"syslog,omitempty"`
	Tacplus  *string `xml:"tacplus,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiDeviceSetupXml struct {
	ContentId  *string `xml:"content-id,omitempty"`
	Hsm        *string `xml:"hsm,omitempty"`
	Interfaces *string `xml:"interfaces,omitempty"`
	Management *string `xml:"management,omitempty"`
	Operations *string `xml:"operations,omitempty"`
	Services   *string `xml:"services,omitempty"`
	Session    *string `xml:"session,omitempty"`
	Telemetry  *string `xml:"telemetry,omitempty"`
	Wildfire   *string `xml:"wildfire,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiGlobalXml struct {
	SystemAlarms *string `xml:"system-alarms,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiMonitorXml struct {
	AppScope                   *string                                              `xml:"app-scope,omitempty"`
	ApplicationReports         *string                                              `xml:"application-reports,omitempty"`
	AutomatedCorrelationEngine *RoleDeviceWebuiMonitorAutomatedCorrelationEngineXml `xml:"automated-correlation-engine,omitempty"`
	BlockIpList                *string                                              `xml:"block-ip-list,omitempty"`
	Botnet                     *string                                              `xml:"botnet,omitempty"`
	CustomReports              *RoleDeviceWebuiMonitorCustomReportsXml              `xml:"custom-reports,omitempty"`
	ExternalLogs               *string                                              `xml:"external-logs,omitempty"`
	GtpReports                 *string                                              `xml:"gtp-reports,omitempty"`
	Logs                       *RoleDeviceWebuiMonitorLogsXml                       `xml:"logs,omitempty"`
	PacketCapture              *string                                              `xml:"packet-capture,omitempty"`
	PdfReports                 *RoleDeviceWebuiMonitorPdfReportsXml                 `xml:"pdf-reports,omitempty"`
	SctpReports                *string                                              `xml:"sctp-reports,omitempty"`
	SessionBrowser             *string                                              `xml:"session-browser,omitempty"`
	ThreatReports              *string                                              `xml:"threat-reports,omitempty"`
	TrafficReports             *string                                              `xml:"traffic-reports,omitempty"`
	UrlFilteringReports        *string                                              `xml:"url-filtering-reports,omitempty"`
	ViewCustomReports          *string                                              `xml:"view-custom-reports,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiMonitorAutomatedCorrelationEngineXml struct {
	CorrelatedEvents   *string `xml:"correlated-events,omitempty"`
	CorrelationObjects *string `xml:"correlation-objects,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiMonitorCustomReportsXml struct {
	ApplicationStatistics *string `xml:"application-statistics,omitempty"`
	Auth                  *string `xml:"auth,omitempty"`
	DataFilteringLog      *string `xml:"data-filtering-log,omitempty"`
	DecryptionLog         *string `xml:"decryption-log,omitempty"`
	DecryptionSummary     *string `xml:"decryption-summary,omitempty"`
	Globalprotect         *string `xml:"globalprotect,omitempty"`
	GtpLog                *string `xml:"gtp-log,omitempty"`
	GtpSummary            *string `xml:"gtp-summary,omitempty"`
	Hipmatch              *string `xml:"hipmatch,omitempty"`
	Iptag                 *string `xml:"iptag,omitempty"`
	SctpLog               *string `xml:"sctp-log,omitempty"`
	SctpSummary           *string `xml:"sctp-summary,omitempty"`
	ThreatLog             *string `xml:"threat-log,omitempty"`
	ThreatSummary         *string `xml:"threat-summary,omitempty"`
	TrafficLog            *string `xml:"traffic-log,omitempty"`
	TrafficSummary        *string `xml:"traffic-summary,omitempty"`
	TunnelLog             *string `xml:"tunnel-log,omitempty"`
	TunnelSummary         *string `xml:"tunnel-summary,omitempty"`
	UrlLog                *string `xml:"url-log,omitempty"`
	UrlSummary            *string `xml:"url-summary,omitempty"`
	Userid                *string `xml:"userid,omitempty"`
	WildfireLog           *string `xml:"wildfire-log,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiMonitorLogsXml struct {
	Alarm          *string `xml:"alarm,omitempty"`
	Authentication *string `xml:"authentication,omitempty"`
	Configuration  *string `xml:"configuration,omitempty"`
	DataFiltering  *string `xml:"data-filtering,omitempty"`
	Decryption     *string `xml:"decryption,omitempty"`
	Globalprotect  *string `xml:"globalprotect,omitempty"`
	Gtp            *string `xml:"gtp,omitempty"`
	Hipmatch       *string `xml:"hipmatch,omitempty"`
	Iptag          *string `xml:"iptag,omitempty"`
	Sctp           *string `xml:"sctp,omitempty"`
	System         *string `xml:"system,omitempty"`
	Threat         *string `xml:"threat,omitempty"`
	Traffic        *string `xml:"traffic,omitempty"`
	Tunnel         *string `xml:"tunnel,omitempty"`
	Url            *string `xml:"url,omitempty"`
	Userid         *string `xml:"userid,omitempty"`
	Wildfire       *string `xml:"wildfire,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiMonitorPdfReportsXml struct {
	EmailScheduler             *string `xml:"email-scheduler,omitempty"`
	ManagePdfSummary           *string `xml:"manage-pdf-summary,omitempty"`
	PdfSummaryReports          *string `xml:"pdf-summary-reports,omitempty"`
	ReportGroups               *string `xml:"report-groups,omitempty"`
	SaasApplicationUsageReport *string `xml:"saas-application-usage-report,omitempty"`
	UserActivityReport         *string `xml:"user-activity-report,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiNetworkXml struct {
	Dhcp                  *string                                   `xml:"dhcp,omitempty"`
	DnsProxy              *string                                   `xml:"dns-proxy,omitempty"`
	GlobalProtect         *RoleDeviceWebuiNetworkGlobalProtectXml   `xml:"global-protect,omitempty"`
	GreTunnels            *string                                   `xml:"gre-tunnels,omitempty"`
	Interfaces            *string                                   `xml:"interfaces,omitempty"`
	IpsecTunnels          *string                                   `xml:"ipsec-tunnels,omitempty"`
	Lldp                  *string                                   `xml:"lldp,omitempty"`
	NetworkProfiles       *RoleDeviceWebuiNetworkNetworkProfilesXml `xml:"network-profiles,omitempty"`
	Qos                   *string                                   `xml:"qos,omitempty"`
	Routing               *RoleDeviceWebuiNetworkRoutingXml         `xml:"routing,omitempty"`
	SdwanInterfaceProfile *string                                   `xml:"sdwan-interface-profile,omitempty"`
	SecureWebGateway      *string                                   `xml:"secure-web-gateway,omitempty"`
	VirtualRouters        *string                                   `xml:"virtual-routers,omitempty"`
	VirtualWires          *string                                   `xml:"virtual-wires,omitempty"`
	Vlans                 *string                                   `xml:"vlans,omitempty"`
	Zones                 *string                                   `xml:"zones,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiNetworkGlobalProtectXml struct {
	ClientlessAppGroups *string `xml:"clientless-app-groups,omitempty"`
	ClientlessApps      *string `xml:"clientless-apps,omitempty"`
	Gateways            *string `xml:"gateways,omitempty"`
	Mdm                 *string `xml:"mdm,omitempty"`
	Portals             *string `xml:"portals,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiNetworkNetworkProfilesXml struct {
	BfdProfile       *string `xml:"bfd-profile,omitempty"`
	GpAppIpsecCrypto *string `xml:"gp-app-ipsec-crypto,omitempty"`
	IkeCrypto        *string `xml:"ike-crypto,omitempty"`
	IkeGateways      *string `xml:"ike-gateways,omitempty"`
	InterfaceMgmt    *string `xml:"interface-mgmt,omitempty"`
	IpsecCrypto      *string `xml:"ipsec-crypto,omitempty"`
	LldpProfile      *string `xml:"lldp-profile,omitempty"`
	QosProfile       *string `xml:"qos-profile,omitempty"`
	TunnelMonitor    *string `xml:"tunnel-monitor,omitempty"`
	ZoneProtection   *string `xml:"zone-protection,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiNetworkRoutingXml struct {
	LogicalRouters  *string                                          `xml:"logical-routers,omitempty"`
	RoutingProfiles *RoleDeviceWebuiNetworkRoutingRoutingProfilesXml `xml:"routing-profiles,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiNetworkRoutingRoutingProfilesXml struct {
	Bfd       *string `xml:"bfd,omitempty"`
	Bgp       *string `xml:"bgp,omitempty"`
	Filters   *string `xml:"filters,omitempty"`
	Multicast *string `xml:"multicast,omitempty"`
	Ospf      *string `xml:"ospf,omitempty"`
	Ospfv3    *string `xml:"ospfv3,omitempty"`
	Ripv2     *string `xml:"ripv2,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiObjectsXml struct {
	AddressGroups         *string                                    `xml:"address-groups,omitempty"`
	Addresses             *string                                    `xml:"addresses,omitempty"`
	ApplicationFilters    *string                                    `xml:"application-filters,omitempty"`
	ApplicationGroups     *string                                    `xml:"application-groups,omitempty"`
	Applications          *string                                    `xml:"applications,omitempty"`
	Authentication        *string                                    `xml:"authentication,omitempty"`
	CustomObjects         *RoleDeviceWebuiObjectsCustomObjectsXml    `xml:"custom-objects,omitempty"`
	Decryption            *RoleDeviceWebuiObjectsDecryptionXml       `xml:"decryption,omitempty"`
	Devices               *string                                    `xml:"devices,omitempty"`
	DynamicBlockLists     *string                                    `xml:"dynamic-block-lists,omitempty"`
	DynamicUserGroups     *string                                    `xml:"dynamic-user-groups,omitempty"`
	GlobalProtect         *RoleDeviceWebuiObjectsGlobalProtectXml    `xml:"global-protect,omitempty"`
	LogForwarding         *string                                    `xml:"log-forwarding,omitempty"`
	PacketBrokerProfile   *string                                    `xml:"packet-broker-profile,omitempty"`
	Regions               *string                                    `xml:"regions,omitempty"`
	Schedules             *string                                    `xml:"schedules,omitempty"`
	Sdwan                 *RoleDeviceWebuiObjectsSdwanXml            `xml:"sdwan,omitempty"`
	SecurityProfileGroups *string                                    `xml:"security-profile-groups,omitempty"`
	SecurityProfiles      *RoleDeviceWebuiObjectsSecurityProfilesXml `xml:"security-profiles,omitempty"`
	ServiceGroups         *string                                    `xml:"service-groups,omitempty"`
	Services              *string                                    `xml:"services,omitempty"`
	Tags                  *string                                    `xml:"tags,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiObjectsCustomObjectsXml struct {
	DataPatterns  *string `xml:"data-patterns,omitempty"`
	Spyware       *string `xml:"spyware,omitempty"`
	UrlCategory   *string `xml:"url-category,omitempty"`
	Vulnerability *string `xml:"vulnerability,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiObjectsDecryptionXml struct {
	DecryptionProfile *string `xml:"decryption-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiObjectsGlobalProtectXml struct {
	HipObjects  *string `xml:"hip-objects,omitempty"`
	HipProfiles *string `xml:"hip-profiles,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiObjectsSdwanXml struct {
	SdwanDistProfile            *string `xml:"sdwan-dist-profile,omitempty"`
	SdwanErrorCorrectionProfile *string `xml:"sdwan-error-correction-profile,omitempty"`
	SdwanProfile                *string `xml:"sdwan-profile,omitempty"`
	SdwanSaasQualityProfile     *string `xml:"sdwan-saas-quality-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiObjectsSecurityProfilesXml struct {
	AntiSpyware             *string `xml:"anti-spyware,omitempty"`
	Antivirus               *string `xml:"antivirus,omitempty"`
	DataFiltering           *string `xml:"data-filtering,omitempty"`
	DosProtection           *string `xml:"dos-protection,omitempty"`
	FileBlocking            *string `xml:"file-blocking,omitempty"`
	GtpProtection           *string `xml:"gtp-protection,omitempty"`
	SctpProtection          *string `xml:"sctp-protection,omitempty"`
	UrlFiltering            *string `xml:"url-filtering,omitempty"`
	VulnerabilityProtection *string `xml:"vulnerability-protection,omitempty"`
	WildfireAnalysis        *string `xml:"wildfire-analysis,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiOperationsXml struct {
	DownloadCoreFiles       *string `xml:"download-core-files,omitempty"`
	DownloadPcapFiles       *string `xml:"download-pcap-files,omitempty"`
	GenerateStatsDumpFile   *string `xml:"generate-stats-dump-file,omitempty"`
	GenerateTechSupportFile *string `xml:"generate-tech-support-file,omitempty"`
	Reboot                  *string `xml:"reboot,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiPoliciesXml struct {
	ApplicationOverrideRulebase *string `xml:"application-override-rulebase,omitempty"`
	AuthenticationRulebase      *string `xml:"authentication-rulebase,omitempty"`
	DosRulebase                 *string `xml:"dos-rulebase,omitempty"`
	NatRulebase                 *string `xml:"nat-rulebase,omitempty"`
	NetworkPacketBrokerRulebase *string `xml:"network-packet-broker-rulebase,omitempty"`
	PbfRulebase                 *string `xml:"pbf-rulebase,omitempty"`
	QosRulebase                 *string `xml:"qos-rulebase,omitempty"`
	RuleHitCountReset           *string `xml:"rule-hit-count-reset,omitempty"`
	SdwanRulebase               *string `xml:"sdwan-rulebase,omitempty"`
	SecurityRulebase            *string `xml:"security-rulebase,omitempty"`
	SslDecryptionRulebase       *string `xml:"ssl-decryption-rulebase,omitempty"`
	TunnelInspectRulebase       *string `xml:"tunnel-inspect-rulebase,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiPrivacyXml struct {
	ShowFullIpAddresses           *string `xml:"show-full-ip-addresses,omitempty"`
	ShowUserNamesInLogsAndReports *string `xml:"show-user-names-in-logs-and-reports,omitempty"`
	ViewPcapFiles                 *string `xml:"view-pcap-files,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceWebuiSaveXml struct {
	ObjectLevelChanges *string `xml:"object-level-changes,omitempty"`
	PartialSave        *string `xml:"partial-save,omitempty"`
	SaveForOtherAdmins *string `xml:"save-for-other-admins,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleDeviceXmlapiXml struct {
	Commit *string `xml:"commit,omitempty"`
	Config *string `xml:"config,omitempty"`
	Export *string `xml:"export,omitempty"`
	Import *string `xml:"import,omitempty"`
	Iot    *string `xml:"iot,omitempty"`
	Log    *string `xml:"log,omitempty"`
	Op     *string `xml:"op,omitempty"`
	Report *string `xml:"report,omitempty"`
	UserId *string `xml:"user-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysXml struct {
	Cli     *string             `xml:"cli,omitempty"`
	Restapi *RoleVsysRestapiXml `xml:"restapi,omitempty"`
	Webui   *RoleVsysWebuiXml   `xml:"webui,omitempty"`
	Xmlapi  *RoleVsysXmlapiXml  `xml:"xmlapi,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysRestapiXml struct {
	Device   *RoleVsysRestapiDeviceXml   `xml:"device,omitempty"`
	Network  *RoleVsysRestapiNetworkXml  `xml:"network,omitempty"`
	Objects  *RoleVsysRestapiObjectsXml  `xml:"objects,omitempty"`
	Policies *RoleVsysRestapiPoliciesXml `xml:"policies,omitempty"`
	System   *RoleVsysRestapiSystemXml   `xml:"system,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysRestapiDeviceXml struct {
	EmailServerProfiles    *string `xml:"email-server-profiles,omitempty"`
	HttpServerProfiles     *string `xml:"http-server-profiles,omitempty"`
	LdapServerProfiles     *string `xml:"ldap-server-profiles,omitempty"`
	LogInterfaceSetting    *string `xml:"log-interface-setting,omitempty"`
	SnmpTrapServerProfiles *string `xml:"snmp-trap-server-profiles,omitempty"`
	SyslogServerProfiles   *string `xml:"syslog-server-profiles,omitempty"`
	VirtualSystems         *string `xml:"virtual-systems,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysRestapiNetworkXml struct {
	GlobalprotectClientlessAppGroups *string `xml:"globalprotect-clientless-app-groups,omitempty"`
	GlobalprotectClientlessApps      *string `xml:"globalprotect-clientless-apps,omitempty"`
	GlobalprotectGateways            *string `xml:"globalprotect-gateways,omitempty"`
	GlobalprotectMdmServers          *string `xml:"globalprotect-mdm-servers,omitempty"`
	GlobalprotectPortals             *string `xml:"globalprotect-portals,omitempty"`
	SdwanInterfaceProfiles           *string `xml:"sdwan-interface-profiles,omitempty"`
	Zones                            *string `xml:"zones,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysRestapiObjectsXml struct {
	AddressGroups                           *string `xml:"address-groups,omitempty"`
	Addresses                               *string `xml:"addresses,omitempty"`
	AntiSpywareSecurityProfiles             *string `xml:"anti-spyware-security-profiles,omitempty"`
	AntivirusSecurityProfiles               *string `xml:"antivirus-security-profiles,omitempty"`
	ApplicationFilters                      *string `xml:"application-filters,omitempty"`
	ApplicationGroups                       *string `xml:"application-groups,omitempty"`
	Applications                            *string `xml:"applications,omitempty"`
	AuthenticationEnforcements              *string `xml:"authentication-enforcements,omitempty"`
	CustomDataPatterns                      *string `xml:"custom-data-patterns,omitempty"`
	CustomSpywareSignatures                 *string `xml:"custom-spyware-signatures,omitempty"`
	CustomUrlCategories                     *string `xml:"custom-url-categories,omitempty"`
	CustomVulnerabilitySignatures           *string `xml:"custom-vulnerability-signatures,omitempty"`
	DataFilteringSecurityProfiles           *string `xml:"data-filtering-security-profiles,omitempty"`
	DecryptionProfiles                      *string `xml:"decryption-profiles,omitempty"`
	Devices                                 *string `xml:"devices,omitempty"`
	DosProtectionSecurityProfiles           *string `xml:"dos-protection-security-profiles,omitempty"`
	DynamicUserGroups                       *string `xml:"dynamic-user-groups,omitempty"`
	ExternalDynamicLists                    *string `xml:"external-dynamic-lists,omitempty"`
	FileBlockingSecurityProfiles            *string `xml:"file-blocking-security-profiles,omitempty"`
	GlobalprotectHipObjects                 *string `xml:"globalprotect-hip-objects,omitempty"`
	GlobalprotectHipProfiles                *string `xml:"globalprotect-hip-profiles,omitempty"`
	GtpProtectionSecurityProfiles           *string `xml:"gtp-protection-security-profiles,omitempty"`
	LogForwardingProfiles                   *string `xml:"log-forwarding-profiles,omitempty"`
	PacketBrokerProfiles                    *string `xml:"packet-broker-profiles,omitempty"`
	Regions                                 *string `xml:"regions,omitempty"`
	Schedules                               *string `xml:"schedules,omitempty"`
	SctpProtectionSecurityProfiles          *string `xml:"sctp-protection-security-profiles,omitempty"`
	SdwanErrorCorrectionProfiles            *string `xml:"sdwan-error-correction-profiles,omitempty"`
	SdwanPathQualityProfiles                *string `xml:"sdwan-path-quality-profiles,omitempty"`
	SdwanSaasQualityProfiles                *string `xml:"sdwan-saas-quality-profiles,omitempty"`
	SdwanTrafficDistributionProfiles        *string `xml:"sdwan-traffic-distribution-profiles,omitempty"`
	SecurityProfileGroups                   *string `xml:"security-profile-groups,omitempty"`
	ServiceGroups                           *string `xml:"service-groups,omitempty"`
	Services                                *string `xml:"services,omitempty"`
	Tags                                    *string `xml:"tags,omitempty"`
	UrlFilteringSecurityProfiles            *string `xml:"url-filtering-security-profiles,omitempty"`
	VulnerabilityProtectionSecurityProfiles *string `xml:"vulnerability-protection-security-profiles,omitempty"`
	WildfireAnalysisSecurityProfiles        *string `xml:"wildfire-analysis-security-profiles,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysRestapiPoliciesXml struct {
	ApplicationOverrideRules   *string `xml:"application-override-rules,omitempty"`
	AuthenticationRules        *string `xml:"authentication-rules,omitempty"`
	DecryptionRules            *string `xml:"decryption-rules,omitempty"`
	DosRules                   *string `xml:"dos-rules,omitempty"`
	NatRules                   *string `xml:"nat-rules,omitempty"`
	NetworkPacketBrokerRules   *string `xml:"network-packet-broker-rules,omitempty"`
	PolicyBasedForwardingRules *string `xml:"policy-based-forwarding-rules,omitempty"`
	QosRules                   *string `xml:"qos-rules,omitempty"`
	SdwanRules                 *string `xml:"sdwan-rules,omitempty"`
	SecurityRules              *string `xml:"security-rules,omitempty"`
	TunnelInspectionRules      *string `xml:"tunnel-inspection-rules,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysRestapiSystemXml struct {
	Configuration *string `xml:"configuration,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiXml struct {
	Acc        *string                     `xml:"acc,omitempty"`
	Commit     *RoleVsysWebuiCommitXml     `xml:"commit,omitempty"`
	Dashboard  *string                     `xml:"dashboard,omitempty"`
	Device     *RoleVsysWebuiDeviceXml     `xml:"device,omitempty"`
	Monitor    *RoleVsysWebuiMonitorXml    `xml:"monitor,omitempty"`
	Network    *RoleVsysWebuiNetworkXml    `xml:"network,omitempty"`
	Objects    *RoleVsysWebuiObjectsXml    `xml:"objects,omitempty"`
	Operations *RoleVsysWebuiOperationsXml `xml:"operations,omitempty"`
	Policies   *RoleVsysWebuiPoliciesXml   `xml:"policies,omitempty"`
	Privacy    *RoleVsysWebuiPrivacyXml    `xml:"privacy,omitempty"`
	Save       *RoleVsysWebuiSaveXml       `xml:"save,omitempty"`
	Tasks      *string                     `xml:"tasks,omitempty"`
	Validate   *string                     `xml:"validate,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiCommitXml struct {
	CommitForOtherAdmins *string `xml:"commit-for-other-admins,omitempty"`
	VirtualSystems       *string `xml:"virtual-systems,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiDeviceXml struct {
	Administrators         *string                                      `xml:"administrators,omitempty"`
	AuthenticationProfile  *string                                      `xml:"authentication-profile,omitempty"`
	AuthenticationSequence *string                                      `xml:"authentication-sequence,omitempty"`
	BlockPages             *string                                      `xml:"block-pages,omitempty"`
	CertificateManagement  *RoleVsysWebuiDeviceCertificateManagementXml `xml:"certificate-management,omitempty"`
	DataRedistribution     *string                                      `xml:"data-redistribution,omitempty"`
	DeviceQuarantine       *string                                      `xml:"device-quarantine,omitempty"`
	DhcpSyslogServer       *string                                      `xml:"dhcp-syslog-server,omitempty"`
	LocalUserDatabase      *RoleVsysWebuiDeviceLocalUserDatabaseXml     `xml:"local-user-database,omitempty"`
	LogSettings            *RoleVsysWebuiDeviceLogSettingsXml           `xml:"log-settings,omitempty"`
	PolicyRecommendations  *RoleVsysWebuiDevicePolicyRecommendationsXml `xml:"policy-recommendations,omitempty"`
	ServerProfile          *RoleVsysWebuiDeviceServerProfileXml         `xml:"server-profile,omitempty"`
	Setup                  *RoleVsysWebuiDeviceSetupXml                 `xml:"setup,omitempty"`
	Troubleshooting        *string                                      `xml:"troubleshooting,omitempty"`
	UserIdentification     *string                                      `xml:"user-identification,omitempty"`
	VmInfoSource           *string                                      `xml:"vm-info-source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiDeviceCertificateManagementXml struct {
	CertificateProfile     *string `xml:"certificate-profile,omitempty"`
	Certificates           *string `xml:"certificates,omitempty"`
	OcspResponder          *string `xml:"ocsp-responder,omitempty"`
	Scep                   *string `xml:"scep,omitempty"`
	SshServiceProfile      *string `xml:"ssh-service-profile,omitempty"`
	SslDecryptionExclusion *string `xml:"ssl-decryption-exclusion,omitempty"`
	SslTlsServiceProfile   *string `xml:"ssl-tls-service-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiDeviceLocalUserDatabaseXml struct {
	UserGroups *string `xml:"user-groups,omitempty"`
	Users      *string `xml:"users,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiDeviceLogSettingsXml struct {
	Config        *string `xml:"config,omitempty"`
	Correlation   *string `xml:"correlation,omitempty"`
	Globalprotect *string `xml:"globalprotect,omitempty"`
	Hipmatch      *string `xml:"hipmatch,omitempty"`
	Iptag         *string `xml:"iptag,omitempty"`
	System        *string `xml:"system,omitempty"`
	UserId        *string `xml:"user-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiDevicePolicyRecommendationsXml struct {
	Iot  *string `xml:"iot,omitempty"`
	Saas *string `xml:"saas,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiDeviceServerProfileXml struct {
	Dns      *string `xml:"dns,omitempty"`
	Email    *string `xml:"email,omitempty"`
	Http     *string `xml:"http,omitempty"`
	Kerberos *string `xml:"kerberos,omitempty"`
	Ldap     *string `xml:"ldap,omitempty"`
	Mfa      *string `xml:"mfa,omitempty"`
	Netflow  *string `xml:"netflow,omitempty"`
	Radius   *string `xml:"radius,omitempty"`
	SamlIdp  *string `xml:"saml_idp,omitempty"`
	Scp      *string `xml:"scp,omitempty"`
	SnmpTrap *string `xml:"snmp-trap,omitempty"`
	Syslog   *string `xml:"syslog,omitempty"`
	Tacplus  *string `xml:"tacplus,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiDeviceSetupXml struct {
	ContentId  *string `xml:"content-id,omitempty"`
	Hsm        *string `xml:"hsm,omitempty"`
	Interfaces *string `xml:"interfaces,omitempty"`
	Management *string `xml:"management,omitempty"`
	Operations *string `xml:"operations,omitempty"`
	Services   *string `xml:"services,omitempty"`
	Session    *string `xml:"session,omitempty"`
	Telemetry  *string `xml:"telemetry,omitempty"`
	Wildfire   *string `xml:"wildfire,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiMonitorXml struct {
	AppScope                   *string                                            `xml:"app-scope,omitempty"`
	AutomatedCorrelationEngine *RoleVsysWebuiMonitorAutomatedCorrelationEngineXml `xml:"automated-correlation-engine,omitempty"`
	BlockIpList                *string                                            `xml:"block-ip-list,omitempty"`
	CustomReports              *RoleVsysWebuiMonitorCustomReportsXml              `xml:"custom-reports,omitempty"`
	ExternalLogs               *string                                            `xml:"external-logs,omitempty"`
	Logs                       *RoleVsysWebuiMonitorLogsXml                       `xml:"logs,omitempty"`
	PdfReports                 *RoleVsysWebuiMonitorPdfReportsXml                 `xml:"pdf-reports,omitempty"`
	SessionBrowser             *string                                            `xml:"session-browser,omitempty"`
	ViewCustomReports          *string                                            `xml:"view-custom-reports,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiMonitorAutomatedCorrelationEngineXml struct {
	CorrelatedEvents   *string `xml:"correlated-events,omitempty"`
	CorrelationObjects *string `xml:"correlation-objects,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiMonitorCustomReportsXml struct {
	ApplicationStatistics *string `xml:"application-statistics,omitempty"`
	Auth                  *string `xml:"auth,omitempty"`
	DataFilteringLog      *string `xml:"data-filtering-log,omitempty"`
	DecryptionLog         *string `xml:"decryption-log,omitempty"`
	DecryptionSummary     *string `xml:"decryption-summary,omitempty"`
	Globalprotect         *string `xml:"globalprotect,omitempty"`
	GtpLog                *string `xml:"gtp-log,omitempty"`
	GtpSummary            *string `xml:"gtp-summary,omitempty"`
	Hipmatch              *string `xml:"hipmatch,omitempty"`
	Iptag                 *string `xml:"iptag,omitempty"`
	SctpLog               *string `xml:"sctp-log,omitempty"`
	SctpSummary           *string `xml:"sctp-summary,omitempty"`
	ThreatLog             *string `xml:"threat-log,omitempty"`
	ThreatSummary         *string `xml:"threat-summary,omitempty"`
	TrafficLog            *string `xml:"traffic-log,omitempty"`
	TrafficSummary        *string `xml:"traffic-summary,omitempty"`
	TunnelLog             *string `xml:"tunnel-log,omitempty"`
	TunnelSummary         *string `xml:"tunnel-summary,omitempty"`
	UrlLog                *string `xml:"url-log,omitempty"`
	UrlSummary            *string `xml:"url-summary,omitempty"`
	Userid                *string `xml:"userid,omitempty"`
	WildfireLog           *string `xml:"wildfire-log,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiMonitorLogsXml struct {
	Authentication *string `xml:"authentication,omitempty"`
	DataFiltering  *string `xml:"data-filtering,omitempty"`
	Decryption     *string `xml:"decryption,omitempty"`
	Globalprotect  *string `xml:"globalprotect,omitempty"`
	Gtp            *string `xml:"gtp,omitempty"`
	Hipmatch       *string `xml:"hipmatch,omitempty"`
	Iptag          *string `xml:"iptag,omitempty"`
	Sctp           *string `xml:"sctp,omitempty"`
	Threat         *string `xml:"threat,omitempty"`
	Traffic        *string `xml:"traffic,omitempty"`
	Tunnel         *string `xml:"tunnel,omitempty"`
	Url            *string `xml:"url,omitempty"`
	Userid         *string `xml:"userid,omitempty"`
	Wildfire       *string `xml:"wildfire,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiMonitorPdfReportsXml struct {
	EmailScheduler             *string `xml:"email-scheduler,omitempty"`
	ManagePdfSummary           *string `xml:"manage-pdf-summary,omitempty"`
	PdfSummaryReports          *string `xml:"pdf-summary-reports,omitempty"`
	ReportGroups               *string `xml:"report-groups,omitempty"`
	SaasApplicationUsageReport *string `xml:"saas-application-usage-report,omitempty"`
	UserActivityReport         *string `xml:"user-activity-report,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiNetworkXml struct {
	GlobalProtect         *RoleVsysWebuiNetworkGlobalProtectXml `xml:"global-protect,omitempty"`
	SdwanInterfaceProfile *string                               `xml:"sdwan-interface-profile,omitempty"`
	Zones                 *string                               `xml:"zones,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiNetworkGlobalProtectXml struct {
	ClientlessAppGroups *string `xml:"clientless-app-groups,omitempty"`
	ClientlessApps      *string `xml:"clientless-apps,omitempty"`
	Gateways            *string `xml:"gateways,omitempty"`
	Mdm                 *string `xml:"mdm,omitempty"`
	Portals             *string `xml:"portals,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiObjectsXml struct {
	AddressGroups         *string                                  `xml:"address-groups,omitempty"`
	Addresses             *string                                  `xml:"addresses,omitempty"`
	ApplicationFilters    *string                                  `xml:"application-filters,omitempty"`
	ApplicationGroups     *string                                  `xml:"application-groups,omitempty"`
	Applications          *string                                  `xml:"applications,omitempty"`
	Authentication        *string                                  `xml:"authentication,omitempty"`
	CustomObjects         *RoleVsysWebuiObjectsCustomObjectsXml    `xml:"custom-objects,omitempty"`
	Decryption            *RoleVsysWebuiObjectsDecryptionXml       `xml:"decryption,omitempty"`
	Devices               *string                                  `xml:"devices,omitempty"`
	DynamicBlockLists     *string                                  `xml:"dynamic-block-lists,omitempty"`
	DynamicUserGroups     *string                                  `xml:"dynamic-user-groups,omitempty"`
	GlobalProtect         *RoleVsysWebuiObjectsGlobalProtectXml    `xml:"global-protect,omitempty"`
	LogForwarding         *string                                  `xml:"log-forwarding,omitempty"`
	PacketBrokerProfile   *string                                  `xml:"packet-broker-profile,omitempty"`
	Regions               *string                                  `xml:"regions,omitempty"`
	Schedules             *string                                  `xml:"schedules,omitempty"`
	Sdwan                 *RoleVsysWebuiObjectsSdwanXml            `xml:"sdwan,omitempty"`
	SecurityProfileGroups *string                                  `xml:"security-profile-groups,omitempty"`
	SecurityProfiles      *RoleVsysWebuiObjectsSecurityProfilesXml `xml:"security-profiles,omitempty"`
	ServiceGroups         *string                                  `xml:"service-groups,omitempty"`
	Services              *string                                  `xml:"services,omitempty"`
	Tags                  *string                                  `xml:"tags,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiObjectsCustomObjectsXml struct {
	DataPatterns  *string `xml:"data-patterns,omitempty"`
	Spyware       *string `xml:"spyware,omitempty"`
	UrlCategory   *string `xml:"url-category,omitempty"`
	Vulnerability *string `xml:"vulnerability,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiObjectsDecryptionXml struct {
	DecryptionProfile *string `xml:"decryption-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiObjectsGlobalProtectXml struct {
	HipObjects  *string `xml:"hip-objects,omitempty"`
	HipProfiles *string `xml:"hip-profiles,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiObjectsSdwanXml struct {
	SdwanDistProfile            *string `xml:"sdwan-dist-profile,omitempty"`
	SdwanErrorCorrectionProfile *string `xml:"sdwan-error-correction-profile,omitempty"`
	SdwanProfile                *string `xml:"sdwan-profile,omitempty"`
	SdwanSaasQualityProfile     *string `xml:"sdwan-saas-quality-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiObjectsSecurityProfilesXml struct {
	AntiSpyware             *string `xml:"anti-spyware,omitempty"`
	Antivirus               *string `xml:"antivirus,omitempty"`
	DataFiltering           *string `xml:"data-filtering,omitempty"`
	DosProtection           *string `xml:"dos-protection,omitempty"`
	FileBlocking            *string `xml:"file-blocking,omitempty"`
	GtpProtection           *string `xml:"gtp-protection,omitempty"`
	SctpProtection          *string `xml:"sctp-protection,omitempty"`
	UrlFiltering            *string `xml:"url-filtering,omitempty"`
	VulnerabilityProtection *string `xml:"vulnerability-protection,omitempty"`
	WildfireAnalysis        *string `xml:"wildfire-analysis,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiOperationsXml struct {
	DownloadCoreFiles       *string `xml:"download-core-files,omitempty"`
	DownloadPcapFiles       *string `xml:"download-pcap-files,omitempty"`
	GenerateStatsDumpFile   *string `xml:"generate-stats-dump-file,omitempty"`
	GenerateTechSupportFile *string `xml:"generate-tech-support-file,omitempty"`
	Reboot                  *string `xml:"reboot,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiPoliciesXml struct {
	ApplicationOverrideRulebase *string `xml:"application-override-rulebase,omitempty"`
	AuthenticationRulebase      *string `xml:"authentication-rulebase,omitempty"`
	DosRulebase                 *string `xml:"dos-rulebase,omitempty"`
	NatRulebase                 *string `xml:"nat-rulebase,omitempty"`
	NetworkPacketBrokerRulebase *string `xml:"network-packet-broker-rulebase,omitempty"`
	PbfRulebase                 *string `xml:"pbf-rulebase,omitempty"`
	QosRulebase                 *string `xml:"qos-rulebase,omitempty"`
	RuleHitCountReset           *string `xml:"rule-hit-count-reset,omitempty"`
	SdwanRulebase               *string `xml:"sdwan-rulebase,omitempty"`
	SecurityRulebase            *string `xml:"security-rulebase,omitempty"`
	SslDecryptionRulebase       *string `xml:"ssl-decryption-rulebase,omitempty"`
	TunnelInspectRulebase       *string `xml:"tunnel-inspect-rulebase,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiPrivacyXml struct {
	ShowFullIpAddresses           *string `xml:"show-full-ip-addresses,omitempty"`
	ShowUserNamesInLogsAndReports *string `xml:"show-user-names-in-logs-and-reports,omitempty"`
	ViewPcapFiles                 *string `xml:"view-pcap-files,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysWebuiSaveXml struct {
	ObjectLevelChanges *string `xml:"object-level-changes,omitempty"`
	PartialSave        *string `xml:"partial-save,omitempty"`
	SaveForOtherAdmins *string `xml:"save-for-other-admins,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoleVsysXmlapiXml struct {
	Commit *string `xml:"commit,omitempty"`
	Config *string `xml:"config,omitempty"`
	Export *string `xml:"export,omitempty"`
	Import *string `xml:"import,omitempty"`
	Iot    *string `xml:"iot,omitempty"`
	Log    *string `xml:"log,omitempty"`
	Op     *string `xml:"op,omitempty"`
	Report *string `xml:"report,omitempty"`
	UserId *string `xml:"user-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "role" || v == "Role" {
		return e.Role, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.Description = o.Description
	var nestedRole *RoleXml
	if o.Role != nil {
		nestedRole = &RoleXml{}
		if _, ok := o.Misc["Role"]; ok {
			nestedRole.Misc = o.Misc["Role"]
		}
		if o.Role.Device != nil {
			nestedRole.Device = &RoleDeviceXml{}
			if _, ok := o.Misc["RoleDevice"]; ok {
				nestedRole.Device.Misc = o.Misc["RoleDevice"]
			}
			if o.Role.Device.Webui != nil {
				nestedRole.Device.Webui = &RoleDeviceWebuiXml{}
				if _, ok := o.Misc["RoleDeviceWebui"]; ok {
					nestedRole.Device.Webui.Misc = o.Misc["RoleDeviceWebui"]
				}
				if o.Role.Device.Webui.Acc != nil {
					nestedRole.Device.Webui.Acc = o.Role.Device.Webui.Acc
				}
				if o.Role.Device.Webui.Dashboard != nil {
					nestedRole.Device.Webui.Dashboard = o.Role.Device.Webui.Dashboard
				}
				if o.Role.Device.Webui.Network != nil {
					nestedRole.Device.Webui.Network = &RoleDeviceWebuiNetworkXml{}
					if _, ok := o.Misc["RoleDeviceWebuiNetwork"]; ok {
						nestedRole.Device.Webui.Network.Misc = o.Misc["RoleDeviceWebuiNetwork"]
					}
					if o.Role.Device.Webui.Network.GreTunnels != nil {
						nestedRole.Device.Webui.Network.GreTunnels = o.Role.Device.Webui.Network.GreTunnels
					}
					if o.Role.Device.Webui.Network.Lldp != nil {
						nestedRole.Device.Webui.Network.Lldp = o.Role.Device.Webui.Network.Lldp
					}
					if o.Role.Device.Webui.Network.Qos != nil {
						nestedRole.Device.Webui.Network.Qos = o.Role.Device.Webui.Network.Qos
					}
					if o.Role.Device.Webui.Network.VirtualRouters != nil {
						nestedRole.Device.Webui.Network.VirtualRouters = o.Role.Device.Webui.Network.VirtualRouters
					}
					if o.Role.Device.Webui.Network.Dhcp != nil {
						nestedRole.Device.Webui.Network.Dhcp = o.Role.Device.Webui.Network.Dhcp
					}
					if o.Role.Device.Webui.Network.GlobalProtect != nil {
						nestedRole.Device.Webui.Network.GlobalProtect = &RoleDeviceWebuiNetworkGlobalProtectXml{}
						if _, ok := o.Misc["RoleDeviceWebuiNetworkGlobalProtect"]; ok {
							nestedRole.Device.Webui.Network.GlobalProtect.Misc = o.Misc["RoleDeviceWebuiNetworkGlobalProtect"]
						}
						if o.Role.Device.Webui.Network.GlobalProtect.ClientlessAppGroups != nil {
							nestedRole.Device.Webui.Network.GlobalProtect.ClientlessAppGroups = o.Role.Device.Webui.Network.GlobalProtect.ClientlessAppGroups
						}
						if o.Role.Device.Webui.Network.GlobalProtect.ClientlessApps != nil {
							nestedRole.Device.Webui.Network.GlobalProtect.ClientlessApps = o.Role.Device.Webui.Network.GlobalProtect.ClientlessApps
						}
						if o.Role.Device.Webui.Network.GlobalProtect.Gateways != nil {
							nestedRole.Device.Webui.Network.GlobalProtect.Gateways = o.Role.Device.Webui.Network.GlobalProtect.Gateways
						}
						if o.Role.Device.Webui.Network.GlobalProtect.Mdm != nil {
							nestedRole.Device.Webui.Network.GlobalProtect.Mdm = o.Role.Device.Webui.Network.GlobalProtect.Mdm
						}
						if o.Role.Device.Webui.Network.GlobalProtect.Portals != nil {
							nestedRole.Device.Webui.Network.GlobalProtect.Portals = o.Role.Device.Webui.Network.GlobalProtect.Portals
						}
					}
					if o.Role.Device.Webui.Network.Zones != nil {
						nestedRole.Device.Webui.Network.Zones = o.Role.Device.Webui.Network.Zones
					}
					if o.Role.Device.Webui.Network.DnsProxy != nil {
						nestedRole.Device.Webui.Network.DnsProxy = o.Role.Device.Webui.Network.DnsProxy
					}
					if o.Role.Device.Webui.Network.Interfaces != nil {
						nestedRole.Device.Webui.Network.Interfaces = o.Role.Device.Webui.Network.Interfaces
					}
					if o.Role.Device.Webui.Network.Routing != nil {
						nestedRole.Device.Webui.Network.Routing = &RoleDeviceWebuiNetworkRoutingXml{}
						if _, ok := o.Misc["RoleDeviceWebuiNetworkRouting"]; ok {
							nestedRole.Device.Webui.Network.Routing.Misc = o.Misc["RoleDeviceWebuiNetworkRouting"]
						}
						if o.Role.Device.Webui.Network.Routing.LogicalRouters != nil {
							nestedRole.Device.Webui.Network.Routing.LogicalRouters = o.Role.Device.Webui.Network.Routing.LogicalRouters
						}
						if o.Role.Device.Webui.Network.Routing.RoutingProfiles != nil {
							nestedRole.Device.Webui.Network.Routing.RoutingProfiles = &RoleDeviceWebuiNetworkRoutingRoutingProfilesXml{}
							if _, ok := o.Misc["RoleDeviceWebuiNetworkRoutingRoutingProfiles"]; ok {
								nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Misc = o.Misc["RoleDeviceWebuiNetworkRoutingRoutingProfiles"]
							}
							if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Bgp != nil {
								nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Bgp = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Bgp
							}
							if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Filters != nil {
								nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Filters = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Filters
							}
							if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Multicast != nil {
								nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Multicast = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Multicast
							}
							if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Ospf != nil {
								nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Ospf = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Ospf
							}
							if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Ospfv3 != nil {
								nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Ospfv3 = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Ospfv3
							}
							if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Ripv2 != nil {
								nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Ripv2 = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Ripv2
							}
							if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Bfd != nil {
								nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Bfd = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Bfd
							}
						}
					}
					if o.Role.Device.Webui.Network.VirtualWires != nil {
						nestedRole.Device.Webui.Network.VirtualWires = o.Role.Device.Webui.Network.VirtualWires
					}
					if o.Role.Device.Webui.Network.IpsecTunnels != nil {
						nestedRole.Device.Webui.Network.IpsecTunnels = o.Role.Device.Webui.Network.IpsecTunnels
					}
					if o.Role.Device.Webui.Network.NetworkProfiles != nil {
						nestedRole.Device.Webui.Network.NetworkProfiles = &RoleDeviceWebuiNetworkNetworkProfilesXml{}
						if _, ok := o.Misc["RoleDeviceWebuiNetworkNetworkProfiles"]; ok {
							nestedRole.Device.Webui.Network.NetworkProfiles.Misc = o.Misc["RoleDeviceWebuiNetworkNetworkProfiles"]
						}
						if o.Role.Device.Webui.Network.NetworkProfiles.TunnelMonitor != nil {
							nestedRole.Device.Webui.Network.NetworkProfiles.TunnelMonitor = o.Role.Device.Webui.Network.NetworkProfiles.TunnelMonitor
						}
						if o.Role.Device.Webui.Network.NetworkProfiles.ZoneProtection != nil {
							nestedRole.Device.Webui.Network.NetworkProfiles.ZoneProtection = o.Role.Device.Webui.Network.NetworkProfiles.ZoneProtection
						}
						if o.Role.Device.Webui.Network.NetworkProfiles.BfdProfile != nil {
							nestedRole.Device.Webui.Network.NetworkProfiles.BfdProfile = o.Role.Device.Webui.Network.NetworkProfiles.BfdProfile
						}
						if o.Role.Device.Webui.Network.NetworkProfiles.IkeCrypto != nil {
							nestedRole.Device.Webui.Network.NetworkProfiles.IkeCrypto = o.Role.Device.Webui.Network.NetworkProfiles.IkeCrypto
						}
						if o.Role.Device.Webui.Network.NetworkProfiles.QosProfile != nil {
							nestedRole.Device.Webui.Network.NetworkProfiles.QosProfile = o.Role.Device.Webui.Network.NetworkProfiles.QosProfile
						}
						if o.Role.Device.Webui.Network.NetworkProfiles.IpsecCrypto != nil {
							nestedRole.Device.Webui.Network.NetworkProfiles.IpsecCrypto = o.Role.Device.Webui.Network.NetworkProfiles.IpsecCrypto
						}
						if o.Role.Device.Webui.Network.NetworkProfiles.LldpProfile != nil {
							nestedRole.Device.Webui.Network.NetworkProfiles.LldpProfile = o.Role.Device.Webui.Network.NetworkProfiles.LldpProfile
						}
						if o.Role.Device.Webui.Network.NetworkProfiles.GpAppIpsecCrypto != nil {
							nestedRole.Device.Webui.Network.NetworkProfiles.GpAppIpsecCrypto = o.Role.Device.Webui.Network.NetworkProfiles.GpAppIpsecCrypto
						}
						if o.Role.Device.Webui.Network.NetworkProfiles.IkeGateways != nil {
							nestedRole.Device.Webui.Network.NetworkProfiles.IkeGateways = o.Role.Device.Webui.Network.NetworkProfiles.IkeGateways
						}
						if o.Role.Device.Webui.Network.NetworkProfiles.InterfaceMgmt != nil {
							nestedRole.Device.Webui.Network.NetworkProfiles.InterfaceMgmt = o.Role.Device.Webui.Network.NetworkProfiles.InterfaceMgmt
						}
					}
					if o.Role.Device.Webui.Network.SecureWebGateway != nil {
						nestedRole.Device.Webui.Network.SecureWebGateway = o.Role.Device.Webui.Network.SecureWebGateway
					}
					if o.Role.Device.Webui.Network.SdwanInterfaceProfile != nil {
						nestedRole.Device.Webui.Network.SdwanInterfaceProfile = o.Role.Device.Webui.Network.SdwanInterfaceProfile
					}
					if o.Role.Device.Webui.Network.Vlans != nil {
						nestedRole.Device.Webui.Network.Vlans = o.Role.Device.Webui.Network.Vlans
					}
				}
				if o.Role.Device.Webui.Privacy != nil {
					nestedRole.Device.Webui.Privacy = &RoleDeviceWebuiPrivacyXml{}
					if _, ok := o.Misc["RoleDeviceWebuiPrivacy"]; ok {
						nestedRole.Device.Webui.Privacy.Misc = o.Misc["RoleDeviceWebuiPrivacy"]
					}
					if o.Role.Device.Webui.Privacy.ViewPcapFiles != nil {
						nestedRole.Device.Webui.Privacy.ViewPcapFiles = o.Role.Device.Webui.Privacy.ViewPcapFiles
					}
					if o.Role.Device.Webui.Privacy.ShowFullIpAddresses != nil {
						nestedRole.Device.Webui.Privacy.ShowFullIpAddresses = o.Role.Device.Webui.Privacy.ShowFullIpAddresses
					}
					if o.Role.Device.Webui.Privacy.ShowUserNamesInLogsAndReports != nil {
						nestedRole.Device.Webui.Privacy.ShowUserNamesInLogsAndReports = o.Role.Device.Webui.Privacy.ShowUserNamesInLogsAndReports
					}
				}
				if o.Role.Device.Webui.Tasks != nil {
					nestedRole.Device.Webui.Tasks = o.Role.Device.Webui.Tasks
				}
				if o.Role.Device.Webui.Commit != nil {
					nestedRole.Device.Webui.Commit = &RoleDeviceWebuiCommitXml{}
					if _, ok := o.Misc["RoleDeviceWebuiCommit"]; ok {
						nestedRole.Device.Webui.Commit.Misc = o.Misc["RoleDeviceWebuiCommit"]
					}
					if o.Role.Device.Webui.Commit.CommitForOtherAdmins != nil {
						nestedRole.Device.Webui.Commit.CommitForOtherAdmins = o.Role.Device.Webui.Commit.CommitForOtherAdmins
					}
					if o.Role.Device.Webui.Commit.Device != nil {
						nestedRole.Device.Webui.Commit.Device = o.Role.Device.Webui.Commit.Device
					}
					if o.Role.Device.Webui.Commit.ObjectLevelChanges != nil {
						nestedRole.Device.Webui.Commit.ObjectLevelChanges = o.Role.Device.Webui.Commit.ObjectLevelChanges
					}
				}
				if o.Role.Device.Webui.Device != nil {
					nestedRole.Device.Webui.Device = &RoleDeviceWebuiDeviceXml{}
					if _, ok := o.Misc["RoleDeviceWebuiDevice"]; ok {
						nestedRole.Device.Webui.Device.Misc = o.Misc["RoleDeviceWebuiDevice"]
					}
					if o.Role.Device.Webui.Device.DataRedistribution != nil {
						nestedRole.Device.Webui.Device.DataRedistribution = o.Role.Device.Webui.Device.DataRedistribution
					}
					if o.Role.Device.Webui.Device.HighAvailability != nil {
						nestedRole.Device.Webui.Device.HighAvailability = o.Role.Device.Webui.Device.HighAvailability
					}
					if o.Role.Device.Webui.Device.LocalUserDatabase != nil {
						nestedRole.Device.Webui.Device.LocalUserDatabase = &RoleDeviceWebuiDeviceLocalUserDatabaseXml{}
						if _, ok := o.Misc["RoleDeviceWebuiDeviceLocalUserDatabase"]; ok {
							nestedRole.Device.Webui.Device.LocalUserDatabase.Misc = o.Misc["RoleDeviceWebuiDeviceLocalUserDatabase"]
						}
						if o.Role.Device.Webui.Device.LocalUserDatabase.UserGroups != nil {
							nestedRole.Device.Webui.Device.LocalUserDatabase.UserGroups = o.Role.Device.Webui.Device.LocalUserDatabase.UserGroups
						}
						if o.Role.Device.Webui.Device.LocalUserDatabase.Users != nil {
							nestedRole.Device.Webui.Device.LocalUserDatabase.Users = o.Role.Device.Webui.Device.LocalUserDatabase.Users
						}
					}
					if o.Role.Device.Webui.Device.MasterKey != nil {
						nestedRole.Device.Webui.Device.MasterKey = o.Role.Device.Webui.Device.MasterKey
					}
					if o.Role.Device.Webui.Device.Plugins != nil {
						nestedRole.Device.Webui.Device.Plugins = o.Role.Device.Webui.Device.Plugins
					}
					if o.Role.Device.Webui.Device.PolicyRecommendations != nil {
						nestedRole.Device.Webui.Device.PolicyRecommendations = &RoleDeviceWebuiDevicePolicyRecommendationsXml{}
						if _, ok := o.Misc["RoleDeviceWebuiDevicePolicyRecommendations"]; ok {
							nestedRole.Device.Webui.Device.PolicyRecommendations.Misc = o.Misc["RoleDeviceWebuiDevicePolicyRecommendations"]
						}
						if o.Role.Device.Webui.Device.PolicyRecommendations.Iot != nil {
							nestedRole.Device.Webui.Device.PolicyRecommendations.Iot = o.Role.Device.Webui.Device.PolicyRecommendations.Iot
						}
						if o.Role.Device.Webui.Device.PolicyRecommendations.Saas != nil {
							nestedRole.Device.Webui.Device.PolicyRecommendations.Saas = o.Role.Device.Webui.Device.PolicyRecommendations.Saas
						}
					}
					if o.Role.Device.Webui.Device.ServerProfile != nil {
						nestedRole.Device.Webui.Device.ServerProfile = &RoleDeviceWebuiDeviceServerProfileXml{}
						if _, ok := o.Misc["RoleDeviceWebuiDeviceServerProfile"]; ok {
							nestedRole.Device.Webui.Device.ServerProfile.Misc = o.Misc["RoleDeviceWebuiDeviceServerProfile"]
						}
						if o.Role.Device.Webui.Device.ServerProfile.Dns != nil {
							nestedRole.Device.Webui.Device.ServerProfile.Dns = o.Role.Device.Webui.Device.ServerProfile.Dns
						}
						if o.Role.Device.Webui.Device.ServerProfile.Email != nil {
							nestedRole.Device.Webui.Device.ServerProfile.Email = o.Role.Device.Webui.Device.ServerProfile.Email
						}
						if o.Role.Device.Webui.Device.ServerProfile.Netflow != nil {
							nestedRole.Device.Webui.Device.ServerProfile.Netflow = o.Role.Device.Webui.Device.ServerProfile.Netflow
						}
						if o.Role.Device.Webui.Device.ServerProfile.SamlIdp != nil {
							nestedRole.Device.Webui.Device.ServerProfile.SamlIdp = o.Role.Device.Webui.Device.ServerProfile.SamlIdp
						}
						if o.Role.Device.Webui.Device.ServerProfile.Tacplus != nil {
							nestedRole.Device.Webui.Device.ServerProfile.Tacplus = o.Role.Device.Webui.Device.ServerProfile.Tacplus
						}
						if o.Role.Device.Webui.Device.ServerProfile.SnmpTrap != nil {
							nestedRole.Device.Webui.Device.ServerProfile.SnmpTrap = o.Role.Device.Webui.Device.ServerProfile.SnmpTrap
						}
						if o.Role.Device.Webui.Device.ServerProfile.Syslog != nil {
							nestedRole.Device.Webui.Device.ServerProfile.Syslog = o.Role.Device.Webui.Device.ServerProfile.Syslog
						}
						if o.Role.Device.Webui.Device.ServerProfile.Http != nil {
							nestedRole.Device.Webui.Device.ServerProfile.Http = o.Role.Device.Webui.Device.ServerProfile.Http
						}
						if o.Role.Device.Webui.Device.ServerProfile.Kerberos != nil {
							nestedRole.Device.Webui.Device.ServerProfile.Kerberos = o.Role.Device.Webui.Device.ServerProfile.Kerberos
						}
						if o.Role.Device.Webui.Device.ServerProfile.Ldap != nil {
							nestedRole.Device.Webui.Device.ServerProfile.Ldap = o.Role.Device.Webui.Device.ServerProfile.Ldap
						}
						if o.Role.Device.Webui.Device.ServerProfile.Mfa != nil {
							nestedRole.Device.Webui.Device.ServerProfile.Mfa = o.Role.Device.Webui.Device.ServerProfile.Mfa
						}
						if o.Role.Device.Webui.Device.ServerProfile.Radius != nil {
							nestedRole.Device.Webui.Device.ServerProfile.Radius = o.Role.Device.Webui.Device.ServerProfile.Radius
						}
						if o.Role.Device.Webui.Device.ServerProfile.Scp != nil {
							nestedRole.Device.Webui.Device.ServerProfile.Scp = o.Role.Device.Webui.Device.ServerProfile.Scp
						}
					}
					if o.Role.Device.Webui.Device.ConfigAudit != nil {
						nestedRole.Device.Webui.Device.ConfigAudit = o.Role.Device.Webui.Device.ConfigAudit
					}
					if o.Role.Device.Webui.Device.Support != nil {
						nestedRole.Device.Webui.Device.Support = o.Role.Device.Webui.Device.Support
					}
					if o.Role.Device.Webui.Device.DynamicUpdates != nil {
						nestedRole.Device.Webui.Device.DynamicUpdates = o.Role.Device.Webui.Device.DynamicUpdates
					}
					if o.Role.Device.Webui.Device.LogSettings != nil {
						nestedRole.Device.Webui.Device.LogSettings = &RoleDeviceWebuiDeviceLogSettingsXml{}
						if _, ok := o.Misc["RoleDeviceWebuiDeviceLogSettings"]; ok {
							nestedRole.Device.Webui.Device.LogSettings.Misc = o.Misc["RoleDeviceWebuiDeviceLogSettings"]
						}
						if o.Role.Device.Webui.Device.LogSettings.Correlation != nil {
							nestedRole.Device.Webui.Device.LogSettings.Correlation = o.Role.Device.Webui.Device.LogSettings.Correlation
						}
						if o.Role.Device.Webui.Device.LogSettings.Globalprotect != nil {
							nestedRole.Device.Webui.Device.LogSettings.Globalprotect = o.Role.Device.Webui.Device.LogSettings.Globalprotect
						}
						if o.Role.Device.Webui.Device.LogSettings.Iptag != nil {
							nestedRole.Device.Webui.Device.LogSettings.Iptag = o.Role.Device.Webui.Device.LogSettings.Iptag
						}
						if o.Role.Device.Webui.Device.LogSettings.ManageLog != nil {
							nestedRole.Device.Webui.Device.LogSettings.ManageLog = o.Role.Device.Webui.Device.LogSettings.ManageLog
						}
						if o.Role.Device.Webui.Device.LogSettings.UserId != nil {
							nestedRole.Device.Webui.Device.LogSettings.UserId = o.Role.Device.Webui.Device.LogSettings.UserId
						}
						if o.Role.Device.Webui.Device.LogSettings.CcAlarm != nil {
							nestedRole.Device.Webui.Device.LogSettings.CcAlarm = o.Role.Device.Webui.Device.LogSettings.CcAlarm
						}
						if o.Role.Device.Webui.Device.LogSettings.Config != nil {
							nestedRole.Device.Webui.Device.LogSettings.Config = o.Role.Device.Webui.Device.LogSettings.Config
						}
						if o.Role.Device.Webui.Device.LogSettings.Hipmatch != nil {
							nestedRole.Device.Webui.Device.LogSettings.Hipmatch = o.Role.Device.Webui.Device.LogSettings.Hipmatch
						}
						if o.Role.Device.Webui.Device.LogSettings.System != nil {
							nestedRole.Device.Webui.Device.LogSettings.System = o.Role.Device.Webui.Device.LogSettings.System
						}
					}
					if o.Role.Device.Webui.Device.VmInfoSource != nil {
						nestedRole.Device.Webui.Device.VmInfoSource = o.Role.Device.Webui.Device.VmInfoSource
					}
					if o.Role.Device.Webui.Device.AdminRoles != nil {
						nestedRole.Device.Webui.Device.AdminRoles = o.Role.Device.Webui.Device.AdminRoles
					}
					if o.Role.Device.Webui.Device.Licenses != nil {
						nestedRole.Device.Webui.Device.Licenses = o.Role.Device.Webui.Device.Licenses
					}
					if o.Role.Device.Webui.Device.VirtualSystems != nil {
						nestedRole.Device.Webui.Device.VirtualSystems = o.Role.Device.Webui.Device.VirtualSystems
					}
					if o.Role.Device.Webui.Device.AuthenticationProfile != nil {
						nestedRole.Device.Webui.Device.AuthenticationProfile = o.Role.Device.Webui.Device.AuthenticationProfile
					}
					if o.Role.Device.Webui.Device.GlobalProtectClient != nil {
						nestedRole.Device.Webui.Device.GlobalProtectClient = o.Role.Device.Webui.Device.GlobalProtectClient
					}
					if o.Role.Device.Webui.Device.BlockPages != nil {
						nestedRole.Device.Webui.Device.BlockPages = o.Role.Device.Webui.Device.BlockPages
					}
					if o.Role.Device.Webui.Device.Administrators != nil {
						nestedRole.Device.Webui.Device.Administrators = o.Role.Device.Webui.Device.Administrators
					}
					if o.Role.Device.Webui.Device.AuthenticationSequence != nil {
						nestedRole.Device.Webui.Device.AuthenticationSequence = o.Role.Device.Webui.Device.AuthenticationSequence
					}
					if o.Role.Device.Webui.Device.AccessDomain != nil {
						nestedRole.Device.Webui.Device.AccessDomain = o.Role.Device.Webui.Device.AccessDomain
					}
					if o.Role.Device.Webui.Device.Setup != nil {
						nestedRole.Device.Webui.Device.Setup = &RoleDeviceWebuiDeviceSetupXml{}
						if _, ok := o.Misc["RoleDeviceWebuiDeviceSetup"]; ok {
							nestedRole.Device.Webui.Device.Setup.Misc = o.Misc["RoleDeviceWebuiDeviceSetup"]
						}
						if o.Role.Device.Webui.Device.Setup.Wildfire != nil {
							nestedRole.Device.Webui.Device.Setup.Wildfire = o.Role.Device.Webui.Device.Setup.Wildfire
						}
						if o.Role.Device.Webui.Device.Setup.ContentId != nil {
							nestedRole.Device.Webui.Device.Setup.ContentId = o.Role.Device.Webui.Device.Setup.ContentId
						}
						if o.Role.Device.Webui.Device.Setup.Management != nil {
							nestedRole.Device.Webui.Device.Setup.Management = o.Role.Device.Webui.Device.Setup.Management
						}
						if o.Role.Device.Webui.Device.Setup.Services != nil {
							nestedRole.Device.Webui.Device.Setup.Services = o.Role.Device.Webui.Device.Setup.Services
						}
						if o.Role.Device.Webui.Device.Setup.Session != nil {
							nestedRole.Device.Webui.Device.Setup.Session = o.Role.Device.Webui.Device.Setup.Session
						}
						if o.Role.Device.Webui.Device.Setup.Hsm != nil {
							nestedRole.Device.Webui.Device.Setup.Hsm = o.Role.Device.Webui.Device.Setup.Hsm
						}
						if o.Role.Device.Webui.Device.Setup.Interfaces != nil {
							nestedRole.Device.Webui.Device.Setup.Interfaces = o.Role.Device.Webui.Device.Setup.Interfaces
						}
						if o.Role.Device.Webui.Device.Setup.Operations != nil {
							nestedRole.Device.Webui.Device.Setup.Operations = o.Role.Device.Webui.Device.Setup.Operations
						}
						if o.Role.Device.Webui.Device.Setup.Telemetry != nil {
							nestedRole.Device.Webui.Device.Setup.Telemetry = o.Role.Device.Webui.Device.Setup.Telemetry
						}
					}
					if o.Role.Device.Webui.Device.ScheduledLogExport != nil {
						nestedRole.Device.Webui.Device.ScheduledLogExport = o.Role.Device.Webui.Device.ScheduledLogExport
					}
					if o.Role.Device.Webui.Device.Software != nil {
						nestedRole.Device.Webui.Device.Software = o.Role.Device.Webui.Device.Software
					}
					if o.Role.Device.Webui.Device.Troubleshooting != nil {
						nestedRole.Device.Webui.Device.Troubleshooting = o.Role.Device.Webui.Device.Troubleshooting
					}
					if o.Role.Device.Webui.Device.UserIdentification != nil {
						nestedRole.Device.Webui.Device.UserIdentification = o.Role.Device.Webui.Device.UserIdentification
					}
					if o.Role.Device.Webui.Device.DhcpSyslogServer != nil {
						nestedRole.Device.Webui.Device.DhcpSyslogServer = o.Role.Device.Webui.Device.DhcpSyslogServer
					}
					if o.Role.Device.Webui.Device.CertificateManagement != nil {
						nestedRole.Device.Webui.Device.CertificateManagement = &RoleDeviceWebuiDeviceCertificateManagementXml{}
						if _, ok := o.Misc["RoleDeviceWebuiDeviceCertificateManagement"]; ok {
							nestedRole.Device.Webui.Device.CertificateManagement.Misc = o.Misc["RoleDeviceWebuiDeviceCertificateManagement"]
						}
						if o.Role.Device.Webui.Device.CertificateManagement.Scep != nil {
							nestedRole.Device.Webui.Device.CertificateManagement.Scep = o.Role.Device.Webui.Device.CertificateManagement.Scep
						}
						if o.Role.Device.Webui.Device.CertificateManagement.SshServiceProfile != nil {
							nestedRole.Device.Webui.Device.CertificateManagement.SshServiceProfile = o.Role.Device.Webui.Device.CertificateManagement.SshServiceProfile
						}
						if o.Role.Device.Webui.Device.CertificateManagement.SslDecryptionExclusion != nil {
							nestedRole.Device.Webui.Device.CertificateManagement.SslDecryptionExclusion = o.Role.Device.Webui.Device.CertificateManagement.SslDecryptionExclusion
						}
						if o.Role.Device.Webui.Device.CertificateManagement.SslTlsServiceProfile != nil {
							nestedRole.Device.Webui.Device.CertificateManagement.SslTlsServiceProfile = o.Role.Device.Webui.Device.CertificateManagement.SslTlsServiceProfile
						}
						if o.Role.Device.Webui.Device.CertificateManagement.CertificateProfile != nil {
							nestedRole.Device.Webui.Device.CertificateManagement.CertificateProfile = o.Role.Device.Webui.Device.CertificateManagement.CertificateProfile
						}
						if o.Role.Device.Webui.Device.CertificateManagement.Certificates != nil {
							nestedRole.Device.Webui.Device.CertificateManagement.Certificates = o.Role.Device.Webui.Device.CertificateManagement.Certificates
						}
						if o.Role.Device.Webui.Device.CertificateManagement.OcspResponder != nil {
							nestedRole.Device.Webui.Device.CertificateManagement.OcspResponder = o.Role.Device.Webui.Device.CertificateManagement.OcspResponder
						}
					}
					if o.Role.Device.Webui.Device.LogFwdCard != nil {
						nestedRole.Device.Webui.Device.LogFwdCard = o.Role.Device.Webui.Device.LogFwdCard
					}
					if o.Role.Device.Webui.Device.SharedGateways != nil {
						nestedRole.Device.Webui.Device.SharedGateways = o.Role.Device.Webui.Device.SharedGateways
					}
					if o.Role.Device.Webui.Device.DeviceQuarantine != nil {
						nestedRole.Device.Webui.Device.DeviceQuarantine = o.Role.Device.Webui.Device.DeviceQuarantine
					}
				}
				if o.Role.Device.Webui.Monitor != nil {
					nestedRole.Device.Webui.Monitor = &RoleDeviceWebuiMonitorXml{}
					if _, ok := o.Misc["RoleDeviceWebuiMonitor"]; ok {
						nestedRole.Device.Webui.Monitor.Misc = o.Misc["RoleDeviceWebuiMonitor"]
					}
					if o.Role.Device.Webui.Monitor.SctpReports != nil {
						nestedRole.Device.Webui.Monitor.SctpReports = o.Role.Device.Webui.Monitor.SctpReports
					}
					if o.Role.Device.Webui.Monitor.SessionBrowser != nil {
						nestedRole.Device.Webui.Monitor.SessionBrowser = o.Role.Device.Webui.Monitor.SessionBrowser
					}
					if o.Role.Device.Webui.Monitor.AutomatedCorrelationEngine != nil {
						nestedRole.Device.Webui.Monitor.AutomatedCorrelationEngine = &RoleDeviceWebuiMonitorAutomatedCorrelationEngineXml{}
						if _, ok := o.Misc["RoleDeviceWebuiMonitorAutomatedCorrelationEngine"]; ok {
							nestedRole.Device.Webui.Monitor.AutomatedCorrelationEngine.Misc = o.Misc["RoleDeviceWebuiMonitorAutomatedCorrelationEngine"]
						}
						if o.Role.Device.Webui.Monitor.AutomatedCorrelationEngine.CorrelatedEvents != nil {
							nestedRole.Device.Webui.Monitor.AutomatedCorrelationEngine.CorrelatedEvents = o.Role.Device.Webui.Monitor.AutomatedCorrelationEngine.CorrelatedEvents
						}
						if o.Role.Device.Webui.Monitor.AutomatedCorrelationEngine.CorrelationObjects != nil {
							nestedRole.Device.Webui.Monitor.AutomatedCorrelationEngine.CorrelationObjects = o.Role.Device.Webui.Monitor.AutomatedCorrelationEngine.CorrelationObjects
						}
					}
					if o.Role.Device.Webui.Monitor.BlockIpList != nil {
						nestedRole.Device.Webui.Monitor.BlockIpList = o.Role.Device.Webui.Monitor.BlockIpList
					}
					if o.Role.Device.Webui.Monitor.PdfReports != nil {
						nestedRole.Device.Webui.Monitor.PdfReports = &RoleDeviceWebuiMonitorPdfReportsXml{}
						if _, ok := o.Misc["RoleDeviceWebuiMonitorPdfReports"]; ok {
							nestedRole.Device.Webui.Monitor.PdfReports.Misc = o.Misc["RoleDeviceWebuiMonitorPdfReports"]
						}
						if o.Role.Device.Webui.Monitor.PdfReports.EmailScheduler != nil {
							nestedRole.Device.Webui.Monitor.PdfReports.EmailScheduler = o.Role.Device.Webui.Monitor.PdfReports.EmailScheduler
						}
						if o.Role.Device.Webui.Monitor.PdfReports.ManagePdfSummary != nil {
							nestedRole.Device.Webui.Monitor.PdfReports.ManagePdfSummary = o.Role.Device.Webui.Monitor.PdfReports.ManagePdfSummary
						}
						if o.Role.Device.Webui.Monitor.PdfReports.PdfSummaryReports != nil {
							nestedRole.Device.Webui.Monitor.PdfReports.PdfSummaryReports = o.Role.Device.Webui.Monitor.PdfReports.PdfSummaryReports
						}
						if o.Role.Device.Webui.Monitor.PdfReports.ReportGroups != nil {
							nestedRole.Device.Webui.Monitor.PdfReports.ReportGroups = o.Role.Device.Webui.Monitor.PdfReports.ReportGroups
						}
						if o.Role.Device.Webui.Monitor.PdfReports.SaasApplicationUsageReport != nil {
							nestedRole.Device.Webui.Monitor.PdfReports.SaasApplicationUsageReport = o.Role.Device.Webui.Monitor.PdfReports.SaasApplicationUsageReport
						}
						if o.Role.Device.Webui.Monitor.PdfReports.UserActivityReport != nil {
							nestedRole.Device.Webui.Monitor.PdfReports.UserActivityReport = o.Role.Device.Webui.Monitor.PdfReports.UserActivityReport
						}
					}
					if o.Role.Device.Webui.Monitor.CustomReports != nil {
						nestedRole.Device.Webui.Monitor.CustomReports = &RoleDeviceWebuiMonitorCustomReportsXml{}
						if _, ok := o.Misc["RoleDeviceWebuiMonitorCustomReports"]; ok {
							nestedRole.Device.Webui.Monitor.CustomReports.Misc = o.Misc["RoleDeviceWebuiMonitorCustomReports"]
						}
						if o.Role.Device.Webui.Monitor.CustomReports.UrlLog != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.UrlLog = o.Role.Device.Webui.Monitor.CustomReports.UrlLog
						}
						if o.Role.Device.Webui.Monitor.CustomReports.Auth != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.Auth = o.Role.Device.Webui.Monitor.CustomReports.Auth
						}
						if o.Role.Device.Webui.Monitor.CustomReports.DataFilteringLog != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.DataFilteringLog = o.Role.Device.Webui.Monitor.CustomReports.DataFilteringLog
						}
						if o.Role.Device.Webui.Monitor.CustomReports.Globalprotect != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.Globalprotect = o.Role.Device.Webui.Monitor.CustomReports.Globalprotect
						}
						if o.Role.Device.Webui.Monitor.CustomReports.GtpSummary != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.GtpSummary = o.Role.Device.Webui.Monitor.CustomReports.GtpSummary
						}
						if o.Role.Device.Webui.Monitor.CustomReports.SctpSummary != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.SctpSummary = o.Role.Device.Webui.Monitor.CustomReports.SctpSummary
						}
						if o.Role.Device.Webui.Monitor.CustomReports.TrafficSummary != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.TrafficSummary = o.Role.Device.Webui.Monitor.CustomReports.TrafficSummary
						}
						if o.Role.Device.Webui.Monitor.CustomReports.TunnelSummary != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.TunnelSummary = o.Role.Device.Webui.Monitor.CustomReports.TunnelSummary
						}
						if o.Role.Device.Webui.Monitor.CustomReports.DecryptionLog != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.DecryptionLog = o.Role.Device.Webui.Monitor.CustomReports.DecryptionLog
						}
						if o.Role.Device.Webui.Monitor.CustomReports.DecryptionSummary != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.DecryptionSummary = o.Role.Device.Webui.Monitor.CustomReports.DecryptionSummary
						}
						if o.Role.Device.Webui.Monitor.CustomReports.Iptag != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.Iptag = o.Role.Device.Webui.Monitor.CustomReports.Iptag
						}
						if o.Role.Device.Webui.Monitor.CustomReports.SctpLog != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.SctpLog = o.Role.Device.Webui.Monitor.CustomReports.SctpLog
						}
						if o.Role.Device.Webui.Monitor.CustomReports.Userid != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.Userid = o.Role.Device.Webui.Monitor.CustomReports.Userid
						}
						if o.Role.Device.Webui.Monitor.CustomReports.GtpLog != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.GtpLog = o.Role.Device.Webui.Monitor.CustomReports.GtpLog
						}
						if o.Role.Device.Webui.Monitor.CustomReports.ThreatSummary != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.ThreatSummary = o.Role.Device.Webui.Monitor.CustomReports.ThreatSummary
						}
						if o.Role.Device.Webui.Monitor.CustomReports.UrlSummary != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.UrlSummary = o.Role.Device.Webui.Monitor.CustomReports.UrlSummary
						}
						if o.Role.Device.Webui.Monitor.CustomReports.WildfireLog != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.WildfireLog = o.Role.Device.Webui.Monitor.CustomReports.WildfireLog
						}
						if o.Role.Device.Webui.Monitor.CustomReports.ApplicationStatistics != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.ApplicationStatistics = o.Role.Device.Webui.Monitor.CustomReports.ApplicationStatistics
						}
						if o.Role.Device.Webui.Monitor.CustomReports.Hipmatch != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.Hipmatch = o.Role.Device.Webui.Monitor.CustomReports.Hipmatch
						}
						if o.Role.Device.Webui.Monitor.CustomReports.ThreatLog != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.ThreatLog = o.Role.Device.Webui.Monitor.CustomReports.ThreatLog
						}
						if o.Role.Device.Webui.Monitor.CustomReports.TrafficLog != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.TrafficLog = o.Role.Device.Webui.Monitor.CustomReports.TrafficLog
						}
						if o.Role.Device.Webui.Monitor.CustomReports.TunnelLog != nil {
							nestedRole.Device.Webui.Monitor.CustomReports.TunnelLog = o.Role.Device.Webui.Monitor.CustomReports.TunnelLog
						}
					}
					if o.Role.Device.Webui.Monitor.ThreatReports != nil {
						nestedRole.Device.Webui.Monitor.ThreatReports = o.Role.Device.Webui.Monitor.ThreatReports
					}
					if o.Role.Device.Webui.Monitor.ViewCustomReports != nil {
						nestedRole.Device.Webui.Monitor.ViewCustomReports = o.Role.Device.Webui.Monitor.ViewCustomReports
					}
					if o.Role.Device.Webui.Monitor.AppScope != nil {
						nestedRole.Device.Webui.Monitor.AppScope = o.Role.Device.Webui.Monitor.AppScope
					}
					if o.Role.Device.Webui.Monitor.ApplicationReports != nil {
						nestedRole.Device.Webui.Monitor.ApplicationReports = o.Role.Device.Webui.Monitor.ApplicationReports
					}
					if o.Role.Device.Webui.Monitor.GtpReports != nil {
						nestedRole.Device.Webui.Monitor.GtpReports = o.Role.Device.Webui.Monitor.GtpReports
					}
					if o.Role.Device.Webui.Monitor.PacketCapture != nil {
						nestedRole.Device.Webui.Monitor.PacketCapture = o.Role.Device.Webui.Monitor.PacketCapture
					}
					if o.Role.Device.Webui.Monitor.TrafficReports != nil {
						nestedRole.Device.Webui.Monitor.TrafficReports = o.Role.Device.Webui.Monitor.TrafficReports
					}
					if o.Role.Device.Webui.Monitor.UrlFilteringReports != nil {
						nestedRole.Device.Webui.Monitor.UrlFilteringReports = o.Role.Device.Webui.Monitor.UrlFilteringReports
					}
					if o.Role.Device.Webui.Monitor.Botnet != nil {
						nestedRole.Device.Webui.Monitor.Botnet = o.Role.Device.Webui.Monitor.Botnet
					}
					if o.Role.Device.Webui.Monitor.ExternalLogs != nil {
						nestedRole.Device.Webui.Monitor.ExternalLogs = o.Role.Device.Webui.Monitor.ExternalLogs
					}
					if o.Role.Device.Webui.Monitor.Logs != nil {
						nestedRole.Device.Webui.Monitor.Logs = &RoleDeviceWebuiMonitorLogsXml{}
						if _, ok := o.Misc["RoleDeviceWebuiMonitorLogs"]; ok {
							nestedRole.Device.Webui.Monitor.Logs.Misc = o.Misc["RoleDeviceWebuiMonitorLogs"]
						}
						if o.Role.Device.Webui.Monitor.Logs.Decryption != nil {
							nestedRole.Device.Webui.Monitor.Logs.Decryption = o.Role.Device.Webui.Monitor.Logs.Decryption
						}
						if o.Role.Device.Webui.Monitor.Logs.Gtp != nil {
							nestedRole.Device.Webui.Monitor.Logs.Gtp = o.Role.Device.Webui.Monitor.Logs.Gtp
						}
						if o.Role.Device.Webui.Monitor.Logs.DataFiltering != nil {
							nestedRole.Device.Webui.Monitor.Logs.DataFiltering = o.Role.Device.Webui.Monitor.Logs.DataFiltering
						}
						if o.Role.Device.Webui.Monitor.Logs.Iptag != nil {
							nestedRole.Device.Webui.Monitor.Logs.Iptag = o.Role.Device.Webui.Monitor.Logs.Iptag
						}
						if o.Role.Device.Webui.Monitor.Logs.Threat != nil {
							nestedRole.Device.Webui.Monitor.Logs.Threat = o.Role.Device.Webui.Monitor.Logs.Threat
						}
						if o.Role.Device.Webui.Monitor.Logs.Userid != nil {
							nestedRole.Device.Webui.Monitor.Logs.Userid = o.Role.Device.Webui.Monitor.Logs.Userid
						}
						if o.Role.Device.Webui.Monitor.Logs.Configuration != nil {
							nestedRole.Device.Webui.Monitor.Logs.Configuration = o.Role.Device.Webui.Monitor.Logs.Configuration
						}
						if o.Role.Device.Webui.Monitor.Logs.Sctp != nil {
							nestedRole.Device.Webui.Monitor.Logs.Sctp = o.Role.Device.Webui.Monitor.Logs.Sctp
						}
						if o.Role.Device.Webui.Monitor.Logs.System != nil {
							nestedRole.Device.Webui.Monitor.Logs.System = o.Role.Device.Webui.Monitor.Logs.System
						}
						if o.Role.Device.Webui.Monitor.Logs.Traffic != nil {
							nestedRole.Device.Webui.Monitor.Logs.Traffic = o.Role.Device.Webui.Monitor.Logs.Traffic
						}
						if o.Role.Device.Webui.Monitor.Logs.Tunnel != nil {
							nestedRole.Device.Webui.Monitor.Logs.Tunnel = o.Role.Device.Webui.Monitor.Logs.Tunnel
						}
						if o.Role.Device.Webui.Monitor.Logs.Url != nil {
							nestedRole.Device.Webui.Monitor.Logs.Url = o.Role.Device.Webui.Monitor.Logs.Url
						}
						if o.Role.Device.Webui.Monitor.Logs.Alarm != nil {
							nestedRole.Device.Webui.Monitor.Logs.Alarm = o.Role.Device.Webui.Monitor.Logs.Alarm
						}
						if o.Role.Device.Webui.Monitor.Logs.Authentication != nil {
							nestedRole.Device.Webui.Monitor.Logs.Authentication = o.Role.Device.Webui.Monitor.Logs.Authentication
						}
						if o.Role.Device.Webui.Monitor.Logs.Globalprotect != nil {
							nestedRole.Device.Webui.Monitor.Logs.Globalprotect = o.Role.Device.Webui.Monitor.Logs.Globalprotect
						}
						if o.Role.Device.Webui.Monitor.Logs.Hipmatch != nil {
							nestedRole.Device.Webui.Monitor.Logs.Hipmatch = o.Role.Device.Webui.Monitor.Logs.Hipmatch
						}
						if o.Role.Device.Webui.Monitor.Logs.Wildfire != nil {
							nestedRole.Device.Webui.Monitor.Logs.Wildfire = o.Role.Device.Webui.Monitor.Logs.Wildfire
						}
					}
				}
				if o.Role.Device.Webui.Operations != nil {
					nestedRole.Device.Webui.Operations = &RoleDeviceWebuiOperationsXml{}
					if _, ok := o.Misc["RoleDeviceWebuiOperations"]; ok {
						nestedRole.Device.Webui.Operations.Misc = o.Misc["RoleDeviceWebuiOperations"]
					}
					if o.Role.Device.Webui.Operations.DownloadCoreFiles != nil {
						nestedRole.Device.Webui.Operations.DownloadCoreFiles = o.Role.Device.Webui.Operations.DownloadCoreFiles
					}
					if o.Role.Device.Webui.Operations.DownloadPcapFiles != nil {
						nestedRole.Device.Webui.Operations.DownloadPcapFiles = o.Role.Device.Webui.Operations.DownloadPcapFiles
					}
					if o.Role.Device.Webui.Operations.GenerateStatsDumpFile != nil {
						nestedRole.Device.Webui.Operations.GenerateStatsDumpFile = o.Role.Device.Webui.Operations.GenerateStatsDumpFile
					}
					if o.Role.Device.Webui.Operations.GenerateTechSupportFile != nil {
						nestedRole.Device.Webui.Operations.GenerateTechSupportFile = o.Role.Device.Webui.Operations.GenerateTechSupportFile
					}
					if o.Role.Device.Webui.Operations.Reboot != nil {
						nestedRole.Device.Webui.Operations.Reboot = o.Role.Device.Webui.Operations.Reboot
					}
				}
				if o.Role.Device.Webui.Save != nil {
					nestedRole.Device.Webui.Save = &RoleDeviceWebuiSaveXml{}
					if _, ok := o.Misc["RoleDeviceWebuiSave"]; ok {
						nestedRole.Device.Webui.Save.Misc = o.Misc["RoleDeviceWebuiSave"]
					}
					if o.Role.Device.Webui.Save.PartialSave != nil {
						nestedRole.Device.Webui.Save.PartialSave = o.Role.Device.Webui.Save.PartialSave
					}
					if o.Role.Device.Webui.Save.SaveForOtherAdmins != nil {
						nestedRole.Device.Webui.Save.SaveForOtherAdmins = o.Role.Device.Webui.Save.SaveForOtherAdmins
					}
					if o.Role.Device.Webui.Save.ObjectLevelChanges != nil {
						nestedRole.Device.Webui.Save.ObjectLevelChanges = o.Role.Device.Webui.Save.ObjectLevelChanges
					}
				}
				if o.Role.Device.Webui.Validate != nil {
					nestedRole.Device.Webui.Validate = o.Role.Device.Webui.Validate
				}
				if o.Role.Device.Webui.Global != nil {
					nestedRole.Device.Webui.Global = &RoleDeviceWebuiGlobalXml{}
					if _, ok := o.Misc["RoleDeviceWebuiGlobal"]; ok {
						nestedRole.Device.Webui.Global.Misc = o.Misc["RoleDeviceWebuiGlobal"]
					}
					if o.Role.Device.Webui.Global.SystemAlarms != nil {
						nestedRole.Device.Webui.Global.SystemAlarms = o.Role.Device.Webui.Global.SystemAlarms
					}
				}
				if o.Role.Device.Webui.Objects != nil {
					nestedRole.Device.Webui.Objects = &RoleDeviceWebuiObjectsXml{}
					if _, ok := o.Misc["RoleDeviceWebuiObjects"]; ok {
						nestedRole.Device.Webui.Objects.Misc = o.Misc["RoleDeviceWebuiObjects"]
					}
					if o.Role.Device.Webui.Objects.DynamicUserGroups != nil {
						nestedRole.Device.Webui.Objects.DynamicUserGroups = o.Role.Device.Webui.Objects.DynamicUserGroups
					}
					if o.Role.Device.Webui.Objects.Services != nil {
						nestedRole.Device.Webui.Objects.Services = o.Role.Device.Webui.Objects.Services
					}
					if o.Role.Device.Webui.Objects.AddressGroups != nil {
						nestedRole.Device.Webui.Objects.AddressGroups = o.Role.Device.Webui.Objects.AddressGroups
					}
					if o.Role.Device.Webui.Objects.ApplicationGroups != nil {
						nestedRole.Device.Webui.Objects.ApplicationGroups = o.Role.Device.Webui.Objects.ApplicationGroups
					}
					if o.Role.Device.Webui.Objects.DynamicBlockLists != nil {
						nestedRole.Device.Webui.Objects.DynamicBlockLists = o.Role.Device.Webui.Objects.DynamicBlockLists
					}
					if o.Role.Device.Webui.Objects.SecurityProfiles != nil {
						nestedRole.Device.Webui.Objects.SecurityProfiles = &RoleDeviceWebuiObjectsSecurityProfilesXml{}
						if _, ok := o.Misc["RoleDeviceWebuiObjectsSecurityProfiles"]; ok {
							nestedRole.Device.Webui.Objects.SecurityProfiles.Misc = o.Misc["RoleDeviceWebuiObjectsSecurityProfiles"]
						}
						if o.Role.Device.Webui.Objects.SecurityProfiles.SctpProtection != nil {
							nestedRole.Device.Webui.Objects.SecurityProfiles.SctpProtection = o.Role.Device.Webui.Objects.SecurityProfiles.SctpProtection
						}
						if o.Role.Device.Webui.Objects.SecurityProfiles.UrlFiltering != nil {
							nestedRole.Device.Webui.Objects.SecurityProfiles.UrlFiltering = o.Role.Device.Webui.Objects.SecurityProfiles.UrlFiltering
						}
						if o.Role.Device.Webui.Objects.SecurityProfiles.VulnerabilityProtection != nil {
							nestedRole.Device.Webui.Objects.SecurityProfiles.VulnerabilityProtection = o.Role.Device.Webui.Objects.SecurityProfiles.VulnerabilityProtection
						}
						if o.Role.Device.Webui.Objects.SecurityProfiles.WildfireAnalysis != nil {
							nestedRole.Device.Webui.Objects.SecurityProfiles.WildfireAnalysis = o.Role.Device.Webui.Objects.SecurityProfiles.WildfireAnalysis
						}
						if o.Role.Device.Webui.Objects.SecurityProfiles.AntiSpyware != nil {
							nestedRole.Device.Webui.Objects.SecurityProfiles.AntiSpyware = o.Role.Device.Webui.Objects.SecurityProfiles.AntiSpyware
						}
						if o.Role.Device.Webui.Objects.SecurityProfiles.DataFiltering != nil {
							nestedRole.Device.Webui.Objects.SecurityProfiles.DataFiltering = o.Role.Device.Webui.Objects.SecurityProfiles.DataFiltering
						}
						if o.Role.Device.Webui.Objects.SecurityProfiles.FileBlocking != nil {
							nestedRole.Device.Webui.Objects.SecurityProfiles.FileBlocking = o.Role.Device.Webui.Objects.SecurityProfiles.FileBlocking
						}
						if o.Role.Device.Webui.Objects.SecurityProfiles.GtpProtection != nil {
							nestedRole.Device.Webui.Objects.SecurityProfiles.GtpProtection = o.Role.Device.Webui.Objects.SecurityProfiles.GtpProtection
						}
						if o.Role.Device.Webui.Objects.SecurityProfiles.Antivirus != nil {
							nestedRole.Device.Webui.Objects.SecurityProfiles.Antivirus = o.Role.Device.Webui.Objects.SecurityProfiles.Antivirus
						}
						if o.Role.Device.Webui.Objects.SecurityProfiles.DosProtection != nil {
							nestedRole.Device.Webui.Objects.SecurityProfiles.DosProtection = o.Role.Device.Webui.Objects.SecurityProfiles.DosProtection
						}
					}
					if o.Role.Device.Webui.Objects.Applications != nil {
						nestedRole.Device.Webui.Objects.Applications = o.Role.Device.Webui.Objects.Applications
					}
					if o.Role.Device.Webui.Objects.LogForwarding != nil {
						nestedRole.Device.Webui.Objects.LogForwarding = o.Role.Device.Webui.Objects.LogForwarding
					}
					if o.Role.Device.Webui.Objects.Sdwan != nil {
						nestedRole.Device.Webui.Objects.Sdwan = &RoleDeviceWebuiObjectsSdwanXml{}
						if _, ok := o.Misc["RoleDeviceWebuiObjectsSdwan"]; ok {
							nestedRole.Device.Webui.Objects.Sdwan.Misc = o.Misc["RoleDeviceWebuiObjectsSdwan"]
						}
						if o.Role.Device.Webui.Objects.Sdwan.SdwanProfile != nil {
							nestedRole.Device.Webui.Objects.Sdwan.SdwanProfile = o.Role.Device.Webui.Objects.Sdwan.SdwanProfile
						}
						if o.Role.Device.Webui.Objects.Sdwan.SdwanSaasQualityProfile != nil {
							nestedRole.Device.Webui.Objects.Sdwan.SdwanSaasQualityProfile = o.Role.Device.Webui.Objects.Sdwan.SdwanSaasQualityProfile
						}
						if o.Role.Device.Webui.Objects.Sdwan.SdwanDistProfile != nil {
							nestedRole.Device.Webui.Objects.Sdwan.SdwanDistProfile = o.Role.Device.Webui.Objects.Sdwan.SdwanDistProfile
						}
						if o.Role.Device.Webui.Objects.Sdwan.SdwanErrorCorrectionProfile != nil {
							nestedRole.Device.Webui.Objects.Sdwan.SdwanErrorCorrectionProfile = o.Role.Device.Webui.Objects.Sdwan.SdwanErrorCorrectionProfile
						}
					}
					if o.Role.Device.Webui.Objects.GlobalProtect != nil {
						nestedRole.Device.Webui.Objects.GlobalProtect = &RoleDeviceWebuiObjectsGlobalProtectXml{}
						if _, ok := o.Misc["RoleDeviceWebuiObjectsGlobalProtect"]; ok {
							nestedRole.Device.Webui.Objects.GlobalProtect.Misc = o.Misc["RoleDeviceWebuiObjectsGlobalProtect"]
						}
						if o.Role.Device.Webui.Objects.GlobalProtect.HipObjects != nil {
							nestedRole.Device.Webui.Objects.GlobalProtect.HipObjects = o.Role.Device.Webui.Objects.GlobalProtect.HipObjects
						}
						if o.Role.Device.Webui.Objects.GlobalProtect.HipProfiles != nil {
							nestedRole.Device.Webui.Objects.GlobalProtect.HipProfiles = o.Role.Device.Webui.Objects.GlobalProtect.HipProfiles
						}
					}
					if o.Role.Device.Webui.Objects.PacketBrokerProfile != nil {
						nestedRole.Device.Webui.Objects.PacketBrokerProfile = o.Role.Device.Webui.Objects.PacketBrokerProfile
					}
					if o.Role.Device.Webui.Objects.Regions != nil {
						nestedRole.Device.Webui.Objects.Regions = o.Role.Device.Webui.Objects.Regions
					}
					if o.Role.Device.Webui.Objects.Schedules != nil {
						nestedRole.Device.Webui.Objects.Schedules = o.Role.Device.Webui.Objects.Schedules
					}
					if o.Role.Device.Webui.Objects.ServiceGroups != nil {
						nestedRole.Device.Webui.Objects.ServiceGroups = o.Role.Device.Webui.Objects.ServiceGroups
					}
					if o.Role.Device.Webui.Objects.ApplicationFilters != nil {
						nestedRole.Device.Webui.Objects.ApplicationFilters = o.Role.Device.Webui.Objects.ApplicationFilters
					}
					if o.Role.Device.Webui.Objects.Decryption != nil {
						nestedRole.Device.Webui.Objects.Decryption = &RoleDeviceWebuiObjectsDecryptionXml{}
						if _, ok := o.Misc["RoleDeviceWebuiObjectsDecryption"]; ok {
							nestedRole.Device.Webui.Objects.Decryption.Misc = o.Misc["RoleDeviceWebuiObjectsDecryption"]
						}
						if o.Role.Device.Webui.Objects.Decryption.DecryptionProfile != nil {
							nestedRole.Device.Webui.Objects.Decryption.DecryptionProfile = o.Role.Device.Webui.Objects.Decryption.DecryptionProfile
						}
					}
					if o.Role.Device.Webui.Objects.Devices != nil {
						nestedRole.Device.Webui.Objects.Devices = o.Role.Device.Webui.Objects.Devices
					}
					if o.Role.Device.Webui.Objects.Tags != nil {
						nestedRole.Device.Webui.Objects.Tags = o.Role.Device.Webui.Objects.Tags
					}
					if o.Role.Device.Webui.Objects.SecurityProfileGroups != nil {
						nestedRole.Device.Webui.Objects.SecurityProfileGroups = o.Role.Device.Webui.Objects.SecurityProfileGroups
					}
					if o.Role.Device.Webui.Objects.Addresses != nil {
						nestedRole.Device.Webui.Objects.Addresses = o.Role.Device.Webui.Objects.Addresses
					}
					if o.Role.Device.Webui.Objects.Authentication != nil {
						nestedRole.Device.Webui.Objects.Authentication = o.Role.Device.Webui.Objects.Authentication
					}
					if o.Role.Device.Webui.Objects.CustomObjects != nil {
						nestedRole.Device.Webui.Objects.CustomObjects = &RoleDeviceWebuiObjectsCustomObjectsXml{}
						if _, ok := o.Misc["RoleDeviceWebuiObjectsCustomObjects"]; ok {
							nestedRole.Device.Webui.Objects.CustomObjects.Misc = o.Misc["RoleDeviceWebuiObjectsCustomObjects"]
						}
						if o.Role.Device.Webui.Objects.CustomObjects.DataPatterns != nil {
							nestedRole.Device.Webui.Objects.CustomObjects.DataPatterns = o.Role.Device.Webui.Objects.CustomObjects.DataPatterns
						}
						if o.Role.Device.Webui.Objects.CustomObjects.Spyware != nil {
							nestedRole.Device.Webui.Objects.CustomObjects.Spyware = o.Role.Device.Webui.Objects.CustomObjects.Spyware
						}
						if o.Role.Device.Webui.Objects.CustomObjects.UrlCategory != nil {
							nestedRole.Device.Webui.Objects.CustomObjects.UrlCategory = o.Role.Device.Webui.Objects.CustomObjects.UrlCategory
						}
						if o.Role.Device.Webui.Objects.CustomObjects.Vulnerability != nil {
							nestedRole.Device.Webui.Objects.CustomObjects.Vulnerability = o.Role.Device.Webui.Objects.CustomObjects.Vulnerability
						}
					}
				}
				if o.Role.Device.Webui.Policies != nil {
					nestedRole.Device.Webui.Policies = &RoleDeviceWebuiPoliciesXml{}
					if _, ok := o.Misc["RoleDeviceWebuiPolicies"]; ok {
						nestedRole.Device.Webui.Policies.Misc = o.Misc["RoleDeviceWebuiPolicies"]
					}
					if o.Role.Device.Webui.Policies.NatRulebase != nil {
						nestedRole.Device.Webui.Policies.NatRulebase = o.Role.Device.Webui.Policies.NatRulebase
					}
					if o.Role.Device.Webui.Policies.NetworkPacketBrokerRulebase != nil {
						nestedRole.Device.Webui.Policies.NetworkPacketBrokerRulebase = o.Role.Device.Webui.Policies.NetworkPacketBrokerRulebase
					}
					if o.Role.Device.Webui.Policies.SecurityRulebase != nil {
						nestedRole.Device.Webui.Policies.SecurityRulebase = o.Role.Device.Webui.Policies.SecurityRulebase
					}
					if o.Role.Device.Webui.Policies.SslDecryptionRulebase != nil {
						nestedRole.Device.Webui.Policies.SslDecryptionRulebase = o.Role.Device.Webui.Policies.SslDecryptionRulebase
					}
					if o.Role.Device.Webui.Policies.AuthenticationRulebase != nil {
						nestedRole.Device.Webui.Policies.AuthenticationRulebase = o.Role.Device.Webui.Policies.AuthenticationRulebase
					}
					if o.Role.Device.Webui.Policies.DosRulebase != nil {
						nestedRole.Device.Webui.Policies.DosRulebase = o.Role.Device.Webui.Policies.DosRulebase
					}
					if o.Role.Device.Webui.Policies.PbfRulebase != nil {
						nestedRole.Device.Webui.Policies.PbfRulebase = o.Role.Device.Webui.Policies.PbfRulebase
					}
					if o.Role.Device.Webui.Policies.QosRulebase != nil {
						nestedRole.Device.Webui.Policies.QosRulebase = o.Role.Device.Webui.Policies.QosRulebase
					}
					if o.Role.Device.Webui.Policies.RuleHitCountReset != nil {
						nestedRole.Device.Webui.Policies.RuleHitCountReset = o.Role.Device.Webui.Policies.RuleHitCountReset
					}
					if o.Role.Device.Webui.Policies.SdwanRulebase != nil {
						nestedRole.Device.Webui.Policies.SdwanRulebase = o.Role.Device.Webui.Policies.SdwanRulebase
					}
					if o.Role.Device.Webui.Policies.TunnelInspectRulebase != nil {
						nestedRole.Device.Webui.Policies.TunnelInspectRulebase = o.Role.Device.Webui.Policies.TunnelInspectRulebase
					}
					if o.Role.Device.Webui.Policies.ApplicationOverrideRulebase != nil {
						nestedRole.Device.Webui.Policies.ApplicationOverrideRulebase = o.Role.Device.Webui.Policies.ApplicationOverrideRulebase
					}
				}
			}
			if o.Role.Device.Xmlapi != nil {
				nestedRole.Device.Xmlapi = &RoleDeviceXmlapiXml{}
				if _, ok := o.Misc["RoleDeviceXmlapi"]; ok {
					nestedRole.Device.Xmlapi.Misc = o.Misc["RoleDeviceXmlapi"]
				}
				if o.Role.Device.Xmlapi.Commit != nil {
					nestedRole.Device.Xmlapi.Commit = o.Role.Device.Xmlapi.Commit
				}
				if o.Role.Device.Xmlapi.Import != nil {
					nestedRole.Device.Xmlapi.Import = o.Role.Device.Xmlapi.Import
				}
				if o.Role.Device.Xmlapi.Op != nil {
					nestedRole.Device.Xmlapi.Op = o.Role.Device.Xmlapi.Op
				}
				if o.Role.Device.Xmlapi.Report != nil {
					nestedRole.Device.Xmlapi.Report = o.Role.Device.Xmlapi.Report
				}
				if o.Role.Device.Xmlapi.UserId != nil {
					nestedRole.Device.Xmlapi.UserId = o.Role.Device.Xmlapi.UserId
				}
				if o.Role.Device.Xmlapi.Config != nil {
					nestedRole.Device.Xmlapi.Config = o.Role.Device.Xmlapi.Config
				}
				if o.Role.Device.Xmlapi.Export != nil {
					nestedRole.Device.Xmlapi.Export = o.Role.Device.Xmlapi.Export
				}
				if o.Role.Device.Xmlapi.Iot != nil {
					nestedRole.Device.Xmlapi.Iot = o.Role.Device.Xmlapi.Iot
				}
				if o.Role.Device.Xmlapi.Log != nil {
					nestedRole.Device.Xmlapi.Log = o.Role.Device.Xmlapi.Log
				}
			}
			if o.Role.Device.Cli != nil {
				nestedRole.Device.Cli = o.Role.Device.Cli
			}
			if o.Role.Device.Restapi != nil {
				nestedRole.Device.Restapi = &RoleDeviceRestapiXml{}
				if _, ok := o.Misc["RoleDeviceRestapi"]; ok {
					nestedRole.Device.Restapi.Misc = o.Misc["RoleDeviceRestapi"]
				}
				if o.Role.Device.Restapi.Device != nil {
					nestedRole.Device.Restapi.Device = &RoleDeviceRestapiDeviceXml{}
					if _, ok := o.Misc["RoleDeviceRestapiDevice"]; ok {
						nestedRole.Device.Restapi.Device.Misc = o.Misc["RoleDeviceRestapiDevice"]
					}
					if o.Role.Device.Restapi.Device.SyslogServerProfiles != nil {
						nestedRole.Device.Restapi.Device.SyslogServerProfiles = o.Role.Device.Restapi.Device.SyslogServerProfiles
					}
					if o.Role.Device.Restapi.Device.VirtualSystems != nil {
						nestedRole.Device.Restapi.Device.VirtualSystems = o.Role.Device.Restapi.Device.VirtualSystems
					}
					if o.Role.Device.Restapi.Device.EmailServerProfiles != nil {
						nestedRole.Device.Restapi.Device.EmailServerProfiles = o.Role.Device.Restapi.Device.EmailServerProfiles
					}
					if o.Role.Device.Restapi.Device.HttpServerProfiles != nil {
						nestedRole.Device.Restapi.Device.HttpServerProfiles = o.Role.Device.Restapi.Device.HttpServerProfiles
					}
					if o.Role.Device.Restapi.Device.LdapServerProfiles != nil {
						nestedRole.Device.Restapi.Device.LdapServerProfiles = o.Role.Device.Restapi.Device.LdapServerProfiles
					}
					if o.Role.Device.Restapi.Device.LogInterfaceSetting != nil {
						nestedRole.Device.Restapi.Device.LogInterfaceSetting = o.Role.Device.Restapi.Device.LogInterfaceSetting
					}
					if o.Role.Device.Restapi.Device.SnmpTrapServerProfiles != nil {
						nestedRole.Device.Restapi.Device.SnmpTrapServerProfiles = o.Role.Device.Restapi.Device.SnmpTrapServerProfiles
					}
				}
				if o.Role.Device.Restapi.Network != nil {
					nestedRole.Device.Restapi.Network = &RoleDeviceRestapiNetworkXml{}
					if _, ok := o.Misc["RoleDeviceRestapiNetwork"]; ok {
						nestedRole.Device.Restapi.Network.Misc = o.Misc["RoleDeviceRestapiNetwork"]
					}
					if o.Role.Device.Restapi.Network.Lldp != nil {
						nestedRole.Device.Restapi.Network.Lldp = o.Role.Device.Restapi.Network.Lldp
					}
					if o.Role.Device.Restapi.Network.QosNetworkProfiles != nil {
						nestedRole.Device.Restapi.Network.QosNetworkProfiles = o.Role.Device.Restapi.Network.QosNetworkProfiles
					}
					if o.Role.Device.Restapi.Network.TunnelMonitorNetworkProfiles != nil {
						nestedRole.Device.Restapi.Network.TunnelMonitorNetworkProfiles = o.Role.Device.Restapi.Network.TunnelMonitorNetworkProfiles
					}
					if o.Role.Device.Restapi.Network.BgpRoutingProfiles != nil {
						nestedRole.Device.Restapi.Network.BgpRoutingProfiles = o.Role.Device.Restapi.Network.BgpRoutingProfiles
					}
					if o.Role.Device.Restapi.Network.GlobalprotectPortals != nil {
						nestedRole.Device.Restapi.Network.GlobalprotectPortals = o.Role.Device.Restapi.Network.GlobalprotectPortals
					}
					if o.Role.Device.Restapi.Network.IkeGatewayNetworkProfiles != nil {
						nestedRole.Device.Restapi.Network.IkeGatewayNetworkProfiles = o.Role.Device.Restapi.Network.IkeGatewayNetworkProfiles
					}
					if o.Role.Device.Restapi.Network.IpsecCryptoNetworkProfiles != nil {
						nestedRole.Device.Restapi.Network.IpsecCryptoNetworkProfiles = o.Role.Device.Restapi.Network.IpsecCryptoNetworkProfiles
					}
					if o.Role.Device.Restapi.Network.SdwanInterfaceProfiles != nil {
						nestedRole.Device.Restapi.Network.SdwanInterfaceProfiles = o.Role.Device.Restapi.Network.SdwanInterfaceProfiles
					}
					if o.Role.Device.Restapi.Network.DhcpServers != nil {
						nestedRole.Device.Restapi.Network.DhcpServers = o.Role.Device.Restapi.Network.DhcpServers
					}
					if o.Role.Device.Restapi.Network.DnsProxies != nil {
						nestedRole.Device.Restapi.Network.DnsProxies = o.Role.Device.Restapi.Network.DnsProxies
					}
					if o.Role.Device.Restapi.Network.EthernetInterfaces != nil {
						nestedRole.Device.Restapi.Network.EthernetInterfaces = o.Role.Device.Restapi.Network.EthernetInterfaces
					}
					if o.Role.Device.Restapi.Network.GlobalprotectClientlessApps != nil {
						nestedRole.Device.Restapi.Network.GlobalprotectClientlessApps = o.Role.Device.Restapi.Network.GlobalprotectClientlessApps
					}
					if o.Role.Device.Restapi.Network.DhcpRelays != nil {
						nestedRole.Device.Restapi.Network.DhcpRelays = o.Role.Device.Restapi.Network.DhcpRelays
					}
					if o.Role.Device.Restapi.Network.LogicalRouters != nil {
						nestedRole.Device.Restapi.Network.LogicalRouters = o.Role.Device.Restapi.Network.LogicalRouters
					}
					if o.Role.Device.Restapi.Network.SdwanInterfaces != nil {
						nestedRole.Device.Restapi.Network.SdwanInterfaces = o.Role.Device.Restapi.Network.SdwanInterfaces
					}
					if o.Role.Device.Restapi.Network.Vlans != nil {
						nestedRole.Device.Restapi.Network.Vlans = o.Role.Device.Restapi.Network.Vlans
					}
					if o.Role.Device.Restapi.Network.AggregateEthernetInterfaces != nil {
						nestedRole.Device.Restapi.Network.AggregateEthernetInterfaces = o.Role.Device.Restapi.Network.AggregateEthernetInterfaces
					}
					if o.Role.Device.Restapi.Network.GlobalprotectIpsecCryptoNetworkProfiles != nil {
						nestedRole.Device.Restapi.Network.GlobalprotectIpsecCryptoNetworkProfiles = o.Role.Device.Restapi.Network.GlobalprotectIpsecCryptoNetworkProfiles
					}
					if o.Role.Device.Restapi.Network.LoopbackInterfaces != nil {
						nestedRole.Device.Restapi.Network.LoopbackInterfaces = o.Role.Device.Restapi.Network.LoopbackInterfaces
					}
					if o.Role.Device.Restapi.Network.QosInterfaces != nil {
						nestedRole.Device.Restapi.Network.QosInterfaces = o.Role.Device.Restapi.Network.QosInterfaces
					}
					if o.Role.Device.Restapi.Network.LldpNetworkProfiles != nil {
						nestedRole.Device.Restapi.Network.LldpNetworkProfiles = o.Role.Device.Restapi.Network.LldpNetworkProfiles
					}
					if o.Role.Device.Restapi.Network.VlanInterfaces != nil {
						nestedRole.Device.Restapi.Network.VlanInterfaces = o.Role.Device.Restapi.Network.VlanInterfaces
					}
					if o.Role.Device.Restapi.Network.Zones != nil {
						nestedRole.Device.Restapi.Network.Zones = o.Role.Device.Restapi.Network.Zones
					}
					if o.Role.Device.Restapi.Network.BfdNetworkProfiles != nil {
						nestedRole.Device.Restapi.Network.BfdNetworkProfiles = o.Role.Device.Restapi.Network.BfdNetworkProfiles
					}
					if o.Role.Device.Restapi.Network.GreTunnels != nil {
						nestedRole.Device.Restapi.Network.GreTunnels = o.Role.Device.Restapi.Network.GreTunnels
					}
					if o.Role.Device.Restapi.Network.IkeCryptoNetworkProfiles != nil {
						nestedRole.Device.Restapi.Network.IkeCryptoNetworkProfiles = o.Role.Device.Restapi.Network.IkeCryptoNetworkProfiles
					}
					if o.Role.Device.Restapi.Network.IpsecTunnels != nil {
						nestedRole.Device.Restapi.Network.IpsecTunnels = o.Role.Device.Restapi.Network.IpsecTunnels
					}
					if o.Role.Device.Restapi.Network.GlobalprotectClientlessAppGroups != nil {
						nestedRole.Device.Restapi.Network.GlobalprotectClientlessAppGroups = o.Role.Device.Restapi.Network.GlobalprotectClientlessAppGroups
					}
					if o.Role.Device.Restapi.Network.GlobalprotectGateways != nil {
						nestedRole.Device.Restapi.Network.GlobalprotectGateways = o.Role.Device.Restapi.Network.GlobalprotectGateways
					}
					if o.Role.Device.Restapi.Network.GlobalprotectMdmServers != nil {
						nestedRole.Device.Restapi.Network.GlobalprotectMdmServers = o.Role.Device.Restapi.Network.GlobalprotectMdmServers
					}
					if o.Role.Device.Restapi.Network.ZoneProtectionNetworkProfiles != nil {
						nestedRole.Device.Restapi.Network.ZoneProtectionNetworkProfiles = o.Role.Device.Restapi.Network.ZoneProtectionNetworkProfiles
					}
					if o.Role.Device.Restapi.Network.InterfaceManagementNetworkProfiles != nil {
						nestedRole.Device.Restapi.Network.InterfaceManagementNetworkProfiles = o.Role.Device.Restapi.Network.InterfaceManagementNetworkProfiles
					}
					if o.Role.Device.Restapi.Network.TunnelInterfaces != nil {
						nestedRole.Device.Restapi.Network.TunnelInterfaces = o.Role.Device.Restapi.Network.TunnelInterfaces
					}
					if o.Role.Device.Restapi.Network.VirtualRouters != nil {
						nestedRole.Device.Restapi.Network.VirtualRouters = o.Role.Device.Restapi.Network.VirtualRouters
					}
					if o.Role.Device.Restapi.Network.VirtualWires != nil {
						nestedRole.Device.Restapi.Network.VirtualWires = o.Role.Device.Restapi.Network.VirtualWires
					}
				}
				if o.Role.Device.Restapi.Objects != nil {
					nestedRole.Device.Restapi.Objects = &RoleDeviceRestapiObjectsXml{}
					if _, ok := o.Misc["RoleDeviceRestapiObjects"]; ok {
						nestedRole.Device.Restapi.Objects.Misc = o.Misc["RoleDeviceRestapiObjects"]
					}
					if o.Role.Device.Restapi.Objects.ExternalDynamicLists != nil {
						nestedRole.Device.Restapi.Objects.ExternalDynamicLists = o.Role.Device.Restapi.Objects.ExternalDynamicLists
					}
					if o.Role.Device.Restapi.Objects.AddressGroups != nil {
						nestedRole.Device.Restapi.Objects.AddressGroups = o.Role.Device.Restapi.Objects.AddressGroups
					}
					if o.Role.Device.Restapi.Objects.AntivirusSecurityProfiles != nil {
						nestedRole.Device.Restapi.Objects.AntivirusSecurityProfiles = o.Role.Device.Restapi.Objects.AntivirusSecurityProfiles
					}
					if o.Role.Device.Restapi.Objects.SecurityProfileGroups != nil {
						nestedRole.Device.Restapi.Objects.SecurityProfileGroups = o.Role.Device.Restapi.Objects.SecurityProfileGroups
					}
					if o.Role.Device.Restapi.Objects.ServiceGroups != nil {
						nestedRole.Device.Restapi.Objects.ServiceGroups = o.Role.Device.Restapi.Objects.ServiceGroups
					}
					if o.Role.Device.Restapi.Objects.Addresses != nil {
						nestedRole.Device.Restapi.Objects.Addresses = o.Role.Device.Restapi.Objects.Addresses
					}
					if o.Role.Device.Restapi.Objects.ApplicationFilters != nil {
						nestedRole.Device.Restapi.Objects.ApplicationFilters = o.Role.Device.Restapi.Objects.ApplicationFilters
					}
					if o.Role.Device.Restapi.Objects.AuthenticationEnforcements != nil {
						nestedRole.Device.Restapi.Objects.AuthenticationEnforcements = o.Role.Device.Restapi.Objects.AuthenticationEnforcements
					}
					if o.Role.Device.Restapi.Objects.FileBlockingSecurityProfiles != nil {
						nestedRole.Device.Restapi.Objects.FileBlockingSecurityProfiles = o.Role.Device.Restapi.Objects.FileBlockingSecurityProfiles
					}
					if o.Role.Device.Restapi.Objects.PacketBrokerProfiles != nil {
						nestedRole.Device.Restapi.Objects.PacketBrokerProfiles = o.Role.Device.Restapi.Objects.PacketBrokerProfiles
					}
					if o.Role.Device.Restapi.Objects.Schedules != nil {
						nestedRole.Device.Restapi.Objects.Schedules = o.Role.Device.Restapi.Objects.Schedules
					}
					if o.Role.Device.Restapi.Objects.SctpProtectionSecurityProfiles != nil {
						nestedRole.Device.Restapi.Objects.SctpProtectionSecurityProfiles = o.Role.Device.Restapi.Objects.SctpProtectionSecurityProfiles
					}
					if o.Role.Device.Restapi.Objects.SdwanSaasQualityProfiles != nil {
						nestedRole.Device.Restapi.Objects.SdwanSaasQualityProfiles = o.Role.Device.Restapi.Objects.SdwanSaasQualityProfiles
					}
					if o.Role.Device.Restapi.Objects.AntiSpywareSecurityProfiles != nil {
						nestedRole.Device.Restapi.Objects.AntiSpywareSecurityProfiles = o.Role.Device.Restapi.Objects.AntiSpywareSecurityProfiles
					}
					if o.Role.Device.Restapi.Objects.ApplicationGroups != nil {
						nestedRole.Device.Restapi.Objects.ApplicationGroups = o.Role.Device.Restapi.Objects.ApplicationGroups
					}
					if o.Role.Device.Restapi.Objects.CustomDataPatterns != nil {
						nestedRole.Device.Restapi.Objects.CustomDataPatterns = o.Role.Device.Restapi.Objects.CustomDataPatterns
					}
					if o.Role.Device.Restapi.Objects.DataFilteringSecurityProfiles != nil {
						nestedRole.Device.Restapi.Objects.DataFilteringSecurityProfiles = o.Role.Device.Restapi.Objects.DataFilteringSecurityProfiles
					}
					if o.Role.Device.Restapi.Objects.GlobalprotectHipProfiles != nil {
						nestedRole.Device.Restapi.Objects.GlobalprotectHipProfiles = o.Role.Device.Restapi.Objects.GlobalprotectHipProfiles
					}
					if o.Role.Device.Restapi.Objects.LogForwardingProfiles != nil {
						nestedRole.Device.Restapi.Objects.LogForwardingProfiles = o.Role.Device.Restapi.Objects.LogForwardingProfiles
					}
					if o.Role.Device.Restapi.Objects.WildfireAnalysisSecurityProfiles != nil {
						nestedRole.Device.Restapi.Objects.WildfireAnalysisSecurityProfiles = o.Role.Device.Restapi.Objects.WildfireAnalysisSecurityProfiles
					}
					if o.Role.Device.Restapi.Objects.CustomSpywareSignatures != nil {
						nestedRole.Device.Restapi.Objects.CustomSpywareSignatures = o.Role.Device.Restapi.Objects.CustomSpywareSignatures
					}
					if o.Role.Device.Restapi.Objects.DecryptionProfiles != nil {
						nestedRole.Device.Restapi.Objects.DecryptionProfiles = o.Role.Device.Restapi.Objects.DecryptionProfiles
					}
					if o.Role.Device.Restapi.Objects.DynamicUserGroups != nil {
						nestedRole.Device.Restapi.Objects.DynamicUserGroups = o.Role.Device.Restapi.Objects.DynamicUserGroups
					}
					if o.Role.Device.Restapi.Objects.UrlFilteringSecurityProfiles != nil {
						nestedRole.Device.Restapi.Objects.UrlFilteringSecurityProfiles = o.Role.Device.Restapi.Objects.UrlFilteringSecurityProfiles
					}
					if o.Role.Device.Restapi.Objects.SdwanPathQualityProfiles != nil {
						nestedRole.Device.Restapi.Objects.SdwanPathQualityProfiles = o.Role.Device.Restapi.Objects.SdwanPathQualityProfiles
					}
					if o.Role.Device.Restapi.Objects.SdwanTrafficDistributionProfiles != nil {
						nestedRole.Device.Restapi.Objects.SdwanTrafficDistributionProfiles = o.Role.Device.Restapi.Objects.SdwanTrafficDistributionProfiles
					}
					if o.Role.Device.Restapi.Objects.VulnerabilityProtectionSecurityProfiles != nil {
						nestedRole.Device.Restapi.Objects.VulnerabilityProtectionSecurityProfiles = o.Role.Device.Restapi.Objects.VulnerabilityProtectionSecurityProfiles
					}
					if o.Role.Device.Restapi.Objects.Applications != nil {
						nestedRole.Device.Restapi.Objects.Applications = o.Role.Device.Restapi.Objects.Applications
					}
					if o.Role.Device.Restapi.Objects.CustomVulnerabilitySignatures != nil {
						nestedRole.Device.Restapi.Objects.CustomVulnerabilitySignatures = o.Role.Device.Restapi.Objects.CustomVulnerabilitySignatures
					}
					if o.Role.Device.Restapi.Objects.GlobalprotectHipObjects != nil {
						nestedRole.Device.Restapi.Objects.GlobalprotectHipObjects = o.Role.Device.Restapi.Objects.GlobalprotectHipObjects
					}
					if o.Role.Device.Restapi.Objects.SdwanErrorCorrectionProfiles != nil {
						nestedRole.Device.Restapi.Objects.SdwanErrorCorrectionProfiles = o.Role.Device.Restapi.Objects.SdwanErrorCorrectionProfiles
					}
					if o.Role.Device.Restapi.Objects.Services != nil {
						nestedRole.Device.Restapi.Objects.Services = o.Role.Device.Restapi.Objects.Services
					}
					if o.Role.Device.Restapi.Objects.Tags != nil {
						nestedRole.Device.Restapi.Objects.Tags = o.Role.Device.Restapi.Objects.Tags
					}
					if o.Role.Device.Restapi.Objects.CustomUrlCategories != nil {
						nestedRole.Device.Restapi.Objects.CustomUrlCategories = o.Role.Device.Restapi.Objects.CustomUrlCategories
					}
					if o.Role.Device.Restapi.Objects.Devices != nil {
						nestedRole.Device.Restapi.Objects.Devices = o.Role.Device.Restapi.Objects.Devices
					}
					if o.Role.Device.Restapi.Objects.DosProtectionSecurityProfiles != nil {
						nestedRole.Device.Restapi.Objects.DosProtectionSecurityProfiles = o.Role.Device.Restapi.Objects.DosProtectionSecurityProfiles
					}
					if o.Role.Device.Restapi.Objects.GtpProtectionSecurityProfiles != nil {
						nestedRole.Device.Restapi.Objects.GtpProtectionSecurityProfiles = o.Role.Device.Restapi.Objects.GtpProtectionSecurityProfiles
					}
					if o.Role.Device.Restapi.Objects.Regions != nil {
						nestedRole.Device.Restapi.Objects.Regions = o.Role.Device.Restapi.Objects.Regions
					}
				}
				if o.Role.Device.Restapi.Policies != nil {
					nestedRole.Device.Restapi.Policies = &RoleDeviceRestapiPoliciesXml{}
					if _, ok := o.Misc["RoleDeviceRestapiPolicies"]; ok {
						nestedRole.Device.Restapi.Policies.Misc = o.Misc["RoleDeviceRestapiPolicies"]
					}
					if o.Role.Device.Restapi.Policies.TunnelInspectionRules != nil {
						nestedRole.Device.Restapi.Policies.TunnelInspectionRules = o.Role.Device.Restapi.Policies.TunnelInspectionRules
					}
					if o.Role.Device.Restapi.Policies.AuthenticationRules != nil {
						nestedRole.Device.Restapi.Policies.AuthenticationRules = o.Role.Device.Restapi.Policies.AuthenticationRules
					}
					if o.Role.Device.Restapi.Policies.DosRules != nil {
						nestedRole.Device.Restapi.Policies.DosRules = o.Role.Device.Restapi.Policies.DosRules
					}
					if o.Role.Device.Restapi.Policies.NatRules != nil {
						nestedRole.Device.Restapi.Policies.NatRules = o.Role.Device.Restapi.Policies.NatRules
					}
					if o.Role.Device.Restapi.Policies.NetworkPacketBrokerRules != nil {
						nestedRole.Device.Restapi.Policies.NetworkPacketBrokerRules = o.Role.Device.Restapi.Policies.NetworkPacketBrokerRules
					}
					if o.Role.Device.Restapi.Policies.SdwanRules != nil {
						nestedRole.Device.Restapi.Policies.SdwanRules = o.Role.Device.Restapi.Policies.SdwanRules
					}
					if o.Role.Device.Restapi.Policies.SecurityRules != nil {
						nestedRole.Device.Restapi.Policies.SecurityRules = o.Role.Device.Restapi.Policies.SecurityRules
					}
					if o.Role.Device.Restapi.Policies.ApplicationOverrideRules != nil {
						nestedRole.Device.Restapi.Policies.ApplicationOverrideRules = o.Role.Device.Restapi.Policies.ApplicationOverrideRules
					}
					if o.Role.Device.Restapi.Policies.DecryptionRules != nil {
						nestedRole.Device.Restapi.Policies.DecryptionRules = o.Role.Device.Restapi.Policies.DecryptionRules
					}
					if o.Role.Device.Restapi.Policies.PolicyBasedForwardingRules != nil {
						nestedRole.Device.Restapi.Policies.PolicyBasedForwardingRules = o.Role.Device.Restapi.Policies.PolicyBasedForwardingRules
					}
					if o.Role.Device.Restapi.Policies.QosRules != nil {
						nestedRole.Device.Restapi.Policies.QosRules = o.Role.Device.Restapi.Policies.QosRules
					}
				}
				if o.Role.Device.Restapi.System != nil {
					nestedRole.Device.Restapi.System = &RoleDeviceRestapiSystemXml{}
					if _, ok := o.Misc["RoleDeviceRestapiSystem"]; ok {
						nestedRole.Device.Restapi.System.Misc = o.Misc["RoleDeviceRestapiSystem"]
					}
					if o.Role.Device.Restapi.System.Configuration != nil {
						nestedRole.Device.Restapi.System.Configuration = o.Role.Device.Restapi.System.Configuration
					}
				}
			}
		}
		if o.Role.Vsys != nil {
			nestedRole.Vsys = &RoleVsysXml{}
			if _, ok := o.Misc["RoleVsys"]; ok {
				nestedRole.Vsys.Misc = o.Misc["RoleVsys"]
			}
			if o.Role.Vsys.Cli != nil {
				nestedRole.Vsys.Cli = o.Role.Vsys.Cli
			}
			if o.Role.Vsys.Restapi != nil {
				nestedRole.Vsys.Restapi = &RoleVsysRestapiXml{}
				if _, ok := o.Misc["RoleVsysRestapi"]; ok {
					nestedRole.Vsys.Restapi.Misc = o.Misc["RoleVsysRestapi"]
				}
				if o.Role.Vsys.Restapi.Device != nil {
					nestedRole.Vsys.Restapi.Device = &RoleVsysRestapiDeviceXml{}
					if _, ok := o.Misc["RoleVsysRestapiDevice"]; ok {
						nestedRole.Vsys.Restapi.Device.Misc = o.Misc["RoleVsysRestapiDevice"]
					}
					if o.Role.Vsys.Restapi.Device.VirtualSystems != nil {
						nestedRole.Vsys.Restapi.Device.VirtualSystems = o.Role.Vsys.Restapi.Device.VirtualSystems
					}
					if o.Role.Vsys.Restapi.Device.EmailServerProfiles != nil {
						nestedRole.Vsys.Restapi.Device.EmailServerProfiles = o.Role.Vsys.Restapi.Device.EmailServerProfiles
					}
					if o.Role.Vsys.Restapi.Device.HttpServerProfiles != nil {
						nestedRole.Vsys.Restapi.Device.HttpServerProfiles = o.Role.Vsys.Restapi.Device.HttpServerProfiles
					}
					if o.Role.Vsys.Restapi.Device.LdapServerProfiles != nil {
						nestedRole.Vsys.Restapi.Device.LdapServerProfiles = o.Role.Vsys.Restapi.Device.LdapServerProfiles
					}
					if o.Role.Vsys.Restapi.Device.LogInterfaceSetting != nil {
						nestedRole.Vsys.Restapi.Device.LogInterfaceSetting = o.Role.Vsys.Restapi.Device.LogInterfaceSetting
					}
					if o.Role.Vsys.Restapi.Device.SnmpTrapServerProfiles != nil {
						nestedRole.Vsys.Restapi.Device.SnmpTrapServerProfiles = o.Role.Vsys.Restapi.Device.SnmpTrapServerProfiles
					}
					if o.Role.Vsys.Restapi.Device.SyslogServerProfiles != nil {
						nestedRole.Vsys.Restapi.Device.SyslogServerProfiles = o.Role.Vsys.Restapi.Device.SyslogServerProfiles
					}
				}
				if o.Role.Vsys.Restapi.Network != nil {
					nestedRole.Vsys.Restapi.Network = &RoleVsysRestapiNetworkXml{}
					if _, ok := o.Misc["RoleVsysRestapiNetwork"]; ok {
						nestedRole.Vsys.Restapi.Network.Misc = o.Misc["RoleVsysRestapiNetwork"]
					}
					if o.Role.Vsys.Restapi.Network.Zones != nil {
						nestedRole.Vsys.Restapi.Network.Zones = o.Role.Vsys.Restapi.Network.Zones
					}
					if o.Role.Vsys.Restapi.Network.SdwanInterfaceProfiles != nil {
						nestedRole.Vsys.Restapi.Network.SdwanInterfaceProfiles = o.Role.Vsys.Restapi.Network.SdwanInterfaceProfiles
					}
					if o.Role.Vsys.Restapi.Network.GlobalprotectClientlessAppGroups != nil {
						nestedRole.Vsys.Restapi.Network.GlobalprotectClientlessAppGroups = o.Role.Vsys.Restapi.Network.GlobalprotectClientlessAppGroups
					}
					if o.Role.Vsys.Restapi.Network.GlobalprotectClientlessApps != nil {
						nestedRole.Vsys.Restapi.Network.GlobalprotectClientlessApps = o.Role.Vsys.Restapi.Network.GlobalprotectClientlessApps
					}
					if o.Role.Vsys.Restapi.Network.GlobalprotectGateways != nil {
						nestedRole.Vsys.Restapi.Network.GlobalprotectGateways = o.Role.Vsys.Restapi.Network.GlobalprotectGateways
					}
					if o.Role.Vsys.Restapi.Network.GlobalprotectMdmServers != nil {
						nestedRole.Vsys.Restapi.Network.GlobalprotectMdmServers = o.Role.Vsys.Restapi.Network.GlobalprotectMdmServers
					}
					if o.Role.Vsys.Restapi.Network.GlobalprotectPortals != nil {
						nestedRole.Vsys.Restapi.Network.GlobalprotectPortals = o.Role.Vsys.Restapi.Network.GlobalprotectPortals
					}
				}
				if o.Role.Vsys.Restapi.Objects != nil {
					nestedRole.Vsys.Restapi.Objects = &RoleVsysRestapiObjectsXml{}
					if _, ok := o.Misc["RoleVsysRestapiObjects"]; ok {
						nestedRole.Vsys.Restapi.Objects.Misc = o.Misc["RoleVsysRestapiObjects"]
					}
					if o.Role.Vsys.Restapi.Objects.SdwanErrorCorrectionProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.SdwanErrorCorrectionProfiles = o.Role.Vsys.Restapi.Objects.SdwanErrorCorrectionProfiles
					}
					if o.Role.Vsys.Restapi.Objects.Tags != nil {
						nestedRole.Vsys.Restapi.Objects.Tags = o.Role.Vsys.Restapi.Objects.Tags
					}
					if o.Role.Vsys.Restapi.Objects.SdwanSaasQualityProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.SdwanSaasQualityProfiles = o.Role.Vsys.Restapi.Objects.SdwanSaasQualityProfiles
					}
					if o.Role.Vsys.Restapi.Objects.CustomDataPatterns != nil {
						nestedRole.Vsys.Restapi.Objects.CustomDataPatterns = o.Role.Vsys.Restapi.Objects.CustomDataPatterns
					}
					if o.Role.Vsys.Restapi.Objects.CustomVulnerabilitySignatures != nil {
						nestedRole.Vsys.Restapi.Objects.CustomVulnerabilitySignatures = o.Role.Vsys.Restapi.Objects.CustomVulnerabilitySignatures
					}
					if o.Role.Vsys.Restapi.Objects.DynamicUserGroups != nil {
						nestedRole.Vsys.Restapi.Objects.DynamicUserGroups = o.Role.Vsys.Restapi.Objects.DynamicUserGroups
					}
					if o.Role.Vsys.Restapi.Objects.FileBlockingSecurityProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.FileBlockingSecurityProfiles = o.Role.Vsys.Restapi.Objects.FileBlockingSecurityProfiles
					}
					if o.Role.Vsys.Restapi.Objects.GlobalprotectHipProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.GlobalprotectHipProfiles = o.Role.Vsys.Restapi.Objects.GlobalprotectHipProfiles
					}
					if o.Role.Vsys.Restapi.Objects.LogForwardingProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.LogForwardingProfiles = o.Role.Vsys.Restapi.Objects.LogForwardingProfiles
					}
					if o.Role.Vsys.Restapi.Objects.SdwanPathQualityProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.SdwanPathQualityProfiles = o.Role.Vsys.Restapi.Objects.SdwanPathQualityProfiles
					}
					if o.Role.Vsys.Restapi.Objects.AntivirusSecurityProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.AntivirusSecurityProfiles = o.Role.Vsys.Restapi.Objects.AntivirusSecurityProfiles
					}
					if o.Role.Vsys.Restapi.Objects.CustomSpywareSignatures != nil {
						nestedRole.Vsys.Restapi.Objects.CustomSpywareSignatures = o.Role.Vsys.Restapi.Objects.CustomSpywareSignatures
					}
					if o.Role.Vsys.Restapi.Objects.DosProtectionSecurityProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.DosProtectionSecurityProfiles = o.Role.Vsys.Restapi.Objects.DosProtectionSecurityProfiles
					}
					if o.Role.Vsys.Restapi.Objects.PacketBrokerProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.PacketBrokerProfiles = o.Role.Vsys.Restapi.Objects.PacketBrokerProfiles
					}
					if o.Role.Vsys.Restapi.Objects.SdwanTrafficDistributionProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.SdwanTrafficDistributionProfiles = o.Role.Vsys.Restapi.Objects.SdwanTrafficDistributionProfiles
					}
					if o.Role.Vsys.Restapi.Objects.AntiSpywareSecurityProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.AntiSpywareSecurityProfiles = o.Role.Vsys.Restapi.Objects.AntiSpywareSecurityProfiles
					}
					if o.Role.Vsys.Restapi.Objects.Applications != nil {
						nestedRole.Vsys.Restapi.Objects.Applications = o.Role.Vsys.Restapi.Objects.Applications
					}
					if o.Role.Vsys.Restapi.Objects.Regions != nil {
						nestedRole.Vsys.Restapi.Objects.Regions = o.Role.Vsys.Restapi.Objects.Regions
					}
					if o.Role.Vsys.Restapi.Objects.VulnerabilityProtectionSecurityProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.VulnerabilityProtectionSecurityProfiles = o.Role.Vsys.Restapi.Objects.VulnerabilityProtectionSecurityProfiles
					}
					if o.Role.Vsys.Restapi.Objects.AuthenticationEnforcements != nil {
						nestedRole.Vsys.Restapi.Objects.AuthenticationEnforcements = o.Role.Vsys.Restapi.Objects.AuthenticationEnforcements
					}
					if o.Role.Vsys.Restapi.Objects.DataFilteringSecurityProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.DataFilteringSecurityProfiles = o.Role.Vsys.Restapi.Objects.DataFilteringSecurityProfiles
					}
					if o.Role.Vsys.Restapi.Objects.ApplicationGroups != nil {
						nestedRole.Vsys.Restapi.Objects.ApplicationGroups = o.Role.Vsys.Restapi.Objects.ApplicationGroups
					}
					if o.Role.Vsys.Restapi.Objects.Devices != nil {
						nestedRole.Vsys.Restapi.Objects.Devices = o.Role.Vsys.Restapi.Objects.Devices
					}
					if o.Role.Vsys.Restapi.Objects.Schedules != nil {
						nestedRole.Vsys.Restapi.Objects.Schedules = o.Role.Vsys.Restapi.Objects.Schedules
					}
					if o.Role.Vsys.Restapi.Objects.SecurityProfileGroups != nil {
						nestedRole.Vsys.Restapi.Objects.SecurityProfileGroups = o.Role.Vsys.Restapi.Objects.SecurityProfileGroups
					}
					if o.Role.Vsys.Restapi.Objects.Services != nil {
						nestedRole.Vsys.Restapi.Objects.Services = o.Role.Vsys.Restapi.Objects.Services
					}
					if o.Role.Vsys.Restapi.Objects.WildfireAnalysisSecurityProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.WildfireAnalysisSecurityProfiles = o.Role.Vsys.Restapi.Objects.WildfireAnalysisSecurityProfiles
					}
					if o.Role.Vsys.Restapi.Objects.AddressGroups != nil {
						nestedRole.Vsys.Restapi.Objects.AddressGroups = o.Role.Vsys.Restapi.Objects.AddressGroups
					}
					if o.Role.Vsys.Restapi.Objects.ApplicationFilters != nil {
						nestedRole.Vsys.Restapi.Objects.ApplicationFilters = o.Role.Vsys.Restapi.Objects.ApplicationFilters
					}
					if o.Role.Vsys.Restapi.Objects.DecryptionProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.DecryptionProfiles = o.Role.Vsys.Restapi.Objects.DecryptionProfiles
					}
					if o.Role.Vsys.Restapi.Objects.GlobalprotectHipObjects != nil {
						nestedRole.Vsys.Restapi.Objects.GlobalprotectHipObjects = o.Role.Vsys.Restapi.Objects.GlobalprotectHipObjects
					}
					if o.Role.Vsys.Restapi.Objects.GtpProtectionSecurityProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.GtpProtectionSecurityProfiles = o.Role.Vsys.Restapi.Objects.GtpProtectionSecurityProfiles
					}
					if o.Role.Vsys.Restapi.Objects.ServiceGroups != nil {
						nestedRole.Vsys.Restapi.Objects.ServiceGroups = o.Role.Vsys.Restapi.Objects.ServiceGroups
					}
					if o.Role.Vsys.Restapi.Objects.Addresses != nil {
						nestedRole.Vsys.Restapi.Objects.Addresses = o.Role.Vsys.Restapi.Objects.Addresses
					}
					if o.Role.Vsys.Restapi.Objects.CustomUrlCategories != nil {
						nestedRole.Vsys.Restapi.Objects.CustomUrlCategories = o.Role.Vsys.Restapi.Objects.CustomUrlCategories
					}
					if o.Role.Vsys.Restapi.Objects.UrlFilteringSecurityProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.UrlFilteringSecurityProfiles = o.Role.Vsys.Restapi.Objects.UrlFilteringSecurityProfiles
					}
					if o.Role.Vsys.Restapi.Objects.ExternalDynamicLists != nil {
						nestedRole.Vsys.Restapi.Objects.ExternalDynamicLists = o.Role.Vsys.Restapi.Objects.ExternalDynamicLists
					}
					if o.Role.Vsys.Restapi.Objects.SctpProtectionSecurityProfiles != nil {
						nestedRole.Vsys.Restapi.Objects.SctpProtectionSecurityProfiles = o.Role.Vsys.Restapi.Objects.SctpProtectionSecurityProfiles
					}
				}
				if o.Role.Vsys.Restapi.Policies != nil {
					nestedRole.Vsys.Restapi.Policies = &RoleVsysRestapiPoliciesXml{}
					if _, ok := o.Misc["RoleVsysRestapiPolicies"]; ok {
						nestedRole.Vsys.Restapi.Policies.Misc = o.Misc["RoleVsysRestapiPolicies"]
					}
					if o.Role.Vsys.Restapi.Policies.PolicyBasedForwardingRules != nil {
						nestedRole.Vsys.Restapi.Policies.PolicyBasedForwardingRules = o.Role.Vsys.Restapi.Policies.PolicyBasedForwardingRules
					}
					if o.Role.Vsys.Restapi.Policies.QosRules != nil {
						nestedRole.Vsys.Restapi.Policies.QosRules = o.Role.Vsys.Restapi.Policies.QosRules
					}
					if o.Role.Vsys.Restapi.Policies.SecurityRules != nil {
						nestedRole.Vsys.Restapi.Policies.SecurityRules = o.Role.Vsys.Restapi.Policies.SecurityRules
					}
					if o.Role.Vsys.Restapi.Policies.TunnelInspectionRules != nil {
						nestedRole.Vsys.Restapi.Policies.TunnelInspectionRules = o.Role.Vsys.Restapi.Policies.TunnelInspectionRules
					}
					if o.Role.Vsys.Restapi.Policies.ApplicationOverrideRules != nil {
						nestedRole.Vsys.Restapi.Policies.ApplicationOverrideRules = o.Role.Vsys.Restapi.Policies.ApplicationOverrideRules
					}
					if o.Role.Vsys.Restapi.Policies.AuthenticationRules != nil {
						nestedRole.Vsys.Restapi.Policies.AuthenticationRules = o.Role.Vsys.Restapi.Policies.AuthenticationRules
					}
					if o.Role.Vsys.Restapi.Policies.DecryptionRules != nil {
						nestedRole.Vsys.Restapi.Policies.DecryptionRules = o.Role.Vsys.Restapi.Policies.DecryptionRules
					}
					if o.Role.Vsys.Restapi.Policies.NatRules != nil {
						nestedRole.Vsys.Restapi.Policies.NatRules = o.Role.Vsys.Restapi.Policies.NatRules
					}
					if o.Role.Vsys.Restapi.Policies.DosRules != nil {
						nestedRole.Vsys.Restapi.Policies.DosRules = o.Role.Vsys.Restapi.Policies.DosRules
					}
					if o.Role.Vsys.Restapi.Policies.NetworkPacketBrokerRules != nil {
						nestedRole.Vsys.Restapi.Policies.NetworkPacketBrokerRules = o.Role.Vsys.Restapi.Policies.NetworkPacketBrokerRules
					}
					if o.Role.Vsys.Restapi.Policies.SdwanRules != nil {
						nestedRole.Vsys.Restapi.Policies.SdwanRules = o.Role.Vsys.Restapi.Policies.SdwanRules
					}
				}
				if o.Role.Vsys.Restapi.System != nil {
					nestedRole.Vsys.Restapi.System = &RoleVsysRestapiSystemXml{}
					if _, ok := o.Misc["RoleVsysRestapiSystem"]; ok {
						nestedRole.Vsys.Restapi.System.Misc = o.Misc["RoleVsysRestapiSystem"]
					}
					if o.Role.Vsys.Restapi.System.Configuration != nil {
						nestedRole.Vsys.Restapi.System.Configuration = o.Role.Vsys.Restapi.System.Configuration
					}
				}
			}
			if o.Role.Vsys.Webui != nil {
				nestedRole.Vsys.Webui = &RoleVsysWebuiXml{}
				if _, ok := o.Misc["RoleVsysWebui"]; ok {
					nestedRole.Vsys.Webui.Misc = o.Misc["RoleVsysWebui"]
				}
				if o.Role.Vsys.Webui.Network != nil {
					nestedRole.Vsys.Webui.Network = &RoleVsysWebuiNetworkXml{}
					if _, ok := o.Misc["RoleVsysWebuiNetwork"]; ok {
						nestedRole.Vsys.Webui.Network.Misc = o.Misc["RoleVsysWebuiNetwork"]
					}
					if o.Role.Vsys.Webui.Network.GlobalProtect != nil {
						nestedRole.Vsys.Webui.Network.GlobalProtect = &RoleVsysWebuiNetworkGlobalProtectXml{}
						if _, ok := o.Misc["RoleVsysWebuiNetworkGlobalProtect"]; ok {
							nestedRole.Vsys.Webui.Network.GlobalProtect.Misc = o.Misc["RoleVsysWebuiNetworkGlobalProtect"]
						}
						if o.Role.Vsys.Webui.Network.GlobalProtect.ClientlessAppGroups != nil {
							nestedRole.Vsys.Webui.Network.GlobalProtect.ClientlessAppGroups = o.Role.Vsys.Webui.Network.GlobalProtect.ClientlessAppGroups
						}
						if o.Role.Vsys.Webui.Network.GlobalProtect.ClientlessApps != nil {
							nestedRole.Vsys.Webui.Network.GlobalProtect.ClientlessApps = o.Role.Vsys.Webui.Network.GlobalProtect.ClientlessApps
						}
						if o.Role.Vsys.Webui.Network.GlobalProtect.Gateways != nil {
							nestedRole.Vsys.Webui.Network.GlobalProtect.Gateways = o.Role.Vsys.Webui.Network.GlobalProtect.Gateways
						}
						if o.Role.Vsys.Webui.Network.GlobalProtect.Mdm != nil {
							nestedRole.Vsys.Webui.Network.GlobalProtect.Mdm = o.Role.Vsys.Webui.Network.GlobalProtect.Mdm
						}
						if o.Role.Vsys.Webui.Network.GlobalProtect.Portals != nil {
							nestedRole.Vsys.Webui.Network.GlobalProtect.Portals = o.Role.Vsys.Webui.Network.GlobalProtect.Portals
						}
					}
					if o.Role.Vsys.Webui.Network.SdwanInterfaceProfile != nil {
						nestedRole.Vsys.Webui.Network.SdwanInterfaceProfile = o.Role.Vsys.Webui.Network.SdwanInterfaceProfile
					}
					if o.Role.Vsys.Webui.Network.Zones != nil {
						nestedRole.Vsys.Webui.Network.Zones = o.Role.Vsys.Webui.Network.Zones
					}
				}
				if o.Role.Vsys.Webui.Operations != nil {
					nestedRole.Vsys.Webui.Operations = &RoleVsysWebuiOperationsXml{}
					if _, ok := o.Misc["RoleVsysWebuiOperations"]; ok {
						nestedRole.Vsys.Webui.Operations.Misc = o.Misc["RoleVsysWebuiOperations"]
					}
					if o.Role.Vsys.Webui.Operations.GenerateStatsDumpFile != nil {
						nestedRole.Vsys.Webui.Operations.GenerateStatsDumpFile = o.Role.Vsys.Webui.Operations.GenerateStatsDumpFile
					}
					if o.Role.Vsys.Webui.Operations.GenerateTechSupportFile != nil {
						nestedRole.Vsys.Webui.Operations.GenerateTechSupportFile = o.Role.Vsys.Webui.Operations.GenerateTechSupportFile
					}
					if o.Role.Vsys.Webui.Operations.Reboot != nil {
						nestedRole.Vsys.Webui.Operations.Reboot = o.Role.Vsys.Webui.Operations.Reboot
					}
					if o.Role.Vsys.Webui.Operations.DownloadCoreFiles != nil {
						nestedRole.Vsys.Webui.Operations.DownloadCoreFiles = o.Role.Vsys.Webui.Operations.DownloadCoreFiles
					}
					if o.Role.Vsys.Webui.Operations.DownloadPcapFiles != nil {
						nestedRole.Vsys.Webui.Operations.DownloadPcapFiles = o.Role.Vsys.Webui.Operations.DownloadPcapFiles
					}
				}
				if o.Role.Vsys.Webui.Policies != nil {
					nestedRole.Vsys.Webui.Policies = &RoleVsysWebuiPoliciesXml{}
					if _, ok := o.Misc["RoleVsysWebuiPolicies"]; ok {
						nestedRole.Vsys.Webui.Policies.Misc = o.Misc["RoleVsysWebuiPolicies"]
					}
					if o.Role.Vsys.Webui.Policies.ApplicationOverrideRulebase != nil {
						nestedRole.Vsys.Webui.Policies.ApplicationOverrideRulebase = o.Role.Vsys.Webui.Policies.ApplicationOverrideRulebase
					}
					if o.Role.Vsys.Webui.Policies.AuthenticationRulebase != nil {
						nestedRole.Vsys.Webui.Policies.AuthenticationRulebase = o.Role.Vsys.Webui.Policies.AuthenticationRulebase
					}
					if o.Role.Vsys.Webui.Policies.DosRulebase != nil {
						nestedRole.Vsys.Webui.Policies.DosRulebase = o.Role.Vsys.Webui.Policies.DosRulebase
					}
					if o.Role.Vsys.Webui.Policies.PbfRulebase != nil {
						nestedRole.Vsys.Webui.Policies.PbfRulebase = o.Role.Vsys.Webui.Policies.PbfRulebase
					}
					if o.Role.Vsys.Webui.Policies.TunnelInspectRulebase != nil {
						nestedRole.Vsys.Webui.Policies.TunnelInspectRulebase = o.Role.Vsys.Webui.Policies.TunnelInspectRulebase
					}
					if o.Role.Vsys.Webui.Policies.SecurityRulebase != nil {
						nestedRole.Vsys.Webui.Policies.SecurityRulebase = o.Role.Vsys.Webui.Policies.SecurityRulebase
					}
					if o.Role.Vsys.Webui.Policies.SslDecryptionRulebase != nil {
						nestedRole.Vsys.Webui.Policies.SslDecryptionRulebase = o.Role.Vsys.Webui.Policies.SslDecryptionRulebase
					}
					if o.Role.Vsys.Webui.Policies.NatRulebase != nil {
						nestedRole.Vsys.Webui.Policies.NatRulebase = o.Role.Vsys.Webui.Policies.NatRulebase
					}
					if o.Role.Vsys.Webui.Policies.NetworkPacketBrokerRulebase != nil {
						nestedRole.Vsys.Webui.Policies.NetworkPacketBrokerRulebase = o.Role.Vsys.Webui.Policies.NetworkPacketBrokerRulebase
					}
					if o.Role.Vsys.Webui.Policies.QosRulebase != nil {
						nestedRole.Vsys.Webui.Policies.QosRulebase = o.Role.Vsys.Webui.Policies.QosRulebase
					}
					if o.Role.Vsys.Webui.Policies.RuleHitCountReset != nil {
						nestedRole.Vsys.Webui.Policies.RuleHitCountReset = o.Role.Vsys.Webui.Policies.RuleHitCountReset
					}
					if o.Role.Vsys.Webui.Policies.SdwanRulebase != nil {
						nestedRole.Vsys.Webui.Policies.SdwanRulebase = o.Role.Vsys.Webui.Policies.SdwanRulebase
					}
				}
				if o.Role.Vsys.Webui.Save != nil {
					nestedRole.Vsys.Webui.Save = &RoleVsysWebuiSaveXml{}
					if _, ok := o.Misc["RoleVsysWebuiSave"]; ok {
						nestedRole.Vsys.Webui.Save.Misc = o.Misc["RoleVsysWebuiSave"]
					}
					if o.Role.Vsys.Webui.Save.SaveForOtherAdmins != nil {
						nestedRole.Vsys.Webui.Save.SaveForOtherAdmins = o.Role.Vsys.Webui.Save.SaveForOtherAdmins
					}
					if o.Role.Vsys.Webui.Save.ObjectLevelChanges != nil {
						nestedRole.Vsys.Webui.Save.ObjectLevelChanges = o.Role.Vsys.Webui.Save.ObjectLevelChanges
					}
					if o.Role.Vsys.Webui.Save.PartialSave != nil {
						nestedRole.Vsys.Webui.Save.PartialSave = o.Role.Vsys.Webui.Save.PartialSave
					}
				}
				if o.Role.Vsys.Webui.Acc != nil {
					nestedRole.Vsys.Webui.Acc = o.Role.Vsys.Webui.Acc
				}
				if o.Role.Vsys.Webui.Dashboard != nil {
					nestedRole.Vsys.Webui.Dashboard = o.Role.Vsys.Webui.Dashboard
				}
				if o.Role.Vsys.Webui.Device != nil {
					nestedRole.Vsys.Webui.Device = &RoleVsysWebuiDeviceXml{}
					if _, ok := o.Misc["RoleVsysWebuiDevice"]; ok {
						nestedRole.Vsys.Webui.Device.Misc = o.Misc["RoleVsysWebuiDevice"]
					}
					if o.Role.Vsys.Webui.Device.Administrators != nil {
						nestedRole.Vsys.Webui.Device.Administrators = o.Role.Vsys.Webui.Device.Administrators
					}
					if o.Role.Vsys.Webui.Device.AuthenticationProfile != nil {
						nestedRole.Vsys.Webui.Device.AuthenticationProfile = o.Role.Vsys.Webui.Device.AuthenticationProfile
					}
					if o.Role.Vsys.Webui.Device.ServerProfile != nil {
						nestedRole.Vsys.Webui.Device.ServerProfile = &RoleVsysWebuiDeviceServerProfileXml{}
						if _, ok := o.Misc["RoleVsysWebuiDeviceServerProfile"]; ok {
							nestedRole.Vsys.Webui.Device.ServerProfile.Misc = o.Misc["RoleVsysWebuiDeviceServerProfile"]
						}
						if o.Role.Vsys.Webui.Device.ServerProfile.Radius != nil {
							nestedRole.Vsys.Webui.Device.ServerProfile.Radius = o.Role.Vsys.Webui.Device.ServerProfile.Radius
						}
						if o.Role.Vsys.Webui.Device.ServerProfile.SamlIdp != nil {
							nestedRole.Vsys.Webui.Device.ServerProfile.SamlIdp = o.Role.Vsys.Webui.Device.ServerProfile.SamlIdp
						}
						if o.Role.Vsys.Webui.Device.ServerProfile.SnmpTrap != nil {
							nestedRole.Vsys.Webui.Device.ServerProfile.SnmpTrap = o.Role.Vsys.Webui.Device.ServerProfile.SnmpTrap
						}
						if o.Role.Vsys.Webui.Device.ServerProfile.Dns != nil {
							nestedRole.Vsys.Webui.Device.ServerProfile.Dns = o.Role.Vsys.Webui.Device.ServerProfile.Dns
						}
						if o.Role.Vsys.Webui.Device.ServerProfile.Http != nil {
							nestedRole.Vsys.Webui.Device.ServerProfile.Http = o.Role.Vsys.Webui.Device.ServerProfile.Http
						}
						if o.Role.Vsys.Webui.Device.ServerProfile.Kerberos != nil {
							nestedRole.Vsys.Webui.Device.ServerProfile.Kerberos = o.Role.Vsys.Webui.Device.ServerProfile.Kerberos
						}
						if o.Role.Vsys.Webui.Device.ServerProfile.Netflow != nil {
							nestedRole.Vsys.Webui.Device.ServerProfile.Netflow = o.Role.Vsys.Webui.Device.ServerProfile.Netflow
						}
						if o.Role.Vsys.Webui.Device.ServerProfile.Syslog != nil {
							nestedRole.Vsys.Webui.Device.ServerProfile.Syslog = o.Role.Vsys.Webui.Device.ServerProfile.Syslog
						}
						if o.Role.Vsys.Webui.Device.ServerProfile.Tacplus != nil {
							nestedRole.Vsys.Webui.Device.ServerProfile.Tacplus = o.Role.Vsys.Webui.Device.ServerProfile.Tacplus
						}
						if o.Role.Vsys.Webui.Device.ServerProfile.Email != nil {
							nestedRole.Vsys.Webui.Device.ServerProfile.Email = o.Role.Vsys.Webui.Device.ServerProfile.Email
						}
						if o.Role.Vsys.Webui.Device.ServerProfile.Ldap != nil {
							nestedRole.Vsys.Webui.Device.ServerProfile.Ldap = o.Role.Vsys.Webui.Device.ServerProfile.Ldap
						}
						if o.Role.Vsys.Webui.Device.ServerProfile.Mfa != nil {
							nestedRole.Vsys.Webui.Device.ServerProfile.Mfa = o.Role.Vsys.Webui.Device.ServerProfile.Mfa
						}
						if o.Role.Vsys.Webui.Device.ServerProfile.Scp != nil {
							nestedRole.Vsys.Webui.Device.ServerProfile.Scp = o.Role.Vsys.Webui.Device.ServerProfile.Scp
						}
					}
					if o.Role.Vsys.Webui.Device.AuthenticationSequence != nil {
						nestedRole.Vsys.Webui.Device.AuthenticationSequence = o.Role.Vsys.Webui.Device.AuthenticationSequence
					}
					if o.Role.Vsys.Webui.Device.LocalUserDatabase != nil {
						nestedRole.Vsys.Webui.Device.LocalUserDatabase = &RoleVsysWebuiDeviceLocalUserDatabaseXml{}
						if _, ok := o.Misc["RoleVsysWebuiDeviceLocalUserDatabase"]; ok {
							nestedRole.Vsys.Webui.Device.LocalUserDatabase.Misc = o.Misc["RoleVsysWebuiDeviceLocalUserDatabase"]
						}
						if o.Role.Vsys.Webui.Device.LocalUserDatabase.UserGroups != nil {
							nestedRole.Vsys.Webui.Device.LocalUserDatabase.UserGroups = o.Role.Vsys.Webui.Device.LocalUserDatabase.UserGroups
						}
						if o.Role.Vsys.Webui.Device.LocalUserDatabase.Users != nil {
							nestedRole.Vsys.Webui.Device.LocalUserDatabase.Users = o.Role.Vsys.Webui.Device.LocalUserDatabase.Users
						}
					}
					if o.Role.Vsys.Webui.Device.Troubleshooting != nil {
						nestedRole.Vsys.Webui.Device.Troubleshooting = o.Role.Vsys.Webui.Device.Troubleshooting
					}
					if o.Role.Vsys.Webui.Device.VmInfoSource != nil {
						nestedRole.Vsys.Webui.Device.VmInfoSource = o.Role.Vsys.Webui.Device.VmInfoSource
					}
					if o.Role.Vsys.Webui.Device.BlockPages != nil {
						nestedRole.Vsys.Webui.Device.BlockPages = o.Role.Vsys.Webui.Device.BlockPages
					}
					if o.Role.Vsys.Webui.Device.CertificateManagement != nil {
						nestedRole.Vsys.Webui.Device.CertificateManagement = &RoleVsysWebuiDeviceCertificateManagementXml{}
						if _, ok := o.Misc["RoleVsysWebuiDeviceCertificateManagement"]; ok {
							nestedRole.Vsys.Webui.Device.CertificateManagement.Misc = o.Misc["RoleVsysWebuiDeviceCertificateManagement"]
						}
						if o.Role.Vsys.Webui.Device.CertificateManagement.SslDecryptionExclusion != nil {
							nestedRole.Vsys.Webui.Device.CertificateManagement.SslDecryptionExclusion = o.Role.Vsys.Webui.Device.CertificateManagement.SslDecryptionExclusion
						}
						if o.Role.Vsys.Webui.Device.CertificateManagement.SslTlsServiceProfile != nil {
							nestedRole.Vsys.Webui.Device.CertificateManagement.SslTlsServiceProfile = o.Role.Vsys.Webui.Device.CertificateManagement.SslTlsServiceProfile
						}
						if o.Role.Vsys.Webui.Device.CertificateManagement.CertificateProfile != nil {
							nestedRole.Vsys.Webui.Device.CertificateManagement.CertificateProfile = o.Role.Vsys.Webui.Device.CertificateManagement.CertificateProfile
						}
						if o.Role.Vsys.Webui.Device.CertificateManagement.Certificates != nil {
							nestedRole.Vsys.Webui.Device.CertificateManagement.Certificates = o.Role.Vsys.Webui.Device.CertificateManagement.Certificates
						}
						if o.Role.Vsys.Webui.Device.CertificateManagement.OcspResponder != nil {
							nestedRole.Vsys.Webui.Device.CertificateManagement.OcspResponder = o.Role.Vsys.Webui.Device.CertificateManagement.OcspResponder
						}
						if o.Role.Vsys.Webui.Device.CertificateManagement.Scep != nil {
							nestedRole.Vsys.Webui.Device.CertificateManagement.Scep = o.Role.Vsys.Webui.Device.CertificateManagement.Scep
						}
						if o.Role.Vsys.Webui.Device.CertificateManagement.SshServiceProfile != nil {
							nestedRole.Vsys.Webui.Device.CertificateManagement.SshServiceProfile = o.Role.Vsys.Webui.Device.CertificateManagement.SshServiceProfile
						}
					}
					if o.Role.Vsys.Webui.Device.DeviceQuarantine != nil {
						nestedRole.Vsys.Webui.Device.DeviceQuarantine = o.Role.Vsys.Webui.Device.DeviceQuarantine
					}
					if o.Role.Vsys.Webui.Device.LogSettings != nil {
						nestedRole.Vsys.Webui.Device.LogSettings = &RoleVsysWebuiDeviceLogSettingsXml{}
						if _, ok := o.Misc["RoleVsysWebuiDeviceLogSettings"]; ok {
							nestedRole.Vsys.Webui.Device.LogSettings.Misc = o.Misc["RoleVsysWebuiDeviceLogSettings"]
						}
						if o.Role.Vsys.Webui.Device.LogSettings.Hipmatch != nil {
							nestedRole.Vsys.Webui.Device.LogSettings.Hipmatch = o.Role.Vsys.Webui.Device.LogSettings.Hipmatch
						}
						if o.Role.Vsys.Webui.Device.LogSettings.Iptag != nil {
							nestedRole.Vsys.Webui.Device.LogSettings.Iptag = o.Role.Vsys.Webui.Device.LogSettings.Iptag
						}
						if o.Role.Vsys.Webui.Device.LogSettings.System != nil {
							nestedRole.Vsys.Webui.Device.LogSettings.System = o.Role.Vsys.Webui.Device.LogSettings.System
						}
						if o.Role.Vsys.Webui.Device.LogSettings.UserId != nil {
							nestedRole.Vsys.Webui.Device.LogSettings.UserId = o.Role.Vsys.Webui.Device.LogSettings.UserId
						}
						if o.Role.Vsys.Webui.Device.LogSettings.Config != nil {
							nestedRole.Vsys.Webui.Device.LogSettings.Config = o.Role.Vsys.Webui.Device.LogSettings.Config
						}
						if o.Role.Vsys.Webui.Device.LogSettings.Correlation != nil {
							nestedRole.Vsys.Webui.Device.LogSettings.Correlation = o.Role.Vsys.Webui.Device.LogSettings.Correlation
						}
						if o.Role.Vsys.Webui.Device.LogSettings.Globalprotect != nil {
							nestedRole.Vsys.Webui.Device.LogSettings.Globalprotect = o.Role.Vsys.Webui.Device.LogSettings.Globalprotect
						}
					}
					if o.Role.Vsys.Webui.Device.PolicyRecommendations != nil {
						nestedRole.Vsys.Webui.Device.PolicyRecommendations = &RoleVsysWebuiDevicePolicyRecommendationsXml{}
						if _, ok := o.Misc["RoleVsysWebuiDevicePolicyRecommendations"]; ok {
							nestedRole.Vsys.Webui.Device.PolicyRecommendations.Misc = o.Misc["RoleVsysWebuiDevicePolicyRecommendations"]
						}
						if o.Role.Vsys.Webui.Device.PolicyRecommendations.Iot != nil {
							nestedRole.Vsys.Webui.Device.PolicyRecommendations.Iot = o.Role.Vsys.Webui.Device.PolicyRecommendations.Iot
						}
						if o.Role.Vsys.Webui.Device.PolicyRecommendations.Saas != nil {
							nestedRole.Vsys.Webui.Device.PolicyRecommendations.Saas = o.Role.Vsys.Webui.Device.PolicyRecommendations.Saas
						}
					}
					if o.Role.Vsys.Webui.Device.Setup != nil {
						nestedRole.Vsys.Webui.Device.Setup = &RoleVsysWebuiDeviceSetupXml{}
						if _, ok := o.Misc["RoleVsysWebuiDeviceSetup"]; ok {
							nestedRole.Vsys.Webui.Device.Setup.Misc = o.Misc["RoleVsysWebuiDeviceSetup"]
						}
						if o.Role.Vsys.Webui.Device.Setup.ContentId != nil {
							nestedRole.Vsys.Webui.Device.Setup.ContentId = o.Role.Vsys.Webui.Device.Setup.ContentId
						}
						if o.Role.Vsys.Webui.Device.Setup.Hsm != nil {
							nestedRole.Vsys.Webui.Device.Setup.Hsm = o.Role.Vsys.Webui.Device.Setup.Hsm
						}
						if o.Role.Vsys.Webui.Device.Setup.Management != nil {
							nestedRole.Vsys.Webui.Device.Setup.Management = o.Role.Vsys.Webui.Device.Setup.Management
						}
						if o.Role.Vsys.Webui.Device.Setup.Services != nil {
							nestedRole.Vsys.Webui.Device.Setup.Services = o.Role.Vsys.Webui.Device.Setup.Services
						}
						if o.Role.Vsys.Webui.Device.Setup.Interfaces != nil {
							nestedRole.Vsys.Webui.Device.Setup.Interfaces = o.Role.Vsys.Webui.Device.Setup.Interfaces
						}
						if o.Role.Vsys.Webui.Device.Setup.Operations != nil {
							nestedRole.Vsys.Webui.Device.Setup.Operations = o.Role.Vsys.Webui.Device.Setup.Operations
						}
						if o.Role.Vsys.Webui.Device.Setup.Session != nil {
							nestedRole.Vsys.Webui.Device.Setup.Session = o.Role.Vsys.Webui.Device.Setup.Session
						}
						if o.Role.Vsys.Webui.Device.Setup.Telemetry != nil {
							nestedRole.Vsys.Webui.Device.Setup.Telemetry = o.Role.Vsys.Webui.Device.Setup.Telemetry
						}
						if o.Role.Vsys.Webui.Device.Setup.Wildfire != nil {
							nestedRole.Vsys.Webui.Device.Setup.Wildfire = o.Role.Vsys.Webui.Device.Setup.Wildfire
						}
					}
					if o.Role.Vsys.Webui.Device.DataRedistribution != nil {
						nestedRole.Vsys.Webui.Device.DataRedistribution = o.Role.Vsys.Webui.Device.DataRedistribution
					}
					if o.Role.Vsys.Webui.Device.UserIdentification != nil {
						nestedRole.Vsys.Webui.Device.UserIdentification = o.Role.Vsys.Webui.Device.UserIdentification
					}
					if o.Role.Vsys.Webui.Device.DhcpSyslogServer != nil {
						nestedRole.Vsys.Webui.Device.DhcpSyslogServer = o.Role.Vsys.Webui.Device.DhcpSyslogServer
					}
				}
				if o.Role.Vsys.Webui.Monitor != nil {
					nestedRole.Vsys.Webui.Monitor = &RoleVsysWebuiMonitorXml{}
					if _, ok := o.Misc["RoleVsysWebuiMonitor"]; ok {
						nestedRole.Vsys.Webui.Monitor.Misc = o.Misc["RoleVsysWebuiMonitor"]
					}
					if o.Role.Vsys.Webui.Monitor.AutomatedCorrelationEngine != nil {
						nestedRole.Vsys.Webui.Monitor.AutomatedCorrelationEngine = &RoleVsysWebuiMonitorAutomatedCorrelationEngineXml{}
						if _, ok := o.Misc["RoleVsysWebuiMonitorAutomatedCorrelationEngine"]; ok {
							nestedRole.Vsys.Webui.Monitor.AutomatedCorrelationEngine.Misc = o.Misc["RoleVsysWebuiMonitorAutomatedCorrelationEngine"]
						}
						if o.Role.Vsys.Webui.Monitor.AutomatedCorrelationEngine.CorrelatedEvents != nil {
							nestedRole.Vsys.Webui.Monitor.AutomatedCorrelationEngine.CorrelatedEvents = o.Role.Vsys.Webui.Monitor.AutomatedCorrelationEngine.CorrelatedEvents
						}
						if o.Role.Vsys.Webui.Monitor.AutomatedCorrelationEngine.CorrelationObjects != nil {
							nestedRole.Vsys.Webui.Monitor.AutomatedCorrelationEngine.CorrelationObjects = o.Role.Vsys.Webui.Monitor.AutomatedCorrelationEngine.CorrelationObjects
						}
					}
					if o.Role.Vsys.Webui.Monitor.BlockIpList != nil {
						nestedRole.Vsys.Webui.Monitor.BlockIpList = o.Role.Vsys.Webui.Monitor.BlockIpList
					}
					if o.Role.Vsys.Webui.Monitor.ExternalLogs != nil {
						nestedRole.Vsys.Webui.Monitor.ExternalLogs = o.Role.Vsys.Webui.Monitor.ExternalLogs
					}
					if o.Role.Vsys.Webui.Monitor.PdfReports != nil {
						nestedRole.Vsys.Webui.Monitor.PdfReports = &RoleVsysWebuiMonitorPdfReportsXml{}
						if _, ok := o.Misc["RoleVsysWebuiMonitorPdfReports"]; ok {
							nestedRole.Vsys.Webui.Monitor.PdfReports.Misc = o.Misc["RoleVsysWebuiMonitorPdfReports"]
						}
						if o.Role.Vsys.Webui.Monitor.PdfReports.SaasApplicationUsageReport != nil {
							nestedRole.Vsys.Webui.Monitor.PdfReports.SaasApplicationUsageReport = o.Role.Vsys.Webui.Monitor.PdfReports.SaasApplicationUsageReport
						}
						if o.Role.Vsys.Webui.Monitor.PdfReports.UserActivityReport != nil {
							nestedRole.Vsys.Webui.Monitor.PdfReports.UserActivityReport = o.Role.Vsys.Webui.Monitor.PdfReports.UserActivityReport
						}
						if o.Role.Vsys.Webui.Monitor.PdfReports.EmailScheduler != nil {
							nestedRole.Vsys.Webui.Monitor.PdfReports.EmailScheduler = o.Role.Vsys.Webui.Monitor.PdfReports.EmailScheduler
						}
						if o.Role.Vsys.Webui.Monitor.PdfReports.ManagePdfSummary != nil {
							nestedRole.Vsys.Webui.Monitor.PdfReports.ManagePdfSummary = o.Role.Vsys.Webui.Monitor.PdfReports.ManagePdfSummary
						}
						if o.Role.Vsys.Webui.Monitor.PdfReports.PdfSummaryReports != nil {
							nestedRole.Vsys.Webui.Monitor.PdfReports.PdfSummaryReports = o.Role.Vsys.Webui.Monitor.PdfReports.PdfSummaryReports
						}
						if o.Role.Vsys.Webui.Monitor.PdfReports.ReportGroups != nil {
							nestedRole.Vsys.Webui.Monitor.PdfReports.ReportGroups = o.Role.Vsys.Webui.Monitor.PdfReports.ReportGroups
						}
					}
					if o.Role.Vsys.Webui.Monitor.AppScope != nil {
						nestedRole.Vsys.Webui.Monitor.AppScope = o.Role.Vsys.Webui.Monitor.AppScope
					}
					if o.Role.Vsys.Webui.Monitor.CustomReports != nil {
						nestedRole.Vsys.Webui.Monitor.CustomReports = &RoleVsysWebuiMonitorCustomReportsXml{}
						if _, ok := o.Misc["RoleVsysWebuiMonitorCustomReports"]; ok {
							nestedRole.Vsys.Webui.Monitor.CustomReports.Misc = o.Misc["RoleVsysWebuiMonitorCustomReports"]
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.SctpSummary != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.SctpSummary = o.Role.Vsys.Webui.Monitor.CustomReports.SctpSummary
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.TrafficLog != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.TrafficLog = o.Role.Vsys.Webui.Monitor.CustomReports.TrafficLog
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.Userid != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.Userid = o.Role.Vsys.Webui.Monitor.CustomReports.Userid
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.GtpLog != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.GtpLog = o.Role.Vsys.Webui.Monitor.CustomReports.GtpLog
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.GtpSummary != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.GtpSummary = o.Role.Vsys.Webui.Monitor.CustomReports.GtpSummary
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.SctpLog != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.SctpLog = o.Role.Vsys.Webui.Monitor.CustomReports.SctpLog
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.Hipmatch != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.Hipmatch = o.Role.Vsys.Webui.Monitor.CustomReports.Hipmatch
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.Iptag != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.Iptag = o.Role.Vsys.Webui.Monitor.CustomReports.Iptag
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.ThreatSummary != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.ThreatSummary = o.Role.Vsys.Webui.Monitor.CustomReports.ThreatSummary
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.TrafficSummary != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.TrafficSummary = o.Role.Vsys.Webui.Monitor.CustomReports.TrafficSummary
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.UrlLog != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.UrlLog = o.Role.Vsys.Webui.Monitor.CustomReports.UrlLog
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.DataFilteringLog != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.DataFilteringLog = o.Role.Vsys.Webui.Monitor.CustomReports.DataFilteringLog
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.DecryptionSummary != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.DecryptionSummary = o.Role.Vsys.Webui.Monitor.CustomReports.DecryptionSummary
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.Globalprotect != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.Globalprotect = o.Role.Vsys.Webui.Monitor.CustomReports.Globalprotect
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.UrlSummary != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.UrlSummary = o.Role.Vsys.Webui.Monitor.CustomReports.UrlSummary
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.WildfireLog != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.WildfireLog = o.Role.Vsys.Webui.Monitor.CustomReports.WildfireLog
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.DecryptionLog != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.DecryptionLog = o.Role.Vsys.Webui.Monitor.CustomReports.DecryptionLog
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.ThreatLog != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.ThreatLog = o.Role.Vsys.Webui.Monitor.CustomReports.ThreatLog
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.TunnelSummary != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.TunnelSummary = o.Role.Vsys.Webui.Monitor.CustomReports.TunnelSummary
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.ApplicationStatistics != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.ApplicationStatistics = o.Role.Vsys.Webui.Monitor.CustomReports.ApplicationStatistics
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.Auth != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.Auth = o.Role.Vsys.Webui.Monitor.CustomReports.Auth
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports.TunnelLog != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports.TunnelLog = o.Role.Vsys.Webui.Monitor.CustomReports.TunnelLog
						}
					}
					if o.Role.Vsys.Webui.Monitor.Logs != nil {
						nestedRole.Vsys.Webui.Monitor.Logs = &RoleVsysWebuiMonitorLogsXml{}
						if _, ok := o.Misc["RoleVsysWebuiMonitorLogs"]; ok {
							nestedRole.Vsys.Webui.Monitor.Logs.Misc = o.Misc["RoleVsysWebuiMonitorLogs"]
						}
						if o.Role.Vsys.Webui.Monitor.Logs.Threat != nil {
							nestedRole.Vsys.Webui.Monitor.Logs.Threat = o.Role.Vsys.Webui.Monitor.Logs.Threat
						}
						if o.Role.Vsys.Webui.Monitor.Logs.Wildfire != nil {
							nestedRole.Vsys.Webui.Monitor.Logs.Wildfire = o.Role.Vsys.Webui.Monitor.Logs.Wildfire
						}
						if o.Role.Vsys.Webui.Monitor.Logs.DataFiltering != nil {
							nestedRole.Vsys.Webui.Monitor.Logs.DataFiltering = o.Role.Vsys.Webui.Monitor.Logs.DataFiltering
						}
						if o.Role.Vsys.Webui.Monitor.Logs.Globalprotect != nil {
							nestedRole.Vsys.Webui.Monitor.Logs.Globalprotect = o.Role.Vsys.Webui.Monitor.Logs.Globalprotect
						}
						if o.Role.Vsys.Webui.Monitor.Logs.Hipmatch != nil {
							nestedRole.Vsys.Webui.Monitor.Logs.Hipmatch = o.Role.Vsys.Webui.Monitor.Logs.Hipmatch
						}
						if o.Role.Vsys.Webui.Monitor.Logs.Url != nil {
							nestedRole.Vsys.Webui.Monitor.Logs.Url = o.Role.Vsys.Webui.Monitor.Logs.Url
						}
						if o.Role.Vsys.Webui.Monitor.Logs.Userid != nil {
							nestedRole.Vsys.Webui.Monitor.Logs.Userid = o.Role.Vsys.Webui.Monitor.Logs.Userid
						}
						if o.Role.Vsys.Webui.Monitor.Logs.Authentication != nil {
							nestedRole.Vsys.Webui.Monitor.Logs.Authentication = o.Role.Vsys.Webui.Monitor.Logs.Authentication
						}
						if o.Role.Vsys.Webui.Monitor.Logs.Decryption != nil {
							nestedRole.Vsys.Webui.Monitor.Logs.Decryption = o.Role.Vsys.Webui.Monitor.Logs.Decryption
						}
						if o.Role.Vsys.Webui.Monitor.Logs.Gtp != nil {
							nestedRole.Vsys.Webui.Monitor.Logs.Gtp = o.Role.Vsys.Webui.Monitor.Logs.Gtp
						}
						if o.Role.Vsys.Webui.Monitor.Logs.Iptag != nil {
							nestedRole.Vsys.Webui.Monitor.Logs.Iptag = o.Role.Vsys.Webui.Monitor.Logs.Iptag
						}
						if o.Role.Vsys.Webui.Monitor.Logs.Traffic != nil {
							nestedRole.Vsys.Webui.Monitor.Logs.Traffic = o.Role.Vsys.Webui.Monitor.Logs.Traffic
						}
						if o.Role.Vsys.Webui.Monitor.Logs.Tunnel != nil {
							nestedRole.Vsys.Webui.Monitor.Logs.Tunnel = o.Role.Vsys.Webui.Monitor.Logs.Tunnel
						}
						if o.Role.Vsys.Webui.Monitor.Logs.Sctp != nil {
							nestedRole.Vsys.Webui.Monitor.Logs.Sctp = o.Role.Vsys.Webui.Monitor.Logs.Sctp
						}
					}
					if o.Role.Vsys.Webui.Monitor.SessionBrowser != nil {
						nestedRole.Vsys.Webui.Monitor.SessionBrowser = o.Role.Vsys.Webui.Monitor.SessionBrowser
					}
					if o.Role.Vsys.Webui.Monitor.ViewCustomReports != nil {
						nestedRole.Vsys.Webui.Monitor.ViewCustomReports = o.Role.Vsys.Webui.Monitor.ViewCustomReports
					}
				}
				if o.Role.Vsys.Webui.Tasks != nil {
					nestedRole.Vsys.Webui.Tasks = o.Role.Vsys.Webui.Tasks
				}
				if o.Role.Vsys.Webui.Validate != nil {
					nestedRole.Vsys.Webui.Validate = o.Role.Vsys.Webui.Validate
				}
				if o.Role.Vsys.Webui.Commit != nil {
					nestedRole.Vsys.Webui.Commit = &RoleVsysWebuiCommitXml{}
					if _, ok := o.Misc["RoleVsysWebuiCommit"]; ok {
						nestedRole.Vsys.Webui.Commit.Misc = o.Misc["RoleVsysWebuiCommit"]
					}
					if o.Role.Vsys.Webui.Commit.CommitForOtherAdmins != nil {
						nestedRole.Vsys.Webui.Commit.CommitForOtherAdmins = o.Role.Vsys.Webui.Commit.CommitForOtherAdmins
					}
					if o.Role.Vsys.Webui.Commit.VirtualSystems != nil {
						nestedRole.Vsys.Webui.Commit.VirtualSystems = o.Role.Vsys.Webui.Commit.VirtualSystems
					}
				}
				if o.Role.Vsys.Webui.Objects != nil {
					nestedRole.Vsys.Webui.Objects = &RoleVsysWebuiObjectsXml{}
					if _, ok := o.Misc["RoleVsysWebuiObjects"]; ok {
						nestedRole.Vsys.Webui.Objects.Misc = o.Misc["RoleVsysWebuiObjects"]
					}
					if o.Role.Vsys.Webui.Objects.Services != nil {
						nestedRole.Vsys.Webui.Objects.Services = o.Role.Vsys.Webui.Objects.Services
					}
					if o.Role.Vsys.Webui.Objects.Tags != nil {
						nestedRole.Vsys.Webui.Objects.Tags = o.Role.Vsys.Webui.Objects.Tags
					}
					if o.Role.Vsys.Webui.Objects.AddressGroups != nil {
						nestedRole.Vsys.Webui.Objects.AddressGroups = o.Role.Vsys.Webui.Objects.AddressGroups
					}
					if o.Role.Vsys.Webui.Objects.Addresses != nil {
						nestedRole.Vsys.Webui.Objects.Addresses = o.Role.Vsys.Webui.Objects.Addresses
					}
					if o.Role.Vsys.Webui.Objects.ApplicationGroups != nil {
						nestedRole.Vsys.Webui.Objects.ApplicationGroups = o.Role.Vsys.Webui.Objects.ApplicationGroups
					}
					if o.Role.Vsys.Webui.Objects.ServiceGroups != nil {
						nestedRole.Vsys.Webui.Objects.ServiceGroups = o.Role.Vsys.Webui.Objects.ServiceGroups
					}
					if o.Role.Vsys.Webui.Objects.SecurityProfiles != nil {
						nestedRole.Vsys.Webui.Objects.SecurityProfiles = &RoleVsysWebuiObjectsSecurityProfilesXml{}
						if _, ok := o.Misc["RoleVsysWebuiObjectsSecurityProfiles"]; ok {
							nestedRole.Vsys.Webui.Objects.SecurityProfiles.Misc = o.Misc["RoleVsysWebuiObjectsSecurityProfiles"]
						}
						if o.Role.Vsys.Webui.Objects.SecurityProfiles.AntiSpyware != nil {
							nestedRole.Vsys.Webui.Objects.SecurityProfiles.AntiSpyware = o.Role.Vsys.Webui.Objects.SecurityProfiles.AntiSpyware
						}
						if o.Role.Vsys.Webui.Objects.SecurityProfiles.DataFiltering != nil {
							nestedRole.Vsys.Webui.Objects.SecurityProfiles.DataFiltering = o.Role.Vsys.Webui.Objects.SecurityProfiles.DataFiltering
						}
						if o.Role.Vsys.Webui.Objects.SecurityProfiles.DosProtection != nil {
							nestedRole.Vsys.Webui.Objects.SecurityProfiles.DosProtection = o.Role.Vsys.Webui.Objects.SecurityProfiles.DosProtection
						}
						if o.Role.Vsys.Webui.Objects.SecurityProfiles.UrlFiltering != nil {
							nestedRole.Vsys.Webui.Objects.SecurityProfiles.UrlFiltering = o.Role.Vsys.Webui.Objects.SecurityProfiles.UrlFiltering
						}
						if o.Role.Vsys.Webui.Objects.SecurityProfiles.Antivirus != nil {
							nestedRole.Vsys.Webui.Objects.SecurityProfiles.Antivirus = o.Role.Vsys.Webui.Objects.SecurityProfiles.Antivirus
						}
						if o.Role.Vsys.Webui.Objects.SecurityProfiles.FileBlocking != nil {
							nestedRole.Vsys.Webui.Objects.SecurityProfiles.FileBlocking = o.Role.Vsys.Webui.Objects.SecurityProfiles.FileBlocking
						}
						if o.Role.Vsys.Webui.Objects.SecurityProfiles.GtpProtection != nil {
							nestedRole.Vsys.Webui.Objects.SecurityProfiles.GtpProtection = o.Role.Vsys.Webui.Objects.SecurityProfiles.GtpProtection
						}
						if o.Role.Vsys.Webui.Objects.SecurityProfiles.SctpProtection != nil {
							nestedRole.Vsys.Webui.Objects.SecurityProfiles.SctpProtection = o.Role.Vsys.Webui.Objects.SecurityProfiles.SctpProtection
						}
						if o.Role.Vsys.Webui.Objects.SecurityProfiles.VulnerabilityProtection != nil {
							nestedRole.Vsys.Webui.Objects.SecurityProfiles.VulnerabilityProtection = o.Role.Vsys.Webui.Objects.SecurityProfiles.VulnerabilityProtection
						}
						if o.Role.Vsys.Webui.Objects.SecurityProfiles.WildfireAnalysis != nil {
							nestedRole.Vsys.Webui.Objects.SecurityProfiles.WildfireAnalysis = o.Role.Vsys.Webui.Objects.SecurityProfiles.WildfireAnalysis
						}
					}
					if o.Role.Vsys.Webui.Objects.Decryption != nil {
						nestedRole.Vsys.Webui.Objects.Decryption = &RoleVsysWebuiObjectsDecryptionXml{}
						if _, ok := o.Misc["RoleVsysWebuiObjectsDecryption"]; ok {
							nestedRole.Vsys.Webui.Objects.Decryption.Misc = o.Misc["RoleVsysWebuiObjectsDecryption"]
						}
						if o.Role.Vsys.Webui.Objects.Decryption.DecryptionProfile != nil {
							nestedRole.Vsys.Webui.Objects.Decryption.DecryptionProfile = o.Role.Vsys.Webui.Objects.Decryption.DecryptionProfile
						}
					}
					if o.Role.Vsys.Webui.Objects.Devices != nil {
						nestedRole.Vsys.Webui.Objects.Devices = o.Role.Vsys.Webui.Objects.Devices
					}
					if o.Role.Vsys.Webui.Objects.DynamicUserGroups != nil {
						nestedRole.Vsys.Webui.Objects.DynamicUserGroups = o.Role.Vsys.Webui.Objects.DynamicUserGroups
					}
					if o.Role.Vsys.Webui.Objects.Schedules != nil {
						nestedRole.Vsys.Webui.Objects.Schedules = o.Role.Vsys.Webui.Objects.Schedules
					}
					if o.Role.Vsys.Webui.Objects.Applications != nil {
						nestedRole.Vsys.Webui.Objects.Applications = o.Role.Vsys.Webui.Objects.Applications
					}
					if o.Role.Vsys.Webui.Objects.Regions != nil {
						nestedRole.Vsys.Webui.Objects.Regions = o.Role.Vsys.Webui.Objects.Regions
					}
					if o.Role.Vsys.Webui.Objects.Sdwan != nil {
						nestedRole.Vsys.Webui.Objects.Sdwan = &RoleVsysWebuiObjectsSdwanXml{}
						if _, ok := o.Misc["RoleVsysWebuiObjectsSdwan"]; ok {
							nestedRole.Vsys.Webui.Objects.Sdwan.Misc = o.Misc["RoleVsysWebuiObjectsSdwan"]
						}
						if o.Role.Vsys.Webui.Objects.Sdwan.SdwanErrorCorrectionProfile != nil {
							nestedRole.Vsys.Webui.Objects.Sdwan.SdwanErrorCorrectionProfile = o.Role.Vsys.Webui.Objects.Sdwan.SdwanErrorCorrectionProfile
						}
						if o.Role.Vsys.Webui.Objects.Sdwan.SdwanProfile != nil {
							nestedRole.Vsys.Webui.Objects.Sdwan.SdwanProfile = o.Role.Vsys.Webui.Objects.Sdwan.SdwanProfile
						}
						if o.Role.Vsys.Webui.Objects.Sdwan.SdwanSaasQualityProfile != nil {
							nestedRole.Vsys.Webui.Objects.Sdwan.SdwanSaasQualityProfile = o.Role.Vsys.Webui.Objects.Sdwan.SdwanSaasQualityProfile
						}
						if o.Role.Vsys.Webui.Objects.Sdwan.SdwanDistProfile != nil {
							nestedRole.Vsys.Webui.Objects.Sdwan.SdwanDistProfile = o.Role.Vsys.Webui.Objects.Sdwan.SdwanDistProfile
						}
					}
					if o.Role.Vsys.Webui.Objects.SecurityProfileGroups != nil {
						nestedRole.Vsys.Webui.Objects.SecurityProfileGroups = o.Role.Vsys.Webui.Objects.SecurityProfileGroups
					}
					if o.Role.Vsys.Webui.Objects.GlobalProtect != nil {
						nestedRole.Vsys.Webui.Objects.GlobalProtect = &RoleVsysWebuiObjectsGlobalProtectXml{}
						if _, ok := o.Misc["RoleVsysWebuiObjectsGlobalProtect"]; ok {
							nestedRole.Vsys.Webui.Objects.GlobalProtect.Misc = o.Misc["RoleVsysWebuiObjectsGlobalProtect"]
						}
						if o.Role.Vsys.Webui.Objects.GlobalProtect.HipObjects != nil {
							nestedRole.Vsys.Webui.Objects.GlobalProtect.HipObjects = o.Role.Vsys.Webui.Objects.GlobalProtect.HipObjects
						}
						if o.Role.Vsys.Webui.Objects.GlobalProtect.HipProfiles != nil {
							nestedRole.Vsys.Webui.Objects.GlobalProtect.HipProfiles = o.Role.Vsys.Webui.Objects.GlobalProtect.HipProfiles
						}
					}
					if o.Role.Vsys.Webui.Objects.LogForwarding != nil {
						nestedRole.Vsys.Webui.Objects.LogForwarding = o.Role.Vsys.Webui.Objects.LogForwarding
					}
					if o.Role.Vsys.Webui.Objects.PacketBrokerProfile != nil {
						nestedRole.Vsys.Webui.Objects.PacketBrokerProfile = o.Role.Vsys.Webui.Objects.PacketBrokerProfile
					}
					if o.Role.Vsys.Webui.Objects.ApplicationFilters != nil {
						nestedRole.Vsys.Webui.Objects.ApplicationFilters = o.Role.Vsys.Webui.Objects.ApplicationFilters
					}
					if o.Role.Vsys.Webui.Objects.Authentication != nil {
						nestedRole.Vsys.Webui.Objects.Authentication = o.Role.Vsys.Webui.Objects.Authentication
					}
					if o.Role.Vsys.Webui.Objects.CustomObjects != nil {
						nestedRole.Vsys.Webui.Objects.CustomObjects = &RoleVsysWebuiObjectsCustomObjectsXml{}
						if _, ok := o.Misc["RoleVsysWebuiObjectsCustomObjects"]; ok {
							nestedRole.Vsys.Webui.Objects.CustomObjects.Misc = o.Misc["RoleVsysWebuiObjectsCustomObjects"]
						}
						if o.Role.Vsys.Webui.Objects.CustomObjects.DataPatterns != nil {
							nestedRole.Vsys.Webui.Objects.CustomObjects.DataPatterns = o.Role.Vsys.Webui.Objects.CustomObjects.DataPatterns
						}
						if o.Role.Vsys.Webui.Objects.CustomObjects.Spyware != nil {
							nestedRole.Vsys.Webui.Objects.CustomObjects.Spyware = o.Role.Vsys.Webui.Objects.CustomObjects.Spyware
						}
						if o.Role.Vsys.Webui.Objects.CustomObjects.UrlCategory != nil {
							nestedRole.Vsys.Webui.Objects.CustomObjects.UrlCategory = o.Role.Vsys.Webui.Objects.CustomObjects.UrlCategory
						}
						if o.Role.Vsys.Webui.Objects.CustomObjects.Vulnerability != nil {
							nestedRole.Vsys.Webui.Objects.CustomObjects.Vulnerability = o.Role.Vsys.Webui.Objects.CustomObjects.Vulnerability
						}
					}
					if o.Role.Vsys.Webui.Objects.DynamicBlockLists != nil {
						nestedRole.Vsys.Webui.Objects.DynamicBlockLists = o.Role.Vsys.Webui.Objects.DynamicBlockLists
					}
				}
				if o.Role.Vsys.Webui.Privacy != nil {
					nestedRole.Vsys.Webui.Privacy = &RoleVsysWebuiPrivacyXml{}
					if _, ok := o.Misc["RoleVsysWebuiPrivacy"]; ok {
						nestedRole.Vsys.Webui.Privacy.Misc = o.Misc["RoleVsysWebuiPrivacy"]
					}
					if o.Role.Vsys.Webui.Privacy.ShowFullIpAddresses != nil {
						nestedRole.Vsys.Webui.Privacy.ShowFullIpAddresses = o.Role.Vsys.Webui.Privacy.ShowFullIpAddresses
					}
					if o.Role.Vsys.Webui.Privacy.ShowUserNamesInLogsAndReports != nil {
						nestedRole.Vsys.Webui.Privacy.ShowUserNamesInLogsAndReports = o.Role.Vsys.Webui.Privacy.ShowUserNamesInLogsAndReports
					}
					if o.Role.Vsys.Webui.Privacy.ViewPcapFiles != nil {
						nestedRole.Vsys.Webui.Privacy.ViewPcapFiles = o.Role.Vsys.Webui.Privacy.ViewPcapFiles
					}
				}
			}
			if o.Role.Vsys.Xmlapi != nil {
				nestedRole.Vsys.Xmlapi = &RoleVsysXmlapiXml{}
				if _, ok := o.Misc["RoleVsysXmlapi"]; ok {
					nestedRole.Vsys.Xmlapi.Misc = o.Misc["RoleVsysXmlapi"]
				}
				if o.Role.Vsys.Xmlapi.Commit != nil {
					nestedRole.Vsys.Xmlapi.Commit = o.Role.Vsys.Xmlapi.Commit
				}
				if o.Role.Vsys.Xmlapi.Config != nil {
					nestedRole.Vsys.Xmlapi.Config = o.Role.Vsys.Xmlapi.Config
				}
				if o.Role.Vsys.Xmlapi.Import != nil {
					nestedRole.Vsys.Xmlapi.Import = o.Role.Vsys.Xmlapi.Import
				}
				if o.Role.Vsys.Xmlapi.Iot != nil {
					nestedRole.Vsys.Xmlapi.Iot = o.Role.Vsys.Xmlapi.Iot
				}
				if o.Role.Vsys.Xmlapi.Log != nil {
					nestedRole.Vsys.Xmlapi.Log = o.Role.Vsys.Xmlapi.Log
				}
				if o.Role.Vsys.Xmlapi.Op != nil {
					nestedRole.Vsys.Xmlapi.Op = o.Role.Vsys.Xmlapi.Op
				}
				if o.Role.Vsys.Xmlapi.Export != nil {
					nestedRole.Vsys.Xmlapi.Export = o.Role.Vsys.Xmlapi.Export
				}
				if o.Role.Vsys.Xmlapi.Report != nil {
					nestedRole.Vsys.Xmlapi.Report = o.Role.Vsys.Xmlapi.Report
				}
				if o.Role.Vsys.Xmlapi.UserId != nil {
					nestedRole.Vsys.Xmlapi.UserId = o.Role.Vsys.Xmlapi.UserId
				}
			}
		}
	}
	entry.Role = nestedRole

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
		entry.Description = o.Description
		var nestedRole *Role
		if o.Role != nil {
			nestedRole = &Role{}
			if o.Role.Misc != nil {
				entry.Misc["Role"] = o.Role.Misc
			}
			if o.Role.Device != nil {
				nestedRole.Device = &RoleDevice{}
				if o.Role.Device.Misc != nil {
					entry.Misc["RoleDevice"] = o.Role.Device.Misc
				}
				if o.Role.Device.Cli != nil {
					nestedRole.Device.Cli = o.Role.Device.Cli
				}
				if o.Role.Device.Restapi != nil {
					nestedRole.Device.Restapi = &RoleDeviceRestapi{}
					if o.Role.Device.Restapi.Misc != nil {
						entry.Misc["RoleDeviceRestapi"] = o.Role.Device.Restapi.Misc
					}
					if o.Role.Device.Restapi.Objects != nil {
						nestedRole.Device.Restapi.Objects = &RoleDeviceRestapiObjects{}
						if o.Role.Device.Restapi.Objects.Misc != nil {
							entry.Misc["RoleDeviceRestapiObjects"] = o.Role.Device.Restapi.Objects.Misc
						}
						if o.Role.Device.Restapi.Objects.ExternalDynamicLists != nil {
							nestedRole.Device.Restapi.Objects.ExternalDynamicLists = o.Role.Device.Restapi.Objects.ExternalDynamicLists
						}
						if o.Role.Device.Restapi.Objects.SecurityProfileGroups != nil {
							nestedRole.Device.Restapi.Objects.SecurityProfileGroups = o.Role.Device.Restapi.Objects.SecurityProfileGroups
						}
						if o.Role.Device.Restapi.Objects.AddressGroups != nil {
							nestedRole.Device.Restapi.Objects.AddressGroups = o.Role.Device.Restapi.Objects.AddressGroups
						}
						if o.Role.Device.Restapi.Objects.AntivirusSecurityProfiles != nil {
							nestedRole.Device.Restapi.Objects.AntivirusSecurityProfiles = o.Role.Device.Restapi.Objects.AntivirusSecurityProfiles
						}
						if o.Role.Device.Restapi.Objects.AuthenticationEnforcements != nil {
							nestedRole.Device.Restapi.Objects.AuthenticationEnforcements = o.Role.Device.Restapi.Objects.AuthenticationEnforcements
						}
						if o.Role.Device.Restapi.Objects.FileBlockingSecurityProfiles != nil {
							nestedRole.Device.Restapi.Objects.FileBlockingSecurityProfiles = o.Role.Device.Restapi.Objects.FileBlockingSecurityProfiles
						}
						if o.Role.Device.Restapi.Objects.PacketBrokerProfiles != nil {
							nestedRole.Device.Restapi.Objects.PacketBrokerProfiles = o.Role.Device.Restapi.Objects.PacketBrokerProfiles
						}
						if o.Role.Device.Restapi.Objects.Schedules != nil {
							nestedRole.Device.Restapi.Objects.Schedules = o.Role.Device.Restapi.Objects.Schedules
						}
						if o.Role.Device.Restapi.Objects.ServiceGroups != nil {
							nestedRole.Device.Restapi.Objects.ServiceGroups = o.Role.Device.Restapi.Objects.ServiceGroups
						}
						if o.Role.Device.Restapi.Objects.Addresses != nil {
							nestedRole.Device.Restapi.Objects.Addresses = o.Role.Device.Restapi.Objects.Addresses
						}
						if o.Role.Device.Restapi.Objects.ApplicationFilters != nil {
							nestedRole.Device.Restapi.Objects.ApplicationFilters = o.Role.Device.Restapi.Objects.ApplicationFilters
						}
						if o.Role.Device.Restapi.Objects.CustomDataPatterns != nil {
							nestedRole.Device.Restapi.Objects.CustomDataPatterns = o.Role.Device.Restapi.Objects.CustomDataPatterns
						}
						if o.Role.Device.Restapi.Objects.DataFilteringSecurityProfiles != nil {
							nestedRole.Device.Restapi.Objects.DataFilteringSecurityProfiles = o.Role.Device.Restapi.Objects.DataFilteringSecurityProfiles
						}
						if o.Role.Device.Restapi.Objects.GlobalprotectHipProfiles != nil {
							nestedRole.Device.Restapi.Objects.GlobalprotectHipProfiles = o.Role.Device.Restapi.Objects.GlobalprotectHipProfiles
						}
						if o.Role.Device.Restapi.Objects.LogForwardingProfiles != nil {
							nestedRole.Device.Restapi.Objects.LogForwardingProfiles = o.Role.Device.Restapi.Objects.LogForwardingProfiles
						}
						if o.Role.Device.Restapi.Objects.SctpProtectionSecurityProfiles != nil {
							nestedRole.Device.Restapi.Objects.SctpProtectionSecurityProfiles = o.Role.Device.Restapi.Objects.SctpProtectionSecurityProfiles
						}
						if o.Role.Device.Restapi.Objects.SdwanSaasQualityProfiles != nil {
							nestedRole.Device.Restapi.Objects.SdwanSaasQualityProfiles = o.Role.Device.Restapi.Objects.SdwanSaasQualityProfiles
						}
						if o.Role.Device.Restapi.Objects.AntiSpywareSecurityProfiles != nil {
							nestedRole.Device.Restapi.Objects.AntiSpywareSecurityProfiles = o.Role.Device.Restapi.Objects.AntiSpywareSecurityProfiles
						}
						if o.Role.Device.Restapi.Objects.ApplicationGroups != nil {
							nestedRole.Device.Restapi.Objects.ApplicationGroups = o.Role.Device.Restapi.Objects.ApplicationGroups
						}
						if o.Role.Device.Restapi.Objects.WildfireAnalysisSecurityProfiles != nil {
							nestedRole.Device.Restapi.Objects.WildfireAnalysisSecurityProfiles = o.Role.Device.Restapi.Objects.WildfireAnalysisSecurityProfiles
						}
						if o.Role.Device.Restapi.Objects.DynamicUserGroups != nil {
							nestedRole.Device.Restapi.Objects.DynamicUserGroups = o.Role.Device.Restapi.Objects.DynamicUserGroups
						}
						if o.Role.Device.Restapi.Objects.UrlFilteringSecurityProfiles != nil {
							nestedRole.Device.Restapi.Objects.UrlFilteringSecurityProfiles = o.Role.Device.Restapi.Objects.UrlFilteringSecurityProfiles
						}
						if o.Role.Device.Restapi.Objects.CustomSpywareSignatures != nil {
							nestedRole.Device.Restapi.Objects.CustomSpywareSignatures = o.Role.Device.Restapi.Objects.CustomSpywareSignatures
						}
						if o.Role.Device.Restapi.Objects.DecryptionProfiles != nil {
							nestedRole.Device.Restapi.Objects.DecryptionProfiles = o.Role.Device.Restapi.Objects.DecryptionProfiles
						}
						if o.Role.Device.Restapi.Objects.SdwanPathQualityProfiles != nil {
							nestedRole.Device.Restapi.Objects.SdwanPathQualityProfiles = o.Role.Device.Restapi.Objects.SdwanPathQualityProfiles
						}
						if o.Role.Device.Restapi.Objects.SdwanTrafficDistributionProfiles != nil {
							nestedRole.Device.Restapi.Objects.SdwanTrafficDistributionProfiles = o.Role.Device.Restapi.Objects.SdwanTrafficDistributionProfiles
						}
						if o.Role.Device.Restapi.Objects.GlobalprotectHipObjects != nil {
							nestedRole.Device.Restapi.Objects.GlobalprotectHipObjects = o.Role.Device.Restapi.Objects.GlobalprotectHipObjects
						}
						if o.Role.Device.Restapi.Objects.SdwanErrorCorrectionProfiles != nil {
							nestedRole.Device.Restapi.Objects.SdwanErrorCorrectionProfiles = o.Role.Device.Restapi.Objects.SdwanErrorCorrectionProfiles
						}
						if o.Role.Device.Restapi.Objects.Services != nil {
							nestedRole.Device.Restapi.Objects.Services = o.Role.Device.Restapi.Objects.Services
						}
						if o.Role.Device.Restapi.Objects.Tags != nil {
							nestedRole.Device.Restapi.Objects.Tags = o.Role.Device.Restapi.Objects.Tags
						}
						if o.Role.Device.Restapi.Objects.VulnerabilityProtectionSecurityProfiles != nil {
							nestedRole.Device.Restapi.Objects.VulnerabilityProtectionSecurityProfiles = o.Role.Device.Restapi.Objects.VulnerabilityProtectionSecurityProfiles
						}
						if o.Role.Device.Restapi.Objects.Applications != nil {
							nestedRole.Device.Restapi.Objects.Applications = o.Role.Device.Restapi.Objects.Applications
						}
						if o.Role.Device.Restapi.Objects.CustomVulnerabilitySignatures != nil {
							nestedRole.Device.Restapi.Objects.CustomVulnerabilitySignatures = o.Role.Device.Restapi.Objects.CustomVulnerabilitySignatures
						}
						if o.Role.Device.Restapi.Objects.DosProtectionSecurityProfiles != nil {
							nestedRole.Device.Restapi.Objects.DosProtectionSecurityProfiles = o.Role.Device.Restapi.Objects.DosProtectionSecurityProfiles
						}
						if o.Role.Device.Restapi.Objects.GtpProtectionSecurityProfiles != nil {
							nestedRole.Device.Restapi.Objects.GtpProtectionSecurityProfiles = o.Role.Device.Restapi.Objects.GtpProtectionSecurityProfiles
						}
						if o.Role.Device.Restapi.Objects.Regions != nil {
							nestedRole.Device.Restapi.Objects.Regions = o.Role.Device.Restapi.Objects.Regions
						}
						if o.Role.Device.Restapi.Objects.CustomUrlCategories != nil {
							nestedRole.Device.Restapi.Objects.CustomUrlCategories = o.Role.Device.Restapi.Objects.CustomUrlCategories
						}
						if o.Role.Device.Restapi.Objects.Devices != nil {
							nestedRole.Device.Restapi.Objects.Devices = o.Role.Device.Restapi.Objects.Devices
						}
					}
					if o.Role.Device.Restapi.Policies != nil {
						nestedRole.Device.Restapi.Policies = &RoleDeviceRestapiPolicies{}
						if o.Role.Device.Restapi.Policies.Misc != nil {
							entry.Misc["RoleDeviceRestapiPolicies"] = o.Role.Device.Restapi.Policies.Misc
						}
						if o.Role.Device.Restapi.Policies.AuthenticationRules != nil {
							nestedRole.Device.Restapi.Policies.AuthenticationRules = o.Role.Device.Restapi.Policies.AuthenticationRules
						}
						if o.Role.Device.Restapi.Policies.DosRules != nil {
							nestedRole.Device.Restapi.Policies.DosRules = o.Role.Device.Restapi.Policies.DosRules
						}
						if o.Role.Device.Restapi.Policies.NatRules != nil {
							nestedRole.Device.Restapi.Policies.NatRules = o.Role.Device.Restapi.Policies.NatRules
						}
						if o.Role.Device.Restapi.Policies.NetworkPacketBrokerRules != nil {
							nestedRole.Device.Restapi.Policies.NetworkPacketBrokerRules = o.Role.Device.Restapi.Policies.NetworkPacketBrokerRules
						}
						if o.Role.Device.Restapi.Policies.TunnelInspectionRules != nil {
							nestedRole.Device.Restapi.Policies.TunnelInspectionRules = o.Role.Device.Restapi.Policies.TunnelInspectionRules
						}
						if o.Role.Device.Restapi.Policies.SecurityRules != nil {
							nestedRole.Device.Restapi.Policies.SecurityRules = o.Role.Device.Restapi.Policies.SecurityRules
						}
						if o.Role.Device.Restapi.Policies.ApplicationOverrideRules != nil {
							nestedRole.Device.Restapi.Policies.ApplicationOverrideRules = o.Role.Device.Restapi.Policies.ApplicationOverrideRules
						}
						if o.Role.Device.Restapi.Policies.DecryptionRules != nil {
							nestedRole.Device.Restapi.Policies.DecryptionRules = o.Role.Device.Restapi.Policies.DecryptionRules
						}
						if o.Role.Device.Restapi.Policies.PolicyBasedForwardingRules != nil {
							nestedRole.Device.Restapi.Policies.PolicyBasedForwardingRules = o.Role.Device.Restapi.Policies.PolicyBasedForwardingRules
						}
						if o.Role.Device.Restapi.Policies.QosRules != nil {
							nestedRole.Device.Restapi.Policies.QosRules = o.Role.Device.Restapi.Policies.QosRules
						}
						if o.Role.Device.Restapi.Policies.SdwanRules != nil {
							nestedRole.Device.Restapi.Policies.SdwanRules = o.Role.Device.Restapi.Policies.SdwanRules
						}
					}
					if o.Role.Device.Restapi.System != nil {
						nestedRole.Device.Restapi.System = &RoleDeviceRestapiSystem{}
						if o.Role.Device.Restapi.System.Misc != nil {
							entry.Misc["RoleDeviceRestapiSystem"] = o.Role.Device.Restapi.System.Misc
						}
						if o.Role.Device.Restapi.System.Configuration != nil {
							nestedRole.Device.Restapi.System.Configuration = o.Role.Device.Restapi.System.Configuration
						}
					}
					if o.Role.Device.Restapi.Device != nil {
						nestedRole.Device.Restapi.Device = &RoleDeviceRestapiDevice{}
						if o.Role.Device.Restapi.Device.Misc != nil {
							entry.Misc["RoleDeviceRestapiDevice"] = o.Role.Device.Restapi.Device.Misc
						}
						if o.Role.Device.Restapi.Device.LogInterfaceSetting != nil {
							nestedRole.Device.Restapi.Device.LogInterfaceSetting = o.Role.Device.Restapi.Device.LogInterfaceSetting
						}
						if o.Role.Device.Restapi.Device.SnmpTrapServerProfiles != nil {
							nestedRole.Device.Restapi.Device.SnmpTrapServerProfiles = o.Role.Device.Restapi.Device.SnmpTrapServerProfiles
						}
						if o.Role.Device.Restapi.Device.SyslogServerProfiles != nil {
							nestedRole.Device.Restapi.Device.SyslogServerProfiles = o.Role.Device.Restapi.Device.SyslogServerProfiles
						}
						if o.Role.Device.Restapi.Device.VirtualSystems != nil {
							nestedRole.Device.Restapi.Device.VirtualSystems = o.Role.Device.Restapi.Device.VirtualSystems
						}
						if o.Role.Device.Restapi.Device.EmailServerProfiles != nil {
							nestedRole.Device.Restapi.Device.EmailServerProfiles = o.Role.Device.Restapi.Device.EmailServerProfiles
						}
						if o.Role.Device.Restapi.Device.HttpServerProfiles != nil {
							nestedRole.Device.Restapi.Device.HttpServerProfiles = o.Role.Device.Restapi.Device.HttpServerProfiles
						}
						if o.Role.Device.Restapi.Device.LdapServerProfiles != nil {
							nestedRole.Device.Restapi.Device.LdapServerProfiles = o.Role.Device.Restapi.Device.LdapServerProfiles
						}
					}
					if o.Role.Device.Restapi.Network != nil {
						nestedRole.Device.Restapi.Network = &RoleDeviceRestapiNetwork{}
						if o.Role.Device.Restapi.Network.Misc != nil {
							entry.Misc["RoleDeviceRestapiNetwork"] = o.Role.Device.Restapi.Network.Misc
						}
						if o.Role.Device.Restapi.Network.Lldp != nil {
							nestedRole.Device.Restapi.Network.Lldp = o.Role.Device.Restapi.Network.Lldp
						}
						if o.Role.Device.Restapi.Network.TunnelMonitorNetworkProfiles != nil {
							nestedRole.Device.Restapi.Network.TunnelMonitorNetworkProfiles = o.Role.Device.Restapi.Network.TunnelMonitorNetworkProfiles
						}
						if o.Role.Device.Restapi.Network.BgpRoutingProfiles != nil {
							nestedRole.Device.Restapi.Network.BgpRoutingProfiles = o.Role.Device.Restapi.Network.BgpRoutingProfiles
						}
						if o.Role.Device.Restapi.Network.GlobalprotectPortals != nil {
							nestedRole.Device.Restapi.Network.GlobalprotectPortals = o.Role.Device.Restapi.Network.GlobalprotectPortals
						}
						if o.Role.Device.Restapi.Network.IkeGatewayNetworkProfiles != nil {
							nestedRole.Device.Restapi.Network.IkeGatewayNetworkProfiles = o.Role.Device.Restapi.Network.IkeGatewayNetworkProfiles
						}
						if o.Role.Device.Restapi.Network.IpsecCryptoNetworkProfiles != nil {
							nestedRole.Device.Restapi.Network.IpsecCryptoNetworkProfiles = o.Role.Device.Restapi.Network.IpsecCryptoNetworkProfiles
						}
						if o.Role.Device.Restapi.Network.QosNetworkProfiles != nil {
							nestedRole.Device.Restapi.Network.QosNetworkProfiles = o.Role.Device.Restapi.Network.QosNetworkProfiles
						}
						if o.Role.Device.Restapi.Network.DhcpServers != nil {
							nestedRole.Device.Restapi.Network.DhcpServers = o.Role.Device.Restapi.Network.DhcpServers
						}
						if o.Role.Device.Restapi.Network.DnsProxies != nil {
							nestedRole.Device.Restapi.Network.DnsProxies = o.Role.Device.Restapi.Network.DnsProxies
						}
						if o.Role.Device.Restapi.Network.EthernetInterfaces != nil {
							nestedRole.Device.Restapi.Network.EthernetInterfaces = o.Role.Device.Restapi.Network.EthernetInterfaces
						}
						if o.Role.Device.Restapi.Network.GlobalprotectClientlessApps != nil {
							nestedRole.Device.Restapi.Network.GlobalprotectClientlessApps = o.Role.Device.Restapi.Network.GlobalprotectClientlessApps
						}
						if o.Role.Device.Restapi.Network.SdwanInterfaceProfiles != nil {
							nestedRole.Device.Restapi.Network.SdwanInterfaceProfiles = o.Role.Device.Restapi.Network.SdwanInterfaceProfiles
						}
						if o.Role.Device.Restapi.Network.DhcpRelays != nil {
							nestedRole.Device.Restapi.Network.DhcpRelays = o.Role.Device.Restapi.Network.DhcpRelays
						}
						if o.Role.Device.Restapi.Network.LogicalRouters != nil {
							nestedRole.Device.Restapi.Network.LogicalRouters = o.Role.Device.Restapi.Network.LogicalRouters
						}
						if o.Role.Device.Restapi.Network.Vlans != nil {
							nestedRole.Device.Restapi.Network.Vlans = o.Role.Device.Restapi.Network.Vlans
						}
						if o.Role.Device.Restapi.Network.AggregateEthernetInterfaces != nil {
							nestedRole.Device.Restapi.Network.AggregateEthernetInterfaces = o.Role.Device.Restapi.Network.AggregateEthernetInterfaces
						}
						if o.Role.Device.Restapi.Network.GlobalprotectIpsecCryptoNetworkProfiles != nil {
							nestedRole.Device.Restapi.Network.GlobalprotectIpsecCryptoNetworkProfiles = o.Role.Device.Restapi.Network.GlobalprotectIpsecCryptoNetworkProfiles
						}
						if o.Role.Device.Restapi.Network.LoopbackInterfaces != nil {
							nestedRole.Device.Restapi.Network.LoopbackInterfaces = o.Role.Device.Restapi.Network.LoopbackInterfaces
						}
						if o.Role.Device.Restapi.Network.QosInterfaces != nil {
							nestedRole.Device.Restapi.Network.QosInterfaces = o.Role.Device.Restapi.Network.QosInterfaces
						}
						if o.Role.Device.Restapi.Network.SdwanInterfaces != nil {
							nestedRole.Device.Restapi.Network.SdwanInterfaces = o.Role.Device.Restapi.Network.SdwanInterfaces
						}
						if o.Role.Device.Restapi.Network.VlanInterfaces != nil {
							nestedRole.Device.Restapi.Network.VlanInterfaces = o.Role.Device.Restapi.Network.VlanInterfaces
						}
						if o.Role.Device.Restapi.Network.Zones != nil {
							nestedRole.Device.Restapi.Network.Zones = o.Role.Device.Restapi.Network.Zones
						}
						if o.Role.Device.Restapi.Network.BfdNetworkProfiles != nil {
							nestedRole.Device.Restapi.Network.BfdNetworkProfiles = o.Role.Device.Restapi.Network.BfdNetworkProfiles
						}
						if o.Role.Device.Restapi.Network.GreTunnels != nil {
							nestedRole.Device.Restapi.Network.GreTunnels = o.Role.Device.Restapi.Network.GreTunnels
						}
						if o.Role.Device.Restapi.Network.IkeCryptoNetworkProfiles != nil {
							nestedRole.Device.Restapi.Network.IkeCryptoNetworkProfiles = o.Role.Device.Restapi.Network.IkeCryptoNetworkProfiles
						}
						if o.Role.Device.Restapi.Network.IpsecTunnels != nil {
							nestedRole.Device.Restapi.Network.IpsecTunnels = o.Role.Device.Restapi.Network.IpsecTunnels
						}
						if o.Role.Device.Restapi.Network.LldpNetworkProfiles != nil {
							nestedRole.Device.Restapi.Network.LldpNetworkProfiles = o.Role.Device.Restapi.Network.LldpNetworkProfiles
						}
						if o.Role.Device.Restapi.Network.GlobalprotectClientlessAppGroups != nil {
							nestedRole.Device.Restapi.Network.GlobalprotectClientlessAppGroups = o.Role.Device.Restapi.Network.GlobalprotectClientlessAppGroups
						}
						if o.Role.Device.Restapi.Network.GlobalprotectGateways != nil {
							nestedRole.Device.Restapi.Network.GlobalprotectGateways = o.Role.Device.Restapi.Network.GlobalprotectGateways
						}
						if o.Role.Device.Restapi.Network.GlobalprotectMdmServers != nil {
							nestedRole.Device.Restapi.Network.GlobalprotectMdmServers = o.Role.Device.Restapi.Network.GlobalprotectMdmServers
						}
						if o.Role.Device.Restapi.Network.InterfaceManagementNetworkProfiles != nil {
							nestedRole.Device.Restapi.Network.InterfaceManagementNetworkProfiles = o.Role.Device.Restapi.Network.InterfaceManagementNetworkProfiles
						}
						if o.Role.Device.Restapi.Network.TunnelInterfaces != nil {
							nestedRole.Device.Restapi.Network.TunnelInterfaces = o.Role.Device.Restapi.Network.TunnelInterfaces
						}
						if o.Role.Device.Restapi.Network.VirtualRouters != nil {
							nestedRole.Device.Restapi.Network.VirtualRouters = o.Role.Device.Restapi.Network.VirtualRouters
						}
						if o.Role.Device.Restapi.Network.VirtualWires != nil {
							nestedRole.Device.Restapi.Network.VirtualWires = o.Role.Device.Restapi.Network.VirtualWires
						}
						if o.Role.Device.Restapi.Network.ZoneProtectionNetworkProfiles != nil {
							nestedRole.Device.Restapi.Network.ZoneProtectionNetworkProfiles = o.Role.Device.Restapi.Network.ZoneProtectionNetworkProfiles
						}
					}
				}
				if o.Role.Device.Webui != nil {
					nestedRole.Device.Webui = &RoleDeviceWebui{}
					if o.Role.Device.Webui.Misc != nil {
						entry.Misc["RoleDeviceWebui"] = o.Role.Device.Webui.Misc
					}
					if o.Role.Device.Webui.Commit != nil {
						nestedRole.Device.Webui.Commit = &RoleDeviceWebuiCommit{}
						if o.Role.Device.Webui.Commit.Misc != nil {
							entry.Misc["RoleDeviceWebuiCommit"] = o.Role.Device.Webui.Commit.Misc
						}
						if o.Role.Device.Webui.Commit.CommitForOtherAdmins != nil {
							nestedRole.Device.Webui.Commit.CommitForOtherAdmins = o.Role.Device.Webui.Commit.CommitForOtherAdmins
						}
						if o.Role.Device.Webui.Commit.Device != nil {
							nestedRole.Device.Webui.Commit.Device = o.Role.Device.Webui.Commit.Device
						}
						if o.Role.Device.Webui.Commit.ObjectLevelChanges != nil {
							nestedRole.Device.Webui.Commit.ObjectLevelChanges = o.Role.Device.Webui.Commit.ObjectLevelChanges
						}
					}
					if o.Role.Device.Webui.Device != nil {
						nestedRole.Device.Webui.Device = &RoleDeviceWebuiDevice{}
						if o.Role.Device.Webui.Device.Misc != nil {
							entry.Misc["RoleDeviceWebuiDevice"] = o.Role.Device.Webui.Device.Misc
						}
						if o.Role.Device.Webui.Device.GlobalProtectClient != nil {
							nestedRole.Device.Webui.Device.GlobalProtectClient = o.Role.Device.Webui.Device.GlobalProtectClient
						}
						if o.Role.Device.Webui.Device.BlockPages != nil {
							nestedRole.Device.Webui.Device.BlockPages = o.Role.Device.Webui.Device.BlockPages
						}
						if o.Role.Device.Webui.Device.Administrators != nil {
							nestedRole.Device.Webui.Device.Administrators = o.Role.Device.Webui.Device.Administrators
						}
						if o.Role.Device.Webui.Device.AuthenticationSequence != nil {
							nestedRole.Device.Webui.Device.AuthenticationSequence = o.Role.Device.Webui.Device.AuthenticationSequence
						}
						if o.Role.Device.Webui.Device.AccessDomain != nil {
							nestedRole.Device.Webui.Device.AccessDomain = o.Role.Device.Webui.Device.AccessDomain
						}
						if o.Role.Device.Webui.Device.Setup != nil {
							nestedRole.Device.Webui.Device.Setup = &RoleDeviceWebuiDeviceSetup{}
							if o.Role.Device.Webui.Device.Setup.Misc != nil {
								entry.Misc["RoleDeviceWebuiDeviceSetup"] = o.Role.Device.Webui.Device.Setup.Misc
							}
							if o.Role.Device.Webui.Device.Setup.ContentId != nil {
								nestedRole.Device.Webui.Device.Setup.ContentId = o.Role.Device.Webui.Device.Setup.ContentId
							}
							if o.Role.Device.Webui.Device.Setup.Management != nil {
								nestedRole.Device.Webui.Device.Setup.Management = o.Role.Device.Webui.Device.Setup.Management
							}
							if o.Role.Device.Webui.Device.Setup.Services != nil {
								nestedRole.Device.Webui.Device.Setup.Services = o.Role.Device.Webui.Device.Setup.Services
							}
							if o.Role.Device.Webui.Device.Setup.Session != nil {
								nestedRole.Device.Webui.Device.Setup.Session = o.Role.Device.Webui.Device.Setup.Session
							}
							if o.Role.Device.Webui.Device.Setup.Wildfire != nil {
								nestedRole.Device.Webui.Device.Setup.Wildfire = o.Role.Device.Webui.Device.Setup.Wildfire
							}
							if o.Role.Device.Webui.Device.Setup.Hsm != nil {
								nestedRole.Device.Webui.Device.Setup.Hsm = o.Role.Device.Webui.Device.Setup.Hsm
							}
							if o.Role.Device.Webui.Device.Setup.Interfaces != nil {
								nestedRole.Device.Webui.Device.Setup.Interfaces = o.Role.Device.Webui.Device.Setup.Interfaces
							}
							if o.Role.Device.Webui.Device.Setup.Operations != nil {
								nestedRole.Device.Webui.Device.Setup.Operations = o.Role.Device.Webui.Device.Setup.Operations
							}
							if o.Role.Device.Webui.Device.Setup.Telemetry != nil {
								nestedRole.Device.Webui.Device.Setup.Telemetry = o.Role.Device.Webui.Device.Setup.Telemetry
							}
						}
						if o.Role.Device.Webui.Device.ScheduledLogExport != nil {
							nestedRole.Device.Webui.Device.ScheduledLogExport = o.Role.Device.Webui.Device.ScheduledLogExport
						}
						if o.Role.Device.Webui.Device.Software != nil {
							nestedRole.Device.Webui.Device.Software = o.Role.Device.Webui.Device.Software
						}
						if o.Role.Device.Webui.Device.Troubleshooting != nil {
							nestedRole.Device.Webui.Device.Troubleshooting = o.Role.Device.Webui.Device.Troubleshooting
						}
						if o.Role.Device.Webui.Device.UserIdentification != nil {
							nestedRole.Device.Webui.Device.UserIdentification = o.Role.Device.Webui.Device.UserIdentification
						}
						if o.Role.Device.Webui.Device.DhcpSyslogServer != nil {
							nestedRole.Device.Webui.Device.DhcpSyslogServer = o.Role.Device.Webui.Device.DhcpSyslogServer
						}
						if o.Role.Device.Webui.Device.CertificateManagement != nil {
							nestedRole.Device.Webui.Device.CertificateManagement = &RoleDeviceWebuiDeviceCertificateManagement{}
							if o.Role.Device.Webui.Device.CertificateManagement.Misc != nil {
								entry.Misc["RoleDeviceWebuiDeviceCertificateManagement"] = o.Role.Device.Webui.Device.CertificateManagement.Misc
							}
							if o.Role.Device.Webui.Device.CertificateManagement.Certificates != nil {
								nestedRole.Device.Webui.Device.CertificateManagement.Certificates = o.Role.Device.Webui.Device.CertificateManagement.Certificates
							}
							if o.Role.Device.Webui.Device.CertificateManagement.OcspResponder != nil {
								nestedRole.Device.Webui.Device.CertificateManagement.OcspResponder = o.Role.Device.Webui.Device.CertificateManagement.OcspResponder
							}
							if o.Role.Device.Webui.Device.CertificateManagement.Scep != nil {
								nestedRole.Device.Webui.Device.CertificateManagement.Scep = o.Role.Device.Webui.Device.CertificateManagement.Scep
							}
							if o.Role.Device.Webui.Device.CertificateManagement.SshServiceProfile != nil {
								nestedRole.Device.Webui.Device.CertificateManagement.SshServiceProfile = o.Role.Device.Webui.Device.CertificateManagement.SshServiceProfile
							}
							if o.Role.Device.Webui.Device.CertificateManagement.SslDecryptionExclusion != nil {
								nestedRole.Device.Webui.Device.CertificateManagement.SslDecryptionExclusion = o.Role.Device.Webui.Device.CertificateManagement.SslDecryptionExclusion
							}
							if o.Role.Device.Webui.Device.CertificateManagement.SslTlsServiceProfile != nil {
								nestedRole.Device.Webui.Device.CertificateManagement.SslTlsServiceProfile = o.Role.Device.Webui.Device.CertificateManagement.SslTlsServiceProfile
							}
							if o.Role.Device.Webui.Device.CertificateManagement.CertificateProfile != nil {
								nestedRole.Device.Webui.Device.CertificateManagement.CertificateProfile = o.Role.Device.Webui.Device.CertificateManagement.CertificateProfile
							}
						}
						if o.Role.Device.Webui.Device.LogFwdCard != nil {
							nestedRole.Device.Webui.Device.LogFwdCard = o.Role.Device.Webui.Device.LogFwdCard
						}
						if o.Role.Device.Webui.Device.SharedGateways != nil {
							nestedRole.Device.Webui.Device.SharedGateways = o.Role.Device.Webui.Device.SharedGateways
						}
						if o.Role.Device.Webui.Device.DeviceQuarantine != nil {
							nestedRole.Device.Webui.Device.DeviceQuarantine = o.Role.Device.Webui.Device.DeviceQuarantine
						}
						if o.Role.Device.Webui.Device.DataRedistribution != nil {
							nestedRole.Device.Webui.Device.DataRedistribution = o.Role.Device.Webui.Device.DataRedistribution
						}
						if o.Role.Device.Webui.Device.HighAvailability != nil {
							nestedRole.Device.Webui.Device.HighAvailability = o.Role.Device.Webui.Device.HighAvailability
						}
						if o.Role.Device.Webui.Device.LocalUserDatabase != nil {
							nestedRole.Device.Webui.Device.LocalUserDatabase = &RoleDeviceWebuiDeviceLocalUserDatabase{}
							if o.Role.Device.Webui.Device.LocalUserDatabase.Misc != nil {
								entry.Misc["RoleDeviceWebuiDeviceLocalUserDatabase"] = o.Role.Device.Webui.Device.LocalUserDatabase.Misc
							}
							if o.Role.Device.Webui.Device.LocalUserDatabase.UserGroups != nil {
								nestedRole.Device.Webui.Device.LocalUserDatabase.UserGroups = o.Role.Device.Webui.Device.LocalUserDatabase.UserGroups
							}
							if o.Role.Device.Webui.Device.LocalUserDatabase.Users != nil {
								nestedRole.Device.Webui.Device.LocalUserDatabase.Users = o.Role.Device.Webui.Device.LocalUserDatabase.Users
							}
						}
						if o.Role.Device.Webui.Device.MasterKey != nil {
							nestedRole.Device.Webui.Device.MasterKey = o.Role.Device.Webui.Device.MasterKey
						}
						if o.Role.Device.Webui.Device.Plugins != nil {
							nestedRole.Device.Webui.Device.Plugins = o.Role.Device.Webui.Device.Plugins
						}
						if o.Role.Device.Webui.Device.PolicyRecommendations != nil {
							nestedRole.Device.Webui.Device.PolicyRecommendations = &RoleDeviceWebuiDevicePolicyRecommendations{}
							if o.Role.Device.Webui.Device.PolicyRecommendations.Misc != nil {
								entry.Misc["RoleDeviceWebuiDevicePolicyRecommendations"] = o.Role.Device.Webui.Device.PolicyRecommendations.Misc
							}
							if o.Role.Device.Webui.Device.PolicyRecommendations.Iot != nil {
								nestedRole.Device.Webui.Device.PolicyRecommendations.Iot = o.Role.Device.Webui.Device.PolicyRecommendations.Iot
							}
							if o.Role.Device.Webui.Device.PolicyRecommendations.Saas != nil {
								nestedRole.Device.Webui.Device.PolicyRecommendations.Saas = o.Role.Device.Webui.Device.PolicyRecommendations.Saas
							}
						}
						if o.Role.Device.Webui.Device.ServerProfile != nil {
							nestedRole.Device.Webui.Device.ServerProfile = &RoleDeviceWebuiDeviceServerProfile{}
							if o.Role.Device.Webui.Device.ServerProfile.Misc != nil {
								entry.Misc["RoleDeviceWebuiDeviceServerProfile"] = o.Role.Device.Webui.Device.ServerProfile.Misc
							}
							if o.Role.Device.Webui.Device.ServerProfile.Dns != nil {
								nestedRole.Device.Webui.Device.ServerProfile.Dns = o.Role.Device.Webui.Device.ServerProfile.Dns
							}
							if o.Role.Device.Webui.Device.ServerProfile.Email != nil {
								nestedRole.Device.Webui.Device.ServerProfile.Email = o.Role.Device.Webui.Device.ServerProfile.Email
							}
							if o.Role.Device.Webui.Device.ServerProfile.Netflow != nil {
								nestedRole.Device.Webui.Device.ServerProfile.Netflow = o.Role.Device.Webui.Device.ServerProfile.Netflow
							}
							if o.Role.Device.Webui.Device.ServerProfile.SamlIdp != nil {
								nestedRole.Device.Webui.Device.ServerProfile.SamlIdp = o.Role.Device.Webui.Device.ServerProfile.SamlIdp
							}
							if o.Role.Device.Webui.Device.ServerProfile.Tacplus != nil {
								nestedRole.Device.Webui.Device.ServerProfile.Tacplus = o.Role.Device.Webui.Device.ServerProfile.Tacplus
							}
							if o.Role.Device.Webui.Device.ServerProfile.Syslog != nil {
								nestedRole.Device.Webui.Device.ServerProfile.Syslog = o.Role.Device.Webui.Device.ServerProfile.Syslog
							}
							if o.Role.Device.Webui.Device.ServerProfile.Http != nil {
								nestedRole.Device.Webui.Device.ServerProfile.Http = o.Role.Device.Webui.Device.ServerProfile.Http
							}
							if o.Role.Device.Webui.Device.ServerProfile.Kerberos != nil {
								nestedRole.Device.Webui.Device.ServerProfile.Kerberos = o.Role.Device.Webui.Device.ServerProfile.Kerberos
							}
							if o.Role.Device.Webui.Device.ServerProfile.Ldap != nil {
								nestedRole.Device.Webui.Device.ServerProfile.Ldap = o.Role.Device.Webui.Device.ServerProfile.Ldap
							}
							if o.Role.Device.Webui.Device.ServerProfile.Mfa != nil {
								nestedRole.Device.Webui.Device.ServerProfile.Mfa = o.Role.Device.Webui.Device.ServerProfile.Mfa
							}
							if o.Role.Device.Webui.Device.ServerProfile.Radius != nil {
								nestedRole.Device.Webui.Device.ServerProfile.Radius = o.Role.Device.Webui.Device.ServerProfile.Radius
							}
							if o.Role.Device.Webui.Device.ServerProfile.Scp != nil {
								nestedRole.Device.Webui.Device.ServerProfile.Scp = o.Role.Device.Webui.Device.ServerProfile.Scp
							}
							if o.Role.Device.Webui.Device.ServerProfile.SnmpTrap != nil {
								nestedRole.Device.Webui.Device.ServerProfile.SnmpTrap = o.Role.Device.Webui.Device.ServerProfile.SnmpTrap
							}
						}
						if o.Role.Device.Webui.Device.ConfigAudit != nil {
							nestedRole.Device.Webui.Device.ConfigAudit = o.Role.Device.Webui.Device.ConfigAudit
						}
						if o.Role.Device.Webui.Device.Support != nil {
							nestedRole.Device.Webui.Device.Support = o.Role.Device.Webui.Device.Support
						}
						if o.Role.Device.Webui.Device.DynamicUpdates != nil {
							nestedRole.Device.Webui.Device.DynamicUpdates = o.Role.Device.Webui.Device.DynamicUpdates
						}
						if o.Role.Device.Webui.Device.LogSettings != nil {
							nestedRole.Device.Webui.Device.LogSettings = &RoleDeviceWebuiDeviceLogSettings{}
							if o.Role.Device.Webui.Device.LogSettings.Misc != nil {
								entry.Misc["RoleDeviceWebuiDeviceLogSettings"] = o.Role.Device.Webui.Device.LogSettings.Misc
							}
							if o.Role.Device.Webui.Device.LogSettings.CcAlarm != nil {
								nestedRole.Device.Webui.Device.LogSettings.CcAlarm = o.Role.Device.Webui.Device.LogSettings.CcAlarm
							}
							if o.Role.Device.Webui.Device.LogSettings.Config != nil {
								nestedRole.Device.Webui.Device.LogSettings.Config = o.Role.Device.Webui.Device.LogSettings.Config
							}
							if o.Role.Device.Webui.Device.LogSettings.Hipmatch != nil {
								nestedRole.Device.Webui.Device.LogSettings.Hipmatch = o.Role.Device.Webui.Device.LogSettings.Hipmatch
							}
							if o.Role.Device.Webui.Device.LogSettings.System != nil {
								nestedRole.Device.Webui.Device.LogSettings.System = o.Role.Device.Webui.Device.LogSettings.System
							}
							if o.Role.Device.Webui.Device.LogSettings.Correlation != nil {
								nestedRole.Device.Webui.Device.LogSettings.Correlation = o.Role.Device.Webui.Device.LogSettings.Correlation
							}
							if o.Role.Device.Webui.Device.LogSettings.Globalprotect != nil {
								nestedRole.Device.Webui.Device.LogSettings.Globalprotect = o.Role.Device.Webui.Device.LogSettings.Globalprotect
							}
							if o.Role.Device.Webui.Device.LogSettings.Iptag != nil {
								nestedRole.Device.Webui.Device.LogSettings.Iptag = o.Role.Device.Webui.Device.LogSettings.Iptag
							}
							if o.Role.Device.Webui.Device.LogSettings.ManageLog != nil {
								nestedRole.Device.Webui.Device.LogSettings.ManageLog = o.Role.Device.Webui.Device.LogSettings.ManageLog
							}
							if o.Role.Device.Webui.Device.LogSettings.UserId != nil {
								nestedRole.Device.Webui.Device.LogSettings.UserId = o.Role.Device.Webui.Device.LogSettings.UserId
							}
						}
						if o.Role.Device.Webui.Device.VmInfoSource != nil {
							nestedRole.Device.Webui.Device.VmInfoSource = o.Role.Device.Webui.Device.VmInfoSource
						}
						if o.Role.Device.Webui.Device.AdminRoles != nil {
							nestedRole.Device.Webui.Device.AdminRoles = o.Role.Device.Webui.Device.AdminRoles
						}
						if o.Role.Device.Webui.Device.Licenses != nil {
							nestedRole.Device.Webui.Device.Licenses = o.Role.Device.Webui.Device.Licenses
						}
						if o.Role.Device.Webui.Device.VirtualSystems != nil {
							nestedRole.Device.Webui.Device.VirtualSystems = o.Role.Device.Webui.Device.VirtualSystems
						}
						if o.Role.Device.Webui.Device.AuthenticationProfile != nil {
							nestedRole.Device.Webui.Device.AuthenticationProfile = o.Role.Device.Webui.Device.AuthenticationProfile
						}
					}
					if o.Role.Device.Webui.Monitor != nil {
						nestedRole.Device.Webui.Monitor = &RoleDeviceWebuiMonitor{}
						if o.Role.Device.Webui.Monitor.Misc != nil {
							entry.Misc["RoleDeviceWebuiMonitor"] = o.Role.Device.Webui.Monitor.Misc
						}
						if o.Role.Device.Webui.Monitor.CustomReports != nil {
							nestedRole.Device.Webui.Monitor.CustomReports = &RoleDeviceWebuiMonitorCustomReports{}
							if o.Role.Device.Webui.Monitor.CustomReports.Misc != nil {
								entry.Misc["RoleDeviceWebuiMonitorCustomReports"] = o.Role.Device.Webui.Monitor.CustomReports.Misc
							}
							if o.Role.Device.Webui.Monitor.CustomReports.DecryptionLog != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.DecryptionLog = o.Role.Device.Webui.Monitor.CustomReports.DecryptionLog
							}
							if o.Role.Device.Webui.Monitor.CustomReports.DecryptionSummary != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.DecryptionSummary = o.Role.Device.Webui.Monitor.CustomReports.DecryptionSummary
							}
							if o.Role.Device.Webui.Monitor.CustomReports.Iptag != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.Iptag = o.Role.Device.Webui.Monitor.CustomReports.Iptag
							}
							if o.Role.Device.Webui.Monitor.CustomReports.SctpLog != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.SctpLog = o.Role.Device.Webui.Monitor.CustomReports.SctpLog
							}
							if o.Role.Device.Webui.Monitor.CustomReports.Userid != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.Userid = o.Role.Device.Webui.Monitor.CustomReports.Userid
							}
							if o.Role.Device.Webui.Monitor.CustomReports.GtpLog != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.GtpLog = o.Role.Device.Webui.Monitor.CustomReports.GtpLog
							}
							if o.Role.Device.Webui.Monitor.CustomReports.ThreatSummary != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.ThreatSummary = o.Role.Device.Webui.Monitor.CustomReports.ThreatSummary
							}
							if o.Role.Device.Webui.Monitor.CustomReports.UrlSummary != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.UrlSummary = o.Role.Device.Webui.Monitor.CustomReports.UrlSummary
							}
							if o.Role.Device.Webui.Monitor.CustomReports.WildfireLog != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.WildfireLog = o.Role.Device.Webui.Monitor.CustomReports.WildfireLog
							}
							if o.Role.Device.Webui.Monitor.CustomReports.ApplicationStatistics != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.ApplicationStatistics = o.Role.Device.Webui.Monitor.CustomReports.ApplicationStatistics
							}
							if o.Role.Device.Webui.Monitor.CustomReports.Hipmatch != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.Hipmatch = o.Role.Device.Webui.Monitor.CustomReports.Hipmatch
							}
							if o.Role.Device.Webui.Monitor.CustomReports.ThreatLog != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.ThreatLog = o.Role.Device.Webui.Monitor.CustomReports.ThreatLog
							}
							if o.Role.Device.Webui.Monitor.CustomReports.TrafficLog != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.TrafficLog = o.Role.Device.Webui.Monitor.CustomReports.TrafficLog
							}
							if o.Role.Device.Webui.Monitor.CustomReports.TunnelLog != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.TunnelLog = o.Role.Device.Webui.Monitor.CustomReports.TunnelLog
							}
							if o.Role.Device.Webui.Monitor.CustomReports.TrafficSummary != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.TrafficSummary = o.Role.Device.Webui.Monitor.CustomReports.TrafficSummary
							}
							if o.Role.Device.Webui.Monitor.CustomReports.TunnelSummary != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.TunnelSummary = o.Role.Device.Webui.Monitor.CustomReports.TunnelSummary
							}
							if o.Role.Device.Webui.Monitor.CustomReports.UrlLog != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.UrlLog = o.Role.Device.Webui.Monitor.CustomReports.UrlLog
							}
							if o.Role.Device.Webui.Monitor.CustomReports.Auth != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.Auth = o.Role.Device.Webui.Monitor.CustomReports.Auth
							}
							if o.Role.Device.Webui.Monitor.CustomReports.DataFilteringLog != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.DataFilteringLog = o.Role.Device.Webui.Monitor.CustomReports.DataFilteringLog
							}
							if o.Role.Device.Webui.Monitor.CustomReports.Globalprotect != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.Globalprotect = o.Role.Device.Webui.Monitor.CustomReports.Globalprotect
							}
							if o.Role.Device.Webui.Monitor.CustomReports.GtpSummary != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.GtpSummary = o.Role.Device.Webui.Monitor.CustomReports.GtpSummary
							}
							if o.Role.Device.Webui.Monitor.CustomReports.SctpSummary != nil {
								nestedRole.Device.Webui.Monitor.CustomReports.SctpSummary = o.Role.Device.Webui.Monitor.CustomReports.SctpSummary
							}
						}
						if o.Role.Device.Webui.Monitor.ThreatReports != nil {
							nestedRole.Device.Webui.Monitor.ThreatReports = o.Role.Device.Webui.Monitor.ThreatReports
						}
						if o.Role.Device.Webui.Monitor.AppScope != nil {
							nestedRole.Device.Webui.Monitor.AppScope = o.Role.Device.Webui.Monitor.AppScope
						}
						if o.Role.Device.Webui.Monitor.ApplicationReports != nil {
							nestedRole.Device.Webui.Monitor.ApplicationReports = o.Role.Device.Webui.Monitor.ApplicationReports
						}
						if o.Role.Device.Webui.Monitor.GtpReports != nil {
							nestedRole.Device.Webui.Monitor.GtpReports = o.Role.Device.Webui.Monitor.GtpReports
						}
						if o.Role.Device.Webui.Monitor.ViewCustomReports != nil {
							nestedRole.Device.Webui.Monitor.ViewCustomReports = o.Role.Device.Webui.Monitor.ViewCustomReports
						}
						if o.Role.Device.Webui.Monitor.Botnet != nil {
							nestedRole.Device.Webui.Monitor.Botnet = o.Role.Device.Webui.Monitor.Botnet
						}
						if o.Role.Device.Webui.Monitor.ExternalLogs != nil {
							nestedRole.Device.Webui.Monitor.ExternalLogs = o.Role.Device.Webui.Monitor.ExternalLogs
						}
						if o.Role.Device.Webui.Monitor.Logs != nil {
							nestedRole.Device.Webui.Monitor.Logs = &RoleDeviceWebuiMonitorLogs{}
							if o.Role.Device.Webui.Monitor.Logs.Misc != nil {
								entry.Misc["RoleDeviceWebuiMonitorLogs"] = o.Role.Device.Webui.Monitor.Logs.Misc
							}
							if o.Role.Device.Webui.Monitor.Logs.Decryption != nil {
								nestedRole.Device.Webui.Monitor.Logs.Decryption = o.Role.Device.Webui.Monitor.Logs.Decryption
							}
							if o.Role.Device.Webui.Monitor.Logs.Gtp != nil {
								nestedRole.Device.Webui.Monitor.Logs.Gtp = o.Role.Device.Webui.Monitor.Logs.Gtp
							}
							if o.Role.Device.Webui.Monitor.Logs.Userid != nil {
								nestedRole.Device.Webui.Monitor.Logs.Userid = o.Role.Device.Webui.Monitor.Logs.Userid
							}
							if o.Role.Device.Webui.Monitor.Logs.DataFiltering != nil {
								nestedRole.Device.Webui.Monitor.Logs.DataFiltering = o.Role.Device.Webui.Monitor.Logs.DataFiltering
							}
							if o.Role.Device.Webui.Monitor.Logs.Iptag != nil {
								nestedRole.Device.Webui.Monitor.Logs.Iptag = o.Role.Device.Webui.Monitor.Logs.Iptag
							}
							if o.Role.Device.Webui.Monitor.Logs.Threat != nil {
								nestedRole.Device.Webui.Monitor.Logs.Threat = o.Role.Device.Webui.Monitor.Logs.Threat
							}
							if o.Role.Device.Webui.Monitor.Logs.Configuration != nil {
								nestedRole.Device.Webui.Monitor.Logs.Configuration = o.Role.Device.Webui.Monitor.Logs.Configuration
							}
							if o.Role.Device.Webui.Monitor.Logs.Sctp != nil {
								nestedRole.Device.Webui.Monitor.Logs.Sctp = o.Role.Device.Webui.Monitor.Logs.Sctp
							}
							if o.Role.Device.Webui.Monitor.Logs.Hipmatch != nil {
								nestedRole.Device.Webui.Monitor.Logs.Hipmatch = o.Role.Device.Webui.Monitor.Logs.Hipmatch
							}
							if o.Role.Device.Webui.Monitor.Logs.System != nil {
								nestedRole.Device.Webui.Monitor.Logs.System = o.Role.Device.Webui.Monitor.Logs.System
							}
							if o.Role.Device.Webui.Monitor.Logs.Traffic != nil {
								nestedRole.Device.Webui.Monitor.Logs.Traffic = o.Role.Device.Webui.Monitor.Logs.Traffic
							}
							if o.Role.Device.Webui.Monitor.Logs.Tunnel != nil {
								nestedRole.Device.Webui.Monitor.Logs.Tunnel = o.Role.Device.Webui.Monitor.Logs.Tunnel
							}
							if o.Role.Device.Webui.Monitor.Logs.Url != nil {
								nestedRole.Device.Webui.Monitor.Logs.Url = o.Role.Device.Webui.Monitor.Logs.Url
							}
							if o.Role.Device.Webui.Monitor.Logs.Alarm != nil {
								nestedRole.Device.Webui.Monitor.Logs.Alarm = o.Role.Device.Webui.Monitor.Logs.Alarm
							}
							if o.Role.Device.Webui.Monitor.Logs.Authentication != nil {
								nestedRole.Device.Webui.Monitor.Logs.Authentication = o.Role.Device.Webui.Monitor.Logs.Authentication
							}
							if o.Role.Device.Webui.Monitor.Logs.Globalprotect != nil {
								nestedRole.Device.Webui.Monitor.Logs.Globalprotect = o.Role.Device.Webui.Monitor.Logs.Globalprotect
							}
							if o.Role.Device.Webui.Monitor.Logs.Wildfire != nil {
								nestedRole.Device.Webui.Monitor.Logs.Wildfire = o.Role.Device.Webui.Monitor.Logs.Wildfire
							}
						}
						if o.Role.Device.Webui.Monitor.PacketCapture != nil {
							nestedRole.Device.Webui.Monitor.PacketCapture = o.Role.Device.Webui.Monitor.PacketCapture
						}
						if o.Role.Device.Webui.Monitor.TrafficReports != nil {
							nestedRole.Device.Webui.Monitor.TrafficReports = o.Role.Device.Webui.Monitor.TrafficReports
						}
						if o.Role.Device.Webui.Monitor.UrlFilteringReports != nil {
							nestedRole.Device.Webui.Monitor.UrlFilteringReports = o.Role.Device.Webui.Monitor.UrlFilteringReports
						}
						if o.Role.Device.Webui.Monitor.AutomatedCorrelationEngine != nil {
							nestedRole.Device.Webui.Monitor.AutomatedCorrelationEngine = &RoleDeviceWebuiMonitorAutomatedCorrelationEngine{}
							if o.Role.Device.Webui.Monitor.AutomatedCorrelationEngine.Misc != nil {
								entry.Misc["RoleDeviceWebuiMonitorAutomatedCorrelationEngine"] = o.Role.Device.Webui.Monitor.AutomatedCorrelationEngine.Misc
							}
							if o.Role.Device.Webui.Monitor.AutomatedCorrelationEngine.CorrelatedEvents != nil {
								nestedRole.Device.Webui.Monitor.AutomatedCorrelationEngine.CorrelatedEvents = o.Role.Device.Webui.Monitor.AutomatedCorrelationEngine.CorrelatedEvents
							}
							if o.Role.Device.Webui.Monitor.AutomatedCorrelationEngine.CorrelationObjects != nil {
								nestedRole.Device.Webui.Monitor.AutomatedCorrelationEngine.CorrelationObjects = o.Role.Device.Webui.Monitor.AutomatedCorrelationEngine.CorrelationObjects
							}
						}
						if o.Role.Device.Webui.Monitor.BlockIpList != nil {
							nestedRole.Device.Webui.Monitor.BlockIpList = o.Role.Device.Webui.Monitor.BlockIpList
						}
						if o.Role.Device.Webui.Monitor.PdfReports != nil {
							nestedRole.Device.Webui.Monitor.PdfReports = &RoleDeviceWebuiMonitorPdfReports{}
							if o.Role.Device.Webui.Monitor.PdfReports.Misc != nil {
								entry.Misc["RoleDeviceWebuiMonitorPdfReports"] = o.Role.Device.Webui.Monitor.PdfReports.Misc
							}
							if o.Role.Device.Webui.Monitor.PdfReports.SaasApplicationUsageReport != nil {
								nestedRole.Device.Webui.Monitor.PdfReports.SaasApplicationUsageReport = o.Role.Device.Webui.Monitor.PdfReports.SaasApplicationUsageReport
							}
							if o.Role.Device.Webui.Monitor.PdfReports.UserActivityReport != nil {
								nestedRole.Device.Webui.Monitor.PdfReports.UserActivityReport = o.Role.Device.Webui.Monitor.PdfReports.UserActivityReport
							}
							if o.Role.Device.Webui.Monitor.PdfReports.EmailScheduler != nil {
								nestedRole.Device.Webui.Monitor.PdfReports.EmailScheduler = o.Role.Device.Webui.Monitor.PdfReports.EmailScheduler
							}
							if o.Role.Device.Webui.Monitor.PdfReports.ManagePdfSummary != nil {
								nestedRole.Device.Webui.Monitor.PdfReports.ManagePdfSummary = o.Role.Device.Webui.Monitor.PdfReports.ManagePdfSummary
							}
							if o.Role.Device.Webui.Monitor.PdfReports.PdfSummaryReports != nil {
								nestedRole.Device.Webui.Monitor.PdfReports.PdfSummaryReports = o.Role.Device.Webui.Monitor.PdfReports.PdfSummaryReports
							}
							if o.Role.Device.Webui.Monitor.PdfReports.ReportGroups != nil {
								nestedRole.Device.Webui.Monitor.PdfReports.ReportGroups = o.Role.Device.Webui.Monitor.PdfReports.ReportGroups
							}
						}
						if o.Role.Device.Webui.Monitor.SctpReports != nil {
							nestedRole.Device.Webui.Monitor.SctpReports = o.Role.Device.Webui.Monitor.SctpReports
						}
						if o.Role.Device.Webui.Monitor.SessionBrowser != nil {
							nestedRole.Device.Webui.Monitor.SessionBrowser = o.Role.Device.Webui.Monitor.SessionBrowser
						}
					}
					if o.Role.Device.Webui.Operations != nil {
						nestedRole.Device.Webui.Operations = &RoleDeviceWebuiOperations{}
						if o.Role.Device.Webui.Operations.Misc != nil {
							entry.Misc["RoleDeviceWebuiOperations"] = o.Role.Device.Webui.Operations.Misc
						}
						if o.Role.Device.Webui.Operations.GenerateStatsDumpFile != nil {
							nestedRole.Device.Webui.Operations.GenerateStatsDumpFile = o.Role.Device.Webui.Operations.GenerateStatsDumpFile
						}
						if o.Role.Device.Webui.Operations.GenerateTechSupportFile != nil {
							nestedRole.Device.Webui.Operations.GenerateTechSupportFile = o.Role.Device.Webui.Operations.GenerateTechSupportFile
						}
						if o.Role.Device.Webui.Operations.Reboot != nil {
							nestedRole.Device.Webui.Operations.Reboot = o.Role.Device.Webui.Operations.Reboot
						}
						if o.Role.Device.Webui.Operations.DownloadCoreFiles != nil {
							nestedRole.Device.Webui.Operations.DownloadCoreFiles = o.Role.Device.Webui.Operations.DownloadCoreFiles
						}
						if o.Role.Device.Webui.Operations.DownloadPcapFiles != nil {
							nestedRole.Device.Webui.Operations.DownloadPcapFiles = o.Role.Device.Webui.Operations.DownloadPcapFiles
						}
					}
					if o.Role.Device.Webui.Save != nil {
						nestedRole.Device.Webui.Save = &RoleDeviceWebuiSave{}
						if o.Role.Device.Webui.Save.Misc != nil {
							entry.Misc["RoleDeviceWebuiSave"] = o.Role.Device.Webui.Save.Misc
						}
						if o.Role.Device.Webui.Save.PartialSave != nil {
							nestedRole.Device.Webui.Save.PartialSave = o.Role.Device.Webui.Save.PartialSave
						}
						if o.Role.Device.Webui.Save.SaveForOtherAdmins != nil {
							nestedRole.Device.Webui.Save.SaveForOtherAdmins = o.Role.Device.Webui.Save.SaveForOtherAdmins
						}
						if o.Role.Device.Webui.Save.ObjectLevelChanges != nil {
							nestedRole.Device.Webui.Save.ObjectLevelChanges = o.Role.Device.Webui.Save.ObjectLevelChanges
						}
					}
					if o.Role.Device.Webui.Validate != nil {
						nestedRole.Device.Webui.Validate = o.Role.Device.Webui.Validate
					}
					if o.Role.Device.Webui.Global != nil {
						nestedRole.Device.Webui.Global = &RoleDeviceWebuiGlobal{}
						if o.Role.Device.Webui.Global.Misc != nil {
							entry.Misc["RoleDeviceWebuiGlobal"] = o.Role.Device.Webui.Global.Misc
						}
						if o.Role.Device.Webui.Global.SystemAlarms != nil {
							nestedRole.Device.Webui.Global.SystemAlarms = o.Role.Device.Webui.Global.SystemAlarms
						}
					}
					if o.Role.Device.Webui.Objects != nil {
						nestedRole.Device.Webui.Objects = &RoleDeviceWebuiObjects{}
						if o.Role.Device.Webui.Objects.Misc != nil {
							entry.Misc["RoleDeviceWebuiObjects"] = o.Role.Device.Webui.Objects.Misc
						}
						if o.Role.Device.Webui.Objects.ApplicationFilters != nil {
							nestedRole.Device.Webui.Objects.ApplicationFilters = o.Role.Device.Webui.Objects.ApplicationFilters
						}
						if o.Role.Device.Webui.Objects.Decryption != nil {
							nestedRole.Device.Webui.Objects.Decryption = &RoleDeviceWebuiObjectsDecryption{}
							if o.Role.Device.Webui.Objects.Decryption.Misc != nil {
								entry.Misc["RoleDeviceWebuiObjectsDecryption"] = o.Role.Device.Webui.Objects.Decryption.Misc
							}
							if o.Role.Device.Webui.Objects.Decryption.DecryptionProfile != nil {
								nestedRole.Device.Webui.Objects.Decryption.DecryptionProfile = o.Role.Device.Webui.Objects.Decryption.DecryptionProfile
							}
						}
						if o.Role.Device.Webui.Objects.Devices != nil {
							nestedRole.Device.Webui.Objects.Devices = o.Role.Device.Webui.Objects.Devices
						}
						if o.Role.Device.Webui.Objects.GlobalProtect != nil {
							nestedRole.Device.Webui.Objects.GlobalProtect = &RoleDeviceWebuiObjectsGlobalProtect{}
							if o.Role.Device.Webui.Objects.GlobalProtect.Misc != nil {
								entry.Misc["RoleDeviceWebuiObjectsGlobalProtect"] = o.Role.Device.Webui.Objects.GlobalProtect.Misc
							}
							if o.Role.Device.Webui.Objects.GlobalProtect.HipObjects != nil {
								nestedRole.Device.Webui.Objects.GlobalProtect.HipObjects = o.Role.Device.Webui.Objects.GlobalProtect.HipObjects
							}
							if o.Role.Device.Webui.Objects.GlobalProtect.HipProfiles != nil {
								nestedRole.Device.Webui.Objects.GlobalProtect.HipProfiles = o.Role.Device.Webui.Objects.GlobalProtect.HipProfiles
							}
						}
						if o.Role.Device.Webui.Objects.PacketBrokerProfile != nil {
							nestedRole.Device.Webui.Objects.PacketBrokerProfile = o.Role.Device.Webui.Objects.PacketBrokerProfile
						}
						if o.Role.Device.Webui.Objects.Regions != nil {
							nestedRole.Device.Webui.Objects.Regions = o.Role.Device.Webui.Objects.Regions
						}
						if o.Role.Device.Webui.Objects.Schedules != nil {
							nestedRole.Device.Webui.Objects.Schedules = o.Role.Device.Webui.Objects.Schedules
						}
						if o.Role.Device.Webui.Objects.ServiceGroups != nil {
							nestedRole.Device.Webui.Objects.ServiceGroups = o.Role.Device.Webui.Objects.ServiceGroups
						}
						if o.Role.Device.Webui.Objects.Tags != nil {
							nestedRole.Device.Webui.Objects.Tags = o.Role.Device.Webui.Objects.Tags
						}
						if o.Role.Device.Webui.Objects.Addresses != nil {
							nestedRole.Device.Webui.Objects.Addresses = o.Role.Device.Webui.Objects.Addresses
						}
						if o.Role.Device.Webui.Objects.Authentication != nil {
							nestedRole.Device.Webui.Objects.Authentication = o.Role.Device.Webui.Objects.Authentication
						}
						if o.Role.Device.Webui.Objects.CustomObjects != nil {
							nestedRole.Device.Webui.Objects.CustomObjects = &RoleDeviceWebuiObjectsCustomObjects{}
							if o.Role.Device.Webui.Objects.CustomObjects.Misc != nil {
								entry.Misc["RoleDeviceWebuiObjectsCustomObjects"] = o.Role.Device.Webui.Objects.CustomObjects.Misc
							}
							if o.Role.Device.Webui.Objects.CustomObjects.DataPatterns != nil {
								nestedRole.Device.Webui.Objects.CustomObjects.DataPatterns = o.Role.Device.Webui.Objects.CustomObjects.DataPatterns
							}
							if o.Role.Device.Webui.Objects.CustomObjects.Spyware != nil {
								nestedRole.Device.Webui.Objects.CustomObjects.Spyware = o.Role.Device.Webui.Objects.CustomObjects.Spyware
							}
							if o.Role.Device.Webui.Objects.CustomObjects.UrlCategory != nil {
								nestedRole.Device.Webui.Objects.CustomObjects.UrlCategory = o.Role.Device.Webui.Objects.CustomObjects.UrlCategory
							}
							if o.Role.Device.Webui.Objects.CustomObjects.Vulnerability != nil {
								nestedRole.Device.Webui.Objects.CustomObjects.Vulnerability = o.Role.Device.Webui.Objects.CustomObjects.Vulnerability
							}
						}
						if o.Role.Device.Webui.Objects.SecurityProfileGroups != nil {
							nestedRole.Device.Webui.Objects.SecurityProfileGroups = o.Role.Device.Webui.Objects.SecurityProfileGroups
						}
						if o.Role.Device.Webui.Objects.AddressGroups != nil {
							nestedRole.Device.Webui.Objects.AddressGroups = o.Role.Device.Webui.Objects.AddressGroups
						}
						if o.Role.Device.Webui.Objects.ApplicationGroups != nil {
							nestedRole.Device.Webui.Objects.ApplicationGroups = o.Role.Device.Webui.Objects.ApplicationGroups
						}
						if o.Role.Device.Webui.Objects.DynamicBlockLists != nil {
							nestedRole.Device.Webui.Objects.DynamicBlockLists = o.Role.Device.Webui.Objects.DynamicBlockLists
						}
						if o.Role.Device.Webui.Objects.DynamicUserGroups != nil {
							nestedRole.Device.Webui.Objects.DynamicUserGroups = o.Role.Device.Webui.Objects.DynamicUserGroups
						}
						if o.Role.Device.Webui.Objects.Services != nil {
							nestedRole.Device.Webui.Objects.Services = o.Role.Device.Webui.Objects.Services
						}
						if o.Role.Device.Webui.Objects.Applications != nil {
							nestedRole.Device.Webui.Objects.Applications = o.Role.Device.Webui.Objects.Applications
						}
						if o.Role.Device.Webui.Objects.LogForwarding != nil {
							nestedRole.Device.Webui.Objects.LogForwarding = o.Role.Device.Webui.Objects.LogForwarding
						}
						if o.Role.Device.Webui.Objects.Sdwan != nil {
							nestedRole.Device.Webui.Objects.Sdwan = &RoleDeviceWebuiObjectsSdwan{}
							if o.Role.Device.Webui.Objects.Sdwan.Misc != nil {
								entry.Misc["RoleDeviceWebuiObjectsSdwan"] = o.Role.Device.Webui.Objects.Sdwan.Misc
							}
							if o.Role.Device.Webui.Objects.Sdwan.SdwanDistProfile != nil {
								nestedRole.Device.Webui.Objects.Sdwan.SdwanDistProfile = o.Role.Device.Webui.Objects.Sdwan.SdwanDistProfile
							}
							if o.Role.Device.Webui.Objects.Sdwan.SdwanErrorCorrectionProfile != nil {
								nestedRole.Device.Webui.Objects.Sdwan.SdwanErrorCorrectionProfile = o.Role.Device.Webui.Objects.Sdwan.SdwanErrorCorrectionProfile
							}
							if o.Role.Device.Webui.Objects.Sdwan.SdwanProfile != nil {
								nestedRole.Device.Webui.Objects.Sdwan.SdwanProfile = o.Role.Device.Webui.Objects.Sdwan.SdwanProfile
							}
							if o.Role.Device.Webui.Objects.Sdwan.SdwanSaasQualityProfile != nil {
								nestedRole.Device.Webui.Objects.Sdwan.SdwanSaasQualityProfile = o.Role.Device.Webui.Objects.Sdwan.SdwanSaasQualityProfile
							}
						}
						if o.Role.Device.Webui.Objects.SecurityProfiles != nil {
							nestedRole.Device.Webui.Objects.SecurityProfiles = &RoleDeviceWebuiObjectsSecurityProfiles{}
							if o.Role.Device.Webui.Objects.SecurityProfiles.Misc != nil {
								entry.Misc["RoleDeviceWebuiObjectsSecurityProfiles"] = o.Role.Device.Webui.Objects.SecurityProfiles.Misc
							}
							if o.Role.Device.Webui.Objects.SecurityProfiles.Antivirus != nil {
								nestedRole.Device.Webui.Objects.SecurityProfiles.Antivirus = o.Role.Device.Webui.Objects.SecurityProfiles.Antivirus
							}
							if o.Role.Device.Webui.Objects.SecurityProfiles.DosProtection != nil {
								nestedRole.Device.Webui.Objects.SecurityProfiles.DosProtection = o.Role.Device.Webui.Objects.SecurityProfiles.DosProtection
							}
							if o.Role.Device.Webui.Objects.SecurityProfiles.FileBlocking != nil {
								nestedRole.Device.Webui.Objects.SecurityProfiles.FileBlocking = o.Role.Device.Webui.Objects.SecurityProfiles.FileBlocking
							}
							if o.Role.Device.Webui.Objects.SecurityProfiles.GtpProtection != nil {
								nestedRole.Device.Webui.Objects.SecurityProfiles.GtpProtection = o.Role.Device.Webui.Objects.SecurityProfiles.GtpProtection
							}
							if o.Role.Device.Webui.Objects.SecurityProfiles.WildfireAnalysis != nil {
								nestedRole.Device.Webui.Objects.SecurityProfiles.WildfireAnalysis = o.Role.Device.Webui.Objects.SecurityProfiles.WildfireAnalysis
							}
							if o.Role.Device.Webui.Objects.SecurityProfiles.AntiSpyware != nil {
								nestedRole.Device.Webui.Objects.SecurityProfiles.AntiSpyware = o.Role.Device.Webui.Objects.SecurityProfiles.AntiSpyware
							}
							if o.Role.Device.Webui.Objects.SecurityProfiles.DataFiltering != nil {
								nestedRole.Device.Webui.Objects.SecurityProfiles.DataFiltering = o.Role.Device.Webui.Objects.SecurityProfiles.DataFiltering
							}
							if o.Role.Device.Webui.Objects.SecurityProfiles.SctpProtection != nil {
								nestedRole.Device.Webui.Objects.SecurityProfiles.SctpProtection = o.Role.Device.Webui.Objects.SecurityProfiles.SctpProtection
							}
							if o.Role.Device.Webui.Objects.SecurityProfiles.UrlFiltering != nil {
								nestedRole.Device.Webui.Objects.SecurityProfiles.UrlFiltering = o.Role.Device.Webui.Objects.SecurityProfiles.UrlFiltering
							}
							if o.Role.Device.Webui.Objects.SecurityProfiles.VulnerabilityProtection != nil {
								nestedRole.Device.Webui.Objects.SecurityProfiles.VulnerabilityProtection = o.Role.Device.Webui.Objects.SecurityProfiles.VulnerabilityProtection
							}
						}
					}
					if o.Role.Device.Webui.Policies != nil {
						nestedRole.Device.Webui.Policies = &RoleDeviceWebuiPolicies{}
						if o.Role.Device.Webui.Policies.Misc != nil {
							entry.Misc["RoleDeviceWebuiPolicies"] = o.Role.Device.Webui.Policies.Misc
						}
						if o.Role.Device.Webui.Policies.NetworkPacketBrokerRulebase != nil {
							nestedRole.Device.Webui.Policies.NetworkPacketBrokerRulebase = o.Role.Device.Webui.Policies.NetworkPacketBrokerRulebase
						}
						if o.Role.Device.Webui.Policies.SecurityRulebase != nil {
							nestedRole.Device.Webui.Policies.SecurityRulebase = o.Role.Device.Webui.Policies.SecurityRulebase
						}
						if o.Role.Device.Webui.Policies.SslDecryptionRulebase != nil {
							nestedRole.Device.Webui.Policies.SslDecryptionRulebase = o.Role.Device.Webui.Policies.SslDecryptionRulebase
						}
						if o.Role.Device.Webui.Policies.AuthenticationRulebase != nil {
							nestedRole.Device.Webui.Policies.AuthenticationRulebase = o.Role.Device.Webui.Policies.AuthenticationRulebase
						}
						if o.Role.Device.Webui.Policies.NatRulebase != nil {
							nestedRole.Device.Webui.Policies.NatRulebase = o.Role.Device.Webui.Policies.NatRulebase
						}
						if o.Role.Device.Webui.Policies.PbfRulebase != nil {
							nestedRole.Device.Webui.Policies.PbfRulebase = o.Role.Device.Webui.Policies.PbfRulebase
						}
						if o.Role.Device.Webui.Policies.QosRulebase != nil {
							nestedRole.Device.Webui.Policies.QosRulebase = o.Role.Device.Webui.Policies.QosRulebase
						}
						if o.Role.Device.Webui.Policies.RuleHitCountReset != nil {
							nestedRole.Device.Webui.Policies.RuleHitCountReset = o.Role.Device.Webui.Policies.RuleHitCountReset
						}
						if o.Role.Device.Webui.Policies.SdwanRulebase != nil {
							nestedRole.Device.Webui.Policies.SdwanRulebase = o.Role.Device.Webui.Policies.SdwanRulebase
						}
						if o.Role.Device.Webui.Policies.TunnelInspectRulebase != nil {
							nestedRole.Device.Webui.Policies.TunnelInspectRulebase = o.Role.Device.Webui.Policies.TunnelInspectRulebase
						}
						if o.Role.Device.Webui.Policies.ApplicationOverrideRulebase != nil {
							nestedRole.Device.Webui.Policies.ApplicationOverrideRulebase = o.Role.Device.Webui.Policies.ApplicationOverrideRulebase
						}
						if o.Role.Device.Webui.Policies.DosRulebase != nil {
							nestedRole.Device.Webui.Policies.DosRulebase = o.Role.Device.Webui.Policies.DosRulebase
						}
					}
					if o.Role.Device.Webui.Acc != nil {
						nestedRole.Device.Webui.Acc = o.Role.Device.Webui.Acc
					}
					if o.Role.Device.Webui.Dashboard != nil {
						nestedRole.Device.Webui.Dashboard = o.Role.Device.Webui.Dashboard
					}
					if o.Role.Device.Webui.Network != nil {
						nestedRole.Device.Webui.Network = &RoleDeviceWebuiNetwork{}
						if o.Role.Device.Webui.Network.Misc != nil {
							entry.Misc["RoleDeviceWebuiNetwork"] = o.Role.Device.Webui.Network.Misc
						}
						if o.Role.Device.Webui.Network.Lldp != nil {
							nestedRole.Device.Webui.Network.Lldp = o.Role.Device.Webui.Network.Lldp
						}
						if o.Role.Device.Webui.Network.Qos != nil {
							nestedRole.Device.Webui.Network.Qos = o.Role.Device.Webui.Network.Qos
						}
						if o.Role.Device.Webui.Network.VirtualRouters != nil {
							nestedRole.Device.Webui.Network.VirtualRouters = o.Role.Device.Webui.Network.VirtualRouters
						}
						if o.Role.Device.Webui.Network.Dhcp != nil {
							nestedRole.Device.Webui.Network.Dhcp = o.Role.Device.Webui.Network.Dhcp
						}
						if o.Role.Device.Webui.Network.GlobalProtect != nil {
							nestedRole.Device.Webui.Network.GlobalProtect = &RoleDeviceWebuiNetworkGlobalProtect{}
							if o.Role.Device.Webui.Network.GlobalProtect.Misc != nil {
								entry.Misc["RoleDeviceWebuiNetworkGlobalProtect"] = o.Role.Device.Webui.Network.GlobalProtect.Misc
							}
							if o.Role.Device.Webui.Network.GlobalProtect.ClientlessApps != nil {
								nestedRole.Device.Webui.Network.GlobalProtect.ClientlessApps = o.Role.Device.Webui.Network.GlobalProtect.ClientlessApps
							}
							if o.Role.Device.Webui.Network.GlobalProtect.Gateways != nil {
								nestedRole.Device.Webui.Network.GlobalProtect.Gateways = o.Role.Device.Webui.Network.GlobalProtect.Gateways
							}
							if o.Role.Device.Webui.Network.GlobalProtect.Mdm != nil {
								nestedRole.Device.Webui.Network.GlobalProtect.Mdm = o.Role.Device.Webui.Network.GlobalProtect.Mdm
							}
							if o.Role.Device.Webui.Network.GlobalProtect.Portals != nil {
								nestedRole.Device.Webui.Network.GlobalProtect.Portals = o.Role.Device.Webui.Network.GlobalProtect.Portals
							}
							if o.Role.Device.Webui.Network.GlobalProtect.ClientlessAppGroups != nil {
								nestedRole.Device.Webui.Network.GlobalProtect.ClientlessAppGroups = o.Role.Device.Webui.Network.GlobalProtect.ClientlessAppGroups
							}
						}
						if o.Role.Device.Webui.Network.GreTunnels != nil {
							nestedRole.Device.Webui.Network.GreTunnels = o.Role.Device.Webui.Network.GreTunnels
						}
						if o.Role.Device.Webui.Network.DnsProxy != nil {
							nestedRole.Device.Webui.Network.DnsProxy = o.Role.Device.Webui.Network.DnsProxy
						}
						if o.Role.Device.Webui.Network.Interfaces != nil {
							nestedRole.Device.Webui.Network.Interfaces = o.Role.Device.Webui.Network.Interfaces
						}
						if o.Role.Device.Webui.Network.Zones != nil {
							nestedRole.Device.Webui.Network.Zones = o.Role.Device.Webui.Network.Zones
						}
						if o.Role.Device.Webui.Network.VirtualWires != nil {
							nestedRole.Device.Webui.Network.VirtualWires = o.Role.Device.Webui.Network.VirtualWires
						}
						if o.Role.Device.Webui.Network.IpsecTunnels != nil {
							nestedRole.Device.Webui.Network.IpsecTunnels = o.Role.Device.Webui.Network.IpsecTunnels
						}
						if o.Role.Device.Webui.Network.NetworkProfiles != nil {
							nestedRole.Device.Webui.Network.NetworkProfiles = &RoleDeviceWebuiNetworkNetworkProfiles{}
							if o.Role.Device.Webui.Network.NetworkProfiles.Misc != nil {
								entry.Misc["RoleDeviceWebuiNetworkNetworkProfiles"] = o.Role.Device.Webui.Network.NetworkProfiles.Misc
							}
							if o.Role.Device.Webui.Network.NetworkProfiles.BfdProfile != nil {
								nestedRole.Device.Webui.Network.NetworkProfiles.BfdProfile = o.Role.Device.Webui.Network.NetworkProfiles.BfdProfile
							}
							if o.Role.Device.Webui.Network.NetworkProfiles.IkeCrypto != nil {
								nestedRole.Device.Webui.Network.NetworkProfiles.IkeCrypto = o.Role.Device.Webui.Network.NetworkProfiles.IkeCrypto
							}
							if o.Role.Device.Webui.Network.NetworkProfiles.QosProfile != nil {
								nestedRole.Device.Webui.Network.NetworkProfiles.QosProfile = o.Role.Device.Webui.Network.NetworkProfiles.QosProfile
							}
							if o.Role.Device.Webui.Network.NetworkProfiles.TunnelMonitor != nil {
								nestedRole.Device.Webui.Network.NetworkProfiles.TunnelMonitor = o.Role.Device.Webui.Network.NetworkProfiles.TunnelMonitor
							}
							if o.Role.Device.Webui.Network.NetworkProfiles.ZoneProtection != nil {
								nestedRole.Device.Webui.Network.NetworkProfiles.ZoneProtection = o.Role.Device.Webui.Network.NetworkProfiles.ZoneProtection
							}
							if o.Role.Device.Webui.Network.NetworkProfiles.GpAppIpsecCrypto != nil {
								nestedRole.Device.Webui.Network.NetworkProfiles.GpAppIpsecCrypto = o.Role.Device.Webui.Network.NetworkProfiles.GpAppIpsecCrypto
							}
							if o.Role.Device.Webui.Network.NetworkProfiles.IkeGateways != nil {
								nestedRole.Device.Webui.Network.NetworkProfiles.IkeGateways = o.Role.Device.Webui.Network.NetworkProfiles.IkeGateways
							}
							if o.Role.Device.Webui.Network.NetworkProfiles.InterfaceMgmt != nil {
								nestedRole.Device.Webui.Network.NetworkProfiles.InterfaceMgmt = o.Role.Device.Webui.Network.NetworkProfiles.InterfaceMgmt
							}
							if o.Role.Device.Webui.Network.NetworkProfiles.IpsecCrypto != nil {
								nestedRole.Device.Webui.Network.NetworkProfiles.IpsecCrypto = o.Role.Device.Webui.Network.NetworkProfiles.IpsecCrypto
							}
							if o.Role.Device.Webui.Network.NetworkProfiles.LldpProfile != nil {
								nestedRole.Device.Webui.Network.NetworkProfiles.LldpProfile = o.Role.Device.Webui.Network.NetworkProfiles.LldpProfile
							}
						}
						if o.Role.Device.Webui.Network.Routing != nil {
							nestedRole.Device.Webui.Network.Routing = &RoleDeviceWebuiNetworkRouting{}
							if o.Role.Device.Webui.Network.Routing.Misc != nil {
								entry.Misc["RoleDeviceWebuiNetworkRouting"] = o.Role.Device.Webui.Network.Routing.Misc
							}
							if o.Role.Device.Webui.Network.Routing.LogicalRouters != nil {
								nestedRole.Device.Webui.Network.Routing.LogicalRouters = o.Role.Device.Webui.Network.Routing.LogicalRouters
							}
							if o.Role.Device.Webui.Network.Routing.RoutingProfiles != nil {
								nestedRole.Device.Webui.Network.Routing.RoutingProfiles = &RoleDeviceWebuiNetworkRoutingRoutingProfiles{}
								if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Misc != nil {
									entry.Misc["RoleDeviceWebuiNetworkRoutingRoutingProfiles"] = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Misc
								}
								if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Bgp != nil {
									nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Bgp = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Bgp
								}
								if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Filters != nil {
									nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Filters = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Filters
								}
								if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Multicast != nil {
									nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Multicast = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Multicast
								}
								if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Ospf != nil {
									nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Ospf = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Ospf
								}
								if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Ospfv3 != nil {
									nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Ospfv3 = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Ospfv3
								}
								if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Ripv2 != nil {
									nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Ripv2 = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Ripv2
								}
								if o.Role.Device.Webui.Network.Routing.RoutingProfiles.Bfd != nil {
									nestedRole.Device.Webui.Network.Routing.RoutingProfiles.Bfd = o.Role.Device.Webui.Network.Routing.RoutingProfiles.Bfd
								}
							}
						}
						if o.Role.Device.Webui.Network.SdwanInterfaceProfile != nil {
							nestedRole.Device.Webui.Network.SdwanInterfaceProfile = o.Role.Device.Webui.Network.SdwanInterfaceProfile
						}
						if o.Role.Device.Webui.Network.Vlans != nil {
							nestedRole.Device.Webui.Network.Vlans = o.Role.Device.Webui.Network.Vlans
						}
						if o.Role.Device.Webui.Network.SecureWebGateway != nil {
							nestedRole.Device.Webui.Network.SecureWebGateway = o.Role.Device.Webui.Network.SecureWebGateway
						}
					}
					if o.Role.Device.Webui.Privacy != nil {
						nestedRole.Device.Webui.Privacy = &RoleDeviceWebuiPrivacy{}
						if o.Role.Device.Webui.Privacy.Misc != nil {
							entry.Misc["RoleDeviceWebuiPrivacy"] = o.Role.Device.Webui.Privacy.Misc
						}
						if o.Role.Device.Webui.Privacy.ShowFullIpAddresses != nil {
							nestedRole.Device.Webui.Privacy.ShowFullIpAddresses = o.Role.Device.Webui.Privacy.ShowFullIpAddresses
						}
						if o.Role.Device.Webui.Privacy.ShowUserNamesInLogsAndReports != nil {
							nestedRole.Device.Webui.Privacy.ShowUserNamesInLogsAndReports = o.Role.Device.Webui.Privacy.ShowUserNamesInLogsAndReports
						}
						if o.Role.Device.Webui.Privacy.ViewPcapFiles != nil {
							nestedRole.Device.Webui.Privacy.ViewPcapFiles = o.Role.Device.Webui.Privacy.ViewPcapFiles
						}
					}
					if o.Role.Device.Webui.Tasks != nil {
						nestedRole.Device.Webui.Tasks = o.Role.Device.Webui.Tasks
					}
				}
				if o.Role.Device.Xmlapi != nil {
					nestedRole.Device.Xmlapi = &RoleDeviceXmlapi{}
					if o.Role.Device.Xmlapi.Misc != nil {
						entry.Misc["RoleDeviceXmlapi"] = o.Role.Device.Xmlapi.Misc
					}
					if o.Role.Device.Xmlapi.Commit != nil {
						nestedRole.Device.Xmlapi.Commit = o.Role.Device.Xmlapi.Commit
					}
					if o.Role.Device.Xmlapi.Import != nil {
						nestedRole.Device.Xmlapi.Import = o.Role.Device.Xmlapi.Import
					}
					if o.Role.Device.Xmlapi.Op != nil {
						nestedRole.Device.Xmlapi.Op = o.Role.Device.Xmlapi.Op
					}
					if o.Role.Device.Xmlapi.Report != nil {
						nestedRole.Device.Xmlapi.Report = o.Role.Device.Xmlapi.Report
					}
					if o.Role.Device.Xmlapi.UserId != nil {
						nestedRole.Device.Xmlapi.UserId = o.Role.Device.Xmlapi.UserId
					}
					if o.Role.Device.Xmlapi.Config != nil {
						nestedRole.Device.Xmlapi.Config = o.Role.Device.Xmlapi.Config
					}
					if o.Role.Device.Xmlapi.Export != nil {
						nestedRole.Device.Xmlapi.Export = o.Role.Device.Xmlapi.Export
					}
					if o.Role.Device.Xmlapi.Iot != nil {
						nestedRole.Device.Xmlapi.Iot = o.Role.Device.Xmlapi.Iot
					}
					if o.Role.Device.Xmlapi.Log != nil {
						nestedRole.Device.Xmlapi.Log = o.Role.Device.Xmlapi.Log
					}
				}
			}
			if o.Role.Vsys != nil {
				nestedRole.Vsys = &RoleVsys{}
				if o.Role.Vsys.Misc != nil {
					entry.Misc["RoleVsys"] = o.Role.Vsys.Misc
				}
				if o.Role.Vsys.Cli != nil {
					nestedRole.Vsys.Cli = o.Role.Vsys.Cli
				}
				if o.Role.Vsys.Restapi != nil {
					nestedRole.Vsys.Restapi = &RoleVsysRestapi{}
					if o.Role.Vsys.Restapi.Misc != nil {
						entry.Misc["RoleVsysRestapi"] = o.Role.Vsys.Restapi.Misc
					}
					if o.Role.Vsys.Restapi.Device != nil {
						nestedRole.Vsys.Restapi.Device = &RoleVsysRestapiDevice{}
						if o.Role.Vsys.Restapi.Device.Misc != nil {
							entry.Misc["RoleVsysRestapiDevice"] = o.Role.Vsys.Restapi.Device.Misc
						}
						if o.Role.Vsys.Restapi.Device.EmailServerProfiles != nil {
							nestedRole.Vsys.Restapi.Device.EmailServerProfiles = o.Role.Vsys.Restapi.Device.EmailServerProfiles
						}
						if o.Role.Vsys.Restapi.Device.HttpServerProfiles != nil {
							nestedRole.Vsys.Restapi.Device.HttpServerProfiles = o.Role.Vsys.Restapi.Device.HttpServerProfiles
						}
						if o.Role.Vsys.Restapi.Device.LdapServerProfiles != nil {
							nestedRole.Vsys.Restapi.Device.LdapServerProfiles = o.Role.Vsys.Restapi.Device.LdapServerProfiles
						}
						if o.Role.Vsys.Restapi.Device.LogInterfaceSetting != nil {
							nestedRole.Vsys.Restapi.Device.LogInterfaceSetting = o.Role.Vsys.Restapi.Device.LogInterfaceSetting
						}
						if o.Role.Vsys.Restapi.Device.SnmpTrapServerProfiles != nil {
							nestedRole.Vsys.Restapi.Device.SnmpTrapServerProfiles = o.Role.Vsys.Restapi.Device.SnmpTrapServerProfiles
						}
						if o.Role.Vsys.Restapi.Device.SyslogServerProfiles != nil {
							nestedRole.Vsys.Restapi.Device.SyslogServerProfiles = o.Role.Vsys.Restapi.Device.SyslogServerProfiles
						}
						if o.Role.Vsys.Restapi.Device.VirtualSystems != nil {
							nestedRole.Vsys.Restapi.Device.VirtualSystems = o.Role.Vsys.Restapi.Device.VirtualSystems
						}
					}
					if o.Role.Vsys.Restapi.Network != nil {
						nestedRole.Vsys.Restapi.Network = &RoleVsysRestapiNetwork{}
						if o.Role.Vsys.Restapi.Network.Misc != nil {
							entry.Misc["RoleVsysRestapiNetwork"] = o.Role.Vsys.Restapi.Network.Misc
						}
						if o.Role.Vsys.Restapi.Network.GlobalprotectMdmServers != nil {
							nestedRole.Vsys.Restapi.Network.GlobalprotectMdmServers = o.Role.Vsys.Restapi.Network.GlobalprotectMdmServers
						}
						if o.Role.Vsys.Restapi.Network.GlobalprotectPortals != nil {
							nestedRole.Vsys.Restapi.Network.GlobalprotectPortals = o.Role.Vsys.Restapi.Network.GlobalprotectPortals
						}
						if o.Role.Vsys.Restapi.Network.Zones != nil {
							nestedRole.Vsys.Restapi.Network.Zones = o.Role.Vsys.Restapi.Network.Zones
						}
						if o.Role.Vsys.Restapi.Network.SdwanInterfaceProfiles != nil {
							nestedRole.Vsys.Restapi.Network.SdwanInterfaceProfiles = o.Role.Vsys.Restapi.Network.SdwanInterfaceProfiles
						}
						if o.Role.Vsys.Restapi.Network.GlobalprotectClientlessAppGroups != nil {
							nestedRole.Vsys.Restapi.Network.GlobalprotectClientlessAppGroups = o.Role.Vsys.Restapi.Network.GlobalprotectClientlessAppGroups
						}
						if o.Role.Vsys.Restapi.Network.GlobalprotectClientlessApps != nil {
							nestedRole.Vsys.Restapi.Network.GlobalprotectClientlessApps = o.Role.Vsys.Restapi.Network.GlobalprotectClientlessApps
						}
						if o.Role.Vsys.Restapi.Network.GlobalprotectGateways != nil {
							nestedRole.Vsys.Restapi.Network.GlobalprotectGateways = o.Role.Vsys.Restapi.Network.GlobalprotectGateways
						}
					}
					if o.Role.Vsys.Restapi.Objects != nil {
						nestedRole.Vsys.Restapi.Objects = &RoleVsysRestapiObjects{}
						if o.Role.Vsys.Restapi.Objects.Misc != nil {
							entry.Misc["RoleVsysRestapiObjects"] = o.Role.Vsys.Restapi.Objects.Misc
						}
						if o.Role.Vsys.Restapi.Objects.ExternalDynamicLists != nil {
							nestedRole.Vsys.Restapi.Objects.ExternalDynamicLists = o.Role.Vsys.Restapi.Objects.ExternalDynamicLists
						}
						if o.Role.Vsys.Restapi.Objects.SctpProtectionSecurityProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.SctpProtectionSecurityProfiles = o.Role.Vsys.Restapi.Objects.SctpProtectionSecurityProfiles
						}
						if o.Role.Vsys.Restapi.Objects.UrlFilteringSecurityProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.UrlFilteringSecurityProfiles = o.Role.Vsys.Restapi.Objects.UrlFilteringSecurityProfiles
						}
						if o.Role.Vsys.Restapi.Objects.SdwanErrorCorrectionProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.SdwanErrorCorrectionProfiles = o.Role.Vsys.Restapi.Objects.SdwanErrorCorrectionProfiles
						}
						if o.Role.Vsys.Restapi.Objects.Tags != nil {
							nestedRole.Vsys.Restapi.Objects.Tags = o.Role.Vsys.Restapi.Objects.Tags
						}
						if o.Role.Vsys.Restapi.Objects.CustomDataPatterns != nil {
							nestedRole.Vsys.Restapi.Objects.CustomDataPatterns = o.Role.Vsys.Restapi.Objects.CustomDataPatterns
						}
						if o.Role.Vsys.Restapi.Objects.CustomVulnerabilitySignatures != nil {
							nestedRole.Vsys.Restapi.Objects.CustomVulnerabilitySignatures = o.Role.Vsys.Restapi.Objects.CustomVulnerabilitySignatures
						}
						if o.Role.Vsys.Restapi.Objects.SdwanSaasQualityProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.SdwanSaasQualityProfiles = o.Role.Vsys.Restapi.Objects.SdwanSaasQualityProfiles
						}
						if o.Role.Vsys.Restapi.Objects.AntivirusSecurityProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.AntivirusSecurityProfiles = o.Role.Vsys.Restapi.Objects.AntivirusSecurityProfiles
						}
						if o.Role.Vsys.Restapi.Objects.CustomSpywareSignatures != nil {
							nestedRole.Vsys.Restapi.Objects.CustomSpywareSignatures = o.Role.Vsys.Restapi.Objects.CustomSpywareSignatures
						}
						if o.Role.Vsys.Restapi.Objects.DynamicUserGroups != nil {
							nestedRole.Vsys.Restapi.Objects.DynamicUserGroups = o.Role.Vsys.Restapi.Objects.DynamicUserGroups
						}
						if o.Role.Vsys.Restapi.Objects.FileBlockingSecurityProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.FileBlockingSecurityProfiles = o.Role.Vsys.Restapi.Objects.FileBlockingSecurityProfiles
						}
						if o.Role.Vsys.Restapi.Objects.GlobalprotectHipProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.GlobalprotectHipProfiles = o.Role.Vsys.Restapi.Objects.GlobalprotectHipProfiles
						}
						if o.Role.Vsys.Restapi.Objects.LogForwardingProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.LogForwardingProfiles = o.Role.Vsys.Restapi.Objects.LogForwardingProfiles
						}
						if o.Role.Vsys.Restapi.Objects.SdwanPathQualityProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.SdwanPathQualityProfiles = o.Role.Vsys.Restapi.Objects.SdwanPathQualityProfiles
						}
						if o.Role.Vsys.Restapi.Objects.AntiSpywareSecurityProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.AntiSpywareSecurityProfiles = o.Role.Vsys.Restapi.Objects.AntiSpywareSecurityProfiles
						}
						if o.Role.Vsys.Restapi.Objects.Applications != nil {
							nestedRole.Vsys.Restapi.Objects.Applications = o.Role.Vsys.Restapi.Objects.Applications
						}
						if o.Role.Vsys.Restapi.Objects.DosProtectionSecurityProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.DosProtectionSecurityProfiles = o.Role.Vsys.Restapi.Objects.DosProtectionSecurityProfiles
						}
						if o.Role.Vsys.Restapi.Objects.PacketBrokerProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.PacketBrokerProfiles = o.Role.Vsys.Restapi.Objects.PacketBrokerProfiles
						}
						if o.Role.Vsys.Restapi.Objects.SdwanTrafficDistributionProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.SdwanTrafficDistributionProfiles = o.Role.Vsys.Restapi.Objects.SdwanTrafficDistributionProfiles
						}
						if o.Role.Vsys.Restapi.Objects.AuthenticationEnforcements != nil {
							nestedRole.Vsys.Restapi.Objects.AuthenticationEnforcements = o.Role.Vsys.Restapi.Objects.AuthenticationEnforcements
						}
						if o.Role.Vsys.Restapi.Objects.DataFilteringSecurityProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.DataFilteringSecurityProfiles = o.Role.Vsys.Restapi.Objects.DataFilteringSecurityProfiles
						}
						if o.Role.Vsys.Restapi.Objects.Regions != nil {
							nestedRole.Vsys.Restapi.Objects.Regions = o.Role.Vsys.Restapi.Objects.Regions
						}
						if o.Role.Vsys.Restapi.Objects.VulnerabilityProtectionSecurityProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.VulnerabilityProtectionSecurityProfiles = o.Role.Vsys.Restapi.Objects.VulnerabilityProtectionSecurityProfiles
						}
						if o.Role.Vsys.Restapi.Objects.AddressGroups != nil {
							nestedRole.Vsys.Restapi.Objects.AddressGroups = o.Role.Vsys.Restapi.Objects.AddressGroups
						}
						if o.Role.Vsys.Restapi.Objects.ApplicationFilters != nil {
							nestedRole.Vsys.Restapi.Objects.ApplicationFilters = o.Role.Vsys.Restapi.Objects.ApplicationFilters
						}
						if o.Role.Vsys.Restapi.Objects.ApplicationGroups != nil {
							nestedRole.Vsys.Restapi.Objects.ApplicationGroups = o.Role.Vsys.Restapi.Objects.ApplicationGroups
						}
						if o.Role.Vsys.Restapi.Objects.Devices != nil {
							nestedRole.Vsys.Restapi.Objects.Devices = o.Role.Vsys.Restapi.Objects.Devices
						}
						if o.Role.Vsys.Restapi.Objects.Schedules != nil {
							nestedRole.Vsys.Restapi.Objects.Schedules = o.Role.Vsys.Restapi.Objects.Schedules
						}
						if o.Role.Vsys.Restapi.Objects.SecurityProfileGroups != nil {
							nestedRole.Vsys.Restapi.Objects.SecurityProfileGroups = o.Role.Vsys.Restapi.Objects.SecurityProfileGroups
						}
						if o.Role.Vsys.Restapi.Objects.Services != nil {
							nestedRole.Vsys.Restapi.Objects.Services = o.Role.Vsys.Restapi.Objects.Services
						}
						if o.Role.Vsys.Restapi.Objects.WildfireAnalysisSecurityProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.WildfireAnalysisSecurityProfiles = o.Role.Vsys.Restapi.Objects.WildfireAnalysisSecurityProfiles
						}
						if o.Role.Vsys.Restapi.Objects.Addresses != nil {
							nestedRole.Vsys.Restapi.Objects.Addresses = o.Role.Vsys.Restapi.Objects.Addresses
						}
						if o.Role.Vsys.Restapi.Objects.CustomUrlCategories != nil {
							nestedRole.Vsys.Restapi.Objects.CustomUrlCategories = o.Role.Vsys.Restapi.Objects.CustomUrlCategories
						}
						if o.Role.Vsys.Restapi.Objects.DecryptionProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.DecryptionProfiles = o.Role.Vsys.Restapi.Objects.DecryptionProfiles
						}
						if o.Role.Vsys.Restapi.Objects.GlobalprotectHipObjects != nil {
							nestedRole.Vsys.Restapi.Objects.GlobalprotectHipObjects = o.Role.Vsys.Restapi.Objects.GlobalprotectHipObjects
						}
						if o.Role.Vsys.Restapi.Objects.GtpProtectionSecurityProfiles != nil {
							nestedRole.Vsys.Restapi.Objects.GtpProtectionSecurityProfiles = o.Role.Vsys.Restapi.Objects.GtpProtectionSecurityProfiles
						}
						if o.Role.Vsys.Restapi.Objects.ServiceGroups != nil {
							nestedRole.Vsys.Restapi.Objects.ServiceGroups = o.Role.Vsys.Restapi.Objects.ServiceGroups
						}
					}
					if o.Role.Vsys.Restapi.Policies != nil {
						nestedRole.Vsys.Restapi.Policies = &RoleVsysRestapiPolicies{}
						if o.Role.Vsys.Restapi.Policies.Misc != nil {
							entry.Misc["RoleVsysRestapiPolicies"] = o.Role.Vsys.Restapi.Policies.Misc
						}
						if o.Role.Vsys.Restapi.Policies.QosRules != nil {
							nestedRole.Vsys.Restapi.Policies.QosRules = o.Role.Vsys.Restapi.Policies.QosRules
						}
						if o.Role.Vsys.Restapi.Policies.SecurityRules != nil {
							nestedRole.Vsys.Restapi.Policies.SecurityRules = o.Role.Vsys.Restapi.Policies.SecurityRules
						}
						if o.Role.Vsys.Restapi.Policies.TunnelInspectionRules != nil {
							nestedRole.Vsys.Restapi.Policies.TunnelInspectionRules = o.Role.Vsys.Restapi.Policies.TunnelInspectionRules
						}
						if o.Role.Vsys.Restapi.Policies.ApplicationOverrideRules != nil {
							nestedRole.Vsys.Restapi.Policies.ApplicationOverrideRules = o.Role.Vsys.Restapi.Policies.ApplicationOverrideRules
						}
						if o.Role.Vsys.Restapi.Policies.AuthenticationRules != nil {
							nestedRole.Vsys.Restapi.Policies.AuthenticationRules = o.Role.Vsys.Restapi.Policies.AuthenticationRules
						}
						if o.Role.Vsys.Restapi.Policies.DecryptionRules != nil {
							nestedRole.Vsys.Restapi.Policies.DecryptionRules = o.Role.Vsys.Restapi.Policies.DecryptionRules
						}
						if o.Role.Vsys.Restapi.Policies.NatRules != nil {
							nestedRole.Vsys.Restapi.Policies.NatRules = o.Role.Vsys.Restapi.Policies.NatRules
						}
						if o.Role.Vsys.Restapi.Policies.PolicyBasedForwardingRules != nil {
							nestedRole.Vsys.Restapi.Policies.PolicyBasedForwardingRules = o.Role.Vsys.Restapi.Policies.PolicyBasedForwardingRules
						}
						if o.Role.Vsys.Restapi.Policies.DosRules != nil {
							nestedRole.Vsys.Restapi.Policies.DosRules = o.Role.Vsys.Restapi.Policies.DosRules
						}
						if o.Role.Vsys.Restapi.Policies.NetworkPacketBrokerRules != nil {
							nestedRole.Vsys.Restapi.Policies.NetworkPacketBrokerRules = o.Role.Vsys.Restapi.Policies.NetworkPacketBrokerRules
						}
						if o.Role.Vsys.Restapi.Policies.SdwanRules != nil {
							nestedRole.Vsys.Restapi.Policies.SdwanRules = o.Role.Vsys.Restapi.Policies.SdwanRules
						}
					}
					if o.Role.Vsys.Restapi.System != nil {
						nestedRole.Vsys.Restapi.System = &RoleVsysRestapiSystem{}
						if o.Role.Vsys.Restapi.System.Misc != nil {
							entry.Misc["RoleVsysRestapiSystem"] = o.Role.Vsys.Restapi.System.Misc
						}
						if o.Role.Vsys.Restapi.System.Configuration != nil {
							nestedRole.Vsys.Restapi.System.Configuration = o.Role.Vsys.Restapi.System.Configuration
						}
					}
				}
				if o.Role.Vsys.Webui != nil {
					nestedRole.Vsys.Webui = &RoleVsysWebui{}
					if o.Role.Vsys.Webui.Misc != nil {
						entry.Misc["RoleVsysWebui"] = o.Role.Vsys.Webui.Misc
					}
					if o.Role.Vsys.Webui.Commit != nil {
						nestedRole.Vsys.Webui.Commit = &RoleVsysWebuiCommit{}
						if o.Role.Vsys.Webui.Commit.Misc != nil {
							entry.Misc["RoleVsysWebuiCommit"] = o.Role.Vsys.Webui.Commit.Misc
						}
						if o.Role.Vsys.Webui.Commit.CommitForOtherAdmins != nil {
							nestedRole.Vsys.Webui.Commit.CommitForOtherAdmins = o.Role.Vsys.Webui.Commit.CommitForOtherAdmins
						}
						if o.Role.Vsys.Webui.Commit.VirtualSystems != nil {
							nestedRole.Vsys.Webui.Commit.VirtualSystems = o.Role.Vsys.Webui.Commit.VirtualSystems
						}
					}
					if o.Role.Vsys.Webui.Objects != nil {
						nestedRole.Vsys.Webui.Objects = &RoleVsysWebuiObjects{}
						if o.Role.Vsys.Webui.Objects.Misc != nil {
							entry.Misc["RoleVsysWebuiObjects"] = o.Role.Vsys.Webui.Objects.Misc
						}
						if o.Role.Vsys.Webui.Objects.Applications != nil {
							nestedRole.Vsys.Webui.Objects.Applications = o.Role.Vsys.Webui.Objects.Applications
						}
						if o.Role.Vsys.Webui.Objects.Regions != nil {
							nestedRole.Vsys.Webui.Objects.Regions = o.Role.Vsys.Webui.Objects.Regions
						}
						if o.Role.Vsys.Webui.Objects.Sdwan != nil {
							nestedRole.Vsys.Webui.Objects.Sdwan = &RoleVsysWebuiObjectsSdwan{}
							if o.Role.Vsys.Webui.Objects.Sdwan.Misc != nil {
								entry.Misc["RoleVsysWebuiObjectsSdwan"] = o.Role.Vsys.Webui.Objects.Sdwan.Misc
							}
							if o.Role.Vsys.Webui.Objects.Sdwan.SdwanProfile != nil {
								nestedRole.Vsys.Webui.Objects.Sdwan.SdwanProfile = o.Role.Vsys.Webui.Objects.Sdwan.SdwanProfile
							}
							if o.Role.Vsys.Webui.Objects.Sdwan.SdwanSaasQualityProfile != nil {
								nestedRole.Vsys.Webui.Objects.Sdwan.SdwanSaasQualityProfile = o.Role.Vsys.Webui.Objects.Sdwan.SdwanSaasQualityProfile
							}
							if o.Role.Vsys.Webui.Objects.Sdwan.SdwanDistProfile != nil {
								nestedRole.Vsys.Webui.Objects.Sdwan.SdwanDistProfile = o.Role.Vsys.Webui.Objects.Sdwan.SdwanDistProfile
							}
							if o.Role.Vsys.Webui.Objects.Sdwan.SdwanErrorCorrectionProfile != nil {
								nestedRole.Vsys.Webui.Objects.Sdwan.SdwanErrorCorrectionProfile = o.Role.Vsys.Webui.Objects.Sdwan.SdwanErrorCorrectionProfile
							}
						}
						if o.Role.Vsys.Webui.Objects.SecurityProfileGroups != nil {
							nestedRole.Vsys.Webui.Objects.SecurityProfileGroups = o.Role.Vsys.Webui.Objects.SecurityProfileGroups
						}
						if o.Role.Vsys.Webui.Objects.ApplicationFilters != nil {
							nestedRole.Vsys.Webui.Objects.ApplicationFilters = o.Role.Vsys.Webui.Objects.ApplicationFilters
						}
						if o.Role.Vsys.Webui.Objects.Authentication != nil {
							nestedRole.Vsys.Webui.Objects.Authentication = o.Role.Vsys.Webui.Objects.Authentication
						}
						if o.Role.Vsys.Webui.Objects.CustomObjects != nil {
							nestedRole.Vsys.Webui.Objects.CustomObjects = &RoleVsysWebuiObjectsCustomObjects{}
							if o.Role.Vsys.Webui.Objects.CustomObjects.Misc != nil {
								entry.Misc["RoleVsysWebuiObjectsCustomObjects"] = o.Role.Vsys.Webui.Objects.CustomObjects.Misc
							}
							if o.Role.Vsys.Webui.Objects.CustomObjects.Spyware != nil {
								nestedRole.Vsys.Webui.Objects.CustomObjects.Spyware = o.Role.Vsys.Webui.Objects.CustomObjects.Spyware
							}
							if o.Role.Vsys.Webui.Objects.CustomObjects.UrlCategory != nil {
								nestedRole.Vsys.Webui.Objects.CustomObjects.UrlCategory = o.Role.Vsys.Webui.Objects.CustomObjects.UrlCategory
							}
							if o.Role.Vsys.Webui.Objects.CustomObjects.Vulnerability != nil {
								nestedRole.Vsys.Webui.Objects.CustomObjects.Vulnerability = o.Role.Vsys.Webui.Objects.CustomObjects.Vulnerability
							}
							if o.Role.Vsys.Webui.Objects.CustomObjects.DataPatterns != nil {
								nestedRole.Vsys.Webui.Objects.CustomObjects.DataPatterns = o.Role.Vsys.Webui.Objects.CustomObjects.DataPatterns
							}
						}
						if o.Role.Vsys.Webui.Objects.DynamicBlockLists != nil {
							nestedRole.Vsys.Webui.Objects.DynamicBlockLists = o.Role.Vsys.Webui.Objects.DynamicBlockLists
						}
						if o.Role.Vsys.Webui.Objects.GlobalProtect != nil {
							nestedRole.Vsys.Webui.Objects.GlobalProtect = &RoleVsysWebuiObjectsGlobalProtect{}
							if o.Role.Vsys.Webui.Objects.GlobalProtect.Misc != nil {
								entry.Misc["RoleVsysWebuiObjectsGlobalProtect"] = o.Role.Vsys.Webui.Objects.GlobalProtect.Misc
							}
							if o.Role.Vsys.Webui.Objects.GlobalProtect.HipObjects != nil {
								nestedRole.Vsys.Webui.Objects.GlobalProtect.HipObjects = o.Role.Vsys.Webui.Objects.GlobalProtect.HipObjects
							}
							if o.Role.Vsys.Webui.Objects.GlobalProtect.HipProfiles != nil {
								nestedRole.Vsys.Webui.Objects.GlobalProtect.HipProfiles = o.Role.Vsys.Webui.Objects.GlobalProtect.HipProfiles
							}
						}
						if o.Role.Vsys.Webui.Objects.LogForwarding != nil {
							nestedRole.Vsys.Webui.Objects.LogForwarding = o.Role.Vsys.Webui.Objects.LogForwarding
						}
						if o.Role.Vsys.Webui.Objects.PacketBrokerProfile != nil {
							nestedRole.Vsys.Webui.Objects.PacketBrokerProfile = o.Role.Vsys.Webui.Objects.PacketBrokerProfile
						}
						if o.Role.Vsys.Webui.Objects.AddressGroups != nil {
							nestedRole.Vsys.Webui.Objects.AddressGroups = o.Role.Vsys.Webui.Objects.AddressGroups
						}
						if o.Role.Vsys.Webui.Objects.Addresses != nil {
							nestedRole.Vsys.Webui.Objects.Addresses = o.Role.Vsys.Webui.Objects.Addresses
						}
						if o.Role.Vsys.Webui.Objects.ApplicationGroups != nil {
							nestedRole.Vsys.Webui.Objects.ApplicationGroups = o.Role.Vsys.Webui.Objects.ApplicationGroups
						}
						if o.Role.Vsys.Webui.Objects.ServiceGroups != nil {
							nestedRole.Vsys.Webui.Objects.ServiceGroups = o.Role.Vsys.Webui.Objects.ServiceGroups
						}
						if o.Role.Vsys.Webui.Objects.Services != nil {
							nestedRole.Vsys.Webui.Objects.Services = o.Role.Vsys.Webui.Objects.Services
						}
						if o.Role.Vsys.Webui.Objects.Tags != nil {
							nestedRole.Vsys.Webui.Objects.Tags = o.Role.Vsys.Webui.Objects.Tags
						}
						if o.Role.Vsys.Webui.Objects.Decryption != nil {
							nestedRole.Vsys.Webui.Objects.Decryption = &RoleVsysWebuiObjectsDecryption{}
							if o.Role.Vsys.Webui.Objects.Decryption.Misc != nil {
								entry.Misc["RoleVsysWebuiObjectsDecryption"] = o.Role.Vsys.Webui.Objects.Decryption.Misc
							}
							if o.Role.Vsys.Webui.Objects.Decryption.DecryptionProfile != nil {
								nestedRole.Vsys.Webui.Objects.Decryption.DecryptionProfile = o.Role.Vsys.Webui.Objects.Decryption.DecryptionProfile
							}
						}
						if o.Role.Vsys.Webui.Objects.Devices != nil {
							nestedRole.Vsys.Webui.Objects.Devices = o.Role.Vsys.Webui.Objects.Devices
						}
						if o.Role.Vsys.Webui.Objects.DynamicUserGroups != nil {
							nestedRole.Vsys.Webui.Objects.DynamicUserGroups = o.Role.Vsys.Webui.Objects.DynamicUserGroups
						}
						if o.Role.Vsys.Webui.Objects.Schedules != nil {
							nestedRole.Vsys.Webui.Objects.Schedules = o.Role.Vsys.Webui.Objects.Schedules
						}
						if o.Role.Vsys.Webui.Objects.SecurityProfiles != nil {
							nestedRole.Vsys.Webui.Objects.SecurityProfiles = &RoleVsysWebuiObjectsSecurityProfiles{}
							if o.Role.Vsys.Webui.Objects.SecurityProfiles.Misc != nil {
								entry.Misc["RoleVsysWebuiObjectsSecurityProfiles"] = o.Role.Vsys.Webui.Objects.SecurityProfiles.Misc
							}
							if o.Role.Vsys.Webui.Objects.SecurityProfiles.DosProtection != nil {
								nestedRole.Vsys.Webui.Objects.SecurityProfiles.DosProtection = o.Role.Vsys.Webui.Objects.SecurityProfiles.DosProtection
							}
							if o.Role.Vsys.Webui.Objects.SecurityProfiles.UrlFiltering != nil {
								nestedRole.Vsys.Webui.Objects.SecurityProfiles.UrlFiltering = o.Role.Vsys.Webui.Objects.SecurityProfiles.UrlFiltering
							}
							if o.Role.Vsys.Webui.Objects.SecurityProfiles.AntiSpyware != nil {
								nestedRole.Vsys.Webui.Objects.SecurityProfiles.AntiSpyware = o.Role.Vsys.Webui.Objects.SecurityProfiles.AntiSpyware
							}
							if o.Role.Vsys.Webui.Objects.SecurityProfiles.DataFiltering != nil {
								nestedRole.Vsys.Webui.Objects.SecurityProfiles.DataFiltering = o.Role.Vsys.Webui.Objects.SecurityProfiles.DataFiltering
							}
							if o.Role.Vsys.Webui.Objects.SecurityProfiles.GtpProtection != nil {
								nestedRole.Vsys.Webui.Objects.SecurityProfiles.GtpProtection = o.Role.Vsys.Webui.Objects.SecurityProfiles.GtpProtection
							}
							if o.Role.Vsys.Webui.Objects.SecurityProfiles.SctpProtection != nil {
								nestedRole.Vsys.Webui.Objects.SecurityProfiles.SctpProtection = o.Role.Vsys.Webui.Objects.SecurityProfiles.SctpProtection
							}
							if o.Role.Vsys.Webui.Objects.SecurityProfiles.VulnerabilityProtection != nil {
								nestedRole.Vsys.Webui.Objects.SecurityProfiles.VulnerabilityProtection = o.Role.Vsys.Webui.Objects.SecurityProfiles.VulnerabilityProtection
							}
							if o.Role.Vsys.Webui.Objects.SecurityProfiles.WildfireAnalysis != nil {
								nestedRole.Vsys.Webui.Objects.SecurityProfiles.WildfireAnalysis = o.Role.Vsys.Webui.Objects.SecurityProfiles.WildfireAnalysis
							}
							if o.Role.Vsys.Webui.Objects.SecurityProfiles.Antivirus != nil {
								nestedRole.Vsys.Webui.Objects.SecurityProfiles.Antivirus = o.Role.Vsys.Webui.Objects.SecurityProfiles.Antivirus
							}
							if o.Role.Vsys.Webui.Objects.SecurityProfiles.FileBlocking != nil {
								nestedRole.Vsys.Webui.Objects.SecurityProfiles.FileBlocking = o.Role.Vsys.Webui.Objects.SecurityProfiles.FileBlocking
							}
						}
					}
					if o.Role.Vsys.Webui.Privacy != nil {
						nestedRole.Vsys.Webui.Privacy = &RoleVsysWebuiPrivacy{}
						if o.Role.Vsys.Webui.Privacy.Misc != nil {
							entry.Misc["RoleVsysWebuiPrivacy"] = o.Role.Vsys.Webui.Privacy.Misc
						}
						if o.Role.Vsys.Webui.Privacy.ShowFullIpAddresses != nil {
							nestedRole.Vsys.Webui.Privacy.ShowFullIpAddresses = o.Role.Vsys.Webui.Privacy.ShowFullIpAddresses
						}
						if o.Role.Vsys.Webui.Privacy.ShowUserNamesInLogsAndReports != nil {
							nestedRole.Vsys.Webui.Privacy.ShowUserNamesInLogsAndReports = o.Role.Vsys.Webui.Privacy.ShowUserNamesInLogsAndReports
						}
						if o.Role.Vsys.Webui.Privacy.ViewPcapFiles != nil {
							nestedRole.Vsys.Webui.Privacy.ViewPcapFiles = o.Role.Vsys.Webui.Privacy.ViewPcapFiles
						}
					}
					if o.Role.Vsys.Webui.Policies != nil {
						nestedRole.Vsys.Webui.Policies = &RoleVsysWebuiPolicies{}
						if o.Role.Vsys.Webui.Policies.Misc != nil {
							entry.Misc["RoleVsysWebuiPolicies"] = o.Role.Vsys.Webui.Policies.Misc
						}
						if o.Role.Vsys.Webui.Policies.AuthenticationRulebase != nil {
							nestedRole.Vsys.Webui.Policies.AuthenticationRulebase = o.Role.Vsys.Webui.Policies.AuthenticationRulebase
						}
						if o.Role.Vsys.Webui.Policies.DosRulebase != nil {
							nestedRole.Vsys.Webui.Policies.DosRulebase = o.Role.Vsys.Webui.Policies.DosRulebase
						}
						if o.Role.Vsys.Webui.Policies.PbfRulebase != nil {
							nestedRole.Vsys.Webui.Policies.PbfRulebase = o.Role.Vsys.Webui.Policies.PbfRulebase
						}
						if o.Role.Vsys.Webui.Policies.TunnelInspectRulebase != nil {
							nestedRole.Vsys.Webui.Policies.TunnelInspectRulebase = o.Role.Vsys.Webui.Policies.TunnelInspectRulebase
						}
						if o.Role.Vsys.Webui.Policies.ApplicationOverrideRulebase != nil {
							nestedRole.Vsys.Webui.Policies.ApplicationOverrideRulebase = o.Role.Vsys.Webui.Policies.ApplicationOverrideRulebase
						}
						if o.Role.Vsys.Webui.Policies.NetworkPacketBrokerRulebase != nil {
							nestedRole.Vsys.Webui.Policies.NetworkPacketBrokerRulebase = o.Role.Vsys.Webui.Policies.NetworkPacketBrokerRulebase
						}
						if o.Role.Vsys.Webui.Policies.QosRulebase != nil {
							nestedRole.Vsys.Webui.Policies.QosRulebase = o.Role.Vsys.Webui.Policies.QosRulebase
						}
						if o.Role.Vsys.Webui.Policies.RuleHitCountReset != nil {
							nestedRole.Vsys.Webui.Policies.RuleHitCountReset = o.Role.Vsys.Webui.Policies.RuleHitCountReset
						}
						if o.Role.Vsys.Webui.Policies.SdwanRulebase != nil {
							nestedRole.Vsys.Webui.Policies.SdwanRulebase = o.Role.Vsys.Webui.Policies.SdwanRulebase
						}
						if o.Role.Vsys.Webui.Policies.SecurityRulebase != nil {
							nestedRole.Vsys.Webui.Policies.SecurityRulebase = o.Role.Vsys.Webui.Policies.SecurityRulebase
						}
						if o.Role.Vsys.Webui.Policies.SslDecryptionRulebase != nil {
							nestedRole.Vsys.Webui.Policies.SslDecryptionRulebase = o.Role.Vsys.Webui.Policies.SslDecryptionRulebase
						}
						if o.Role.Vsys.Webui.Policies.NatRulebase != nil {
							nestedRole.Vsys.Webui.Policies.NatRulebase = o.Role.Vsys.Webui.Policies.NatRulebase
						}
					}
					if o.Role.Vsys.Webui.Save != nil {
						nestedRole.Vsys.Webui.Save = &RoleVsysWebuiSave{}
						if o.Role.Vsys.Webui.Save.Misc != nil {
							entry.Misc["RoleVsysWebuiSave"] = o.Role.Vsys.Webui.Save.Misc
						}
						if o.Role.Vsys.Webui.Save.PartialSave != nil {
							nestedRole.Vsys.Webui.Save.PartialSave = o.Role.Vsys.Webui.Save.PartialSave
						}
						if o.Role.Vsys.Webui.Save.SaveForOtherAdmins != nil {
							nestedRole.Vsys.Webui.Save.SaveForOtherAdmins = o.Role.Vsys.Webui.Save.SaveForOtherAdmins
						}
						if o.Role.Vsys.Webui.Save.ObjectLevelChanges != nil {
							nestedRole.Vsys.Webui.Save.ObjectLevelChanges = o.Role.Vsys.Webui.Save.ObjectLevelChanges
						}
					}
					if o.Role.Vsys.Webui.Acc != nil {
						nestedRole.Vsys.Webui.Acc = o.Role.Vsys.Webui.Acc
					}
					if o.Role.Vsys.Webui.Dashboard != nil {
						nestedRole.Vsys.Webui.Dashboard = o.Role.Vsys.Webui.Dashboard
					}
					if o.Role.Vsys.Webui.Device != nil {
						nestedRole.Vsys.Webui.Device = &RoleVsysWebuiDevice{}
						if o.Role.Vsys.Webui.Device.Misc != nil {
							entry.Misc["RoleVsysWebuiDevice"] = o.Role.Vsys.Webui.Device.Misc
						}
						if o.Role.Vsys.Webui.Device.AuthenticationSequence != nil {
							nestedRole.Vsys.Webui.Device.AuthenticationSequence = o.Role.Vsys.Webui.Device.AuthenticationSequence
						}
						if o.Role.Vsys.Webui.Device.LocalUserDatabase != nil {
							nestedRole.Vsys.Webui.Device.LocalUserDatabase = &RoleVsysWebuiDeviceLocalUserDatabase{}
							if o.Role.Vsys.Webui.Device.LocalUserDatabase.Misc != nil {
								entry.Misc["RoleVsysWebuiDeviceLocalUserDatabase"] = o.Role.Vsys.Webui.Device.LocalUserDatabase.Misc
							}
							if o.Role.Vsys.Webui.Device.LocalUserDatabase.UserGroups != nil {
								nestedRole.Vsys.Webui.Device.LocalUserDatabase.UserGroups = o.Role.Vsys.Webui.Device.LocalUserDatabase.UserGroups
							}
							if o.Role.Vsys.Webui.Device.LocalUserDatabase.Users != nil {
								nestedRole.Vsys.Webui.Device.LocalUserDatabase.Users = o.Role.Vsys.Webui.Device.LocalUserDatabase.Users
							}
						}
						if o.Role.Vsys.Webui.Device.Troubleshooting != nil {
							nestedRole.Vsys.Webui.Device.Troubleshooting = o.Role.Vsys.Webui.Device.Troubleshooting
						}
						if o.Role.Vsys.Webui.Device.VmInfoSource != nil {
							nestedRole.Vsys.Webui.Device.VmInfoSource = o.Role.Vsys.Webui.Device.VmInfoSource
						}
						if o.Role.Vsys.Webui.Device.BlockPages != nil {
							nestedRole.Vsys.Webui.Device.BlockPages = o.Role.Vsys.Webui.Device.BlockPages
						}
						if o.Role.Vsys.Webui.Device.CertificateManagement != nil {
							nestedRole.Vsys.Webui.Device.CertificateManagement = &RoleVsysWebuiDeviceCertificateManagement{}
							if o.Role.Vsys.Webui.Device.CertificateManagement.Misc != nil {
								entry.Misc["RoleVsysWebuiDeviceCertificateManagement"] = o.Role.Vsys.Webui.Device.CertificateManagement.Misc
							}
							if o.Role.Vsys.Webui.Device.CertificateManagement.Scep != nil {
								nestedRole.Vsys.Webui.Device.CertificateManagement.Scep = o.Role.Vsys.Webui.Device.CertificateManagement.Scep
							}
							if o.Role.Vsys.Webui.Device.CertificateManagement.SshServiceProfile != nil {
								nestedRole.Vsys.Webui.Device.CertificateManagement.SshServiceProfile = o.Role.Vsys.Webui.Device.CertificateManagement.SshServiceProfile
							}
							if o.Role.Vsys.Webui.Device.CertificateManagement.SslDecryptionExclusion != nil {
								nestedRole.Vsys.Webui.Device.CertificateManagement.SslDecryptionExclusion = o.Role.Vsys.Webui.Device.CertificateManagement.SslDecryptionExclusion
							}
							if o.Role.Vsys.Webui.Device.CertificateManagement.SslTlsServiceProfile != nil {
								nestedRole.Vsys.Webui.Device.CertificateManagement.SslTlsServiceProfile = o.Role.Vsys.Webui.Device.CertificateManagement.SslTlsServiceProfile
							}
							if o.Role.Vsys.Webui.Device.CertificateManagement.CertificateProfile != nil {
								nestedRole.Vsys.Webui.Device.CertificateManagement.CertificateProfile = o.Role.Vsys.Webui.Device.CertificateManagement.CertificateProfile
							}
							if o.Role.Vsys.Webui.Device.CertificateManagement.Certificates != nil {
								nestedRole.Vsys.Webui.Device.CertificateManagement.Certificates = o.Role.Vsys.Webui.Device.CertificateManagement.Certificates
							}
							if o.Role.Vsys.Webui.Device.CertificateManagement.OcspResponder != nil {
								nestedRole.Vsys.Webui.Device.CertificateManagement.OcspResponder = o.Role.Vsys.Webui.Device.CertificateManagement.OcspResponder
							}
						}
						if o.Role.Vsys.Webui.Device.DeviceQuarantine != nil {
							nestedRole.Vsys.Webui.Device.DeviceQuarantine = o.Role.Vsys.Webui.Device.DeviceQuarantine
						}
						if o.Role.Vsys.Webui.Device.LogSettings != nil {
							nestedRole.Vsys.Webui.Device.LogSettings = &RoleVsysWebuiDeviceLogSettings{}
							if o.Role.Vsys.Webui.Device.LogSettings.Misc != nil {
								entry.Misc["RoleVsysWebuiDeviceLogSettings"] = o.Role.Vsys.Webui.Device.LogSettings.Misc
							}
							if o.Role.Vsys.Webui.Device.LogSettings.Correlation != nil {
								nestedRole.Vsys.Webui.Device.LogSettings.Correlation = o.Role.Vsys.Webui.Device.LogSettings.Correlation
							}
							if o.Role.Vsys.Webui.Device.LogSettings.Globalprotect != nil {
								nestedRole.Vsys.Webui.Device.LogSettings.Globalprotect = o.Role.Vsys.Webui.Device.LogSettings.Globalprotect
							}
							if o.Role.Vsys.Webui.Device.LogSettings.Hipmatch != nil {
								nestedRole.Vsys.Webui.Device.LogSettings.Hipmatch = o.Role.Vsys.Webui.Device.LogSettings.Hipmatch
							}
							if o.Role.Vsys.Webui.Device.LogSettings.Iptag != nil {
								nestedRole.Vsys.Webui.Device.LogSettings.Iptag = o.Role.Vsys.Webui.Device.LogSettings.Iptag
							}
							if o.Role.Vsys.Webui.Device.LogSettings.System != nil {
								nestedRole.Vsys.Webui.Device.LogSettings.System = o.Role.Vsys.Webui.Device.LogSettings.System
							}
							if o.Role.Vsys.Webui.Device.LogSettings.UserId != nil {
								nestedRole.Vsys.Webui.Device.LogSettings.UserId = o.Role.Vsys.Webui.Device.LogSettings.UserId
							}
							if o.Role.Vsys.Webui.Device.LogSettings.Config != nil {
								nestedRole.Vsys.Webui.Device.LogSettings.Config = o.Role.Vsys.Webui.Device.LogSettings.Config
							}
						}
						if o.Role.Vsys.Webui.Device.PolicyRecommendations != nil {
							nestedRole.Vsys.Webui.Device.PolicyRecommendations = &RoleVsysWebuiDevicePolicyRecommendations{}
							if o.Role.Vsys.Webui.Device.PolicyRecommendations.Misc != nil {
								entry.Misc["RoleVsysWebuiDevicePolicyRecommendations"] = o.Role.Vsys.Webui.Device.PolicyRecommendations.Misc
							}
							if o.Role.Vsys.Webui.Device.PolicyRecommendations.Iot != nil {
								nestedRole.Vsys.Webui.Device.PolicyRecommendations.Iot = o.Role.Vsys.Webui.Device.PolicyRecommendations.Iot
							}
							if o.Role.Vsys.Webui.Device.PolicyRecommendations.Saas != nil {
								nestedRole.Vsys.Webui.Device.PolicyRecommendations.Saas = o.Role.Vsys.Webui.Device.PolicyRecommendations.Saas
							}
						}
						if o.Role.Vsys.Webui.Device.Setup != nil {
							nestedRole.Vsys.Webui.Device.Setup = &RoleVsysWebuiDeviceSetup{}
							if o.Role.Vsys.Webui.Device.Setup.Misc != nil {
								entry.Misc["RoleVsysWebuiDeviceSetup"] = o.Role.Vsys.Webui.Device.Setup.Misc
							}
							if o.Role.Vsys.Webui.Device.Setup.Hsm != nil {
								nestedRole.Vsys.Webui.Device.Setup.Hsm = o.Role.Vsys.Webui.Device.Setup.Hsm
							}
							if o.Role.Vsys.Webui.Device.Setup.Management != nil {
								nestedRole.Vsys.Webui.Device.Setup.Management = o.Role.Vsys.Webui.Device.Setup.Management
							}
							if o.Role.Vsys.Webui.Device.Setup.Services != nil {
								nestedRole.Vsys.Webui.Device.Setup.Services = o.Role.Vsys.Webui.Device.Setup.Services
							}
							if o.Role.Vsys.Webui.Device.Setup.ContentId != nil {
								nestedRole.Vsys.Webui.Device.Setup.ContentId = o.Role.Vsys.Webui.Device.Setup.ContentId
							}
							if o.Role.Vsys.Webui.Device.Setup.Operations != nil {
								nestedRole.Vsys.Webui.Device.Setup.Operations = o.Role.Vsys.Webui.Device.Setup.Operations
							}
							if o.Role.Vsys.Webui.Device.Setup.Session != nil {
								nestedRole.Vsys.Webui.Device.Setup.Session = o.Role.Vsys.Webui.Device.Setup.Session
							}
							if o.Role.Vsys.Webui.Device.Setup.Telemetry != nil {
								nestedRole.Vsys.Webui.Device.Setup.Telemetry = o.Role.Vsys.Webui.Device.Setup.Telemetry
							}
							if o.Role.Vsys.Webui.Device.Setup.Wildfire != nil {
								nestedRole.Vsys.Webui.Device.Setup.Wildfire = o.Role.Vsys.Webui.Device.Setup.Wildfire
							}
							if o.Role.Vsys.Webui.Device.Setup.Interfaces != nil {
								nestedRole.Vsys.Webui.Device.Setup.Interfaces = o.Role.Vsys.Webui.Device.Setup.Interfaces
							}
						}
						if o.Role.Vsys.Webui.Device.DataRedistribution != nil {
							nestedRole.Vsys.Webui.Device.DataRedistribution = o.Role.Vsys.Webui.Device.DataRedistribution
						}
						if o.Role.Vsys.Webui.Device.UserIdentification != nil {
							nestedRole.Vsys.Webui.Device.UserIdentification = o.Role.Vsys.Webui.Device.UserIdentification
						}
						if o.Role.Vsys.Webui.Device.DhcpSyslogServer != nil {
							nestedRole.Vsys.Webui.Device.DhcpSyslogServer = o.Role.Vsys.Webui.Device.DhcpSyslogServer
						}
						if o.Role.Vsys.Webui.Device.Administrators != nil {
							nestedRole.Vsys.Webui.Device.Administrators = o.Role.Vsys.Webui.Device.Administrators
						}
						if o.Role.Vsys.Webui.Device.AuthenticationProfile != nil {
							nestedRole.Vsys.Webui.Device.AuthenticationProfile = o.Role.Vsys.Webui.Device.AuthenticationProfile
						}
						if o.Role.Vsys.Webui.Device.ServerProfile != nil {
							nestedRole.Vsys.Webui.Device.ServerProfile = &RoleVsysWebuiDeviceServerProfile{}
							if o.Role.Vsys.Webui.Device.ServerProfile.Misc != nil {
								entry.Misc["RoleVsysWebuiDeviceServerProfile"] = o.Role.Vsys.Webui.Device.ServerProfile.Misc
							}
							if o.Role.Vsys.Webui.Device.ServerProfile.Netflow != nil {
								nestedRole.Vsys.Webui.Device.ServerProfile.Netflow = o.Role.Vsys.Webui.Device.ServerProfile.Netflow
							}
							if o.Role.Vsys.Webui.Device.ServerProfile.Radius != nil {
								nestedRole.Vsys.Webui.Device.ServerProfile.Radius = o.Role.Vsys.Webui.Device.ServerProfile.Radius
							}
							if o.Role.Vsys.Webui.Device.ServerProfile.SamlIdp != nil {
								nestedRole.Vsys.Webui.Device.ServerProfile.SamlIdp = o.Role.Vsys.Webui.Device.ServerProfile.SamlIdp
							}
							if o.Role.Vsys.Webui.Device.ServerProfile.SnmpTrap != nil {
								nestedRole.Vsys.Webui.Device.ServerProfile.SnmpTrap = o.Role.Vsys.Webui.Device.ServerProfile.SnmpTrap
							}
							if o.Role.Vsys.Webui.Device.ServerProfile.Dns != nil {
								nestedRole.Vsys.Webui.Device.ServerProfile.Dns = o.Role.Vsys.Webui.Device.ServerProfile.Dns
							}
							if o.Role.Vsys.Webui.Device.ServerProfile.Http != nil {
								nestedRole.Vsys.Webui.Device.ServerProfile.Http = o.Role.Vsys.Webui.Device.ServerProfile.Http
							}
							if o.Role.Vsys.Webui.Device.ServerProfile.Kerberos != nil {
								nestedRole.Vsys.Webui.Device.ServerProfile.Kerberos = o.Role.Vsys.Webui.Device.ServerProfile.Kerberos
							}
							if o.Role.Vsys.Webui.Device.ServerProfile.Scp != nil {
								nestedRole.Vsys.Webui.Device.ServerProfile.Scp = o.Role.Vsys.Webui.Device.ServerProfile.Scp
							}
							if o.Role.Vsys.Webui.Device.ServerProfile.Syslog != nil {
								nestedRole.Vsys.Webui.Device.ServerProfile.Syslog = o.Role.Vsys.Webui.Device.ServerProfile.Syslog
							}
							if o.Role.Vsys.Webui.Device.ServerProfile.Tacplus != nil {
								nestedRole.Vsys.Webui.Device.ServerProfile.Tacplus = o.Role.Vsys.Webui.Device.ServerProfile.Tacplus
							}
							if o.Role.Vsys.Webui.Device.ServerProfile.Email != nil {
								nestedRole.Vsys.Webui.Device.ServerProfile.Email = o.Role.Vsys.Webui.Device.ServerProfile.Email
							}
							if o.Role.Vsys.Webui.Device.ServerProfile.Ldap != nil {
								nestedRole.Vsys.Webui.Device.ServerProfile.Ldap = o.Role.Vsys.Webui.Device.ServerProfile.Ldap
							}
							if o.Role.Vsys.Webui.Device.ServerProfile.Mfa != nil {
								nestedRole.Vsys.Webui.Device.ServerProfile.Mfa = o.Role.Vsys.Webui.Device.ServerProfile.Mfa
							}
						}
					}
					if o.Role.Vsys.Webui.Monitor != nil {
						nestedRole.Vsys.Webui.Monitor = &RoleVsysWebuiMonitor{}
						if o.Role.Vsys.Webui.Monitor.Misc != nil {
							entry.Misc["RoleVsysWebuiMonitor"] = o.Role.Vsys.Webui.Monitor.Misc
						}
						if o.Role.Vsys.Webui.Monitor.SessionBrowser != nil {
							nestedRole.Vsys.Webui.Monitor.SessionBrowser = o.Role.Vsys.Webui.Monitor.SessionBrowser
						}
						if o.Role.Vsys.Webui.Monitor.ViewCustomReports != nil {
							nestedRole.Vsys.Webui.Monitor.ViewCustomReports = o.Role.Vsys.Webui.Monitor.ViewCustomReports
						}
						if o.Role.Vsys.Webui.Monitor.AppScope != nil {
							nestedRole.Vsys.Webui.Monitor.AppScope = o.Role.Vsys.Webui.Monitor.AppScope
						}
						if o.Role.Vsys.Webui.Monitor.CustomReports != nil {
							nestedRole.Vsys.Webui.Monitor.CustomReports = &RoleVsysWebuiMonitorCustomReports{}
							if o.Role.Vsys.Webui.Monitor.CustomReports.Misc != nil {
								entry.Misc["RoleVsysWebuiMonitorCustomReports"] = o.Role.Vsys.Webui.Monitor.CustomReports.Misc
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.Hipmatch != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.Hipmatch = o.Role.Vsys.Webui.Monitor.CustomReports.Hipmatch
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.Iptag != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.Iptag = o.Role.Vsys.Webui.Monitor.CustomReports.Iptag
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.ThreatSummary != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.ThreatSummary = o.Role.Vsys.Webui.Monitor.CustomReports.ThreatSummary
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.TrafficSummary != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.TrafficSummary = o.Role.Vsys.Webui.Monitor.CustomReports.TrafficSummary
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.UrlLog != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.UrlLog = o.Role.Vsys.Webui.Monitor.CustomReports.UrlLog
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.DataFilteringLog != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.DataFilteringLog = o.Role.Vsys.Webui.Monitor.CustomReports.DataFilteringLog
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.DecryptionSummary != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.DecryptionSummary = o.Role.Vsys.Webui.Monitor.CustomReports.DecryptionSummary
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.Globalprotect != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.Globalprotect = o.Role.Vsys.Webui.Monitor.CustomReports.Globalprotect
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.UrlSummary != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.UrlSummary = o.Role.Vsys.Webui.Monitor.CustomReports.UrlSummary
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.WildfireLog != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.WildfireLog = o.Role.Vsys.Webui.Monitor.CustomReports.WildfireLog
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.DecryptionLog != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.DecryptionLog = o.Role.Vsys.Webui.Monitor.CustomReports.DecryptionLog
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.ThreatLog != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.ThreatLog = o.Role.Vsys.Webui.Monitor.CustomReports.ThreatLog
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.TunnelSummary != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.TunnelSummary = o.Role.Vsys.Webui.Monitor.CustomReports.TunnelSummary
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.ApplicationStatistics != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.ApplicationStatistics = o.Role.Vsys.Webui.Monitor.CustomReports.ApplicationStatistics
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.Auth != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.Auth = o.Role.Vsys.Webui.Monitor.CustomReports.Auth
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.TunnelLog != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.TunnelLog = o.Role.Vsys.Webui.Monitor.CustomReports.TunnelLog
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.SctpSummary != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.SctpSummary = o.Role.Vsys.Webui.Monitor.CustomReports.SctpSummary
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.TrafficLog != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.TrafficLog = o.Role.Vsys.Webui.Monitor.CustomReports.TrafficLog
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.Userid != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.Userid = o.Role.Vsys.Webui.Monitor.CustomReports.Userid
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.GtpLog != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.GtpLog = o.Role.Vsys.Webui.Monitor.CustomReports.GtpLog
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.GtpSummary != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.GtpSummary = o.Role.Vsys.Webui.Monitor.CustomReports.GtpSummary
							}
							if o.Role.Vsys.Webui.Monitor.CustomReports.SctpLog != nil {
								nestedRole.Vsys.Webui.Monitor.CustomReports.SctpLog = o.Role.Vsys.Webui.Monitor.CustomReports.SctpLog
							}
						}
						if o.Role.Vsys.Webui.Monitor.Logs != nil {
							nestedRole.Vsys.Webui.Monitor.Logs = &RoleVsysWebuiMonitorLogs{}
							if o.Role.Vsys.Webui.Monitor.Logs.Misc != nil {
								entry.Misc["RoleVsysWebuiMonitorLogs"] = o.Role.Vsys.Webui.Monitor.Logs.Misc
							}
							if o.Role.Vsys.Webui.Monitor.Logs.Threat != nil {
								nestedRole.Vsys.Webui.Monitor.Logs.Threat = o.Role.Vsys.Webui.Monitor.Logs.Threat
							}
							if o.Role.Vsys.Webui.Monitor.Logs.Wildfire != nil {
								nestedRole.Vsys.Webui.Monitor.Logs.Wildfire = o.Role.Vsys.Webui.Monitor.Logs.Wildfire
							}
							if o.Role.Vsys.Webui.Monitor.Logs.DataFiltering != nil {
								nestedRole.Vsys.Webui.Monitor.Logs.DataFiltering = o.Role.Vsys.Webui.Monitor.Logs.DataFiltering
							}
							if o.Role.Vsys.Webui.Monitor.Logs.Globalprotect != nil {
								nestedRole.Vsys.Webui.Monitor.Logs.Globalprotect = o.Role.Vsys.Webui.Monitor.Logs.Globalprotect
							}
							if o.Role.Vsys.Webui.Monitor.Logs.Hipmatch != nil {
								nestedRole.Vsys.Webui.Monitor.Logs.Hipmatch = o.Role.Vsys.Webui.Monitor.Logs.Hipmatch
							}
							if o.Role.Vsys.Webui.Monitor.Logs.Url != nil {
								nestedRole.Vsys.Webui.Monitor.Logs.Url = o.Role.Vsys.Webui.Monitor.Logs.Url
							}
							if o.Role.Vsys.Webui.Monitor.Logs.Userid != nil {
								nestedRole.Vsys.Webui.Monitor.Logs.Userid = o.Role.Vsys.Webui.Monitor.Logs.Userid
							}
							if o.Role.Vsys.Webui.Monitor.Logs.Tunnel != nil {
								nestedRole.Vsys.Webui.Monitor.Logs.Tunnel = o.Role.Vsys.Webui.Monitor.Logs.Tunnel
							}
							if o.Role.Vsys.Webui.Monitor.Logs.Authentication != nil {
								nestedRole.Vsys.Webui.Monitor.Logs.Authentication = o.Role.Vsys.Webui.Monitor.Logs.Authentication
							}
							if o.Role.Vsys.Webui.Monitor.Logs.Decryption != nil {
								nestedRole.Vsys.Webui.Monitor.Logs.Decryption = o.Role.Vsys.Webui.Monitor.Logs.Decryption
							}
							if o.Role.Vsys.Webui.Monitor.Logs.Gtp != nil {
								nestedRole.Vsys.Webui.Monitor.Logs.Gtp = o.Role.Vsys.Webui.Monitor.Logs.Gtp
							}
							if o.Role.Vsys.Webui.Monitor.Logs.Iptag != nil {
								nestedRole.Vsys.Webui.Monitor.Logs.Iptag = o.Role.Vsys.Webui.Monitor.Logs.Iptag
							}
							if o.Role.Vsys.Webui.Monitor.Logs.Traffic != nil {
								nestedRole.Vsys.Webui.Monitor.Logs.Traffic = o.Role.Vsys.Webui.Monitor.Logs.Traffic
							}
							if o.Role.Vsys.Webui.Monitor.Logs.Sctp != nil {
								nestedRole.Vsys.Webui.Monitor.Logs.Sctp = o.Role.Vsys.Webui.Monitor.Logs.Sctp
							}
						}
						if o.Role.Vsys.Webui.Monitor.PdfReports != nil {
							nestedRole.Vsys.Webui.Monitor.PdfReports = &RoleVsysWebuiMonitorPdfReports{}
							if o.Role.Vsys.Webui.Monitor.PdfReports.Misc != nil {
								entry.Misc["RoleVsysWebuiMonitorPdfReports"] = o.Role.Vsys.Webui.Monitor.PdfReports.Misc
							}
							if o.Role.Vsys.Webui.Monitor.PdfReports.SaasApplicationUsageReport != nil {
								nestedRole.Vsys.Webui.Monitor.PdfReports.SaasApplicationUsageReport = o.Role.Vsys.Webui.Monitor.PdfReports.SaasApplicationUsageReport
							}
							if o.Role.Vsys.Webui.Monitor.PdfReports.UserActivityReport != nil {
								nestedRole.Vsys.Webui.Monitor.PdfReports.UserActivityReport = o.Role.Vsys.Webui.Monitor.PdfReports.UserActivityReport
							}
							if o.Role.Vsys.Webui.Monitor.PdfReports.EmailScheduler != nil {
								nestedRole.Vsys.Webui.Monitor.PdfReports.EmailScheduler = o.Role.Vsys.Webui.Monitor.PdfReports.EmailScheduler
							}
							if o.Role.Vsys.Webui.Monitor.PdfReports.ManagePdfSummary != nil {
								nestedRole.Vsys.Webui.Monitor.PdfReports.ManagePdfSummary = o.Role.Vsys.Webui.Monitor.PdfReports.ManagePdfSummary
							}
							if o.Role.Vsys.Webui.Monitor.PdfReports.PdfSummaryReports != nil {
								nestedRole.Vsys.Webui.Monitor.PdfReports.PdfSummaryReports = o.Role.Vsys.Webui.Monitor.PdfReports.PdfSummaryReports
							}
							if o.Role.Vsys.Webui.Monitor.PdfReports.ReportGroups != nil {
								nestedRole.Vsys.Webui.Monitor.PdfReports.ReportGroups = o.Role.Vsys.Webui.Monitor.PdfReports.ReportGroups
							}
						}
						if o.Role.Vsys.Webui.Monitor.AutomatedCorrelationEngine != nil {
							nestedRole.Vsys.Webui.Monitor.AutomatedCorrelationEngine = &RoleVsysWebuiMonitorAutomatedCorrelationEngine{}
							if o.Role.Vsys.Webui.Monitor.AutomatedCorrelationEngine.Misc != nil {
								entry.Misc["RoleVsysWebuiMonitorAutomatedCorrelationEngine"] = o.Role.Vsys.Webui.Monitor.AutomatedCorrelationEngine.Misc
							}
							if o.Role.Vsys.Webui.Monitor.AutomatedCorrelationEngine.CorrelatedEvents != nil {
								nestedRole.Vsys.Webui.Monitor.AutomatedCorrelationEngine.CorrelatedEvents = o.Role.Vsys.Webui.Monitor.AutomatedCorrelationEngine.CorrelatedEvents
							}
							if o.Role.Vsys.Webui.Monitor.AutomatedCorrelationEngine.CorrelationObjects != nil {
								nestedRole.Vsys.Webui.Monitor.AutomatedCorrelationEngine.CorrelationObjects = o.Role.Vsys.Webui.Monitor.AutomatedCorrelationEngine.CorrelationObjects
							}
						}
						if o.Role.Vsys.Webui.Monitor.BlockIpList != nil {
							nestedRole.Vsys.Webui.Monitor.BlockIpList = o.Role.Vsys.Webui.Monitor.BlockIpList
						}
						if o.Role.Vsys.Webui.Monitor.ExternalLogs != nil {
							nestedRole.Vsys.Webui.Monitor.ExternalLogs = o.Role.Vsys.Webui.Monitor.ExternalLogs
						}
					}
					if o.Role.Vsys.Webui.Network != nil {
						nestedRole.Vsys.Webui.Network = &RoleVsysWebuiNetwork{}
						if o.Role.Vsys.Webui.Network.Misc != nil {
							entry.Misc["RoleVsysWebuiNetwork"] = o.Role.Vsys.Webui.Network.Misc
						}
						if o.Role.Vsys.Webui.Network.SdwanInterfaceProfile != nil {
							nestedRole.Vsys.Webui.Network.SdwanInterfaceProfile = o.Role.Vsys.Webui.Network.SdwanInterfaceProfile
						}
						if o.Role.Vsys.Webui.Network.Zones != nil {
							nestedRole.Vsys.Webui.Network.Zones = o.Role.Vsys.Webui.Network.Zones
						}
						if o.Role.Vsys.Webui.Network.GlobalProtect != nil {
							nestedRole.Vsys.Webui.Network.GlobalProtect = &RoleVsysWebuiNetworkGlobalProtect{}
							if o.Role.Vsys.Webui.Network.GlobalProtect.Misc != nil {
								entry.Misc["RoleVsysWebuiNetworkGlobalProtect"] = o.Role.Vsys.Webui.Network.GlobalProtect.Misc
							}
							if o.Role.Vsys.Webui.Network.GlobalProtect.Gateways != nil {
								nestedRole.Vsys.Webui.Network.GlobalProtect.Gateways = o.Role.Vsys.Webui.Network.GlobalProtect.Gateways
							}
							if o.Role.Vsys.Webui.Network.GlobalProtect.Mdm != nil {
								nestedRole.Vsys.Webui.Network.GlobalProtect.Mdm = o.Role.Vsys.Webui.Network.GlobalProtect.Mdm
							}
							if o.Role.Vsys.Webui.Network.GlobalProtect.Portals != nil {
								nestedRole.Vsys.Webui.Network.GlobalProtect.Portals = o.Role.Vsys.Webui.Network.GlobalProtect.Portals
							}
							if o.Role.Vsys.Webui.Network.GlobalProtect.ClientlessAppGroups != nil {
								nestedRole.Vsys.Webui.Network.GlobalProtect.ClientlessAppGroups = o.Role.Vsys.Webui.Network.GlobalProtect.ClientlessAppGroups
							}
							if o.Role.Vsys.Webui.Network.GlobalProtect.ClientlessApps != nil {
								nestedRole.Vsys.Webui.Network.GlobalProtect.ClientlessApps = o.Role.Vsys.Webui.Network.GlobalProtect.ClientlessApps
							}
						}
					}
					if o.Role.Vsys.Webui.Operations != nil {
						nestedRole.Vsys.Webui.Operations = &RoleVsysWebuiOperations{}
						if o.Role.Vsys.Webui.Operations.Misc != nil {
							entry.Misc["RoleVsysWebuiOperations"] = o.Role.Vsys.Webui.Operations.Misc
						}
						if o.Role.Vsys.Webui.Operations.Reboot != nil {
							nestedRole.Vsys.Webui.Operations.Reboot = o.Role.Vsys.Webui.Operations.Reboot
						}
						if o.Role.Vsys.Webui.Operations.DownloadCoreFiles != nil {
							nestedRole.Vsys.Webui.Operations.DownloadCoreFiles = o.Role.Vsys.Webui.Operations.DownloadCoreFiles
						}
						if o.Role.Vsys.Webui.Operations.DownloadPcapFiles != nil {
							nestedRole.Vsys.Webui.Operations.DownloadPcapFiles = o.Role.Vsys.Webui.Operations.DownloadPcapFiles
						}
						if o.Role.Vsys.Webui.Operations.GenerateStatsDumpFile != nil {
							nestedRole.Vsys.Webui.Operations.GenerateStatsDumpFile = o.Role.Vsys.Webui.Operations.GenerateStatsDumpFile
						}
						if o.Role.Vsys.Webui.Operations.GenerateTechSupportFile != nil {
							nestedRole.Vsys.Webui.Operations.GenerateTechSupportFile = o.Role.Vsys.Webui.Operations.GenerateTechSupportFile
						}
					}
					if o.Role.Vsys.Webui.Tasks != nil {
						nestedRole.Vsys.Webui.Tasks = o.Role.Vsys.Webui.Tasks
					}
					if o.Role.Vsys.Webui.Validate != nil {
						nestedRole.Vsys.Webui.Validate = o.Role.Vsys.Webui.Validate
					}
				}
				if o.Role.Vsys.Xmlapi != nil {
					nestedRole.Vsys.Xmlapi = &RoleVsysXmlapi{}
					if o.Role.Vsys.Xmlapi.Misc != nil {
						entry.Misc["RoleVsysXmlapi"] = o.Role.Vsys.Xmlapi.Misc
					}
					if o.Role.Vsys.Xmlapi.Import != nil {
						nestedRole.Vsys.Xmlapi.Import = o.Role.Vsys.Xmlapi.Import
					}
					if o.Role.Vsys.Xmlapi.Iot != nil {
						nestedRole.Vsys.Xmlapi.Iot = o.Role.Vsys.Xmlapi.Iot
					}
					if o.Role.Vsys.Xmlapi.Log != nil {
						nestedRole.Vsys.Xmlapi.Log = o.Role.Vsys.Xmlapi.Log
					}
					if o.Role.Vsys.Xmlapi.Op != nil {
						nestedRole.Vsys.Xmlapi.Op = o.Role.Vsys.Xmlapi.Op
					}
					if o.Role.Vsys.Xmlapi.Commit != nil {
						nestedRole.Vsys.Xmlapi.Commit = o.Role.Vsys.Xmlapi.Commit
					}
					if o.Role.Vsys.Xmlapi.Config != nil {
						nestedRole.Vsys.Xmlapi.Config = o.Role.Vsys.Xmlapi.Config
					}
					if o.Role.Vsys.Xmlapi.UserId != nil {
						nestedRole.Vsys.Xmlapi.UserId = o.Role.Vsys.Xmlapi.UserId
					}
					if o.Role.Vsys.Xmlapi.Export != nil {
						nestedRole.Vsys.Xmlapi.Export = o.Role.Vsys.Xmlapi.Export
					}
					if o.Role.Vsys.Xmlapi.Report != nil {
						nestedRole.Vsys.Xmlapi.Report = o.Role.Vsys.Xmlapi.Report
					}
				}
			}
		}
		entry.Role = nestedRole

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
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !matchRole(a.Role, b.Role) {
		return false
	}

	return true
}

func matchRoleDeviceRestapiDevice(a *RoleDeviceRestapiDevice, b *RoleDeviceRestapiDevice) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.SnmpTrapServerProfiles, b.SnmpTrapServerProfiles) {
		return false
	}
	if !util.StringsMatch(a.SyslogServerProfiles, b.SyslogServerProfiles) {
		return false
	}
	if !util.StringsMatch(a.VirtualSystems, b.VirtualSystems) {
		return false
	}
	if !util.StringsMatch(a.EmailServerProfiles, b.EmailServerProfiles) {
		return false
	}
	if !util.StringsMatch(a.HttpServerProfiles, b.HttpServerProfiles) {
		return false
	}
	if !util.StringsMatch(a.LdapServerProfiles, b.LdapServerProfiles) {
		return false
	}
	if !util.StringsMatch(a.LogInterfaceSetting, b.LogInterfaceSetting) {
		return false
	}
	return true
}
func matchRoleDeviceRestapiNetwork(a *RoleDeviceRestapiNetwork, b *RoleDeviceRestapiNetwork) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.GlobalprotectIpsecCryptoNetworkProfiles, b.GlobalprotectIpsecCryptoNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(a.LoopbackInterfaces, b.LoopbackInterfaces) {
		return false
	}
	if !util.StringsMatch(a.QosInterfaces, b.QosInterfaces) {
		return false
	}
	if !util.StringsMatch(a.SdwanInterfaces, b.SdwanInterfaces) {
		return false
	}
	if !util.StringsMatch(a.Vlans, b.Vlans) {
		return false
	}
	if !util.StringsMatch(a.AggregateEthernetInterfaces, b.AggregateEthernetInterfaces) {
		return false
	}
	if !util.StringsMatch(a.GreTunnels, b.GreTunnels) {
		return false
	}
	if !util.StringsMatch(a.IkeCryptoNetworkProfiles, b.IkeCryptoNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(a.IpsecTunnels, b.IpsecTunnels) {
		return false
	}
	if !util.StringsMatch(a.LldpNetworkProfiles, b.LldpNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(a.VlanInterfaces, b.VlanInterfaces) {
		return false
	}
	if !util.StringsMatch(a.Zones, b.Zones) {
		return false
	}
	if !util.StringsMatch(a.BfdNetworkProfiles, b.BfdNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(a.GlobalprotectGateways, b.GlobalprotectGateways) {
		return false
	}
	if !util.StringsMatch(a.GlobalprotectMdmServers, b.GlobalprotectMdmServers) {
		return false
	}
	if !util.StringsMatch(a.GlobalprotectClientlessAppGroups, b.GlobalprotectClientlessAppGroups) {
		return false
	}
	if !util.StringsMatch(a.TunnelInterfaces, b.TunnelInterfaces) {
		return false
	}
	if !util.StringsMatch(a.VirtualRouters, b.VirtualRouters) {
		return false
	}
	if !util.StringsMatch(a.VirtualWires, b.VirtualWires) {
		return false
	}
	if !util.StringsMatch(a.ZoneProtectionNetworkProfiles, b.ZoneProtectionNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(a.InterfaceManagementNetworkProfiles, b.InterfaceManagementNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(a.Lldp, b.Lldp) {
		return false
	}
	if !util.StringsMatch(a.GlobalprotectPortals, b.GlobalprotectPortals) {
		return false
	}
	if !util.StringsMatch(a.IkeGatewayNetworkProfiles, b.IkeGatewayNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(a.IpsecCryptoNetworkProfiles, b.IpsecCryptoNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(a.QosNetworkProfiles, b.QosNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(a.TunnelMonitorNetworkProfiles, b.TunnelMonitorNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(a.BgpRoutingProfiles, b.BgpRoutingProfiles) {
		return false
	}
	if !util.StringsMatch(a.DnsProxies, b.DnsProxies) {
		return false
	}
	if !util.StringsMatch(a.EthernetInterfaces, b.EthernetInterfaces) {
		return false
	}
	if !util.StringsMatch(a.GlobalprotectClientlessApps, b.GlobalprotectClientlessApps) {
		return false
	}
	if !util.StringsMatch(a.SdwanInterfaceProfiles, b.SdwanInterfaceProfiles) {
		return false
	}
	if !util.StringsMatch(a.DhcpServers, b.DhcpServers) {
		return false
	}
	if !util.StringsMatch(a.LogicalRouters, b.LogicalRouters) {
		return false
	}
	if !util.StringsMatch(a.DhcpRelays, b.DhcpRelays) {
		return false
	}
	return true
}
func matchRoleDeviceRestapiObjects(a *RoleDeviceRestapiObjects, b *RoleDeviceRestapiObjects) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ExternalDynamicLists, b.ExternalDynamicLists) {
		return false
	}
	if !util.StringsMatch(a.AddressGroups, b.AddressGroups) {
		return false
	}
	if !util.StringsMatch(a.AntivirusSecurityProfiles, b.AntivirusSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.SecurityProfileGroups, b.SecurityProfileGroups) {
		return false
	}
	if !util.StringsMatch(a.PacketBrokerProfiles, b.PacketBrokerProfiles) {
		return false
	}
	if !util.StringsMatch(a.Schedules, b.Schedules) {
		return false
	}
	if !util.StringsMatch(a.ServiceGroups, b.ServiceGroups) {
		return false
	}
	if !util.StringsMatch(a.Addresses, b.Addresses) {
		return false
	}
	if !util.StringsMatch(a.ApplicationFilters, b.ApplicationFilters) {
		return false
	}
	if !util.StringsMatch(a.AuthenticationEnforcements, b.AuthenticationEnforcements) {
		return false
	}
	if !util.StringsMatch(a.FileBlockingSecurityProfiles, b.FileBlockingSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.GlobalprotectHipProfiles, b.GlobalprotectHipProfiles) {
		return false
	}
	if !util.StringsMatch(a.LogForwardingProfiles, b.LogForwardingProfiles) {
		return false
	}
	if !util.StringsMatch(a.SctpProtectionSecurityProfiles, b.SctpProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.SdwanSaasQualityProfiles, b.SdwanSaasQualityProfiles) {
		return false
	}
	if !util.StringsMatch(a.AntiSpywareSecurityProfiles, b.AntiSpywareSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.ApplicationGroups, b.ApplicationGroups) {
		return false
	}
	if !util.StringsMatch(a.CustomDataPatterns, b.CustomDataPatterns) {
		return false
	}
	if !util.StringsMatch(a.DataFilteringSecurityProfiles, b.DataFilteringSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.WildfireAnalysisSecurityProfiles, b.WildfireAnalysisSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.CustomSpywareSignatures, b.CustomSpywareSignatures) {
		return false
	}
	if !util.StringsMatch(a.DecryptionProfiles, b.DecryptionProfiles) {
		return false
	}
	if !util.StringsMatch(a.DynamicUserGroups, b.DynamicUserGroups) {
		return false
	}
	if !util.StringsMatch(a.UrlFilteringSecurityProfiles, b.UrlFilteringSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.SdwanPathQualityProfiles, b.SdwanPathQualityProfiles) {
		return false
	}
	if !util.StringsMatch(a.SdwanTrafficDistributionProfiles, b.SdwanTrafficDistributionProfiles) {
		return false
	}
	if !util.StringsMatch(a.Services, b.Services) {
		return false
	}
	if !util.StringsMatch(a.Tags, b.Tags) {
		return false
	}
	if !util.StringsMatch(a.VulnerabilityProtectionSecurityProfiles, b.VulnerabilityProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.Applications, b.Applications) {
		return false
	}
	if !util.StringsMatch(a.CustomVulnerabilitySignatures, b.CustomVulnerabilitySignatures) {
		return false
	}
	if !util.StringsMatch(a.GlobalprotectHipObjects, b.GlobalprotectHipObjects) {
		return false
	}
	if !util.StringsMatch(a.SdwanErrorCorrectionProfiles, b.SdwanErrorCorrectionProfiles) {
		return false
	}
	if !util.StringsMatch(a.Regions, b.Regions) {
		return false
	}
	if !util.StringsMatch(a.CustomUrlCategories, b.CustomUrlCategories) {
		return false
	}
	if !util.StringsMatch(a.Devices, b.Devices) {
		return false
	}
	if !util.StringsMatch(a.DosProtectionSecurityProfiles, b.DosProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.GtpProtectionSecurityProfiles, b.GtpProtectionSecurityProfiles) {
		return false
	}
	return true
}
func matchRoleDeviceRestapiPolicies(a *RoleDeviceRestapiPolicies, b *RoleDeviceRestapiPolicies) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.AuthenticationRules, b.AuthenticationRules) {
		return false
	}
	if !util.StringsMatch(a.DosRules, b.DosRules) {
		return false
	}
	if !util.StringsMatch(a.NatRules, b.NatRules) {
		return false
	}
	if !util.StringsMatch(a.NetworkPacketBrokerRules, b.NetworkPacketBrokerRules) {
		return false
	}
	if !util.StringsMatch(a.TunnelInspectionRules, b.TunnelInspectionRules) {
		return false
	}
	if !util.StringsMatch(a.ApplicationOverrideRules, b.ApplicationOverrideRules) {
		return false
	}
	if !util.StringsMatch(a.DecryptionRules, b.DecryptionRules) {
		return false
	}
	if !util.StringsMatch(a.PolicyBasedForwardingRules, b.PolicyBasedForwardingRules) {
		return false
	}
	if !util.StringsMatch(a.QosRules, b.QosRules) {
		return false
	}
	if !util.StringsMatch(a.SdwanRules, b.SdwanRules) {
		return false
	}
	if !util.StringsMatch(a.SecurityRules, b.SecurityRules) {
		return false
	}
	return true
}
func matchRoleDeviceRestapiSystem(a *RoleDeviceRestapiSystem, b *RoleDeviceRestapiSystem) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Configuration, b.Configuration) {
		return false
	}
	return true
}
func matchRoleDeviceRestapi(a *RoleDeviceRestapi, b *RoleDeviceRestapi) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoleDeviceRestapiPolicies(a.Policies, b.Policies) {
		return false
	}
	if !matchRoleDeviceRestapiSystem(a.System, b.System) {
		return false
	}
	if !matchRoleDeviceRestapiDevice(a.Device, b.Device) {
		return false
	}
	if !matchRoleDeviceRestapiNetwork(a.Network, b.Network) {
		return false
	}
	if !matchRoleDeviceRestapiObjects(a.Objects, b.Objects) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiNetworkNetworkProfiles(a *RoleDeviceWebuiNetworkNetworkProfiles, b *RoleDeviceWebuiNetworkNetworkProfiles) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.TunnelMonitor, b.TunnelMonitor) {
		return false
	}
	if !util.StringsMatch(a.ZoneProtection, b.ZoneProtection) {
		return false
	}
	if !util.StringsMatch(a.BfdProfile, b.BfdProfile) {
		return false
	}
	if !util.StringsMatch(a.IkeCrypto, b.IkeCrypto) {
		return false
	}
	if !util.StringsMatch(a.QosProfile, b.QosProfile) {
		return false
	}
	if !util.StringsMatch(a.IpsecCrypto, b.IpsecCrypto) {
		return false
	}
	if !util.StringsMatch(a.LldpProfile, b.LldpProfile) {
		return false
	}
	if !util.StringsMatch(a.GpAppIpsecCrypto, b.GpAppIpsecCrypto) {
		return false
	}
	if !util.StringsMatch(a.IkeGateways, b.IkeGateways) {
		return false
	}
	if !util.StringsMatch(a.InterfaceMgmt, b.InterfaceMgmt) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiNetworkRoutingRoutingProfiles(a *RoleDeviceWebuiNetworkRoutingRoutingProfiles, b *RoleDeviceWebuiNetworkRoutingRoutingProfiles) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Ospfv3, b.Ospfv3) {
		return false
	}
	if !util.StringsMatch(a.Ripv2, b.Ripv2) {
		return false
	}
	if !util.StringsMatch(a.Bfd, b.Bfd) {
		return false
	}
	if !util.StringsMatch(a.Bgp, b.Bgp) {
		return false
	}
	if !util.StringsMatch(a.Filters, b.Filters) {
		return false
	}
	if !util.StringsMatch(a.Multicast, b.Multicast) {
		return false
	}
	if !util.StringsMatch(a.Ospf, b.Ospf) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiNetworkRouting(a *RoleDeviceWebuiNetworkRouting, b *RoleDeviceWebuiNetworkRouting) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.LogicalRouters, b.LogicalRouters) {
		return false
	}
	if !matchRoleDeviceWebuiNetworkRoutingRoutingProfiles(a.RoutingProfiles, b.RoutingProfiles) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiNetworkGlobalProtect(a *RoleDeviceWebuiNetworkGlobalProtect, b *RoleDeviceWebuiNetworkGlobalProtect) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ClientlessAppGroups, b.ClientlessAppGroups) {
		return false
	}
	if !util.StringsMatch(a.ClientlessApps, b.ClientlessApps) {
		return false
	}
	if !util.StringsMatch(a.Gateways, b.Gateways) {
		return false
	}
	if !util.StringsMatch(a.Mdm, b.Mdm) {
		return false
	}
	if !util.StringsMatch(a.Portals, b.Portals) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiNetwork(a *RoleDeviceWebuiNetwork, b *RoleDeviceWebuiNetwork) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Qos, b.Qos) {
		return false
	}
	if !util.StringsMatch(a.VirtualRouters, b.VirtualRouters) {
		return false
	}
	if !util.StringsMatch(a.Dhcp, b.Dhcp) {
		return false
	}
	if !matchRoleDeviceWebuiNetworkGlobalProtect(a.GlobalProtect, b.GlobalProtect) {
		return false
	}
	if !util.StringsMatch(a.GreTunnels, b.GreTunnels) {
		return false
	}
	if !util.StringsMatch(a.Lldp, b.Lldp) {
		return false
	}
	if !util.StringsMatch(a.DnsProxy, b.DnsProxy) {
		return false
	}
	if !util.StringsMatch(a.Interfaces, b.Interfaces) {
		return false
	}
	if !util.StringsMatch(a.Zones, b.Zones) {
		return false
	}
	if !util.StringsMatch(a.IpsecTunnels, b.IpsecTunnels) {
		return false
	}
	if !matchRoleDeviceWebuiNetworkNetworkProfiles(a.NetworkProfiles, b.NetworkProfiles) {
		return false
	}
	if !matchRoleDeviceWebuiNetworkRouting(a.Routing, b.Routing) {
		return false
	}
	if !util.StringsMatch(a.VirtualWires, b.VirtualWires) {
		return false
	}
	if !util.StringsMatch(a.SdwanInterfaceProfile, b.SdwanInterfaceProfile) {
		return false
	}
	if !util.StringsMatch(a.Vlans, b.Vlans) {
		return false
	}
	if !util.StringsMatch(a.SecureWebGateway, b.SecureWebGateway) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiPrivacy(a *RoleDeviceWebuiPrivacy, b *RoleDeviceWebuiPrivacy) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ShowFullIpAddresses, b.ShowFullIpAddresses) {
		return false
	}
	if !util.StringsMatch(a.ShowUserNamesInLogsAndReports, b.ShowUserNamesInLogsAndReports) {
		return false
	}
	if !util.StringsMatch(a.ViewPcapFiles, b.ViewPcapFiles) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiCommit(a *RoleDeviceWebuiCommit, b *RoleDeviceWebuiCommit) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Device, b.Device) {
		return false
	}
	if !util.StringsMatch(a.ObjectLevelChanges, b.ObjectLevelChanges) {
		return false
	}
	if !util.StringsMatch(a.CommitForOtherAdmins, b.CommitForOtherAdmins) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiDeviceLogSettings(a *RoleDeviceWebuiDeviceLogSettings, b *RoleDeviceWebuiDeviceLogSettings) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Correlation, b.Correlation) {
		return false
	}
	if !util.StringsMatch(a.Globalprotect, b.Globalprotect) {
		return false
	}
	if !util.StringsMatch(a.Iptag, b.Iptag) {
		return false
	}
	if !util.StringsMatch(a.ManageLog, b.ManageLog) {
		return false
	}
	if !util.StringsMatch(a.UserId, b.UserId) {
		return false
	}
	if !util.StringsMatch(a.CcAlarm, b.CcAlarm) {
		return false
	}
	if !util.StringsMatch(a.Config, b.Config) {
		return false
	}
	if !util.StringsMatch(a.Hipmatch, b.Hipmatch) {
		return false
	}
	if !util.StringsMatch(a.System, b.System) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiDeviceSetup(a *RoleDeviceWebuiDeviceSetup, b *RoleDeviceWebuiDeviceSetup) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Hsm, b.Hsm) {
		return false
	}
	if !util.StringsMatch(a.Interfaces, b.Interfaces) {
		return false
	}
	if !util.StringsMatch(a.Operations, b.Operations) {
		return false
	}
	if !util.StringsMatch(a.Telemetry, b.Telemetry) {
		return false
	}
	if !util.StringsMatch(a.ContentId, b.ContentId) {
		return false
	}
	if !util.StringsMatch(a.Management, b.Management) {
		return false
	}
	if !util.StringsMatch(a.Services, b.Services) {
		return false
	}
	if !util.StringsMatch(a.Session, b.Session) {
		return false
	}
	if !util.StringsMatch(a.Wildfire, b.Wildfire) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiDeviceCertificateManagement(a *RoleDeviceWebuiDeviceCertificateManagement, b *RoleDeviceWebuiDeviceCertificateManagement) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Scep, b.Scep) {
		return false
	}
	if !util.StringsMatch(a.SshServiceProfile, b.SshServiceProfile) {
		return false
	}
	if !util.StringsMatch(a.SslDecryptionExclusion, b.SslDecryptionExclusion) {
		return false
	}
	if !util.StringsMatch(a.SslTlsServiceProfile, b.SslTlsServiceProfile) {
		return false
	}
	if !util.StringsMatch(a.CertificateProfile, b.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(a.Certificates, b.Certificates) {
		return false
	}
	if !util.StringsMatch(a.OcspResponder, b.OcspResponder) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiDevicePolicyRecommendations(a *RoleDeviceWebuiDevicePolicyRecommendations, b *RoleDeviceWebuiDevicePolicyRecommendations) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Iot, b.Iot) {
		return false
	}
	if !util.StringsMatch(a.Saas, b.Saas) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiDeviceServerProfile(a *RoleDeviceWebuiDeviceServerProfile, b *RoleDeviceWebuiDeviceServerProfile) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Tacplus, b.Tacplus) {
		return false
	}
	if !util.StringsMatch(a.Dns, b.Dns) {
		return false
	}
	if !util.StringsMatch(a.Email, b.Email) {
		return false
	}
	if !util.StringsMatch(a.Netflow, b.Netflow) {
		return false
	}
	if !util.StringsMatch(a.SamlIdp, b.SamlIdp) {
		return false
	}
	if !util.StringsMatch(a.Radius, b.Radius) {
		return false
	}
	if !util.StringsMatch(a.Scp, b.Scp) {
		return false
	}
	if !util.StringsMatch(a.SnmpTrap, b.SnmpTrap) {
		return false
	}
	if !util.StringsMatch(a.Syslog, b.Syslog) {
		return false
	}
	if !util.StringsMatch(a.Http, b.Http) {
		return false
	}
	if !util.StringsMatch(a.Kerberos, b.Kerberos) {
		return false
	}
	if !util.StringsMatch(a.Ldap, b.Ldap) {
		return false
	}
	if !util.StringsMatch(a.Mfa, b.Mfa) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiDeviceLocalUserDatabase(a *RoleDeviceWebuiDeviceLocalUserDatabase, b *RoleDeviceWebuiDeviceLocalUserDatabase) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Users, b.Users) {
		return false
	}
	if !util.StringsMatch(a.UserGroups, b.UserGroups) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiDevice(a *RoleDeviceWebuiDevice, b *RoleDeviceWebuiDevice) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.HighAvailability, b.HighAvailability) {
		return false
	}
	if !matchRoleDeviceWebuiDeviceLocalUserDatabase(a.LocalUserDatabase, b.LocalUserDatabase) {
		return false
	}
	if !util.StringsMatch(a.MasterKey, b.MasterKey) {
		return false
	}
	if !util.StringsMatch(a.Plugins, b.Plugins) {
		return false
	}
	if !matchRoleDeviceWebuiDevicePolicyRecommendations(a.PolicyRecommendations, b.PolicyRecommendations) {
		return false
	}
	if !matchRoleDeviceWebuiDeviceServerProfile(a.ServerProfile, b.ServerProfile) {
		return false
	}
	if !util.StringsMatch(a.ConfigAudit, b.ConfigAudit) {
		return false
	}
	if !util.StringsMatch(a.DataRedistribution, b.DataRedistribution) {
		return false
	}
	if !util.StringsMatch(a.Support, b.Support) {
		return false
	}
	if !matchRoleDeviceWebuiDeviceLogSettings(a.LogSettings, b.LogSettings) {
		return false
	}
	if !util.StringsMatch(a.VmInfoSource, b.VmInfoSource) {
		return false
	}
	if !util.StringsMatch(a.AdminRoles, b.AdminRoles) {
		return false
	}
	if !util.StringsMatch(a.DynamicUpdates, b.DynamicUpdates) {
		return false
	}
	if !util.StringsMatch(a.VirtualSystems, b.VirtualSystems) {
		return false
	}
	if !util.StringsMatch(a.AuthenticationProfile, b.AuthenticationProfile) {
		return false
	}
	if !util.StringsMatch(a.Licenses, b.Licenses) {
		return false
	}
	if !util.StringsMatch(a.BlockPages, b.BlockPages) {
		return false
	}
	if !util.StringsMatch(a.GlobalProtectClient, b.GlobalProtectClient) {
		return false
	}
	if !util.StringsMatch(a.AuthenticationSequence, b.AuthenticationSequence) {
		return false
	}
	if !util.StringsMatch(a.AccessDomain, b.AccessDomain) {
		return false
	}
	if !util.StringsMatch(a.Administrators, b.Administrators) {
		return false
	}
	if !matchRoleDeviceWebuiDeviceSetup(a.Setup, b.Setup) {
		return false
	}
	if !util.StringsMatch(a.Software, b.Software) {
		return false
	}
	if !util.StringsMatch(a.Troubleshooting, b.Troubleshooting) {
		return false
	}
	if !util.StringsMatch(a.UserIdentification, b.UserIdentification) {
		return false
	}
	if !util.StringsMatch(a.DhcpSyslogServer, b.DhcpSyslogServer) {
		return false
	}
	if !matchRoleDeviceWebuiDeviceCertificateManagement(a.CertificateManagement, b.CertificateManagement) {
		return false
	}
	if !util.StringsMatch(a.ScheduledLogExport, b.ScheduledLogExport) {
		return false
	}
	if !util.StringsMatch(a.SharedGateways, b.SharedGateways) {
		return false
	}
	if !util.StringsMatch(a.DeviceQuarantine, b.DeviceQuarantine) {
		return false
	}
	if !util.StringsMatch(a.LogFwdCard, b.LogFwdCard) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiMonitorLogs(a *RoleDeviceWebuiMonitorLogs, b *RoleDeviceWebuiMonitorLogs) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Configuration, b.Configuration) {
		return false
	}
	if !util.StringsMatch(a.Sctp, b.Sctp) {
		return false
	}
	if !util.StringsMatch(a.Url, b.Url) {
		return false
	}
	if !util.StringsMatch(a.Alarm, b.Alarm) {
		return false
	}
	if !util.StringsMatch(a.Authentication, b.Authentication) {
		return false
	}
	if !util.StringsMatch(a.Globalprotect, b.Globalprotect) {
		return false
	}
	if !util.StringsMatch(a.Hipmatch, b.Hipmatch) {
		return false
	}
	if !util.StringsMatch(a.System, b.System) {
		return false
	}
	if !util.StringsMatch(a.Traffic, b.Traffic) {
		return false
	}
	if !util.StringsMatch(a.Tunnel, b.Tunnel) {
		return false
	}
	if !util.StringsMatch(a.Wildfire, b.Wildfire) {
		return false
	}
	if !util.StringsMatch(a.Decryption, b.Decryption) {
		return false
	}
	if !util.StringsMatch(a.Gtp, b.Gtp) {
		return false
	}
	if !util.StringsMatch(a.DataFiltering, b.DataFiltering) {
		return false
	}
	if !util.StringsMatch(a.Iptag, b.Iptag) {
		return false
	}
	if !util.StringsMatch(a.Threat, b.Threat) {
		return false
	}
	if !util.StringsMatch(a.Userid, b.Userid) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiMonitorPdfReports(a *RoleDeviceWebuiMonitorPdfReports, b *RoleDeviceWebuiMonitorPdfReports) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.EmailScheduler, b.EmailScheduler) {
		return false
	}
	if !util.StringsMatch(a.ManagePdfSummary, b.ManagePdfSummary) {
		return false
	}
	if !util.StringsMatch(a.PdfSummaryReports, b.PdfSummaryReports) {
		return false
	}
	if !util.StringsMatch(a.ReportGroups, b.ReportGroups) {
		return false
	}
	if !util.StringsMatch(a.SaasApplicationUsageReport, b.SaasApplicationUsageReport) {
		return false
	}
	if !util.StringsMatch(a.UserActivityReport, b.UserActivityReport) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiMonitorAutomatedCorrelationEngine(a *RoleDeviceWebuiMonitorAutomatedCorrelationEngine, b *RoleDeviceWebuiMonitorAutomatedCorrelationEngine) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.CorrelationObjects, b.CorrelationObjects) {
		return false
	}
	if !util.StringsMatch(a.CorrelatedEvents, b.CorrelatedEvents) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiMonitorCustomReports(a *RoleDeviceWebuiMonitorCustomReports, b *RoleDeviceWebuiMonitorCustomReports) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.GtpLog, b.GtpLog) {
		return false
	}
	if !util.StringsMatch(a.ThreatSummary, b.ThreatSummary) {
		return false
	}
	if !util.StringsMatch(a.UrlSummary, b.UrlSummary) {
		return false
	}
	if !util.StringsMatch(a.WildfireLog, b.WildfireLog) {
		return false
	}
	if !util.StringsMatch(a.ApplicationStatistics, b.ApplicationStatistics) {
		return false
	}
	if !util.StringsMatch(a.Hipmatch, b.Hipmatch) {
		return false
	}
	if !util.StringsMatch(a.ThreatLog, b.ThreatLog) {
		return false
	}
	if !util.StringsMatch(a.TrafficLog, b.TrafficLog) {
		return false
	}
	if !util.StringsMatch(a.TunnelLog, b.TunnelLog) {
		return false
	}
	if !util.StringsMatch(a.TunnelSummary, b.TunnelSummary) {
		return false
	}
	if !util.StringsMatch(a.UrlLog, b.UrlLog) {
		return false
	}
	if !util.StringsMatch(a.Auth, b.Auth) {
		return false
	}
	if !util.StringsMatch(a.DataFilteringLog, b.DataFilteringLog) {
		return false
	}
	if !util.StringsMatch(a.Globalprotect, b.Globalprotect) {
		return false
	}
	if !util.StringsMatch(a.GtpSummary, b.GtpSummary) {
		return false
	}
	if !util.StringsMatch(a.SctpSummary, b.SctpSummary) {
		return false
	}
	if !util.StringsMatch(a.TrafficSummary, b.TrafficSummary) {
		return false
	}
	if !util.StringsMatch(a.DecryptionLog, b.DecryptionLog) {
		return false
	}
	if !util.StringsMatch(a.DecryptionSummary, b.DecryptionSummary) {
		return false
	}
	if !util.StringsMatch(a.Iptag, b.Iptag) {
		return false
	}
	if !util.StringsMatch(a.SctpLog, b.SctpLog) {
		return false
	}
	if !util.StringsMatch(a.Userid, b.Userid) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiMonitor(a *RoleDeviceWebuiMonitor, b *RoleDeviceWebuiMonitor) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.AppScope, b.AppScope) {
		return false
	}
	if !util.StringsMatch(a.ApplicationReports, b.ApplicationReports) {
		return false
	}
	if !util.StringsMatch(a.GtpReports, b.GtpReports) {
		return false
	}
	if !util.StringsMatch(a.ViewCustomReports, b.ViewCustomReports) {
		return false
	}
	if !util.StringsMatch(a.Botnet, b.Botnet) {
		return false
	}
	if !util.StringsMatch(a.ExternalLogs, b.ExternalLogs) {
		return false
	}
	if !matchRoleDeviceWebuiMonitorLogs(a.Logs, b.Logs) {
		return false
	}
	if !util.StringsMatch(a.PacketCapture, b.PacketCapture) {
		return false
	}
	if !util.StringsMatch(a.TrafficReports, b.TrafficReports) {
		return false
	}
	if !util.StringsMatch(a.UrlFilteringReports, b.UrlFilteringReports) {
		return false
	}
	if !matchRoleDeviceWebuiMonitorAutomatedCorrelationEngine(a.AutomatedCorrelationEngine, b.AutomatedCorrelationEngine) {
		return false
	}
	if !util.StringsMatch(a.BlockIpList, b.BlockIpList) {
		return false
	}
	if !matchRoleDeviceWebuiMonitorPdfReports(a.PdfReports, b.PdfReports) {
		return false
	}
	if !util.StringsMatch(a.SctpReports, b.SctpReports) {
		return false
	}
	if !util.StringsMatch(a.SessionBrowser, b.SessionBrowser) {
		return false
	}
	if !matchRoleDeviceWebuiMonitorCustomReports(a.CustomReports, b.CustomReports) {
		return false
	}
	if !util.StringsMatch(a.ThreatReports, b.ThreatReports) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiOperations(a *RoleDeviceWebuiOperations, b *RoleDeviceWebuiOperations) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DownloadCoreFiles, b.DownloadCoreFiles) {
		return false
	}
	if !util.StringsMatch(a.DownloadPcapFiles, b.DownloadPcapFiles) {
		return false
	}
	if !util.StringsMatch(a.GenerateStatsDumpFile, b.GenerateStatsDumpFile) {
		return false
	}
	if !util.StringsMatch(a.GenerateTechSupportFile, b.GenerateTechSupportFile) {
		return false
	}
	if !util.StringsMatch(a.Reboot, b.Reboot) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiSave(a *RoleDeviceWebuiSave, b *RoleDeviceWebuiSave) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ObjectLevelChanges, b.ObjectLevelChanges) {
		return false
	}
	if !util.StringsMatch(a.PartialSave, b.PartialSave) {
		return false
	}
	if !util.StringsMatch(a.SaveForOtherAdmins, b.SaveForOtherAdmins) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiGlobal(a *RoleDeviceWebuiGlobal, b *RoleDeviceWebuiGlobal) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.SystemAlarms, b.SystemAlarms) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiObjectsGlobalProtect(a *RoleDeviceWebuiObjectsGlobalProtect, b *RoleDeviceWebuiObjectsGlobalProtect) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.HipObjects, b.HipObjects) {
		return false
	}
	if !util.StringsMatch(a.HipProfiles, b.HipProfiles) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiObjectsDecryption(a *RoleDeviceWebuiObjectsDecryption, b *RoleDeviceWebuiObjectsDecryption) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DecryptionProfile, b.DecryptionProfile) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiObjectsCustomObjects(a *RoleDeviceWebuiObjectsCustomObjects, b *RoleDeviceWebuiObjectsCustomObjects) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DataPatterns, b.DataPatterns) {
		return false
	}
	if !util.StringsMatch(a.Spyware, b.Spyware) {
		return false
	}
	if !util.StringsMatch(a.UrlCategory, b.UrlCategory) {
		return false
	}
	if !util.StringsMatch(a.Vulnerability, b.Vulnerability) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiObjectsSecurityProfiles(a *RoleDeviceWebuiObjectsSecurityProfiles, b *RoleDeviceWebuiObjectsSecurityProfiles) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Antivirus, b.Antivirus) {
		return false
	}
	if !util.StringsMatch(a.DosProtection, b.DosProtection) {
		return false
	}
	if !util.StringsMatch(a.FileBlocking, b.FileBlocking) {
		return false
	}
	if !util.StringsMatch(a.GtpProtection, b.GtpProtection) {
		return false
	}
	if !util.StringsMatch(a.VulnerabilityProtection, b.VulnerabilityProtection) {
		return false
	}
	if !util.StringsMatch(a.WildfireAnalysis, b.WildfireAnalysis) {
		return false
	}
	if !util.StringsMatch(a.AntiSpyware, b.AntiSpyware) {
		return false
	}
	if !util.StringsMatch(a.DataFiltering, b.DataFiltering) {
		return false
	}
	if !util.StringsMatch(a.SctpProtection, b.SctpProtection) {
		return false
	}
	if !util.StringsMatch(a.UrlFiltering, b.UrlFiltering) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiObjectsSdwan(a *RoleDeviceWebuiObjectsSdwan, b *RoleDeviceWebuiObjectsSdwan) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.SdwanErrorCorrectionProfile, b.SdwanErrorCorrectionProfile) {
		return false
	}
	if !util.StringsMatch(a.SdwanProfile, b.SdwanProfile) {
		return false
	}
	if !util.StringsMatch(a.SdwanSaasQualityProfile, b.SdwanSaasQualityProfile) {
		return false
	}
	if !util.StringsMatch(a.SdwanDistProfile, b.SdwanDistProfile) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiObjects(a *RoleDeviceWebuiObjects, b *RoleDeviceWebuiObjects) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Applications, b.Applications) {
		return false
	}
	if !util.StringsMatch(a.LogForwarding, b.LogForwarding) {
		return false
	}
	if !matchRoleDeviceWebuiObjectsSdwan(a.Sdwan, b.Sdwan) {
		return false
	}
	if !matchRoleDeviceWebuiObjectsSecurityProfiles(a.SecurityProfiles, b.SecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.ApplicationFilters, b.ApplicationFilters) {
		return false
	}
	if !matchRoleDeviceWebuiObjectsDecryption(a.Decryption, b.Decryption) {
		return false
	}
	if !util.StringsMatch(a.Devices, b.Devices) {
		return false
	}
	if !matchRoleDeviceWebuiObjectsGlobalProtect(a.GlobalProtect, b.GlobalProtect) {
		return false
	}
	if !util.StringsMatch(a.PacketBrokerProfile, b.PacketBrokerProfile) {
		return false
	}
	if !util.StringsMatch(a.Regions, b.Regions) {
		return false
	}
	if !util.StringsMatch(a.Schedules, b.Schedules) {
		return false
	}
	if !util.StringsMatch(a.ServiceGroups, b.ServiceGroups) {
		return false
	}
	if !util.StringsMatch(a.Tags, b.Tags) {
		return false
	}
	if !util.StringsMatch(a.Addresses, b.Addresses) {
		return false
	}
	if !util.StringsMatch(a.Authentication, b.Authentication) {
		return false
	}
	if !matchRoleDeviceWebuiObjectsCustomObjects(a.CustomObjects, b.CustomObjects) {
		return false
	}
	if !util.StringsMatch(a.SecurityProfileGroups, b.SecurityProfileGroups) {
		return false
	}
	if !util.StringsMatch(a.AddressGroups, b.AddressGroups) {
		return false
	}
	if !util.StringsMatch(a.ApplicationGroups, b.ApplicationGroups) {
		return false
	}
	if !util.StringsMatch(a.DynamicBlockLists, b.DynamicBlockLists) {
		return false
	}
	if !util.StringsMatch(a.DynamicUserGroups, b.DynamicUserGroups) {
		return false
	}
	if !util.StringsMatch(a.Services, b.Services) {
		return false
	}
	return true
}
func matchRoleDeviceWebuiPolicies(a *RoleDeviceWebuiPolicies, b *RoleDeviceWebuiPolicies) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.AuthenticationRulebase, b.AuthenticationRulebase) {
		return false
	}
	if !util.StringsMatch(a.NatRulebase, b.NatRulebase) {
		return false
	}
	if !util.StringsMatch(a.NetworkPacketBrokerRulebase, b.NetworkPacketBrokerRulebase) {
		return false
	}
	if !util.StringsMatch(a.SecurityRulebase, b.SecurityRulebase) {
		return false
	}
	if !util.StringsMatch(a.SslDecryptionRulebase, b.SslDecryptionRulebase) {
		return false
	}
	if !util.StringsMatch(a.SdwanRulebase, b.SdwanRulebase) {
		return false
	}
	if !util.StringsMatch(a.TunnelInspectRulebase, b.TunnelInspectRulebase) {
		return false
	}
	if !util.StringsMatch(a.ApplicationOverrideRulebase, b.ApplicationOverrideRulebase) {
		return false
	}
	if !util.StringsMatch(a.DosRulebase, b.DosRulebase) {
		return false
	}
	if !util.StringsMatch(a.PbfRulebase, b.PbfRulebase) {
		return false
	}
	if !util.StringsMatch(a.QosRulebase, b.QosRulebase) {
		return false
	}
	if !util.StringsMatch(a.RuleHitCountReset, b.RuleHitCountReset) {
		return false
	}
	return true
}
func matchRoleDeviceWebui(a *RoleDeviceWebui, b *RoleDeviceWebui) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Acc, b.Acc) {
		return false
	}
	if !util.StringsMatch(a.Dashboard, b.Dashboard) {
		return false
	}
	if !matchRoleDeviceWebuiNetwork(a.Network, b.Network) {
		return false
	}
	if !matchRoleDeviceWebuiPrivacy(a.Privacy, b.Privacy) {
		return false
	}
	if !util.StringsMatch(a.Tasks, b.Tasks) {
		return false
	}
	if !matchRoleDeviceWebuiCommit(a.Commit, b.Commit) {
		return false
	}
	if !matchRoleDeviceWebuiDevice(a.Device, b.Device) {
		return false
	}
	if !matchRoleDeviceWebuiMonitor(a.Monitor, b.Monitor) {
		return false
	}
	if !matchRoleDeviceWebuiOperations(a.Operations, b.Operations) {
		return false
	}
	if !matchRoleDeviceWebuiSave(a.Save, b.Save) {
		return false
	}
	if !util.StringsMatch(a.Validate, b.Validate) {
		return false
	}
	if !matchRoleDeviceWebuiGlobal(a.Global, b.Global) {
		return false
	}
	if !matchRoleDeviceWebuiObjects(a.Objects, b.Objects) {
		return false
	}
	if !matchRoleDeviceWebuiPolicies(a.Policies, b.Policies) {
		return false
	}
	return true
}
func matchRoleDeviceXmlapi(a *RoleDeviceXmlapi, b *RoleDeviceXmlapi) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Op, b.Op) {
		return false
	}
	if !util.StringsMatch(a.Commit, b.Commit) {
		return false
	}
	if !util.StringsMatch(a.Import, b.Import) {
		return false
	}
	if !util.StringsMatch(a.Iot, b.Iot) {
		return false
	}
	if !util.StringsMatch(a.Log, b.Log) {
		return false
	}
	if !util.StringsMatch(a.Report, b.Report) {
		return false
	}
	if !util.StringsMatch(a.UserId, b.UserId) {
		return false
	}
	if !util.StringsMatch(a.Config, b.Config) {
		return false
	}
	if !util.StringsMatch(a.Export, b.Export) {
		return false
	}
	return true
}
func matchRoleDevice(a *RoleDevice, b *RoleDevice) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoleDeviceRestapi(a.Restapi, b.Restapi) {
		return false
	}
	if !matchRoleDeviceWebui(a.Webui, b.Webui) {
		return false
	}
	if !matchRoleDeviceXmlapi(a.Xmlapi, b.Xmlapi) {
		return false
	}
	if !util.StringsMatch(a.Cli, b.Cli) {
		return false
	}
	return true
}
func matchRoleVsysRestapiDevice(a *RoleVsysRestapiDevice, b *RoleVsysRestapiDevice) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.EmailServerProfiles, b.EmailServerProfiles) {
		return false
	}
	if !util.StringsMatch(a.HttpServerProfiles, b.HttpServerProfiles) {
		return false
	}
	if !util.StringsMatch(a.LdapServerProfiles, b.LdapServerProfiles) {
		return false
	}
	if !util.StringsMatch(a.LogInterfaceSetting, b.LogInterfaceSetting) {
		return false
	}
	if !util.StringsMatch(a.SnmpTrapServerProfiles, b.SnmpTrapServerProfiles) {
		return false
	}
	if !util.StringsMatch(a.SyslogServerProfiles, b.SyslogServerProfiles) {
		return false
	}
	if !util.StringsMatch(a.VirtualSystems, b.VirtualSystems) {
		return false
	}
	return true
}
func matchRoleVsysRestapiNetwork(a *RoleVsysRestapiNetwork, b *RoleVsysRestapiNetwork) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Zones, b.Zones) {
		return false
	}
	if !util.StringsMatch(a.SdwanInterfaceProfiles, b.SdwanInterfaceProfiles) {
		return false
	}
	if !util.StringsMatch(a.GlobalprotectClientlessAppGroups, b.GlobalprotectClientlessAppGroups) {
		return false
	}
	if !util.StringsMatch(a.GlobalprotectClientlessApps, b.GlobalprotectClientlessApps) {
		return false
	}
	if !util.StringsMatch(a.GlobalprotectGateways, b.GlobalprotectGateways) {
		return false
	}
	if !util.StringsMatch(a.GlobalprotectMdmServers, b.GlobalprotectMdmServers) {
		return false
	}
	if !util.StringsMatch(a.GlobalprotectPortals, b.GlobalprotectPortals) {
		return false
	}
	return true
}
func matchRoleVsysRestapiObjects(a *RoleVsysRestapiObjects, b *RoleVsysRestapiObjects) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DynamicUserGroups, b.DynamicUserGroups) {
		return false
	}
	if !util.StringsMatch(a.FileBlockingSecurityProfiles, b.FileBlockingSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.GlobalprotectHipProfiles, b.GlobalprotectHipProfiles) {
		return false
	}
	if !util.StringsMatch(a.LogForwardingProfiles, b.LogForwardingProfiles) {
		return false
	}
	if !util.StringsMatch(a.SdwanPathQualityProfiles, b.SdwanPathQualityProfiles) {
		return false
	}
	if !util.StringsMatch(a.AntivirusSecurityProfiles, b.AntivirusSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.CustomSpywareSignatures, b.CustomSpywareSignatures) {
		return false
	}
	if !util.StringsMatch(a.DosProtectionSecurityProfiles, b.DosProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.PacketBrokerProfiles, b.PacketBrokerProfiles) {
		return false
	}
	if !util.StringsMatch(a.SdwanTrafficDistributionProfiles, b.SdwanTrafficDistributionProfiles) {
		return false
	}
	if !util.StringsMatch(a.AntiSpywareSecurityProfiles, b.AntiSpywareSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.Applications, b.Applications) {
		return false
	}
	if !util.StringsMatch(a.Regions, b.Regions) {
		return false
	}
	if !util.StringsMatch(a.VulnerabilityProtectionSecurityProfiles, b.VulnerabilityProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.AuthenticationEnforcements, b.AuthenticationEnforcements) {
		return false
	}
	if !util.StringsMatch(a.DataFilteringSecurityProfiles, b.DataFilteringSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.ApplicationGroups, b.ApplicationGroups) {
		return false
	}
	if !util.StringsMatch(a.Devices, b.Devices) {
		return false
	}
	if !util.StringsMatch(a.Schedules, b.Schedules) {
		return false
	}
	if !util.StringsMatch(a.SecurityProfileGroups, b.SecurityProfileGroups) {
		return false
	}
	if !util.StringsMatch(a.Services, b.Services) {
		return false
	}
	if !util.StringsMatch(a.WildfireAnalysisSecurityProfiles, b.WildfireAnalysisSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.AddressGroups, b.AddressGroups) {
		return false
	}
	if !util.StringsMatch(a.ApplicationFilters, b.ApplicationFilters) {
		return false
	}
	if !util.StringsMatch(a.DecryptionProfiles, b.DecryptionProfiles) {
		return false
	}
	if !util.StringsMatch(a.GlobalprotectHipObjects, b.GlobalprotectHipObjects) {
		return false
	}
	if !util.StringsMatch(a.GtpProtectionSecurityProfiles, b.GtpProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.ServiceGroups, b.ServiceGroups) {
		return false
	}
	if !util.StringsMatch(a.Addresses, b.Addresses) {
		return false
	}
	if !util.StringsMatch(a.CustomUrlCategories, b.CustomUrlCategories) {
		return false
	}
	if !util.StringsMatch(a.UrlFilteringSecurityProfiles, b.UrlFilteringSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.ExternalDynamicLists, b.ExternalDynamicLists) {
		return false
	}
	if !util.StringsMatch(a.SctpProtectionSecurityProfiles, b.SctpProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(a.SdwanErrorCorrectionProfiles, b.SdwanErrorCorrectionProfiles) {
		return false
	}
	if !util.StringsMatch(a.Tags, b.Tags) {
		return false
	}
	if !util.StringsMatch(a.SdwanSaasQualityProfiles, b.SdwanSaasQualityProfiles) {
		return false
	}
	if !util.StringsMatch(a.CustomDataPatterns, b.CustomDataPatterns) {
		return false
	}
	if !util.StringsMatch(a.CustomVulnerabilitySignatures, b.CustomVulnerabilitySignatures) {
		return false
	}
	return true
}
func matchRoleVsysRestapiPolicies(a *RoleVsysRestapiPolicies, b *RoleVsysRestapiPolicies) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ApplicationOverrideRules, b.ApplicationOverrideRules) {
		return false
	}
	if !util.StringsMatch(a.AuthenticationRules, b.AuthenticationRules) {
		return false
	}
	if !util.StringsMatch(a.DecryptionRules, b.DecryptionRules) {
		return false
	}
	if !util.StringsMatch(a.NatRules, b.NatRules) {
		return false
	}
	if !util.StringsMatch(a.PolicyBasedForwardingRules, b.PolicyBasedForwardingRules) {
		return false
	}
	if !util.StringsMatch(a.QosRules, b.QosRules) {
		return false
	}
	if !util.StringsMatch(a.SecurityRules, b.SecurityRules) {
		return false
	}
	if !util.StringsMatch(a.TunnelInspectionRules, b.TunnelInspectionRules) {
		return false
	}
	if !util.StringsMatch(a.DosRules, b.DosRules) {
		return false
	}
	if !util.StringsMatch(a.NetworkPacketBrokerRules, b.NetworkPacketBrokerRules) {
		return false
	}
	if !util.StringsMatch(a.SdwanRules, b.SdwanRules) {
		return false
	}
	return true
}
func matchRoleVsysRestapiSystem(a *RoleVsysRestapiSystem, b *RoleVsysRestapiSystem) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Configuration, b.Configuration) {
		return false
	}
	return true
}
func matchRoleVsysRestapi(a *RoleVsysRestapi, b *RoleVsysRestapi) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoleVsysRestapiDevice(a.Device, b.Device) {
		return false
	}
	if !matchRoleVsysRestapiNetwork(a.Network, b.Network) {
		return false
	}
	if !matchRoleVsysRestapiObjects(a.Objects, b.Objects) {
		return false
	}
	if !matchRoleVsysRestapiPolicies(a.Policies, b.Policies) {
		return false
	}
	if !matchRoleVsysRestapiSystem(a.System, b.System) {
		return false
	}
	return true
}
func matchRoleVsysWebuiDeviceLocalUserDatabase(a *RoleVsysWebuiDeviceLocalUserDatabase, b *RoleVsysWebuiDeviceLocalUserDatabase) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.UserGroups, b.UserGroups) {
		return false
	}
	if !util.StringsMatch(a.Users, b.Users) {
		return false
	}
	return true
}
func matchRoleVsysWebuiDeviceCertificateManagement(a *RoleVsysWebuiDeviceCertificateManagement, b *RoleVsysWebuiDeviceCertificateManagement) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.SslTlsServiceProfile, b.SslTlsServiceProfile) {
		return false
	}
	if !util.StringsMatch(a.CertificateProfile, b.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(a.Certificates, b.Certificates) {
		return false
	}
	if !util.StringsMatch(a.OcspResponder, b.OcspResponder) {
		return false
	}
	if !util.StringsMatch(a.Scep, b.Scep) {
		return false
	}
	if !util.StringsMatch(a.SshServiceProfile, b.SshServiceProfile) {
		return false
	}
	if !util.StringsMatch(a.SslDecryptionExclusion, b.SslDecryptionExclusion) {
		return false
	}
	return true
}
func matchRoleVsysWebuiDeviceLogSettings(a *RoleVsysWebuiDeviceLogSettings, b *RoleVsysWebuiDeviceLogSettings) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.UserId, b.UserId) {
		return false
	}
	if !util.StringsMatch(a.Config, b.Config) {
		return false
	}
	if !util.StringsMatch(a.Correlation, b.Correlation) {
		return false
	}
	if !util.StringsMatch(a.Globalprotect, b.Globalprotect) {
		return false
	}
	if !util.StringsMatch(a.Hipmatch, b.Hipmatch) {
		return false
	}
	if !util.StringsMatch(a.Iptag, b.Iptag) {
		return false
	}
	if !util.StringsMatch(a.System, b.System) {
		return false
	}
	return true
}
func matchRoleVsysWebuiDevicePolicyRecommendations(a *RoleVsysWebuiDevicePolicyRecommendations, b *RoleVsysWebuiDevicePolicyRecommendations) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Iot, b.Iot) {
		return false
	}
	if !util.StringsMatch(a.Saas, b.Saas) {
		return false
	}
	return true
}
func matchRoleVsysWebuiDeviceSetup(a *RoleVsysWebuiDeviceSetup, b *RoleVsysWebuiDeviceSetup) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Operations, b.Operations) {
		return false
	}
	if !util.StringsMatch(a.Session, b.Session) {
		return false
	}
	if !util.StringsMatch(a.Telemetry, b.Telemetry) {
		return false
	}
	if !util.StringsMatch(a.Wildfire, b.Wildfire) {
		return false
	}
	if !util.StringsMatch(a.Interfaces, b.Interfaces) {
		return false
	}
	if !util.StringsMatch(a.Hsm, b.Hsm) {
		return false
	}
	if !util.StringsMatch(a.Management, b.Management) {
		return false
	}
	if !util.StringsMatch(a.Services, b.Services) {
		return false
	}
	if !util.StringsMatch(a.ContentId, b.ContentId) {
		return false
	}
	return true
}
func matchRoleVsysWebuiDeviceServerProfile(a *RoleVsysWebuiDeviceServerProfile, b *RoleVsysWebuiDeviceServerProfile) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Http, b.Http) {
		return false
	}
	if !util.StringsMatch(a.Kerberos, b.Kerberos) {
		return false
	}
	if !util.StringsMatch(a.Netflow, b.Netflow) {
		return false
	}
	if !util.StringsMatch(a.Radius, b.Radius) {
		return false
	}
	if !util.StringsMatch(a.SamlIdp, b.SamlIdp) {
		return false
	}
	if !util.StringsMatch(a.SnmpTrap, b.SnmpTrap) {
		return false
	}
	if !util.StringsMatch(a.Dns, b.Dns) {
		return false
	}
	if !util.StringsMatch(a.Ldap, b.Ldap) {
		return false
	}
	if !util.StringsMatch(a.Mfa, b.Mfa) {
		return false
	}
	if !util.StringsMatch(a.Scp, b.Scp) {
		return false
	}
	if !util.StringsMatch(a.Syslog, b.Syslog) {
		return false
	}
	if !util.StringsMatch(a.Tacplus, b.Tacplus) {
		return false
	}
	if !util.StringsMatch(a.Email, b.Email) {
		return false
	}
	return true
}
func matchRoleVsysWebuiDevice(a *RoleVsysWebuiDevice, b *RoleVsysWebuiDevice) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DataRedistribution, b.DataRedistribution) {
		return false
	}
	if !util.StringsMatch(a.UserIdentification, b.UserIdentification) {
		return false
	}
	if !util.StringsMatch(a.DhcpSyslogServer, b.DhcpSyslogServer) {
		return false
	}
	if !util.StringsMatch(a.Administrators, b.Administrators) {
		return false
	}
	if !util.StringsMatch(a.AuthenticationProfile, b.AuthenticationProfile) {
		return false
	}
	if !matchRoleVsysWebuiDeviceServerProfile(a.ServerProfile, b.ServerProfile) {
		return false
	}
	if !util.StringsMatch(a.AuthenticationSequence, b.AuthenticationSequence) {
		return false
	}
	if !matchRoleVsysWebuiDeviceLocalUserDatabase(a.LocalUserDatabase, b.LocalUserDatabase) {
		return false
	}
	if !util.StringsMatch(a.Troubleshooting, b.Troubleshooting) {
		return false
	}
	if !util.StringsMatch(a.VmInfoSource, b.VmInfoSource) {
		return false
	}
	if !matchRoleVsysWebuiDevicePolicyRecommendations(a.PolicyRecommendations, b.PolicyRecommendations) {
		return false
	}
	if !matchRoleVsysWebuiDeviceSetup(a.Setup, b.Setup) {
		return false
	}
	if !util.StringsMatch(a.BlockPages, b.BlockPages) {
		return false
	}
	if !matchRoleVsysWebuiDeviceCertificateManagement(a.CertificateManagement, b.CertificateManagement) {
		return false
	}
	if !util.StringsMatch(a.DeviceQuarantine, b.DeviceQuarantine) {
		return false
	}
	if !matchRoleVsysWebuiDeviceLogSettings(a.LogSettings, b.LogSettings) {
		return false
	}
	return true
}
func matchRoleVsysWebuiMonitorPdfReports(a *RoleVsysWebuiMonitorPdfReports, b *RoleVsysWebuiMonitorPdfReports) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ReportGroups, b.ReportGroups) {
		return false
	}
	if !util.StringsMatch(a.SaasApplicationUsageReport, b.SaasApplicationUsageReport) {
		return false
	}
	if !util.StringsMatch(a.UserActivityReport, b.UserActivityReport) {
		return false
	}
	if !util.StringsMatch(a.EmailScheduler, b.EmailScheduler) {
		return false
	}
	if !util.StringsMatch(a.ManagePdfSummary, b.ManagePdfSummary) {
		return false
	}
	if !util.StringsMatch(a.PdfSummaryReports, b.PdfSummaryReports) {
		return false
	}
	return true
}
func matchRoleVsysWebuiMonitorAutomatedCorrelationEngine(a *RoleVsysWebuiMonitorAutomatedCorrelationEngine, b *RoleVsysWebuiMonitorAutomatedCorrelationEngine) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.CorrelatedEvents, b.CorrelatedEvents) {
		return false
	}
	if !util.StringsMatch(a.CorrelationObjects, b.CorrelationObjects) {
		return false
	}
	return true
}
func matchRoleVsysWebuiMonitorCustomReports(a *RoleVsysWebuiMonitorCustomReports, b *RoleVsysWebuiMonitorCustomReports) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ApplicationStatistics, b.ApplicationStatistics) {
		return false
	}
	if !util.StringsMatch(a.Auth, b.Auth) {
		return false
	}
	if !util.StringsMatch(a.TunnelLog, b.TunnelLog) {
		return false
	}
	if !util.StringsMatch(a.TunnelSummary, b.TunnelSummary) {
		return false
	}
	if !util.StringsMatch(a.GtpLog, b.GtpLog) {
		return false
	}
	if !util.StringsMatch(a.GtpSummary, b.GtpSummary) {
		return false
	}
	if !util.StringsMatch(a.SctpLog, b.SctpLog) {
		return false
	}
	if !util.StringsMatch(a.SctpSummary, b.SctpSummary) {
		return false
	}
	if !util.StringsMatch(a.TrafficLog, b.TrafficLog) {
		return false
	}
	if !util.StringsMatch(a.Userid, b.Userid) {
		return false
	}
	if !util.StringsMatch(a.TrafficSummary, b.TrafficSummary) {
		return false
	}
	if !util.StringsMatch(a.UrlLog, b.UrlLog) {
		return false
	}
	if !util.StringsMatch(a.DataFilteringLog, b.DataFilteringLog) {
		return false
	}
	if !util.StringsMatch(a.DecryptionSummary, b.DecryptionSummary) {
		return false
	}
	if !util.StringsMatch(a.Globalprotect, b.Globalprotect) {
		return false
	}
	if !util.StringsMatch(a.Hipmatch, b.Hipmatch) {
		return false
	}
	if !util.StringsMatch(a.Iptag, b.Iptag) {
		return false
	}
	if !util.StringsMatch(a.ThreatSummary, b.ThreatSummary) {
		return false
	}
	if !util.StringsMatch(a.UrlSummary, b.UrlSummary) {
		return false
	}
	if !util.StringsMatch(a.WildfireLog, b.WildfireLog) {
		return false
	}
	if !util.StringsMatch(a.DecryptionLog, b.DecryptionLog) {
		return false
	}
	if !util.StringsMatch(a.ThreatLog, b.ThreatLog) {
		return false
	}
	return true
}
func matchRoleVsysWebuiMonitorLogs(a *RoleVsysWebuiMonitorLogs, b *RoleVsysWebuiMonitorLogs) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Threat, b.Threat) {
		return false
	}
	if !util.StringsMatch(a.Wildfire, b.Wildfire) {
		return false
	}
	if !util.StringsMatch(a.DataFiltering, b.DataFiltering) {
		return false
	}
	if !util.StringsMatch(a.Globalprotect, b.Globalprotect) {
		return false
	}
	if !util.StringsMatch(a.Hipmatch, b.Hipmatch) {
		return false
	}
	if !util.StringsMatch(a.Url, b.Url) {
		return false
	}
	if !util.StringsMatch(a.Userid, b.Userid) {
		return false
	}
	if !util.StringsMatch(a.Authentication, b.Authentication) {
		return false
	}
	if !util.StringsMatch(a.Decryption, b.Decryption) {
		return false
	}
	if !util.StringsMatch(a.Gtp, b.Gtp) {
		return false
	}
	if !util.StringsMatch(a.Iptag, b.Iptag) {
		return false
	}
	if !util.StringsMatch(a.Traffic, b.Traffic) {
		return false
	}
	if !util.StringsMatch(a.Tunnel, b.Tunnel) {
		return false
	}
	if !util.StringsMatch(a.Sctp, b.Sctp) {
		return false
	}
	return true
}
func matchRoleVsysWebuiMonitor(a *RoleVsysWebuiMonitor, b *RoleVsysWebuiMonitor) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoleVsysWebuiMonitorAutomatedCorrelationEngine(a.AutomatedCorrelationEngine, b.AutomatedCorrelationEngine) {
		return false
	}
	if !util.StringsMatch(a.BlockIpList, b.BlockIpList) {
		return false
	}
	if !util.StringsMatch(a.ExternalLogs, b.ExternalLogs) {
		return false
	}
	if !matchRoleVsysWebuiMonitorPdfReports(a.PdfReports, b.PdfReports) {
		return false
	}
	if !util.StringsMatch(a.AppScope, b.AppScope) {
		return false
	}
	if !matchRoleVsysWebuiMonitorCustomReports(a.CustomReports, b.CustomReports) {
		return false
	}
	if !matchRoleVsysWebuiMonitorLogs(a.Logs, b.Logs) {
		return false
	}
	if !util.StringsMatch(a.SessionBrowser, b.SessionBrowser) {
		return false
	}
	if !util.StringsMatch(a.ViewCustomReports, b.ViewCustomReports) {
		return false
	}
	return true
}
func matchRoleVsysWebuiNetworkGlobalProtect(a *RoleVsysWebuiNetworkGlobalProtect, b *RoleVsysWebuiNetworkGlobalProtect) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ClientlessApps, b.ClientlessApps) {
		return false
	}
	if !util.StringsMatch(a.Gateways, b.Gateways) {
		return false
	}
	if !util.StringsMatch(a.Mdm, b.Mdm) {
		return false
	}
	if !util.StringsMatch(a.Portals, b.Portals) {
		return false
	}
	if !util.StringsMatch(a.ClientlessAppGroups, b.ClientlessAppGroups) {
		return false
	}
	return true
}
func matchRoleVsysWebuiNetwork(a *RoleVsysWebuiNetwork, b *RoleVsysWebuiNetwork) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoleVsysWebuiNetworkGlobalProtect(a.GlobalProtect, b.GlobalProtect) {
		return false
	}
	if !util.StringsMatch(a.SdwanInterfaceProfile, b.SdwanInterfaceProfile) {
		return false
	}
	if !util.StringsMatch(a.Zones, b.Zones) {
		return false
	}
	return true
}
func matchRoleVsysWebuiOperations(a *RoleVsysWebuiOperations, b *RoleVsysWebuiOperations) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DownloadCoreFiles, b.DownloadCoreFiles) {
		return false
	}
	if !util.StringsMatch(a.DownloadPcapFiles, b.DownloadPcapFiles) {
		return false
	}
	if !util.StringsMatch(a.GenerateStatsDumpFile, b.GenerateStatsDumpFile) {
		return false
	}
	if !util.StringsMatch(a.GenerateTechSupportFile, b.GenerateTechSupportFile) {
		return false
	}
	if !util.StringsMatch(a.Reboot, b.Reboot) {
		return false
	}
	return true
}
func matchRoleVsysWebuiPolicies(a *RoleVsysWebuiPolicies, b *RoleVsysWebuiPolicies) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ApplicationOverrideRulebase, b.ApplicationOverrideRulebase) {
		return false
	}
	if !util.StringsMatch(a.AuthenticationRulebase, b.AuthenticationRulebase) {
		return false
	}
	if !util.StringsMatch(a.DosRulebase, b.DosRulebase) {
		return false
	}
	if !util.StringsMatch(a.PbfRulebase, b.PbfRulebase) {
		return false
	}
	if !util.StringsMatch(a.TunnelInspectRulebase, b.TunnelInspectRulebase) {
		return false
	}
	if !util.StringsMatch(a.SecurityRulebase, b.SecurityRulebase) {
		return false
	}
	if !util.StringsMatch(a.SslDecryptionRulebase, b.SslDecryptionRulebase) {
		return false
	}
	if !util.StringsMatch(a.NatRulebase, b.NatRulebase) {
		return false
	}
	if !util.StringsMatch(a.NetworkPacketBrokerRulebase, b.NetworkPacketBrokerRulebase) {
		return false
	}
	if !util.StringsMatch(a.QosRulebase, b.QosRulebase) {
		return false
	}
	if !util.StringsMatch(a.RuleHitCountReset, b.RuleHitCountReset) {
		return false
	}
	if !util.StringsMatch(a.SdwanRulebase, b.SdwanRulebase) {
		return false
	}
	return true
}
func matchRoleVsysWebuiSave(a *RoleVsysWebuiSave, b *RoleVsysWebuiSave) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.SaveForOtherAdmins, b.SaveForOtherAdmins) {
		return false
	}
	if !util.StringsMatch(a.ObjectLevelChanges, b.ObjectLevelChanges) {
		return false
	}
	if !util.StringsMatch(a.PartialSave, b.PartialSave) {
		return false
	}
	return true
}
func matchRoleVsysWebuiCommit(a *RoleVsysWebuiCommit, b *RoleVsysWebuiCommit) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.CommitForOtherAdmins, b.CommitForOtherAdmins) {
		return false
	}
	if !util.StringsMatch(a.VirtualSystems, b.VirtualSystems) {
		return false
	}
	return true
}
func matchRoleVsysWebuiObjectsCustomObjects(a *RoleVsysWebuiObjectsCustomObjects, b *RoleVsysWebuiObjectsCustomObjects) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Spyware, b.Spyware) {
		return false
	}
	if !util.StringsMatch(a.UrlCategory, b.UrlCategory) {
		return false
	}
	if !util.StringsMatch(a.Vulnerability, b.Vulnerability) {
		return false
	}
	if !util.StringsMatch(a.DataPatterns, b.DataPatterns) {
		return false
	}
	return true
}
func matchRoleVsysWebuiObjectsGlobalProtect(a *RoleVsysWebuiObjectsGlobalProtect, b *RoleVsysWebuiObjectsGlobalProtect) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.HipObjects, b.HipObjects) {
		return false
	}
	if !util.StringsMatch(a.HipProfiles, b.HipProfiles) {
		return false
	}
	return true
}
func matchRoleVsysWebuiObjectsDecryption(a *RoleVsysWebuiObjectsDecryption, b *RoleVsysWebuiObjectsDecryption) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DecryptionProfile, b.DecryptionProfile) {
		return false
	}
	return true
}
func matchRoleVsysWebuiObjectsSecurityProfiles(a *RoleVsysWebuiObjectsSecurityProfiles, b *RoleVsysWebuiObjectsSecurityProfiles) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DosProtection, b.DosProtection) {
		return false
	}
	if !util.StringsMatch(a.UrlFiltering, b.UrlFiltering) {
		return false
	}
	if !util.StringsMatch(a.AntiSpyware, b.AntiSpyware) {
		return false
	}
	if !util.StringsMatch(a.DataFiltering, b.DataFiltering) {
		return false
	}
	if !util.StringsMatch(a.GtpProtection, b.GtpProtection) {
		return false
	}
	if !util.StringsMatch(a.SctpProtection, b.SctpProtection) {
		return false
	}
	if !util.StringsMatch(a.VulnerabilityProtection, b.VulnerabilityProtection) {
		return false
	}
	if !util.StringsMatch(a.WildfireAnalysis, b.WildfireAnalysis) {
		return false
	}
	if !util.StringsMatch(a.Antivirus, b.Antivirus) {
		return false
	}
	if !util.StringsMatch(a.FileBlocking, b.FileBlocking) {
		return false
	}
	return true
}
func matchRoleVsysWebuiObjectsSdwan(a *RoleVsysWebuiObjectsSdwan, b *RoleVsysWebuiObjectsSdwan) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.SdwanDistProfile, b.SdwanDistProfile) {
		return false
	}
	if !util.StringsMatch(a.SdwanErrorCorrectionProfile, b.SdwanErrorCorrectionProfile) {
		return false
	}
	if !util.StringsMatch(a.SdwanProfile, b.SdwanProfile) {
		return false
	}
	if !util.StringsMatch(a.SdwanSaasQualityProfile, b.SdwanSaasQualityProfile) {
		return false
	}
	return true
}
func matchRoleVsysWebuiObjects(a *RoleVsysWebuiObjects, b *RoleVsysWebuiObjects) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Applications, b.Applications) {
		return false
	}
	if !util.StringsMatch(a.Regions, b.Regions) {
		return false
	}
	if !matchRoleVsysWebuiObjectsSdwan(a.Sdwan, b.Sdwan) {
		return false
	}
	if !util.StringsMatch(a.SecurityProfileGroups, b.SecurityProfileGroups) {
		return false
	}
	if !util.StringsMatch(a.LogForwarding, b.LogForwarding) {
		return false
	}
	if !util.StringsMatch(a.PacketBrokerProfile, b.PacketBrokerProfile) {
		return false
	}
	if !util.StringsMatch(a.ApplicationFilters, b.ApplicationFilters) {
		return false
	}
	if !util.StringsMatch(a.Authentication, b.Authentication) {
		return false
	}
	if !matchRoleVsysWebuiObjectsCustomObjects(a.CustomObjects, b.CustomObjects) {
		return false
	}
	if !util.StringsMatch(a.DynamicBlockLists, b.DynamicBlockLists) {
		return false
	}
	if !matchRoleVsysWebuiObjectsGlobalProtect(a.GlobalProtect, b.GlobalProtect) {
		return false
	}
	if !util.StringsMatch(a.Tags, b.Tags) {
		return false
	}
	if !util.StringsMatch(a.AddressGroups, b.AddressGroups) {
		return false
	}
	if !util.StringsMatch(a.Addresses, b.Addresses) {
		return false
	}
	if !util.StringsMatch(a.ApplicationGroups, b.ApplicationGroups) {
		return false
	}
	if !util.StringsMatch(a.ServiceGroups, b.ServiceGroups) {
		return false
	}
	if !util.StringsMatch(a.Services, b.Services) {
		return false
	}
	if !matchRoleVsysWebuiObjectsDecryption(a.Decryption, b.Decryption) {
		return false
	}
	if !util.StringsMatch(a.Devices, b.Devices) {
		return false
	}
	if !util.StringsMatch(a.DynamicUserGroups, b.DynamicUserGroups) {
		return false
	}
	if !util.StringsMatch(a.Schedules, b.Schedules) {
		return false
	}
	if !matchRoleVsysWebuiObjectsSecurityProfiles(a.SecurityProfiles, b.SecurityProfiles) {
		return false
	}
	return true
}
func matchRoleVsysWebuiPrivacy(a *RoleVsysWebuiPrivacy, b *RoleVsysWebuiPrivacy) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.ShowUserNamesInLogsAndReports, b.ShowUserNamesInLogsAndReports) {
		return false
	}
	if !util.StringsMatch(a.ViewPcapFiles, b.ViewPcapFiles) {
		return false
	}
	if !util.StringsMatch(a.ShowFullIpAddresses, b.ShowFullIpAddresses) {
		return false
	}
	return true
}
func matchRoleVsysWebui(a *RoleVsysWebui, b *RoleVsysWebui) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoleVsysWebuiSave(a.Save, b.Save) {
		return false
	}
	if !util.StringsMatch(a.Acc, b.Acc) {
		return false
	}
	if !util.StringsMatch(a.Dashboard, b.Dashboard) {
		return false
	}
	if !matchRoleVsysWebuiDevice(a.Device, b.Device) {
		return false
	}
	if !matchRoleVsysWebuiMonitor(a.Monitor, b.Monitor) {
		return false
	}
	if !matchRoleVsysWebuiNetwork(a.Network, b.Network) {
		return false
	}
	if !matchRoleVsysWebuiOperations(a.Operations, b.Operations) {
		return false
	}
	if !matchRoleVsysWebuiPolicies(a.Policies, b.Policies) {
		return false
	}
	if !util.StringsMatch(a.Tasks, b.Tasks) {
		return false
	}
	if !util.StringsMatch(a.Validate, b.Validate) {
		return false
	}
	if !matchRoleVsysWebuiCommit(a.Commit, b.Commit) {
		return false
	}
	if !matchRoleVsysWebuiObjects(a.Objects, b.Objects) {
		return false
	}
	if !matchRoleVsysWebuiPrivacy(a.Privacy, b.Privacy) {
		return false
	}
	return true
}
func matchRoleVsysXmlapi(a *RoleVsysXmlapi, b *RoleVsysXmlapi) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Config, b.Config) {
		return false
	}
	if !util.StringsMatch(a.Import, b.Import) {
		return false
	}
	if !util.StringsMatch(a.Iot, b.Iot) {
		return false
	}
	if !util.StringsMatch(a.Log, b.Log) {
		return false
	}
	if !util.StringsMatch(a.Op, b.Op) {
		return false
	}
	if !util.StringsMatch(a.Commit, b.Commit) {
		return false
	}
	if !util.StringsMatch(a.Report, b.Report) {
		return false
	}
	if !util.StringsMatch(a.UserId, b.UserId) {
		return false
	}
	if !util.StringsMatch(a.Export, b.Export) {
		return false
	}
	return true
}
func matchRoleVsys(a *RoleVsys, b *RoleVsys) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Cli, b.Cli) {
		return false
	}
	if !matchRoleVsysRestapi(a.Restapi, b.Restapi) {
		return false
	}
	if !matchRoleVsysWebui(a.Webui, b.Webui) {
		return false
	}
	if !matchRoleVsysXmlapi(a.Xmlapi, b.Xmlapi) {
		return false
	}
	return true
}
func matchRole(a *Role, b *Role) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoleDevice(a.Device, b.Device) {
		return false
	}
	if !matchRoleVsys(a.Vsys, b.Vsys) {
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
