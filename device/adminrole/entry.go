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
	suffix = []string{"admin-role", "$name"}
)

type Entry struct {
	Name        string
	Description *string
	Role        *Role
	Misc        []generic.Xml
}
type Role struct {
	Device *RoleDevice
	Vsys   *RoleVsys
	Misc   []generic.Xml
}
type RoleDevice struct {
	Cli     *string
	Restapi *RoleDeviceRestapi
	Webui   *RoleDeviceWebui
	Xmlapi  *RoleDeviceXmlapi
	Misc    []generic.Xml
}
type RoleDeviceRestapi struct {
	Device   *RoleDeviceRestapiDevice
	Network  *RoleDeviceRestapiNetwork
	Objects  *RoleDeviceRestapiObjects
	Policies *RoleDeviceRestapiPolicies
	System   *RoleDeviceRestapiSystem
	Misc     []generic.Xml
}
type RoleDeviceRestapiDevice struct {
	EmailServerProfiles    *string
	HttpServerProfiles     *string
	LdapServerProfiles     *string
	LogInterfaceSetting    *string
	SnmpTrapServerProfiles *string
	SyslogServerProfiles   *string
	VirtualSystems         *string
	Misc                   []generic.Xml
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
	Misc                                    []generic.Xml
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
	Misc                                    []generic.Xml
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
	Misc                       []generic.Xml
}
type RoleDeviceRestapiSystem struct {
	Configuration *string
	Misc          []generic.Xml
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
	Misc       []generic.Xml
}
type RoleDeviceWebuiCommit struct {
	CommitForOtherAdmins *string
	Device               *string
	ObjectLevelChanges   *string
	Misc                 []generic.Xml
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
	Misc                   []generic.Xml
}
type RoleDeviceWebuiDeviceCertificateManagement struct {
	CertificateProfile     *string
	Certificates           *string
	OcspResponder          *string
	Scep                   *string
	SshServiceProfile      *string
	SslDecryptionExclusion *string
	SslTlsServiceProfile   *string
	Misc                   []generic.Xml
}
type RoleDeviceWebuiDeviceLocalUserDatabase struct {
	UserGroups *string
	Users      *string
	Misc       []generic.Xml
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
	Misc          []generic.Xml
}
type RoleDeviceWebuiDevicePolicyRecommendations struct {
	Iot  *string
	Saas *string
	Misc []generic.Xml
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
	Misc     []generic.Xml
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
	Misc       []generic.Xml
}
type RoleDeviceWebuiGlobal struct {
	SystemAlarms *string
	Misc         []generic.Xml
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
	Misc                       []generic.Xml
}
type RoleDeviceWebuiMonitorAutomatedCorrelationEngine struct {
	CorrelatedEvents   *string
	CorrelationObjects *string
	Misc               []generic.Xml
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
	Misc                  []generic.Xml
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
	Misc           []generic.Xml
}
type RoleDeviceWebuiMonitorPdfReports struct {
	EmailScheduler             *string
	ManagePdfSummary           *string
	PdfSummaryReports          *string
	ReportGroups               *string
	SaasApplicationUsageReport *string
	UserActivityReport         *string
	Misc                       []generic.Xml
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
	Misc                  []generic.Xml
}
type RoleDeviceWebuiNetworkGlobalProtect struct {
	ClientlessAppGroups *string
	ClientlessApps      *string
	Gateways            *string
	Mdm                 *string
	Portals             *string
	Misc                []generic.Xml
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
	Misc             []generic.Xml
}
type RoleDeviceWebuiNetworkRouting struct {
	LogicalRouters  *string
	RoutingProfiles *RoleDeviceWebuiNetworkRoutingRoutingProfiles
	Misc            []generic.Xml
}
type RoleDeviceWebuiNetworkRoutingRoutingProfiles struct {
	Bfd       *string
	Bgp       *string
	Filters   *string
	Multicast *string
	Ospf      *string
	Ospfv3    *string
	Ripv2     *string
	Misc      []generic.Xml
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
	Misc                  []generic.Xml
}
type RoleDeviceWebuiObjectsCustomObjects struct {
	DataPatterns  *string
	Spyware       *string
	UrlCategory   *string
	Vulnerability *string
	Misc          []generic.Xml
}
type RoleDeviceWebuiObjectsDecryption struct {
	DecryptionProfile *string
	Misc              []generic.Xml
}
type RoleDeviceWebuiObjectsGlobalProtect struct {
	HipObjects  *string
	HipProfiles *string
	Misc        []generic.Xml
}
type RoleDeviceWebuiObjectsSdwan struct {
	SdwanDistProfile            *string
	SdwanErrorCorrectionProfile *string
	SdwanProfile                *string
	SdwanSaasQualityProfile     *string
	Misc                        []generic.Xml
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
	Misc                    []generic.Xml
}
type RoleDeviceWebuiOperations struct {
	DownloadCoreFiles       *string
	DownloadPcapFiles       *string
	GenerateStatsDumpFile   *string
	GenerateTechSupportFile *string
	Reboot                  *string
	Misc                    []generic.Xml
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
	Misc                        []generic.Xml
}
type RoleDeviceWebuiPrivacy struct {
	ShowFullIpAddresses           *string
	ShowUserNamesInLogsAndReports *string
	ViewPcapFiles                 *string
	Misc                          []generic.Xml
}
type RoleDeviceWebuiSave struct {
	ObjectLevelChanges *string
	PartialSave        *string
	SaveForOtherAdmins *string
	Misc               []generic.Xml
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
	Misc   []generic.Xml
}
type RoleVsys struct {
	Cli     *string
	Restapi *RoleVsysRestapi
	Webui   *RoleVsysWebui
	Xmlapi  *RoleVsysXmlapi
	Misc    []generic.Xml
}
type RoleVsysRestapi struct {
	Device   *RoleVsysRestapiDevice
	Network  *RoleVsysRestapiNetwork
	Objects  *RoleVsysRestapiObjects
	Policies *RoleVsysRestapiPolicies
	System   *RoleVsysRestapiSystem
	Misc     []generic.Xml
}
type RoleVsysRestapiDevice struct {
	EmailServerProfiles    *string
	HttpServerProfiles     *string
	LdapServerProfiles     *string
	LogInterfaceSetting    *string
	SnmpTrapServerProfiles *string
	SyslogServerProfiles   *string
	VirtualSystems         *string
	Misc                   []generic.Xml
}
type RoleVsysRestapiNetwork struct {
	GlobalprotectClientlessAppGroups *string
	GlobalprotectClientlessApps      *string
	GlobalprotectGateways            *string
	GlobalprotectMdmServers          *string
	GlobalprotectPortals             *string
	SdwanInterfaceProfiles           *string
	Zones                            *string
	Misc                             []generic.Xml
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
	Misc                                    []generic.Xml
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
	Misc                       []generic.Xml
}
type RoleVsysRestapiSystem struct {
	Configuration *string
	Misc          []generic.Xml
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
	Misc       []generic.Xml
}
type RoleVsysWebuiCommit struct {
	CommitForOtherAdmins *string
	VirtualSystems       *string
	Misc                 []generic.Xml
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
	Misc                   []generic.Xml
}
type RoleVsysWebuiDeviceCertificateManagement struct {
	CertificateProfile     *string
	Certificates           *string
	OcspResponder          *string
	Scep                   *string
	SshServiceProfile      *string
	SslDecryptionExclusion *string
	SslTlsServiceProfile   *string
	Misc                   []generic.Xml
}
type RoleVsysWebuiDeviceLocalUserDatabase struct {
	UserGroups *string
	Users      *string
	Misc       []generic.Xml
}
type RoleVsysWebuiDeviceLogSettings struct {
	Config        *string
	Correlation   *string
	Globalprotect *string
	Hipmatch      *string
	Iptag         *string
	System        *string
	UserId        *string
	Misc          []generic.Xml
}
type RoleVsysWebuiDevicePolicyRecommendations struct {
	Iot  *string
	Saas *string
	Misc []generic.Xml
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
	Misc     []generic.Xml
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
	Misc       []generic.Xml
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
	Misc                       []generic.Xml
}
type RoleVsysWebuiMonitorAutomatedCorrelationEngine struct {
	CorrelatedEvents   *string
	CorrelationObjects *string
	Misc               []generic.Xml
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
	Misc                  []generic.Xml
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
	Misc           []generic.Xml
}
type RoleVsysWebuiMonitorPdfReports struct {
	EmailScheduler             *string
	ManagePdfSummary           *string
	PdfSummaryReports          *string
	ReportGroups               *string
	SaasApplicationUsageReport *string
	UserActivityReport         *string
	Misc                       []generic.Xml
}
type RoleVsysWebuiNetwork struct {
	GlobalProtect         *RoleVsysWebuiNetworkGlobalProtect
	SdwanInterfaceProfile *string
	Zones                 *string
	Misc                  []generic.Xml
}
type RoleVsysWebuiNetworkGlobalProtect struct {
	ClientlessAppGroups *string
	ClientlessApps      *string
	Gateways            *string
	Mdm                 *string
	Portals             *string
	Misc                []generic.Xml
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
	Misc                  []generic.Xml
}
type RoleVsysWebuiObjectsCustomObjects struct {
	DataPatterns  *string
	Spyware       *string
	UrlCategory   *string
	Vulnerability *string
	Misc          []generic.Xml
}
type RoleVsysWebuiObjectsDecryption struct {
	DecryptionProfile *string
	Misc              []generic.Xml
}
type RoleVsysWebuiObjectsGlobalProtect struct {
	HipObjects  *string
	HipProfiles *string
	Misc        []generic.Xml
}
type RoleVsysWebuiObjectsSdwan struct {
	SdwanDistProfile            *string
	SdwanErrorCorrectionProfile *string
	SdwanProfile                *string
	SdwanSaasQualityProfile     *string
	Misc                        []generic.Xml
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
	Misc                    []generic.Xml
}
type RoleVsysWebuiOperations struct {
	DownloadCoreFiles       *string
	DownloadPcapFiles       *string
	GenerateStatsDumpFile   *string
	GenerateTechSupportFile *string
	Reboot                  *string
	Misc                    []generic.Xml
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
	Misc                        []generic.Xml
}
type RoleVsysWebuiPrivacy struct {
	ShowFullIpAddresses           *string
	ShowUserNamesInLogsAndReports *string
	ViewPcapFiles                 *string
	Misc                          []generic.Xml
}
type RoleVsysWebuiSave struct {
	ObjectLevelChanges *string
	PartialSave        *string
	SaveForOtherAdmins *string
	Misc               []generic.Xml
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
	Misc   []generic.Xml
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXmlContainer_11_0_2 struct {
	Answer []entryXml_11_0_2 `xml:"entry"`
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
func (o *entryXmlContainer_11_0_2) Normalize() ([]*Entry, error) {
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
func specifyEntry_11_0_2(source *Entry) (any, error) {
	var obj entryXml_11_0_2
	obj.MarshalFromObject(*source)
	return obj, nil
}

type entryXml struct {
	XMLName     xml.Name      `xml:"entry"`
	Name        string        `xml:"name,attr"`
	Description *string       `xml:"description,omitempty"`
	Role        *roleXml      `xml:"role,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type roleXml struct {
	Device *roleDeviceXml `xml:"device,omitempty"`
	Vsys   *roleVsysXml   `xml:"vsys,omitempty"`
	Misc   []generic.Xml  `xml:",any"`
}
type roleDeviceXml struct {
	Cli     *string               `xml:"cli,omitempty"`
	Restapi *roleDeviceRestapiXml `xml:"restapi,omitempty"`
	Webui   *roleDeviceWebuiXml   `xml:"webui,omitempty"`
	Xmlapi  *roleDeviceXmlapiXml  `xml:"xmlapi,omitempty"`
	Misc    []generic.Xml         `xml:",any"`
}
type roleDeviceRestapiXml struct {
	Device   *roleDeviceRestapiDeviceXml   `xml:"device,omitempty"`
	Network  *roleDeviceRestapiNetworkXml  `xml:"network,omitempty"`
	Objects  *roleDeviceRestapiObjectsXml  `xml:"objects,omitempty"`
	Policies *roleDeviceRestapiPoliciesXml `xml:"policies,omitempty"`
	System   *roleDeviceRestapiSystemXml   `xml:"system,omitempty"`
	Misc     []generic.Xml                 `xml:",any"`
}
type roleDeviceRestapiDeviceXml struct {
	EmailServerProfiles    *string       `xml:"email-server-profiles,omitempty"`
	HttpServerProfiles     *string       `xml:"http-server-profiles,omitempty"`
	LdapServerProfiles     *string       `xml:"ldap-server-profiles,omitempty"`
	LogInterfaceSetting    *string       `xml:"log-interface-setting,omitempty"`
	SnmpTrapServerProfiles *string       `xml:"snmp-trap-server-profiles,omitempty"`
	SyslogServerProfiles   *string       `xml:"syslog-server-profiles,omitempty"`
	VirtualSystems         *string       `xml:"virtual-systems,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type roleDeviceRestapiNetworkXml struct {
	AggregateEthernetInterfaces             *string       `xml:"aggregate-ethernet-interfaces,omitempty"`
	BfdNetworkProfiles                      *string       `xml:"bfd-network-profiles,omitempty"`
	BgpRoutingProfiles                      *string       `xml:"bgp-routing-profiles,omitempty"`
	DhcpRelays                              *string       `xml:"dhcp-relays,omitempty"`
	DhcpServers                             *string       `xml:"dhcp-servers,omitempty"`
	DnsProxies                              *string       `xml:"dns-proxies,omitempty"`
	EthernetInterfaces                      *string       `xml:"ethernet-interfaces,omitempty"`
	GlobalprotectClientlessAppGroups        *string       `xml:"globalprotect-clientless-app-groups,omitempty"`
	GlobalprotectClientlessApps             *string       `xml:"globalprotect-clientless-apps,omitempty"`
	GlobalprotectGateways                   *string       `xml:"globalprotect-gateways,omitempty"`
	GlobalprotectIpsecCryptoNetworkProfiles *string       `xml:"globalprotect-ipsec-crypto-network-profiles,omitempty"`
	GlobalprotectMdmServers                 *string       `xml:"globalprotect-mdm-servers,omitempty"`
	GlobalprotectPortals                    *string       `xml:"globalprotect-portals,omitempty"`
	GreTunnels                              *string       `xml:"gre-tunnels,omitempty"`
	IkeCryptoNetworkProfiles                *string       `xml:"ike-crypto-network-profiles,omitempty"`
	IkeGatewayNetworkProfiles               *string       `xml:"ike-gateway-network-profiles,omitempty"`
	InterfaceManagementNetworkProfiles      *string       `xml:"interface-management-network-profiles,omitempty"`
	IpsecCryptoNetworkProfiles              *string       `xml:"ipsec-crypto-network-profiles,omitempty"`
	IpsecTunnels                            *string       `xml:"ipsec-tunnels,omitempty"`
	Lldp                                    *string       `xml:"lldp,omitempty"`
	LldpNetworkProfiles                     *string       `xml:"lldp-network-profiles,omitempty"`
	LogicalRouters                          *string       `xml:"logical-routers,omitempty"`
	LoopbackInterfaces                      *string       `xml:"loopback-interfaces,omitempty"`
	QosInterfaces                           *string       `xml:"qos-interfaces,omitempty"`
	QosNetworkProfiles                      *string       `xml:"qos-network-profiles,omitempty"`
	SdwanInterfaceProfiles                  *string       `xml:"sdwan-interface-profiles,omitempty"`
	SdwanInterfaces                         *string       `xml:"sdwan-interfaces,omitempty"`
	TunnelInterfaces                        *string       `xml:"tunnel-interfaces,omitempty"`
	TunnelMonitorNetworkProfiles            *string       `xml:"tunnel-monitor-network-profiles,omitempty"`
	VirtualRouters                          *string       `xml:"virtual-routers,omitempty"`
	VirtualWires                            *string       `xml:"virtual-wires,omitempty"`
	VlanInterfaces                          *string       `xml:"vlan-interfaces,omitempty"`
	Vlans                                   *string       `xml:"vlans,omitempty"`
	ZoneProtectionNetworkProfiles           *string       `xml:"zone-protection-network-profiles,omitempty"`
	Zones                                   *string       `xml:"zones,omitempty"`
	Misc                                    []generic.Xml `xml:",any"`
}
type roleDeviceRestapiObjectsXml struct {
	AddressGroups                           *string       `xml:"address-groups,omitempty"`
	Addresses                               *string       `xml:"addresses,omitempty"`
	AntiSpywareSecurityProfiles             *string       `xml:"anti-spyware-security-profiles,omitempty"`
	AntivirusSecurityProfiles               *string       `xml:"antivirus-security-profiles,omitempty"`
	ApplicationFilters                      *string       `xml:"application-filters,omitempty"`
	ApplicationGroups                       *string       `xml:"application-groups,omitempty"`
	Applications                            *string       `xml:"applications,omitempty"`
	AuthenticationEnforcements              *string       `xml:"authentication-enforcements,omitempty"`
	CustomDataPatterns                      *string       `xml:"custom-data-patterns,omitempty"`
	CustomSpywareSignatures                 *string       `xml:"custom-spyware-signatures,omitempty"`
	CustomUrlCategories                     *string       `xml:"custom-url-categories,omitempty"`
	CustomVulnerabilitySignatures           *string       `xml:"custom-vulnerability-signatures,omitempty"`
	DataFilteringSecurityProfiles           *string       `xml:"data-filtering-security-profiles,omitempty"`
	DecryptionProfiles                      *string       `xml:"decryption-profiles,omitempty"`
	Devices                                 *string       `xml:"devices,omitempty"`
	DosProtectionSecurityProfiles           *string       `xml:"dos-protection-security-profiles,omitempty"`
	DynamicUserGroups                       *string       `xml:"dynamic-user-groups,omitempty"`
	ExternalDynamicLists                    *string       `xml:"external-dynamic-lists,omitempty"`
	FileBlockingSecurityProfiles            *string       `xml:"file-blocking-security-profiles,omitempty"`
	GlobalprotectHipObjects                 *string       `xml:"globalprotect-hip-objects,omitempty"`
	GlobalprotectHipProfiles                *string       `xml:"globalprotect-hip-profiles,omitempty"`
	GtpProtectionSecurityProfiles           *string       `xml:"gtp-protection-security-profiles,omitempty"`
	LogForwardingProfiles                   *string       `xml:"log-forwarding-profiles,omitempty"`
	PacketBrokerProfiles                    *string       `xml:"packet-broker-profiles,omitempty"`
	Regions                                 *string       `xml:"regions,omitempty"`
	Schedules                               *string       `xml:"schedules,omitempty"`
	SctpProtectionSecurityProfiles          *string       `xml:"sctp-protection-security-profiles,omitempty"`
	SdwanErrorCorrectionProfiles            *string       `xml:"sdwan-error-correction-profiles,omitempty"`
	SdwanPathQualityProfiles                *string       `xml:"sdwan-path-quality-profiles,omitempty"`
	SdwanSaasQualityProfiles                *string       `xml:"sdwan-saas-quality-profiles,omitempty"`
	SdwanTrafficDistributionProfiles        *string       `xml:"sdwan-traffic-distribution-profiles,omitempty"`
	SecurityProfileGroups                   *string       `xml:"security-profile-groups,omitempty"`
	ServiceGroups                           *string       `xml:"service-groups,omitempty"`
	Services                                *string       `xml:"services,omitempty"`
	Tags                                    *string       `xml:"tags,omitempty"`
	UrlFilteringSecurityProfiles            *string       `xml:"url-filtering-security-profiles,omitempty"`
	VulnerabilityProtectionSecurityProfiles *string       `xml:"vulnerability-protection-security-profiles,omitempty"`
	WildfireAnalysisSecurityProfiles        *string       `xml:"wildfire-analysis-security-profiles,omitempty"`
	Misc                                    []generic.Xml `xml:",any"`
}
type roleDeviceRestapiPoliciesXml struct {
	ApplicationOverrideRules   *string       `xml:"application-override-rules,omitempty"`
	AuthenticationRules        *string       `xml:"authentication-rules,omitempty"`
	DecryptionRules            *string       `xml:"decryption-rules,omitempty"`
	DosRules                   *string       `xml:"dos-rules,omitempty"`
	NatRules                   *string       `xml:"nat-rules,omitempty"`
	NetworkPacketBrokerRules   *string       `xml:"network-packet-broker-rules,omitempty"`
	PolicyBasedForwardingRules *string       `xml:"policy-based-forwarding-rules,omitempty"`
	QosRules                   *string       `xml:"qos-rules,omitempty"`
	SdwanRules                 *string       `xml:"sdwan-rules,omitempty"`
	SecurityRules              *string       `xml:"security-rules,omitempty"`
	TunnelInspectionRules      *string       `xml:"tunnel-inspection-rules,omitempty"`
	Misc                       []generic.Xml `xml:",any"`
}
type roleDeviceRestapiSystemXml struct {
	Configuration *string       `xml:"configuration,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type roleDeviceWebuiXml struct {
	Acc        *string                       `xml:"acc,omitempty"`
	Commit     *roleDeviceWebuiCommitXml     `xml:"commit,omitempty"`
	Dashboard  *string                       `xml:"dashboard,omitempty"`
	Device     *roleDeviceWebuiDeviceXml     `xml:"device,omitempty"`
	Global     *roleDeviceWebuiGlobalXml     `xml:"global,omitempty"`
	Monitor    *roleDeviceWebuiMonitorXml    `xml:"monitor,omitempty"`
	Network    *roleDeviceWebuiNetworkXml    `xml:"network,omitempty"`
	Objects    *roleDeviceWebuiObjectsXml    `xml:"objects,omitempty"`
	Operations *roleDeviceWebuiOperationsXml `xml:"operations,omitempty"`
	Policies   *roleDeviceWebuiPoliciesXml   `xml:"policies,omitempty"`
	Privacy    *roleDeviceWebuiPrivacyXml    `xml:"privacy,omitempty"`
	Save       *roleDeviceWebuiSaveXml       `xml:"save,omitempty"`
	Tasks      *string                       `xml:"tasks,omitempty"`
	Validate   *string                       `xml:"validate,omitempty"`
	Misc       []generic.Xml                 `xml:",any"`
}
type roleDeviceWebuiCommitXml struct {
	CommitForOtherAdmins *string       `xml:"commit-for-other-admins,omitempty"`
	Device               *string       `xml:"device,omitempty"`
	ObjectLevelChanges   *string       `xml:"object-level-changes,omitempty"`
	Misc                 []generic.Xml `xml:",any"`
}
type roleDeviceWebuiDeviceXml struct {
	AccessDomain           *string                                        `xml:"access-domain,omitempty"`
	AdminRoles             *string                                        `xml:"admin-roles,omitempty"`
	Administrators         *string                                        `xml:"administrators,omitempty"`
	AuthenticationProfile  *string                                        `xml:"authentication-profile,omitempty"`
	AuthenticationSequence *string                                        `xml:"authentication-sequence,omitempty"`
	BlockPages             *string                                        `xml:"block-pages,omitempty"`
	CertificateManagement  *roleDeviceWebuiDeviceCertificateManagementXml `xml:"certificate-management,omitempty"`
	ConfigAudit            *string                                        `xml:"config-audit,omitempty"`
	DataRedistribution     *string                                        `xml:"data-redistribution,omitempty"`
	DeviceQuarantine       *string                                        `xml:"device-quarantine,omitempty"`
	DhcpSyslogServer       *string                                        `xml:"dhcp-syslog-server,omitempty"`
	DynamicUpdates         *string                                        `xml:"dynamic-updates,omitempty"`
	GlobalProtectClient    *string                                        `xml:"global-protect-client,omitempty"`
	HighAvailability       *string                                        `xml:"high-availability,omitempty"`
	Licenses               *string                                        `xml:"licenses,omitempty"`
	LocalUserDatabase      *roleDeviceWebuiDeviceLocalUserDatabaseXml     `xml:"local-user-database,omitempty"`
	LogFwdCard             *string                                        `xml:"log-fwd-card,omitempty"`
	LogSettings            *roleDeviceWebuiDeviceLogSettingsXml           `xml:"log-settings,omitempty"`
	MasterKey              *string                                        `xml:"master-key,omitempty"`
	Plugins                *string                                        `xml:"plugins,omitempty"`
	PolicyRecommendations  *roleDeviceWebuiDevicePolicyRecommendationsXml `xml:"policy-recommendations,omitempty"`
	ScheduledLogExport     *string                                        `xml:"scheduled-log-export,omitempty"`
	ServerProfile          *roleDeviceWebuiDeviceServerProfileXml         `xml:"server-profile,omitempty"`
	Setup                  *roleDeviceWebuiDeviceSetupXml                 `xml:"setup,omitempty"`
	SharedGateways         *string                                        `xml:"shared-gateways,omitempty"`
	Software               *string                                        `xml:"software,omitempty"`
	Support                *string                                        `xml:"support,omitempty"`
	Troubleshooting        *string                                        `xml:"troubleshooting,omitempty"`
	UserIdentification     *string                                        `xml:"user-identification,omitempty"`
	VirtualSystems         *string                                        `xml:"virtual-systems,omitempty"`
	VmInfoSource           *string                                        `xml:"vm-info-source,omitempty"`
	Misc                   []generic.Xml                                  `xml:",any"`
}
type roleDeviceWebuiDeviceCertificateManagementXml struct {
	CertificateProfile     *string       `xml:"certificate-profile,omitempty"`
	Certificates           *string       `xml:"certificates,omitempty"`
	OcspResponder          *string       `xml:"ocsp-responder,omitempty"`
	Scep                   *string       `xml:"scep,omitempty"`
	SshServiceProfile      *string       `xml:"ssh-service-profile,omitempty"`
	SslDecryptionExclusion *string       `xml:"ssl-decryption-exclusion,omitempty"`
	SslTlsServiceProfile   *string       `xml:"ssl-tls-service-profile,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type roleDeviceWebuiDeviceLocalUserDatabaseXml struct {
	UserGroups *string       `xml:"user-groups,omitempty"`
	Users      *string       `xml:"users,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type roleDeviceWebuiDeviceLogSettingsXml struct {
	CcAlarm       *string       `xml:"cc-alarm,omitempty"`
	Config        *string       `xml:"config,omitempty"`
	Correlation   *string       `xml:"correlation,omitempty"`
	Globalprotect *string       `xml:"globalprotect,omitempty"`
	Hipmatch      *string       `xml:"hipmatch,omitempty"`
	Iptag         *string       `xml:"iptag,omitempty"`
	ManageLog     *string       `xml:"manage-log,omitempty"`
	System        *string       `xml:"system,omitempty"`
	UserId        *string       `xml:"user-id,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type roleDeviceWebuiDevicePolicyRecommendationsXml struct {
	Iot  *string       `xml:"iot,omitempty"`
	Saas *string       `xml:"saas,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type roleDeviceWebuiDeviceServerProfileXml struct {
	Dns      *string       `xml:"dns,omitempty"`
	Email    *string       `xml:"email,omitempty"`
	Http     *string       `xml:"http,omitempty"`
	Kerberos *string       `xml:"kerberos,omitempty"`
	Ldap     *string       `xml:"ldap,omitempty"`
	Mfa      *string       `xml:"mfa,omitempty"`
	Netflow  *string       `xml:"netflow,omitempty"`
	Radius   *string       `xml:"radius,omitempty"`
	SamlIdp  *string       `xml:"saml_idp,omitempty"`
	Scp      *string       `xml:"scp,omitempty"`
	SnmpTrap *string       `xml:"snmp-trap,omitempty"`
	Syslog   *string       `xml:"syslog,omitempty"`
	Tacplus  *string       `xml:"tacplus,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type roleDeviceWebuiDeviceSetupXml struct {
	ContentId  *string       `xml:"content-id,omitempty"`
	Hsm        *string       `xml:"hsm,omitempty"`
	Interfaces *string       `xml:"interfaces,omitempty"`
	Management *string       `xml:"management,omitempty"`
	Operations *string       `xml:"operations,omitempty"`
	Services   *string       `xml:"services,omitempty"`
	Session    *string       `xml:"session,omitempty"`
	Telemetry  *string       `xml:"telemetry,omitempty"`
	Wildfire   *string       `xml:"wildfire,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type roleDeviceWebuiGlobalXml struct {
	SystemAlarms *string       `xml:"system-alarms,omitempty"`
	Misc         []generic.Xml `xml:",any"`
}
type roleDeviceWebuiMonitorXml struct {
	AppScope                   *string                                              `xml:"app-scope,omitempty"`
	ApplicationReports         *string                                              `xml:"application-reports,omitempty"`
	AutomatedCorrelationEngine *roleDeviceWebuiMonitorAutomatedCorrelationEngineXml `xml:"automated-correlation-engine,omitempty"`
	BlockIpList                *string                                              `xml:"block-ip-list,omitempty"`
	Botnet                     *string                                              `xml:"botnet,omitempty"`
	CustomReports              *roleDeviceWebuiMonitorCustomReportsXml              `xml:"custom-reports,omitempty"`
	ExternalLogs               *string                                              `xml:"external-logs,omitempty"`
	GtpReports                 *string                                              `xml:"gtp-reports,omitempty"`
	Logs                       *roleDeviceWebuiMonitorLogsXml                       `xml:"logs,omitempty"`
	PacketCapture              *string                                              `xml:"packet-capture,omitempty"`
	PdfReports                 *roleDeviceWebuiMonitorPdfReportsXml                 `xml:"pdf-reports,omitempty"`
	SctpReports                *string                                              `xml:"sctp-reports,omitempty"`
	SessionBrowser             *string                                              `xml:"session-browser,omitempty"`
	ThreatReports              *string                                              `xml:"threat-reports,omitempty"`
	TrafficReports             *string                                              `xml:"traffic-reports,omitempty"`
	UrlFilteringReports        *string                                              `xml:"url-filtering-reports,omitempty"`
	ViewCustomReports          *string                                              `xml:"view-custom-reports,omitempty"`
	Misc                       []generic.Xml                                        `xml:",any"`
}
type roleDeviceWebuiMonitorAutomatedCorrelationEngineXml struct {
	CorrelatedEvents   *string       `xml:"correlated-events,omitempty"`
	CorrelationObjects *string       `xml:"correlation-objects,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type roleDeviceWebuiMonitorCustomReportsXml struct {
	ApplicationStatistics *string       `xml:"application-statistics,omitempty"`
	Auth                  *string       `xml:"auth,omitempty"`
	DataFilteringLog      *string       `xml:"data-filtering-log,omitempty"`
	DecryptionLog         *string       `xml:"decryption-log,omitempty"`
	DecryptionSummary     *string       `xml:"decryption-summary,omitempty"`
	Globalprotect         *string       `xml:"globalprotect,omitempty"`
	GtpLog                *string       `xml:"gtp-log,omitempty"`
	GtpSummary            *string       `xml:"gtp-summary,omitempty"`
	Hipmatch              *string       `xml:"hipmatch,omitempty"`
	Iptag                 *string       `xml:"iptag,omitempty"`
	SctpLog               *string       `xml:"sctp-log,omitempty"`
	SctpSummary           *string       `xml:"sctp-summary,omitempty"`
	ThreatLog             *string       `xml:"threat-log,omitempty"`
	ThreatSummary         *string       `xml:"threat-summary,omitempty"`
	TrafficLog            *string       `xml:"traffic-log,omitempty"`
	TrafficSummary        *string       `xml:"traffic-summary,omitempty"`
	TunnelLog             *string       `xml:"tunnel-log,omitempty"`
	TunnelSummary         *string       `xml:"tunnel-summary,omitempty"`
	UrlLog                *string       `xml:"url-log,omitempty"`
	UrlSummary            *string       `xml:"url-summary,omitempty"`
	Userid                *string       `xml:"userid,omitempty"`
	WildfireLog           *string       `xml:"wildfire-log,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type roleDeviceWebuiMonitorLogsXml struct {
	Alarm          *string       `xml:"alarm,omitempty"`
	Authentication *string       `xml:"authentication,omitempty"`
	Configuration  *string       `xml:"configuration,omitempty"`
	DataFiltering  *string       `xml:"data-filtering,omitempty"`
	Decryption     *string       `xml:"decryption,omitempty"`
	Globalprotect  *string       `xml:"globalprotect,omitempty"`
	Gtp            *string       `xml:"gtp,omitempty"`
	Hipmatch       *string       `xml:"hipmatch,omitempty"`
	Iptag          *string       `xml:"iptag,omitempty"`
	Sctp           *string       `xml:"sctp,omitempty"`
	System         *string       `xml:"system,omitempty"`
	Threat         *string       `xml:"threat,omitempty"`
	Traffic        *string       `xml:"traffic,omitempty"`
	Tunnel         *string       `xml:"tunnel,omitempty"`
	Url            *string       `xml:"url,omitempty"`
	Userid         *string       `xml:"userid,omitempty"`
	Wildfire       *string       `xml:"wildfire,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type roleDeviceWebuiMonitorPdfReportsXml struct {
	EmailScheduler             *string       `xml:"email-scheduler,omitempty"`
	ManagePdfSummary           *string       `xml:"manage-pdf-summary,omitempty"`
	PdfSummaryReports          *string       `xml:"pdf-summary-reports,omitempty"`
	ReportGroups               *string       `xml:"report-groups,omitempty"`
	SaasApplicationUsageReport *string       `xml:"saas-application-usage-report,omitempty"`
	UserActivityReport         *string       `xml:"user-activity-report,omitempty"`
	Misc                       []generic.Xml `xml:",any"`
}
type roleDeviceWebuiNetworkXml struct {
	Dhcp                  *string                                   `xml:"dhcp,omitempty"`
	DnsProxy              *string                                   `xml:"dns-proxy,omitempty"`
	GlobalProtect         *roleDeviceWebuiNetworkGlobalProtectXml   `xml:"global-protect,omitempty"`
	GreTunnels            *string                                   `xml:"gre-tunnels,omitempty"`
	Interfaces            *string                                   `xml:"interfaces,omitempty"`
	IpsecTunnels          *string                                   `xml:"ipsec-tunnels,omitempty"`
	Lldp                  *string                                   `xml:"lldp,omitempty"`
	NetworkProfiles       *roleDeviceWebuiNetworkNetworkProfilesXml `xml:"network-profiles,omitempty"`
	Qos                   *string                                   `xml:"qos,omitempty"`
	Routing               *roleDeviceWebuiNetworkRoutingXml         `xml:"routing,omitempty"`
	SdwanInterfaceProfile *string                                   `xml:"sdwan-interface-profile,omitempty"`
	SecureWebGateway      *string                                   `xml:"secure-web-gateway,omitempty"`
	VirtualRouters        *string                                   `xml:"virtual-routers,omitempty"`
	VirtualWires          *string                                   `xml:"virtual-wires,omitempty"`
	Vlans                 *string                                   `xml:"vlans,omitempty"`
	Zones                 *string                                   `xml:"zones,omitempty"`
	Misc                  []generic.Xml                             `xml:",any"`
}
type roleDeviceWebuiNetworkGlobalProtectXml struct {
	ClientlessAppGroups *string       `xml:"clientless-app-groups,omitempty"`
	ClientlessApps      *string       `xml:"clientless-apps,omitempty"`
	Gateways            *string       `xml:"gateways,omitempty"`
	Mdm                 *string       `xml:"mdm,omitempty"`
	Portals             *string       `xml:"portals,omitempty"`
	Misc                []generic.Xml `xml:",any"`
}
type roleDeviceWebuiNetworkNetworkProfilesXml struct {
	BfdProfile       *string       `xml:"bfd-profile,omitempty"`
	GpAppIpsecCrypto *string       `xml:"gp-app-ipsec-crypto,omitempty"`
	IkeCrypto        *string       `xml:"ike-crypto,omitempty"`
	IkeGateways      *string       `xml:"ike-gateways,omitempty"`
	InterfaceMgmt    *string       `xml:"interface-mgmt,omitempty"`
	IpsecCrypto      *string       `xml:"ipsec-crypto,omitempty"`
	LldpProfile      *string       `xml:"lldp-profile,omitempty"`
	QosProfile       *string       `xml:"qos-profile,omitempty"`
	TunnelMonitor    *string       `xml:"tunnel-monitor,omitempty"`
	ZoneProtection   *string       `xml:"zone-protection,omitempty"`
	Misc             []generic.Xml `xml:",any"`
}
type roleDeviceWebuiNetworkRoutingXml struct {
	LogicalRouters  *string                                          `xml:"logical-routers,omitempty"`
	RoutingProfiles *roleDeviceWebuiNetworkRoutingRoutingProfilesXml `xml:"routing-profiles,omitempty"`
	Misc            []generic.Xml                                    `xml:",any"`
}
type roleDeviceWebuiNetworkRoutingRoutingProfilesXml struct {
	Bfd       *string       `xml:"bfd,omitempty"`
	Bgp       *string       `xml:"bgp,omitempty"`
	Filters   *string       `xml:"filters,omitempty"`
	Multicast *string       `xml:"multicast,omitempty"`
	Ospf      *string       `xml:"ospf,omitempty"`
	Ospfv3    *string       `xml:"ospfv3,omitempty"`
	Ripv2     *string       `xml:"ripv2,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type roleDeviceWebuiObjectsXml struct {
	AddressGroups         *string                                    `xml:"address-groups,omitempty"`
	Addresses             *string                                    `xml:"addresses,omitempty"`
	ApplicationFilters    *string                                    `xml:"application-filters,omitempty"`
	ApplicationGroups     *string                                    `xml:"application-groups,omitempty"`
	Applications          *string                                    `xml:"applications,omitempty"`
	Authentication        *string                                    `xml:"authentication,omitempty"`
	CustomObjects         *roleDeviceWebuiObjectsCustomObjectsXml    `xml:"custom-objects,omitempty"`
	Decryption            *roleDeviceWebuiObjectsDecryptionXml       `xml:"decryption,omitempty"`
	Devices               *string                                    `xml:"devices,omitempty"`
	DynamicBlockLists     *string                                    `xml:"dynamic-block-lists,omitempty"`
	DynamicUserGroups     *string                                    `xml:"dynamic-user-groups,omitempty"`
	GlobalProtect         *roleDeviceWebuiObjectsGlobalProtectXml    `xml:"global-protect,omitempty"`
	LogForwarding         *string                                    `xml:"log-forwarding,omitempty"`
	PacketBrokerProfile   *string                                    `xml:"packet-broker-profile,omitempty"`
	Regions               *string                                    `xml:"regions,omitempty"`
	Schedules             *string                                    `xml:"schedules,omitempty"`
	Sdwan                 *roleDeviceWebuiObjectsSdwanXml            `xml:"sdwan,omitempty"`
	SecurityProfileGroups *string                                    `xml:"security-profile-groups,omitempty"`
	SecurityProfiles      *roleDeviceWebuiObjectsSecurityProfilesXml `xml:"security-profiles,omitempty"`
	ServiceGroups         *string                                    `xml:"service-groups,omitempty"`
	Services              *string                                    `xml:"services,omitempty"`
	Tags                  *string                                    `xml:"tags,omitempty"`
	Misc                  []generic.Xml                              `xml:",any"`
}
type roleDeviceWebuiObjectsCustomObjectsXml struct {
	DataPatterns  *string       `xml:"data-patterns,omitempty"`
	Spyware       *string       `xml:"spyware,omitempty"`
	UrlCategory   *string       `xml:"url-category,omitempty"`
	Vulnerability *string       `xml:"vulnerability,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type roleDeviceWebuiObjectsDecryptionXml struct {
	DecryptionProfile *string       `xml:"decryption-profile,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type roleDeviceWebuiObjectsGlobalProtectXml struct {
	HipObjects  *string       `xml:"hip-objects,omitempty"`
	HipProfiles *string       `xml:"hip-profiles,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type roleDeviceWebuiObjectsSdwanXml struct {
	SdwanDistProfile            *string       `xml:"sdwan-dist-profile,omitempty"`
	SdwanErrorCorrectionProfile *string       `xml:"sdwan-error-correction-profile,omitempty"`
	SdwanProfile                *string       `xml:"sdwan-profile,omitempty"`
	SdwanSaasQualityProfile     *string       `xml:"sdwan-saas-quality-profile,omitempty"`
	Misc                        []generic.Xml `xml:",any"`
}
type roleDeviceWebuiObjectsSecurityProfilesXml struct {
	AntiSpyware             *string       `xml:"anti-spyware,omitempty"`
	Antivirus               *string       `xml:"antivirus,omitempty"`
	DataFiltering           *string       `xml:"data-filtering,omitempty"`
	DosProtection           *string       `xml:"dos-protection,omitempty"`
	FileBlocking            *string       `xml:"file-blocking,omitempty"`
	GtpProtection           *string       `xml:"gtp-protection,omitempty"`
	SctpProtection          *string       `xml:"sctp-protection,omitempty"`
	UrlFiltering            *string       `xml:"url-filtering,omitempty"`
	VulnerabilityProtection *string       `xml:"vulnerability-protection,omitempty"`
	WildfireAnalysis        *string       `xml:"wildfire-analysis,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
}
type roleDeviceWebuiOperationsXml struct {
	DownloadCoreFiles       *string       `xml:"download-core-files,omitempty"`
	DownloadPcapFiles       *string       `xml:"download-pcap-files,omitempty"`
	GenerateStatsDumpFile   *string       `xml:"generate-stats-dump-file,omitempty"`
	GenerateTechSupportFile *string       `xml:"generate-tech-support-file,omitempty"`
	Reboot                  *string       `xml:"reboot,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
}
type roleDeviceWebuiPoliciesXml struct {
	ApplicationOverrideRulebase *string       `xml:"application-override-rulebase,omitempty"`
	AuthenticationRulebase      *string       `xml:"authentication-rulebase,omitempty"`
	DosRulebase                 *string       `xml:"dos-rulebase,omitempty"`
	NatRulebase                 *string       `xml:"nat-rulebase,omitempty"`
	NetworkPacketBrokerRulebase *string       `xml:"network-packet-broker-rulebase,omitempty"`
	PbfRulebase                 *string       `xml:"pbf-rulebase,omitempty"`
	QosRulebase                 *string       `xml:"qos-rulebase,omitempty"`
	RuleHitCountReset           *string       `xml:"rule-hit-count-reset,omitempty"`
	SdwanRulebase               *string       `xml:"sdwan-rulebase,omitempty"`
	SecurityRulebase            *string       `xml:"security-rulebase,omitempty"`
	SslDecryptionRulebase       *string       `xml:"ssl-decryption-rulebase,omitempty"`
	TunnelInspectRulebase       *string       `xml:"tunnel-inspect-rulebase,omitempty"`
	Misc                        []generic.Xml `xml:",any"`
}
type roleDeviceWebuiPrivacyXml struct {
	ShowFullIpAddresses           *string       `xml:"show-full-ip-addresses,omitempty"`
	ShowUserNamesInLogsAndReports *string       `xml:"show-user-names-in-logs-and-reports,omitempty"`
	ViewPcapFiles                 *string       `xml:"view-pcap-files,omitempty"`
	Misc                          []generic.Xml `xml:",any"`
}
type roleDeviceWebuiSaveXml struct {
	ObjectLevelChanges *string       `xml:"object-level-changes,omitempty"`
	PartialSave        *string       `xml:"partial-save,omitempty"`
	SaveForOtherAdmins *string       `xml:"save-for-other-admins,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type roleDeviceXmlapiXml struct {
	Commit *string       `xml:"commit,omitempty"`
	Config *string       `xml:"config,omitempty"`
	Export *string       `xml:"export,omitempty"`
	Import *string       `xml:"import,omitempty"`
	Iot    *string       `xml:"iot,omitempty"`
	Log    *string       `xml:"log,omitempty"`
	Op     *string       `xml:"op,omitempty"`
	Report *string       `xml:"report,omitempty"`
	UserId *string       `xml:"user-id,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type roleVsysXml struct {
	Cli     *string             `xml:"cli,omitempty"`
	Restapi *roleVsysRestapiXml `xml:"restapi,omitempty"`
	Webui   *roleVsysWebuiXml   `xml:"webui,omitempty"`
	Xmlapi  *roleVsysXmlapiXml  `xml:"xmlapi,omitempty"`
	Misc    []generic.Xml       `xml:",any"`
}
type roleVsysRestapiXml struct {
	Device   *roleVsysRestapiDeviceXml   `xml:"device,omitempty"`
	Network  *roleVsysRestapiNetworkXml  `xml:"network,omitempty"`
	Objects  *roleVsysRestapiObjectsXml  `xml:"objects,omitempty"`
	Policies *roleVsysRestapiPoliciesXml `xml:"policies,omitempty"`
	System   *roleVsysRestapiSystemXml   `xml:"system,omitempty"`
	Misc     []generic.Xml               `xml:",any"`
}
type roleVsysRestapiDeviceXml struct {
	EmailServerProfiles    *string       `xml:"email-server-profiles,omitempty"`
	HttpServerProfiles     *string       `xml:"http-server-profiles,omitempty"`
	LdapServerProfiles     *string       `xml:"ldap-server-profiles,omitempty"`
	LogInterfaceSetting    *string       `xml:"log-interface-setting,omitempty"`
	SnmpTrapServerProfiles *string       `xml:"snmp-trap-server-profiles,omitempty"`
	SyslogServerProfiles   *string       `xml:"syslog-server-profiles,omitempty"`
	VirtualSystems         *string       `xml:"virtual-systems,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type roleVsysRestapiNetworkXml struct {
	GlobalprotectClientlessAppGroups *string       `xml:"globalprotect-clientless-app-groups,omitempty"`
	GlobalprotectClientlessApps      *string       `xml:"globalprotect-clientless-apps,omitempty"`
	GlobalprotectGateways            *string       `xml:"globalprotect-gateways,omitempty"`
	GlobalprotectMdmServers          *string       `xml:"globalprotect-mdm-servers,omitempty"`
	GlobalprotectPortals             *string       `xml:"globalprotect-portals,omitempty"`
	SdwanInterfaceProfiles           *string       `xml:"sdwan-interface-profiles,omitempty"`
	Zones                            *string       `xml:"zones,omitempty"`
	Misc                             []generic.Xml `xml:",any"`
}
type roleVsysRestapiObjectsXml struct {
	AddressGroups                           *string       `xml:"address-groups,omitempty"`
	Addresses                               *string       `xml:"addresses,omitempty"`
	AntiSpywareSecurityProfiles             *string       `xml:"anti-spyware-security-profiles,omitempty"`
	AntivirusSecurityProfiles               *string       `xml:"antivirus-security-profiles,omitempty"`
	ApplicationFilters                      *string       `xml:"application-filters,omitempty"`
	ApplicationGroups                       *string       `xml:"application-groups,omitempty"`
	Applications                            *string       `xml:"applications,omitempty"`
	AuthenticationEnforcements              *string       `xml:"authentication-enforcements,omitempty"`
	CustomDataPatterns                      *string       `xml:"custom-data-patterns,omitempty"`
	CustomSpywareSignatures                 *string       `xml:"custom-spyware-signatures,omitempty"`
	CustomUrlCategories                     *string       `xml:"custom-url-categories,omitempty"`
	CustomVulnerabilitySignatures           *string       `xml:"custom-vulnerability-signatures,omitempty"`
	DataFilteringSecurityProfiles           *string       `xml:"data-filtering-security-profiles,omitempty"`
	DecryptionProfiles                      *string       `xml:"decryption-profiles,omitempty"`
	Devices                                 *string       `xml:"devices,omitempty"`
	DosProtectionSecurityProfiles           *string       `xml:"dos-protection-security-profiles,omitempty"`
	DynamicUserGroups                       *string       `xml:"dynamic-user-groups,omitempty"`
	ExternalDynamicLists                    *string       `xml:"external-dynamic-lists,omitempty"`
	FileBlockingSecurityProfiles            *string       `xml:"file-blocking-security-profiles,omitempty"`
	GlobalprotectHipObjects                 *string       `xml:"globalprotect-hip-objects,omitempty"`
	GlobalprotectHipProfiles                *string       `xml:"globalprotect-hip-profiles,omitempty"`
	GtpProtectionSecurityProfiles           *string       `xml:"gtp-protection-security-profiles,omitempty"`
	LogForwardingProfiles                   *string       `xml:"log-forwarding-profiles,omitempty"`
	PacketBrokerProfiles                    *string       `xml:"packet-broker-profiles,omitempty"`
	Regions                                 *string       `xml:"regions,omitempty"`
	Schedules                               *string       `xml:"schedules,omitempty"`
	SctpProtectionSecurityProfiles          *string       `xml:"sctp-protection-security-profiles,omitempty"`
	SdwanErrorCorrectionProfiles            *string       `xml:"sdwan-error-correction-profiles,omitempty"`
	SdwanPathQualityProfiles                *string       `xml:"sdwan-path-quality-profiles,omitempty"`
	SdwanSaasQualityProfiles                *string       `xml:"sdwan-saas-quality-profiles,omitempty"`
	SdwanTrafficDistributionProfiles        *string       `xml:"sdwan-traffic-distribution-profiles,omitempty"`
	SecurityProfileGroups                   *string       `xml:"security-profile-groups,omitempty"`
	ServiceGroups                           *string       `xml:"service-groups,omitempty"`
	Services                                *string       `xml:"services,omitempty"`
	Tags                                    *string       `xml:"tags,omitempty"`
	UrlFilteringSecurityProfiles            *string       `xml:"url-filtering-security-profiles,omitempty"`
	VulnerabilityProtectionSecurityProfiles *string       `xml:"vulnerability-protection-security-profiles,omitempty"`
	WildfireAnalysisSecurityProfiles        *string       `xml:"wildfire-analysis-security-profiles,omitempty"`
	Misc                                    []generic.Xml `xml:",any"`
}
type roleVsysRestapiPoliciesXml struct {
	ApplicationOverrideRules   *string       `xml:"application-override-rules,omitempty"`
	AuthenticationRules        *string       `xml:"authentication-rules,omitempty"`
	DecryptionRules            *string       `xml:"decryption-rules,omitempty"`
	DosRules                   *string       `xml:"dos-rules,omitempty"`
	NatRules                   *string       `xml:"nat-rules,omitempty"`
	NetworkPacketBrokerRules   *string       `xml:"network-packet-broker-rules,omitempty"`
	PolicyBasedForwardingRules *string       `xml:"policy-based-forwarding-rules,omitempty"`
	QosRules                   *string       `xml:"qos-rules,omitempty"`
	SdwanRules                 *string       `xml:"sdwan-rules,omitempty"`
	SecurityRules              *string       `xml:"security-rules,omitempty"`
	TunnelInspectionRules      *string       `xml:"tunnel-inspection-rules,omitempty"`
	Misc                       []generic.Xml `xml:",any"`
}
type roleVsysRestapiSystemXml struct {
	Configuration *string       `xml:"configuration,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type roleVsysWebuiXml struct {
	Acc        *string                     `xml:"acc,omitempty"`
	Commit     *roleVsysWebuiCommitXml     `xml:"commit,omitempty"`
	Dashboard  *string                     `xml:"dashboard,omitempty"`
	Device     *roleVsysWebuiDeviceXml     `xml:"device,omitempty"`
	Monitor    *roleVsysWebuiMonitorXml    `xml:"monitor,omitempty"`
	Network    *roleVsysWebuiNetworkXml    `xml:"network,omitempty"`
	Objects    *roleVsysWebuiObjectsXml    `xml:"objects,omitempty"`
	Operations *roleVsysWebuiOperationsXml `xml:"operations,omitempty"`
	Policies   *roleVsysWebuiPoliciesXml   `xml:"policies,omitempty"`
	Privacy    *roleVsysWebuiPrivacyXml    `xml:"privacy,omitempty"`
	Save       *roleVsysWebuiSaveXml       `xml:"save,omitempty"`
	Tasks      *string                     `xml:"tasks,omitempty"`
	Validate   *string                     `xml:"validate,omitempty"`
	Misc       []generic.Xml               `xml:",any"`
}
type roleVsysWebuiCommitXml struct {
	CommitForOtherAdmins *string       `xml:"commit-for-other-admins,omitempty"`
	VirtualSystems       *string       `xml:"virtual-systems,omitempty"`
	Misc                 []generic.Xml `xml:",any"`
}
type roleVsysWebuiDeviceXml struct {
	Administrators         *string                                      `xml:"administrators,omitempty"`
	AuthenticationProfile  *string                                      `xml:"authentication-profile,omitempty"`
	AuthenticationSequence *string                                      `xml:"authentication-sequence,omitempty"`
	BlockPages             *string                                      `xml:"block-pages,omitempty"`
	CertificateManagement  *roleVsysWebuiDeviceCertificateManagementXml `xml:"certificate-management,omitempty"`
	DataRedistribution     *string                                      `xml:"data-redistribution,omitempty"`
	DeviceQuarantine       *string                                      `xml:"device-quarantine,omitempty"`
	DhcpSyslogServer       *string                                      `xml:"dhcp-syslog-server,omitempty"`
	LocalUserDatabase      *roleVsysWebuiDeviceLocalUserDatabaseXml     `xml:"local-user-database,omitempty"`
	LogSettings            *roleVsysWebuiDeviceLogSettingsXml           `xml:"log-settings,omitempty"`
	PolicyRecommendations  *roleVsysWebuiDevicePolicyRecommendationsXml `xml:"policy-recommendations,omitempty"`
	ServerProfile          *roleVsysWebuiDeviceServerProfileXml         `xml:"server-profile,omitempty"`
	Setup                  *roleVsysWebuiDeviceSetupXml                 `xml:"setup,omitempty"`
	Troubleshooting        *string                                      `xml:"troubleshooting,omitempty"`
	UserIdentification     *string                                      `xml:"user-identification,omitempty"`
	VmInfoSource           *string                                      `xml:"vm-info-source,omitempty"`
	Misc                   []generic.Xml                                `xml:",any"`
}
type roleVsysWebuiDeviceCertificateManagementXml struct {
	CertificateProfile     *string       `xml:"certificate-profile,omitempty"`
	Certificates           *string       `xml:"certificates,omitempty"`
	OcspResponder          *string       `xml:"ocsp-responder,omitempty"`
	Scep                   *string       `xml:"scep,omitempty"`
	SshServiceProfile      *string       `xml:"ssh-service-profile,omitempty"`
	SslDecryptionExclusion *string       `xml:"ssl-decryption-exclusion,omitempty"`
	SslTlsServiceProfile   *string       `xml:"ssl-tls-service-profile,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type roleVsysWebuiDeviceLocalUserDatabaseXml struct {
	UserGroups *string       `xml:"user-groups,omitempty"`
	Users      *string       `xml:"users,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type roleVsysWebuiDeviceLogSettingsXml struct {
	Config        *string       `xml:"config,omitempty"`
	Correlation   *string       `xml:"correlation,omitempty"`
	Globalprotect *string       `xml:"globalprotect,omitempty"`
	Hipmatch      *string       `xml:"hipmatch,omitempty"`
	Iptag         *string       `xml:"iptag,omitempty"`
	System        *string       `xml:"system,omitempty"`
	UserId        *string       `xml:"user-id,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type roleVsysWebuiDevicePolicyRecommendationsXml struct {
	Iot  *string       `xml:"iot,omitempty"`
	Saas *string       `xml:"saas,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type roleVsysWebuiDeviceServerProfileXml struct {
	Dns      *string       `xml:"dns,omitempty"`
	Email    *string       `xml:"email,omitempty"`
	Http     *string       `xml:"http,omitempty"`
	Kerberos *string       `xml:"kerberos,omitempty"`
	Ldap     *string       `xml:"ldap,omitempty"`
	Mfa      *string       `xml:"mfa,omitempty"`
	Netflow  *string       `xml:"netflow,omitempty"`
	Radius   *string       `xml:"radius,omitempty"`
	SamlIdp  *string       `xml:"saml_idp,omitempty"`
	Scp      *string       `xml:"scp,omitempty"`
	SnmpTrap *string       `xml:"snmp-trap,omitempty"`
	Syslog   *string       `xml:"syslog,omitempty"`
	Tacplus  *string       `xml:"tacplus,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type roleVsysWebuiDeviceSetupXml struct {
	ContentId  *string       `xml:"content-id,omitempty"`
	Hsm        *string       `xml:"hsm,omitempty"`
	Interfaces *string       `xml:"interfaces,omitempty"`
	Management *string       `xml:"management,omitempty"`
	Operations *string       `xml:"operations,omitempty"`
	Services   *string       `xml:"services,omitempty"`
	Session    *string       `xml:"session,omitempty"`
	Telemetry  *string       `xml:"telemetry,omitempty"`
	Wildfire   *string       `xml:"wildfire,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type roleVsysWebuiMonitorXml struct {
	AppScope                   *string                                            `xml:"app-scope,omitempty"`
	AutomatedCorrelationEngine *roleVsysWebuiMonitorAutomatedCorrelationEngineXml `xml:"automated-correlation-engine,omitempty"`
	BlockIpList                *string                                            `xml:"block-ip-list,omitempty"`
	CustomReports              *roleVsysWebuiMonitorCustomReportsXml              `xml:"custom-reports,omitempty"`
	ExternalLogs               *string                                            `xml:"external-logs,omitempty"`
	Logs                       *roleVsysWebuiMonitorLogsXml                       `xml:"logs,omitempty"`
	PdfReports                 *roleVsysWebuiMonitorPdfReportsXml                 `xml:"pdf-reports,omitempty"`
	SessionBrowser             *string                                            `xml:"session-browser,omitempty"`
	ViewCustomReports          *string                                            `xml:"view-custom-reports,omitempty"`
	Misc                       []generic.Xml                                      `xml:",any"`
}
type roleVsysWebuiMonitorAutomatedCorrelationEngineXml struct {
	CorrelatedEvents   *string       `xml:"correlated-events,omitempty"`
	CorrelationObjects *string       `xml:"correlation-objects,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type roleVsysWebuiMonitorCustomReportsXml struct {
	ApplicationStatistics *string       `xml:"application-statistics,omitempty"`
	Auth                  *string       `xml:"auth,omitempty"`
	DataFilteringLog      *string       `xml:"data-filtering-log,omitempty"`
	DecryptionLog         *string       `xml:"decryption-log,omitempty"`
	DecryptionSummary     *string       `xml:"decryption-summary,omitempty"`
	Globalprotect         *string       `xml:"globalprotect,omitempty"`
	GtpLog                *string       `xml:"gtp-log,omitempty"`
	GtpSummary            *string       `xml:"gtp-summary,omitempty"`
	Hipmatch              *string       `xml:"hipmatch,omitempty"`
	Iptag                 *string       `xml:"iptag,omitempty"`
	SctpLog               *string       `xml:"sctp-log,omitempty"`
	SctpSummary           *string       `xml:"sctp-summary,omitempty"`
	ThreatLog             *string       `xml:"threat-log,omitempty"`
	ThreatSummary         *string       `xml:"threat-summary,omitempty"`
	TrafficLog            *string       `xml:"traffic-log,omitempty"`
	TrafficSummary        *string       `xml:"traffic-summary,omitempty"`
	TunnelLog             *string       `xml:"tunnel-log,omitempty"`
	TunnelSummary         *string       `xml:"tunnel-summary,omitempty"`
	UrlLog                *string       `xml:"url-log,omitempty"`
	UrlSummary            *string       `xml:"url-summary,omitempty"`
	Userid                *string       `xml:"userid,omitempty"`
	WildfireLog           *string       `xml:"wildfire-log,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type roleVsysWebuiMonitorLogsXml struct {
	Authentication *string       `xml:"authentication,omitempty"`
	DataFiltering  *string       `xml:"data-filtering,omitempty"`
	Decryption     *string       `xml:"decryption,omitempty"`
	Globalprotect  *string       `xml:"globalprotect,omitempty"`
	Gtp            *string       `xml:"gtp,omitempty"`
	Hipmatch       *string       `xml:"hipmatch,omitempty"`
	Iptag          *string       `xml:"iptag,omitempty"`
	Sctp           *string       `xml:"sctp,omitempty"`
	Threat         *string       `xml:"threat,omitempty"`
	Traffic        *string       `xml:"traffic,omitempty"`
	Tunnel         *string       `xml:"tunnel,omitempty"`
	Url            *string       `xml:"url,omitempty"`
	Userid         *string       `xml:"userid,omitempty"`
	Wildfire       *string       `xml:"wildfire,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type roleVsysWebuiMonitorPdfReportsXml struct {
	EmailScheduler             *string       `xml:"email-scheduler,omitempty"`
	ManagePdfSummary           *string       `xml:"manage-pdf-summary,omitempty"`
	PdfSummaryReports          *string       `xml:"pdf-summary-reports,omitempty"`
	ReportGroups               *string       `xml:"report-groups,omitempty"`
	SaasApplicationUsageReport *string       `xml:"saas-application-usage-report,omitempty"`
	UserActivityReport         *string       `xml:"user-activity-report,omitempty"`
	Misc                       []generic.Xml `xml:",any"`
}
type roleVsysWebuiNetworkXml struct {
	GlobalProtect         *roleVsysWebuiNetworkGlobalProtectXml `xml:"global-protect,omitempty"`
	SdwanInterfaceProfile *string                               `xml:"sdwan-interface-profile,omitempty"`
	Zones                 *string                               `xml:"zones,omitempty"`
	Misc                  []generic.Xml                         `xml:",any"`
}
type roleVsysWebuiNetworkGlobalProtectXml struct {
	ClientlessAppGroups *string       `xml:"clientless-app-groups,omitempty"`
	ClientlessApps      *string       `xml:"clientless-apps,omitempty"`
	Gateways            *string       `xml:"gateways,omitempty"`
	Mdm                 *string       `xml:"mdm,omitempty"`
	Portals             *string       `xml:"portals,omitempty"`
	Misc                []generic.Xml `xml:",any"`
}
type roleVsysWebuiObjectsXml struct {
	AddressGroups         *string                                  `xml:"address-groups,omitempty"`
	Addresses             *string                                  `xml:"addresses,omitempty"`
	ApplicationFilters    *string                                  `xml:"application-filters,omitempty"`
	ApplicationGroups     *string                                  `xml:"application-groups,omitempty"`
	Applications          *string                                  `xml:"applications,omitempty"`
	Authentication        *string                                  `xml:"authentication,omitempty"`
	CustomObjects         *roleVsysWebuiObjectsCustomObjectsXml    `xml:"custom-objects,omitempty"`
	Decryption            *roleVsysWebuiObjectsDecryptionXml       `xml:"decryption,omitempty"`
	Devices               *string                                  `xml:"devices,omitempty"`
	DynamicBlockLists     *string                                  `xml:"dynamic-block-lists,omitempty"`
	DynamicUserGroups     *string                                  `xml:"dynamic-user-groups,omitempty"`
	GlobalProtect         *roleVsysWebuiObjectsGlobalProtectXml    `xml:"global-protect,omitempty"`
	LogForwarding         *string                                  `xml:"log-forwarding,omitempty"`
	PacketBrokerProfile   *string                                  `xml:"packet-broker-profile,omitempty"`
	Regions               *string                                  `xml:"regions,omitempty"`
	Schedules             *string                                  `xml:"schedules,omitempty"`
	Sdwan                 *roleVsysWebuiObjectsSdwanXml            `xml:"sdwan,omitempty"`
	SecurityProfileGroups *string                                  `xml:"security-profile-groups,omitempty"`
	SecurityProfiles      *roleVsysWebuiObjectsSecurityProfilesXml `xml:"security-profiles,omitempty"`
	ServiceGroups         *string                                  `xml:"service-groups,omitempty"`
	Services              *string                                  `xml:"services,omitempty"`
	Tags                  *string                                  `xml:"tags,omitempty"`
	Misc                  []generic.Xml                            `xml:",any"`
}
type roleVsysWebuiObjectsCustomObjectsXml struct {
	DataPatterns  *string       `xml:"data-patterns,omitempty"`
	Spyware       *string       `xml:"spyware,omitempty"`
	UrlCategory   *string       `xml:"url-category,omitempty"`
	Vulnerability *string       `xml:"vulnerability,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type roleVsysWebuiObjectsDecryptionXml struct {
	DecryptionProfile *string       `xml:"decryption-profile,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type roleVsysWebuiObjectsGlobalProtectXml struct {
	HipObjects  *string       `xml:"hip-objects,omitempty"`
	HipProfiles *string       `xml:"hip-profiles,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type roleVsysWebuiObjectsSdwanXml struct {
	SdwanDistProfile            *string       `xml:"sdwan-dist-profile,omitempty"`
	SdwanErrorCorrectionProfile *string       `xml:"sdwan-error-correction-profile,omitempty"`
	SdwanProfile                *string       `xml:"sdwan-profile,omitempty"`
	SdwanSaasQualityProfile     *string       `xml:"sdwan-saas-quality-profile,omitempty"`
	Misc                        []generic.Xml `xml:",any"`
}
type roleVsysWebuiObjectsSecurityProfilesXml struct {
	AntiSpyware             *string       `xml:"anti-spyware,omitempty"`
	Antivirus               *string       `xml:"antivirus,omitempty"`
	DataFiltering           *string       `xml:"data-filtering,omitempty"`
	DosProtection           *string       `xml:"dos-protection,omitempty"`
	FileBlocking            *string       `xml:"file-blocking,omitempty"`
	GtpProtection           *string       `xml:"gtp-protection,omitempty"`
	SctpProtection          *string       `xml:"sctp-protection,omitempty"`
	UrlFiltering            *string       `xml:"url-filtering,omitempty"`
	VulnerabilityProtection *string       `xml:"vulnerability-protection,omitempty"`
	WildfireAnalysis        *string       `xml:"wildfire-analysis,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
}
type roleVsysWebuiOperationsXml struct {
	DownloadCoreFiles       *string       `xml:"download-core-files,omitempty"`
	DownloadPcapFiles       *string       `xml:"download-pcap-files,omitempty"`
	GenerateStatsDumpFile   *string       `xml:"generate-stats-dump-file,omitempty"`
	GenerateTechSupportFile *string       `xml:"generate-tech-support-file,omitempty"`
	Reboot                  *string       `xml:"reboot,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
}
type roleVsysWebuiPoliciesXml struct {
	ApplicationOverrideRulebase *string       `xml:"application-override-rulebase,omitempty"`
	AuthenticationRulebase      *string       `xml:"authentication-rulebase,omitempty"`
	DosRulebase                 *string       `xml:"dos-rulebase,omitempty"`
	NatRulebase                 *string       `xml:"nat-rulebase,omitempty"`
	NetworkPacketBrokerRulebase *string       `xml:"network-packet-broker-rulebase,omitempty"`
	PbfRulebase                 *string       `xml:"pbf-rulebase,omitempty"`
	QosRulebase                 *string       `xml:"qos-rulebase,omitempty"`
	RuleHitCountReset           *string       `xml:"rule-hit-count-reset,omitempty"`
	SdwanRulebase               *string       `xml:"sdwan-rulebase,omitempty"`
	SecurityRulebase            *string       `xml:"security-rulebase,omitempty"`
	SslDecryptionRulebase       *string       `xml:"ssl-decryption-rulebase,omitempty"`
	TunnelInspectRulebase       *string       `xml:"tunnel-inspect-rulebase,omitempty"`
	Misc                        []generic.Xml `xml:",any"`
}
type roleVsysWebuiPrivacyXml struct {
	ShowFullIpAddresses           *string       `xml:"show-full-ip-addresses,omitempty"`
	ShowUserNamesInLogsAndReports *string       `xml:"show-user-names-in-logs-and-reports,omitempty"`
	ViewPcapFiles                 *string       `xml:"view-pcap-files,omitempty"`
	Misc                          []generic.Xml `xml:",any"`
}
type roleVsysWebuiSaveXml struct {
	ObjectLevelChanges *string       `xml:"object-level-changes,omitempty"`
	PartialSave        *string       `xml:"partial-save,omitempty"`
	SaveForOtherAdmins *string       `xml:"save-for-other-admins,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type roleVsysXmlapiXml struct {
	Commit *string       `xml:"commit,omitempty"`
	Config *string       `xml:"config,omitempty"`
	Export *string       `xml:"export,omitempty"`
	Import *string       `xml:"import,omitempty"`
	Iot    *string       `xml:"iot,omitempty"`
	Log    *string       `xml:"log,omitempty"`
	Op     *string       `xml:"op,omitempty"`
	Report *string       `xml:"report,omitempty"`
	UserId *string       `xml:"user-id,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type entryXml_11_0_2 struct {
	XMLName     xml.Name        `xml:"entry"`
	Name        string          `xml:"name,attr"`
	Description *string         `xml:"description,omitempty"`
	Role        *roleXml_11_0_2 `xml:"role,omitempty"`
	Misc        []generic.Xml   `xml:",any"`
}
type roleXml_11_0_2 struct {
	Device *roleDeviceXml_11_0_2 `xml:"device,omitempty"`
	Vsys   *roleVsysXml_11_0_2   `xml:"vsys,omitempty"`
	Misc   []generic.Xml         `xml:",any"`
}
type roleDeviceXml_11_0_2 struct {
	Cli     *string                      `xml:"cli,omitempty"`
	Restapi *roleDeviceRestapiXml_11_0_2 `xml:"restapi,omitempty"`
	Webui   *roleDeviceWebuiXml_11_0_2   `xml:"webui,omitempty"`
	Xmlapi  *roleDeviceXmlapiXml_11_0_2  `xml:"xmlapi,omitempty"`
	Misc    []generic.Xml                `xml:",any"`
}
type roleDeviceRestapiXml_11_0_2 struct {
	Device   *roleDeviceRestapiDeviceXml_11_0_2   `xml:"device,omitempty"`
	Network  *roleDeviceRestapiNetworkXml_11_0_2  `xml:"network,omitempty"`
	Objects  *roleDeviceRestapiObjectsXml_11_0_2  `xml:"objects,omitempty"`
	Policies *roleDeviceRestapiPoliciesXml_11_0_2 `xml:"policies,omitempty"`
	System   *roleDeviceRestapiSystemXml_11_0_2   `xml:"system,omitempty"`
	Misc     []generic.Xml                        `xml:",any"`
}
type roleDeviceRestapiDeviceXml_11_0_2 struct {
	EmailServerProfiles    *string       `xml:"email-server-profiles,omitempty"`
	HttpServerProfiles     *string       `xml:"http-server-profiles,omitempty"`
	LdapServerProfiles     *string       `xml:"ldap-server-profiles,omitempty"`
	LogInterfaceSetting    *string       `xml:"log-interface-setting,omitempty"`
	SnmpTrapServerProfiles *string       `xml:"snmp-trap-server-profiles,omitempty"`
	SyslogServerProfiles   *string       `xml:"syslog-server-profiles,omitempty"`
	VirtualSystems         *string       `xml:"virtual-systems,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type roleDeviceRestapiNetworkXml_11_0_2 struct {
	AggregateEthernetInterfaces             *string       `xml:"aggregate-ethernet-interfaces,omitempty"`
	BfdNetworkProfiles                      *string       `xml:"bfd-network-profiles,omitempty"`
	BgpRoutingProfiles                      *string       `xml:"bgp-routing-profiles,omitempty"`
	DhcpRelays                              *string       `xml:"dhcp-relays,omitempty"`
	DhcpServers                             *string       `xml:"dhcp-servers,omitempty"`
	DnsProxies                              *string       `xml:"dns-proxies,omitempty"`
	EthernetInterfaces                      *string       `xml:"ethernet-interfaces,omitempty"`
	GlobalprotectClientlessAppGroups        *string       `xml:"globalprotect-clientless-app-groups,omitempty"`
	GlobalprotectClientlessApps             *string       `xml:"globalprotect-clientless-apps,omitempty"`
	GlobalprotectGateways                   *string       `xml:"globalprotect-gateways,omitempty"`
	GlobalprotectIpsecCryptoNetworkProfiles *string       `xml:"globalprotect-ipsec-crypto-network-profiles,omitempty"`
	GlobalprotectMdmServers                 *string       `xml:"globalprotect-mdm-servers,omitempty"`
	GlobalprotectPortals                    *string       `xml:"globalprotect-portals,omitempty"`
	GreTunnels                              *string       `xml:"gre-tunnels,omitempty"`
	IkeCryptoNetworkProfiles                *string       `xml:"ike-crypto-network-profiles,omitempty"`
	IkeGatewayNetworkProfiles               *string       `xml:"ike-gateway-network-profiles,omitempty"`
	InterfaceManagementNetworkProfiles      *string       `xml:"interface-management-network-profiles,omitempty"`
	IpsecCryptoNetworkProfiles              *string       `xml:"ipsec-crypto-network-profiles,omitempty"`
	IpsecTunnels                            *string       `xml:"ipsec-tunnels,omitempty"`
	Lldp                                    *string       `xml:"lldp,omitempty"`
	LldpNetworkProfiles                     *string       `xml:"lldp-network-profiles,omitempty"`
	LogicalRouters                          *string       `xml:"logical-routers,omitempty"`
	LoopbackInterfaces                      *string       `xml:"loopback-interfaces,omitempty"`
	QosInterfaces                           *string       `xml:"qos-interfaces,omitempty"`
	QosNetworkProfiles                      *string       `xml:"qos-network-profiles,omitempty"`
	SdwanInterfaceProfiles                  *string       `xml:"sdwan-interface-profiles,omitempty"`
	SdwanInterfaces                         *string       `xml:"sdwan-interfaces,omitempty"`
	TunnelInterfaces                        *string       `xml:"tunnel-interfaces,omitempty"`
	TunnelMonitorNetworkProfiles            *string       `xml:"tunnel-monitor-network-profiles,omitempty"`
	VirtualRouters                          *string       `xml:"virtual-routers,omitempty"`
	VirtualWires                            *string       `xml:"virtual-wires,omitempty"`
	VlanInterfaces                          *string       `xml:"vlan-interfaces,omitempty"`
	Vlans                                   *string       `xml:"vlans,omitempty"`
	ZoneProtectionNetworkProfiles           *string       `xml:"zone-protection-network-profiles,omitempty"`
	Zones                                   *string       `xml:"zones,omitempty"`
	Misc                                    []generic.Xml `xml:",any"`
}
type roleDeviceRestapiObjectsXml_11_0_2 struct {
	AddressGroups                           *string       `xml:"address-groups,omitempty"`
	Addresses                               *string       `xml:"addresses,omitempty"`
	AntiSpywareSecurityProfiles             *string       `xml:"anti-spyware-security-profiles,omitempty"`
	AntivirusSecurityProfiles               *string       `xml:"antivirus-security-profiles,omitempty"`
	ApplicationFilters                      *string       `xml:"application-filters,omitempty"`
	ApplicationGroups                       *string       `xml:"application-groups,omitempty"`
	Applications                            *string       `xml:"applications,omitempty"`
	AuthenticationEnforcements              *string       `xml:"authentication-enforcements,omitempty"`
	CustomDataPatterns                      *string       `xml:"custom-data-patterns,omitempty"`
	CustomSpywareSignatures                 *string       `xml:"custom-spyware-signatures,omitempty"`
	CustomUrlCategories                     *string       `xml:"custom-url-categories,omitempty"`
	CustomVulnerabilitySignatures           *string       `xml:"custom-vulnerability-signatures,omitempty"`
	DataFilteringSecurityProfiles           *string       `xml:"data-filtering-security-profiles,omitempty"`
	DecryptionProfiles                      *string       `xml:"decryption-profiles,omitempty"`
	Devices                                 *string       `xml:"devices,omitempty"`
	DosProtectionSecurityProfiles           *string       `xml:"dos-protection-security-profiles,omitempty"`
	DynamicUserGroups                       *string       `xml:"dynamic-user-groups,omitempty"`
	ExternalDynamicLists                    *string       `xml:"external-dynamic-lists,omitempty"`
	FileBlockingSecurityProfiles            *string       `xml:"file-blocking-security-profiles,omitempty"`
	GlobalprotectHipObjects                 *string       `xml:"globalprotect-hip-objects,omitempty"`
	GlobalprotectHipProfiles                *string       `xml:"globalprotect-hip-profiles,omitempty"`
	GtpProtectionSecurityProfiles           *string       `xml:"gtp-protection-security-profiles,omitempty"`
	LogForwardingProfiles                   *string       `xml:"log-forwarding-profiles,omitempty"`
	PacketBrokerProfiles                    *string       `xml:"packet-broker-profiles,omitempty"`
	Regions                                 *string       `xml:"regions,omitempty"`
	Schedules                               *string       `xml:"schedules,omitempty"`
	SctpProtectionSecurityProfiles          *string       `xml:"sctp-protection-security-profiles,omitempty"`
	SdwanErrorCorrectionProfiles            *string       `xml:"sdwan-error-correction-profiles,omitempty"`
	SdwanPathQualityProfiles                *string       `xml:"sdwan-path-quality-profiles,omitempty"`
	SdwanSaasQualityProfiles                *string       `xml:"sdwan-saas-quality-profiles,omitempty"`
	SdwanTrafficDistributionProfiles        *string       `xml:"sdwan-traffic-distribution-profiles,omitempty"`
	SecurityProfileGroups                   *string       `xml:"security-profile-groups,omitempty"`
	ServiceGroups                           *string       `xml:"service-groups,omitempty"`
	Services                                *string       `xml:"services,omitempty"`
	Tags                                    *string       `xml:"tags,omitempty"`
	UrlFilteringSecurityProfiles            *string       `xml:"url-filtering-security-profiles,omitempty"`
	VulnerabilityProtectionSecurityProfiles *string       `xml:"vulnerability-protection-security-profiles,omitempty"`
	WildfireAnalysisSecurityProfiles        *string       `xml:"wildfire-analysis-security-profiles,omitempty"`
	Misc                                    []generic.Xml `xml:",any"`
}
type roleDeviceRestapiPoliciesXml_11_0_2 struct {
	ApplicationOverrideRules   *string       `xml:"application-override-rules,omitempty"`
	AuthenticationRules        *string       `xml:"authentication-rules,omitempty"`
	DecryptionRules            *string       `xml:"decryption-rules,omitempty"`
	DosRules                   *string       `xml:"dos-rules,omitempty"`
	NatRules                   *string       `xml:"nat-rules,omitempty"`
	NetworkPacketBrokerRules   *string       `xml:"network-packet-broker-rules,omitempty"`
	PolicyBasedForwardingRules *string       `xml:"policy-based-forwarding-rules,omitempty"`
	QosRules                   *string       `xml:"qos-rules,omitempty"`
	SdwanRules                 *string       `xml:"sdwan-rules,omitempty"`
	SecurityRules              *string       `xml:"security-rules,omitempty"`
	TunnelInspectionRules      *string       `xml:"tunnel-inspection-rules,omitempty"`
	Misc                       []generic.Xml `xml:",any"`
}
type roleDeviceRestapiSystemXml_11_0_2 struct {
	Configuration *string       `xml:"configuration,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type roleDeviceWebuiXml_11_0_2 struct {
	Acc        *string                              `xml:"acc,omitempty"`
	Commit     *roleDeviceWebuiCommitXml_11_0_2     `xml:"commit,omitempty"`
	Dashboard  *string                              `xml:"dashboard,omitempty"`
	Device     *roleDeviceWebuiDeviceXml_11_0_2     `xml:"device,omitempty"`
	Global     *roleDeviceWebuiGlobalXml_11_0_2     `xml:"global,omitempty"`
	Monitor    *roleDeviceWebuiMonitorXml_11_0_2    `xml:"monitor,omitempty"`
	Network    *roleDeviceWebuiNetworkXml_11_0_2    `xml:"network,omitempty"`
	Objects    *roleDeviceWebuiObjectsXml_11_0_2    `xml:"objects,omitempty"`
	Operations *roleDeviceWebuiOperationsXml_11_0_2 `xml:"operations,omitempty"`
	Policies   *roleDeviceWebuiPoliciesXml_11_0_2   `xml:"policies,omitempty"`
	Privacy    *roleDeviceWebuiPrivacyXml_11_0_2    `xml:"privacy,omitempty"`
	Save       *roleDeviceWebuiSaveXml_11_0_2       `xml:"save,omitempty"`
	Tasks      *string                              `xml:"tasks,omitempty"`
	Validate   *string                              `xml:"validate,omitempty"`
	Misc       []generic.Xml                        `xml:",any"`
}
type roleDeviceWebuiCommitXml_11_0_2 struct {
	CommitForOtherAdmins *string       `xml:"commit-for-other-admins,omitempty"`
	Device               *string       `xml:"device,omitempty"`
	ObjectLevelChanges   *string       `xml:"object-level-changes,omitempty"`
	Misc                 []generic.Xml `xml:",any"`
}
type roleDeviceWebuiDeviceXml_11_0_2 struct {
	AccessDomain           *string                                               `xml:"access-domain,omitempty"`
	AdminRoles             *string                                               `xml:"admin-roles,omitempty"`
	Administrators         *string                                               `xml:"administrators,omitempty"`
	AuthenticationProfile  *string                                               `xml:"authentication-profile,omitempty"`
	AuthenticationSequence *string                                               `xml:"authentication-sequence,omitempty"`
	BlockPages             *string                                               `xml:"block-pages,omitempty"`
	CertificateManagement  *roleDeviceWebuiDeviceCertificateManagementXml_11_0_2 `xml:"certificate-management,omitempty"`
	ConfigAudit            *string                                               `xml:"config-audit,omitempty"`
	DataRedistribution     *string                                               `xml:"data-redistribution,omitempty"`
	DeviceQuarantine       *string                                               `xml:"device-quarantine,omitempty"`
	DhcpSyslogServer       *string                                               `xml:"dhcp-syslog-server,omitempty"`
	DynamicUpdates         *string                                               `xml:"dynamic-updates,omitempty"`
	GlobalProtectClient    *string                                               `xml:"global-protect-client,omitempty"`
	HighAvailability       *string                                               `xml:"high-availability,omitempty"`
	Licenses               *string                                               `xml:"licenses,omitempty"`
	LocalUserDatabase      *roleDeviceWebuiDeviceLocalUserDatabaseXml_11_0_2     `xml:"local-user-database,omitempty"`
	LogFwdCard             *string                                               `xml:"log-fwd-card,omitempty"`
	LogSettings            *roleDeviceWebuiDeviceLogSettingsXml_11_0_2           `xml:"log-settings,omitempty"`
	MasterKey              *string                                               `xml:"master-key,omitempty"`
	Plugins                *string                                               `xml:"plugins,omitempty"`
	PolicyRecommendations  *roleDeviceWebuiDevicePolicyRecommendationsXml_11_0_2 `xml:"policy-recommendations,omitempty"`
	ScheduledLogExport     *string                                               `xml:"scheduled-log-export,omitempty"`
	ServerProfile          *roleDeviceWebuiDeviceServerProfileXml_11_0_2         `xml:"server-profile,omitempty"`
	Setup                  *roleDeviceWebuiDeviceSetupXml_11_0_2                 `xml:"setup,omitempty"`
	SharedGateways         *string                                               `xml:"shared-gateways,omitempty"`
	Software               *string                                               `xml:"software,omitempty"`
	Support                *string                                               `xml:"support,omitempty"`
	Troubleshooting        *string                                               `xml:"troubleshooting,omitempty"`
	UserIdentification     *string                                               `xml:"user-identification,omitempty"`
	VirtualSystems         *string                                               `xml:"virtual-systems,omitempty"`
	VmInfoSource           *string                                               `xml:"vm-info-source,omitempty"`
	Misc                   []generic.Xml                                         `xml:",any"`
}
type roleDeviceWebuiDeviceCertificateManagementXml_11_0_2 struct {
	CertificateProfile     *string       `xml:"certificate-profile,omitempty"`
	Certificates           *string       `xml:"certificates,omitempty"`
	OcspResponder          *string       `xml:"ocsp-responder,omitempty"`
	Scep                   *string       `xml:"scep,omitempty"`
	SshServiceProfile      *string       `xml:"ssh-service-profile,omitempty"`
	SslDecryptionExclusion *string       `xml:"ssl-decryption-exclusion,omitempty"`
	SslTlsServiceProfile   *string       `xml:"ssl-tls-service-profile,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type roleDeviceWebuiDeviceLocalUserDatabaseXml_11_0_2 struct {
	UserGroups *string       `xml:"user-groups,omitempty"`
	Users      *string       `xml:"users,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type roleDeviceWebuiDeviceLogSettingsXml_11_0_2 struct {
	CcAlarm       *string       `xml:"cc-alarm,omitempty"`
	Config        *string       `xml:"config,omitempty"`
	Correlation   *string       `xml:"correlation,omitempty"`
	Globalprotect *string       `xml:"globalprotect,omitempty"`
	Hipmatch      *string       `xml:"hipmatch,omitempty"`
	Iptag         *string       `xml:"iptag,omitempty"`
	ManageLog     *string       `xml:"manage-log,omitempty"`
	System        *string       `xml:"system,omitempty"`
	UserId        *string       `xml:"user-id,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type roleDeviceWebuiDevicePolicyRecommendationsXml_11_0_2 struct {
	Iot  *string       `xml:"iot,omitempty"`
	Saas *string       `xml:"saas,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type roleDeviceWebuiDeviceServerProfileXml_11_0_2 struct {
	Dns      *string       `xml:"dns,omitempty"`
	Email    *string       `xml:"email,omitempty"`
	Http     *string       `xml:"http,omitempty"`
	Kerberos *string       `xml:"kerberos,omitempty"`
	Ldap     *string       `xml:"ldap,omitempty"`
	Mfa      *string       `xml:"mfa,omitempty"`
	Netflow  *string       `xml:"netflow,omitempty"`
	Radius   *string       `xml:"radius,omitempty"`
	SamlIdp  *string       `xml:"saml_idp,omitempty"`
	Scp      *string       `xml:"scp,omitempty"`
	SnmpTrap *string       `xml:"snmp-trap,omitempty"`
	Syslog   *string       `xml:"syslog,omitempty"`
	Tacplus  *string       `xml:"tacplus,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type roleDeviceWebuiDeviceSetupXml_11_0_2 struct {
	ContentId  *string       `xml:"content-id,omitempty"`
	Hsm        *string       `xml:"hsm,omitempty"`
	Interfaces *string       `xml:"interfaces,omitempty"`
	Management *string       `xml:"management,omitempty"`
	Operations *string       `xml:"operations,omitempty"`
	Services   *string       `xml:"services,omitempty"`
	Session    *string       `xml:"session,omitempty"`
	Telemetry  *string       `xml:"telemetry,omitempty"`
	Wildfire   *string       `xml:"wildfire,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type roleDeviceWebuiGlobalXml_11_0_2 struct {
	SystemAlarms *string       `xml:"system-alarms,omitempty"`
	Misc         []generic.Xml `xml:",any"`
}
type roleDeviceWebuiMonitorXml_11_0_2 struct {
	AppScope                   *string                                                     `xml:"app-scope,omitempty"`
	ApplicationReports         *string                                                     `xml:"application-reports,omitempty"`
	AutomatedCorrelationEngine *roleDeviceWebuiMonitorAutomatedCorrelationEngineXml_11_0_2 `xml:"automated-correlation-engine,omitempty"`
	BlockIpList                *string                                                     `xml:"block-ip-list,omitempty"`
	Botnet                     *string                                                     `xml:"botnet,omitempty"`
	CustomReports              *roleDeviceWebuiMonitorCustomReportsXml_11_0_2              `xml:"custom-reports,omitempty"`
	ExternalLogs               *string                                                     `xml:"external-logs,omitempty"`
	GtpReports                 *string                                                     `xml:"gtp-reports,omitempty"`
	Logs                       *roleDeviceWebuiMonitorLogsXml_11_0_2                       `xml:"logs,omitempty"`
	PacketCapture              *string                                                     `xml:"packet-capture,omitempty"`
	PdfReports                 *roleDeviceWebuiMonitorPdfReportsXml_11_0_2                 `xml:"pdf-reports,omitempty"`
	SctpReports                *string                                                     `xml:"sctp-reports,omitempty"`
	SessionBrowser             *string                                                     `xml:"session-browser,omitempty"`
	ThreatReports              *string                                                     `xml:"threat-reports,omitempty"`
	TrafficReports             *string                                                     `xml:"traffic-reports,omitempty"`
	UrlFilteringReports        *string                                                     `xml:"url-filtering-reports,omitempty"`
	ViewCustomReports          *string                                                     `xml:"view-custom-reports,omitempty"`
	Misc                       []generic.Xml                                               `xml:",any"`
}
type roleDeviceWebuiMonitorAutomatedCorrelationEngineXml_11_0_2 struct {
	CorrelatedEvents   *string       `xml:"correlated-events,omitempty"`
	CorrelationObjects *string       `xml:"correlation-objects,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type roleDeviceWebuiMonitorCustomReportsXml_11_0_2 struct {
	ApplicationStatistics *string       `xml:"application-statistics,omitempty"`
	Auth                  *string       `xml:"auth,omitempty"`
	DataFilteringLog      *string       `xml:"data-filtering-log,omitempty"`
	DecryptionLog         *string       `xml:"decryption-log,omitempty"`
	DecryptionSummary     *string       `xml:"decryption-summary,omitempty"`
	Globalprotect         *string       `xml:"globalprotect,omitempty"`
	GtpLog                *string       `xml:"gtp-log,omitempty"`
	GtpSummary            *string       `xml:"gtp-summary,omitempty"`
	Hipmatch              *string       `xml:"hipmatch,omitempty"`
	Iptag                 *string       `xml:"iptag,omitempty"`
	SctpLog               *string       `xml:"sctp-log,omitempty"`
	SctpSummary           *string       `xml:"sctp-summary,omitempty"`
	ThreatLog             *string       `xml:"threat-log,omitempty"`
	ThreatSummary         *string       `xml:"threat-summary,omitempty"`
	TrafficLog            *string       `xml:"traffic-log,omitempty"`
	TrafficSummary        *string       `xml:"traffic-summary,omitempty"`
	TunnelLog             *string       `xml:"tunnel-log,omitempty"`
	TunnelSummary         *string       `xml:"tunnel-summary,omitempty"`
	UrlLog                *string       `xml:"url-log,omitempty"`
	UrlSummary            *string       `xml:"url-summary,omitempty"`
	Userid                *string       `xml:"userid,omitempty"`
	WildfireLog           *string       `xml:"wildfire-log,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type roleDeviceWebuiMonitorLogsXml_11_0_2 struct {
	Alarm          *string       `xml:"alarm,omitempty"`
	Authentication *string       `xml:"authentication,omitempty"`
	Configuration  *string       `xml:"configuration,omitempty"`
	DataFiltering  *string       `xml:"data-filtering,omitempty"`
	Decryption     *string       `xml:"decryption,omitempty"`
	Globalprotect  *string       `xml:"globalprotect,omitempty"`
	Gtp            *string       `xml:"gtp,omitempty"`
	Hipmatch       *string       `xml:"hipmatch,omitempty"`
	Iptag          *string       `xml:"iptag,omitempty"`
	Sctp           *string       `xml:"sctp,omitempty"`
	System         *string       `xml:"system,omitempty"`
	Threat         *string       `xml:"threat,omitempty"`
	Traffic        *string       `xml:"traffic,omitempty"`
	Tunnel         *string       `xml:"tunnel,omitempty"`
	Url            *string       `xml:"url,omitempty"`
	Userid         *string       `xml:"userid,omitempty"`
	Wildfire       *string       `xml:"wildfire,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type roleDeviceWebuiMonitorPdfReportsXml_11_0_2 struct {
	EmailScheduler             *string       `xml:"email-scheduler,omitempty"`
	ManagePdfSummary           *string       `xml:"manage-pdf-summary,omitempty"`
	PdfSummaryReports          *string       `xml:"pdf-summary-reports,omitempty"`
	ReportGroups               *string       `xml:"report-groups,omitempty"`
	SaasApplicationUsageReport *string       `xml:"saas-application-usage-report,omitempty"`
	UserActivityReport         *string       `xml:"user-activity-report,omitempty"`
	Misc                       []generic.Xml `xml:",any"`
}
type roleDeviceWebuiNetworkXml_11_0_2 struct {
	Dhcp                  *string                                          `xml:"dhcp,omitempty"`
	DnsProxy              *string                                          `xml:"dns-proxy,omitempty"`
	GlobalProtect         *roleDeviceWebuiNetworkGlobalProtectXml_11_0_2   `xml:"global-protect,omitempty"`
	GreTunnels            *string                                          `xml:"gre-tunnels,omitempty"`
	Interfaces            *string                                          `xml:"interfaces,omitempty"`
	IpsecTunnels          *string                                          `xml:"ipsec-tunnels,omitempty"`
	Lldp                  *string                                          `xml:"lldp,omitempty"`
	NetworkProfiles       *roleDeviceWebuiNetworkNetworkProfilesXml_11_0_2 `xml:"network-profiles,omitempty"`
	Qos                   *string                                          `xml:"qos,omitempty"`
	Routing               *roleDeviceWebuiNetworkRoutingXml_11_0_2         `xml:"routing,omitempty"`
	SdwanInterfaceProfile *string                                          `xml:"sdwan-interface-profile,omitempty"`
	SecureWebGateway      *string                                          `xml:"secure-web-gateway,omitempty"`
	VirtualRouters        *string                                          `xml:"virtual-routers,omitempty"`
	VirtualWires          *string                                          `xml:"virtual-wires,omitempty"`
	Vlans                 *string                                          `xml:"vlans,omitempty"`
	Zones                 *string                                          `xml:"zones,omitempty"`
	Misc                  []generic.Xml                                    `xml:",any"`
}
type roleDeviceWebuiNetworkGlobalProtectXml_11_0_2 struct {
	ClientlessAppGroups *string       `xml:"clientless-app-groups,omitempty"`
	ClientlessApps      *string       `xml:"clientless-apps,omitempty"`
	Gateways            *string       `xml:"gateways,omitempty"`
	Mdm                 *string       `xml:"mdm,omitempty"`
	Portals             *string       `xml:"portals,omitempty"`
	Misc                []generic.Xml `xml:",any"`
}
type roleDeviceWebuiNetworkNetworkProfilesXml_11_0_2 struct {
	BfdProfile       *string       `xml:"bfd-profile,omitempty"`
	GpAppIpsecCrypto *string       `xml:"gp-app-ipsec-crypto,omitempty"`
	IkeCrypto        *string       `xml:"ike-crypto,omitempty"`
	IkeGateways      *string       `xml:"ike-gateways,omitempty"`
	InterfaceMgmt    *string       `xml:"interface-mgmt,omitempty"`
	IpsecCrypto      *string       `xml:"ipsec-crypto,omitempty"`
	LldpProfile      *string       `xml:"lldp-profile,omitempty"`
	QosProfile       *string       `xml:"qos-profile,omitempty"`
	TunnelMonitor    *string       `xml:"tunnel-monitor,omitempty"`
	ZoneProtection   *string       `xml:"zone-protection,omitempty"`
	Misc             []generic.Xml `xml:",any"`
}
type roleDeviceWebuiNetworkRoutingXml_11_0_2 struct {
	LogicalRouters  *string                                                 `xml:"logical-routers,omitempty"`
	RoutingProfiles *roleDeviceWebuiNetworkRoutingRoutingProfilesXml_11_0_2 `xml:"routing-profiles,omitempty"`
	Misc            []generic.Xml                                           `xml:",any"`
}
type roleDeviceWebuiNetworkRoutingRoutingProfilesXml_11_0_2 struct {
	Bfd       *string       `xml:"bfd,omitempty"`
	Bgp       *string       `xml:"bgp,omitempty"`
	Filters   *string       `xml:"filters,omitempty"`
	Multicast *string       `xml:"multicast,omitempty"`
	Ospf      *string       `xml:"ospf,omitempty"`
	Ospfv3    *string       `xml:"ospfv3,omitempty"`
	Ripv2     *string       `xml:"ripv2,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type roleDeviceWebuiObjectsXml_11_0_2 struct {
	AddressGroups         *string                                           `xml:"address-groups,omitempty"`
	Addresses             *string                                           `xml:"addresses,omitempty"`
	ApplicationFilters    *string                                           `xml:"application-filters,omitempty"`
	ApplicationGroups     *string                                           `xml:"application-groups,omitempty"`
	Applications          *string                                           `xml:"applications,omitempty"`
	Authentication        *string                                           `xml:"authentication,omitempty"`
	CustomObjects         *roleDeviceWebuiObjectsCustomObjectsXml_11_0_2    `xml:"custom-objects,omitempty"`
	Decryption            *roleDeviceWebuiObjectsDecryptionXml_11_0_2       `xml:"decryption,omitempty"`
	Devices               *string                                           `xml:"devices,omitempty"`
	DynamicBlockLists     *string                                           `xml:"dynamic-block-lists,omitempty"`
	DynamicUserGroups     *string                                           `xml:"dynamic-user-groups,omitempty"`
	GlobalProtect         *roleDeviceWebuiObjectsGlobalProtectXml_11_0_2    `xml:"global-protect,omitempty"`
	LogForwarding         *string                                           `xml:"log-forwarding,omitempty"`
	PacketBrokerProfile   *string                                           `xml:"packet-broker-profile,omitempty"`
	Regions               *string                                           `xml:"regions,omitempty"`
	Schedules             *string                                           `xml:"schedules,omitempty"`
	Sdwan                 *roleDeviceWebuiObjectsSdwanXml_11_0_2            `xml:"sdwan,omitempty"`
	SecurityProfileGroups *string                                           `xml:"security-profile-groups,omitempty"`
	SecurityProfiles      *roleDeviceWebuiObjectsSecurityProfilesXml_11_0_2 `xml:"security-profiles,omitempty"`
	ServiceGroups         *string                                           `xml:"service-groups,omitempty"`
	Services              *string                                           `xml:"services,omitempty"`
	Tags                  *string                                           `xml:"tags,omitempty"`
	Misc                  []generic.Xml                                     `xml:",any"`
}
type roleDeviceWebuiObjectsCustomObjectsXml_11_0_2 struct {
	DataPatterns  *string       `xml:"data-patterns,omitempty"`
	Spyware       *string       `xml:"spyware,omitempty"`
	UrlCategory   *string       `xml:"url-category,omitempty"`
	Vulnerability *string       `xml:"vulnerability,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type roleDeviceWebuiObjectsDecryptionXml_11_0_2 struct {
	DecryptionProfile *string       `xml:"decryption-profile,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type roleDeviceWebuiObjectsGlobalProtectXml_11_0_2 struct {
	HipObjects  *string       `xml:"hip-objects,omitempty"`
	HipProfiles *string       `xml:"hip-profiles,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type roleDeviceWebuiObjectsSdwanXml_11_0_2 struct {
	SdwanDistProfile            *string       `xml:"sdwan-dist-profile,omitempty"`
	SdwanErrorCorrectionProfile *string       `xml:"sdwan-error-correction-profile,omitempty"`
	SdwanProfile                *string       `xml:"sdwan-profile,omitempty"`
	SdwanSaasQualityProfile     *string       `xml:"sdwan-saas-quality-profile,omitempty"`
	Misc                        []generic.Xml `xml:",any"`
}
type roleDeviceWebuiObjectsSecurityProfilesXml_11_0_2 struct {
	AntiSpyware             *string       `xml:"anti-spyware,omitempty"`
	Antivirus               *string       `xml:"antivirus,omitempty"`
	DataFiltering           *string       `xml:"data-filtering,omitempty"`
	DosProtection           *string       `xml:"dos-protection,omitempty"`
	FileBlocking            *string       `xml:"file-blocking,omitempty"`
	GtpProtection           *string       `xml:"gtp-protection,omitempty"`
	SctpProtection          *string       `xml:"sctp-protection,omitempty"`
	UrlFiltering            *string       `xml:"url-filtering,omitempty"`
	VulnerabilityProtection *string       `xml:"vulnerability-protection,omitempty"`
	WildfireAnalysis        *string       `xml:"wildfire-analysis,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
}
type roleDeviceWebuiOperationsXml_11_0_2 struct {
	DownloadCoreFiles       *string       `xml:"download-core-files,omitempty"`
	DownloadPcapFiles       *string       `xml:"download-pcap-files,omitempty"`
	GenerateStatsDumpFile   *string       `xml:"generate-stats-dump-file,omitempty"`
	GenerateTechSupportFile *string       `xml:"generate-tech-support-file,omitempty"`
	Reboot                  *string       `xml:"reboot,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
}
type roleDeviceWebuiPoliciesXml_11_0_2 struct {
	ApplicationOverrideRulebase *string       `xml:"application-override-rulebase,omitempty"`
	AuthenticationRulebase      *string       `xml:"authentication-rulebase,omitempty"`
	DosRulebase                 *string       `xml:"dos-rulebase,omitempty"`
	NatRulebase                 *string       `xml:"nat-rulebase,omitempty"`
	NetworkPacketBrokerRulebase *string       `xml:"network-packet-broker-rulebase,omitempty"`
	PbfRulebase                 *string       `xml:"pbf-rulebase,omitempty"`
	QosRulebase                 *string       `xml:"qos-rulebase,omitempty"`
	RuleHitCountReset           *string       `xml:"rule-hit-count-reset,omitempty"`
	SdwanRulebase               *string       `xml:"sdwan-rulebase,omitempty"`
	SecurityRulebase            *string       `xml:"security-rulebase,omitempty"`
	SslDecryptionRulebase       *string       `xml:"ssl-decryption-rulebase,omitempty"`
	TunnelInspectRulebase       *string       `xml:"tunnel-inspect-rulebase,omitempty"`
	Misc                        []generic.Xml `xml:",any"`
}
type roleDeviceWebuiPrivacyXml_11_0_2 struct {
	ShowFullIpAddresses           *string       `xml:"show-full-ip-addresses,omitempty"`
	ShowUserNamesInLogsAndReports *string       `xml:"show-user-names-in-logs-and-reports,omitempty"`
	ViewPcapFiles                 *string       `xml:"view-pcap-files,omitempty"`
	Misc                          []generic.Xml `xml:",any"`
}
type roleDeviceWebuiSaveXml_11_0_2 struct {
	ObjectLevelChanges *string       `xml:"object-level-changes,omitempty"`
	PartialSave        *string       `xml:"partial-save,omitempty"`
	SaveForOtherAdmins *string       `xml:"save-for-other-admins,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type roleDeviceXmlapiXml_11_0_2 struct {
	Commit *string       `xml:"commit,omitempty"`
	Config *string       `xml:"config,omitempty"`
	Export *string       `xml:"export,omitempty"`
	Import *string       `xml:"import,omitempty"`
	Iot    *string       `xml:"iot,omitempty"`
	Log    *string       `xml:"log,omitempty"`
	Op     *string       `xml:"op,omitempty"`
	Report *string       `xml:"report,omitempty"`
	UserId *string       `xml:"user-id,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type roleVsysXml_11_0_2 struct {
	Cli     *string                    `xml:"cli,omitempty"`
	Restapi *roleVsysRestapiXml_11_0_2 `xml:"restapi,omitempty"`
	Webui   *roleVsysWebuiXml_11_0_2   `xml:"webui,omitempty"`
	Xmlapi  *roleVsysXmlapiXml_11_0_2  `xml:"xmlapi,omitempty"`
	Misc    []generic.Xml              `xml:",any"`
}
type roleVsysRestapiXml_11_0_2 struct {
	Device   *roleVsysRestapiDeviceXml_11_0_2   `xml:"device,omitempty"`
	Network  *roleVsysRestapiNetworkXml_11_0_2  `xml:"network,omitempty"`
	Objects  *roleVsysRestapiObjectsXml_11_0_2  `xml:"objects,omitempty"`
	Policies *roleVsysRestapiPoliciesXml_11_0_2 `xml:"policies,omitempty"`
	System   *roleVsysRestapiSystemXml_11_0_2   `xml:"system,omitempty"`
	Misc     []generic.Xml                      `xml:",any"`
}
type roleVsysRestapiDeviceXml_11_0_2 struct {
	EmailServerProfiles    *string       `xml:"email-server-profiles,omitempty"`
	HttpServerProfiles     *string       `xml:"http-server-profiles,omitempty"`
	LdapServerProfiles     *string       `xml:"ldap-server-profiles,omitempty"`
	LogInterfaceSetting    *string       `xml:"log-interface-setting,omitempty"`
	SnmpTrapServerProfiles *string       `xml:"snmp-trap-server-profiles,omitempty"`
	SyslogServerProfiles   *string       `xml:"syslog-server-profiles,omitempty"`
	VirtualSystems         *string       `xml:"virtual-systems,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type roleVsysRestapiNetworkXml_11_0_2 struct {
	GlobalprotectClientlessAppGroups *string       `xml:"globalprotect-clientless-app-groups,omitempty"`
	GlobalprotectClientlessApps      *string       `xml:"globalprotect-clientless-apps,omitempty"`
	GlobalprotectGateways            *string       `xml:"globalprotect-gateways,omitempty"`
	GlobalprotectMdmServers          *string       `xml:"globalprotect-mdm-servers,omitempty"`
	GlobalprotectPortals             *string       `xml:"globalprotect-portals,omitempty"`
	SdwanInterfaceProfiles           *string       `xml:"sdwan-interface-profiles,omitempty"`
	Zones                            *string       `xml:"zones,omitempty"`
	Misc                             []generic.Xml `xml:",any"`
}
type roleVsysRestapiObjectsXml_11_0_2 struct {
	AddressGroups                           *string       `xml:"address-groups,omitempty"`
	Addresses                               *string       `xml:"addresses,omitempty"`
	AntiSpywareSecurityProfiles             *string       `xml:"anti-spyware-security-profiles,omitempty"`
	AntivirusSecurityProfiles               *string       `xml:"antivirus-security-profiles,omitempty"`
	ApplicationFilters                      *string       `xml:"application-filters,omitempty"`
	ApplicationGroups                       *string       `xml:"application-groups,omitempty"`
	Applications                            *string       `xml:"applications,omitempty"`
	AuthenticationEnforcements              *string       `xml:"authentication-enforcements,omitempty"`
	CustomDataPatterns                      *string       `xml:"custom-data-patterns,omitempty"`
	CustomSpywareSignatures                 *string       `xml:"custom-spyware-signatures,omitempty"`
	CustomUrlCategories                     *string       `xml:"custom-url-categories,omitempty"`
	CustomVulnerabilitySignatures           *string       `xml:"custom-vulnerability-signatures,omitempty"`
	DataFilteringSecurityProfiles           *string       `xml:"data-filtering-security-profiles,omitempty"`
	DecryptionProfiles                      *string       `xml:"decryption-profiles,omitempty"`
	Devices                                 *string       `xml:"devices,omitempty"`
	DosProtectionSecurityProfiles           *string       `xml:"dos-protection-security-profiles,omitempty"`
	DynamicUserGroups                       *string       `xml:"dynamic-user-groups,omitempty"`
	ExternalDynamicLists                    *string       `xml:"external-dynamic-lists,omitempty"`
	FileBlockingSecurityProfiles            *string       `xml:"file-blocking-security-profiles,omitempty"`
	GlobalprotectHipObjects                 *string       `xml:"globalprotect-hip-objects,omitempty"`
	GlobalprotectHipProfiles                *string       `xml:"globalprotect-hip-profiles,omitempty"`
	GtpProtectionSecurityProfiles           *string       `xml:"gtp-protection-security-profiles,omitempty"`
	LogForwardingProfiles                   *string       `xml:"log-forwarding-profiles,omitempty"`
	PacketBrokerProfiles                    *string       `xml:"packet-broker-profiles,omitempty"`
	Regions                                 *string       `xml:"regions,omitempty"`
	Schedules                               *string       `xml:"schedules,omitempty"`
	SctpProtectionSecurityProfiles          *string       `xml:"sctp-protection-security-profiles,omitempty"`
	SdwanErrorCorrectionProfiles            *string       `xml:"sdwan-error-correction-profiles,omitempty"`
	SdwanPathQualityProfiles                *string       `xml:"sdwan-path-quality-profiles,omitempty"`
	SdwanSaasQualityProfiles                *string       `xml:"sdwan-saas-quality-profiles,omitempty"`
	SdwanTrafficDistributionProfiles        *string       `xml:"sdwan-traffic-distribution-profiles,omitempty"`
	SecurityProfileGroups                   *string       `xml:"security-profile-groups,omitempty"`
	ServiceGroups                           *string       `xml:"service-groups,omitempty"`
	Services                                *string       `xml:"services,omitempty"`
	Tags                                    *string       `xml:"tags,omitempty"`
	UrlFilteringSecurityProfiles            *string       `xml:"url-filtering-security-profiles,omitempty"`
	VulnerabilityProtectionSecurityProfiles *string       `xml:"vulnerability-protection-security-profiles,omitempty"`
	WildfireAnalysisSecurityProfiles        *string       `xml:"wildfire-analysis-security-profiles,omitempty"`
	Misc                                    []generic.Xml `xml:",any"`
}
type roleVsysRestapiPoliciesXml_11_0_2 struct {
	ApplicationOverrideRules   *string       `xml:"application-override-rules,omitempty"`
	AuthenticationRules        *string       `xml:"authentication-rules,omitempty"`
	DecryptionRules            *string       `xml:"decryption-rules,omitempty"`
	DosRules                   *string       `xml:"dos-rules,omitempty"`
	NatRules                   *string       `xml:"nat-rules,omitempty"`
	NetworkPacketBrokerRules   *string       `xml:"network-packet-broker-rules,omitempty"`
	PolicyBasedForwardingRules *string       `xml:"policy-based-forwarding-rules,omitempty"`
	QosRules                   *string       `xml:"qos-rules,omitempty"`
	SdwanRules                 *string       `xml:"sdwan-rules,omitempty"`
	SecurityRules              *string       `xml:"security-rules,omitempty"`
	TunnelInspectionRules      *string       `xml:"tunnel-inspection-rules,omitempty"`
	Misc                       []generic.Xml `xml:",any"`
}
type roleVsysRestapiSystemXml_11_0_2 struct {
	Configuration *string       `xml:"configuration,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type roleVsysWebuiXml_11_0_2 struct {
	Acc        *string                            `xml:"acc,omitempty"`
	Commit     *roleVsysWebuiCommitXml_11_0_2     `xml:"commit,omitempty"`
	Dashboard  *string                            `xml:"dashboard,omitempty"`
	Device     *roleVsysWebuiDeviceXml_11_0_2     `xml:"device,omitempty"`
	Monitor    *roleVsysWebuiMonitorXml_11_0_2    `xml:"monitor,omitempty"`
	Network    *roleVsysWebuiNetworkXml_11_0_2    `xml:"network,omitempty"`
	Objects    *roleVsysWebuiObjectsXml_11_0_2    `xml:"objects,omitempty"`
	Operations *roleVsysWebuiOperationsXml_11_0_2 `xml:"operations,omitempty"`
	Policies   *roleVsysWebuiPoliciesXml_11_0_2   `xml:"policies,omitempty"`
	Privacy    *roleVsysWebuiPrivacyXml_11_0_2    `xml:"privacy,omitempty"`
	Save       *roleVsysWebuiSaveXml_11_0_2       `xml:"save,omitempty"`
	Tasks      *string                            `xml:"tasks,omitempty"`
	Validate   *string                            `xml:"validate,omitempty"`
	Misc       []generic.Xml                      `xml:",any"`
}
type roleVsysWebuiCommitXml_11_0_2 struct {
	CommitForOtherAdmins *string       `xml:"commit-for-other-admins,omitempty"`
	VirtualSystems       *string       `xml:"virtual-systems,omitempty"`
	Misc                 []generic.Xml `xml:",any"`
}
type roleVsysWebuiDeviceXml_11_0_2 struct {
	Administrators         *string                                             `xml:"administrators,omitempty"`
	AuthenticationProfile  *string                                             `xml:"authentication-profile,omitempty"`
	AuthenticationSequence *string                                             `xml:"authentication-sequence,omitempty"`
	BlockPages             *string                                             `xml:"block-pages,omitempty"`
	CertificateManagement  *roleVsysWebuiDeviceCertificateManagementXml_11_0_2 `xml:"certificate-management,omitempty"`
	DataRedistribution     *string                                             `xml:"data-redistribution,omitempty"`
	DeviceQuarantine       *string                                             `xml:"device-quarantine,omitempty"`
	DhcpSyslogServer       *string                                             `xml:"dhcp-syslog-server,omitempty"`
	LocalUserDatabase      *roleVsysWebuiDeviceLocalUserDatabaseXml_11_0_2     `xml:"local-user-database,omitempty"`
	LogSettings            *roleVsysWebuiDeviceLogSettingsXml_11_0_2           `xml:"log-settings,omitempty"`
	PolicyRecommendations  *roleVsysWebuiDevicePolicyRecommendationsXml_11_0_2 `xml:"policy-recommendations,omitempty"`
	ServerProfile          *roleVsysWebuiDeviceServerProfileXml_11_0_2         `xml:"server-profile,omitempty"`
	Setup                  *roleVsysWebuiDeviceSetupXml_11_0_2                 `xml:"setup,omitempty"`
	Troubleshooting        *string                                             `xml:"troubleshooting,omitempty"`
	UserIdentification     *string                                             `xml:"user-identification,omitempty"`
	VmInfoSource           *string                                             `xml:"vm-info-source,omitempty"`
	Misc                   []generic.Xml                                       `xml:",any"`
}
type roleVsysWebuiDeviceCertificateManagementXml_11_0_2 struct {
	CertificateProfile     *string       `xml:"certificate-profile,omitempty"`
	Certificates           *string       `xml:"certificates,omitempty"`
	OcspResponder          *string       `xml:"ocsp-responder,omitempty"`
	Scep                   *string       `xml:"scep,omitempty"`
	SshServiceProfile      *string       `xml:"ssh-service-profile,omitempty"`
	SslDecryptionExclusion *string       `xml:"ssl-decryption-exclusion,omitempty"`
	SslTlsServiceProfile   *string       `xml:"ssl-tls-service-profile,omitempty"`
	Misc                   []generic.Xml `xml:",any"`
}
type roleVsysWebuiDeviceLocalUserDatabaseXml_11_0_2 struct {
	UserGroups *string       `xml:"user-groups,omitempty"`
	Users      *string       `xml:"users,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type roleVsysWebuiDeviceLogSettingsXml_11_0_2 struct {
	Config        *string       `xml:"config,omitempty"`
	Correlation   *string       `xml:"correlation,omitempty"`
	Globalprotect *string       `xml:"globalprotect,omitempty"`
	Hipmatch      *string       `xml:"hipmatch,omitempty"`
	Iptag         *string       `xml:"iptag,omitempty"`
	System        *string       `xml:"system,omitempty"`
	UserId        *string       `xml:"user-id,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type roleVsysWebuiDevicePolicyRecommendationsXml_11_0_2 struct {
	Iot  *string       `xml:"iot,omitempty"`
	Saas *string       `xml:"saas,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type roleVsysWebuiDeviceServerProfileXml_11_0_2 struct {
	Dns      *string       `xml:"dns,omitempty"`
	Email    *string       `xml:"email,omitempty"`
	Http     *string       `xml:"http,omitempty"`
	Kerberos *string       `xml:"kerberos,omitempty"`
	Ldap     *string       `xml:"ldap,omitempty"`
	Mfa      *string       `xml:"mfa,omitempty"`
	Netflow  *string       `xml:"netflow,omitempty"`
	Radius   *string       `xml:"radius,omitempty"`
	SamlIdp  *string       `xml:"saml_idp,omitempty"`
	Scp      *string       `xml:"scp,omitempty"`
	SnmpTrap *string       `xml:"snmp-trap,omitempty"`
	Syslog   *string       `xml:"syslog,omitempty"`
	Tacplus  *string       `xml:"tacplus,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type roleVsysWebuiDeviceSetupXml_11_0_2 struct {
	ContentId  *string       `xml:"content-id,omitempty"`
	Hsm        *string       `xml:"hsm,omitempty"`
	Interfaces *string       `xml:"interfaces,omitempty"`
	Management *string       `xml:"management,omitempty"`
	Operations *string       `xml:"operations,omitempty"`
	Services   *string       `xml:"services,omitempty"`
	Session    *string       `xml:"session,omitempty"`
	Telemetry  *string       `xml:"telemetry,omitempty"`
	Wildfire   *string       `xml:"wildfire,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type roleVsysWebuiMonitorXml_11_0_2 struct {
	AppScope                   *string                                                   `xml:"app-scope,omitempty"`
	AutomatedCorrelationEngine *roleVsysWebuiMonitorAutomatedCorrelationEngineXml_11_0_2 `xml:"automated-correlation-engine,omitempty"`
	BlockIpList                *string                                                   `xml:"block-ip-list,omitempty"`
	CustomReports              *roleVsysWebuiMonitorCustomReportsXml_11_0_2              `xml:"custom-reports,omitempty"`
	ExternalLogs               *string                                                   `xml:"external-logs,omitempty"`
	Logs                       *roleVsysWebuiMonitorLogsXml_11_0_2                       `xml:"logs,omitempty"`
	PdfReports                 *roleVsysWebuiMonitorPdfReportsXml_11_0_2                 `xml:"pdf-reports,omitempty"`
	SessionBrowser             *string                                                   `xml:"session-browser,omitempty"`
	ViewCustomReports          *string                                                   `xml:"view-custom-reports,omitempty"`
	Misc                       []generic.Xml                                             `xml:",any"`
}
type roleVsysWebuiMonitorAutomatedCorrelationEngineXml_11_0_2 struct {
	CorrelatedEvents   *string       `xml:"correlated-events,omitempty"`
	CorrelationObjects *string       `xml:"correlation-objects,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type roleVsysWebuiMonitorCustomReportsXml_11_0_2 struct {
	ApplicationStatistics *string       `xml:"application-statistics,omitempty"`
	Auth                  *string       `xml:"auth,omitempty"`
	DataFilteringLog      *string       `xml:"data-filtering-log,omitempty"`
	DecryptionLog         *string       `xml:"decryption-log,omitempty"`
	DecryptionSummary     *string       `xml:"decryption-summary,omitempty"`
	Globalprotect         *string       `xml:"globalprotect,omitempty"`
	GtpLog                *string       `xml:"gtp-log,omitempty"`
	GtpSummary            *string       `xml:"gtp-summary,omitempty"`
	Hipmatch              *string       `xml:"hipmatch,omitempty"`
	Iptag                 *string       `xml:"iptag,omitempty"`
	SctpLog               *string       `xml:"sctp-log,omitempty"`
	SctpSummary           *string       `xml:"sctp-summary,omitempty"`
	ThreatLog             *string       `xml:"threat-log,omitempty"`
	ThreatSummary         *string       `xml:"threat-summary,omitempty"`
	TrafficLog            *string       `xml:"traffic-log,omitempty"`
	TrafficSummary        *string       `xml:"traffic-summary,omitempty"`
	TunnelLog             *string       `xml:"tunnel-log,omitempty"`
	TunnelSummary         *string       `xml:"tunnel-summary,omitempty"`
	UrlLog                *string       `xml:"url-log,omitempty"`
	UrlSummary            *string       `xml:"url-summary,omitempty"`
	Userid                *string       `xml:"userid,omitempty"`
	WildfireLog           *string       `xml:"wildfire-log,omitempty"`
	Misc                  []generic.Xml `xml:",any"`
}
type roleVsysWebuiMonitorLogsXml_11_0_2 struct {
	Authentication *string       `xml:"authentication,omitempty"`
	DataFiltering  *string       `xml:"data-filtering,omitempty"`
	Decryption     *string       `xml:"decryption,omitempty"`
	Globalprotect  *string       `xml:"globalprotect,omitempty"`
	Gtp            *string       `xml:"gtp,omitempty"`
	Hipmatch       *string       `xml:"hipmatch,omitempty"`
	Iptag          *string       `xml:"iptag,omitempty"`
	Sctp           *string       `xml:"sctp,omitempty"`
	Threat         *string       `xml:"threat,omitempty"`
	Traffic        *string       `xml:"traffic,omitempty"`
	Tunnel         *string       `xml:"tunnel,omitempty"`
	Url            *string       `xml:"url,omitempty"`
	Userid         *string       `xml:"userid,omitempty"`
	Wildfire       *string       `xml:"wildfire,omitempty"`
	Misc           []generic.Xml `xml:",any"`
}
type roleVsysWebuiMonitorPdfReportsXml_11_0_2 struct {
	EmailScheduler             *string       `xml:"email-scheduler,omitempty"`
	ManagePdfSummary           *string       `xml:"manage-pdf-summary,omitempty"`
	PdfSummaryReports          *string       `xml:"pdf-summary-reports,omitempty"`
	ReportGroups               *string       `xml:"report-groups,omitempty"`
	SaasApplicationUsageReport *string       `xml:"saas-application-usage-report,omitempty"`
	UserActivityReport         *string       `xml:"user-activity-report,omitempty"`
	Misc                       []generic.Xml `xml:",any"`
}
type roleVsysWebuiNetworkXml_11_0_2 struct {
	GlobalProtect         *roleVsysWebuiNetworkGlobalProtectXml_11_0_2 `xml:"global-protect,omitempty"`
	SdwanInterfaceProfile *string                                      `xml:"sdwan-interface-profile,omitempty"`
	Zones                 *string                                      `xml:"zones,omitempty"`
	Misc                  []generic.Xml                                `xml:",any"`
}
type roleVsysWebuiNetworkGlobalProtectXml_11_0_2 struct {
	ClientlessAppGroups *string       `xml:"clientless-app-groups,omitempty"`
	ClientlessApps      *string       `xml:"clientless-apps,omitempty"`
	Gateways            *string       `xml:"gateways,omitempty"`
	Mdm                 *string       `xml:"mdm,omitempty"`
	Portals             *string       `xml:"portals,omitempty"`
	Misc                []generic.Xml `xml:",any"`
}
type roleVsysWebuiObjectsXml_11_0_2 struct {
	AddressGroups         *string                                         `xml:"address-groups,omitempty"`
	Addresses             *string                                         `xml:"addresses,omitempty"`
	ApplicationFilters    *string                                         `xml:"application-filters,omitempty"`
	ApplicationGroups     *string                                         `xml:"application-groups,omitempty"`
	Applications          *string                                         `xml:"applications,omitempty"`
	Authentication        *string                                         `xml:"authentication,omitempty"`
	CustomObjects         *roleVsysWebuiObjectsCustomObjectsXml_11_0_2    `xml:"custom-objects,omitempty"`
	Decryption            *roleVsysWebuiObjectsDecryptionXml_11_0_2       `xml:"decryption,omitempty"`
	Devices               *string                                         `xml:"devices,omitempty"`
	DynamicBlockLists     *string                                         `xml:"dynamic-block-lists,omitempty"`
	DynamicUserGroups     *string                                         `xml:"dynamic-user-groups,omitempty"`
	GlobalProtect         *roleVsysWebuiObjectsGlobalProtectXml_11_0_2    `xml:"global-protect,omitempty"`
	LogForwarding         *string                                         `xml:"log-forwarding,omitempty"`
	PacketBrokerProfile   *string                                         `xml:"packet-broker-profile,omitempty"`
	Regions               *string                                         `xml:"regions,omitempty"`
	Schedules             *string                                         `xml:"schedules,omitempty"`
	Sdwan                 *roleVsysWebuiObjectsSdwanXml_11_0_2            `xml:"sdwan,omitempty"`
	SecurityProfileGroups *string                                         `xml:"security-profile-groups,omitempty"`
	SecurityProfiles      *roleVsysWebuiObjectsSecurityProfilesXml_11_0_2 `xml:"security-profiles,omitempty"`
	ServiceGroups         *string                                         `xml:"service-groups,omitempty"`
	Services              *string                                         `xml:"services,omitempty"`
	Tags                  *string                                         `xml:"tags,omitempty"`
	Misc                  []generic.Xml                                   `xml:",any"`
}
type roleVsysWebuiObjectsCustomObjectsXml_11_0_2 struct {
	DataPatterns  *string       `xml:"data-patterns,omitempty"`
	Spyware       *string       `xml:"spyware,omitempty"`
	UrlCategory   *string       `xml:"url-category,omitempty"`
	Vulnerability *string       `xml:"vulnerability,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type roleVsysWebuiObjectsDecryptionXml_11_0_2 struct {
	DecryptionProfile *string       `xml:"decryption-profile,omitempty"`
	Misc              []generic.Xml `xml:",any"`
}
type roleVsysWebuiObjectsGlobalProtectXml_11_0_2 struct {
	HipObjects  *string       `xml:"hip-objects,omitempty"`
	HipProfiles *string       `xml:"hip-profiles,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type roleVsysWebuiObjectsSdwanXml_11_0_2 struct {
	SdwanDistProfile            *string       `xml:"sdwan-dist-profile,omitempty"`
	SdwanErrorCorrectionProfile *string       `xml:"sdwan-error-correction-profile,omitempty"`
	SdwanProfile                *string       `xml:"sdwan-profile,omitempty"`
	SdwanSaasQualityProfile     *string       `xml:"sdwan-saas-quality-profile,omitempty"`
	Misc                        []generic.Xml `xml:",any"`
}
type roleVsysWebuiObjectsSecurityProfilesXml_11_0_2 struct {
	AntiSpyware             *string       `xml:"anti-spyware,omitempty"`
	Antivirus               *string       `xml:"antivirus,omitempty"`
	DataFiltering           *string       `xml:"data-filtering,omitempty"`
	DosProtection           *string       `xml:"dos-protection,omitempty"`
	FileBlocking            *string       `xml:"file-blocking,omitempty"`
	GtpProtection           *string       `xml:"gtp-protection,omitempty"`
	SctpProtection          *string       `xml:"sctp-protection,omitempty"`
	UrlFiltering            *string       `xml:"url-filtering,omitempty"`
	VulnerabilityProtection *string       `xml:"vulnerability-protection,omitempty"`
	WildfireAnalysis        *string       `xml:"wildfire-analysis,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
}
type roleVsysWebuiOperationsXml_11_0_2 struct {
	DownloadCoreFiles       *string       `xml:"download-core-files,omitempty"`
	DownloadPcapFiles       *string       `xml:"download-pcap-files,omitempty"`
	GenerateStatsDumpFile   *string       `xml:"generate-stats-dump-file,omitempty"`
	GenerateTechSupportFile *string       `xml:"generate-tech-support-file,omitempty"`
	Reboot                  *string       `xml:"reboot,omitempty"`
	Misc                    []generic.Xml `xml:",any"`
}
type roleVsysWebuiPoliciesXml_11_0_2 struct {
	ApplicationOverrideRulebase *string       `xml:"application-override-rulebase,omitempty"`
	AuthenticationRulebase      *string       `xml:"authentication-rulebase,omitempty"`
	DosRulebase                 *string       `xml:"dos-rulebase,omitempty"`
	NatRulebase                 *string       `xml:"nat-rulebase,omitempty"`
	NetworkPacketBrokerRulebase *string       `xml:"network-packet-broker-rulebase,omitempty"`
	PbfRulebase                 *string       `xml:"pbf-rulebase,omitempty"`
	QosRulebase                 *string       `xml:"qos-rulebase,omitempty"`
	RuleHitCountReset           *string       `xml:"rule-hit-count-reset,omitempty"`
	SdwanRulebase               *string       `xml:"sdwan-rulebase,omitempty"`
	SecurityRulebase            *string       `xml:"security-rulebase,omitempty"`
	SslDecryptionRulebase       *string       `xml:"ssl-decryption-rulebase,omitempty"`
	TunnelInspectRulebase       *string       `xml:"tunnel-inspect-rulebase,omitempty"`
	Misc                        []generic.Xml `xml:",any"`
}
type roleVsysWebuiPrivacyXml_11_0_2 struct {
	ShowFullIpAddresses           *string       `xml:"show-full-ip-addresses,omitempty"`
	ShowUserNamesInLogsAndReports *string       `xml:"show-user-names-in-logs-and-reports,omitempty"`
	ViewPcapFiles                 *string       `xml:"view-pcap-files,omitempty"`
	Misc                          []generic.Xml `xml:",any"`
}
type roleVsysWebuiSaveXml_11_0_2 struct {
	ObjectLevelChanges *string       `xml:"object-level-changes,omitempty"`
	PartialSave        *string       `xml:"partial-save,omitempty"`
	SaveForOtherAdmins *string       `xml:"save-for-other-admins,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type roleVsysXmlapiXml_11_0_2 struct {
	Commit *string       `xml:"commit,omitempty"`
	Config *string       `xml:"config,omitempty"`
	Export *string       `xml:"export,omitempty"`
	Import *string       `xml:"import,omitempty"`
	Iot    *string       `xml:"iot,omitempty"`
	Log    *string       `xml:"log,omitempty"`
	Op     *string       `xml:"op,omitempty"`
	Report *string       `xml:"report,omitempty"`
	UserId *string       `xml:"user-id,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	if s.Role != nil {
		var obj roleXml
		obj.MarshalFromObject(*s.Role)
		o.Role = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var roleVal *Role
	if o.Role != nil {
		obj, err := o.Role.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		roleVal = obj
	}

	result := &Entry{
		Name:        o.Name,
		Description: o.Description,
		Role:        roleVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *roleXml) MarshalFromObject(s Role) {
	if s.Device != nil {
		var obj roleDeviceXml
		obj.MarshalFromObject(*s.Device)
		o.Device = &obj
	}
	if s.Vsys != nil {
		var obj roleVsysXml
		obj.MarshalFromObject(*s.Vsys)
		o.Vsys = &obj
	}
	o.Misc = s.Misc
}

func (o roleXml) UnmarshalToObject() (*Role, error) {
	var deviceVal *RoleDevice
	if o.Device != nil {
		obj, err := o.Device.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		deviceVal = obj
	}
	var vsysVal *RoleVsys
	if o.Vsys != nil {
		obj, err := o.Vsys.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		vsysVal = obj
	}

	result := &Role{
		Device: deviceVal,
		Vsys:   vsysVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *roleDeviceXml) MarshalFromObject(s RoleDevice) {
	o.Cli = s.Cli
	if s.Restapi != nil {
		var obj roleDeviceRestapiXml
		obj.MarshalFromObject(*s.Restapi)
		o.Restapi = &obj
	}
	if s.Webui != nil {
		var obj roleDeviceWebuiXml
		obj.MarshalFromObject(*s.Webui)
		o.Webui = &obj
	}
	if s.Xmlapi != nil {
		var obj roleDeviceXmlapiXml
		obj.MarshalFromObject(*s.Xmlapi)
		o.Xmlapi = &obj
	}
	o.Misc = s.Misc
}

func (o roleDeviceXml) UnmarshalToObject() (*RoleDevice, error) {
	var restapiVal *RoleDeviceRestapi
	if o.Restapi != nil {
		obj, err := o.Restapi.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		restapiVal = obj
	}
	var webuiVal *RoleDeviceWebui
	if o.Webui != nil {
		obj, err := o.Webui.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		webuiVal = obj
	}
	var xmlapiVal *RoleDeviceXmlapi
	if o.Xmlapi != nil {
		obj, err := o.Xmlapi.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		xmlapiVal = obj
	}

	result := &RoleDevice{
		Cli:     o.Cli,
		Restapi: restapiVal,
		Webui:   webuiVal,
		Xmlapi:  xmlapiVal,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *roleDeviceRestapiXml) MarshalFromObject(s RoleDeviceRestapi) {
	if s.Device != nil {
		var obj roleDeviceRestapiDeviceXml
		obj.MarshalFromObject(*s.Device)
		o.Device = &obj
	}
	if s.Network != nil {
		var obj roleDeviceRestapiNetworkXml
		obj.MarshalFromObject(*s.Network)
		o.Network = &obj
	}
	if s.Objects != nil {
		var obj roleDeviceRestapiObjectsXml
		obj.MarshalFromObject(*s.Objects)
		o.Objects = &obj
	}
	if s.Policies != nil {
		var obj roleDeviceRestapiPoliciesXml
		obj.MarshalFromObject(*s.Policies)
		o.Policies = &obj
	}
	if s.System != nil {
		var obj roleDeviceRestapiSystemXml
		obj.MarshalFromObject(*s.System)
		o.System = &obj
	}
	o.Misc = s.Misc
}

func (o roleDeviceRestapiXml) UnmarshalToObject() (*RoleDeviceRestapi, error) {
	var deviceVal *RoleDeviceRestapiDevice
	if o.Device != nil {
		obj, err := o.Device.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		deviceVal = obj
	}
	var networkVal *RoleDeviceRestapiNetwork
	if o.Network != nil {
		obj, err := o.Network.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		networkVal = obj
	}
	var objectsVal *RoleDeviceRestapiObjects
	if o.Objects != nil {
		obj, err := o.Objects.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		objectsVal = obj
	}
	var policiesVal *RoleDeviceRestapiPolicies
	if o.Policies != nil {
		obj, err := o.Policies.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		policiesVal = obj
	}
	var systemVal *RoleDeviceRestapiSystem
	if o.System != nil {
		obj, err := o.System.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		systemVal = obj
	}

	result := &RoleDeviceRestapi{
		Device:   deviceVal,
		Network:  networkVal,
		Objects:  objectsVal,
		Policies: policiesVal,
		System:   systemVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *roleDeviceRestapiDeviceXml) MarshalFromObject(s RoleDeviceRestapiDevice) {
	o.EmailServerProfiles = s.EmailServerProfiles
	o.HttpServerProfiles = s.HttpServerProfiles
	o.LdapServerProfiles = s.LdapServerProfiles
	o.LogInterfaceSetting = s.LogInterfaceSetting
	o.SnmpTrapServerProfiles = s.SnmpTrapServerProfiles
	o.SyslogServerProfiles = s.SyslogServerProfiles
	o.VirtualSystems = s.VirtualSystems
	o.Misc = s.Misc
}

func (o roleDeviceRestapiDeviceXml) UnmarshalToObject() (*RoleDeviceRestapiDevice, error) {

	result := &RoleDeviceRestapiDevice{
		EmailServerProfiles:    o.EmailServerProfiles,
		HttpServerProfiles:     o.HttpServerProfiles,
		LdapServerProfiles:     o.LdapServerProfiles,
		LogInterfaceSetting:    o.LogInterfaceSetting,
		SnmpTrapServerProfiles: o.SnmpTrapServerProfiles,
		SyslogServerProfiles:   o.SyslogServerProfiles,
		VirtualSystems:         o.VirtualSystems,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *roleDeviceRestapiNetworkXml) MarshalFromObject(s RoleDeviceRestapiNetwork) {
	o.AggregateEthernetInterfaces = s.AggregateEthernetInterfaces
	o.BfdNetworkProfiles = s.BfdNetworkProfiles
	o.BgpRoutingProfiles = s.BgpRoutingProfiles
	o.DhcpRelays = s.DhcpRelays
	o.DhcpServers = s.DhcpServers
	o.DnsProxies = s.DnsProxies
	o.EthernetInterfaces = s.EthernetInterfaces
	o.GlobalprotectClientlessAppGroups = s.GlobalprotectClientlessAppGroups
	o.GlobalprotectClientlessApps = s.GlobalprotectClientlessApps
	o.GlobalprotectGateways = s.GlobalprotectGateways
	o.GlobalprotectIpsecCryptoNetworkProfiles = s.GlobalprotectIpsecCryptoNetworkProfiles
	o.GlobalprotectMdmServers = s.GlobalprotectMdmServers
	o.GlobalprotectPortals = s.GlobalprotectPortals
	o.GreTunnels = s.GreTunnels
	o.IkeCryptoNetworkProfiles = s.IkeCryptoNetworkProfiles
	o.IkeGatewayNetworkProfiles = s.IkeGatewayNetworkProfiles
	o.InterfaceManagementNetworkProfiles = s.InterfaceManagementNetworkProfiles
	o.IpsecCryptoNetworkProfiles = s.IpsecCryptoNetworkProfiles
	o.IpsecTunnels = s.IpsecTunnels
	o.Lldp = s.Lldp
	o.LldpNetworkProfiles = s.LldpNetworkProfiles
	o.LogicalRouters = s.LogicalRouters
	o.LoopbackInterfaces = s.LoopbackInterfaces
	o.QosInterfaces = s.QosInterfaces
	o.QosNetworkProfiles = s.QosNetworkProfiles
	o.SdwanInterfaceProfiles = s.SdwanInterfaceProfiles
	o.SdwanInterfaces = s.SdwanInterfaces
	o.TunnelInterfaces = s.TunnelInterfaces
	o.TunnelMonitorNetworkProfiles = s.TunnelMonitorNetworkProfiles
	o.VirtualRouters = s.VirtualRouters
	o.VirtualWires = s.VirtualWires
	o.VlanInterfaces = s.VlanInterfaces
	o.Vlans = s.Vlans
	o.ZoneProtectionNetworkProfiles = s.ZoneProtectionNetworkProfiles
	o.Zones = s.Zones
	o.Misc = s.Misc
}

func (o roleDeviceRestapiNetworkXml) UnmarshalToObject() (*RoleDeviceRestapiNetwork, error) {

	result := &RoleDeviceRestapiNetwork{
		AggregateEthernetInterfaces:             o.AggregateEthernetInterfaces,
		BfdNetworkProfiles:                      o.BfdNetworkProfiles,
		BgpRoutingProfiles:                      o.BgpRoutingProfiles,
		DhcpRelays:                              o.DhcpRelays,
		DhcpServers:                             o.DhcpServers,
		DnsProxies:                              o.DnsProxies,
		EthernetInterfaces:                      o.EthernetInterfaces,
		GlobalprotectClientlessAppGroups:        o.GlobalprotectClientlessAppGroups,
		GlobalprotectClientlessApps:             o.GlobalprotectClientlessApps,
		GlobalprotectGateways:                   o.GlobalprotectGateways,
		GlobalprotectIpsecCryptoNetworkProfiles: o.GlobalprotectIpsecCryptoNetworkProfiles,
		GlobalprotectMdmServers:                 o.GlobalprotectMdmServers,
		GlobalprotectPortals:                    o.GlobalprotectPortals,
		GreTunnels:                              o.GreTunnels,
		IkeCryptoNetworkProfiles:                o.IkeCryptoNetworkProfiles,
		IkeGatewayNetworkProfiles:               o.IkeGatewayNetworkProfiles,
		InterfaceManagementNetworkProfiles:      o.InterfaceManagementNetworkProfiles,
		IpsecCryptoNetworkProfiles:              o.IpsecCryptoNetworkProfiles,
		IpsecTunnels:                            o.IpsecTunnels,
		Lldp:                                    o.Lldp,
		LldpNetworkProfiles:                     o.LldpNetworkProfiles,
		LogicalRouters:                          o.LogicalRouters,
		LoopbackInterfaces:                      o.LoopbackInterfaces,
		QosInterfaces:                           o.QosInterfaces,
		QosNetworkProfiles:                      o.QosNetworkProfiles,
		SdwanInterfaceProfiles:                  o.SdwanInterfaceProfiles,
		SdwanInterfaces:                         o.SdwanInterfaces,
		TunnelInterfaces:                        o.TunnelInterfaces,
		TunnelMonitorNetworkProfiles:            o.TunnelMonitorNetworkProfiles,
		VirtualRouters:                          o.VirtualRouters,
		VirtualWires:                            o.VirtualWires,
		VlanInterfaces:                          o.VlanInterfaces,
		Vlans:                                   o.Vlans,
		ZoneProtectionNetworkProfiles:           o.ZoneProtectionNetworkProfiles,
		Zones:                                   o.Zones,
		Misc:                                    o.Misc,
	}
	return result, nil
}
func (o *roleDeviceRestapiObjectsXml) MarshalFromObject(s RoleDeviceRestapiObjects) {
	o.AddressGroups = s.AddressGroups
	o.Addresses = s.Addresses
	o.AntiSpywareSecurityProfiles = s.AntiSpywareSecurityProfiles
	o.AntivirusSecurityProfiles = s.AntivirusSecurityProfiles
	o.ApplicationFilters = s.ApplicationFilters
	o.ApplicationGroups = s.ApplicationGroups
	o.Applications = s.Applications
	o.AuthenticationEnforcements = s.AuthenticationEnforcements
	o.CustomDataPatterns = s.CustomDataPatterns
	o.CustomSpywareSignatures = s.CustomSpywareSignatures
	o.CustomUrlCategories = s.CustomUrlCategories
	o.CustomVulnerabilitySignatures = s.CustomVulnerabilitySignatures
	o.DataFilteringSecurityProfiles = s.DataFilteringSecurityProfiles
	o.DecryptionProfiles = s.DecryptionProfiles
	o.Devices = s.Devices
	o.DosProtectionSecurityProfiles = s.DosProtectionSecurityProfiles
	o.DynamicUserGroups = s.DynamicUserGroups
	o.ExternalDynamicLists = s.ExternalDynamicLists
	o.FileBlockingSecurityProfiles = s.FileBlockingSecurityProfiles
	o.GlobalprotectHipObjects = s.GlobalprotectHipObjects
	o.GlobalprotectHipProfiles = s.GlobalprotectHipProfiles
	o.GtpProtectionSecurityProfiles = s.GtpProtectionSecurityProfiles
	o.LogForwardingProfiles = s.LogForwardingProfiles
	o.PacketBrokerProfiles = s.PacketBrokerProfiles
	o.Regions = s.Regions
	o.Schedules = s.Schedules
	o.SctpProtectionSecurityProfiles = s.SctpProtectionSecurityProfiles
	o.SdwanErrorCorrectionProfiles = s.SdwanErrorCorrectionProfiles
	o.SdwanPathQualityProfiles = s.SdwanPathQualityProfiles
	o.SdwanSaasQualityProfiles = s.SdwanSaasQualityProfiles
	o.SdwanTrafficDistributionProfiles = s.SdwanTrafficDistributionProfiles
	o.SecurityProfileGroups = s.SecurityProfileGroups
	o.ServiceGroups = s.ServiceGroups
	o.Services = s.Services
	o.Tags = s.Tags
	o.UrlFilteringSecurityProfiles = s.UrlFilteringSecurityProfiles
	o.VulnerabilityProtectionSecurityProfiles = s.VulnerabilityProtectionSecurityProfiles
	o.WildfireAnalysisSecurityProfiles = s.WildfireAnalysisSecurityProfiles
	o.Misc = s.Misc
}

func (o roleDeviceRestapiObjectsXml) UnmarshalToObject() (*RoleDeviceRestapiObjects, error) {

	result := &RoleDeviceRestapiObjects{
		AddressGroups:                           o.AddressGroups,
		Addresses:                               o.Addresses,
		AntiSpywareSecurityProfiles:             o.AntiSpywareSecurityProfiles,
		AntivirusSecurityProfiles:               o.AntivirusSecurityProfiles,
		ApplicationFilters:                      o.ApplicationFilters,
		ApplicationGroups:                       o.ApplicationGroups,
		Applications:                            o.Applications,
		AuthenticationEnforcements:              o.AuthenticationEnforcements,
		CustomDataPatterns:                      o.CustomDataPatterns,
		CustomSpywareSignatures:                 o.CustomSpywareSignatures,
		CustomUrlCategories:                     o.CustomUrlCategories,
		CustomVulnerabilitySignatures:           o.CustomVulnerabilitySignatures,
		DataFilteringSecurityProfiles:           o.DataFilteringSecurityProfiles,
		DecryptionProfiles:                      o.DecryptionProfiles,
		Devices:                                 o.Devices,
		DosProtectionSecurityProfiles:           o.DosProtectionSecurityProfiles,
		DynamicUserGroups:                       o.DynamicUserGroups,
		ExternalDynamicLists:                    o.ExternalDynamicLists,
		FileBlockingSecurityProfiles:            o.FileBlockingSecurityProfiles,
		GlobalprotectHipObjects:                 o.GlobalprotectHipObjects,
		GlobalprotectHipProfiles:                o.GlobalprotectHipProfiles,
		GtpProtectionSecurityProfiles:           o.GtpProtectionSecurityProfiles,
		LogForwardingProfiles:                   o.LogForwardingProfiles,
		PacketBrokerProfiles:                    o.PacketBrokerProfiles,
		Regions:                                 o.Regions,
		Schedules:                               o.Schedules,
		SctpProtectionSecurityProfiles:          o.SctpProtectionSecurityProfiles,
		SdwanErrorCorrectionProfiles:            o.SdwanErrorCorrectionProfiles,
		SdwanPathQualityProfiles:                o.SdwanPathQualityProfiles,
		SdwanSaasQualityProfiles:                o.SdwanSaasQualityProfiles,
		SdwanTrafficDistributionProfiles:        o.SdwanTrafficDistributionProfiles,
		SecurityProfileGroups:                   o.SecurityProfileGroups,
		ServiceGroups:                           o.ServiceGroups,
		Services:                                o.Services,
		Tags:                                    o.Tags,
		UrlFilteringSecurityProfiles:            o.UrlFilteringSecurityProfiles,
		VulnerabilityProtectionSecurityProfiles: o.VulnerabilityProtectionSecurityProfiles,
		WildfireAnalysisSecurityProfiles:        o.WildfireAnalysisSecurityProfiles,
		Misc:                                    o.Misc,
	}
	return result, nil
}
func (o *roleDeviceRestapiPoliciesXml) MarshalFromObject(s RoleDeviceRestapiPolicies) {
	o.ApplicationOverrideRules = s.ApplicationOverrideRules
	o.AuthenticationRules = s.AuthenticationRules
	o.DecryptionRules = s.DecryptionRules
	o.DosRules = s.DosRules
	o.NatRules = s.NatRules
	o.NetworkPacketBrokerRules = s.NetworkPacketBrokerRules
	o.PolicyBasedForwardingRules = s.PolicyBasedForwardingRules
	o.QosRules = s.QosRules
	o.SdwanRules = s.SdwanRules
	o.SecurityRules = s.SecurityRules
	o.TunnelInspectionRules = s.TunnelInspectionRules
	o.Misc = s.Misc
}

func (o roleDeviceRestapiPoliciesXml) UnmarshalToObject() (*RoleDeviceRestapiPolicies, error) {

	result := &RoleDeviceRestapiPolicies{
		ApplicationOverrideRules:   o.ApplicationOverrideRules,
		AuthenticationRules:        o.AuthenticationRules,
		DecryptionRules:            o.DecryptionRules,
		DosRules:                   o.DosRules,
		NatRules:                   o.NatRules,
		NetworkPacketBrokerRules:   o.NetworkPacketBrokerRules,
		PolicyBasedForwardingRules: o.PolicyBasedForwardingRules,
		QosRules:                   o.QosRules,
		SdwanRules:                 o.SdwanRules,
		SecurityRules:              o.SecurityRules,
		TunnelInspectionRules:      o.TunnelInspectionRules,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *roleDeviceRestapiSystemXml) MarshalFromObject(s RoleDeviceRestapiSystem) {
	o.Configuration = s.Configuration
	o.Misc = s.Misc
}

func (o roleDeviceRestapiSystemXml) UnmarshalToObject() (*RoleDeviceRestapiSystem, error) {

	result := &RoleDeviceRestapiSystem{
		Configuration: o.Configuration,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiXml) MarshalFromObject(s RoleDeviceWebui) {
	o.Acc = s.Acc
	if s.Commit != nil {
		var obj roleDeviceWebuiCommitXml
		obj.MarshalFromObject(*s.Commit)
		o.Commit = &obj
	}
	o.Dashboard = s.Dashboard
	if s.Device != nil {
		var obj roleDeviceWebuiDeviceXml
		obj.MarshalFromObject(*s.Device)
		o.Device = &obj
	}
	if s.Global != nil {
		var obj roleDeviceWebuiGlobalXml
		obj.MarshalFromObject(*s.Global)
		o.Global = &obj
	}
	if s.Monitor != nil {
		var obj roleDeviceWebuiMonitorXml
		obj.MarshalFromObject(*s.Monitor)
		o.Monitor = &obj
	}
	if s.Network != nil {
		var obj roleDeviceWebuiNetworkXml
		obj.MarshalFromObject(*s.Network)
		o.Network = &obj
	}
	if s.Objects != nil {
		var obj roleDeviceWebuiObjectsXml
		obj.MarshalFromObject(*s.Objects)
		o.Objects = &obj
	}
	if s.Operations != nil {
		var obj roleDeviceWebuiOperationsXml
		obj.MarshalFromObject(*s.Operations)
		o.Operations = &obj
	}
	if s.Policies != nil {
		var obj roleDeviceWebuiPoliciesXml
		obj.MarshalFromObject(*s.Policies)
		o.Policies = &obj
	}
	if s.Privacy != nil {
		var obj roleDeviceWebuiPrivacyXml
		obj.MarshalFromObject(*s.Privacy)
		o.Privacy = &obj
	}
	if s.Save != nil {
		var obj roleDeviceWebuiSaveXml
		obj.MarshalFromObject(*s.Save)
		o.Save = &obj
	}
	o.Tasks = s.Tasks
	o.Validate = s.Validate
	o.Misc = s.Misc
}

func (o roleDeviceWebuiXml) UnmarshalToObject() (*RoleDeviceWebui, error) {
	var commitVal *RoleDeviceWebuiCommit
	if o.Commit != nil {
		obj, err := o.Commit.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		commitVal = obj
	}
	var deviceVal *RoleDeviceWebuiDevice
	if o.Device != nil {
		obj, err := o.Device.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		deviceVal = obj
	}
	var globalVal *RoleDeviceWebuiGlobal
	if o.Global != nil {
		obj, err := o.Global.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalVal = obj
	}
	var monitorVal *RoleDeviceWebuiMonitor
	if o.Monitor != nil {
		obj, err := o.Monitor.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		monitorVal = obj
	}
	var networkVal *RoleDeviceWebuiNetwork
	if o.Network != nil {
		obj, err := o.Network.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		networkVal = obj
	}
	var objectsVal *RoleDeviceWebuiObjects
	if o.Objects != nil {
		obj, err := o.Objects.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		objectsVal = obj
	}
	var operationsVal *RoleDeviceWebuiOperations
	if o.Operations != nil {
		obj, err := o.Operations.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		operationsVal = obj
	}
	var policiesVal *RoleDeviceWebuiPolicies
	if o.Policies != nil {
		obj, err := o.Policies.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		policiesVal = obj
	}
	var privacyVal *RoleDeviceWebuiPrivacy
	if o.Privacy != nil {
		obj, err := o.Privacy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		privacyVal = obj
	}
	var saveVal *RoleDeviceWebuiSave
	if o.Save != nil {
		obj, err := o.Save.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		saveVal = obj
	}

	result := &RoleDeviceWebui{
		Acc:        o.Acc,
		Commit:     commitVal,
		Dashboard:  o.Dashboard,
		Device:     deviceVal,
		Global:     globalVal,
		Monitor:    monitorVal,
		Network:    networkVal,
		Objects:    objectsVal,
		Operations: operationsVal,
		Policies:   policiesVal,
		Privacy:    privacyVal,
		Save:       saveVal,
		Tasks:      o.Tasks,
		Validate:   o.Validate,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiCommitXml) MarshalFromObject(s RoleDeviceWebuiCommit) {
	o.CommitForOtherAdmins = s.CommitForOtherAdmins
	o.Device = s.Device
	o.ObjectLevelChanges = s.ObjectLevelChanges
	o.Misc = s.Misc
}

func (o roleDeviceWebuiCommitXml) UnmarshalToObject() (*RoleDeviceWebuiCommit, error) {

	result := &RoleDeviceWebuiCommit{
		CommitForOtherAdmins: o.CommitForOtherAdmins,
		Device:               o.Device,
		ObjectLevelChanges:   o.ObjectLevelChanges,
		Misc:                 o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiDeviceXml) MarshalFromObject(s RoleDeviceWebuiDevice) {
	o.AccessDomain = s.AccessDomain
	o.AdminRoles = s.AdminRoles
	o.Administrators = s.Administrators
	o.AuthenticationProfile = s.AuthenticationProfile
	o.AuthenticationSequence = s.AuthenticationSequence
	o.BlockPages = s.BlockPages
	if s.CertificateManagement != nil {
		var obj roleDeviceWebuiDeviceCertificateManagementXml
		obj.MarshalFromObject(*s.CertificateManagement)
		o.CertificateManagement = &obj
	}
	o.ConfigAudit = s.ConfigAudit
	o.DataRedistribution = s.DataRedistribution
	o.DeviceQuarantine = s.DeviceQuarantine
	o.DhcpSyslogServer = s.DhcpSyslogServer
	o.DynamicUpdates = s.DynamicUpdates
	o.GlobalProtectClient = s.GlobalProtectClient
	o.HighAvailability = s.HighAvailability
	o.Licenses = s.Licenses
	if s.LocalUserDatabase != nil {
		var obj roleDeviceWebuiDeviceLocalUserDatabaseXml
		obj.MarshalFromObject(*s.LocalUserDatabase)
		o.LocalUserDatabase = &obj
	}
	o.LogFwdCard = s.LogFwdCard
	if s.LogSettings != nil {
		var obj roleDeviceWebuiDeviceLogSettingsXml
		obj.MarshalFromObject(*s.LogSettings)
		o.LogSettings = &obj
	}
	o.MasterKey = s.MasterKey
	o.Plugins = s.Plugins
	if s.PolicyRecommendations != nil {
		var obj roleDeviceWebuiDevicePolicyRecommendationsXml
		obj.MarshalFromObject(*s.PolicyRecommendations)
		o.PolicyRecommendations = &obj
	}
	o.ScheduledLogExport = s.ScheduledLogExport
	if s.ServerProfile != nil {
		var obj roleDeviceWebuiDeviceServerProfileXml
		obj.MarshalFromObject(*s.ServerProfile)
		o.ServerProfile = &obj
	}
	if s.Setup != nil {
		var obj roleDeviceWebuiDeviceSetupXml
		obj.MarshalFromObject(*s.Setup)
		o.Setup = &obj
	}
	o.SharedGateways = s.SharedGateways
	o.Software = s.Software
	o.Support = s.Support
	o.Troubleshooting = s.Troubleshooting
	o.UserIdentification = s.UserIdentification
	o.VirtualSystems = s.VirtualSystems
	o.VmInfoSource = s.VmInfoSource
	o.Misc = s.Misc
}

func (o roleDeviceWebuiDeviceXml) UnmarshalToObject() (*RoleDeviceWebuiDevice, error) {
	var certificateManagementVal *RoleDeviceWebuiDeviceCertificateManagement
	if o.CertificateManagement != nil {
		obj, err := o.CertificateManagement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		certificateManagementVal = obj
	}
	var localUserDatabaseVal *RoleDeviceWebuiDeviceLocalUserDatabase
	if o.LocalUserDatabase != nil {
		obj, err := o.LocalUserDatabase.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localUserDatabaseVal = obj
	}
	var logSettingsVal *RoleDeviceWebuiDeviceLogSettings
	if o.LogSettings != nil {
		obj, err := o.LogSettings.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		logSettingsVal = obj
	}
	var policyRecommendationsVal *RoleDeviceWebuiDevicePolicyRecommendations
	if o.PolicyRecommendations != nil {
		obj, err := o.PolicyRecommendations.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		policyRecommendationsVal = obj
	}
	var serverProfileVal *RoleDeviceWebuiDeviceServerProfile
	if o.ServerProfile != nil {
		obj, err := o.ServerProfile.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		serverProfileVal = obj
	}
	var setupVal *RoleDeviceWebuiDeviceSetup
	if o.Setup != nil {
		obj, err := o.Setup.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setupVal = obj
	}

	result := &RoleDeviceWebuiDevice{
		AccessDomain:           o.AccessDomain,
		AdminRoles:             o.AdminRoles,
		Administrators:         o.Administrators,
		AuthenticationProfile:  o.AuthenticationProfile,
		AuthenticationSequence: o.AuthenticationSequence,
		BlockPages:             o.BlockPages,
		CertificateManagement:  certificateManagementVal,
		ConfigAudit:            o.ConfigAudit,
		DataRedistribution:     o.DataRedistribution,
		DeviceQuarantine:       o.DeviceQuarantine,
		DhcpSyslogServer:       o.DhcpSyslogServer,
		DynamicUpdates:         o.DynamicUpdates,
		GlobalProtectClient:    o.GlobalProtectClient,
		HighAvailability:       o.HighAvailability,
		Licenses:               o.Licenses,
		LocalUserDatabase:      localUserDatabaseVal,
		LogFwdCard:             o.LogFwdCard,
		LogSettings:            logSettingsVal,
		MasterKey:              o.MasterKey,
		Plugins:                o.Plugins,
		PolicyRecommendations:  policyRecommendationsVal,
		ScheduledLogExport:     o.ScheduledLogExport,
		ServerProfile:          serverProfileVal,
		Setup:                  setupVal,
		SharedGateways:         o.SharedGateways,
		Software:               o.Software,
		Support:                o.Support,
		Troubleshooting:        o.Troubleshooting,
		UserIdentification:     o.UserIdentification,
		VirtualSystems:         o.VirtualSystems,
		VmInfoSource:           o.VmInfoSource,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiDeviceCertificateManagementXml) MarshalFromObject(s RoleDeviceWebuiDeviceCertificateManagement) {
	o.CertificateProfile = s.CertificateProfile
	o.Certificates = s.Certificates
	o.OcspResponder = s.OcspResponder
	o.Scep = s.Scep
	o.SshServiceProfile = s.SshServiceProfile
	o.SslDecryptionExclusion = s.SslDecryptionExclusion
	o.SslTlsServiceProfile = s.SslTlsServiceProfile
	o.Misc = s.Misc
}

func (o roleDeviceWebuiDeviceCertificateManagementXml) UnmarshalToObject() (*RoleDeviceWebuiDeviceCertificateManagement, error) {

	result := &RoleDeviceWebuiDeviceCertificateManagement{
		CertificateProfile:     o.CertificateProfile,
		Certificates:           o.Certificates,
		OcspResponder:          o.OcspResponder,
		Scep:                   o.Scep,
		SshServiceProfile:      o.SshServiceProfile,
		SslDecryptionExclusion: o.SslDecryptionExclusion,
		SslTlsServiceProfile:   o.SslTlsServiceProfile,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiDeviceLocalUserDatabaseXml) MarshalFromObject(s RoleDeviceWebuiDeviceLocalUserDatabase) {
	o.UserGroups = s.UserGroups
	o.Users = s.Users
	o.Misc = s.Misc
}

func (o roleDeviceWebuiDeviceLocalUserDatabaseXml) UnmarshalToObject() (*RoleDeviceWebuiDeviceLocalUserDatabase, error) {

	result := &RoleDeviceWebuiDeviceLocalUserDatabase{
		UserGroups: o.UserGroups,
		Users:      o.Users,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiDeviceLogSettingsXml) MarshalFromObject(s RoleDeviceWebuiDeviceLogSettings) {
	o.CcAlarm = s.CcAlarm
	o.Config = s.Config
	o.Correlation = s.Correlation
	o.Globalprotect = s.Globalprotect
	o.Hipmatch = s.Hipmatch
	o.Iptag = s.Iptag
	o.ManageLog = s.ManageLog
	o.System = s.System
	o.UserId = s.UserId
	o.Misc = s.Misc
}

func (o roleDeviceWebuiDeviceLogSettingsXml) UnmarshalToObject() (*RoleDeviceWebuiDeviceLogSettings, error) {

	result := &RoleDeviceWebuiDeviceLogSettings{
		CcAlarm:       o.CcAlarm,
		Config:        o.Config,
		Correlation:   o.Correlation,
		Globalprotect: o.Globalprotect,
		Hipmatch:      o.Hipmatch,
		Iptag:         o.Iptag,
		ManageLog:     o.ManageLog,
		System:        o.System,
		UserId:        o.UserId,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiDevicePolicyRecommendationsXml) MarshalFromObject(s RoleDeviceWebuiDevicePolicyRecommendations) {
	o.Iot = s.Iot
	o.Saas = s.Saas
	o.Misc = s.Misc
}

func (o roleDeviceWebuiDevicePolicyRecommendationsXml) UnmarshalToObject() (*RoleDeviceWebuiDevicePolicyRecommendations, error) {

	result := &RoleDeviceWebuiDevicePolicyRecommendations{
		Iot:  o.Iot,
		Saas: o.Saas,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiDeviceServerProfileXml) MarshalFromObject(s RoleDeviceWebuiDeviceServerProfile) {
	o.Dns = s.Dns
	o.Email = s.Email
	o.Http = s.Http
	o.Kerberos = s.Kerberos
	o.Ldap = s.Ldap
	o.Mfa = s.Mfa
	o.Netflow = s.Netflow
	o.Radius = s.Radius
	o.SamlIdp = s.SamlIdp
	o.Scp = s.Scp
	o.SnmpTrap = s.SnmpTrap
	o.Syslog = s.Syslog
	o.Tacplus = s.Tacplus
	o.Misc = s.Misc
}

func (o roleDeviceWebuiDeviceServerProfileXml) UnmarshalToObject() (*RoleDeviceWebuiDeviceServerProfile, error) {

	result := &RoleDeviceWebuiDeviceServerProfile{
		Dns:      o.Dns,
		Email:    o.Email,
		Http:     o.Http,
		Kerberos: o.Kerberos,
		Ldap:     o.Ldap,
		Mfa:      o.Mfa,
		Netflow:  o.Netflow,
		Radius:   o.Radius,
		SamlIdp:  o.SamlIdp,
		Scp:      o.Scp,
		SnmpTrap: o.SnmpTrap,
		Syslog:   o.Syslog,
		Tacplus:  o.Tacplus,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiDeviceSetupXml) MarshalFromObject(s RoleDeviceWebuiDeviceSetup) {
	o.ContentId = s.ContentId
	o.Hsm = s.Hsm
	o.Interfaces = s.Interfaces
	o.Management = s.Management
	o.Operations = s.Operations
	o.Services = s.Services
	o.Session = s.Session
	o.Telemetry = s.Telemetry
	o.Wildfire = s.Wildfire
	o.Misc = s.Misc
}

func (o roleDeviceWebuiDeviceSetupXml) UnmarshalToObject() (*RoleDeviceWebuiDeviceSetup, error) {

	result := &RoleDeviceWebuiDeviceSetup{
		ContentId:  o.ContentId,
		Hsm:        o.Hsm,
		Interfaces: o.Interfaces,
		Management: o.Management,
		Operations: o.Operations,
		Services:   o.Services,
		Session:    o.Session,
		Telemetry:  o.Telemetry,
		Wildfire:   o.Wildfire,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiGlobalXml) MarshalFromObject(s RoleDeviceWebuiGlobal) {
	o.SystemAlarms = s.SystemAlarms
	o.Misc = s.Misc
}

func (o roleDeviceWebuiGlobalXml) UnmarshalToObject() (*RoleDeviceWebuiGlobal, error) {

	result := &RoleDeviceWebuiGlobal{
		SystemAlarms: o.SystemAlarms,
		Misc:         o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiMonitorXml) MarshalFromObject(s RoleDeviceWebuiMonitor) {
	o.AppScope = s.AppScope
	o.ApplicationReports = s.ApplicationReports
	if s.AutomatedCorrelationEngine != nil {
		var obj roleDeviceWebuiMonitorAutomatedCorrelationEngineXml
		obj.MarshalFromObject(*s.AutomatedCorrelationEngine)
		o.AutomatedCorrelationEngine = &obj
	}
	o.BlockIpList = s.BlockIpList
	o.Botnet = s.Botnet
	if s.CustomReports != nil {
		var obj roleDeviceWebuiMonitorCustomReportsXml
		obj.MarshalFromObject(*s.CustomReports)
		o.CustomReports = &obj
	}
	o.ExternalLogs = s.ExternalLogs
	o.GtpReports = s.GtpReports
	if s.Logs != nil {
		var obj roleDeviceWebuiMonitorLogsXml
		obj.MarshalFromObject(*s.Logs)
		o.Logs = &obj
	}
	o.PacketCapture = s.PacketCapture
	if s.PdfReports != nil {
		var obj roleDeviceWebuiMonitorPdfReportsXml
		obj.MarshalFromObject(*s.PdfReports)
		o.PdfReports = &obj
	}
	o.SctpReports = s.SctpReports
	o.SessionBrowser = s.SessionBrowser
	o.ThreatReports = s.ThreatReports
	o.TrafficReports = s.TrafficReports
	o.UrlFilteringReports = s.UrlFilteringReports
	o.ViewCustomReports = s.ViewCustomReports
	o.Misc = s.Misc
}

func (o roleDeviceWebuiMonitorXml) UnmarshalToObject() (*RoleDeviceWebuiMonitor, error) {
	var automatedCorrelationEngineVal *RoleDeviceWebuiMonitorAutomatedCorrelationEngine
	if o.AutomatedCorrelationEngine != nil {
		obj, err := o.AutomatedCorrelationEngine.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		automatedCorrelationEngineVal = obj
	}
	var customReportsVal *RoleDeviceWebuiMonitorCustomReports
	if o.CustomReports != nil {
		obj, err := o.CustomReports.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		customReportsVal = obj
	}
	var logsVal *RoleDeviceWebuiMonitorLogs
	if o.Logs != nil {
		obj, err := o.Logs.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		logsVal = obj
	}
	var pdfReportsVal *RoleDeviceWebuiMonitorPdfReports
	if o.PdfReports != nil {
		obj, err := o.PdfReports.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		pdfReportsVal = obj
	}

	result := &RoleDeviceWebuiMonitor{
		AppScope:                   o.AppScope,
		ApplicationReports:         o.ApplicationReports,
		AutomatedCorrelationEngine: automatedCorrelationEngineVal,
		BlockIpList:                o.BlockIpList,
		Botnet:                     o.Botnet,
		CustomReports:              customReportsVal,
		ExternalLogs:               o.ExternalLogs,
		GtpReports:                 o.GtpReports,
		Logs:                       logsVal,
		PacketCapture:              o.PacketCapture,
		PdfReports:                 pdfReportsVal,
		SctpReports:                o.SctpReports,
		SessionBrowser:             o.SessionBrowser,
		ThreatReports:              o.ThreatReports,
		TrafficReports:             o.TrafficReports,
		UrlFilteringReports:        o.UrlFilteringReports,
		ViewCustomReports:          o.ViewCustomReports,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiMonitorAutomatedCorrelationEngineXml) MarshalFromObject(s RoleDeviceWebuiMonitorAutomatedCorrelationEngine) {
	o.CorrelatedEvents = s.CorrelatedEvents
	o.CorrelationObjects = s.CorrelationObjects
	o.Misc = s.Misc
}

func (o roleDeviceWebuiMonitorAutomatedCorrelationEngineXml) UnmarshalToObject() (*RoleDeviceWebuiMonitorAutomatedCorrelationEngine, error) {

	result := &RoleDeviceWebuiMonitorAutomatedCorrelationEngine{
		CorrelatedEvents:   o.CorrelatedEvents,
		CorrelationObjects: o.CorrelationObjects,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiMonitorCustomReportsXml) MarshalFromObject(s RoleDeviceWebuiMonitorCustomReports) {
	o.ApplicationStatistics = s.ApplicationStatistics
	o.Auth = s.Auth
	o.DataFilteringLog = s.DataFilteringLog
	o.DecryptionLog = s.DecryptionLog
	o.DecryptionSummary = s.DecryptionSummary
	o.Globalprotect = s.Globalprotect
	o.GtpLog = s.GtpLog
	o.GtpSummary = s.GtpSummary
	o.Hipmatch = s.Hipmatch
	o.Iptag = s.Iptag
	o.SctpLog = s.SctpLog
	o.SctpSummary = s.SctpSummary
	o.ThreatLog = s.ThreatLog
	o.ThreatSummary = s.ThreatSummary
	o.TrafficLog = s.TrafficLog
	o.TrafficSummary = s.TrafficSummary
	o.TunnelLog = s.TunnelLog
	o.TunnelSummary = s.TunnelSummary
	o.UrlLog = s.UrlLog
	o.UrlSummary = s.UrlSummary
	o.Userid = s.Userid
	o.WildfireLog = s.WildfireLog
	o.Misc = s.Misc
}

func (o roleDeviceWebuiMonitorCustomReportsXml) UnmarshalToObject() (*RoleDeviceWebuiMonitorCustomReports, error) {

	result := &RoleDeviceWebuiMonitorCustomReports{
		ApplicationStatistics: o.ApplicationStatistics,
		Auth:                  o.Auth,
		DataFilteringLog:      o.DataFilteringLog,
		DecryptionLog:         o.DecryptionLog,
		DecryptionSummary:     o.DecryptionSummary,
		Globalprotect:         o.Globalprotect,
		GtpLog:                o.GtpLog,
		GtpSummary:            o.GtpSummary,
		Hipmatch:              o.Hipmatch,
		Iptag:                 o.Iptag,
		SctpLog:               o.SctpLog,
		SctpSummary:           o.SctpSummary,
		ThreatLog:             o.ThreatLog,
		ThreatSummary:         o.ThreatSummary,
		TrafficLog:            o.TrafficLog,
		TrafficSummary:        o.TrafficSummary,
		TunnelLog:             o.TunnelLog,
		TunnelSummary:         o.TunnelSummary,
		UrlLog:                o.UrlLog,
		UrlSummary:            o.UrlSummary,
		Userid:                o.Userid,
		WildfireLog:           o.WildfireLog,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiMonitorLogsXml) MarshalFromObject(s RoleDeviceWebuiMonitorLogs) {
	o.Alarm = s.Alarm
	o.Authentication = s.Authentication
	o.Configuration = s.Configuration
	o.DataFiltering = s.DataFiltering
	o.Decryption = s.Decryption
	o.Globalprotect = s.Globalprotect
	o.Gtp = s.Gtp
	o.Hipmatch = s.Hipmatch
	o.Iptag = s.Iptag
	o.Sctp = s.Sctp
	o.System = s.System
	o.Threat = s.Threat
	o.Traffic = s.Traffic
	o.Tunnel = s.Tunnel
	o.Url = s.Url
	o.Userid = s.Userid
	o.Wildfire = s.Wildfire
	o.Misc = s.Misc
}

func (o roleDeviceWebuiMonitorLogsXml) UnmarshalToObject() (*RoleDeviceWebuiMonitorLogs, error) {

	result := &RoleDeviceWebuiMonitorLogs{
		Alarm:          o.Alarm,
		Authentication: o.Authentication,
		Configuration:  o.Configuration,
		DataFiltering:  o.DataFiltering,
		Decryption:     o.Decryption,
		Globalprotect:  o.Globalprotect,
		Gtp:            o.Gtp,
		Hipmatch:       o.Hipmatch,
		Iptag:          o.Iptag,
		Sctp:           o.Sctp,
		System:         o.System,
		Threat:         o.Threat,
		Traffic:        o.Traffic,
		Tunnel:         o.Tunnel,
		Url:            o.Url,
		Userid:         o.Userid,
		Wildfire:       o.Wildfire,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiMonitorPdfReportsXml) MarshalFromObject(s RoleDeviceWebuiMonitorPdfReports) {
	o.EmailScheduler = s.EmailScheduler
	o.ManagePdfSummary = s.ManagePdfSummary
	o.PdfSummaryReports = s.PdfSummaryReports
	o.ReportGroups = s.ReportGroups
	o.SaasApplicationUsageReport = s.SaasApplicationUsageReport
	o.UserActivityReport = s.UserActivityReport
	o.Misc = s.Misc
}

func (o roleDeviceWebuiMonitorPdfReportsXml) UnmarshalToObject() (*RoleDeviceWebuiMonitorPdfReports, error) {

	result := &RoleDeviceWebuiMonitorPdfReports{
		EmailScheduler:             o.EmailScheduler,
		ManagePdfSummary:           o.ManagePdfSummary,
		PdfSummaryReports:          o.PdfSummaryReports,
		ReportGroups:               o.ReportGroups,
		SaasApplicationUsageReport: o.SaasApplicationUsageReport,
		UserActivityReport:         o.UserActivityReport,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiNetworkXml) MarshalFromObject(s RoleDeviceWebuiNetwork) {
	o.Dhcp = s.Dhcp
	o.DnsProxy = s.DnsProxy
	if s.GlobalProtect != nil {
		var obj roleDeviceWebuiNetworkGlobalProtectXml
		obj.MarshalFromObject(*s.GlobalProtect)
		o.GlobalProtect = &obj
	}
	o.GreTunnels = s.GreTunnels
	o.Interfaces = s.Interfaces
	o.IpsecTunnels = s.IpsecTunnels
	o.Lldp = s.Lldp
	if s.NetworkProfiles != nil {
		var obj roleDeviceWebuiNetworkNetworkProfilesXml
		obj.MarshalFromObject(*s.NetworkProfiles)
		o.NetworkProfiles = &obj
	}
	o.Qos = s.Qos
	if s.Routing != nil {
		var obj roleDeviceWebuiNetworkRoutingXml
		obj.MarshalFromObject(*s.Routing)
		o.Routing = &obj
	}
	o.SdwanInterfaceProfile = s.SdwanInterfaceProfile
	o.SecureWebGateway = s.SecureWebGateway
	o.VirtualRouters = s.VirtualRouters
	o.VirtualWires = s.VirtualWires
	o.Vlans = s.Vlans
	o.Zones = s.Zones
	o.Misc = s.Misc
}

func (o roleDeviceWebuiNetworkXml) UnmarshalToObject() (*RoleDeviceWebuiNetwork, error) {
	var globalProtectVal *RoleDeviceWebuiNetworkGlobalProtect
	if o.GlobalProtect != nil {
		obj, err := o.GlobalProtect.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalProtectVal = obj
	}
	var networkProfilesVal *RoleDeviceWebuiNetworkNetworkProfiles
	if o.NetworkProfiles != nil {
		obj, err := o.NetworkProfiles.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		networkProfilesVal = obj
	}
	var routingVal *RoleDeviceWebuiNetworkRouting
	if o.Routing != nil {
		obj, err := o.Routing.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routingVal = obj
	}

	result := &RoleDeviceWebuiNetwork{
		Dhcp:                  o.Dhcp,
		DnsProxy:              o.DnsProxy,
		GlobalProtect:         globalProtectVal,
		GreTunnels:            o.GreTunnels,
		Interfaces:            o.Interfaces,
		IpsecTunnels:          o.IpsecTunnels,
		Lldp:                  o.Lldp,
		NetworkProfiles:       networkProfilesVal,
		Qos:                   o.Qos,
		Routing:               routingVal,
		SdwanInterfaceProfile: o.SdwanInterfaceProfile,
		SecureWebGateway:      o.SecureWebGateway,
		VirtualRouters:        o.VirtualRouters,
		VirtualWires:          o.VirtualWires,
		Vlans:                 o.Vlans,
		Zones:                 o.Zones,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiNetworkGlobalProtectXml) MarshalFromObject(s RoleDeviceWebuiNetworkGlobalProtect) {
	o.ClientlessAppGroups = s.ClientlessAppGroups
	o.ClientlessApps = s.ClientlessApps
	o.Gateways = s.Gateways
	o.Mdm = s.Mdm
	o.Portals = s.Portals
	o.Misc = s.Misc
}

func (o roleDeviceWebuiNetworkGlobalProtectXml) UnmarshalToObject() (*RoleDeviceWebuiNetworkGlobalProtect, error) {

	result := &RoleDeviceWebuiNetworkGlobalProtect{
		ClientlessAppGroups: o.ClientlessAppGroups,
		ClientlessApps:      o.ClientlessApps,
		Gateways:            o.Gateways,
		Mdm:                 o.Mdm,
		Portals:             o.Portals,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiNetworkNetworkProfilesXml) MarshalFromObject(s RoleDeviceWebuiNetworkNetworkProfiles) {
	o.BfdProfile = s.BfdProfile
	o.GpAppIpsecCrypto = s.GpAppIpsecCrypto
	o.IkeCrypto = s.IkeCrypto
	o.IkeGateways = s.IkeGateways
	o.InterfaceMgmt = s.InterfaceMgmt
	o.IpsecCrypto = s.IpsecCrypto
	o.LldpProfile = s.LldpProfile
	o.QosProfile = s.QosProfile
	o.TunnelMonitor = s.TunnelMonitor
	o.ZoneProtection = s.ZoneProtection
	o.Misc = s.Misc
}

func (o roleDeviceWebuiNetworkNetworkProfilesXml) UnmarshalToObject() (*RoleDeviceWebuiNetworkNetworkProfiles, error) {

	result := &RoleDeviceWebuiNetworkNetworkProfiles{
		BfdProfile:       o.BfdProfile,
		GpAppIpsecCrypto: o.GpAppIpsecCrypto,
		IkeCrypto:        o.IkeCrypto,
		IkeGateways:      o.IkeGateways,
		InterfaceMgmt:    o.InterfaceMgmt,
		IpsecCrypto:      o.IpsecCrypto,
		LldpProfile:      o.LldpProfile,
		QosProfile:       o.QosProfile,
		TunnelMonitor:    o.TunnelMonitor,
		ZoneProtection:   o.ZoneProtection,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiNetworkRoutingXml) MarshalFromObject(s RoleDeviceWebuiNetworkRouting) {
	o.LogicalRouters = s.LogicalRouters
	if s.RoutingProfiles != nil {
		var obj roleDeviceWebuiNetworkRoutingRoutingProfilesXml
		obj.MarshalFromObject(*s.RoutingProfiles)
		o.RoutingProfiles = &obj
	}
	o.Misc = s.Misc
}

func (o roleDeviceWebuiNetworkRoutingXml) UnmarshalToObject() (*RoleDeviceWebuiNetworkRouting, error) {
	var routingProfilesVal *RoleDeviceWebuiNetworkRoutingRoutingProfiles
	if o.RoutingProfiles != nil {
		obj, err := o.RoutingProfiles.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routingProfilesVal = obj
	}

	result := &RoleDeviceWebuiNetworkRouting{
		LogicalRouters:  o.LogicalRouters,
		RoutingProfiles: routingProfilesVal,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiNetworkRoutingRoutingProfilesXml) MarshalFromObject(s RoleDeviceWebuiNetworkRoutingRoutingProfiles) {
	o.Bfd = s.Bfd
	o.Bgp = s.Bgp
	o.Filters = s.Filters
	o.Multicast = s.Multicast
	o.Ospf = s.Ospf
	o.Ospfv3 = s.Ospfv3
	o.Ripv2 = s.Ripv2
	o.Misc = s.Misc
}

func (o roleDeviceWebuiNetworkRoutingRoutingProfilesXml) UnmarshalToObject() (*RoleDeviceWebuiNetworkRoutingRoutingProfiles, error) {

	result := &RoleDeviceWebuiNetworkRoutingRoutingProfiles{
		Bfd:       o.Bfd,
		Bgp:       o.Bgp,
		Filters:   o.Filters,
		Multicast: o.Multicast,
		Ospf:      o.Ospf,
		Ospfv3:    o.Ospfv3,
		Ripv2:     o.Ripv2,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiObjectsXml) MarshalFromObject(s RoleDeviceWebuiObjects) {
	o.AddressGroups = s.AddressGroups
	o.Addresses = s.Addresses
	o.ApplicationFilters = s.ApplicationFilters
	o.ApplicationGroups = s.ApplicationGroups
	o.Applications = s.Applications
	o.Authentication = s.Authentication
	if s.CustomObjects != nil {
		var obj roleDeviceWebuiObjectsCustomObjectsXml
		obj.MarshalFromObject(*s.CustomObjects)
		o.CustomObjects = &obj
	}
	if s.Decryption != nil {
		var obj roleDeviceWebuiObjectsDecryptionXml
		obj.MarshalFromObject(*s.Decryption)
		o.Decryption = &obj
	}
	o.Devices = s.Devices
	o.DynamicBlockLists = s.DynamicBlockLists
	o.DynamicUserGroups = s.DynamicUserGroups
	if s.GlobalProtect != nil {
		var obj roleDeviceWebuiObjectsGlobalProtectXml
		obj.MarshalFromObject(*s.GlobalProtect)
		o.GlobalProtect = &obj
	}
	o.LogForwarding = s.LogForwarding
	o.PacketBrokerProfile = s.PacketBrokerProfile
	o.Regions = s.Regions
	o.Schedules = s.Schedules
	if s.Sdwan != nil {
		var obj roleDeviceWebuiObjectsSdwanXml
		obj.MarshalFromObject(*s.Sdwan)
		o.Sdwan = &obj
	}
	o.SecurityProfileGroups = s.SecurityProfileGroups
	if s.SecurityProfiles != nil {
		var obj roleDeviceWebuiObjectsSecurityProfilesXml
		obj.MarshalFromObject(*s.SecurityProfiles)
		o.SecurityProfiles = &obj
	}
	o.ServiceGroups = s.ServiceGroups
	o.Services = s.Services
	o.Tags = s.Tags
	o.Misc = s.Misc
}

func (o roleDeviceWebuiObjectsXml) UnmarshalToObject() (*RoleDeviceWebuiObjects, error) {
	var customObjectsVal *RoleDeviceWebuiObjectsCustomObjects
	if o.CustomObjects != nil {
		obj, err := o.CustomObjects.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		customObjectsVal = obj
	}
	var decryptionVal *RoleDeviceWebuiObjectsDecryption
	if o.Decryption != nil {
		obj, err := o.Decryption.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		decryptionVal = obj
	}
	var globalProtectVal *RoleDeviceWebuiObjectsGlobalProtect
	if o.GlobalProtect != nil {
		obj, err := o.GlobalProtect.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalProtectVal = obj
	}
	var sdwanVal *RoleDeviceWebuiObjectsSdwan
	if o.Sdwan != nil {
		obj, err := o.Sdwan.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sdwanVal = obj
	}
	var securityProfilesVal *RoleDeviceWebuiObjectsSecurityProfiles
	if o.SecurityProfiles != nil {
		obj, err := o.SecurityProfiles.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		securityProfilesVal = obj
	}

	result := &RoleDeviceWebuiObjects{
		AddressGroups:         o.AddressGroups,
		Addresses:             o.Addresses,
		ApplicationFilters:    o.ApplicationFilters,
		ApplicationGroups:     o.ApplicationGroups,
		Applications:          o.Applications,
		Authentication:        o.Authentication,
		CustomObjects:         customObjectsVal,
		Decryption:            decryptionVal,
		Devices:               o.Devices,
		DynamicBlockLists:     o.DynamicBlockLists,
		DynamicUserGroups:     o.DynamicUserGroups,
		GlobalProtect:         globalProtectVal,
		LogForwarding:         o.LogForwarding,
		PacketBrokerProfile:   o.PacketBrokerProfile,
		Regions:               o.Regions,
		Schedules:             o.Schedules,
		Sdwan:                 sdwanVal,
		SecurityProfileGroups: o.SecurityProfileGroups,
		SecurityProfiles:      securityProfilesVal,
		ServiceGroups:         o.ServiceGroups,
		Services:              o.Services,
		Tags:                  o.Tags,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiObjectsCustomObjectsXml) MarshalFromObject(s RoleDeviceWebuiObjectsCustomObjects) {
	o.DataPatterns = s.DataPatterns
	o.Spyware = s.Spyware
	o.UrlCategory = s.UrlCategory
	o.Vulnerability = s.Vulnerability
	o.Misc = s.Misc
}

func (o roleDeviceWebuiObjectsCustomObjectsXml) UnmarshalToObject() (*RoleDeviceWebuiObjectsCustomObjects, error) {

	result := &RoleDeviceWebuiObjectsCustomObjects{
		DataPatterns:  o.DataPatterns,
		Spyware:       o.Spyware,
		UrlCategory:   o.UrlCategory,
		Vulnerability: o.Vulnerability,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiObjectsDecryptionXml) MarshalFromObject(s RoleDeviceWebuiObjectsDecryption) {
	o.DecryptionProfile = s.DecryptionProfile
	o.Misc = s.Misc
}

func (o roleDeviceWebuiObjectsDecryptionXml) UnmarshalToObject() (*RoleDeviceWebuiObjectsDecryption, error) {

	result := &RoleDeviceWebuiObjectsDecryption{
		DecryptionProfile: o.DecryptionProfile,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiObjectsGlobalProtectXml) MarshalFromObject(s RoleDeviceWebuiObjectsGlobalProtect) {
	o.HipObjects = s.HipObjects
	o.HipProfiles = s.HipProfiles
	o.Misc = s.Misc
}

func (o roleDeviceWebuiObjectsGlobalProtectXml) UnmarshalToObject() (*RoleDeviceWebuiObjectsGlobalProtect, error) {

	result := &RoleDeviceWebuiObjectsGlobalProtect{
		HipObjects:  o.HipObjects,
		HipProfiles: o.HipProfiles,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiObjectsSdwanXml) MarshalFromObject(s RoleDeviceWebuiObjectsSdwan) {
	o.SdwanDistProfile = s.SdwanDistProfile
	o.SdwanErrorCorrectionProfile = s.SdwanErrorCorrectionProfile
	o.SdwanProfile = s.SdwanProfile
	o.SdwanSaasQualityProfile = s.SdwanSaasQualityProfile
	o.Misc = s.Misc
}

func (o roleDeviceWebuiObjectsSdwanXml) UnmarshalToObject() (*RoleDeviceWebuiObjectsSdwan, error) {

	result := &RoleDeviceWebuiObjectsSdwan{
		SdwanDistProfile:            o.SdwanDistProfile,
		SdwanErrorCorrectionProfile: o.SdwanErrorCorrectionProfile,
		SdwanProfile:                o.SdwanProfile,
		SdwanSaasQualityProfile:     o.SdwanSaasQualityProfile,
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiObjectsSecurityProfilesXml) MarshalFromObject(s RoleDeviceWebuiObjectsSecurityProfiles) {
	o.AntiSpyware = s.AntiSpyware
	o.Antivirus = s.Antivirus
	o.DataFiltering = s.DataFiltering
	o.DosProtection = s.DosProtection
	o.FileBlocking = s.FileBlocking
	o.GtpProtection = s.GtpProtection
	o.SctpProtection = s.SctpProtection
	o.UrlFiltering = s.UrlFiltering
	o.VulnerabilityProtection = s.VulnerabilityProtection
	o.WildfireAnalysis = s.WildfireAnalysis
	o.Misc = s.Misc
}

func (o roleDeviceWebuiObjectsSecurityProfilesXml) UnmarshalToObject() (*RoleDeviceWebuiObjectsSecurityProfiles, error) {

	result := &RoleDeviceWebuiObjectsSecurityProfiles{
		AntiSpyware:             o.AntiSpyware,
		Antivirus:               o.Antivirus,
		DataFiltering:           o.DataFiltering,
		DosProtection:           o.DosProtection,
		FileBlocking:            o.FileBlocking,
		GtpProtection:           o.GtpProtection,
		SctpProtection:          o.SctpProtection,
		UrlFiltering:            o.UrlFiltering,
		VulnerabilityProtection: o.VulnerabilityProtection,
		WildfireAnalysis:        o.WildfireAnalysis,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiOperationsXml) MarshalFromObject(s RoleDeviceWebuiOperations) {
	o.DownloadCoreFiles = s.DownloadCoreFiles
	o.DownloadPcapFiles = s.DownloadPcapFiles
	o.GenerateStatsDumpFile = s.GenerateStatsDumpFile
	o.GenerateTechSupportFile = s.GenerateTechSupportFile
	o.Reboot = s.Reboot
	o.Misc = s.Misc
}

func (o roleDeviceWebuiOperationsXml) UnmarshalToObject() (*RoleDeviceWebuiOperations, error) {

	result := &RoleDeviceWebuiOperations{
		DownloadCoreFiles:       o.DownloadCoreFiles,
		DownloadPcapFiles:       o.DownloadPcapFiles,
		GenerateStatsDumpFile:   o.GenerateStatsDumpFile,
		GenerateTechSupportFile: o.GenerateTechSupportFile,
		Reboot:                  o.Reboot,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiPoliciesXml) MarshalFromObject(s RoleDeviceWebuiPolicies) {
	o.ApplicationOverrideRulebase = s.ApplicationOverrideRulebase
	o.AuthenticationRulebase = s.AuthenticationRulebase
	o.DosRulebase = s.DosRulebase
	o.NatRulebase = s.NatRulebase
	o.NetworkPacketBrokerRulebase = s.NetworkPacketBrokerRulebase
	o.PbfRulebase = s.PbfRulebase
	o.QosRulebase = s.QosRulebase
	o.RuleHitCountReset = s.RuleHitCountReset
	o.SdwanRulebase = s.SdwanRulebase
	o.SecurityRulebase = s.SecurityRulebase
	o.SslDecryptionRulebase = s.SslDecryptionRulebase
	o.TunnelInspectRulebase = s.TunnelInspectRulebase
	o.Misc = s.Misc
}

func (o roleDeviceWebuiPoliciesXml) UnmarshalToObject() (*RoleDeviceWebuiPolicies, error) {

	result := &RoleDeviceWebuiPolicies{
		ApplicationOverrideRulebase: o.ApplicationOverrideRulebase,
		AuthenticationRulebase:      o.AuthenticationRulebase,
		DosRulebase:                 o.DosRulebase,
		NatRulebase:                 o.NatRulebase,
		NetworkPacketBrokerRulebase: o.NetworkPacketBrokerRulebase,
		PbfRulebase:                 o.PbfRulebase,
		QosRulebase:                 o.QosRulebase,
		RuleHitCountReset:           o.RuleHitCountReset,
		SdwanRulebase:               o.SdwanRulebase,
		SecurityRulebase:            o.SecurityRulebase,
		SslDecryptionRulebase:       o.SslDecryptionRulebase,
		TunnelInspectRulebase:       o.TunnelInspectRulebase,
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiPrivacyXml) MarshalFromObject(s RoleDeviceWebuiPrivacy) {
	o.ShowFullIpAddresses = s.ShowFullIpAddresses
	o.ShowUserNamesInLogsAndReports = s.ShowUserNamesInLogsAndReports
	o.ViewPcapFiles = s.ViewPcapFiles
	o.Misc = s.Misc
}

func (o roleDeviceWebuiPrivacyXml) UnmarshalToObject() (*RoleDeviceWebuiPrivacy, error) {

	result := &RoleDeviceWebuiPrivacy{
		ShowFullIpAddresses:           o.ShowFullIpAddresses,
		ShowUserNamesInLogsAndReports: o.ShowUserNamesInLogsAndReports,
		ViewPcapFiles:                 o.ViewPcapFiles,
		Misc:                          o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiSaveXml) MarshalFromObject(s RoleDeviceWebuiSave) {
	o.ObjectLevelChanges = s.ObjectLevelChanges
	o.PartialSave = s.PartialSave
	o.SaveForOtherAdmins = s.SaveForOtherAdmins
	o.Misc = s.Misc
}

func (o roleDeviceWebuiSaveXml) UnmarshalToObject() (*RoleDeviceWebuiSave, error) {

	result := &RoleDeviceWebuiSave{
		ObjectLevelChanges: o.ObjectLevelChanges,
		PartialSave:        o.PartialSave,
		SaveForOtherAdmins: o.SaveForOtherAdmins,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *roleDeviceXmlapiXml) MarshalFromObject(s RoleDeviceXmlapi) {
	o.Commit = s.Commit
	o.Config = s.Config
	o.Export = s.Export
	o.Import = s.Import
	o.Iot = s.Iot
	o.Log = s.Log
	o.Op = s.Op
	o.Report = s.Report
	o.UserId = s.UserId
	o.Misc = s.Misc
}

func (o roleDeviceXmlapiXml) UnmarshalToObject() (*RoleDeviceXmlapi, error) {

	result := &RoleDeviceXmlapi{
		Commit: o.Commit,
		Config: o.Config,
		Export: o.Export,
		Import: o.Import,
		Iot:    o.Iot,
		Log:    o.Log,
		Op:     o.Op,
		Report: o.Report,
		UserId: o.UserId,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *roleVsysXml) MarshalFromObject(s RoleVsys) {
	o.Cli = s.Cli
	if s.Restapi != nil {
		var obj roleVsysRestapiXml
		obj.MarshalFromObject(*s.Restapi)
		o.Restapi = &obj
	}
	if s.Webui != nil {
		var obj roleVsysWebuiXml
		obj.MarshalFromObject(*s.Webui)
		o.Webui = &obj
	}
	if s.Xmlapi != nil {
		var obj roleVsysXmlapiXml
		obj.MarshalFromObject(*s.Xmlapi)
		o.Xmlapi = &obj
	}
	o.Misc = s.Misc
}

func (o roleVsysXml) UnmarshalToObject() (*RoleVsys, error) {
	var restapiVal *RoleVsysRestapi
	if o.Restapi != nil {
		obj, err := o.Restapi.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		restapiVal = obj
	}
	var webuiVal *RoleVsysWebui
	if o.Webui != nil {
		obj, err := o.Webui.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		webuiVal = obj
	}
	var xmlapiVal *RoleVsysXmlapi
	if o.Xmlapi != nil {
		obj, err := o.Xmlapi.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		xmlapiVal = obj
	}

	result := &RoleVsys{
		Cli:     o.Cli,
		Restapi: restapiVal,
		Webui:   webuiVal,
		Xmlapi:  xmlapiVal,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *roleVsysRestapiXml) MarshalFromObject(s RoleVsysRestapi) {
	if s.Device != nil {
		var obj roleVsysRestapiDeviceXml
		obj.MarshalFromObject(*s.Device)
		o.Device = &obj
	}
	if s.Network != nil {
		var obj roleVsysRestapiNetworkXml
		obj.MarshalFromObject(*s.Network)
		o.Network = &obj
	}
	if s.Objects != nil {
		var obj roleVsysRestapiObjectsXml
		obj.MarshalFromObject(*s.Objects)
		o.Objects = &obj
	}
	if s.Policies != nil {
		var obj roleVsysRestapiPoliciesXml
		obj.MarshalFromObject(*s.Policies)
		o.Policies = &obj
	}
	if s.System != nil {
		var obj roleVsysRestapiSystemXml
		obj.MarshalFromObject(*s.System)
		o.System = &obj
	}
	o.Misc = s.Misc
}

func (o roleVsysRestapiXml) UnmarshalToObject() (*RoleVsysRestapi, error) {
	var deviceVal *RoleVsysRestapiDevice
	if o.Device != nil {
		obj, err := o.Device.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		deviceVal = obj
	}
	var networkVal *RoleVsysRestapiNetwork
	if o.Network != nil {
		obj, err := o.Network.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		networkVal = obj
	}
	var objectsVal *RoleVsysRestapiObjects
	if o.Objects != nil {
		obj, err := o.Objects.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		objectsVal = obj
	}
	var policiesVal *RoleVsysRestapiPolicies
	if o.Policies != nil {
		obj, err := o.Policies.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		policiesVal = obj
	}
	var systemVal *RoleVsysRestapiSystem
	if o.System != nil {
		obj, err := o.System.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		systemVal = obj
	}

	result := &RoleVsysRestapi{
		Device:   deviceVal,
		Network:  networkVal,
		Objects:  objectsVal,
		Policies: policiesVal,
		System:   systemVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *roleVsysRestapiDeviceXml) MarshalFromObject(s RoleVsysRestapiDevice) {
	o.EmailServerProfiles = s.EmailServerProfiles
	o.HttpServerProfiles = s.HttpServerProfiles
	o.LdapServerProfiles = s.LdapServerProfiles
	o.LogInterfaceSetting = s.LogInterfaceSetting
	o.SnmpTrapServerProfiles = s.SnmpTrapServerProfiles
	o.SyslogServerProfiles = s.SyslogServerProfiles
	o.VirtualSystems = s.VirtualSystems
	o.Misc = s.Misc
}

func (o roleVsysRestapiDeviceXml) UnmarshalToObject() (*RoleVsysRestapiDevice, error) {

	result := &RoleVsysRestapiDevice{
		EmailServerProfiles:    o.EmailServerProfiles,
		HttpServerProfiles:     o.HttpServerProfiles,
		LdapServerProfiles:     o.LdapServerProfiles,
		LogInterfaceSetting:    o.LogInterfaceSetting,
		SnmpTrapServerProfiles: o.SnmpTrapServerProfiles,
		SyslogServerProfiles:   o.SyslogServerProfiles,
		VirtualSystems:         o.VirtualSystems,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *roleVsysRestapiNetworkXml) MarshalFromObject(s RoleVsysRestapiNetwork) {
	o.GlobalprotectClientlessAppGroups = s.GlobalprotectClientlessAppGroups
	o.GlobalprotectClientlessApps = s.GlobalprotectClientlessApps
	o.GlobalprotectGateways = s.GlobalprotectGateways
	o.GlobalprotectMdmServers = s.GlobalprotectMdmServers
	o.GlobalprotectPortals = s.GlobalprotectPortals
	o.SdwanInterfaceProfiles = s.SdwanInterfaceProfiles
	o.Zones = s.Zones
	o.Misc = s.Misc
}

func (o roleVsysRestapiNetworkXml) UnmarshalToObject() (*RoleVsysRestapiNetwork, error) {

	result := &RoleVsysRestapiNetwork{
		GlobalprotectClientlessAppGroups: o.GlobalprotectClientlessAppGroups,
		GlobalprotectClientlessApps:      o.GlobalprotectClientlessApps,
		GlobalprotectGateways:            o.GlobalprotectGateways,
		GlobalprotectMdmServers:          o.GlobalprotectMdmServers,
		GlobalprotectPortals:             o.GlobalprotectPortals,
		SdwanInterfaceProfiles:           o.SdwanInterfaceProfiles,
		Zones:                            o.Zones,
		Misc:                             o.Misc,
	}
	return result, nil
}
func (o *roleVsysRestapiObjectsXml) MarshalFromObject(s RoleVsysRestapiObjects) {
	o.AddressGroups = s.AddressGroups
	o.Addresses = s.Addresses
	o.AntiSpywareSecurityProfiles = s.AntiSpywareSecurityProfiles
	o.AntivirusSecurityProfiles = s.AntivirusSecurityProfiles
	o.ApplicationFilters = s.ApplicationFilters
	o.ApplicationGroups = s.ApplicationGroups
	o.Applications = s.Applications
	o.AuthenticationEnforcements = s.AuthenticationEnforcements
	o.CustomDataPatterns = s.CustomDataPatterns
	o.CustomSpywareSignatures = s.CustomSpywareSignatures
	o.CustomUrlCategories = s.CustomUrlCategories
	o.CustomVulnerabilitySignatures = s.CustomVulnerabilitySignatures
	o.DataFilteringSecurityProfiles = s.DataFilteringSecurityProfiles
	o.DecryptionProfiles = s.DecryptionProfiles
	o.Devices = s.Devices
	o.DosProtectionSecurityProfiles = s.DosProtectionSecurityProfiles
	o.DynamicUserGroups = s.DynamicUserGroups
	o.ExternalDynamicLists = s.ExternalDynamicLists
	o.FileBlockingSecurityProfiles = s.FileBlockingSecurityProfiles
	o.GlobalprotectHipObjects = s.GlobalprotectHipObjects
	o.GlobalprotectHipProfiles = s.GlobalprotectHipProfiles
	o.GtpProtectionSecurityProfiles = s.GtpProtectionSecurityProfiles
	o.LogForwardingProfiles = s.LogForwardingProfiles
	o.PacketBrokerProfiles = s.PacketBrokerProfiles
	o.Regions = s.Regions
	o.Schedules = s.Schedules
	o.SctpProtectionSecurityProfiles = s.SctpProtectionSecurityProfiles
	o.SdwanErrorCorrectionProfiles = s.SdwanErrorCorrectionProfiles
	o.SdwanPathQualityProfiles = s.SdwanPathQualityProfiles
	o.SdwanSaasQualityProfiles = s.SdwanSaasQualityProfiles
	o.SdwanTrafficDistributionProfiles = s.SdwanTrafficDistributionProfiles
	o.SecurityProfileGroups = s.SecurityProfileGroups
	o.ServiceGroups = s.ServiceGroups
	o.Services = s.Services
	o.Tags = s.Tags
	o.UrlFilteringSecurityProfiles = s.UrlFilteringSecurityProfiles
	o.VulnerabilityProtectionSecurityProfiles = s.VulnerabilityProtectionSecurityProfiles
	o.WildfireAnalysisSecurityProfiles = s.WildfireAnalysisSecurityProfiles
	o.Misc = s.Misc
}

func (o roleVsysRestapiObjectsXml) UnmarshalToObject() (*RoleVsysRestapiObjects, error) {

	result := &RoleVsysRestapiObjects{
		AddressGroups:                           o.AddressGroups,
		Addresses:                               o.Addresses,
		AntiSpywareSecurityProfiles:             o.AntiSpywareSecurityProfiles,
		AntivirusSecurityProfiles:               o.AntivirusSecurityProfiles,
		ApplicationFilters:                      o.ApplicationFilters,
		ApplicationGroups:                       o.ApplicationGroups,
		Applications:                            o.Applications,
		AuthenticationEnforcements:              o.AuthenticationEnforcements,
		CustomDataPatterns:                      o.CustomDataPatterns,
		CustomSpywareSignatures:                 o.CustomSpywareSignatures,
		CustomUrlCategories:                     o.CustomUrlCategories,
		CustomVulnerabilitySignatures:           o.CustomVulnerabilitySignatures,
		DataFilteringSecurityProfiles:           o.DataFilteringSecurityProfiles,
		DecryptionProfiles:                      o.DecryptionProfiles,
		Devices:                                 o.Devices,
		DosProtectionSecurityProfiles:           o.DosProtectionSecurityProfiles,
		DynamicUserGroups:                       o.DynamicUserGroups,
		ExternalDynamicLists:                    o.ExternalDynamicLists,
		FileBlockingSecurityProfiles:            o.FileBlockingSecurityProfiles,
		GlobalprotectHipObjects:                 o.GlobalprotectHipObjects,
		GlobalprotectHipProfiles:                o.GlobalprotectHipProfiles,
		GtpProtectionSecurityProfiles:           o.GtpProtectionSecurityProfiles,
		LogForwardingProfiles:                   o.LogForwardingProfiles,
		PacketBrokerProfiles:                    o.PacketBrokerProfiles,
		Regions:                                 o.Regions,
		Schedules:                               o.Schedules,
		SctpProtectionSecurityProfiles:          o.SctpProtectionSecurityProfiles,
		SdwanErrorCorrectionProfiles:            o.SdwanErrorCorrectionProfiles,
		SdwanPathQualityProfiles:                o.SdwanPathQualityProfiles,
		SdwanSaasQualityProfiles:                o.SdwanSaasQualityProfiles,
		SdwanTrafficDistributionProfiles:        o.SdwanTrafficDistributionProfiles,
		SecurityProfileGroups:                   o.SecurityProfileGroups,
		ServiceGroups:                           o.ServiceGroups,
		Services:                                o.Services,
		Tags:                                    o.Tags,
		UrlFilteringSecurityProfiles:            o.UrlFilteringSecurityProfiles,
		VulnerabilityProtectionSecurityProfiles: o.VulnerabilityProtectionSecurityProfiles,
		WildfireAnalysisSecurityProfiles:        o.WildfireAnalysisSecurityProfiles,
		Misc:                                    o.Misc,
	}
	return result, nil
}
func (o *roleVsysRestapiPoliciesXml) MarshalFromObject(s RoleVsysRestapiPolicies) {
	o.ApplicationOverrideRules = s.ApplicationOverrideRules
	o.AuthenticationRules = s.AuthenticationRules
	o.DecryptionRules = s.DecryptionRules
	o.DosRules = s.DosRules
	o.NatRules = s.NatRules
	o.NetworkPacketBrokerRules = s.NetworkPacketBrokerRules
	o.PolicyBasedForwardingRules = s.PolicyBasedForwardingRules
	o.QosRules = s.QosRules
	o.SdwanRules = s.SdwanRules
	o.SecurityRules = s.SecurityRules
	o.TunnelInspectionRules = s.TunnelInspectionRules
	o.Misc = s.Misc
}

func (o roleVsysRestapiPoliciesXml) UnmarshalToObject() (*RoleVsysRestapiPolicies, error) {

	result := &RoleVsysRestapiPolicies{
		ApplicationOverrideRules:   o.ApplicationOverrideRules,
		AuthenticationRules:        o.AuthenticationRules,
		DecryptionRules:            o.DecryptionRules,
		DosRules:                   o.DosRules,
		NatRules:                   o.NatRules,
		NetworkPacketBrokerRules:   o.NetworkPacketBrokerRules,
		PolicyBasedForwardingRules: o.PolicyBasedForwardingRules,
		QosRules:                   o.QosRules,
		SdwanRules:                 o.SdwanRules,
		SecurityRules:              o.SecurityRules,
		TunnelInspectionRules:      o.TunnelInspectionRules,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *roleVsysRestapiSystemXml) MarshalFromObject(s RoleVsysRestapiSystem) {
	o.Configuration = s.Configuration
	o.Misc = s.Misc
}

func (o roleVsysRestapiSystemXml) UnmarshalToObject() (*RoleVsysRestapiSystem, error) {

	result := &RoleVsysRestapiSystem{
		Configuration: o.Configuration,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiXml) MarshalFromObject(s RoleVsysWebui) {
	o.Acc = s.Acc
	if s.Commit != nil {
		var obj roleVsysWebuiCommitXml
		obj.MarshalFromObject(*s.Commit)
		o.Commit = &obj
	}
	o.Dashboard = s.Dashboard
	if s.Device != nil {
		var obj roleVsysWebuiDeviceXml
		obj.MarshalFromObject(*s.Device)
		o.Device = &obj
	}
	if s.Monitor != nil {
		var obj roleVsysWebuiMonitorXml
		obj.MarshalFromObject(*s.Monitor)
		o.Monitor = &obj
	}
	if s.Network != nil {
		var obj roleVsysWebuiNetworkXml
		obj.MarshalFromObject(*s.Network)
		o.Network = &obj
	}
	if s.Objects != nil {
		var obj roleVsysWebuiObjectsXml
		obj.MarshalFromObject(*s.Objects)
		o.Objects = &obj
	}
	if s.Operations != nil {
		var obj roleVsysWebuiOperationsXml
		obj.MarshalFromObject(*s.Operations)
		o.Operations = &obj
	}
	if s.Policies != nil {
		var obj roleVsysWebuiPoliciesXml
		obj.MarshalFromObject(*s.Policies)
		o.Policies = &obj
	}
	if s.Privacy != nil {
		var obj roleVsysWebuiPrivacyXml
		obj.MarshalFromObject(*s.Privacy)
		o.Privacy = &obj
	}
	if s.Save != nil {
		var obj roleVsysWebuiSaveXml
		obj.MarshalFromObject(*s.Save)
		o.Save = &obj
	}
	o.Tasks = s.Tasks
	o.Validate = s.Validate
	o.Misc = s.Misc
}

func (o roleVsysWebuiXml) UnmarshalToObject() (*RoleVsysWebui, error) {
	var commitVal *RoleVsysWebuiCommit
	if o.Commit != nil {
		obj, err := o.Commit.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		commitVal = obj
	}
	var deviceVal *RoleVsysWebuiDevice
	if o.Device != nil {
		obj, err := o.Device.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		deviceVal = obj
	}
	var monitorVal *RoleVsysWebuiMonitor
	if o.Monitor != nil {
		obj, err := o.Monitor.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		monitorVal = obj
	}
	var networkVal *RoleVsysWebuiNetwork
	if o.Network != nil {
		obj, err := o.Network.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		networkVal = obj
	}
	var objectsVal *RoleVsysWebuiObjects
	if o.Objects != nil {
		obj, err := o.Objects.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		objectsVal = obj
	}
	var operationsVal *RoleVsysWebuiOperations
	if o.Operations != nil {
		obj, err := o.Operations.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		operationsVal = obj
	}
	var policiesVal *RoleVsysWebuiPolicies
	if o.Policies != nil {
		obj, err := o.Policies.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		policiesVal = obj
	}
	var privacyVal *RoleVsysWebuiPrivacy
	if o.Privacy != nil {
		obj, err := o.Privacy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		privacyVal = obj
	}
	var saveVal *RoleVsysWebuiSave
	if o.Save != nil {
		obj, err := o.Save.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		saveVal = obj
	}

	result := &RoleVsysWebui{
		Acc:        o.Acc,
		Commit:     commitVal,
		Dashboard:  o.Dashboard,
		Device:     deviceVal,
		Monitor:    monitorVal,
		Network:    networkVal,
		Objects:    objectsVal,
		Operations: operationsVal,
		Policies:   policiesVal,
		Privacy:    privacyVal,
		Save:       saveVal,
		Tasks:      o.Tasks,
		Validate:   o.Validate,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiCommitXml) MarshalFromObject(s RoleVsysWebuiCommit) {
	o.CommitForOtherAdmins = s.CommitForOtherAdmins
	o.VirtualSystems = s.VirtualSystems
	o.Misc = s.Misc
}

func (o roleVsysWebuiCommitXml) UnmarshalToObject() (*RoleVsysWebuiCommit, error) {

	result := &RoleVsysWebuiCommit{
		CommitForOtherAdmins: o.CommitForOtherAdmins,
		VirtualSystems:       o.VirtualSystems,
		Misc:                 o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiDeviceXml) MarshalFromObject(s RoleVsysWebuiDevice) {
	o.Administrators = s.Administrators
	o.AuthenticationProfile = s.AuthenticationProfile
	o.AuthenticationSequence = s.AuthenticationSequence
	o.BlockPages = s.BlockPages
	if s.CertificateManagement != nil {
		var obj roleVsysWebuiDeviceCertificateManagementXml
		obj.MarshalFromObject(*s.CertificateManagement)
		o.CertificateManagement = &obj
	}
	o.DataRedistribution = s.DataRedistribution
	o.DeviceQuarantine = s.DeviceQuarantine
	o.DhcpSyslogServer = s.DhcpSyslogServer
	if s.LocalUserDatabase != nil {
		var obj roleVsysWebuiDeviceLocalUserDatabaseXml
		obj.MarshalFromObject(*s.LocalUserDatabase)
		o.LocalUserDatabase = &obj
	}
	if s.LogSettings != nil {
		var obj roleVsysWebuiDeviceLogSettingsXml
		obj.MarshalFromObject(*s.LogSettings)
		o.LogSettings = &obj
	}
	if s.PolicyRecommendations != nil {
		var obj roleVsysWebuiDevicePolicyRecommendationsXml
		obj.MarshalFromObject(*s.PolicyRecommendations)
		o.PolicyRecommendations = &obj
	}
	if s.ServerProfile != nil {
		var obj roleVsysWebuiDeviceServerProfileXml
		obj.MarshalFromObject(*s.ServerProfile)
		o.ServerProfile = &obj
	}
	if s.Setup != nil {
		var obj roleVsysWebuiDeviceSetupXml
		obj.MarshalFromObject(*s.Setup)
		o.Setup = &obj
	}
	o.Troubleshooting = s.Troubleshooting
	o.UserIdentification = s.UserIdentification
	o.VmInfoSource = s.VmInfoSource
	o.Misc = s.Misc
}

func (o roleVsysWebuiDeviceXml) UnmarshalToObject() (*RoleVsysWebuiDevice, error) {
	var certificateManagementVal *RoleVsysWebuiDeviceCertificateManagement
	if o.CertificateManagement != nil {
		obj, err := o.CertificateManagement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		certificateManagementVal = obj
	}
	var localUserDatabaseVal *RoleVsysWebuiDeviceLocalUserDatabase
	if o.LocalUserDatabase != nil {
		obj, err := o.LocalUserDatabase.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localUserDatabaseVal = obj
	}
	var logSettingsVal *RoleVsysWebuiDeviceLogSettings
	if o.LogSettings != nil {
		obj, err := o.LogSettings.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		logSettingsVal = obj
	}
	var policyRecommendationsVal *RoleVsysWebuiDevicePolicyRecommendations
	if o.PolicyRecommendations != nil {
		obj, err := o.PolicyRecommendations.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		policyRecommendationsVal = obj
	}
	var serverProfileVal *RoleVsysWebuiDeviceServerProfile
	if o.ServerProfile != nil {
		obj, err := o.ServerProfile.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		serverProfileVal = obj
	}
	var setupVal *RoleVsysWebuiDeviceSetup
	if o.Setup != nil {
		obj, err := o.Setup.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setupVal = obj
	}

	result := &RoleVsysWebuiDevice{
		Administrators:         o.Administrators,
		AuthenticationProfile:  o.AuthenticationProfile,
		AuthenticationSequence: o.AuthenticationSequence,
		BlockPages:             o.BlockPages,
		CertificateManagement:  certificateManagementVal,
		DataRedistribution:     o.DataRedistribution,
		DeviceQuarantine:       o.DeviceQuarantine,
		DhcpSyslogServer:       o.DhcpSyslogServer,
		LocalUserDatabase:      localUserDatabaseVal,
		LogSettings:            logSettingsVal,
		PolicyRecommendations:  policyRecommendationsVal,
		ServerProfile:          serverProfileVal,
		Setup:                  setupVal,
		Troubleshooting:        o.Troubleshooting,
		UserIdentification:     o.UserIdentification,
		VmInfoSource:           o.VmInfoSource,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiDeviceCertificateManagementXml) MarshalFromObject(s RoleVsysWebuiDeviceCertificateManagement) {
	o.CertificateProfile = s.CertificateProfile
	o.Certificates = s.Certificates
	o.OcspResponder = s.OcspResponder
	o.Scep = s.Scep
	o.SshServiceProfile = s.SshServiceProfile
	o.SslDecryptionExclusion = s.SslDecryptionExclusion
	o.SslTlsServiceProfile = s.SslTlsServiceProfile
	o.Misc = s.Misc
}

func (o roleVsysWebuiDeviceCertificateManagementXml) UnmarshalToObject() (*RoleVsysWebuiDeviceCertificateManagement, error) {

	result := &RoleVsysWebuiDeviceCertificateManagement{
		CertificateProfile:     o.CertificateProfile,
		Certificates:           o.Certificates,
		OcspResponder:          o.OcspResponder,
		Scep:                   o.Scep,
		SshServiceProfile:      o.SshServiceProfile,
		SslDecryptionExclusion: o.SslDecryptionExclusion,
		SslTlsServiceProfile:   o.SslTlsServiceProfile,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiDeviceLocalUserDatabaseXml) MarshalFromObject(s RoleVsysWebuiDeviceLocalUserDatabase) {
	o.UserGroups = s.UserGroups
	o.Users = s.Users
	o.Misc = s.Misc
}

func (o roleVsysWebuiDeviceLocalUserDatabaseXml) UnmarshalToObject() (*RoleVsysWebuiDeviceLocalUserDatabase, error) {

	result := &RoleVsysWebuiDeviceLocalUserDatabase{
		UserGroups: o.UserGroups,
		Users:      o.Users,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiDeviceLogSettingsXml) MarshalFromObject(s RoleVsysWebuiDeviceLogSettings) {
	o.Config = s.Config
	o.Correlation = s.Correlation
	o.Globalprotect = s.Globalprotect
	o.Hipmatch = s.Hipmatch
	o.Iptag = s.Iptag
	o.System = s.System
	o.UserId = s.UserId
	o.Misc = s.Misc
}

func (o roleVsysWebuiDeviceLogSettingsXml) UnmarshalToObject() (*RoleVsysWebuiDeviceLogSettings, error) {

	result := &RoleVsysWebuiDeviceLogSettings{
		Config:        o.Config,
		Correlation:   o.Correlation,
		Globalprotect: o.Globalprotect,
		Hipmatch:      o.Hipmatch,
		Iptag:         o.Iptag,
		System:        o.System,
		UserId:        o.UserId,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiDevicePolicyRecommendationsXml) MarshalFromObject(s RoleVsysWebuiDevicePolicyRecommendations) {
	o.Iot = s.Iot
	o.Saas = s.Saas
	o.Misc = s.Misc
}

func (o roleVsysWebuiDevicePolicyRecommendationsXml) UnmarshalToObject() (*RoleVsysWebuiDevicePolicyRecommendations, error) {

	result := &RoleVsysWebuiDevicePolicyRecommendations{
		Iot:  o.Iot,
		Saas: o.Saas,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiDeviceServerProfileXml) MarshalFromObject(s RoleVsysWebuiDeviceServerProfile) {
	o.Dns = s.Dns
	o.Email = s.Email
	o.Http = s.Http
	o.Kerberos = s.Kerberos
	o.Ldap = s.Ldap
	o.Mfa = s.Mfa
	o.Netflow = s.Netflow
	o.Radius = s.Radius
	o.SamlIdp = s.SamlIdp
	o.Scp = s.Scp
	o.SnmpTrap = s.SnmpTrap
	o.Syslog = s.Syslog
	o.Tacplus = s.Tacplus
	o.Misc = s.Misc
}

func (o roleVsysWebuiDeviceServerProfileXml) UnmarshalToObject() (*RoleVsysWebuiDeviceServerProfile, error) {

	result := &RoleVsysWebuiDeviceServerProfile{
		Dns:      o.Dns,
		Email:    o.Email,
		Http:     o.Http,
		Kerberos: o.Kerberos,
		Ldap:     o.Ldap,
		Mfa:      o.Mfa,
		Netflow:  o.Netflow,
		Radius:   o.Radius,
		SamlIdp:  o.SamlIdp,
		Scp:      o.Scp,
		SnmpTrap: o.SnmpTrap,
		Syslog:   o.Syslog,
		Tacplus:  o.Tacplus,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiDeviceSetupXml) MarshalFromObject(s RoleVsysWebuiDeviceSetup) {
	o.ContentId = s.ContentId
	o.Hsm = s.Hsm
	o.Interfaces = s.Interfaces
	o.Management = s.Management
	o.Operations = s.Operations
	o.Services = s.Services
	o.Session = s.Session
	o.Telemetry = s.Telemetry
	o.Wildfire = s.Wildfire
	o.Misc = s.Misc
}

func (o roleVsysWebuiDeviceSetupXml) UnmarshalToObject() (*RoleVsysWebuiDeviceSetup, error) {

	result := &RoleVsysWebuiDeviceSetup{
		ContentId:  o.ContentId,
		Hsm:        o.Hsm,
		Interfaces: o.Interfaces,
		Management: o.Management,
		Operations: o.Operations,
		Services:   o.Services,
		Session:    o.Session,
		Telemetry:  o.Telemetry,
		Wildfire:   o.Wildfire,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiMonitorXml) MarshalFromObject(s RoleVsysWebuiMonitor) {
	o.AppScope = s.AppScope
	if s.AutomatedCorrelationEngine != nil {
		var obj roleVsysWebuiMonitorAutomatedCorrelationEngineXml
		obj.MarshalFromObject(*s.AutomatedCorrelationEngine)
		o.AutomatedCorrelationEngine = &obj
	}
	o.BlockIpList = s.BlockIpList
	if s.CustomReports != nil {
		var obj roleVsysWebuiMonitorCustomReportsXml
		obj.MarshalFromObject(*s.CustomReports)
		o.CustomReports = &obj
	}
	o.ExternalLogs = s.ExternalLogs
	if s.Logs != nil {
		var obj roleVsysWebuiMonitorLogsXml
		obj.MarshalFromObject(*s.Logs)
		o.Logs = &obj
	}
	if s.PdfReports != nil {
		var obj roleVsysWebuiMonitorPdfReportsXml
		obj.MarshalFromObject(*s.PdfReports)
		o.PdfReports = &obj
	}
	o.SessionBrowser = s.SessionBrowser
	o.ViewCustomReports = s.ViewCustomReports
	o.Misc = s.Misc
}

func (o roleVsysWebuiMonitorXml) UnmarshalToObject() (*RoleVsysWebuiMonitor, error) {
	var automatedCorrelationEngineVal *RoleVsysWebuiMonitorAutomatedCorrelationEngine
	if o.AutomatedCorrelationEngine != nil {
		obj, err := o.AutomatedCorrelationEngine.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		automatedCorrelationEngineVal = obj
	}
	var customReportsVal *RoleVsysWebuiMonitorCustomReports
	if o.CustomReports != nil {
		obj, err := o.CustomReports.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		customReportsVal = obj
	}
	var logsVal *RoleVsysWebuiMonitorLogs
	if o.Logs != nil {
		obj, err := o.Logs.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		logsVal = obj
	}
	var pdfReportsVal *RoleVsysWebuiMonitorPdfReports
	if o.PdfReports != nil {
		obj, err := o.PdfReports.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		pdfReportsVal = obj
	}

	result := &RoleVsysWebuiMonitor{
		AppScope:                   o.AppScope,
		AutomatedCorrelationEngine: automatedCorrelationEngineVal,
		BlockIpList:                o.BlockIpList,
		CustomReports:              customReportsVal,
		ExternalLogs:               o.ExternalLogs,
		Logs:                       logsVal,
		PdfReports:                 pdfReportsVal,
		SessionBrowser:             o.SessionBrowser,
		ViewCustomReports:          o.ViewCustomReports,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiMonitorAutomatedCorrelationEngineXml) MarshalFromObject(s RoleVsysWebuiMonitorAutomatedCorrelationEngine) {
	o.CorrelatedEvents = s.CorrelatedEvents
	o.CorrelationObjects = s.CorrelationObjects
	o.Misc = s.Misc
}

func (o roleVsysWebuiMonitorAutomatedCorrelationEngineXml) UnmarshalToObject() (*RoleVsysWebuiMonitorAutomatedCorrelationEngine, error) {

	result := &RoleVsysWebuiMonitorAutomatedCorrelationEngine{
		CorrelatedEvents:   o.CorrelatedEvents,
		CorrelationObjects: o.CorrelationObjects,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiMonitorCustomReportsXml) MarshalFromObject(s RoleVsysWebuiMonitorCustomReports) {
	o.ApplicationStatistics = s.ApplicationStatistics
	o.Auth = s.Auth
	o.DataFilteringLog = s.DataFilteringLog
	o.DecryptionLog = s.DecryptionLog
	o.DecryptionSummary = s.DecryptionSummary
	o.Globalprotect = s.Globalprotect
	o.GtpLog = s.GtpLog
	o.GtpSummary = s.GtpSummary
	o.Hipmatch = s.Hipmatch
	o.Iptag = s.Iptag
	o.SctpLog = s.SctpLog
	o.SctpSummary = s.SctpSummary
	o.ThreatLog = s.ThreatLog
	o.ThreatSummary = s.ThreatSummary
	o.TrafficLog = s.TrafficLog
	o.TrafficSummary = s.TrafficSummary
	o.TunnelLog = s.TunnelLog
	o.TunnelSummary = s.TunnelSummary
	o.UrlLog = s.UrlLog
	o.UrlSummary = s.UrlSummary
	o.Userid = s.Userid
	o.WildfireLog = s.WildfireLog
	o.Misc = s.Misc
}

func (o roleVsysWebuiMonitorCustomReportsXml) UnmarshalToObject() (*RoleVsysWebuiMonitorCustomReports, error) {

	result := &RoleVsysWebuiMonitorCustomReports{
		ApplicationStatistics: o.ApplicationStatistics,
		Auth:                  o.Auth,
		DataFilteringLog:      o.DataFilteringLog,
		DecryptionLog:         o.DecryptionLog,
		DecryptionSummary:     o.DecryptionSummary,
		Globalprotect:         o.Globalprotect,
		GtpLog:                o.GtpLog,
		GtpSummary:            o.GtpSummary,
		Hipmatch:              o.Hipmatch,
		Iptag:                 o.Iptag,
		SctpLog:               o.SctpLog,
		SctpSummary:           o.SctpSummary,
		ThreatLog:             o.ThreatLog,
		ThreatSummary:         o.ThreatSummary,
		TrafficLog:            o.TrafficLog,
		TrafficSummary:        o.TrafficSummary,
		TunnelLog:             o.TunnelLog,
		TunnelSummary:         o.TunnelSummary,
		UrlLog:                o.UrlLog,
		UrlSummary:            o.UrlSummary,
		Userid:                o.Userid,
		WildfireLog:           o.WildfireLog,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiMonitorLogsXml) MarshalFromObject(s RoleVsysWebuiMonitorLogs) {
	o.Authentication = s.Authentication
	o.DataFiltering = s.DataFiltering
	o.Decryption = s.Decryption
	o.Globalprotect = s.Globalprotect
	o.Gtp = s.Gtp
	o.Hipmatch = s.Hipmatch
	o.Iptag = s.Iptag
	o.Sctp = s.Sctp
	o.Threat = s.Threat
	o.Traffic = s.Traffic
	o.Tunnel = s.Tunnel
	o.Url = s.Url
	o.Userid = s.Userid
	o.Wildfire = s.Wildfire
	o.Misc = s.Misc
}

func (o roleVsysWebuiMonitorLogsXml) UnmarshalToObject() (*RoleVsysWebuiMonitorLogs, error) {

	result := &RoleVsysWebuiMonitorLogs{
		Authentication: o.Authentication,
		DataFiltering:  o.DataFiltering,
		Decryption:     o.Decryption,
		Globalprotect:  o.Globalprotect,
		Gtp:            o.Gtp,
		Hipmatch:       o.Hipmatch,
		Iptag:          o.Iptag,
		Sctp:           o.Sctp,
		Threat:         o.Threat,
		Traffic:        o.Traffic,
		Tunnel:         o.Tunnel,
		Url:            o.Url,
		Userid:         o.Userid,
		Wildfire:       o.Wildfire,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiMonitorPdfReportsXml) MarshalFromObject(s RoleVsysWebuiMonitorPdfReports) {
	o.EmailScheduler = s.EmailScheduler
	o.ManagePdfSummary = s.ManagePdfSummary
	o.PdfSummaryReports = s.PdfSummaryReports
	o.ReportGroups = s.ReportGroups
	o.SaasApplicationUsageReport = s.SaasApplicationUsageReport
	o.UserActivityReport = s.UserActivityReport
	o.Misc = s.Misc
}

func (o roleVsysWebuiMonitorPdfReportsXml) UnmarshalToObject() (*RoleVsysWebuiMonitorPdfReports, error) {

	result := &RoleVsysWebuiMonitorPdfReports{
		EmailScheduler:             o.EmailScheduler,
		ManagePdfSummary:           o.ManagePdfSummary,
		PdfSummaryReports:          o.PdfSummaryReports,
		ReportGroups:               o.ReportGroups,
		SaasApplicationUsageReport: o.SaasApplicationUsageReport,
		UserActivityReport:         o.UserActivityReport,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiNetworkXml) MarshalFromObject(s RoleVsysWebuiNetwork) {
	if s.GlobalProtect != nil {
		var obj roleVsysWebuiNetworkGlobalProtectXml
		obj.MarshalFromObject(*s.GlobalProtect)
		o.GlobalProtect = &obj
	}
	o.SdwanInterfaceProfile = s.SdwanInterfaceProfile
	o.Zones = s.Zones
	o.Misc = s.Misc
}

func (o roleVsysWebuiNetworkXml) UnmarshalToObject() (*RoleVsysWebuiNetwork, error) {
	var globalProtectVal *RoleVsysWebuiNetworkGlobalProtect
	if o.GlobalProtect != nil {
		obj, err := o.GlobalProtect.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalProtectVal = obj
	}

	result := &RoleVsysWebuiNetwork{
		GlobalProtect:         globalProtectVal,
		SdwanInterfaceProfile: o.SdwanInterfaceProfile,
		Zones:                 o.Zones,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiNetworkGlobalProtectXml) MarshalFromObject(s RoleVsysWebuiNetworkGlobalProtect) {
	o.ClientlessAppGroups = s.ClientlessAppGroups
	o.ClientlessApps = s.ClientlessApps
	o.Gateways = s.Gateways
	o.Mdm = s.Mdm
	o.Portals = s.Portals
	o.Misc = s.Misc
}

func (o roleVsysWebuiNetworkGlobalProtectXml) UnmarshalToObject() (*RoleVsysWebuiNetworkGlobalProtect, error) {

	result := &RoleVsysWebuiNetworkGlobalProtect{
		ClientlessAppGroups: o.ClientlessAppGroups,
		ClientlessApps:      o.ClientlessApps,
		Gateways:            o.Gateways,
		Mdm:                 o.Mdm,
		Portals:             o.Portals,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiObjectsXml) MarshalFromObject(s RoleVsysWebuiObjects) {
	o.AddressGroups = s.AddressGroups
	o.Addresses = s.Addresses
	o.ApplicationFilters = s.ApplicationFilters
	o.ApplicationGroups = s.ApplicationGroups
	o.Applications = s.Applications
	o.Authentication = s.Authentication
	if s.CustomObjects != nil {
		var obj roleVsysWebuiObjectsCustomObjectsXml
		obj.MarshalFromObject(*s.CustomObjects)
		o.CustomObjects = &obj
	}
	if s.Decryption != nil {
		var obj roleVsysWebuiObjectsDecryptionXml
		obj.MarshalFromObject(*s.Decryption)
		o.Decryption = &obj
	}
	o.Devices = s.Devices
	o.DynamicBlockLists = s.DynamicBlockLists
	o.DynamicUserGroups = s.DynamicUserGroups
	if s.GlobalProtect != nil {
		var obj roleVsysWebuiObjectsGlobalProtectXml
		obj.MarshalFromObject(*s.GlobalProtect)
		o.GlobalProtect = &obj
	}
	o.LogForwarding = s.LogForwarding
	o.PacketBrokerProfile = s.PacketBrokerProfile
	o.Regions = s.Regions
	o.Schedules = s.Schedules
	if s.Sdwan != nil {
		var obj roleVsysWebuiObjectsSdwanXml
		obj.MarshalFromObject(*s.Sdwan)
		o.Sdwan = &obj
	}
	o.SecurityProfileGroups = s.SecurityProfileGroups
	if s.SecurityProfiles != nil {
		var obj roleVsysWebuiObjectsSecurityProfilesXml
		obj.MarshalFromObject(*s.SecurityProfiles)
		o.SecurityProfiles = &obj
	}
	o.ServiceGroups = s.ServiceGroups
	o.Services = s.Services
	o.Tags = s.Tags
	o.Misc = s.Misc
}

func (o roleVsysWebuiObjectsXml) UnmarshalToObject() (*RoleVsysWebuiObjects, error) {
	var customObjectsVal *RoleVsysWebuiObjectsCustomObjects
	if o.CustomObjects != nil {
		obj, err := o.CustomObjects.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		customObjectsVal = obj
	}
	var decryptionVal *RoleVsysWebuiObjectsDecryption
	if o.Decryption != nil {
		obj, err := o.Decryption.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		decryptionVal = obj
	}
	var globalProtectVal *RoleVsysWebuiObjectsGlobalProtect
	if o.GlobalProtect != nil {
		obj, err := o.GlobalProtect.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalProtectVal = obj
	}
	var sdwanVal *RoleVsysWebuiObjectsSdwan
	if o.Sdwan != nil {
		obj, err := o.Sdwan.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sdwanVal = obj
	}
	var securityProfilesVal *RoleVsysWebuiObjectsSecurityProfiles
	if o.SecurityProfiles != nil {
		obj, err := o.SecurityProfiles.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		securityProfilesVal = obj
	}

	result := &RoleVsysWebuiObjects{
		AddressGroups:         o.AddressGroups,
		Addresses:             o.Addresses,
		ApplicationFilters:    o.ApplicationFilters,
		ApplicationGroups:     o.ApplicationGroups,
		Applications:          o.Applications,
		Authentication:        o.Authentication,
		CustomObjects:         customObjectsVal,
		Decryption:            decryptionVal,
		Devices:               o.Devices,
		DynamicBlockLists:     o.DynamicBlockLists,
		DynamicUserGroups:     o.DynamicUserGroups,
		GlobalProtect:         globalProtectVal,
		LogForwarding:         o.LogForwarding,
		PacketBrokerProfile:   o.PacketBrokerProfile,
		Regions:               o.Regions,
		Schedules:             o.Schedules,
		Sdwan:                 sdwanVal,
		SecurityProfileGroups: o.SecurityProfileGroups,
		SecurityProfiles:      securityProfilesVal,
		ServiceGroups:         o.ServiceGroups,
		Services:              o.Services,
		Tags:                  o.Tags,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiObjectsCustomObjectsXml) MarshalFromObject(s RoleVsysWebuiObjectsCustomObjects) {
	o.DataPatterns = s.DataPatterns
	o.Spyware = s.Spyware
	o.UrlCategory = s.UrlCategory
	o.Vulnerability = s.Vulnerability
	o.Misc = s.Misc
}

func (o roleVsysWebuiObjectsCustomObjectsXml) UnmarshalToObject() (*RoleVsysWebuiObjectsCustomObjects, error) {

	result := &RoleVsysWebuiObjectsCustomObjects{
		DataPatterns:  o.DataPatterns,
		Spyware:       o.Spyware,
		UrlCategory:   o.UrlCategory,
		Vulnerability: o.Vulnerability,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiObjectsDecryptionXml) MarshalFromObject(s RoleVsysWebuiObjectsDecryption) {
	o.DecryptionProfile = s.DecryptionProfile
	o.Misc = s.Misc
}

func (o roleVsysWebuiObjectsDecryptionXml) UnmarshalToObject() (*RoleVsysWebuiObjectsDecryption, error) {

	result := &RoleVsysWebuiObjectsDecryption{
		DecryptionProfile: o.DecryptionProfile,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiObjectsGlobalProtectXml) MarshalFromObject(s RoleVsysWebuiObjectsGlobalProtect) {
	o.HipObjects = s.HipObjects
	o.HipProfiles = s.HipProfiles
	o.Misc = s.Misc
}

func (o roleVsysWebuiObjectsGlobalProtectXml) UnmarshalToObject() (*RoleVsysWebuiObjectsGlobalProtect, error) {

	result := &RoleVsysWebuiObjectsGlobalProtect{
		HipObjects:  o.HipObjects,
		HipProfiles: o.HipProfiles,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiObjectsSdwanXml) MarshalFromObject(s RoleVsysWebuiObjectsSdwan) {
	o.SdwanDistProfile = s.SdwanDistProfile
	o.SdwanErrorCorrectionProfile = s.SdwanErrorCorrectionProfile
	o.SdwanProfile = s.SdwanProfile
	o.SdwanSaasQualityProfile = s.SdwanSaasQualityProfile
	o.Misc = s.Misc
}

func (o roleVsysWebuiObjectsSdwanXml) UnmarshalToObject() (*RoleVsysWebuiObjectsSdwan, error) {

	result := &RoleVsysWebuiObjectsSdwan{
		SdwanDistProfile:            o.SdwanDistProfile,
		SdwanErrorCorrectionProfile: o.SdwanErrorCorrectionProfile,
		SdwanProfile:                o.SdwanProfile,
		SdwanSaasQualityProfile:     o.SdwanSaasQualityProfile,
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiObjectsSecurityProfilesXml) MarshalFromObject(s RoleVsysWebuiObjectsSecurityProfiles) {
	o.AntiSpyware = s.AntiSpyware
	o.Antivirus = s.Antivirus
	o.DataFiltering = s.DataFiltering
	o.DosProtection = s.DosProtection
	o.FileBlocking = s.FileBlocking
	o.GtpProtection = s.GtpProtection
	o.SctpProtection = s.SctpProtection
	o.UrlFiltering = s.UrlFiltering
	o.VulnerabilityProtection = s.VulnerabilityProtection
	o.WildfireAnalysis = s.WildfireAnalysis
	o.Misc = s.Misc
}

func (o roleVsysWebuiObjectsSecurityProfilesXml) UnmarshalToObject() (*RoleVsysWebuiObjectsSecurityProfiles, error) {

	result := &RoleVsysWebuiObjectsSecurityProfiles{
		AntiSpyware:             o.AntiSpyware,
		Antivirus:               o.Antivirus,
		DataFiltering:           o.DataFiltering,
		DosProtection:           o.DosProtection,
		FileBlocking:            o.FileBlocking,
		GtpProtection:           o.GtpProtection,
		SctpProtection:          o.SctpProtection,
		UrlFiltering:            o.UrlFiltering,
		VulnerabilityProtection: o.VulnerabilityProtection,
		WildfireAnalysis:        o.WildfireAnalysis,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiOperationsXml) MarshalFromObject(s RoleVsysWebuiOperations) {
	o.DownloadCoreFiles = s.DownloadCoreFiles
	o.DownloadPcapFiles = s.DownloadPcapFiles
	o.GenerateStatsDumpFile = s.GenerateStatsDumpFile
	o.GenerateTechSupportFile = s.GenerateTechSupportFile
	o.Reboot = s.Reboot
	o.Misc = s.Misc
}

func (o roleVsysWebuiOperationsXml) UnmarshalToObject() (*RoleVsysWebuiOperations, error) {

	result := &RoleVsysWebuiOperations{
		DownloadCoreFiles:       o.DownloadCoreFiles,
		DownloadPcapFiles:       o.DownloadPcapFiles,
		GenerateStatsDumpFile:   o.GenerateStatsDumpFile,
		GenerateTechSupportFile: o.GenerateTechSupportFile,
		Reboot:                  o.Reboot,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiPoliciesXml) MarshalFromObject(s RoleVsysWebuiPolicies) {
	o.ApplicationOverrideRulebase = s.ApplicationOverrideRulebase
	o.AuthenticationRulebase = s.AuthenticationRulebase
	o.DosRulebase = s.DosRulebase
	o.NatRulebase = s.NatRulebase
	o.NetworkPacketBrokerRulebase = s.NetworkPacketBrokerRulebase
	o.PbfRulebase = s.PbfRulebase
	o.QosRulebase = s.QosRulebase
	o.RuleHitCountReset = s.RuleHitCountReset
	o.SdwanRulebase = s.SdwanRulebase
	o.SecurityRulebase = s.SecurityRulebase
	o.SslDecryptionRulebase = s.SslDecryptionRulebase
	o.TunnelInspectRulebase = s.TunnelInspectRulebase
	o.Misc = s.Misc
}

func (o roleVsysWebuiPoliciesXml) UnmarshalToObject() (*RoleVsysWebuiPolicies, error) {

	result := &RoleVsysWebuiPolicies{
		ApplicationOverrideRulebase: o.ApplicationOverrideRulebase,
		AuthenticationRulebase:      o.AuthenticationRulebase,
		DosRulebase:                 o.DosRulebase,
		NatRulebase:                 o.NatRulebase,
		NetworkPacketBrokerRulebase: o.NetworkPacketBrokerRulebase,
		PbfRulebase:                 o.PbfRulebase,
		QosRulebase:                 o.QosRulebase,
		RuleHitCountReset:           o.RuleHitCountReset,
		SdwanRulebase:               o.SdwanRulebase,
		SecurityRulebase:            o.SecurityRulebase,
		SslDecryptionRulebase:       o.SslDecryptionRulebase,
		TunnelInspectRulebase:       o.TunnelInspectRulebase,
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiPrivacyXml) MarshalFromObject(s RoleVsysWebuiPrivacy) {
	o.ShowFullIpAddresses = s.ShowFullIpAddresses
	o.ShowUserNamesInLogsAndReports = s.ShowUserNamesInLogsAndReports
	o.ViewPcapFiles = s.ViewPcapFiles
	o.Misc = s.Misc
}

func (o roleVsysWebuiPrivacyXml) UnmarshalToObject() (*RoleVsysWebuiPrivacy, error) {

	result := &RoleVsysWebuiPrivacy{
		ShowFullIpAddresses:           o.ShowFullIpAddresses,
		ShowUserNamesInLogsAndReports: o.ShowUserNamesInLogsAndReports,
		ViewPcapFiles:                 o.ViewPcapFiles,
		Misc:                          o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiSaveXml) MarshalFromObject(s RoleVsysWebuiSave) {
	o.ObjectLevelChanges = s.ObjectLevelChanges
	o.PartialSave = s.PartialSave
	o.SaveForOtherAdmins = s.SaveForOtherAdmins
	o.Misc = s.Misc
}

func (o roleVsysWebuiSaveXml) UnmarshalToObject() (*RoleVsysWebuiSave, error) {

	result := &RoleVsysWebuiSave{
		ObjectLevelChanges: o.ObjectLevelChanges,
		PartialSave:        o.PartialSave,
		SaveForOtherAdmins: o.SaveForOtherAdmins,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *roleVsysXmlapiXml) MarshalFromObject(s RoleVsysXmlapi) {
	o.Commit = s.Commit
	o.Config = s.Config
	o.Export = s.Export
	o.Import = s.Import
	o.Iot = s.Iot
	o.Log = s.Log
	o.Op = s.Op
	o.Report = s.Report
	o.UserId = s.UserId
	o.Misc = s.Misc
}

func (o roleVsysXmlapiXml) UnmarshalToObject() (*RoleVsysXmlapi, error) {

	result := &RoleVsysXmlapi{
		Commit: o.Commit,
		Config: o.Config,
		Export: o.Export,
		Import: o.Import,
		Iot:    o.Iot,
		Log:    o.Log,
		Op:     o.Op,
		Report: o.Report,
		UserId: o.UserId,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *entryXml_11_0_2) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.Description = s.Description
	if s.Role != nil {
		var obj roleXml_11_0_2
		obj.MarshalFromObject(*s.Role)
		o.Role = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml_11_0_2) UnmarshalToObject() (*Entry, error) {
	var roleVal *Role
	if o.Role != nil {
		obj, err := o.Role.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		roleVal = obj
	}

	result := &Entry{
		Name:        o.Name,
		Description: o.Description,
		Role:        roleVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *roleXml_11_0_2) MarshalFromObject(s Role) {
	if s.Device != nil {
		var obj roleDeviceXml_11_0_2
		obj.MarshalFromObject(*s.Device)
		o.Device = &obj
	}
	if s.Vsys != nil {
		var obj roleVsysXml_11_0_2
		obj.MarshalFromObject(*s.Vsys)
		o.Vsys = &obj
	}
	o.Misc = s.Misc
}

func (o roleXml_11_0_2) UnmarshalToObject() (*Role, error) {
	var deviceVal *RoleDevice
	if o.Device != nil {
		obj, err := o.Device.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		deviceVal = obj
	}
	var vsysVal *RoleVsys
	if o.Vsys != nil {
		obj, err := o.Vsys.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		vsysVal = obj
	}

	result := &Role{
		Device: deviceVal,
		Vsys:   vsysVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *roleDeviceXml_11_0_2) MarshalFromObject(s RoleDevice) {
	o.Cli = s.Cli
	if s.Restapi != nil {
		var obj roleDeviceRestapiXml_11_0_2
		obj.MarshalFromObject(*s.Restapi)
		o.Restapi = &obj
	}
	if s.Webui != nil {
		var obj roleDeviceWebuiXml_11_0_2
		obj.MarshalFromObject(*s.Webui)
		o.Webui = &obj
	}
	if s.Xmlapi != nil {
		var obj roleDeviceXmlapiXml_11_0_2
		obj.MarshalFromObject(*s.Xmlapi)
		o.Xmlapi = &obj
	}
	o.Misc = s.Misc
}

func (o roleDeviceXml_11_0_2) UnmarshalToObject() (*RoleDevice, error) {
	var restapiVal *RoleDeviceRestapi
	if o.Restapi != nil {
		obj, err := o.Restapi.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		restapiVal = obj
	}
	var webuiVal *RoleDeviceWebui
	if o.Webui != nil {
		obj, err := o.Webui.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		webuiVal = obj
	}
	var xmlapiVal *RoleDeviceXmlapi
	if o.Xmlapi != nil {
		obj, err := o.Xmlapi.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		xmlapiVal = obj
	}

	result := &RoleDevice{
		Cli:     o.Cli,
		Restapi: restapiVal,
		Webui:   webuiVal,
		Xmlapi:  xmlapiVal,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *roleDeviceRestapiXml_11_0_2) MarshalFromObject(s RoleDeviceRestapi) {
	if s.Device != nil {
		var obj roleDeviceRestapiDeviceXml_11_0_2
		obj.MarshalFromObject(*s.Device)
		o.Device = &obj
	}
	if s.Network != nil {
		var obj roleDeviceRestapiNetworkXml_11_0_2
		obj.MarshalFromObject(*s.Network)
		o.Network = &obj
	}
	if s.Objects != nil {
		var obj roleDeviceRestapiObjectsXml_11_0_2
		obj.MarshalFromObject(*s.Objects)
		o.Objects = &obj
	}
	if s.Policies != nil {
		var obj roleDeviceRestapiPoliciesXml_11_0_2
		obj.MarshalFromObject(*s.Policies)
		o.Policies = &obj
	}
	if s.System != nil {
		var obj roleDeviceRestapiSystemXml_11_0_2
		obj.MarshalFromObject(*s.System)
		o.System = &obj
	}
	o.Misc = s.Misc
}

func (o roleDeviceRestapiXml_11_0_2) UnmarshalToObject() (*RoleDeviceRestapi, error) {
	var deviceVal *RoleDeviceRestapiDevice
	if o.Device != nil {
		obj, err := o.Device.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		deviceVal = obj
	}
	var networkVal *RoleDeviceRestapiNetwork
	if o.Network != nil {
		obj, err := o.Network.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		networkVal = obj
	}
	var objectsVal *RoleDeviceRestapiObjects
	if o.Objects != nil {
		obj, err := o.Objects.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		objectsVal = obj
	}
	var policiesVal *RoleDeviceRestapiPolicies
	if o.Policies != nil {
		obj, err := o.Policies.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		policiesVal = obj
	}
	var systemVal *RoleDeviceRestapiSystem
	if o.System != nil {
		obj, err := o.System.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		systemVal = obj
	}

	result := &RoleDeviceRestapi{
		Device:   deviceVal,
		Network:  networkVal,
		Objects:  objectsVal,
		Policies: policiesVal,
		System:   systemVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *roleDeviceRestapiDeviceXml_11_0_2) MarshalFromObject(s RoleDeviceRestapiDevice) {
	o.EmailServerProfiles = s.EmailServerProfiles
	o.HttpServerProfiles = s.HttpServerProfiles
	o.LdapServerProfiles = s.LdapServerProfiles
	o.LogInterfaceSetting = s.LogInterfaceSetting
	o.SnmpTrapServerProfiles = s.SnmpTrapServerProfiles
	o.SyslogServerProfiles = s.SyslogServerProfiles
	o.VirtualSystems = s.VirtualSystems
	o.Misc = s.Misc
}

func (o roleDeviceRestapiDeviceXml_11_0_2) UnmarshalToObject() (*RoleDeviceRestapiDevice, error) {

	result := &RoleDeviceRestapiDevice{
		EmailServerProfiles:    o.EmailServerProfiles,
		HttpServerProfiles:     o.HttpServerProfiles,
		LdapServerProfiles:     o.LdapServerProfiles,
		LogInterfaceSetting:    o.LogInterfaceSetting,
		SnmpTrapServerProfiles: o.SnmpTrapServerProfiles,
		SyslogServerProfiles:   o.SyslogServerProfiles,
		VirtualSystems:         o.VirtualSystems,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *roleDeviceRestapiNetworkXml_11_0_2) MarshalFromObject(s RoleDeviceRestapiNetwork) {
	o.AggregateEthernetInterfaces = s.AggregateEthernetInterfaces
	o.BfdNetworkProfiles = s.BfdNetworkProfiles
	o.BgpRoutingProfiles = s.BgpRoutingProfiles
	o.DhcpRelays = s.DhcpRelays
	o.DhcpServers = s.DhcpServers
	o.DnsProxies = s.DnsProxies
	o.EthernetInterfaces = s.EthernetInterfaces
	o.GlobalprotectClientlessAppGroups = s.GlobalprotectClientlessAppGroups
	o.GlobalprotectClientlessApps = s.GlobalprotectClientlessApps
	o.GlobalprotectGateways = s.GlobalprotectGateways
	o.GlobalprotectIpsecCryptoNetworkProfiles = s.GlobalprotectIpsecCryptoNetworkProfiles
	o.GlobalprotectMdmServers = s.GlobalprotectMdmServers
	o.GlobalprotectPortals = s.GlobalprotectPortals
	o.GreTunnels = s.GreTunnels
	o.IkeCryptoNetworkProfiles = s.IkeCryptoNetworkProfiles
	o.IkeGatewayNetworkProfiles = s.IkeGatewayNetworkProfiles
	o.InterfaceManagementNetworkProfiles = s.InterfaceManagementNetworkProfiles
	o.IpsecCryptoNetworkProfiles = s.IpsecCryptoNetworkProfiles
	o.IpsecTunnels = s.IpsecTunnels
	o.Lldp = s.Lldp
	o.LldpNetworkProfiles = s.LldpNetworkProfiles
	o.LogicalRouters = s.LogicalRouters
	o.LoopbackInterfaces = s.LoopbackInterfaces
	o.QosInterfaces = s.QosInterfaces
	o.QosNetworkProfiles = s.QosNetworkProfiles
	o.SdwanInterfaceProfiles = s.SdwanInterfaceProfiles
	o.SdwanInterfaces = s.SdwanInterfaces
	o.TunnelInterfaces = s.TunnelInterfaces
	o.TunnelMonitorNetworkProfiles = s.TunnelMonitorNetworkProfiles
	o.VirtualRouters = s.VirtualRouters
	o.VirtualWires = s.VirtualWires
	o.VlanInterfaces = s.VlanInterfaces
	o.Vlans = s.Vlans
	o.ZoneProtectionNetworkProfiles = s.ZoneProtectionNetworkProfiles
	o.Zones = s.Zones
	o.Misc = s.Misc
}

func (o roleDeviceRestapiNetworkXml_11_0_2) UnmarshalToObject() (*RoleDeviceRestapiNetwork, error) {

	result := &RoleDeviceRestapiNetwork{
		AggregateEthernetInterfaces:             o.AggregateEthernetInterfaces,
		BfdNetworkProfiles:                      o.BfdNetworkProfiles,
		BgpRoutingProfiles:                      o.BgpRoutingProfiles,
		DhcpRelays:                              o.DhcpRelays,
		DhcpServers:                             o.DhcpServers,
		DnsProxies:                              o.DnsProxies,
		EthernetInterfaces:                      o.EthernetInterfaces,
		GlobalprotectClientlessAppGroups:        o.GlobalprotectClientlessAppGroups,
		GlobalprotectClientlessApps:             o.GlobalprotectClientlessApps,
		GlobalprotectGateways:                   o.GlobalprotectGateways,
		GlobalprotectIpsecCryptoNetworkProfiles: o.GlobalprotectIpsecCryptoNetworkProfiles,
		GlobalprotectMdmServers:                 o.GlobalprotectMdmServers,
		GlobalprotectPortals:                    o.GlobalprotectPortals,
		GreTunnels:                              o.GreTunnels,
		IkeCryptoNetworkProfiles:                o.IkeCryptoNetworkProfiles,
		IkeGatewayNetworkProfiles:               o.IkeGatewayNetworkProfiles,
		InterfaceManagementNetworkProfiles:      o.InterfaceManagementNetworkProfiles,
		IpsecCryptoNetworkProfiles:              o.IpsecCryptoNetworkProfiles,
		IpsecTunnels:                            o.IpsecTunnels,
		Lldp:                                    o.Lldp,
		LldpNetworkProfiles:                     o.LldpNetworkProfiles,
		LogicalRouters:                          o.LogicalRouters,
		LoopbackInterfaces:                      o.LoopbackInterfaces,
		QosInterfaces:                           o.QosInterfaces,
		QosNetworkProfiles:                      o.QosNetworkProfiles,
		SdwanInterfaceProfiles:                  o.SdwanInterfaceProfiles,
		SdwanInterfaces:                         o.SdwanInterfaces,
		TunnelInterfaces:                        o.TunnelInterfaces,
		TunnelMonitorNetworkProfiles:            o.TunnelMonitorNetworkProfiles,
		VirtualRouters:                          o.VirtualRouters,
		VirtualWires:                            o.VirtualWires,
		VlanInterfaces:                          o.VlanInterfaces,
		Vlans:                                   o.Vlans,
		ZoneProtectionNetworkProfiles:           o.ZoneProtectionNetworkProfiles,
		Zones:                                   o.Zones,
		Misc:                                    o.Misc,
	}
	return result, nil
}
func (o *roleDeviceRestapiObjectsXml_11_0_2) MarshalFromObject(s RoleDeviceRestapiObjects) {
	o.AddressGroups = s.AddressGroups
	o.Addresses = s.Addresses
	o.AntiSpywareSecurityProfiles = s.AntiSpywareSecurityProfiles
	o.AntivirusSecurityProfiles = s.AntivirusSecurityProfiles
	o.ApplicationFilters = s.ApplicationFilters
	o.ApplicationGroups = s.ApplicationGroups
	o.Applications = s.Applications
	o.AuthenticationEnforcements = s.AuthenticationEnforcements
	o.CustomDataPatterns = s.CustomDataPatterns
	o.CustomSpywareSignatures = s.CustomSpywareSignatures
	o.CustomUrlCategories = s.CustomUrlCategories
	o.CustomVulnerabilitySignatures = s.CustomVulnerabilitySignatures
	o.DataFilteringSecurityProfiles = s.DataFilteringSecurityProfiles
	o.DecryptionProfiles = s.DecryptionProfiles
	o.Devices = s.Devices
	o.DosProtectionSecurityProfiles = s.DosProtectionSecurityProfiles
	o.DynamicUserGroups = s.DynamicUserGroups
	o.ExternalDynamicLists = s.ExternalDynamicLists
	o.FileBlockingSecurityProfiles = s.FileBlockingSecurityProfiles
	o.GlobalprotectHipObjects = s.GlobalprotectHipObjects
	o.GlobalprotectHipProfiles = s.GlobalprotectHipProfiles
	o.GtpProtectionSecurityProfiles = s.GtpProtectionSecurityProfiles
	o.LogForwardingProfiles = s.LogForwardingProfiles
	o.PacketBrokerProfiles = s.PacketBrokerProfiles
	o.Regions = s.Regions
	o.Schedules = s.Schedules
	o.SctpProtectionSecurityProfiles = s.SctpProtectionSecurityProfiles
	o.SdwanErrorCorrectionProfiles = s.SdwanErrorCorrectionProfiles
	o.SdwanPathQualityProfiles = s.SdwanPathQualityProfiles
	o.SdwanSaasQualityProfiles = s.SdwanSaasQualityProfiles
	o.SdwanTrafficDistributionProfiles = s.SdwanTrafficDistributionProfiles
	o.SecurityProfileGroups = s.SecurityProfileGroups
	o.ServiceGroups = s.ServiceGroups
	o.Services = s.Services
	o.Tags = s.Tags
	o.UrlFilteringSecurityProfiles = s.UrlFilteringSecurityProfiles
	o.VulnerabilityProtectionSecurityProfiles = s.VulnerabilityProtectionSecurityProfiles
	o.WildfireAnalysisSecurityProfiles = s.WildfireAnalysisSecurityProfiles
	o.Misc = s.Misc
}

func (o roleDeviceRestapiObjectsXml_11_0_2) UnmarshalToObject() (*RoleDeviceRestapiObjects, error) {

	result := &RoleDeviceRestapiObjects{
		AddressGroups:                           o.AddressGroups,
		Addresses:                               o.Addresses,
		AntiSpywareSecurityProfiles:             o.AntiSpywareSecurityProfiles,
		AntivirusSecurityProfiles:               o.AntivirusSecurityProfiles,
		ApplicationFilters:                      o.ApplicationFilters,
		ApplicationGroups:                       o.ApplicationGroups,
		Applications:                            o.Applications,
		AuthenticationEnforcements:              o.AuthenticationEnforcements,
		CustomDataPatterns:                      o.CustomDataPatterns,
		CustomSpywareSignatures:                 o.CustomSpywareSignatures,
		CustomUrlCategories:                     o.CustomUrlCategories,
		CustomVulnerabilitySignatures:           o.CustomVulnerabilitySignatures,
		DataFilteringSecurityProfiles:           o.DataFilteringSecurityProfiles,
		DecryptionProfiles:                      o.DecryptionProfiles,
		Devices:                                 o.Devices,
		DosProtectionSecurityProfiles:           o.DosProtectionSecurityProfiles,
		DynamicUserGroups:                       o.DynamicUserGroups,
		ExternalDynamicLists:                    o.ExternalDynamicLists,
		FileBlockingSecurityProfiles:            o.FileBlockingSecurityProfiles,
		GlobalprotectHipObjects:                 o.GlobalprotectHipObjects,
		GlobalprotectHipProfiles:                o.GlobalprotectHipProfiles,
		GtpProtectionSecurityProfiles:           o.GtpProtectionSecurityProfiles,
		LogForwardingProfiles:                   o.LogForwardingProfiles,
		PacketBrokerProfiles:                    o.PacketBrokerProfiles,
		Regions:                                 o.Regions,
		Schedules:                               o.Schedules,
		SctpProtectionSecurityProfiles:          o.SctpProtectionSecurityProfiles,
		SdwanErrorCorrectionProfiles:            o.SdwanErrorCorrectionProfiles,
		SdwanPathQualityProfiles:                o.SdwanPathQualityProfiles,
		SdwanSaasQualityProfiles:                o.SdwanSaasQualityProfiles,
		SdwanTrafficDistributionProfiles:        o.SdwanTrafficDistributionProfiles,
		SecurityProfileGroups:                   o.SecurityProfileGroups,
		ServiceGroups:                           o.ServiceGroups,
		Services:                                o.Services,
		Tags:                                    o.Tags,
		UrlFilteringSecurityProfiles:            o.UrlFilteringSecurityProfiles,
		VulnerabilityProtectionSecurityProfiles: o.VulnerabilityProtectionSecurityProfiles,
		WildfireAnalysisSecurityProfiles:        o.WildfireAnalysisSecurityProfiles,
		Misc:                                    o.Misc,
	}
	return result, nil
}
func (o *roleDeviceRestapiPoliciesXml_11_0_2) MarshalFromObject(s RoleDeviceRestapiPolicies) {
	o.ApplicationOverrideRules = s.ApplicationOverrideRules
	o.AuthenticationRules = s.AuthenticationRules
	o.DecryptionRules = s.DecryptionRules
	o.DosRules = s.DosRules
	o.NatRules = s.NatRules
	o.NetworkPacketBrokerRules = s.NetworkPacketBrokerRules
	o.PolicyBasedForwardingRules = s.PolicyBasedForwardingRules
	o.QosRules = s.QosRules
	o.SdwanRules = s.SdwanRules
	o.SecurityRules = s.SecurityRules
	o.TunnelInspectionRules = s.TunnelInspectionRules
	o.Misc = s.Misc
}

func (o roleDeviceRestapiPoliciesXml_11_0_2) UnmarshalToObject() (*RoleDeviceRestapiPolicies, error) {

	result := &RoleDeviceRestapiPolicies{
		ApplicationOverrideRules:   o.ApplicationOverrideRules,
		AuthenticationRules:        o.AuthenticationRules,
		DecryptionRules:            o.DecryptionRules,
		DosRules:                   o.DosRules,
		NatRules:                   o.NatRules,
		NetworkPacketBrokerRules:   o.NetworkPacketBrokerRules,
		PolicyBasedForwardingRules: o.PolicyBasedForwardingRules,
		QosRules:                   o.QosRules,
		SdwanRules:                 o.SdwanRules,
		SecurityRules:              o.SecurityRules,
		TunnelInspectionRules:      o.TunnelInspectionRules,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *roleDeviceRestapiSystemXml_11_0_2) MarshalFromObject(s RoleDeviceRestapiSystem) {
	o.Configuration = s.Configuration
	o.Misc = s.Misc
}

func (o roleDeviceRestapiSystemXml_11_0_2) UnmarshalToObject() (*RoleDeviceRestapiSystem, error) {

	result := &RoleDeviceRestapiSystem{
		Configuration: o.Configuration,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiXml_11_0_2) MarshalFromObject(s RoleDeviceWebui) {
	o.Acc = s.Acc
	if s.Commit != nil {
		var obj roleDeviceWebuiCommitXml_11_0_2
		obj.MarshalFromObject(*s.Commit)
		o.Commit = &obj
	}
	o.Dashboard = s.Dashboard
	if s.Device != nil {
		var obj roleDeviceWebuiDeviceXml_11_0_2
		obj.MarshalFromObject(*s.Device)
		o.Device = &obj
	}
	if s.Global != nil {
		var obj roleDeviceWebuiGlobalXml_11_0_2
		obj.MarshalFromObject(*s.Global)
		o.Global = &obj
	}
	if s.Monitor != nil {
		var obj roleDeviceWebuiMonitorXml_11_0_2
		obj.MarshalFromObject(*s.Monitor)
		o.Monitor = &obj
	}
	if s.Network != nil {
		var obj roleDeviceWebuiNetworkXml_11_0_2
		obj.MarshalFromObject(*s.Network)
		o.Network = &obj
	}
	if s.Objects != nil {
		var obj roleDeviceWebuiObjectsXml_11_0_2
		obj.MarshalFromObject(*s.Objects)
		o.Objects = &obj
	}
	if s.Operations != nil {
		var obj roleDeviceWebuiOperationsXml_11_0_2
		obj.MarshalFromObject(*s.Operations)
		o.Operations = &obj
	}
	if s.Policies != nil {
		var obj roleDeviceWebuiPoliciesXml_11_0_2
		obj.MarshalFromObject(*s.Policies)
		o.Policies = &obj
	}
	if s.Privacy != nil {
		var obj roleDeviceWebuiPrivacyXml_11_0_2
		obj.MarshalFromObject(*s.Privacy)
		o.Privacy = &obj
	}
	if s.Save != nil {
		var obj roleDeviceWebuiSaveXml_11_0_2
		obj.MarshalFromObject(*s.Save)
		o.Save = &obj
	}
	o.Tasks = s.Tasks
	o.Validate = s.Validate
	o.Misc = s.Misc
}

func (o roleDeviceWebuiXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebui, error) {
	var commitVal *RoleDeviceWebuiCommit
	if o.Commit != nil {
		obj, err := o.Commit.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		commitVal = obj
	}
	var deviceVal *RoleDeviceWebuiDevice
	if o.Device != nil {
		obj, err := o.Device.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		deviceVal = obj
	}
	var globalVal *RoleDeviceWebuiGlobal
	if o.Global != nil {
		obj, err := o.Global.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalVal = obj
	}
	var monitorVal *RoleDeviceWebuiMonitor
	if o.Monitor != nil {
		obj, err := o.Monitor.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		monitorVal = obj
	}
	var networkVal *RoleDeviceWebuiNetwork
	if o.Network != nil {
		obj, err := o.Network.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		networkVal = obj
	}
	var objectsVal *RoleDeviceWebuiObjects
	if o.Objects != nil {
		obj, err := o.Objects.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		objectsVal = obj
	}
	var operationsVal *RoleDeviceWebuiOperations
	if o.Operations != nil {
		obj, err := o.Operations.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		operationsVal = obj
	}
	var policiesVal *RoleDeviceWebuiPolicies
	if o.Policies != nil {
		obj, err := o.Policies.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		policiesVal = obj
	}
	var privacyVal *RoleDeviceWebuiPrivacy
	if o.Privacy != nil {
		obj, err := o.Privacy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		privacyVal = obj
	}
	var saveVal *RoleDeviceWebuiSave
	if o.Save != nil {
		obj, err := o.Save.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		saveVal = obj
	}

	result := &RoleDeviceWebui{
		Acc:        o.Acc,
		Commit:     commitVal,
		Dashboard:  o.Dashboard,
		Device:     deviceVal,
		Global:     globalVal,
		Monitor:    monitorVal,
		Network:    networkVal,
		Objects:    objectsVal,
		Operations: operationsVal,
		Policies:   policiesVal,
		Privacy:    privacyVal,
		Save:       saveVal,
		Tasks:      o.Tasks,
		Validate:   o.Validate,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiCommitXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiCommit) {
	o.CommitForOtherAdmins = s.CommitForOtherAdmins
	o.Device = s.Device
	o.ObjectLevelChanges = s.ObjectLevelChanges
	o.Misc = s.Misc
}

func (o roleDeviceWebuiCommitXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiCommit, error) {

	result := &RoleDeviceWebuiCommit{
		CommitForOtherAdmins: o.CommitForOtherAdmins,
		Device:               o.Device,
		ObjectLevelChanges:   o.ObjectLevelChanges,
		Misc:                 o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiDeviceXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiDevice) {
	o.AccessDomain = s.AccessDomain
	o.AdminRoles = s.AdminRoles
	o.Administrators = s.Administrators
	o.AuthenticationProfile = s.AuthenticationProfile
	o.AuthenticationSequence = s.AuthenticationSequence
	o.BlockPages = s.BlockPages
	if s.CertificateManagement != nil {
		var obj roleDeviceWebuiDeviceCertificateManagementXml_11_0_2
		obj.MarshalFromObject(*s.CertificateManagement)
		o.CertificateManagement = &obj
	}
	o.ConfigAudit = s.ConfigAudit
	o.DataRedistribution = s.DataRedistribution
	o.DeviceQuarantine = s.DeviceQuarantine
	o.DhcpSyslogServer = s.DhcpSyslogServer
	o.DynamicUpdates = s.DynamicUpdates
	o.GlobalProtectClient = s.GlobalProtectClient
	o.HighAvailability = s.HighAvailability
	o.Licenses = s.Licenses
	if s.LocalUserDatabase != nil {
		var obj roleDeviceWebuiDeviceLocalUserDatabaseXml_11_0_2
		obj.MarshalFromObject(*s.LocalUserDatabase)
		o.LocalUserDatabase = &obj
	}
	o.LogFwdCard = s.LogFwdCard
	if s.LogSettings != nil {
		var obj roleDeviceWebuiDeviceLogSettingsXml_11_0_2
		obj.MarshalFromObject(*s.LogSettings)
		o.LogSettings = &obj
	}
	o.MasterKey = s.MasterKey
	o.Plugins = s.Plugins
	if s.PolicyRecommendations != nil {
		var obj roleDeviceWebuiDevicePolicyRecommendationsXml_11_0_2
		obj.MarshalFromObject(*s.PolicyRecommendations)
		o.PolicyRecommendations = &obj
	}
	o.ScheduledLogExport = s.ScheduledLogExport
	if s.ServerProfile != nil {
		var obj roleDeviceWebuiDeviceServerProfileXml_11_0_2
		obj.MarshalFromObject(*s.ServerProfile)
		o.ServerProfile = &obj
	}
	if s.Setup != nil {
		var obj roleDeviceWebuiDeviceSetupXml_11_0_2
		obj.MarshalFromObject(*s.Setup)
		o.Setup = &obj
	}
	o.SharedGateways = s.SharedGateways
	o.Software = s.Software
	o.Support = s.Support
	o.Troubleshooting = s.Troubleshooting
	o.UserIdentification = s.UserIdentification
	o.VirtualSystems = s.VirtualSystems
	o.VmInfoSource = s.VmInfoSource
	o.Misc = s.Misc
}

func (o roleDeviceWebuiDeviceXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiDevice, error) {
	var certificateManagementVal *RoleDeviceWebuiDeviceCertificateManagement
	if o.CertificateManagement != nil {
		obj, err := o.CertificateManagement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		certificateManagementVal = obj
	}
	var localUserDatabaseVal *RoleDeviceWebuiDeviceLocalUserDatabase
	if o.LocalUserDatabase != nil {
		obj, err := o.LocalUserDatabase.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localUserDatabaseVal = obj
	}
	var logSettingsVal *RoleDeviceWebuiDeviceLogSettings
	if o.LogSettings != nil {
		obj, err := o.LogSettings.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		logSettingsVal = obj
	}
	var policyRecommendationsVal *RoleDeviceWebuiDevicePolicyRecommendations
	if o.PolicyRecommendations != nil {
		obj, err := o.PolicyRecommendations.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		policyRecommendationsVal = obj
	}
	var serverProfileVal *RoleDeviceWebuiDeviceServerProfile
	if o.ServerProfile != nil {
		obj, err := o.ServerProfile.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		serverProfileVal = obj
	}
	var setupVal *RoleDeviceWebuiDeviceSetup
	if o.Setup != nil {
		obj, err := o.Setup.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setupVal = obj
	}

	result := &RoleDeviceWebuiDevice{
		AccessDomain:           o.AccessDomain,
		AdminRoles:             o.AdminRoles,
		Administrators:         o.Administrators,
		AuthenticationProfile:  o.AuthenticationProfile,
		AuthenticationSequence: o.AuthenticationSequence,
		BlockPages:             o.BlockPages,
		CertificateManagement:  certificateManagementVal,
		ConfigAudit:            o.ConfigAudit,
		DataRedistribution:     o.DataRedistribution,
		DeviceQuarantine:       o.DeviceQuarantine,
		DhcpSyslogServer:       o.DhcpSyslogServer,
		DynamicUpdates:         o.DynamicUpdates,
		GlobalProtectClient:    o.GlobalProtectClient,
		HighAvailability:       o.HighAvailability,
		Licenses:               o.Licenses,
		LocalUserDatabase:      localUserDatabaseVal,
		LogFwdCard:             o.LogFwdCard,
		LogSettings:            logSettingsVal,
		MasterKey:              o.MasterKey,
		Plugins:                o.Plugins,
		PolicyRecommendations:  policyRecommendationsVal,
		ScheduledLogExport:     o.ScheduledLogExport,
		ServerProfile:          serverProfileVal,
		Setup:                  setupVal,
		SharedGateways:         o.SharedGateways,
		Software:               o.Software,
		Support:                o.Support,
		Troubleshooting:        o.Troubleshooting,
		UserIdentification:     o.UserIdentification,
		VirtualSystems:         o.VirtualSystems,
		VmInfoSource:           o.VmInfoSource,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiDeviceCertificateManagementXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiDeviceCertificateManagement) {
	o.CertificateProfile = s.CertificateProfile
	o.Certificates = s.Certificates
	o.OcspResponder = s.OcspResponder
	o.Scep = s.Scep
	o.SshServiceProfile = s.SshServiceProfile
	o.SslDecryptionExclusion = s.SslDecryptionExclusion
	o.SslTlsServiceProfile = s.SslTlsServiceProfile
	o.Misc = s.Misc
}

func (o roleDeviceWebuiDeviceCertificateManagementXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiDeviceCertificateManagement, error) {

	result := &RoleDeviceWebuiDeviceCertificateManagement{
		CertificateProfile:     o.CertificateProfile,
		Certificates:           o.Certificates,
		OcspResponder:          o.OcspResponder,
		Scep:                   o.Scep,
		SshServiceProfile:      o.SshServiceProfile,
		SslDecryptionExclusion: o.SslDecryptionExclusion,
		SslTlsServiceProfile:   o.SslTlsServiceProfile,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiDeviceLocalUserDatabaseXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiDeviceLocalUserDatabase) {
	o.UserGroups = s.UserGroups
	o.Users = s.Users
	o.Misc = s.Misc
}

func (o roleDeviceWebuiDeviceLocalUserDatabaseXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiDeviceLocalUserDatabase, error) {

	result := &RoleDeviceWebuiDeviceLocalUserDatabase{
		UserGroups: o.UserGroups,
		Users:      o.Users,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiDeviceLogSettingsXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiDeviceLogSettings) {
	o.CcAlarm = s.CcAlarm
	o.Config = s.Config
	o.Correlation = s.Correlation
	o.Globalprotect = s.Globalprotect
	o.Hipmatch = s.Hipmatch
	o.Iptag = s.Iptag
	o.ManageLog = s.ManageLog
	o.System = s.System
	o.UserId = s.UserId
	o.Misc = s.Misc
}

func (o roleDeviceWebuiDeviceLogSettingsXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiDeviceLogSettings, error) {

	result := &RoleDeviceWebuiDeviceLogSettings{
		CcAlarm:       o.CcAlarm,
		Config:        o.Config,
		Correlation:   o.Correlation,
		Globalprotect: o.Globalprotect,
		Hipmatch:      o.Hipmatch,
		Iptag:         o.Iptag,
		ManageLog:     o.ManageLog,
		System:        o.System,
		UserId:        o.UserId,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiDevicePolicyRecommendationsXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiDevicePolicyRecommendations) {
	o.Iot = s.Iot
	o.Saas = s.Saas
	o.Misc = s.Misc
}

func (o roleDeviceWebuiDevicePolicyRecommendationsXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiDevicePolicyRecommendations, error) {

	result := &RoleDeviceWebuiDevicePolicyRecommendations{
		Iot:  o.Iot,
		Saas: o.Saas,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiDeviceServerProfileXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiDeviceServerProfile) {
	o.Dns = s.Dns
	o.Email = s.Email
	o.Http = s.Http
	o.Kerberos = s.Kerberos
	o.Ldap = s.Ldap
	o.Mfa = s.Mfa
	o.Netflow = s.Netflow
	o.Radius = s.Radius
	o.SamlIdp = s.SamlIdp
	o.Scp = s.Scp
	o.SnmpTrap = s.SnmpTrap
	o.Syslog = s.Syslog
	o.Tacplus = s.Tacplus
	o.Misc = s.Misc
}

func (o roleDeviceWebuiDeviceServerProfileXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiDeviceServerProfile, error) {

	result := &RoleDeviceWebuiDeviceServerProfile{
		Dns:      o.Dns,
		Email:    o.Email,
		Http:     o.Http,
		Kerberos: o.Kerberos,
		Ldap:     o.Ldap,
		Mfa:      o.Mfa,
		Netflow:  o.Netflow,
		Radius:   o.Radius,
		SamlIdp:  o.SamlIdp,
		Scp:      o.Scp,
		SnmpTrap: o.SnmpTrap,
		Syslog:   o.Syslog,
		Tacplus:  o.Tacplus,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiDeviceSetupXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiDeviceSetup) {
	o.ContentId = s.ContentId
	o.Hsm = s.Hsm
	o.Interfaces = s.Interfaces
	o.Management = s.Management
	o.Operations = s.Operations
	o.Services = s.Services
	o.Session = s.Session
	o.Telemetry = s.Telemetry
	o.Wildfire = s.Wildfire
	o.Misc = s.Misc
}

func (o roleDeviceWebuiDeviceSetupXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiDeviceSetup, error) {

	result := &RoleDeviceWebuiDeviceSetup{
		ContentId:  o.ContentId,
		Hsm:        o.Hsm,
		Interfaces: o.Interfaces,
		Management: o.Management,
		Operations: o.Operations,
		Services:   o.Services,
		Session:    o.Session,
		Telemetry:  o.Telemetry,
		Wildfire:   o.Wildfire,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiGlobalXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiGlobal) {
	o.SystemAlarms = s.SystemAlarms
	o.Misc = s.Misc
}

func (o roleDeviceWebuiGlobalXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiGlobal, error) {

	result := &RoleDeviceWebuiGlobal{
		SystemAlarms: o.SystemAlarms,
		Misc:         o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiMonitorXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiMonitor) {
	o.AppScope = s.AppScope
	o.ApplicationReports = s.ApplicationReports
	if s.AutomatedCorrelationEngine != nil {
		var obj roleDeviceWebuiMonitorAutomatedCorrelationEngineXml_11_0_2
		obj.MarshalFromObject(*s.AutomatedCorrelationEngine)
		o.AutomatedCorrelationEngine = &obj
	}
	o.BlockIpList = s.BlockIpList
	o.Botnet = s.Botnet
	if s.CustomReports != nil {
		var obj roleDeviceWebuiMonitorCustomReportsXml_11_0_2
		obj.MarshalFromObject(*s.CustomReports)
		o.CustomReports = &obj
	}
	o.ExternalLogs = s.ExternalLogs
	o.GtpReports = s.GtpReports
	if s.Logs != nil {
		var obj roleDeviceWebuiMonitorLogsXml_11_0_2
		obj.MarshalFromObject(*s.Logs)
		o.Logs = &obj
	}
	o.PacketCapture = s.PacketCapture
	if s.PdfReports != nil {
		var obj roleDeviceWebuiMonitorPdfReportsXml_11_0_2
		obj.MarshalFromObject(*s.PdfReports)
		o.PdfReports = &obj
	}
	o.SctpReports = s.SctpReports
	o.SessionBrowser = s.SessionBrowser
	o.ThreatReports = s.ThreatReports
	o.TrafficReports = s.TrafficReports
	o.UrlFilteringReports = s.UrlFilteringReports
	o.ViewCustomReports = s.ViewCustomReports
	o.Misc = s.Misc
}

func (o roleDeviceWebuiMonitorXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiMonitor, error) {
	var automatedCorrelationEngineVal *RoleDeviceWebuiMonitorAutomatedCorrelationEngine
	if o.AutomatedCorrelationEngine != nil {
		obj, err := o.AutomatedCorrelationEngine.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		automatedCorrelationEngineVal = obj
	}
	var customReportsVal *RoleDeviceWebuiMonitorCustomReports
	if o.CustomReports != nil {
		obj, err := o.CustomReports.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		customReportsVal = obj
	}
	var logsVal *RoleDeviceWebuiMonitorLogs
	if o.Logs != nil {
		obj, err := o.Logs.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		logsVal = obj
	}
	var pdfReportsVal *RoleDeviceWebuiMonitorPdfReports
	if o.PdfReports != nil {
		obj, err := o.PdfReports.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		pdfReportsVal = obj
	}

	result := &RoleDeviceWebuiMonitor{
		AppScope:                   o.AppScope,
		ApplicationReports:         o.ApplicationReports,
		AutomatedCorrelationEngine: automatedCorrelationEngineVal,
		BlockIpList:                o.BlockIpList,
		Botnet:                     o.Botnet,
		CustomReports:              customReportsVal,
		ExternalLogs:               o.ExternalLogs,
		GtpReports:                 o.GtpReports,
		Logs:                       logsVal,
		PacketCapture:              o.PacketCapture,
		PdfReports:                 pdfReportsVal,
		SctpReports:                o.SctpReports,
		SessionBrowser:             o.SessionBrowser,
		ThreatReports:              o.ThreatReports,
		TrafficReports:             o.TrafficReports,
		UrlFilteringReports:        o.UrlFilteringReports,
		ViewCustomReports:          o.ViewCustomReports,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiMonitorAutomatedCorrelationEngineXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiMonitorAutomatedCorrelationEngine) {
	o.CorrelatedEvents = s.CorrelatedEvents
	o.CorrelationObjects = s.CorrelationObjects
	o.Misc = s.Misc
}

func (o roleDeviceWebuiMonitorAutomatedCorrelationEngineXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiMonitorAutomatedCorrelationEngine, error) {

	result := &RoleDeviceWebuiMonitorAutomatedCorrelationEngine{
		CorrelatedEvents:   o.CorrelatedEvents,
		CorrelationObjects: o.CorrelationObjects,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiMonitorCustomReportsXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiMonitorCustomReports) {
	o.ApplicationStatistics = s.ApplicationStatistics
	o.Auth = s.Auth
	o.DataFilteringLog = s.DataFilteringLog
	o.DecryptionLog = s.DecryptionLog
	o.DecryptionSummary = s.DecryptionSummary
	o.Globalprotect = s.Globalprotect
	o.GtpLog = s.GtpLog
	o.GtpSummary = s.GtpSummary
	o.Hipmatch = s.Hipmatch
	o.Iptag = s.Iptag
	o.SctpLog = s.SctpLog
	o.SctpSummary = s.SctpSummary
	o.ThreatLog = s.ThreatLog
	o.ThreatSummary = s.ThreatSummary
	o.TrafficLog = s.TrafficLog
	o.TrafficSummary = s.TrafficSummary
	o.TunnelLog = s.TunnelLog
	o.TunnelSummary = s.TunnelSummary
	o.UrlLog = s.UrlLog
	o.UrlSummary = s.UrlSummary
	o.Userid = s.Userid
	o.WildfireLog = s.WildfireLog
	o.Misc = s.Misc
}

func (o roleDeviceWebuiMonitorCustomReportsXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiMonitorCustomReports, error) {

	result := &RoleDeviceWebuiMonitorCustomReports{
		ApplicationStatistics: o.ApplicationStatistics,
		Auth:                  o.Auth,
		DataFilteringLog:      o.DataFilteringLog,
		DecryptionLog:         o.DecryptionLog,
		DecryptionSummary:     o.DecryptionSummary,
		Globalprotect:         o.Globalprotect,
		GtpLog:                o.GtpLog,
		GtpSummary:            o.GtpSummary,
		Hipmatch:              o.Hipmatch,
		Iptag:                 o.Iptag,
		SctpLog:               o.SctpLog,
		SctpSummary:           o.SctpSummary,
		ThreatLog:             o.ThreatLog,
		ThreatSummary:         o.ThreatSummary,
		TrafficLog:            o.TrafficLog,
		TrafficSummary:        o.TrafficSummary,
		TunnelLog:             o.TunnelLog,
		TunnelSummary:         o.TunnelSummary,
		UrlLog:                o.UrlLog,
		UrlSummary:            o.UrlSummary,
		Userid:                o.Userid,
		WildfireLog:           o.WildfireLog,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiMonitorLogsXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiMonitorLogs) {
	o.Alarm = s.Alarm
	o.Authentication = s.Authentication
	o.Configuration = s.Configuration
	o.DataFiltering = s.DataFiltering
	o.Decryption = s.Decryption
	o.Globalprotect = s.Globalprotect
	o.Gtp = s.Gtp
	o.Hipmatch = s.Hipmatch
	o.Iptag = s.Iptag
	o.Sctp = s.Sctp
	o.System = s.System
	o.Threat = s.Threat
	o.Traffic = s.Traffic
	o.Tunnel = s.Tunnel
	o.Url = s.Url
	o.Userid = s.Userid
	o.Wildfire = s.Wildfire
	o.Misc = s.Misc
}

func (o roleDeviceWebuiMonitorLogsXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiMonitorLogs, error) {

	result := &RoleDeviceWebuiMonitorLogs{
		Alarm:          o.Alarm,
		Authentication: o.Authentication,
		Configuration:  o.Configuration,
		DataFiltering:  o.DataFiltering,
		Decryption:     o.Decryption,
		Globalprotect:  o.Globalprotect,
		Gtp:            o.Gtp,
		Hipmatch:       o.Hipmatch,
		Iptag:          o.Iptag,
		Sctp:           o.Sctp,
		System:         o.System,
		Threat:         o.Threat,
		Traffic:        o.Traffic,
		Tunnel:         o.Tunnel,
		Url:            o.Url,
		Userid:         o.Userid,
		Wildfire:       o.Wildfire,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiMonitorPdfReportsXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiMonitorPdfReports) {
	o.EmailScheduler = s.EmailScheduler
	o.ManagePdfSummary = s.ManagePdfSummary
	o.PdfSummaryReports = s.PdfSummaryReports
	o.ReportGroups = s.ReportGroups
	o.SaasApplicationUsageReport = s.SaasApplicationUsageReport
	o.UserActivityReport = s.UserActivityReport
	o.Misc = s.Misc
}

func (o roleDeviceWebuiMonitorPdfReportsXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiMonitorPdfReports, error) {

	result := &RoleDeviceWebuiMonitorPdfReports{
		EmailScheduler:             o.EmailScheduler,
		ManagePdfSummary:           o.ManagePdfSummary,
		PdfSummaryReports:          o.PdfSummaryReports,
		ReportGroups:               o.ReportGroups,
		SaasApplicationUsageReport: o.SaasApplicationUsageReport,
		UserActivityReport:         o.UserActivityReport,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiNetworkXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiNetwork) {
	o.Dhcp = s.Dhcp
	o.DnsProxy = s.DnsProxy
	if s.GlobalProtect != nil {
		var obj roleDeviceWebuiNetworkGlobalProtectXml_11_0_2
		obj.MarshalFromObject(*s.GlobalProtect)
		o.GlobalProtect = &obj
	}
	o.GreTunnels = s.GreTunnels
	o.Interfaces = s.Interfaces
	o.IpsecTunnels = s.IpsecTunnels
	o.Lldp = s.Lldp
	if s.NetworkProfiles != nil {
		var obj roleDeviceWebuiNetworkNetworkProfilesXml_11_0_2
		obj.MarshalFromObject(*s.NetworkProfiles)
		o.NetworkProfiles = &obj
	}
	o.Qos = s.Qos
	if s.Routing != nil {
		var obj roleDeviceWebuiNetworkRoutingXml_11_0_2
		obj.MarshalFromObject(*s.Routing)
		o.Routing = &obj
	}
	o.SdwanInterfaceProfile = s.SdwanInterfaceProfile
	o.SecureWebGateway = s.SecureWebGateway
	o.VirtualRouters = s.VirtualRouters
	o.VirtualWires = s.VirtualWires
	o.Vlans = s.Vlans
	o.Zones = s.Zones
	o.Misc = s.Misc
}

func (o roleDeviceWebuiNetworkXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiNetwork, error) {
	var globalProtectVal *RoleDeviceWebuiNetworkGlobalProtect
	if o.GlobalProtect != nil {
		obj, err := o.GlobalProtect.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalProtectVal = obj
	}
	var networkProfilesVal *RoleDeviceWebuiNetworkNetworkProfiles
	if o.NetworkProfiles != nil {
		obj, err := o.NetworkProfiles.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		networkProfilesVal = obj
	}
	var routingVal *RoleDeviceWebuiNetworkRouting
	if o.Routing != nil {
		obj, err := o.Routing.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routingVal = obj
	}

	result := &RoleDeviceWebuiNetwork{
		Dhcp:                  o.Dhcp,
		DnsProxy:              o.DnsProxy,
		GlobalProtect:         globalProtectVal,
		GreTunnels:            o.GreTunnels,
		Interfaces:            o.Interfaces,
		IpsecTunnels:          o.IpsecTunnels,
		Lldp:                  o.Lldp,
		NetworkProfiles:       networkProfilesVal,
		Qos:                   o.Qos,
		Routing:               routingVal,
		SdwanInterfaceProfile: o.SdwanInterfaceProfile,
		SecureWebGateway:      o.SecureWebGateway,
		VirtualRouters:        o.VirtualRouters,
		VirtualWires:          o.VirtualWires,
		Vlans:                 o.Vlans,
		Zones:                 o.Zones,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiNetworkGlobalProtectXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiNetworkGlobalProtect) {
	o.ClientlessAppGroups = s.ClientlessAppGroups
	o.ClientlessApps = s.ClientlessApps
	o.Gateways = s.Gateways
	o.Mdm = s.Mdm
	o.Portals = s.Portals
	o.Misc = s.Misc
}

func (o roleDeviceWebuiNetworkGlobalProtectXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiNetworkGlobalProtect, error) {

	result := &RoleDeviceWebuiNetworkGlobalProtect{
		ClientlessAppGroups: o.ClientlessAppGroups,
		ClientlessApps:      o.ClientlessApps,
		Gateways:            o.Gateways,
		Mdm:                 o.Mdm,
		Portals:             o.Portals,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiNetworkNetworkProfilesXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiNetworkNetworkProfiles) {
	o.BfdProfile = s.BfdProfile
	o.GpAppIpsecCrypto = s.GpAppIpsecCrypto
	o.IkeCrypto = s.IkeCrypto
	o.IkeGateways = s.IkeGateways
	o.InterfaceMgmt = s.InterfaceMgmt
	o.IpsecCrypto = s.IpsecCrypto
	o.LldpProfile = s.LldpProfile
	o.QosProfile = s.QosProfile
	o.TunnelMonitor = s.TunnelMonitor
	o.ZoneProtection = s.ZoneProtection
	o.Misc = s.Misc
}

func (o roleDeviceWebuiNetworkNetworkProfilesXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiNetworkNetworkProfiles, error) {

	result := &RoleDeviceWebuiNetworkNetworkProfiles{
		BfdProfile:       o.BfdProfile,
		GpAppIpsecCrypto: o.GpAppIpsecCrypto,
		IkeCrypto:        o.IkeCrypto,
		IkeGateways:      o.IkeGateways,
		InterfaceMgmt:    o.InterfaceMgmt,
		IpsecCrypto:      o.IpsecCrypto,
		LldpProfile:      o.LldpProfile,
		QosProfile:       o.QosProfile,
		TunnelMonitor:    o.TunnelMonitor,
		ZoneProtection:   o.ZoneProtection,
		Misc:             o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiNetworkRoutingXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiNetworkRouting) {
	o.LogicalRouters = s.LogicalRouters
	if s.RoutingProfiles != nil {
		var obj roleDeviceWebuiNetworkRoutingRoutingProfilesXml_11_0_2
		obj.MarshalFromObject(*s.RoutingProfiles)
		o.RoutingProfiles = &obj
	}
	o.Misc = s.Misc
}

func (o roleDeviceWebuiNetworkRoutingXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiNetworkRouting, error) {
	var routingProfilesVal *RoleDeviceWebuiNetworkRoutingRoutingProfiles
	if o.RoutingProfiles != nil {
		obj, err := o.RoutingProfiles.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		routingProfilesVal = obj
	}

	result := &RoleDeviceWebuiNetworkRouting{
		LogicalRouters:  o.LogicalRouters,
		RoutingProfiles: routingProfilesVal,
		Misc:            o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiNetworkRoutingRoutingProfilesXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiNetworkRoutingRoutingProfiles) {
	o.Bfd = s.Bfd
	o.Bgp = s.Bgp
	o.Filters = s.Filters
	o.Multicast = s.Multicast
	o.Ospf = s.Ospf
	o.Ospfv3 = s.Ospfv3
	o.Ripv2 = s.Ripv2
	o.Misc = s.Misc
}

func (o roleDeviceWebuiNetworkRoutingRoutingProfilesXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiNetworkRoutingRoutingProfiles, error) {

	result := &RoleDeviceWebuiNetworkRoutingRoutingProfiles{
		Bfd:       o.Bfd,
		Bgp:       o.Bgp,
		Filters:   o.Filters,
		Multicast: o.Multicast,
		Ospf:      o.Ospf,
		Ospfv3:    o.Ospfv3,
		Ripv2:     o.Ripv2,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiObjectsXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiObjects) {
	o.AddressGroups = s.AddressGroups
	o.Addresses = s.Addresses
	o.ApplicationFilters = s.ApplicationFilters
	o.ApplicationGroups = s.ApplicationGroups
	o.Applications = s.Applications
	o.Authentication = s.Authentication
	if s.CustomObjects != nil {
		var obj roleDeviceWebuiObjectsCustomObjectsXml_11_0_2
		obj.MarshalFromObject(*s.CustomObjects)
		o.CustomObjects = &obj
	}
	if s.Decryption != nil {
		var obj roleDeviceWebuiObjectsDecryptionXml_11_0_2
		obj.MarshalFromObject(*s.Decryption)
		o.Decryption = &obj
	}
	o.Devices = s.Devices
	o.DynamicBlockLists = s.DynamicBlockLists
	o.DynamicUserGroups = s.DynamicUserGroups
	if s.GlobalProtect != nil {
		var obj roleDeviceWebuiObjectsGlobalProtectXml_11_0_2
		obj.MarshalFromObject(*s.GlobalProtect)
		o.GlobalProtect = &obj
	}
	o.LogForwarding = s.LogForwarding
	o.PacketBrokerProfile = s.PacketBrokerProfile
	o.Regions = s.Regions
	o.Schedules = s.Schedules
	if s.Sdwan != nil {
		var obj roleDeviceWebuiObjectsSdwanXml_11_0_2
		obj.MarshalFromObject(*s.Sdwan)
		o.Sdwan = &obj
	}
	o.SecurityProfileGroups = s.SecurityProfileGroups
	if s.SecurityProfiles != nil {
		var obj roleDeviceWebuiObjectsSecurityProfilesXml_11_0_2
		obj.MarshalFromObject(*s.SecurityProfiles)
		o.SecurityProfiles = &obj
	}
	o.ServiceGroups = s.ServiceGroups
	o.Services = s.Services
	o.Tags = s.Tags
	o.Misc = s.Misc
}

func (o roleDeviceWebuiObjectsXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiObjects, error) {
	var customObjectsVal *RoleDeviceWebuiObjectsCustomObjects
	if o.CustomObjects != nil {
		obj, err := o.CustomObjects.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		customObjectsVal = obj
	}
	var decryptionVal *RoleDeviceWebuiObjectsDecryption
	if o.Decryption != nil {
		obj, err := o.Decryption.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		decryptionVal = obj
	}
	var globalProtectVal *RoleDeviceWebuiObjectsGlobalProtect
	if o.GlobalProtect != nil {
		obj, err := o.GlobalProtect.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalProtectVal = obj
	}
	var sdwanVal *RoleDeviceWebuiObjectsSdwan
	if o.Sdwan != nil {
		obj, err := o.Sdwan.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sdwanVal = obj
	}
	var securityProfilesVal *RoleDeviceWebuiObjectsSecurityProfiles
	if o.SecurityProfiles != nil {
		obj, err := o.SecurityProfiles.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		securityProfilesVal = obj
	}

	result := &RoleDeviceWebuiObjects{
		AddressGroups:         o.AddressGroups,
		Addresses:             o.Addresses,
		ApplicationFilters:    o.ApplicationFilters,
		ApplicationGroups:     o.ApplicationGroups,
		Applications:          o.Applications,
		Authentication:        o.Authentication,
		CustomObjects:         customObjectsVal,
		Decryption:            decryptionVal,
		Devices:               o.Devices,
		DynamicBlockLists:     o.DynamicBlockLists,
		DynamicUserGroups:     o.DynamicUserGroups,
		GlobalProtect:         globalProtectVal,
		LogForwarding:         o.LogForwarding,
		PacketBrokerProfile:   o.PacketBrokerProfile,
		Regions:               o.Regions,
		Schedules:             o.Schedules,
		Sdwan:                 sdwanVal,
		SecurityProfileGroups: o.SecurityProfileGroups,
		SecurityProfiles:      securityProfilesVal,
		ServiceGroups:         o.ServiceGroups,
		Services:              o.Services,
		Tags:                  o.Tags,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiObjectsCustomObjectsXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiObjectsCustomObjects) {
	o.DataPatterns = s.DataPatterns
	o.Spyware = s.Spyware
	o.UrlCategory = s.UrlCategory
	o.Vulnerability = s.Vulnerability
	o.Misc = s.Misc
}

func (o roleDeviceWebuiObjectsCustomObjectsXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiObjectsCustomObjects, error) {

	result := &RoleDeviceWebuiObjectsCustomObjects{
		DataPatterns:  o.DataPatterns,
		Spyware:       o.Spyware,
		UrlCategory:   o.UrlCategory,
		Vulnerability: o.Vulnerability,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiObjectsDecryptionXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiObjectsDecryption) {
	o.DecryptionProfile = s.DecryptionProfile
	o.Misc = s.Misc
}

func (o roleDeviceWebuiObjectsDecryptionXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiObjectsDecryption, error) {

	result := &RoleDeviceWebuiObjectsDecryption{
		DecryptionProfile: o.DecryptionProfile,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiObjectsGlobalProtectXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiObjectsGlobalProtect) {
	o.HipObjects = s.HipObjects
	o.HipProfiles = s.HipProfiles
	o.Misc = s.Misc
}

func (o roleDeviceWebuiObjectsGlobalProtectXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiObjectsGlobalProtect, error) {

	result := &RoleDeviceWebuiObjectsGlobalProtect{
		HipObjects:  o.HipObjects,
		HipProfiles: o.HipProfiles,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiObjectsSdwanXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiObjectsSdwan) {
	o.SdwanDistProfile = s.SdwanDistProfile
	o.SdwanErrorCorrectionProfile = s.SdwanErrorCorrectionProfile
	o.SdwanProfile = s.SdwanProfile
	o.SdwanSaasQualityProfile = s.SdwanSaasQualityProfile
	o.Misc = s.Misc
}

func (o roleDeviceWebuiObjectsSdwanXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiObjectsSdwan, error) {

	result := &RoleDeviceWebuiObjectsSdwan{
		SdwanDistProfile:            o.SdwanDistProfile,
		SdwanErrorCorrectionProfile: o.SdwanErrorCorrectionProfile,
		SdwanProfile:                o.SdwanProfile,
		SdwanSaasQualityProfile:     o.SdwanSaasQualityProfile,
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiObjectsSecurityProfilesXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiObjectsSecurityProfiles) {
	o.AntiSpyware = s.AntiSpyware
	o.Antivirus = s.Antivirus
	o.DataFiltering = s.DataFiltering
	o.DosProtection = s.DosProtection
	o.FileBlocking = s.FileBlocking
	o.GtpProtection = s.GtpProtection
	o.SctpProtection = s.SctpProtection
	o.UrlFiltering = s.UrlFiltering
	o.VulnerabilityProtection = s.VulnerabilityProtection
	o.WildfireAnalysis = s.WildfireAnalysis
	o.Misc = s.Misc
}

func (o roleDeviceWebuiObjectsSecurityProfilesXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiObjectsSecurityProfiles, error) {

	result := &RoleDeviceWebuiObjectsSecurityProfiles{
		AntiSpyware:             o.AntiSpyware,
		Antivirus:               o.Antivirus,
		DataFiltering:           o.DataFiltering,
		DosProtection:           o.DosProtection,
		FileBlocking:            o.FileBlocking,
		GtpProtection:           o.GtpProtection,
		SctpProtection:          o.SctpProtection,
		UrlFiltering:            o.UrlFiltering,
		VulnerabilityProtection: o.VulnerabilityProtection,
		WildfireAnalysis:        o.WildfireAnalysis,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiOperationsXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiOperations) {
	o.DownloadCoreFiles = s.DownloadCoreFiles
	o.DownloadPcapFiles = s.DownloadPcapFiles
	o.GenerateStatsDumpFile = s.GenerateStatsDumpFile
	o.GenerateTechSupportFile = s.GenerateTechSupportFile
	o.Reboot = s.Reboot
	o.Misc = s.Misc
}

func (o roleDeviceWebuiOperationsXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiOperations, error) {

	result := &RoleDeviceWebuiOperations{
		DownloadCoreFiles:       o.DownloadCoreFiles,
		DownloadPcapFiles:       o.DownloadPcapFiles,
		GenerateStatsDumpFile:   o.GenerateStatsDumpFile,
		GenerateTechSupportFile: o.GenerateTechSupportFile,
		Reboot:                  o.Reboot,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiPoliciesXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiPolicies) {
	o.ApplicationOverrideRulebase = s.ApplicationOverrideRulebase
	o.AuthenticationRulebase = s.AuthenticationRulebase
	o.DosRulebase = s.DosRulebase
	o.NatRulebase = s.NatRulebase
	o.NetworkPacketBrokerRulebase = s.NetworkPacketBrokerRulebase
	o.PbfRulebase = s.PbfRulebase
	o.QosRulebase = s.QosRulebase
	o.RuleHitCountReset = s.RuleHitCountReset
	o.SdwanRulebase = s.SdwanRulebase
	o.SecurityRulebase = s.SecurityRulebase
	o.SslDecryptionRulebase = s.SslDecryptionRulebase
	o.TunnelInspectRulebase = s.TunnelInspectRulebase
	o.Misc = s.Misc
}

func (o roleDeviceWebuiPoliciesXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiPolicies, error) {

	result := &RoleDeviceWebuiPolicies{
		ApplicationOverrideRulebase: o.ApplicationOverrideRulebase,
		AuthenticationRulebase:      o.AuthenticationRulebase,
		DosRulebase:                 o.DosRulebase,
		NatRulebase:                 o.NatRulebase,
		NetworkPacketBrokerRulebase: o.NetworkPacketBrokerRulebase,
		PbfRulebase:                 o.PbfRulebase,
		QosRulebase:                 o.QosRulebase,
		RuleHitCountReset:           o.RuleHitCountReset,
		SdwanRulebase:               o.SdwanRulebase,
		SecurityRulebase:            o.SecurityRulebase,
		SslDecryptionRulebase:       o.SslDecryptionRulebase,
		TunnelInspectRulebase:       o.TunnelInspectRulebase,
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiPrivacyXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiPrivacy) {
	o.ShowFullIpAddresses = s.ShowFullIpAddresses
	o.ShowUserNamesInLogsAndReports = s.ShowUserNamesInLogsAndReports
	o.ViewPcapFiles = s.ViewPcapFiles
	o.Misc = s.Misc
}

func (o roleDeviceWebuiPrivacyXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiPrivacy, error) {

	result := &RoleDeviceWebuiPrivacy{
		ShowFullIpAddresses:           o.ShowFullIpAddresses,
		ShowUserNamesInLogsAndReports: o.ShowUserNamesInLogsAndReports,
		ViewPcapFiles:                 o.ViewPcapFiles,
		Misc:                          o.Misc,
	}
	return result, nil
}
func (o *roleDeviceWebuiSaveXml_11_0_2) MarshalFromObject(s RoleDeviceWebuiSave) {
	o.ObjectLevelChanges = s.ObjectLevelChanges
	o.PartialSave = s.PartialSave
	o.SaveForOtherAdmins = s.SaveForOtherAdmins
	o.Misc = s.Misc
}

func (o roleDeviceWebuiSaveXml_11_0_2) UnmarshalToObject() (*RoleDeviceWebuiSave, error) {

	result := &RoleDeviceWebuiSave{
		ObjectLevelChanges: o.ObjectLevelChanges,
		PartialSave:        o.PartialSave,
		SaveForOtherAdmins: o.SaveForOtherAdmins,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *roleDeviceXmlapiXml_11_0_2) MarshalFromObject(s RoleDeviceXmlapi) {
	o.Commit = s.Commit
	o.Config = s.Config
	o.Export = s.Export
	o.Import = s.Import
	o.Iot = s.Iot
	o.Log = s.Log
	o.Op = s.Op
	o.Report = s.Report
	o.UserId = s.UserId
	o.Misc = s.Misc
}

func (o roleDeviceXmlapiXml_11_0_2) UnmarshalToObject() (*RoleDeviceXmlapi, error) {

	result := &RoleDeviceXmlapi{
		Commit: o.Commit,
		Config: o.Config,
		Export: o.Export,
		Import: o.Import,
		Iot:    o.Iot,
		Log:    o.Log,
		Op:     o.Op,
		Report: o.Report,
		UserId: o.UserId,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *roleVsysXml_11_0_2) MarshalFromObject(s RoleVsys) {
	o.Cli = s.Cli
	if s.Restapi != nil {
		var obj roleVsysRestapiXml_11_0_2
		obj.MarshalFromObject(*s.Restapi)
		o.Restapi = &obj
	}
	if s.Webui != nil {
		var obj roleVsysWebuiXml_11_0_2
		obj.MarshalFromObject(*s.Webui)
		o.Webui = &obj
	}
	if s.Xmlapi != nil {
		var obj roleVsysXmlapiXml_11_0_2
		obj.MarshalFromObject(*s.Xmlapi)
		o.Xmlapi = &obj
	}
	o.Misc = s.Misc
}

func (o roleVsysXml_11_0_2) UnmarshalToObject() (*RoleVsys, error) {
	var restapiVal *RoleVsysRestapi
	if o.Restapi != nil {
		obj, err := o.Restapi.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		restapiVal = obj
	}
	var webuiVal *RoleVsysWebui
	if o.Webui != nil {
		obj, err := o.Webui.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		webuiVal = obj
	}
	var xmlapiVal *RoleVsysXmlapi
	if o.Xmlapi != nil {
		obj, err := o.Xmlapi.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		xmlapiVal = obj
	}

	result := &RoleVsys{
		Cli:     o.Cli,
		Restapi: restapiVal,
		Webui:   webuiVal,
		Xmlapi:  xmlapiVal,
		Misc:    o.Misc,
	}
	return result, nil
}
func (o *roleVsysRestapiXml_11_0_2) MarshalFromObject(s RoleVsysRestapi) {
	if s.Device != nil {
		var obj roleVsysRestapiDeviceXml_11_0_2
		obj.MarshalFromObject(*s.Device)
		o.Device = &obj
	}
	if s.Network != nil {
		var obj roleVsysRestapiNetworkXml_11_0_2
		obj.MarshalFromObject(*s.Network)
		o.Network = &obj
	}
	if s.Objects != nil {
		var obj roleVsysRestapiObjectsXml_11_0_2
		obj.MarshalFromObject(*s.Objects)
		o.Objects = &obj
	}
	if s.Policies != nil {
		var obj roleVsysRestapiPoliciesXml_11_0_2
		obj.MarshalFromObject(*s.Policies)
		o.Policies = &obj
	}
	if s.System != nil {
		var obj roleVsysRestapiSystemXml_11_0_2
		obj.MarshalFromObject(*s.System)
		o.System = &obj
	}
	o.Misc = s.Misc
}

func (o roleVsysRestapiXml_11_0_2) UnmarshalToObject() (*RoleVsysRestapi, error) {
	var deviceVal *RoleVsysRestapiDevice
	if o.Device != nil {
		obj, err := o.Device.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		deviceVal = obj
	}
	var networkVal *RoleVsysRestapiNetwork
	if o.Network != nil {
		obj, err := o.Network.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		networkVal = obj
	}
	var objectsVal *RoleVsysRestapiObjects
	if o.Objects != nil {
		obj, err := o.Objects.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		objectsVal = obj
	}
	var policiesVal *RoleVsysRestapiPolicies
	if o.Policies != nil {
		obj, err := o.Policies.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		policiesVal = obj
	}
	var systemVal *RoleVsysRestapiSystem
	if o.System != nil {
		obj, err := o.System.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		systemVal = obj
	}

	result := &RoleVsysRestapi{
		Device:   deviceVal,
		Network:  networkVal,
		Objects:  objectsVal,
		Policies: policiesVal,
		System:   systemVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *roleVsysRestapiDeviceXml_11_0_2) MarshalFromObject(s RoleVsysRestapiDevice) {
	o.EmailServerProfiles = s.EmailServerProfiles
	o.HttpServerProfiles = s.HttpServerProfiles
	o.LdapServerProfiles = s.LdapServerProfiles
	o.LogInterfaceSetting = s.LogInterfaceSetting
	o.SnmpTrapServerProfiles = s.SnmpTrapServerProfiles
	o.SyslogServerProfiles = s.SyslogServerProfiles
	o.VirtualSystems = s.VirtualSystems
	o.Misc = s.Misc
}

func (o roleVsysRestapiDeviceXml_11_0_2) UnmarshalToObject() (*RoleVsysRestapiDevice, error) {

	result := &RoleVsysRestapiDevice{
		EmailServerProfiles:    o.EmailServerProfiles,
		HttpServerProfiles:     o.HttpServerProfiles,
		LdapServerProfiles:     o.LdapServerProfiles,
		LogInterfaceSetting:    o.LogInterfaceSetting,
		SnmpTrapServerProfiles: o.SnmpTrapServerProfiles,
		SyslogServerProfiles:   o.SyslogServerProfiles,
		VirtualSystems:         o.VirtualSystems,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *roleVsysRestapiNetworkXml_11_0_2) MarshalFromObject(s RoleVsysRestapiNetwork) {
	o.GlobalprotectClientlessAppGroups = s.GlobalprotectClientlessAppGroups
	o.GlobalprotectClientlessApps = s.GlobalprotectClientlessApps
	o.GlobalprotectGateways = s.GlobalprotectGateways
	o.GlobalprotectMdmServers = s.GlobalprotectMdmServers
	o.GlobalprotectPortals = s.GlobalprotectPortals
	o.SdwanInterfaceProfiles = s.SdwanInterfaceProfiles
	o.Zones = s.Zones
	o.Misc = s.Misc
}

func (o roleVsysRestapiNetworkXml_11_0_2) UnmarshalToObject() (*RoleVsysRestapiNetwork, error) {

	result := &RoleVsysRestapiNetwork{
		GlobalprotectClientlessAppGroups: o.GlobalprotectClientlessAppGroups,
		GlobalprotectClientlessApps:      o.GlobalprotectClientlessApps,
		GlobalprotectGateways:            o.GlobalprotectGateways,
		GlobalprotectMdmServers:          o.GlobalprotectMdmServers,
		GlobalprotectPortals:             o.GlobalprotectPortals,
		SdwanInterfaceProfiles:           o.SdwanInterfaceProfiles,
		Zones:                            o.Zones,
		Misc:                             o.Misc,
	}
	return result, nil
}
func (o *roleVsysRestapiObjectsXml_11_0_2) MarshalFromObject(s RoleVsysRestapiObjects) {
	o.AddressGroups = s.AddressGroups
	o.Addresses = s.Addresses
	o.AntiSpywareSecurityProfiles = s.AntiSpywareSecurityProfiles
	o.AntivirusSecurityProfiles = s.AntivirusSecurityProfiles
	o.ApplicationFilters = s.ApplicationFilters
	o.ApplicationGroups = s.ApplicationGroups
	o.Applications = s.Applications
	o.AuthenticationEnforcements = s.AuthenticationEnforcements
	o.CustomDataPatterns = s.CustomDataPatterns
	o.CustomSpywareSignatures = s.CustomSpywareSignatures
	o.CustomUrlCategories = s.CustomUrlCategories
	o.CustomVulnerabilitySignatures = s.CustomVulnerabilitySignatures
	o.DataFilteringSecurityProfiles = s.DataFilteringSecurityProfiles
	o.DecryptionProfiles = s.DecryptionProfiles
	o.Devices = s.Devices
	o.DosProtectionSecurityProfiles = s.DosProtectionSecurityProfiles
	o.DynamicUserGroups = s.DynamicUserGroups
	o.ExternalDynamicLists = s.ExternalDynamicLists
	o.FileBlockingSecurityProfiles = s.FileBlockingSecurityProfiles
	o.GlobalprotectHipObjects = s.GlobalprotectHipObjects
	o.GlobalprotectHipProfiles = s.GlobalprotectHipProfiles
	o.GtpProtectionSecurityProfiles = s.GtpProtectionSecurityProfiles
	o.LogForwardingProfiles = s.LogForwardingProfiles
	o.PacketBrokerProfiles = s.PacketBrokerProfiles
	o.Regions = s.Regions
	o.Schedules = s.Schedules
	o.SctpProtectionSecurityProfiles = s.SctpProtectionSecurityProfiles
	o.SdwanErrorCorrectionProfiles = s.SdwanErrorCorrectionProfiles
	o.SdwanPathQualityProfiles = s.SdwanPathQualityProfiles
	o.SdwanSaasQualityProfiles = s.SdwanSaasQualityProfiles
	o.SdwanTrafficDistributionProfiles = s.SdwanTrafficDistributionProfiles
	o.SecurityProfileGroups = s.SecurityProfileGroups
	o.ServiceGroups = s.ServiceGroups
	o.Services = s.Services
	o.Tags = s.Tags
	o.UrlFilteringSecurityProfiles = s.UrlFilteringSecurityProfiles
	o.VulnerabilityProtectionSecurityProfiles = s.VulnerabilityProtectionSecurityProfiles
	o.WildfireAnalysisSecurityProfiles = s.WildfireAnalysisSecurityProfiles
	o.Misc = s.Misc
}

func (o roleVsysRestapiObjectsXml_11_0_2) UnmarshalToObject() (*RoleVsysRestapiObjects, error) {

	result := &RoleVsysRestapiObjects{
		AddressGroups:                           o.AddressGroups,
		Addresses:                               o.Addresses,
		AntiSpywareSecurityProfiles:             o.AntiSpywareSecurityProfiles,
		AntivirusSecurityProfiles:               o.AntivirusSecurityProfiles,
		ApplicationFilters:                      o.ApplicationFilters,
		ApplicationGroups:                       o.ApplicationGroups,
		Applications:                            o.Applications,
		AuthenticationEnforcements:              o.AuthenticationEnforcements,
		CustomDataPatterns:                      o.CustomDataPatterns,
		CustomSpywareSignatures:                 o.CustomSpywareSignatures,
		CustomUrlCategories:                     o.CustomUrlCategories,
		CustomVulnerabilitySignatures:           o.CustomVulnerabilitySignatures,
		DataFilteringSecurityProfiles:           o.DataFilteringSecurityProfiles,
		DecryptionProfiles:                      o.DecryptionProfiles,
		Devices:                                 o.Devices,
		DosProtectionSecurityProfiles:           o.DosProtectionSecurityProfiles,
		DynamicUserGroups:                       o.DynamicUserGroups,
		ExternalDynamicLists:                    o.ExternalDynamicLists,
		FileBlockingSecurityProfiles:            o.FileBlockingSecurityProfiles,
		GlobalprotectHipObjects:                 o.GlobalprotectHipObjects,
		GlobalprotectHipProfiles:                o.GlobalprotectHipProfiles,
		GtpProtectionSecurityProfiles:           o.GtpProtectionSecurityProfiles,
		LogForwardingProfiles:                   o.LogForwardingProfiles,
		PacketBrokerProfiles:                    o.PacketBrokerProfiles,
		Regions:                                 o.Regions,
		Schedules:                               o.Schedules,
		SctpProtectionSecurityProfiles:          o.SctpProtectionSecurityProfiles,
		SdwanErrorCorrectionProfiles:            o.SdwanErrorCorrectionProfiles,
		SdwanPathQualityProfiles:                o.SdwanPathQualityProfiles,
		SdwanSaasQualityProfiles:                o.SdwanSaasQualityProfiles,
		SdwanTrafficDistributionProfiles:        o.SdwanTrafficDistributionProfiles,
		SecurityProfileGroups:                   o.SecurityProfileGroups,
		ServiceGroups:                           o.ServiceGroups,
		Services:                                o.Services,
		Tags:                                    o.Tags,
		UrlFilteringSecurityProfiles:            o.UrlFilteringSecurityProfiles,
		VulnerabilityProtectionSecurityProfiles: o.VulnerabilityProtectionSecurityProfiles,
		WildfireAnalysisSecurityProfiles:        o.WildfireAnalysisSecurityProfiles,
		Misc:                                    o.Misc,
	}
	return result, nil
}
func (o *roleVsysRestapiPoliciesXml_11_0_2) MarshalFromObject(s RoleVsysRestapiPolicies) {
	o.ApplicationOverrideRules = s.ApplicationOverrideRules
	o.AuthenticationRules = s.AuthenticationRules
	o.DecryptionRules = s.DecryptionRules
	o.DosRules = s.DosRules
	o.NatRules = s.NatRules
	o.NetworkPacketBrokerRules = s.NetworkPacketBrokerRules
	o.PolicyBasedForwardingRules = s.PolicyBasedForwardingRules
	o.QosRules = s.QosRules
	o.SdwanRules = s.SdwanRules
	o.SecurityRules = s.SecurityRules
	o.TunnelInspectionRules = s.TunnelInspectionRules
	o.Misc = s.Misc
}

func (o roleVsysRestapiPoliciesXml_11_0_2) UnmarshalToObject() (*RoleVsysRestapiPolicies, error) {

	result := &RoleVsysRestapiPolicies{
		ApplicationOverrideRules:   o.ApplicationOverrideRules,
		AuthenticationRules:        o.AuthenticationRules,
		DecryptionRules:            o.DecryptionRules,
		DosRules:                   o.DosRules,
		NatRules:                   o.NatRules,
		NetworkPacketBrokerRules:   o.NetworkPacketBrokerRules,
		PolicyBasedForwardingRules: o.PolicyBasedForwardingRules,
		QosRules:                   o.QosRules,
		SdwanRules:                 o.SdwanRules,
		SecurityRules:              o.SecurityRules,
		TunnelInspectionRules:      o.TunnelInspectionRules,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *roleVsysRestapiSystemXml_11_0_2) MarshalFromObject(s RoleVsysRestapiSystem) {
	o.Configuration = s.Configuration
	o.Misc = s.Misc
}

func (o roleVsysRestapiSystemXml_11_0_2) UnmarshalToObject() (*RoleVsysRestapiSystem, error) {

	result := &RoleVsysRestapiSystem{
		Configuration: o.Configuration,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiXml_11_0_2) MarshalFromObject(s RoleVsysWebui) {
	o.Acc = s.Acc
	if s.Commit != nil {
		var obj roleVsysWebuiCommitXml_11_0_2
		obj.MarshalFromObject(*s.Commit)
		o.Commit = &obj
	}
	o.Dashboard = s.Dashboard
	if s.Device != nil {
		var obj roleVsysWebuiDeviceXml_11_0_2
		obj.MarshalFromObject(*s.Device)
		o.Device = &obj
	}
	if s.Monitor != nil {
		var obj roleVsysWebuiMonitorXml_11_0_2
		obj.MarshalFromObject(*s.Monitor)
		o.Monitor = &obj
	}
	if s.Network != nil {
		var obj roleVsysWebuiNetworkXml_11_0_2
		obj.MarshalFromObject(*s.Network)
		o.Network = &obj
	}
	if s.Objects != nil {
		var obj roleVsysWebuiObjectsXml_11_0_2
		obj.MarshalFromObject(*s.Objects)
		o.Objects = &obj
	}
	if s.Operations != nil {
		var obj roleVsysWebuiOperationsXml_11_0_2
		obj.MarshalFromObject(*s.Operations)
		o.Operations = &obj
	}
	if s.Policies != nil {
		var obj roleVsysWebuiPoliciesXml_11_0_2
		obj.MarshalFromObject(*s.Policies)
		o.Policies = &obj
	}
	if s.Privacy != nil {
		var obj roleVsysWebuiPrivacyXml_11_0_2
		obj.MarshalFromObject(*s.Privacy)
		o.Privacy = &obj
	}
	if s.Save != nil {
		var obj roleVsysWebuiSaveXml_11_0_2
		obj.MarshalFromObject(*s.Save)
		o.Save = &obj
	}
	o.Tasks = s.Tasks
	o.Validate = s.Validate
	o.Misc = s.Misc
}

func (o roleVsysWebuiXml_11_0_2) UnmarshalToObject() (*RoleVsysWebui, error) {
	var commitVal *RoleVsysWebuiCommit
	if o.Commit != nil {
		obj, err := o.Commit.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		commitVal = obj
	}
	var deviceVal *RoleVsysWebuiDevice
	if o.Device != nil {
		obj, err := o.Device.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		deviceVal = obj
	}
	var monitorVal *RoleVsysWebuiMonitor
	if o.Monitor != nil {
		obj, err := o.Monitor.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		monitorVal = obj
	}
	var networkVal *RoleVsysWebuiNetwork
	if o.Network != nil {
		obj, err := o.Network.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		networkVal = obj
	}
	var objectsVal *RoleVsysWebuiObjects
	if o.Objects != nil {
		obj, err := o.Objects.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		objectsVal = obj
	}
	var operationsVal *RoleVsysWebuiOperations
	if o.Operations != nil {
		obj, err := o.Operations.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		operationsVal = obj
	}
	var policiesVal *RoleVsysWebuiPolicies
	if o.Policies != nil {
		obj, err := o.Policies.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		policiesVal = obj
	}
	var privacyVal *RoleVsysWebuiPrivacy
	if o.Privacy != nil {
		obj, err := o.Privacy.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		privacyVal = obj
	}
	var saveVal *RoleVsysWebuiSave
	if o.Save != nil {
		obj, err := o.Save.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		saveVal = obj
	}

	result := &RoleVsysWebui{
		Acc:        o.Acc,
		Commit:     commitVal,
		Dashboard:  o.Dashboard,
		Device:     deviceVal,
		Monitor:    monitorVal,
		Network:    networkVal,
		Objects:    objectsVal,
		Operations: operationsVal,
		Policies:   policiesVal,
		Privacy:    privacyVal,
		Save:       saveVal,
		Tasks:      o.Tasks,
		Validate:   o.Validate,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiCommitXml_11_0_2) MarshalFromObject(s RoleVsysWebuiCommit) {
	o.CommitForOtherAdmins = s.CommitForOtherAdmins
	o.VirtualSystems = s.VirtualSystems
	o.Misc = s.Misc
}

func (o roleVsysWebuiCommitXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiCommit, error) {

	result := &RoleVsysWebuiCommit{
		CommitForOtherAdmins: o.CommitForOtherAdmins,
		VirtualSystems:       o.VirtualSystems,
		Misc:                 o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiDeviceXml_11_0_2) MarshalFromObject(s RoleVsysWebuiDevice) {
	o.Administrators = s.Administrators
	o.AuthenticationProfile = s.AuthenticationProfile
	o.AuthenticationSequence = s.AuthenticationSequence
	o.BlockPages = s.BlockPages
	if s.CertificateManagement != nil {
		var obj roleVsysWebuiDeviceCertificateManagementXml_11_0_2
		obj.MarshalFromObject(*s.CertificateManagement)
		o.CertificateManagement = &obj
	}
	o.DataRedistribution = s.DataRedistribution
	o.DeviceQuarantine = s.DeviceQuarantine
	o.DhcpSyslogServer = s.DhcpSyslogServer
	if s.LocalUserDatabase != nil {
		var obj roleVsysWebuiDeviceLocalUserDatabaseXml_11_0_2
		obj.MarshalFromObject(*s.LocalUserDatabase)
		o.LocalUserDatabase = &obj
	}
	if s.LogSettings != nil {
		var obj roleVsysWebuiDeviceLogSettingsXml_11_0_2
		obj.MarshalFromObject(*s.LogSettings)
		o.LogSettings = &obj
	}
	if s.PolicyRecommendations != nil {
		var obj roleVsysWebuiDevicePolicyRecommendationsXml_11_0_2
		obj.MarshalFromObject(*s.PolicyRecommendations)
		o.PolicyRecommendations = &obj
	}
	if s.ServerProfile != nil {
		var obj roleVsysWebuiDeviceServerProfileXml_11_0_2
		obj.MarshalFromObject(*s.ServerProfile)
		o.ServerProfile = &obj
	}
	if s.Setup != nil {
		var obj roleVsysWebuiDeviceSetupXml_11_0_2
		obj.MarshalFromObject(*s.Setup)
		o.Setup = &obj
	}
	o.Troubleshooting = s.Troubleshooting
	o.UserIdentification = s.UserIdentification
	o.VmInfoSource = s.VmInfoSource
	o.Misc = s.Misc
}

func (o roleVsysWebuiDeviceXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiDevice, error) {
	var certificateManagementVal *RoleVsysWebuiDeviceCertificateManagement
	if o.CertificateManagement != nil {
		obj, err := o.CertificateManagement.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		certificateManagementVal = obj
	}
	var localUserDatabaseVal *RoleVsysWebuiDeviceLocalUserDatabase
	if o.LocalUserDatabase != nil {
		obj, err := o.LocalUserDatabase.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localUserDatabaseVal = obj
	}
	var logSettingsVal *RoleVsysWebuiDeviceLogSettings
	if o.LogSettings != nil {
		obj, err := o.LogSettings.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		logSettingsVal = obj
	}
	var policyRecommendationsVal *RoleVsysWebuiDevicePolicyRecommendations
	if o.PolicyRecommendations != nil {
		obj, err := o.PolicyRecommendations.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		policyRecommendationsVal = obj
	}
	var serverProfileVal *RoleVsysWebuiDeviceServerProfile
	if o.ServerProfile != nil {
		obj, err := o.ServerProfile.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		serverProfileVal = obj
	}
	var setupVal *RoleVsysWebuiDeviceSetup
	if o.Setup != nil {
		obj, err := o.Setup.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		setupVal = obj
	}

	result := &RoleVsysWebuiDevice{
		Administrators:         o.Administrators,
		AuthenticationProfile:  o.AuthenticationProfile,
		AuthenticationSequence: o.AuthenticationSequence,
		BlockPages:             o.BlockPages,
		CertificateManagement:  certificateManagementVal,
		DataRedistribution:     o.DataRedistribution,
		DeviceQuarantine:       o.DeviceQuarantine,
		DhcpSyslogServer:       o.DhcpSyslogServer,
		LocalUserDatabase:      localUserDatabaseVal,
		LogSettings:            logSettingsVal,
		PolicyRecommendations:  policyRecommendationsVal,
		ServerProfile:          serverProfileVal,
		Setup:                  setupVal,
		Troubleshooting:        o.Troubleshooting,
		UserIdentification:     o.UserIdentification,
		VmInfoSource:           o.VmInfoSource,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiDeviceCertificateManagementXml_11_0_2) MarshalFromObject(s RoleVsysWebuiDeviceCertificateManagement) {
	o.CertificateProfile = s.CertificateProfile
	o.Certificates = s.Certificates
	o.OcspResponder = s.OcspResponder
	o.Scep = s.Scep
	o.SshServiceProfile = s.SshServiceProfile
	o.SslDecryptionExclusion = s.SslDecryptionExclusion
	o.SslTlsServiceProfile = s.SslTlsServiceProfile
	o.Misc = s.Misc
}

func (o roleVsysWebuiDeviceCertificateManagementXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiDeviceCertificateManagement, error) {

	result := &RoleVsysWebuiDeviceCertificateManagement{
		CertificateProfile:     o.CertificateProfile,
		Certificates:           o.Certificates,
		OcspResponder:          o.OcspResponder,
		Scep:                   o.Scep,
		SshServiceProfile:      o.SshServiceProfile,
		SslDecryptionExclusion: o.SslDecryptionExclusion,
		SslTlsServiceProfile:   o.SslTlsServiceProfile,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiDeviceLocalUserDatabaseXml_11_0_2) MarshalFromObject(s RoleVsysWebuiDeviceLocalUserDatabase) {
	o.UserGroups = s.UserGroups
	o.Users = s.Users
	o.Misc = s.Misc
}

func (o roleVsysWebuiDeviceLocalUserDatabaseXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiDeviceLocalUserDatabase, error) {

	result := &RoleVsysWebuiDeviceLocalUserDatabase{
		UserGroups: o.UserGroups,
		Users:      o.Users,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiDeviceLogSettingsXml_11_0_2) MarshalFromObject(s RoleVsysWebuiDeviceLogSettings) {
	o.Config = s.Config
	o.Correlation = s.Correlation
	o.Globalprotect = s.Globalprotect
	o.Hipmatch = s.Hipmatch
	o.Iptag = s.Iptag
	o.System = s.System
	o.UserId = s.UserId
	o.Misc = s.Misc
}

func (o roleVsysWebuiDeviceLogSettingsXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiDeviceLogSettings, error) {

	result := &RoleVsysWebuiDeviceLogSettings{
		Config:        o.Config,
		Correlation:   o.Correlation,
		Globalprotect: o.Globalprotect,
		Hipmatch:      o.Hipmatch,
		Iptag:         o.Iptag,
		System:        o.System,
		UserId:        o.UserId,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiDevicePolicyRecommendationsXml_11_0_2) MarshalFromObject(s RoleVsysWebuiDevicePolicyRecommendations) {
	o.Iot = s.Iot
	o.Saas = s.Saas
	o.Misc = s.Misc
}

func (o roleVsysWebuiDevicePolicyRecommendationsXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiDevicePolicyRecommendations, error) {

	result := &RoleVsysWebuiDevicePolicyRecommendations{
		Iot:  o.Iot,
		Saas: o.Saas,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiDeviceServerProfileXml_11_0_2) MarshalFromObject(s RoleVsysWebuiDeviceServerProfile) {
	o.Dns = s.Dns
	o.Email = s.Email
	o.Http = s.Http
	o.Kerberos = s.Kerberos
	o.Ldap = s.Ldap
	o.Mfa = s.Mfa
	o.Netflow = s.Netflow
	o.Radius = s.Radius
	o.SamlIdp = s.SamlIdp
	o.Scp = s.Scp
	o.SnmpTrap = s.SnmpTrap
	o.Syslog = s.Syslog
	o.Tacplus = s.Tacplus
	o.Misc = s.Misc
}

func (o roleVsysWebuiDeviceServerProfileXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiDeviceServerProfile, error) {

	result := &RoleVsysWebuiDeviceServerProfile{
		Dns:      o.Dns,
		Email:    o.Email,
		Http:     o.Http,
		Kerberos: o.Kerberos,
		Ldap:     o.Ldap,
		Mfa:      o.Mfa,
		Netflow:  o.Netflow,
		Radius:   o.Radius,
		SamlIdp:  o.SamlIdp,
		Scp:      o.Scp,
		SnmpTrap: o.SnmpTrap,
		Syslog:   o.Syslog,
		Tacplus:  o.Tacplus,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiDeviceSetupXml_11_0_2) MarshalFromObject(s RoleVsysWebuiDeviceSetup) {
	o.ContentId = s.ContentId
	o.Hsm = s.Hsm
	o.Interfaces = s.Interfaces
	o.Management = s.Management
	o.Operations = s.Operations
	o.Services = s.Services
	o.Session = s.Session
	o.Telemetry = s.Telemetry
	o.Wildfire = s.Wildfire
	o.Misc = s.Misc
}

func (o roleVsysWebuiDeviceSetupXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiDeviceSetup, error) {

	result := &RoleVsysWebuiDeviceSetup{
		ContentId:  o.ContentId,
		Hsm:        o.Hsm,
		Interfaces: o.Interfaces,
		Management: o.Management,
		Operations: o.Operations,
		Services:   o.Services,
		Session:    o.Session,
		Telemetry:  o.Telemetry,
		Wildfire:   o.Wildfire,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiMonitorXml_11_0_2) MarshalFromObject(s RoleVsysWebuiMonitor) {
	o.AppScope = s.AppScope
	if s.AutomatedCorrelationEngine != nil {
		var obj roleVsysWebuiMonitorAutomatedCorrelationEngineXml_11_0_2
		obj.MarshalFromObject(*s.AutomatedCorrelationEngine)
		o.AutomatedCorrelationEngine = &obj
	}
	o.BlockIpList = s.BlockIpList
	if s.CustomReports != nil {
		var obj roleVsysWebuiMonitorCustomReportsXml_11_0_2
		obj.MarshalFromObject(*s.CustomReports)
		o.CustomReports = &obj
	}
	o.ExternalLogs = s.ExternalLogs
	if s.Logs != nil {
		var obj roleVsysWebuiMonitorLogsXml_11_0_2
		obj.MarshalFromObject(*s.Logs)
		o.Logs = &obj
	}
	if s.PdfReports != nil {
		var obj roleVsysWebuiMonitorPdfReportsXml_11_0_2
		obj.MarshalFromObject(*s.PdfReports)
		o.PdfReports = &obj
	}
	o.SessionBrowser = s.SessionBrowser
	o.ViewCustomReports = s.ViewCustomReports
	o.Misc = s.Misc
}

func (o roleVsysWebuiMonitorXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiMonitor, error) {
	var automatedCorrelationEngineVal *RoleVsysWebuiMonitorAutomatedCorrelationEngine
	if o.AutomatedCorrelationEngine != nil {
		obj, err := o.AutomatedCorrelationEngine.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		automatedCorrelationEngineVal = obj
	}
	var customReportsVal *RoleVsysWebuiMonitorCustomReports
	if o.CustomReports != nil {
		obj, err := o.CustomReports.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		customReportsVal = obj
	}
	var logsVal *RoleVsysWebuiMonitorLogs
	if o.Logs != nil {
		obj, err := o.Logs.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		logsVal = obj
	}
	var pdfReportsVal *RoleVsysWebuiMonitorPdfReports
	if o.PdfReports != nil {
		obj, err := o.PdfReports.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		pdfReportsVal = obj
	}

	result := &RoleVsysWebuiMonitor{
		AppScope:                   o.AppScope,
		AutomatedCorrelationEngine: automatedCorrelationEngineVal,
		BlockIpList:                o.BlockIpList,
		CustomReports:              customReportsVal,
		ExternalLogs:               o.ExternalLogs,
		Logs:                       logsVal,
		PdfReports:                 pdfReportsVal,
		SessionBrowser:             o.SessionBrowser,
		ViewCustomReports:          o.ViewCustomReports,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiMonitorAutomatedCorrelationEngineXml_11_0_2) MarshalFromObject(s RoleVsysWebuiMonitorAutomatedCorrelationEngine) {
	o.CorrelatedEvents = s.CorrelatedEvents
	o.CorrelationObjects = s.CorrelationObjects
	o.Misc = s.Misc
}

func (o roleVsysWebuiMonitorAutomatedCorrelationEngineXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiMonitorAutomatedCorrelationEngine, error) {

	result := &RoleVsysWebuiMonitorAutomatedCorrelationEngine{
		CorrelatedEvents:   o.CorrelatedEvents,
		CorrelationObjects: o.CorrelationObjects,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiMonitorCustomReportsXml_11_0_2) MarshalFromObject(s RoleVsysWebuiMonitorCustomReports) {
	o.ApplicationStatistics = s.ApplicationStatistics
	o.Auth = s.Auth
	o.DataFilteringLog = s.DataFilteringLog
	o.DecryptionLog = s.DecryptionLog
	o.DecryptionSummary = s.DecryptionSummary
	o.Globalprotect = s.Globalprotect
	o.GtpLog = s.GtpLog
	o.GtpSummary = s.GtpSummary
	o.Hipmatch = s.Hipmatch
	o.Iptag = s.Iptag
	o.SctpLog = s.SctpLog
	o.SctpSummary = s.SctpSummary
	o.ThreatLog = s.ThreatLog
	o.ThreatSummary = s.ThreatSummary
	o.TrafficLog = s.TrafficLog
	o.TrafficSummary = s.TrafficSummary
	o.TunnelLog = s.TunnelLog
	o.TunnelSummary = s.TunnelSummary
	o.UrlLog = s.UrlLog
	o.UrlSummary = s.UrlSummary
	o.Userid = s.Userid
	o.WildfireLog = s.WildfireLog
	o.Misc = s.Misc
}

func (o roleVsysWebuiMonitorCustomReportsXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiMonitorCustomReports, error) {

	result := &RoleVsysWebuiMonitorCustomReports{
		ApplicationStatistics: o.ApplicationStatistics,
		Auth:                  o.Auth,
		DataFilteringLog:      o.DataFilteringLog,
		DecryptionLog:         o.DecryptionLog,
		DecryptionSummary:     o.DecryptionSummary,
		Globalprotect:         o.Globalprotect,
		GtpLog:                o.GtpLog,
		GtpSummary:            o.GtpSummary,
		Hipmatch:              o.Hipmatch,
		Iptag:                 o.Iptag,
		SctpLog:               o.SctpLog,
		SctpSummary:           o.SctpSummary,
		ThreatLog:             o.ThreatLog,
		ThreatSummary:         o.ThreatSummary,
		TrafficLog:            o.TrafficLog,
		TrafficSummary:        o.TrafficSummary,
		TunnelLog:             o.TunnelLog,
		TunnelSummary:         o.TunnelSummary,
		UrlLog:                o.UrlLog,
		UrlSummary:            o.UrlSummary,
		Userid:                o.Userid,
		WildfireLog:           o.WildfireLog,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiMonitorLogsXml_11_0_2) MarshalFromObject(s RoleVsysWebuiMonitorLogs) {
	o.Authentication = s.Authentication
	o.DataFiltering = s.DataFiltering
	o.Decryption = s.Decryption
	o.Globalprotect = s.Globalprotect
	o.Gtp = s.Gtp
	o.Hipmatch = s.Hipmatch
	o.Iptag = s.Iptag
	o.Sctp = s.Sctp
	o.Threat = s.Threat
	o.Traffic = s.Traffic
	o.Tunnel = s.Tunnel
	o.Url = s.Url
	o.Userid = s.Userid
	o.Wildfire = s.Wildfire
	o.Misc = s.Misc
}

func (o roleVsysWebuiMonitorLogsXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiMonitorLogs, error) {

	result := &RoleVsysWebuiMonitorLogs{
		Authentication: o.Authentication,
		DataFiltering:  o.DataFiltering,
		Decryption:     o.Decryption,
		Globalprotect:  o.Globalprotect,
		Gtp:            o.Gtp,
		Hipmatch:       o.Hipmatch,
		Iptag:          o.Iptag,
		Sctp:           o.Sctp,
		Threat:         o.Threat,
		Traffic:        o.Traffic,
		Tunnel:         o.Tunnel,
		Url:            o.Url,
		Userid:         o.Userid,
		Wildfire:       o.Wildfire,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiMonitorPdfReportsXml_11_0_2) MarshalFromObject(s RoleVsysWebuiMonitorPdfReports) {
	o.EmailScheduler = s.EmailScheduler
	o.ManagePdfSummary = s.ManagePdfSummary
	o.PdfSummaryReports = s.PdfSummaryReports
	o.ReportGroups = s.ReportGroups
	o.SaasApplicationUsageReport = s.SaasApplicationUsageReport
	o.UserActivityReport = s.UserActivityReport
	o.Misc = s.Misc
}

func (o roleVsysWebuiMonitorPdfReportsXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiMonitorPdfReports, error) {

	result := &RoleVsysWebuiMonitorPdfReports{
		EmailScheduler:             o.EmailScheduler,
		ManagePdfSummary:           o.ManagePdfSummary,
		PdfSummaryReports:          o.PdfSummaryReports,
		ReportGroups:               o.ReportGroups,
		SaasApplicationUsageReport: o.SaasApplicationUsageReport,
		UserActivityReport:         o.UserActivityReport,
		Misc:                       o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiNetworkXml_11_0_2) MarshalFromObject(s RoleVsysWebuiNetwork) {
	if s.GlobalProtect != nil {
		var obj roleVsysWebuiNetworkGlobalProtectXml_11_0_2
		obj.MarshalFromObject(*s.GlobalProtect)
		o.GlobalProtect = &obj
	}
	o.SdwanInterfaceProfile = s.SdwanInterfaceProfile
	o.Zones = s.Zones
	o.Misc = s.Misc
}

func (o roleVsysWebuiNetworkXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiNetwork, error) {
	var globalProtectVal *RoleVsysWebuiNetworkGlobalProtect
	if o.GlobalProtect != nil {
		obj, err := o.GlobalProtect.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalProtectVal = obj
	}

	result := &RoleVsysWebuiNetwork{
		GlobalProtect:         globalProtectVal,
		SdwanInterfaceProfile: o.SdwanInterfaceProfile,
		Zones:                 o.Zones,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiNetworkGlobalProtectXml_11_0_2) MarshalFromObject(s RoleVsysWebuiNetworkGlobalProtect) {
	o.ClientlessAppGroups = s.ClientlessAppGroups
	o.ClientlessApps = s.ClientlessApps
	o.Gateways = s.Gateways
	o.Mdm = s.Mdm
	o.Portals = s.Portals
	o.Misc = s.Misc
}

func (o roleVsysWebuiNetworkGlobalProtectXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiNetworkGlobalProtect, error) {

	result := &RoleVsysWebuiNetworkGlobalProtect{
		ClientlessAppGroups: o.ClientlessAppGroups,
		ClientlessApps:      o.ClientlessApps,
		Gateways:            o.Gateways,
		Mdm:                 o.Mdm,
		Portals:             o.Portals,
		Misc:                o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiObjectsXml_11_0_2) MarshalFromObject(s RoleVsysWebuiObjects) {
	o.AddressGroups = s.AddressGroups
	o.Addresses = s.Addresses
	o.ApplicationFilters = s.ApplicationFilters
	o.ApplicationGroups = s.ApplicationGroups
	o.Applications = s.Applications
	o.Authentication = s.Authentication
	if s.CustomObjects != nil {
		var obj roleVsysWebuiObjectsCustomObjectsXml_11_0_2
		obj.MarshalFromObject(*s.CustomObjects)
		o.CustomObjects = &obj
	}
	if s.Decryption != nil {
		var obj roleVsysWebuiObjectsDecryptionXml_11_0_2
		obj.MarshalFromObject(*s.Decryption)
		o.Decryption = &obj
	}
	o.Devices = s.Devices
	o.DynamicBlockLists = s.DynamicBlockLists
	o.DynamicUserGroups = s.DynamicUserGroups
	if s.GlobalProtect != nil {
		var obj roleVsysWebuiObjectsGlobalProtectXml_11_0_2
		obj.MarshalFromObject(*s.GlobalProtect)
		o.GlobalProtect = &obj
	}
	o.LogForwarding = s.LogForwarding
	o.PacketBrokerProfile = s.PacketBrokerProfile
	o.Regions = s.Regions
	o.Schedules = s.Schedules
	if s.Sdwan != nil {
		var obj roleVsysWebuiObjectsSdwanXml_11_0_2
		obj.MarshalFromObject(*s.Sdwan)
		o.Sdwan = &obj
	}
	o.SecurityProfileGroups = s.SecurityProfileGroups
	if s.SecurityProfiles != nil {
		var obj roleVsysWebuiObjectsSecurityProfilesXml_11_0_2
		obj.MarshalFromObject(*s.SecurityProfiles)
		o.SecurityProfiles = &obj
	}
	o.ServiceGroups = s.ServiceGroups
	o.Services = s.Services
	o.Tags = s.Tags
	o.Misc = s.Misc
}

func (o roleVsysWebuiObjectsXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiObjects, error) {
	var customObjectsVal *RoleVsysWebuiObjectsCustomObjects
	if o.CustomObjects != nil {
		obj, err := o.CustomObjects.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		customObjectsVal = obj
	}
	var decryptionVal *RoleVsysWebuiObjectsDecryption
	if o.Decryption != nil {
		obj, err := o.Decryption.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		decryptionVal = obj
	}
	var globalProtectVal *RoleVsysWebuiObjectsGlobalProtect
	if o.GlobalProtect != nil {
		obj, err := o.GlobalProtect.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalProtectVal = obj
	}
	var sdwanVal *RoleVsysWebuiObjectsSdwan
	if o.Sdwan != nil {
		obj, err := o.Sdwan.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sdwanVal = obj
	}
	var securityProfilesVal *RoleVsysWebuiObjectsSecurityProfiles
	if o.SecurityProfiles != nil {
		obj, err := o.SecurityProfiles.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		securityProfilesVal = obj
	}

	result := &RoleVsysWebuiObjects{
		AddressGroups:         o.AddressGroups,
		Addresses:             o.Addresses,
		ApplicationFilters:    o.ApplicationFilters,
		ApplicationGroups:     o.ApplicationGroups,
		Applications:          o.Applications,
		Authentication:        o.Authentication,
		CustomObjects:         customObjectsVal,
		Decryption:            decryptionVal,
		Devices:               o.Devices,
		DynamicBlockLists:     o.DynamicBlockLists,
		DynamicUserGroups:     o.DynamicUserGroups,
		GlobalProtect:         globalProtectVal,
		LogForwarding:         o.LogForwarding,
		PacketBrokerProfile:   o.PacketBrokerProfile,
		Regions:               o.Regions,
		Schedules:             o.Schedules,
		Sdwan:                 sdwanVal,
		SecurityProfileGroups: o.SecurityProfileGroups,
		SecurityProfiles:      securityProfilesVal,
		ServiceGroups:         o.ServiceGroups,
		Services:              o.Services,
		Tags:                  o.Tags,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiObjectsCustomObjectsXml_11_0_2) MarshalFromObject(s RoleVsysWebuiObjectsCustomObjects) {
	o.DataPatterns = s.DataPatterns
	o.Spyware = s.Spyware
	o.UrlCategory = s.UrlCategory
	o.Vulnerability = s.Vulnerability
	o.Misc = s.Misc
}

func (o roleVsysWebuiObjectsCustomObjectsXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiObjectsCustomObjects, error) {

	result := &RoleVsysWebuiObjectsCustomObjects{
		DataPatterns:  o.DataPatterns,
		Spyware:       o.Spyware,
		UrlCategory:   o.UrlCategory,
		Vulnerability: o.Vulnerability,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiObjectsDecryptionXml_11_0_2) MarshalFromObject(s RoleVsysWebuiObjectsDecryption) {
	o.DecryptionProfile = s.DecryptionProfile
	o.Misc = s.Misc
}

func (o roleVsysWebuiObjectsDecryptionXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiObjectsDecryption, error) {

	result := &RoleVsysWebuiObjectsDecryption{
		DecryptionProfile: o.DecryptionProfile,
		Misc:              o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiObjectsGlobalProtectXml_11_0_2) MarshalFromObject(s RoleVsysWebuiObjectsGlobalProtect) {
	o.HipObjects = s.HipObjects
	o.HipProfiles = s.HipProfiles
	o.Misc = s.Misc
}

func (o roleVsysWebuiObjectsGlobalProtectXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiObjectsGlobalProtect, error) {

	result := &RoleVsysWebuiObjectsGlobalProtect{
		HipObjects:  o.HipObjects,
		HipProfiles: o.HipProfiles,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiObjectsSdwanXml_11_0_2) MarshalFromObject(s RoleVsysWebuiObjectsSdwan) {
	o.SdwanDistProfile = s.SdwanDistProfile
	o.SdwanErrorCorrectionProfile = s.SdwanErrorCorrectionProfile
	o.SdwanProfile = s.SdwanProfile
	o.SdwanSaasQualityProfile = s.SdwanSaasQualityProfile
	o.Misc = s.Misc
}

func (o roleVsysWebuiObjectsSdwanXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiObjectsSdwan, error) {

	result := &RoleVsysWebuiObjectsSdwan{
		SdwanDistProfile:            o.SdwanDistProfile,
		SdwanErrorCorrectionProfile: o.SdwanErrorCorrectionProfile,
		SdwanProfile:                o.SdwanProfile,
		SdwanSaasQualityProfile:     o.SdwanSaasQualityProfile,
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiObjectsSecurityProfilesXml_11_0_2) MarshalFromObject(s RoleVsysWebuiObjectsSecurityProfiles) {
	o.AntiSpyware = s.AntiSpyware
	o.Antivirus = s.Antivirus
	o.DataFiltering = s.DataFiltering
	o.DosProtection = s.DosProtection
	o.FileBlocking = s.FileBlocking
	o.GtpProtection = s.GtpProtection
	o.SctpProtection = s.SctpProtection
	o.UrlFiltering = s.UrlFiltering
	o.VulnerabilityProtection = s.VulnerabilityProtection
	o.WildfireAnalysis = s.WildfireAnalysis
	o.Misc = s.Misc
}

func (o roleVsysWebuiObjectsSecurityProfilesXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiObjectsSecurityProfiles, error) {

	result := &RoleVsysWebuiObjectsSecurityProfiles{
		AntiSpyware:             o.AntiSpyware,
		Antivirus:               o.Antivirus,
		DataFiltering:           o.DataFiltering,
		DosProtection:           o.DosProtection,
		FileBlocking:            o.FileBlocking,
		GtpProtection:           o.GtpProtection,
		SctpProtection:          o.SctpProtection,
		UrlFiltering:            o.UrlFiltering,
		VulnerabilityProtection: o.VulnerabilityProtection,
		WildfireAnalysis:        o.WildfireAnalysis,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiOperationsXml_11_0_2) MarshalFromObject(s RoleVsysWebuiOperations) {
	o.DownloadCoreFiles = s.DownloadCoreFiles
	o.DownloadPcapFiles = s.DownloadPcapFiles
	o.GenerateStatsDumpFile = s.GenerateStatsDumpFile
	o.GenerateTechSupportFile = s.GenerateTechSupportFile
	o.Reboot = s.Reboot
	o.Misc = s.Misc
}

func (o roleVsysWebuiOperationsXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiOperations, error) {

	result := &RoleVsysWebuiOperations{
		DownloadCoreFiles:       o.DownloadCoreFiles,
		DownloadPcapFiles:       o.DownloadPcapFiles,
		GenerateStatsDumpFile:   o.GenerateStatsDumpFile,
		GenerateTechSupportFile: o.GenerateTechSupportFile,
		Reboot:                  o.Reboot,
		Misc:                    o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiPoliciesXml_11_0_2) MarshalFromObject(s RoleVsysWebuiPolicies) {
	o.ApplicationOverrideRulebase = s.ApplicationOverrideRulebase
	o.AuthenticationRulebase = s.AuthenticationRulebase
	o.DosRulebase = s.DosRulebase
	o.NatRulebase = s.NatRulebase
	o.NetworkPacketBrokerRulebase = s.NetworkPacketBrokerRulebase
	o.PbfRulebase = s.PbfRulebase
	o.QosRulebase = s.QosRulebase
	o.RuleHitCountReset = s.RuleHitCountReset
	o.SdwanRulebase = s.SdwanRulebase
	o.SecurityRulebase = s.SecurityRulebase
	o.SslDecryptionRulebase = s.SslDecryptionRulebase
	o.TunnelInspectRulebase = s.TunnelInspectRulebase
	o.Misc = s.Misc
}

func (o roleVsysWebuiPoliciesXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiPolicies, error) {

	result := &RoleVsysWebuiPolicies{
		ApplicationOverrideRulebase: o.ApplicationOverrideRulebase,
		AuthenticationRulebase:      o.AuthenticationRulebase,
		DosRulebase:                 o.DosRulebase,
		NatRulebase:                 o.NatRulebase,
		NetworkPacketBrokerRulebase: o.NetworkPacketBrokerRulebase,
		PbfRulebase:                 o.PbfRulebase,
		QosRulebase:                 o.QosRulebase,
		RuleHitCountReset:           o.RuleHitCountReset,
		SdwanRulebase:               o.SdwanRulebase,
		SecurityRulebase:            o.SecurityRulebase,
		SslDecryptionRulebase:       o.SslDecryptionRulebase,
		TunnelInspectRulebase:       o.TunnelInspectRulebase,
		Misc:                        o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiPrivacyXml_11_0_2) MarshalFromObject(s RoleVsysWebuiPrivacy) {
	o.ShowFullIpAddresses = s.ShowFullIpAddresses
	o.ShowUserNamesInLogsAndReports = s.ShowUserNamesInLogsAndReports
	o.ViewPcapFiles = s.ViewPcapFiles
	o.Misc = s.Misc
}

func (o roleVsysWebuiPrivacyXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiPrivacy, error) {

	result := &RoleVsysWebuiPrivacy{
		ShowFullIpAddresses:           o.ShowFullIpAddresses,
		ShowUserNamesInLogsAndReports: o.ShowUserNamesInLogsAndReports,
		ViewPcapFiles:                 o.ViewPcapFiles,
		Misc:                          o.Misc,
	}
	return result, nil
}
func (o *roleVsysWebuiSaveXml_11_0_2) MarshalFromObject(s RoleVsysWebuiSave) {
	o.ObjectLevelChanges = s.ObjectLevelChanges
	o.PartialSave = s.PartialSave
	o.SaveForOtherAdmins = s.SaveForOtherAdmins
	o.Misc = s.Misc
}

func (o roleVsysWebuiSaveXml_11_0_2) UnmarshalToObject() (*RoleVsysWebuiSave, error) {

	result := &RoleVsysWebuiSave{
		ObjectLevelChanges: o.ObjectLevelChanges,
		PartialSave:        o.PartialSave,
		SaveForOtherAdmins: o.SaveForOtherAdmins,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *roleVsysXmlapiXml_11_0_2) MarshalFromObject(s RoleVsysXmlapi) {
	o.Commit = s.Commit
	o.Config = s.Config
	o.Export = s.Export
	o.Import = s.Import
	o.Iot = s.Iot
	o.Log = s.Log
	o.Op = s.Op
	o.Report = s.Report
	o.UserId = s.UserId
	o.Misc = s.Misc
}

func (o roleVsysXmlapiXml_11_0_2) UnmarshalToObject() (*RoleVsysXmlapi, error) {

	result := &RoleVsysXmlapi{
		Commit: o.Commit,
		Config: o.Config,
		Export: o.Export,
		Import: o.Import,
		Iot:    o.Iot,
		Log:    o.Log,
		Op:     o.Op,
		Report: o.Report,
		UserId: o.UserId,
		Misc:   o.Misc,
	}
	return result, nil
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
	version_11_0_2, _ := version.New("11.0.2")
	version_11_1_0, _ := version.New("11.1.0")
	if vn.Gte(version_11_0_2) && vn.Lt(version_11_1_0) {
		return specifyEntry_11_0_2, &entryXmlContainer_11_0_2{}, nil
	}

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
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !o.Role.matches(other.Role) {
		return false
	}

	return true
}

func (o *Role) matches(other *Role) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Device.matches(other.Device) {
		return false
	}
	if !o.Vsys.matches(other.Vsys) {
		return false
	}

	return true
}

func (o *RoleDevice) matches(other *RoleDevice) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Cli, other.Cli) {
		return false
	}
	if !o.Restapi.matches(other.Restapi) {
		return false
	}
	if !o.Webui.matches(other.Webui) {
		return false
	}
	if !o.Xmlapi.matches(other.Xmlapi) {
		return false
	}

	return true
}

func (o *RoleDeviceRestapi) matches(other *RoleDeviceRestapi) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Device.matches(other.Device) {
		return false
	}
	if !o.Network.matches(other.Network) {
		return false
	}
	if !o.Objects.matches(other.Objects) {
		return false
	}
	if !o.Policies.matches(other.Policies) {
		return false
	}
	if !o.System.matches(other.System) {
		return false
	}

	return true
}

func (o *RoleDeviceRestapiDevice) matches(other *RoleDeviceRestapiDevice) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.EmailServerProfiles, other.EmailServerProfiles) {
		return false
	}
	if !util.StringsMatch(o.HttpServerProfiles, other.HttpServerProfiles) {
		return false
	}
	if !util.StringsMatch(o.LdapServerProfiles, other.LdapServerProfiles) {
		return false
	}
	if !util.StringsMatch(o.LogInterfaceSetting, other.LogInterfaceSetting) {
		return false
	}
	if !util.StringsMatch(o.SnmpTrapServerProfiles, other.SnmpTrapServerProfiles) {
		return false
	}
	if !util.StringsMatch(o.SyslogServerProfiles, other.SyslogServerProfiles) {
		return false
	}
	if !util.StringsMatch(o.VirtualSystems, other.VirtualSystems) {
		return false
	}

	return true
}

func (o *RoleDeviceRestapiNetwork) matches(other *RoleDeviceRestapiNetwork) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AggregateEthernetInterfaces, other.AggregateEthernetInterfaces) {
		return false
	}
	if !util.StringsMatch(o.BfdNetworkProfiles, other.BfdNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(o.BgpRoutingProfiles, other.BgpRoutingProfiles) {
		return false
	}
	if !util.StringsMatch(o.DhcpRelays, other.DhcpRelays) {
		return false
	}
	if !util.StringsMatch(o.DhcpServers, other.DhcpServers) {
		return false
	}
	if !util.StringsMatch(o.DnsProxies, other.DnsProxies) {
		return false
	}
	if !util.StringsMatch(o.EthernetInterfaces, other.EthernetInterfaces) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectClientlessAppGroups, other.GlobalprotectClientlessAppGroups) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectClientlessApps, other.GlobalprotectClientlessApps) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectGateways, other.GlobalprotectGateways) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectIpsecCryptoNetworkProfiles, other.GlobalprotectIpsecCryptoNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectMdmServers, other.GlobalprotectMdmServers) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectPortals, other.GlobalprotectPortals) {
		return false
	}
	if !util.StringsMatch(o.GreTunnels, other.GreTunnels) {
		return false
	}
	if !util.StringsMatch(o.IkeCryptoNetworkProfiles, other.IkeCryptoNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(o.IkeGatewayNetworkProfiles, other.IkeGatewayNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(o.InterfaceManagementNetworkProfiles, other.InterfaceManagementNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(o.IpsecCryptoNetworkProfiles, other.IpsecCryptoNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(o.IpsecTunnels, other.IpsecTunnels) {
		return false
	}
	if !util.StringsMatch(o.Lldp, other.Lldp) {
		return false
	}
	if !util.StringsMatch(o.LldpNetworkProfiles, other.LldpNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(o.LogicalRouters, other.LogicalRouters) {
		return false
	}
	if !util.StringsMatch(o.LoopbackInterfaces, other.LoopbackInterfaces) {
		return false
	}
	if !util.StringsMatch(o.QosInterfaces, other.QosInterfaces) {
		return false
	}
	if !util.StringsMatch(o.QosNetworkProfiles, other.QosNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(o.SdwanInterfaceProfiles, other.SdwanInterfaceProfiles) {
		return false
	}
	if !util.StringsMatch(o.SdwanInterfaces, other.SdwanInterfaces) {
		return false
	}
	if !util.StringsMatch(o.TunnelInterfaces, other.TunnelInterfaces) {
		return false
	}
	if !util.StringsMatch(o.TunnelMonitorNetworkProfiles, other.TunnelMonitorNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(o.VirtualRouters, other.VirtualRouters) {
		return false
	}
	if !util.StringsMatch(o.VirtualWires, other.VirtualWires) {
		return false
	}
	if !util.StringsMatch(o.VlanInterfaces, other.VlanInterfaces) {
		return false
	}
	if !util.StringsMatch(o.Vlans, other.Vlans) {
		return false
	}
	if !util.StringsMatch(o.ZoneProtectionNetworkProfiles, other.ZoneProtectionNetworkProfiles) {
		return false
	}
	if !util.StringsMatch(o.Zones, other.Zones) {
		return false
	}

	return true
}

func (o *RoleDeviceRestapiObjects) matches(other *RoleDeviceRestapiObjects) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AddressGroups, other.AddressGroups) {
		return false
	}
	if !util.StringsMatch(o.Addresses, other.Addresses) {
		return false
	}
	if !util.StringsMatch(o.AntiSpywareSecurityProfiles, other.AntiSpywareSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.AntivirusSecurityProfiles, other.AntivirusSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.ApplicationFilters, other.ApplicationFilters) {
		return false
	}
	if !util.StringsMatch(o.ApplicationGroups, other.ApplicationGroups) {
		return false
	}
	if !util.StringsMatch(o.Applications, other.Applications) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationEnforcements, other.AuthenticationEnforcements) {
		return false
	}
	if !util.StringsMatch(o.CustomDataPatterns, other.CustomDataPatterns) {
		return false
	}
	if !util.StringsMatch(o.CustomSpywareSignatures, other.CustomSpywareSignatures) {
		return false
	}
	if !util.StringsMatch(o.CustomUrlCategories, other.CustomUrlCategories) {
		return false
	}
	if !util.StringsMatch(o.CustomVulnerabilitySignatures, other.CustomVulnerabilitySignatures) {
		return false
	}
	if !util.StringsMatch(o.DataFilteringSecurityProfiles, other.DataFilteringSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.DecryptionProfiles, other.DecryptionProfiles) {
		return false
	}
	if !util.StringsMatch(o.Devices, other.Devices) {
		return false
	}
	if !util.StringsMatch(o.DosProtectionSecurityProfiles, other.DosProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.DynamicUserGroups, other.DynamicUserGroups) {
		return false
	}
	if !util.StringsMatch(o.ExternalDynamicLists, other.ExternalDynamicLists) {
		return false
	}
	if !util.StringsMatch(o.FileBlockingSecurityProfiles, other.FileBlockingSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectHipObjects, other.GlobalprotectHipObjects) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectHipProfiles, other.GlobalprotectHipProfiles) {
		return false
	}
	if !util.StringsMatch(o.GtpProtectionSecurityProfiles, other.GtpProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.LogForwardingProfiles, other.LogForwardingProfiles) {
		return false
	}
	if !util.StringsMatch(o.PacketBrokerProfiles, other.PacketBrokerProfiles) {
		return false
	}
	if !util.StringsMatch(o.Regions, other.Regions) {
		return false
	}
	if !util.StringsMatch(o.Schedules, other.Schedules) {
		return false
	}
	if !util.StringsMatch(o.SctpProtectionSecurityProfiles, other.SctpProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.SdwanErrorCorrectionProfiles, other.SdwanErrorCorrectionProfiles) {
		return false
	}
	if !util.StringsMatch(o.SdwanPathQualityProfiles, other.SdwanPathQualityProfiles) {
		return false
	}
	if !util.StringsMatch(o.SdwanSaasQualityProfiles, other.SdwanSaasQualityProfiles) {
		return false
	}
	if !util.StringsMatch(o.SdwanTrafficDistributionProfiles, other.SdwanTrafficDistributionProfiles) {
		return false
	}
	if !util.StringsMatch(o.SecurityProfileGroups, other.SecurityProfileGroups) {
		return false
	}
	if !util.StringsMatch(o.ServiceGroups, other.ServiceGroups) {
		return false
	}
	if !util.StringsMatch(o.Services, other.Services) {
		return false
	}
	if !util.StringsMatch(o.Tags, other.Tags) {
		return false
	}
	if !util.StringsMatch(o.UrlFilteringSecurityProfiles, other.UrlFilteringSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.VulnerabilityProtectionSecurityProfiles, other.VulnerabilityProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.WildfireAnalysisSecurityProfiles, other.WildfireAnalysisSecurityProfiles) {
		return false
	}

	return true
}

func (o *RoleDeviceRestapiPolicies) matches(other *RoleDeviceRestapiPolicies) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ApplicationOverrideRules, other.ApplicationOverrideRules) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationRules, other.AuthenticationRules) {
		return false
	}
	if !util.StringsMatch(o.DecryptionRules, other.DecryptionRules) {
		return false
	}
	if !util.StringsMatch(o.DosRules, other.DosRules) {
		return false
	}
	if !util.StringsMatch(o.NatRules, other.NatRules) {
		return false
	}
	if !util.StringsMatch(o.NetworkPacketBrokerRules, other.NetworkPacketBrokerRules) {
		return false
	}
	if !util.StringsMatch(o.PolicyBasedForwardingRules, other.PolicyBasedForwardingRules) {
		return false
	}
	if !util.StringsMatch(o.QosRules, other.QosRules) {
		return false
	}
	if !util.StringsMatch(o.SdwanRules, other.SdwanRules) {
		return false
	}
	if !util.StringsMatch(o.SecurityRules, other.SecurityRules) {
		return false
	}
	if !util.StringsMatch(o.TunnelInspectionRules, other.TunnelInspectionRules) {
		return false
	}

	return true
}

func (o *RoleDeviceRestapiSystem) matches(other *RoleDeviceRestapiSystem) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Configuration, other.Configuration) {
		return false
	}

	return true
}

func (o *RoleDeviceWebui) matches(other *RoleDeviceWebui) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Acc, other.Acc) {
		return false
	}
	if !o.Commit.matches(other.Commit) {
		return false
	}
	if !util.StringsMatch(o.Dashboard, other.Dashboard) {
		return false
	}
	if !o.Device.matches(other.Device) {
		return false
	}
	if !o.Global.matches(other.Global) {
		return false
	}
	if !o.Monitor.matches(other.Monitor) {
		return false
	}
	if !o.Network.matches(other.Network) {
		return false
	}
	if !o.Objects.matches(other.Objects) {
		return false
	}
	if !o.Operations.matches(other.Operations) {
		return false
	}
	if !o.Policies.matches(other.Policies) {
		return false
	}
	if !o.Privacy.matches(other.Privacy) {
		return false
	}
	if !o.Save.matches(other.Save) {
		return false
	}
	if !util.StringsMatch(o.Tasks, other.Tasks) {
		return false
	}
	if !util.StringsMatch(o.Validate, other.Validate) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiCommit) matches(other *RoleDeviceWebuiCommit) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.CommitForOtherAdmins, other.CommitForOtherAdmins) {
		return false
	}
	if !util.StringsMatch(o.Device, other.Device) {
		return false
	}
	if !util.StringsMatch(o.ObjectLevelChanges, other.ObjectLevelChanges) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiDevice) matches(other *RoleDeviceWebuiDevice) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AccessDomain, other.AccessDomain) {
		return false
	}
	if !util.StringsMatch(o.AdminRoles, other.AdminRoles) {
		return false
	}
	if !util.StringsMatch(o.Administrators, other.Administrators) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationProfile, other.AuthenticationProfile) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationSequence, other.AuthenticationSequence) {
		return false
	}
	if !util.StringsMatch(o.BlockPages, other.BlockPages) {
		return false
	}
	if !o.CertificateManagement.matches(other.CertificateManagement) {
		return false
	}
	if !util.StringsMatch(o.ConfigAudit, other.ConfigAudit) {
		return false
	}
	if !util.StringsMatch(o.DataRedistribution, other.DataRedistribution) {
		return false
	}
	if !util.StringsMatch(o.DeviceQuarantine, other.DeviceQuarantine) {
		return false
	}
	if !util.StringsMatch(o.DhcpSyslogServer, other.DhcpSyslogServer) {
		return false
	}
	if !util.StringsMatch(o.DynamicUpdates, other.DynamicUpdates) {
		return false
	}
	if !util.StringsMatch(o.GlobalProtectClient, other.GlobalProtectClient) {
		return false
	}
	if !util.StringsMatch(o.HighAvailability, other.HighAvailability) {
		return false
	}
	if !util.StringsMatch(o.Licenses, other.Licenses) {
		return false
	}
	if !o.LocalUserDatabase.matches(other.LocalUserDatabase) {
		return false
	}
	if !util.StringsMatch(o.LogFwdCard, other.LogFwdCard) {
		return false
	}
	if !o.LogSettings.matches(other.LogSettings) {
		return false
	}
	if !util.StringsMatch(o.MasterKey, other.MasterKey) {
		return false
	}
	if !util.StringsMatch(o.Plugins, other.Plugins) {
		return false
	}
	if !o.PolicyRecommendations.matches(other.PolicyRecommendations) {
		return false
	}
	if !util.StringsMatch(o.ScheduledLogExport, other.ScheduledLogExport) {
		return false
	}
	if !o.ServerProfile.matches(other.ServerProfile) {
		return false
	}
	if !o.Setup.matches(other.Setup) {
		return false
	}
	if !util.StringsMatch(o.SharedGateways, other.SharedGateways) {
		return false
	}
	if !util.StringsMatch(o.Software, other.Software) {
		return false
	}
	if !util.StringsMatch(o.Support, other.Support) {
		return false
	}
	if !util.StringsMatch(o.Troubleshooting, other.Troubleshooting) {
		return false
	}
	if !util.StringsMatch(o.UserIdentification, other.UserIdentification) {
		return false
	}
	if !util.StringsMatch(o.VirtualSystems, other.VirtualSystems) {
		return false
	}
	if !util.StringsMatch(o.VmInfoSource, other.VmInfoSource) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiDeviceCertificateManagement) matches(other *RoleDeviceWebuiDeviceCertificateManagement) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.CertificateProfile, other.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(o.Certificates, other.Certificates) {
		return false
	}
	if !util.StringsMatch(o.OcspResponder, other.OcspResponder) {
		return false
	}
	if !util.StringsMatch(o.Scep, other.Scep) {
		return false
	}
	if !util.StringsMatch(o.SshServiceProfile, other.SshServiceProfile) {
		return false
	}
	if !util.StringsMatch(o.SslDecryptionExclusion, other.SslDecryptionExclusion) {
		return false
	}
	if !util.StringsMatch(o.SslTlsServiceProfile, other.SslTlsServiceProfile) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiDeviceLocalUserDatabase) matches(other *RoleDeviceWebuiDeviceLocalUserDatabase) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.UserGroups, other.UserGroups) {
		return false
	}
	if !util.StringsMatch(o.Users, other.Users) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiDeviceLogSettings) matches(other *RoleDeviceWebuiDeviceLogSettings) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.CcAlarm, other.CcAlarm) {
		return false
	}
	if !util.StringsMatch(o.Config, other.Config) {
		return false
	}
	if !util.StringsMatch(o.Correlation, other.Correlation) {
		return false
	}
	if !util.StringsMatch(o.Globalprotect, other.Globalprotect) {
		return false
	}
	if !util.StringsMatch(o.Hipmatch, other.Hipmatch) {
		return false
	}
	if !util.StringsMatch(o.Iptag, other.Iptag) {
		return false
	}
	if !util.StringsMatch(o.ManageLog, other.ManageLog) {
		return false
	}
	if !util.StringsMatch(o.System, other.System) {
		return false
	}
	if !util.StringsMatch(o.UserId, other.UserId) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiDevicePolicyRecommendations) matches(other *RoleDeviceWebuiDevicePolicyRecommendations) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Iot, other.Iot) {
		return false
	}
	if !util.StringsMatch(o.Saas, other.Saas) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiDeviceServerProfile) matches(other *RoleDeviceWebuiDeviceServerProfile) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Dns, other.Dns) {
		return false
	}
	if !util.StringsMatch(o.Email, other.Email) {
		return false
	}
	if !util.StringsMatch(o.Http, other.Http) {
		return false
	}
	if !util.StringsMatch(o.Kerberos, other.Kerberos) {
		return false
	}
	if !util.StringsMatch(o.Ldap, other.Ldap) {
		return false
	}
	if !util.StringsMatch(o.Mfa, other.Mfa) {
		return false
	}
	if !util.StringsMatch(o.Netflow, other.Netflow) {
		return false
	}
	if !util.StringsMatch(o.Radius, other.Radius) {
		return false
	}
	if !util.StringsMatch(o.SamlIdp, other.SamlIdp) {
		return false
	}
	if !util.StringsMatch(o.Scp, other.Scp) {
		return false
	}
	if !util.StringsMatch(o.SnmpTrap, other.SnmpTrap) {
		return false
	}
	if !util.StringsMatch(o.Syslog, other.Syslog) {
		return false
	}
	if !util.StringsMatch(o.Tacplus, other.Tacplus) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiDeviceSetup) matches(other *RoleDeviceWebuiDeviceSetup) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ContentId, other.ContentId) {
		return false
	}
	if !util.StringsMatch(o.Hsm, other.Hsm) {
		return false
	}
	if !util.StringsMatch(o.Interfaces, other.Interfaces) {
		return false
	}
	if !util.StringsMatch(o.Management, other.Management) {
		return false
	}
	if !util.StringsMatch(o.Operations, other.Operations) {
		return false
	}
	if !util.StringsMatch(o.Services, other.Services) {
		return false
	}
	if !util.StringsMatch(o.Session, other.Session) {
		return false
	}
	if !util.StringsMatch(o.Telemetry, other.Telemetry) {
		return false
	}
	if !util.StringsMatch(o.Wildfire, other.Wildfire) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiGlobal) matches(other *RoleDeviceWebuiGlobal) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.SystemAlarms, other.SystemAlarms) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiMonitor) matches(other *RoleDeviceWebuiMonitor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AppScope, other.AppScope) {
		return false
	}
	if !util.StringsMatch(o.ApplicationReports, other.ApplicationReports) {
		return false
	}
	if !o.AutomatedCorrelationEngine.matches(other.AutomatedCorrelationEngine) {
		return false
	}
	if !util.StringsMatch(o.BlockIpList, other.BlockIpList) {
		return false
	}
	if !util.StringsMatch(o.Botnet, other.Botnet) {
		return false
	}
	if !o.CustomReports.matches(other.CustomReports) {
		return false
	}
	if !util.StringsMatch(o.ExternalLogs, other.ExternalLogs) {
		return false
	}
	if !util.StringsMatch(o.GtpReports, other.GtpReports) {
		return false
	}
	if !o.Logs.matches(other.Logs) {
		return false
	}
	if !util.StringsMatch(o.PacketCapture, other.PacketCapture) {
		return false
	}
	if !o.PdfReports.matches(other.PdfReports) {
		return false
	}
	if !util.StringsMatch(o.SctpReports, other.SctpReports) {
		return false
	}
	if !util.StringsMatch(o.SessionBrowser, other.SessionBrowser) {
		return false
	}
	if !util.StringsMatch(o.ThreatReports, other.ThreatReports) {
		return false
	}
	if !util.StringsMatch(o.TrafficReports, other.TrafficReports) {
		return false
	}
	if !util.StringsMatch(o.UrlFilteringReports, other.UrlFilteringReports) {
		return false
	}
	if !util.StringsMatch(o.ViewCustomReports, other.ViewCustomReports) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiMonitorAutomatedCorrelationEngine) matches(other *RoleDeviceWebuiMonitorAutomatedCorrelationEngine) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.CorrelatedEvents, other.CorrelatedEvents) {
		return false
	}
	if !util.StringsMatch(o.CorrelationObjects, other.CorrelationObjects) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiMonitorCustomReports) matches(other *RoleDeviceWebuiMonitorCustomReports) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ApplicationStatistics, other.ApplicationStatistics) {
		return false
	}
	if !util.StringsMatch(o.Auth, other.Auth) {
		return false
	}
	if !util.StringsMatch(o.DataFilteringLog, other.DataFilteringLog) {
		return false
	}
	if !util.StringsMatch(o.DecryptionLog, other.DecryptionLog) {
		return false
	}
	if !util.StringsMatch(o.DecryptionSummary, other.DecryptionSummary) {
		return false
	}
	if !util.StringsMatch(o.Globalprotect, other.Globalprotect) {
		return false
	}
	if !util.StringsMatch(o.GtpLog, other.GtpLog) {
		return false
	}
	if !util.StringsMatch(o.GtpSummary, other.GtpSummary) {
		return false
	}
	if !util.StringsMatch(o.Hipmatch, other.Hipmatch) {
		return false
	}
	if !util.StringsMatch(o.Iptag, other.Iptag) {
		return false
	}
	if !util.StringsMatch(o.SctpLog, other.SctpLog) {
		return false
	}
	if !util.StringsMatch(o.SctpSummary, other.SctpSummary) {
		return false
	}
	if !util.StringsMatch(o.ThreatLog, other.ThreatLog) {
		return false
	}
	if !util.StringsMatch(o.ThreatSummary, other.ThreatSummary) {
		return false
	}
	if !util.StringsMatch(o.TrafficLog, other.TrafficLog) {
		return false
	}
	if !util.StringsMatch(o.TrafficSummary, other.TrafficSummary) {
		return false
	}
	if !util.StringsMatch(o.TunnelLog, other.TunnelLog) {
		return false
	}
	if !util.StringsMatch(o.TunnelSummary, other.TunnelSummary) {
		return false
	}
	if !util.StringsMatch(o.UrlLog, other.UrlLog) {
		return false
	}
	if !util.StringsMatch(o.UrlSummary, other.UrlSummary) {
		return false
	}
	if !util.StringsMatch(o.Userid, other.Userid) {
		return false
	}
	if !util.StringsMatch(o.WildfireLog, other.WildfireLog) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiMonitorLogs) matches(other *RoleDeviceWebuiMonitorLogs) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Alarm, other.Alarm) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !util.StringsMatch(o.Configuration, other.Configuration) {
		return false
	}
	if !util.StringsMatch(o.DataFiltering, other.DataFiltering) {
		return false
	}
	if !util.StringsMatch(o.Decryption, other.Decryption) {
		return false
	}
	if !util.StringsMatch(o.Globalprotect, other.Globalprotect) {
		return false
	}
	if !util.StringsMatch(o.Gtp, other.Gtp) {
		return false
	}
	if !util.StringsMatch(o.Hipmatch, other.Hipmatch) {
		return false
	}
	if !util.StringsMatch(o.Iptag, other.Iptag) {
		return false
	}
	if !util.StringsMatch(o.Sctp, other.Sctp) {
		return false
	}
	if !util.StringsMatch(o.System, other.System) {
		return false
	}
	if !util.StringsMatch(o.Threat, other.Threat) {
		return false
	}
	if !util.StringsMatch(o.Traffic, other.Traffic) {
		return false
	}
	if !util.StringsMatch(o.Tunnel, other.Tunnel) {
		return false
	}
	if !util.StringsMatch(o.Url, other.Url) {
		return false
	}
	if !util.StringsMatch(o.Userid, other.Userid) {
		return false
	}
	if !util.StringsMatch(o.Wildfire, other.Wildfire) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiMonitorPdfReports) matches(other *RoleDeviceWebuiMonitorPdfReports) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.EmailScheduler, other.EmailScheduler) {
		return false
	}
	if !util.StringsMatch(o.ManagePdfSummary, other.ManagePdfSummary) {
		return false
	}
	if !util.StringsMatch(o.PdfSummaryReports, other.PdfSummaryReports) {
		return false
	}
	if !util.StringsMatch(o.ReportGroups, other.ReportGroups) {
		return false
	}
	if !util.StringsMatch(o.SaasApplicationUsageReport, other.SaasApplicationUsageReport) {
		return false
	}
	if !util.StringsMatch(o.UserActivityReport, other.UserActivityReport) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiNetwork) matches(other *RoleDeviceWebuiNetwork) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Dhcp, other.Dhcp) {
		return false
	}
	if !util.StringsMatch(o.DnsProxy, other.DnsProxy) {
		return false
	}
	if !o.GlobalProtect.matches(other.GlobalProtect) {
		return false
	}
	if !util.StringsMatch(o.GreTunnels, other.GreTunnels) {
		return false
	}
	if !util.StringsMatch(o.Interfaces, other.Interfaces) {
		return false
	}
	if !util.StringsMatch(o.IpsecTunnels, other.IpsecTunnels) {
		return false
	}
	if !util.StringsMatch(o.Lldp, other.Lldp) {
		return false
	}
	if !o.NetworkProfiles.matches(other.NetworkProfiles) {
		return false
	}
	if !util.StringsMatch(o.Qos, other.Qos) {
		return false
	}
	if !o.Routing.matches(other.Routing) {
		return false
	}
	if !util.StringsMatch(o.SdwanInterfaceProfile, other.SdwanInterfaceProfile) {
		return false
	}
	if !util.StringsMatch(o.SecureWebGateway, other.SecureWebGateway) {
		return false
	}
	if !util.StringsMatch(o.VirtualRouters, other.VirtualRouters) {
		return false
	}
	if !util.StringsMatch(o.VirtualWires, other.VirtualWires) {
		return false
	}
	if !util.StringsMatch(o.Vlans, other.Vlans) {
		return false
	}
	if !util.StringsMatch(o.Zones, other.Zones) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiNetworkGlobalProtect) matches(other *RoleDeviceWebuiNetworkGlobalProtect) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ClientlessAppGroups, other.ClientlessAppGroups) {
		return false
	}
	if !util.StringsMatch(o.ClientlessApps, other.ClientlessApps) {
		return false
	}
	if !util.StringsMatch(o.Gateways, other.Gateways) {
		return false
	}
	if !util.StringsMatch(o.Mdm, other.Mdm) {
		return false
	}
	if !util.StringsMatch(o.Portals, other.Portals) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiNetworkNetworkProfiles) matches(other *RoleDeviceWebuiNetworkNetworkProfiles) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.BfdProfile, other.BfdProfile) {
		return false
	}
	if !util.StringsMatch(o.GpAppIpsecCrypto, other.GpAppIpsecCrypto) {
		return false
	}
	if !util.StringsMatch(o.IkeCrypto, other.IkeCrypto) {
		return false
	}
	if !util.StringsMatch(o.IkeGateways, other.IkeGateways) {
		return false
	}
	if !util.StringsMatch(o.InterfaceMgmt, other.InterfaceMgmt) {
		return false
	}
	if !util.StringsMatch(o.IpsecCrypto, other.IpsecCrypto) {
		return false
	}
	if !util.StringsMatch(o.LldpProfile, other.LldpProfile) {
		return false
	}
	if !util.StringsMatch(o.QosProfile, other.QosProfile) {
		return false
	}
	if !util.StringsMatch(o.TunnelMonitor, other.TunnelMonitor) {
		return false
	}
	if !util.StringsMatch(o.ZoneProtection, other.ZoneProtection) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiNetworkRouting) matches(other *RoleDeviceWebuiNetworkRouting) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.LogicalRouters, other.LogicalRouters) {
		return false
	}
	if !o.RoutingProfiles.matches(other.RoutingProfiles) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiNetworkRoutingRoutingProfiles) matches(other *RoleDeviceWebuiNetworkRoutingRoutingProfiles) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Bfd, other.Bfd) {
		return false
	}
	if !util.StringsMatch(o.Bgp, other.Bgp) {
		return false
	}
	if !util.StringsMatch(o.Filters, other.Filters) {
		return false
	}
	if !util.StringsMatch(o.Multicast, other.Multicast) {
		return false
	}
	if !util.StringsMatch(o.Ospf, other.Ospf) {
		return false
	}
	if !util.StringsMatch(o.Ospfv3, other.Ospfv3) {
		return false
	}
	if !util.StringsMatch(o.Ripv2, other.Ripv2) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiObjects) matches(other *RoleDeviceWebuiObjects) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AddressGroups, other.AddressGroups) {
		return false
	}
	if !util.StringsMatch(o.Addresses, other.Addresses) {
		return false
	}
	if !util.StringsMatch(o.ApplicationFilters, other.ApplicationFilters) {
		return false
	}
	if !util.StringsMatch(o.ApplicationGroups, other.ApplicationGroups) {
		return false
	}
	if !util.StringsMatch(o.Applications, other.Applications) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !o.CustomObjects.matches(other.CustomObjects) {
		return false
	}
	if !o.Decryption.matches(other.Decryption) {
		return false
	}
	if !util.StringsMatch(o.Devices, other.Devices) {
		return false
	}
	if !util.StringsMatch(o.DynamicBlockLists, other.DynamicBlockLists) {
		return false
	}
	if !util.StringsMatch(o.DynamicUserGroups, other.DynamicUserGroups) {
		return false
	}
	if !o.GlobalProtect.matches(other.GlobalProtect) {
		return false
	}
	if !util.StringsMatch(o.LogForwarding, other.LogForwarding) {
		return false
	}
	if !util.StringsMatch(o.PacketBrokerProfile, other.PacketBrokerProfile) {
		return false
	}
	if !util.StringsMatch(o.Regions, other.Regions) {
		return false
	}
	if !util.StringsMatch(o.Schedules, other.Schedules) {
		return false
	}
	if !o.Sdwan.matches(other.Sdwan) {
		return false
	}
	if !util.StringsMatch(o.SecurityProfileGroups, other.SecurityProfileGroups) {
		return false
	}
	if !o.SecurityProfiles.matches(other.SecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.ServiceGroups, other.ServiceGroups) {
		return false
	}
	if !util.StringsMatch(o.Services, other.Services) {
		return false
	}
	if !util.StringsMatch(o.Tags, other.Tags) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiObjectsCustomObjects) matches(other *RoleDeviceWebuiObjectsCustomObjects) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DataPatterns, other.DataPatterns) {
		return false
	}
	if !util.StringsMatch(o.Spyware, other.Spyware) {
		return false
	}
	if !util.StringsMatch(o.UrlCategory, other.UrlCategory) {
		return false
	}
	if !util.StringsMatch(o.Vulnerability, other.Vulnerability) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiObjectsDecryption) matches(other *RoleDeviceWebuiObjectsDecryption) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DecryptionProfile, other.DecryptionProfile) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiObjectsGlobalProtect) matches(other *RoleDeviceWebuiObjectsGlobalProtect) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.HipObjects, other.HipObjects) {
		return false
	}
	if !util.StringsMatch(o.HipProfiles, other.HipProfiles) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiObjectsSdwan) matches(other *RoleDeviceWebuiObjectsSdwan) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.SdwanDistProfile, other.SdwanDistProfile) {
		return false
	}
	if !util.StringsMatch(o.SdwanErrorCorrectionProfile, other.SdwanErrorCorrectionProfile) {
		return false
	}
	if !util.StringsMatch(o.SdwanProfile, other.SdwanProfile) {
		return false
	}
	if !util.StringsMatch(o.SdwanSaasQualityProfile, other.SdwanSaasQualityProfile) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiObjectsSecurityProfiles) matches(other *RoleDeviceWebuiObjectsSecurityProfiles) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AntiSpyware, other.AntiSpyware) {
		return false
	}
	if !util.StringsMatch(o.Antivirus, other.Antivirus) {
		return false
	}
	if !util.StringsMatch(o.DataFiltering, other.DataFiltering) {
		return false
	}
	if !util.StringsMatch(o.DosProtection, other.DosProtection) {
		return false
	}
	if !util.StringsMatch(o.FileBlocking, other.FileBlocking) {
		return false
	}
	if !util.StringsMatch(o.GtpProtection, other.GtpProtection) {
		return false
	}
	if !util.StringsMatch(o.SctpProtection, other.SctpProtection) {
		return false
	}
	if !util.StringsMatch(o.UrlFiltering, other.UrlFiltering) {
		return false
	}
	if !util.StringsMatch(o.VulnerabilityProtection, other.VulnerabilityProtection) {
		return false
	}
	if !util.StringsMatch(o.WildfireAnalysis, other.WildfireAnalysis) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiOperations) matches(other *RoleDeviceWebuiOperations) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DownloadCoreFiles, other.DownloadCoreFiles) {
		return false
	}
	if !util.StringsMatch(o.DownloadPcapFiles, other.DownloadPcapFiles) {
		return false
	}
	if !util.StringsMatch(o.GenerateStatsDumpFile, other.GenerateStatsDumpFile) {
		return false
	}
	if !util.StringsMatch(o.GenerateTechSupportFile, other.GenerateTechSupportFile) {
		return false
	}
	if !util.StringsMatch(o.Reboot, other.Reboot) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiPolicies) matches(other *RoleDeviceWebuiPolicies) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ApplicationOverrideRulebase, other.ApplicationOverrideRulebase) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationRulebase, other.AuthenticationRulebase) {
		return false
	}
	if !util.StringsMatch(o.DosRulebase, other.DosRulebase) {
		return false
	}
	if !util.StringsMatch(o.NatRulebase, other.NatRulebase) {
		return false
	}
	if !util.StringsMatch(o.NetworkPacketBrokerRulebase, other.NetworkPacketBrokerRulebase) {
		return false
	}
	if !util.StringsMatch(o.PbfRulebase, other.PbfRulebase) {
		return false
	}
	if !util.StringsMatch(o.QosRulebase, other.QosRulebase) {
		return false
	}
	if !util.StringsMatch(o.RuleHitCountReset, other.RuleHitCountReset) {
		return false
	}
	if !util.StringsMatch(o.SdwanRulebase, other.SdwanRulebase) {
		return false
	}
	if !util.StringsMatch(o.SecurityRulebase, other.SecurityRulebase) {
		return false
	}
	if !util.StringsMatch(o.SslDecryptionRulebase, other.SslDecryptionRulebase) {
		return false
	}
	if !util.StringsMatch(o.TunnelInspectRulebase, other.TunnelInspectRulebase) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiPrivacy) matches(other *RoleDeviceWebuiPrivacy) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ShowFullIpAddresses, other.ShowFullIpAddresses) {
		return false
	}
	if !util.StringsMatch(o.ShowUserNamesInLogsAndReports, other.ShowUserNamesInLogsAndReports) {
		return false
	}
	if !util.StringsMatch(o.ViewPcapFiles, other.ViewPcapFiles) {
		return false
	}

	return true
}

func (o *RoleDeviceWebuiSave) matches(other *RoleDeviceWebuiSave) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ObjectLevelChanges, other.ObjectLevelChanges) {
		return false
	}
	if !util.StringsMatch(o.PartialSave, other.PartialSave) {
		return false
	}
	if !util.StringsMatch(o.SaveForOtherAdmins, other.SaveForOtherAdmins) {
		return false
	}

	return true
}

func (o *RoleDeviceXmlapi) matches(other *RoleDeviceXmlapi) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Commit, other.Commit) {
		return false
	}
	if !util.StringsMatch(o.Config, other.Config) {
		return false
	}
	if !util.StringsMatch(o.Export, other.Export) {
		return false
	}
	if !util.StringsMatch(o.Import, other.Import) {
		return false
	}
	if !util.StringsMatch(o.Iot, other.Iot) {
		return false
	}
	if !util.StringsMatch(o.Log, other.Log) {
		return false
	}
	if !util.StringsMatch(o.Op, other.Op) {
		return false
	}
	if !util.StringsMatch(o.Report, other.Report) {
		return false
	}
	if !util.StringsMatch(o.UserId, other.UserId) {
		return false
	}

	return true
}

func (o *RoleVsys) matches(other *RoleVsys) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Cli, other.Cli) {
		return false
	}
	if !o.Restapi.matches(other.Restapi) {
		return false
	}
	if !o.Webui.matches(other.Webui) {
		return false
	}
	if !o.Xmlapi.matches(other.Xmlapi) {
		return false
	}

	return true
}

func (o *RoleVsysRestapi) matches(other *RoleVsysRestapi) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Device.matches(other.Device) {
		return false
	}
	if !o.Network.matches(other.Network) {
		return false
	}
	if !o.Objects.matches(other.Objects) {
		return false
	}
	if !o.Policies.matches(other.Policies) {
		return false
	}
	if !o.System.matches(other.System) {
		return false
	}

	return true
}

func (o *RoleVsysRestapiDevice) matches(other *RoleVsysRestapiDevice) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.EmailServerProfiles, other.EmailServerProfiles) {
		return false
	}
	if !util.StringsMatch(o.HttpServerProfiles, other.HttpServerProfiles) {
		return false
	}
	if !util.StringsMatch(o.LdapServerProfiles, other.LdapServerProfiles) {
		return false
	}
	if !util.StringsMatch(o.LogInterfaceSetting, other.LogInterfaceSetting) {
		return false
	}
	if !util.StringsMatch(o.SnmpTrapServerProfiles, other.SnmpTrapServerProfiles) {
		return false
	}
	if !util.StringsMatch(o.SyslogServerProfiles, other.SyslogServerProfiles) {
		return false
	}
	if !util.StringsMatch(o.VirtualSystems, other.VirtualSystems) {
		return false
	}

	return true
}

func (o *RoleVsysRestapiNetwork) matches(other *RoleVsysRestapiNetwork) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectClientlessAppGroups, other.GlobalprotectClientlessAppGroups) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectClientlessApps, other.GlobalprotectClientlessApps) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectGateways, other.GlobalprotectGateways) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectMdmServers, other.GlobalprotectMdmServers) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectPortals, other.GlobalprotectPortals) {
		return false
	}
	if !util.StringsMatch(o.SdwanInterfaceProfiles, other.SdwanInterfaceProfiles) {
		return false
	}
	if !util.StringsMatch(o.Zones, other.Zones) {
		return false
	}

	return true
}

func (o *RoleVsysRestapiObjects) matches(other *RoleVsysRestapiObjects) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AddressGroups, other.AddressGroups) {
		return false
	}
	if !util.StringsMatch(o.Addresses, other.Addresses) {
		return false
	}
	if !util.StringsMatch(o.AntiSpywareSecurityProfiles, other.AntiSpywareSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.AntivirusSecurityProfiles, other.AntivirusSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.ApplicationFilters, other.ApplicationFilters) {
		return false
	}
	if !util.StringsMatch(o.ApplicationGroups, other.ApplicationGroups) {
		return false
	}
	if !util.StringsMatch(o.Applications, other.Applications) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationEnforcements, other.AuthenticationEnforcements) {
		return false
	}
	if !util.StringsMatch(o.CustomDataPatterns, other.CustomDataPatterns) {
		return false
	}
	if !util.StringsMatch(o.CustomSpywareSignatures, other.CustomSpywareSignatures) {
		return false
	}
	if !util.StringsMatch(o.CustomUrlCategories, other.CustomUrlCategories) {
		return false
	}
	if !util.StringsMatch(o.CustomVulnerabilitySignatures, other.CustomVulnerabilitySignatures) {
		return false
	}
	if !util.StringsMatch(o.DataFilteringSecurityProfiles, other.DataFilteringSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.DecryptionProfiles, other.DecryptionProfiles) {
		return false
	}
	if !util.StringsMatch(o.Devices, other.Devices) {
		return false
	}
	if !util.StringsMatch(o.DosProtectionSecurityProfiles, other.DosProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.DynamicUserGroups, other.DynamicUserGroups) {
		return false
	}
	if !util.StringsMatch(o.ExternalDynamicLists, other.ExternalDynamicLists) {
		return false
	}
	if !util.StringsMatch(o.FileBlockingSecurityProfiles, other.FileBlockingSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectHipObjects, other.GlobalprotectHipObjects) {
		return false
	}
	if !util.StringsMatch(o.GlobalprotectHipProfiles, other.GlobalprotectHipProfiles) {
		return false
	}
	if !util.StringsMatch(o.GtpProtectionSecurityProfiles, other.GtpProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.LogForwardingProfiles, other.LogForwardingProfiles) {
		return false
	}
	if !util.StringsMatch(o.PacketBrokerProfiles, other.PacketBrokerProfiles) {
		return false
	}
	if !util.StringsMatch(o.Regions, other.Regions) {
		return false
	}
	if !util.StringsMatch(o.Schedules, other.Schedules) {
		return false
	}
	if !util.StringsMatch(o.SctpProtectionSecurityProfiles, other.SctpProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.SdwanErrorCorrectionProfiles, other.SdwanErrorCorrectionProfiles) {
		return false
	}
	if !util.StringsMatch(o.SdwanPathQualityProfiles, other.SdwanPathQualityProfiles) {
		return false
	}
	if !util.StringsMatch(o.SdwanSaasQualityProfiles, other.SdwanSaasQualityProfiles) {
		return false
	}
	if !util.StringsMatch(o.SdwanTrafficDistributionProfiles, other.SdwanTrafficDistributionProfiles) {
		return false
	}
	if !util.StringsMatch(o.SecurityProfileGroups, other.SecurityProfileGroups) {
		return false
	}
	if !util.StringsMatch(o.ServiceGroups, other.ServiceGroups) {
		return false
	}
	if !util.StringsMatch(o.Services, other.Services) {
		return false
	}
	if !util.StringsMatch(o.Tags, other.Tags) {
		return false
	}
	if !util.StringsMatch(o.UrlFilteringSecurityProfiles, other.UrlFilteringSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.VulnerabilityProtectionSecurityProfiles, other.VulnerabilityProtectionSecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.WildfireAnalysisSecurityProfiles, other.WildfireAnalysisSecurityProfiles) {
		return false
	}

	return true
}

func (o *RoleVsysRestapiPolicies) matches(other *RoleVsysRestapiPolicies) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ApplicationOverrideRules, other.ApplicationOverrideRules) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationRules, other.AuthenticationRules) {
		return false
	}
	if !util.StringsMatch(o.DecryptionRules, other.DecryptionRules) {
		return false
	}
	if !util.StringsMatch(o.DosRules, other.DosRules) {
		return false
	}
	if !util.StringsMatch(o.NatRules, other.NatRules) {
		return false
	}
	if !util.StringsMatch(o.NetworkPacketBrokerRules, other.NetworkPacketBrokerRules) {
		return false
	}
	if !util.StringsMatch(o.PolicyBasedForwardingRules, other.PolicyBasedForwardingRules) {
		return false
	}
	if !util.StringsMatch(o.QosRules, other.QosRules) {
		return false
	}
	if !util.StringsMatch(o.SdwanRules, other.SdwanRules) {
		return false
	}
	if !util.StringsMatch(o.SecurityRules, other.SecurityRules) {
		return false
	}
	if !util.StringsMatch(o.TunnelInspectionRules, other.TunnelInspectionRules) {
		return false
	}

	return true
}

func (o *RoleVsysRestapiSystem) matches(other *RoleVsysRestapiSystem) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Configuration, other.Configuration) {
		return false
	}

	return true
}

func (o *RoleVsysWebui) matches(other *RoleVsysWebui) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Acc, other.Acc) {
		return false
	}
	if !o.Commit.matches(other.Commit) {
		return false
	}
	if !util.StringsMatch(o.Dashboard, other.Dashboard) {
		return false
	}
	if !o.Device.matches(other.Device) {
		return false
	}
	if !o.Monitor.matches(other.Monitor) {
		return false
	}
	if !o.Network.matches(other.Network) {
		return false
	}
	if !o.Objects.matches(other.Objects) {
		return false
	}
	if !o.Operations.matches(other.Operations) {
		return false
	}
	if !o.Policies.matches(other.Policies) {
		return false
	}
	if !o.Privacy.matches(other.Privacy) {
		return false
	}
	if !o.Save.matches(other.Save) {
		return false
	}
	if !util.StringsMatch(o.Tasks, other.Tasks) {
		return false
	}
	if !util.StringsMatch(o.Validate, other.Validate) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiCommit) matches(other *RoleVsysWebuiCommit) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.CommitForOtherAdmins, other.CommitForOtherAdmins) {
		return false
	}
	if !util.StringsMatch(o.VirtualSystems, other.VirtualSystems) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiDevice) matches(other *RoleVsysWebuiDevice) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Administrators, other.Administrators) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationProfile, other.AuthenticationProfile) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationSequence, other.AuthenticationSequence) {
		return false
	}
	if !util.StringsMatch(o.BlockPages, other.BlockPages) {
		return false
	}
	if !o.CertificateManagement.matches(other.CertificateManagement) {
		return false
	}
	if !util.StringsMatch(o.DataRedistribution, other.DataRedistribution) {
		return false
	}
	if !util.StringsMatch(o.DeviceQuarantine, other.DeviceQuarantine) {
		return false
	}
	if !util.StringsMatch(o.DhcpSyslogServer, other.DhcpSyslogServer) {
		return false
	}
	if !o.LocalUserDatabase.matches(other.LocalUserDatabase) {
		return false
	}
	if !o.LogSettings.matches(other.LogSettings) {
		return false
	}
	if !o.PolicyRecommendations.matches(other.PolicyRecommendations) {
		return false
	}
	if !o.ServerProfile.matches(other.ServerProfile) {
		return false
	}
	if !o.Setup.matches(other.Setup) {
		return false
	}
	if !util.StringsMatch(o.Troubleshooting, other.Troubleshooting) {
		return false
	}
	if !util.StringsMatch(o.UserIdentification, other.UserIdentification) {
		return false
	}
	if !util.StringsMatch(o.VmInfoSource, other.VmInfoSource) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiDeviceCertificateManagement) matches(other *RoleVsysWebuiDeviceCertificateManagement) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.CertificateProfile, other.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(o.Certificates, other.Certificates) {
		return false
	}
	if !util.StringsMatch(o.OcspResponder, other.OcspResponder) {
		return false
	}
	if !util.StringsMatch(o.Scep, other.Scep) {
		return false
	}
	if !util.StringsMatch(o.SshServiceProfile, other.SshServiceProfile) {
		return false
	}
	if !util.StringsMatch(o.SslDecryptionExclusion, other.SslDecryptionExclusion) {
		return false
	}
	if !util.StringsMatch(o.SslTlsServiceProfile, other.SslTlsServiceProfile) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiDeviceLocalUserDatabase) matches(other *RoleVsysWebuiDeviceLocalUserDatabase) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.UserGroups, other.UserGroups) {
		return false
	}
	if !util.StringsMatch(o.Users, other.Users) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiDeviceLogSettings) matches(other *RoleVsysWebuiDeviceLogSettings) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Config, other.Config) {
		return false
	}
	if !util.StringsMatch(o.Correlation, other.Correlation) {
		return false
	}
	if !util.StringsMatch(o.Globalprotect, other.Globalprotect) {
		return false
	}
	if !util.StringsMatch(o.Hipmatch, other.Hipmatch) {
		return false
	}
	if !util.StringsMatch(o.Iptag, other.Iptag) {
		return false
	}
	if !util.StringsMatch(o.System, other.System) {
		return false
	}
	if !util.StringsMatch(o.UserId, other.UserId) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiDevicePolicyRecommendations) matches(other *RoleVsysWebuiDevicePolicyRecommendations) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Iot, other.Iot) {
		return false
	}
	if !util.StringsMatch(o.Saas, other.Saas) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiDeviceServerProfile) matches(other *RoleVsysWebuiDeviceServerProfile) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Dns, other.Dns) {
		return false
	}
	if !util.StringsMatch(o.Email, other.Email) {
		return false
	}
	if !util.StringsMatch(o.Http, other.Http) {
		return false
	}
	if !util.StringsMatch(o.Kerberos, other.Kerberos) {
		return false
	}
	if !util.StringsMatch(o.Ldap, other.Ldap) {
		return false
	}
	if !util.StringsMatch(o.Mfa, other.Mfa) {
		return false
	}
	if !util.StringsMatch(o.Netflow, other.Netflow) {
		return false
	}
	if !util.StringsMatch(o.Radius, other.Radius) {
		return false
	}
	if !util.StringsMatch(o.SamlIdp, other.SamlIdp) {
		return false
	}
	if !util.StringsMatch(o.Scp, other.Scp) {
		return false
	}
	if !util.StringsMatch(o.SnmpTrap, other.SnmpTrap) {
		return false
	}
	if !util.StringsMatch(o.Syslog, other.Syslog) {
		return false
	}
	if !util.StringsMatch(o.Tacplus, other.Tacplus) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiDeviceSetup) matches(other *RoleVsysWebuiDeviceSetup) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ContentId, other.ContentId) {
		return false
	}
	if !util.StringsMatch(o.Hsm, other.Hsm) {
		return false
	}
	if !util.StringsMatch(o.Interfaces, other.Interfaces) {
		return false
	}
	if !util.StringsMatch(o.Management, other.Management) {
		return false
	}
	if !util.StringsMatch(o.Operations, other.Operations) {
		return false
	}
	if !util.StringsMatch(o.Services, other.Services) {
		return false
	}
	if !util.StringsMatch(o.Session, other.Session) {
		return false
	}
	if !util.StringsMatch(o.Telemetry, other.Telemetry) {
		return false
	}
	if !util.StringsMatch(o.Wildfire, other.Wildfire) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiMonitor) matches(other *RoleVsysWebuiMonitor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AppScope, other.AppScope) {
		return false
	}
	if !o.AutomatedCorrelationEngine.matches(other.AutomatedCorrelationEngine) {
		return false
	}
	if !util.StringsMatch(o.BlockIpList, other.BlockIpList) {
		return false
	}
	if !o.CustomReports.matches(other.CustomReports) {
		return false
	}
	if !util.StringsMatch(o.ExternalLogs, other.ExternalLogs) {
		return false
	}
	if !o.Logs.matches(other.Logs) {
		return false
	}
	if !o.PdfReports.matches(other.PdfReports) {
		return false
	}
	if !util.StringsMatch(o.SessionBrowser, other.SessionBrowser) {
		return false
	}
	if !util.StringsMatch(o.ViewCustomReports, other.ViewCustomReports) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiMonitorAutomatedCorrelationEngine) matches(other *RoleVsysWebuiMonitorAutomatedCorrelationEngine) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.CorrelatedEvents, other.CorrelatedEvents) {
		return false
	}
	if !util.StringsMatch(o.CorrelationObjects, other.CorrelationObjects) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiMonitorCustomReports) matches(other *RoleVsysWebuiMonitorCustomReports) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ApplicationStatistics, other.ApplicationStatistics) {
		return false
	}
	if !util.StringsMatch(o.Auth, other.Auth) {
		return false
	}
	if !util.StringsMatch(o.DataFilteringLog, other.DataFilteringLog) {
		return false
	}
	if !util.StringsMatch(o.DecryptionLog, other.DecryptionLog) {
		return false
	}
	if !util.StringsMatch(o.DecryptionSummary, other.DecryptionSummary) {
		return false
	}
	if !util.StringsMatch(o.Globalprotect, other.Globalprotect) {
		return false
	}
	if !util.StringsMatch(o.GtpLog, other.GtpLog) {
		return false
	}
	if !util.StringsMatch(o.GtpSummary, other.GtpSummary) {
		return false
	}
	if !util.StringsMatch(o.Hipmatch, other.Hipmatch) {
		return false
	}
	if !util.StringsMatch(o.Iptag, other.Iptag) {
		return false
	}
	if !util.StringsMatch(o.SctpLog, other.SctpLog) {
		return false
	}
	if !util.StringsMatch(o.SctpSummary, other.SctpSummary) {
		return false
	}
	if !util.StringsMatch(o.ThreatLog, other.ThreatLog) {
		return false
	}
	if !util.StringsMatch(o.ThreatSummary, other.ThreatSummary) {
		return false
	}
	if !util.StringsMatch(o.TrafficLog, other.TrafficLog) {
		return false
	}
	if !util.StringsMatch(o.TrafficSummary, other.TrafficSummary) {
		return false
	}
	if !util.StringsMatch(o.TunnelLog, other.TunnelLog) {
		return false
	}
	if !util.StringsMatch(o.TunnelSummary, other.TunnelSummary) {
		return false
	}
	if !util.StringsMatch(o.UrlLog, other.UrlLog) {
		return false
	}
	if !util.StringsMatch(o.UrlSummary, other.UrlSummary) {
		return false
	}
	if !util.StringsMatch(o.Userid, other.Userid) {
		return false
	}
	if !util.StringsMatch(o.WildfireLog, other.WildfireLog) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiMonitorLogs) matches(other *RoleVsysWebuiMonitorLogs) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !util.StringsMatch(o.DataFiltering, other.DataFiltering) {
		return false
	}
	if !util.StringsMatch(o.Decryption, other.Decryption) {
		return false
	}
	if !util.StringsMatch(o.Globalprotect, other.Globalprotect) {
		return false
	}
	if !util.StringsMatch(o.Gtp, other.Gtp) {
		return false
	}
	if !util.StringsMatch(o.Hipmatch, other.Hipmatch) {
		return false
	}
	if !util.StringsMatch(o.Iptag, other.Iptag) {
		return false
	}
	if !util.StringsMatch(o.Sctp, other.Sctp) {
		return false
	}
	if !util.StringsMatch(o.Threat, other.Threat) {
		return false
	}
	if !util.StringsMatch(o.Traffic, other.Traffic) {
		return false
	}
	if !util.StringsMatch(o.Tunnel, other.Tunnel) {
		return false
	}
	if !util.StringsMatch(o.Url, other.Url) {
		return false
	}
	if !util.StringsMatch(o.Userid, other.Userid) {
		return false
	}
	if !util.StringsMatch(o.Wildfire, other.Wildfire) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiMonitorPdfReports) matches(other *RoleVsysWebuiMonitorPdfReports) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.EmailScheduler, other.EmailScheduler) {
		return false
	}
	if !util.StringsMatch(o.ManagePdfSummary, other.ManagePdfSummary) {
		return false
	}
	if !util.StringsMatch(o.PdfSummaryReports, other.PdfSummaryReports) {
		return false
	}
	if !util.StringsMatch(o.ReportGroups, other.ReportGroups) {
		return false
	}
	if !util.StringsMatch(o.SaasApplicationUsageReport, other.SaasApplicationUsageReport) {
		return false
	}
	if !util.StringsMatch(o.UserActivityReport, other.UserActivityReport) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiNetwork) matches(other *RoleVsysWebuiNetwork) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.GlobalProtect.matches(other.GlobalProtect) {
		return false
	}
	if !util.StringsMatch(o.SdwanInterfaceProfile, other.SdwanInterfaceProfile) {
		return false
	}
	if !util.StringsMatch(o.Zones, other.Zones) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiNetworkGlobalProtect) matches(other *RoleVsysWebuiNetworkGlobalProtect) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ClientlessAppGroups, other.ClientlessAppGroups) {
		return false
	}
	if !util.StringsMatch(o.ClientlessApps, other.ClientlessApps) {
		return false
	}
	if !util.StringsMatch(o.Gateways, other.Gateways) {
		return false
	}
	if !util.StringsMatch(o.Mdm, other.Mdm) {
		return false
	}
	if !util.StringsMatch(o.Portals, other.Portals) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiObjects) matches(other *RoleVsysWebuiObjects) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AddressGroups, other.AddressGroups) {
		return false
	}
	if !util.StringsMatch(o.Addresses, other.Addresses) {
		return false
	}
	if !util.StringsMatch(o.ApplicationFilters, other.ApplicationFilters) {
		return false
	}
	if !util.StringsMatch(o.ApplicationGroups, other.ApplicationGroups) {
		return false
	}
	if !util.StringsMatch(o.Applications, other.Applications) {
		return false
	}
	if !util.StringsMatch(o.Authentication, other.Authentication) {
		return false
	}
	if !o.CustomObjects.matches(other.CustomObjects) {
		return false
	}
	if !o.Decryption.matches(other.Decryption) {
		return false
	}
	if !util.StringsMatch(o.Devices, other.Devices) {
		return false
	}
	if !util.StringsMatch(o.DynamicBlockLists, other.DynamicBlockLists) {
		return false
	}
	if !util.StringsMatch(o.DynamicUserGroups, other.DynamicUserGroups) {
		return false
	}
	if !o.GlobalProtect.matches(other.GlobalProtect) {
		return false
	}
	if !util.StringsMatch(o.LogForwarding, other.LogForwarding) {
		return false
	}
	if !util.StringsMatch(o.PacketBrokerProfile, other.PacketBrokerProfile) {
		return false
	}
	if !util.StringsMatch(o.Regions, other.Regions) {
		return false
	}
	if !util.StringsMatch(o.Schedules, other.Schedules) {
		return false
	}
	if !o.Sdwan.matches(other.Sdwan) {
		return false
	}
	if !util.StringsMatch(o.SecurityProfileGroups, other.SecurityProfileGroups) {
		return false
	}
	if !o.SecurityProfiles.matches(other.SecurityProfiles) {
		return false
	}
	if !util.StringsMatch(o.ServiceGroups, other.ServiceGroups) {
		return false
	}
	if !util.StringsMatch(o.Services, other.Services) {
		return false
	}
	if !util.StringsMatch(o.Tags, other.Tags) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiObjectsCustomObjects) matches(other *RoleVsysWebuiObjectsCustomObjects) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DataPatterns, other.DataPatterns) {
		return false
	}
	if !util.StringsMatch(o.Spyware, other.Spyware) {
		return false
	}
	if !util.StringsMatch(o.UrlCategory, other.UrlCategory) {
		return false
	}
	if !util.StringsMatch(o.Vulnerability, other.Vulnerability) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiObjectsDecryption) matches(other *RoleVsysWebuiObjectsDecryption) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DecryptionProfile, other.DecryptionProfile) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiObjectsGlobalProtect) matches(other *RoleVsysWebuiObjectsGlobalProtect) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.HipObjects, other.HipObjects) {
		return false
	}
	if !util.StringsMatch(o.HipProfiles, other.HipProfiles) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiObjectsSdwan) matches(other *RoleVsysWebuiObjectsSdwan) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.SdwanDistProfile, other.SdwanDistProfile) {
		return false
	}
	if !util.StringsMatch(o.SdwanErrorCorrectionProfile, other.SdwanErrorCorrectionProfile) {
		return false
	}
	if !util.StringsMatch(o.SdwanProfile, other.SdwanProfile) {
		return false
	}
	if !util.StringsMatch(o.SdwanSaasQualityProfile, other.SdwanSaasQualityProfile) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiObjectsSecurityProfiles) matches(other *RoleVsysWebuiObjectsSecurityProfiles) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.AntiSpyware, other.AntiSpyware) {
		return false
	}
	if !util.StringsMatch(o.Antivirus, other.Antivirus) {
		return false
	}
	if !util.StringsMatch(o.DataFiltering, other.DataFiltering) {
		return false
	}
	if !util.StringsMatch(o.DosProtection, other.DosProtection) {
		return false
	}
	if !util.StringsMatch(o.FileBlocking, other.FileBlocking) {
		return false
	}
	if !util.StringsMatch(o.GtpProtection, other.GtpProtection) {
		return false
	}
	if !util.StringsMatch(o.SctpProtection, other.SctpProtection) {
		return false
	}
	if !util.StringsMatch(o.UrlFiltering, other.UrlFiltering) {
		return false
	}
	if !util.StringsMatch(o.VulnerabilityProtection, other.VulnerabilityProtection) {
		return false
	}
	if !util.StringsMatch(o.WildfireAnalysis, other.WildfireAnalysis) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiOperations) matches(other *RoleVsysWebuiOperations) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DownloadCoreFiles, other.DownloadCoreFiles) {
		return false
	}
	if !util.StringsMatch(o.DownloadPcapFiles, other.DownloadPcapFiles) {
		return false
	}
	if !util.StringsMatch(o.GenerateStatsDumpFile, other.GenerateStatsDumpFile) {
		return false
	}
	if !util.StringsMatch(o.GenerateTechSupportFile, other.GenerateTechSupportFile) {
		return false
	}
	if !util.StringsMatch(o.Reboot, other.Reboot) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiPolicies) matches(other *RoleVsysWebuiPolicies) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ApplicationOverrideRulebase, other.ApplicationOverrideRulebase) {
		return false
	}
	if !util.StringsMatch(o.AuthenticationRulebase, other.AuthenticationRulebase) {
		return false
	}
	if !util.StringsMatch(o.DosRulebase, other.DosRulebase) {
		return false
	}
	if !util.StringsMatch(o.NatRulebase, other.NatRulebase) {
		return false
	}
	if !util.StringsMatch(o.NetworkPacketBrokerRulebase, other.NetworkPacketBrokerRulebase) {
		return false
	}
	if !util.StringsMatch(o.PbfRulebase, other.PbfRulebase) {
		return false
	}
	if !util.StringsMatch(o.QosRulebase, other.QosRulebase) {
		return false
	}
	if !util.StringsMatch(o.RuleHitCountReset, other.RuleHitCountReset) {
		return false
	}
	if !util.StringsMatch(o.SdwanRulebase, other.SdwanRulebase) {
		return false
	}
	if !util.StringsMatch(o.SecurityRulebase, other.SecurityRulebase) {
		return false
	}
	if !util.StringsMatch(o.SslDecryptionRulebase, other.SslDecryptionRulebase) {
		return false
	}
	if !util.StringsMatch(o.TunnelInspectRulebase, other.TunnelInspectRulebase) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiPrivacy) matches(other *RoleVsysWebuiPrivacy) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ShowFullIpAddresses, other.ShowFullIpAddresses) {
		return false
	}
	if !util.StringsMatch(o.ShowUserNamesInLogsAndReports, other.ShowUserNamesInLogsAndReports) {
		return false
	}
	if !util.StringsMatch(o.ViewPcapFiles, other.ViewPcapFiles) {
		return false
	}

	return true
}

func (o *RoleVsysWebuiSave) matches(other *RoleVsysWebuiSave) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.ObjectLevelChanges, other.ObjectLevelChanges) {
		return false
	}
	if !util.StringsMatch(o.PartialSave, other.PartialSave) {
		return false
	}
	if !util.StringsMatch(o.SaveForOtherAdmins, other.SaveForOtherAdmins) {
		return false
	}

	return true
}

func (o *RoleVsysXmlapi) matches(other *RoleVsysXmlapi) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Commit, other.Commit) {
		return false
	}
	if !util.StringsMatch(o.Config, other.Config) {
		return false
	}
	if !util.StringsMatch(o.Export, other.Export) {
		return false
	}
	if !util.StringsMatch(o.Import, other.Import) {
		return false
	}
	if !util.StringsMatch(o.Iot, other.Iot) {
		return false
	}
	if !util.StringsMatch(o.Log, other.Log) {
		return false
	}
	if !util.StringsMatch(o.Op, other.Op) {
		return false
	}
	if !util.StringsMatch(o.Report, other.Report) {
		return false
	}
	if !util.StringsMatch(o.UserId, other.UserId) {
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
