package spyware

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
	Suffix = []string{"profiles", "spyware"}
)

type Entry struct {
	Name                     string
	BotnetDomains            *BotnetDomains
	CloudInlineAnalysis      *bool
	Description              *string
	DisableOverride          *string
	InlineExceptionEdlUrl    []string
	InlineExceptionIpAddress []string
	MicaEngineSpywareEnabled []MicaEngineSpywareEnabled
	Rules                    []Rules
	ThreatException          []ThreatException

	Misc map[string][]generic.Xml
}

type BotnetDomains struct {
	DnsSecurityCategories []BotnetDomainsDnsSecurityCategories
	Lists                 []BotnetDomainsLists
	RtypeAction           *BotnetDomainsRtypeAction
	Sinkhole              *BotnetDomainsSinkhole
	ThreatException       []BotnetDomainsThreatException
	Whitelist             []BotnetDomainsWhitelist
}
type BotnetDomainsDnsSecurityCategories struct {
	Action        *string
	LogLevel      *string
	Name          string
	PacketCapture *string
}
type BotnetDomainsLists struct {
	Action        *BotnetDomainsListsAction
	Name          string
	PacketCapture *string
}
type BotnetDomainsListsAction struct {
	Alert    *BotnetDomainsListsActionAlert
	Allow    *BotnetDomainsListsActionAllow
	Block    *BotnetDomainsListsActionBlock
	Sinkhole *BotnetDomainsListsActionSinkhole
}
type BotnetDomainsListsActionAlert struct {
}
type BotnetDomainsListsActionAllow struct {
}
type BotnetDomainsListsActionBlock struct {
}
type BotnetDomainsListsActionSinkhole struct {
}
type BotnetDomainsRtypeAction struct {
	Any   *string
	Https *string
	Svcb  *string
}
type BotnetDomainsSinkhole struct {
	Ipv4Address *string
	Ipv6Address *string
}
type BotnetDomainsThreatException struct {
	Name string
}
type BotnetDomainsWhitelist struct {
	Description *string
	Name        string
}
type MicaEngineSpywareEnabled struct {
	InlinePolicyAction *string
	Name               string
}
type Rules struct {
	Action        *RulesAction
	Category      *string
	Name          string
	PacketCapture *string
	Severity      []string
	ThreatName    *string
}
type RulesAction struct {
	Alert       *RulesActionAlert
	Allow       *RulesActionAllow
	BlockIp     *RulesActionBlockIp
	Default     *RulesActionDefault
	Drop        *RulesActionDrop
	ResetBoth   *RulesActionResetBoth
	ResetClient *RulesActionResetClient
	ResetServer *RulesActionResetServer
}
type RulesActionAlert struct {
}
type RulesActionAllow struct {
}
type RulesActionBlockIp struct {
	Duration *int64
	TrackBy  *string
}
type RulesActionDefault struct {
}
type RulesActionDrop struct {
}
type RulesActionResetBoth struct {
}
type RulesActionResetClient struct {
}
type RulesActionResetServer struct {
}
type ThreatException struct {
	Action        *ThreatExceptionAction
	ExemptIp      []ThreatExceptionExemptIp
	Name          string
	PacketCapture *string
}
type ThreatExceptionAction struct {
	Alert       *ThreatExceptionActionAlert
	Allow       *ThreatExceptionActionAllow
	BlockIp     *ThreatExceptionActionBlockIp
	Default     *ThreatExceptionActionDefault
	Drop        *ThreatExceptionActionDrop
	ResetBoth   *ThreatExceptionActionResetBoth
	ResetClient *ThreatExceptionActionResetClient
	ResetServer *ThreatExceptionActionResetServer
}
type ThreatExceptionActionAlert struct {
}
type ThreatExceptionActionAllow struct {
}
type ThreatExceptionActionBlockIp struct {
	Duration *int64
	TrackBy  *string
}
type ThreatExceptionActionDefault struct {
}
type ThreatExceptionActionDrop struct {
}
type ThreatExceptionActionResetBoth struct {
}
type ThreatExceptionActionResetClient struct {
}
type ThreatExceptionActionResetServer struct {
}
type ThreatExceptionExemptIp struct {
	Name string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXmlContainer_11_0_2 struct {
	Answer []entryXml_11_0_2 `xml:"entry"`
}
type entryXml struct {
	XMLName                  xml.Name                      `xml:"entry"`
	Name                     string                        `xml:"name,attr"`
	BotnetDomains            *BotnetDomainsXml             `xml:"botnet-domains,omitempty"`
	CloudInlineAnalysis      *string                       `xml:"cloud-inline-analysis,omitempty"`
	Description              *string                       `xml:"description,omitempty"`
	DisableOverride          *string                       `xml:"disable-override,omitempty"`
	InlineExceptionEdlUrl    *util.MemberType              `xml:"inline-exception-edl-url,omitempty"`
	InlineExceptionIpAddress *util.MemberType              `xml:"inline-exception-ip-address,omitempty"`
	MicaEngineSpywareEnabled []MicaEngineSpywareEnabledXml `xml:"mica-engine-spyware-enabled>entry,omitempty"`
	Rules                    []RulesXml                    `xml:"rules>entry,omitempty"`
	ThreatException          []ThreatExceptionXml          `xml:"threat-exception>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type entryXml_11_0_2 struct {
	XMLName                  xml.Name                             `xml:"entry"`
	Name                     string                               `xml:"name,attr"`
	BotnetDomains            *BotnetDomainsXml_11_0_2             `xml:"botnet-domains,omitempty"`
	CloudInlineAnalysis      *string                              `xml:"cloud-inline-analysis,omitempty"`
	Description              *string                              `xml:"description,omitempty"`
	DisableOverride          *string                              `xml:"disable-override,omitempty"`
	InlineExceptionEdlUrl    *util.MemberType                     `xml:"inline-exception-edl-url,omitempty"`
	InlineExceptionIpAddress *util.MemberType                     `xml:"inline-exception-ip-address,omitempty"`
	MicaEngineSpywareEnabled []MicaEngineSpywareEnabledXml_11_0_2 `xml:"mica-engine-spyware-enabled>entry,omitempty"`
	Rules                    []RulesXml_11_0_2                    `xml:"rules>entry,omitempty"`
	ThreatException          []ThreatExceptionXml_11_0_2          `xml:"threat-exception>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsXml struct {
	DnsSecurityCategories []BotnetDomainsDnsSecurityCategoriesXml `xml:"dns-security-categories>entry,omitempty"`
	Lists                 []BotnetDomainsListsXml                 `xml:"lists>entry,omitempty"`
	RtypeAction           *BotnetDomainsRtypeActionXml            `xml:"rtype-action,omitempty"`
	Sinkhole              *BotnetDomainsSinkholeXml               `xml:"sinkhole,omitempty"`
	ThreatException       []BotnetDomainsThreatExceptionXml       `xml:"threat-exception>entry,omitempty"`
	Whitelist             []BotnetDomainsWhitelistXml             `xml:"whitelist>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsDnsSecurityCategoriesXml struct {
	Action        *string  `xml:"action,omitempty"`
	LogLevel      *string  `xml:"log-level,omitempty"`
	XMLName       xml.Name `xml:"entry"`
	Name          string   `xml:"name,attr"`
	PacketCapture *string  `xml:"packet-capture,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsListsXml struct {
	Action        *BotnetDomainsListsActionXml `xml:"action,omitempty"`
	XMLName       xml.Name                     `xml:"entry"`
	Name          string                       `xml:"name,attr"`
	PacketCapture *string                      `xml:"packet-capture,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsListsActionXml struct {
	Alert    *BotnetDomainsListsActionAlertXml    `xml:"alert,omitempty"`
	Allow    *BotnetDomainsListsActionAllowXml    `xml:"allow,omitempty"`
	Block    *BotnetDomainsListsActionBlockXml    `xml:"block,omitempty"`
	Sinkhole *BotnetDomainsListsActionSinkholeXml `xml:"sinkhole,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsListsActionAlertXml struct {
	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsListsActionAllowXml struct {
	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsListsActionBlockXml struct {
	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsListsActionSinkholeXml struct {
	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsRtypeActionXml struct {
	Any   *string `xml:"any,omitempty"`
	Https *string `xml:"https,omitempty"`
	Svcb  *string `xml:"svcb,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsSinkholeXml struct {
	Ipv4Address *string `xml:"ipv4-address,omitempty"`
	Ipv6Address *string `xml:"ipv6-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsThreatExceptionXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsWhitelistXml struct {
	Description *string  `xml:"description,omitempty"`
	XMLName     xml.Name `xml:"entry"`
	Name        string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type MicaEngineSpywareEnabledXml struct {
	InlinePolicyAction *string  `xml:"inline-policy-action,omitempty"`
	XMLName            xml.Name `xml:"entry"`
	Name               string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type RulesXml struct {
	Action        *RulesActionXml  `xml:"action,omitempty"`
	Category      *string          `xml:"category,omitempty"`
	XMLName       xml.Name         `xml:"entry"`
	Name          string           `xml:"name,attr"`
	PacketCapture *string          `xml:"packet-capture,omitempty"`
	Severity      *util.MemberType `xml:"severity,omitempty"`
	ThreatName    *string          `xml:"threat-name,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RulesActionXml struct {
	Alert       *RulesActionAlertXml       `xml:"alert,omitempty"`
	Allow       *RulesActionAllowXml       `xml:"allow,omitempty"`
	BlockIp     *RulesActionBlockIpXml     `xml:"block-ip,omitempty"`
	Default     *RulesActionDefaultXml     `xml:"default,omitempty"`
	Drop        *RulesActionDropXml        `xml:"drop,omitempty"`
	ResetBoth   *RulesActionResetBothXml   `xml:"reset-both,omitempty"`
	ResetClient *RulesActionResetClientXml `xml:"reset-client,omitempty"`
	ResetServer *RulesActionResetServerXml `xml:"reset-server,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RulesActionAlertXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RulesActionAllowXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RulesActionBlockIpXml struct {
	Duration *int64  `xml:"duration,omitempty"`
	TrackBy  *string `xml:"track-by,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RulesActionDefaultXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RulesActionDropXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RulesActionResetBothXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RulesActionResetClientXml struct {
	Misc []generic.Xml `xml:",any"`
}
type RulesActionResetServerXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionXml struct {
	Action        *ThreatExceptionActionXml    `xml:"action,omitempty"`
	ExemptIp      []ThreatExceptionExemptIpXml `xml:"exempt-ip>entry,omitempty"`
	XMLName       xml.Name                     `xml:"entry"`
	Name          string                       `xml:"name,attr"`
	PacketCapture *string                      `xml:"packet-capture,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionXml struct {
	Alert       *ThreatExceptionActionAlertXml       `xml:"alert,omitempty"`
	Allow       *ThreatExceptionActionAllowXml       `xml:"allow,omitempty"`
	BlockIp     *ThreatExceptionActionBlockIpXml     `xml:"block-ip,omitempty"`
	Default     *ThreatExceptionActionDefaultXml     `xml:"default,omitempty"`
	Drop        *ThreatExceptionActionDropXml        `xml:"drop,omitempty"`
	ResetBoth   *ThreatExceptionActionResetBothXml   `xml:"reset-both,omitempty"`
	ResetClient *ThreatExceptionActionResetClientXml `xml:"reset-client,omitempty"`
	ResetServer *ThreatExceptionActionResetServerXml `xml:"reset-server,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionAlertXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionAllowXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionBlockIpXml struct {
	Duration *int64  `xml:"duration,omitempty"`
	TrackBy  *string `xml:"track-by,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionDefaultXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionDropXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionResetBothXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionResetClientXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionResetServerXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionExemptIpXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsXml_11_0_2 struct {
	DnsSecurityCategories []BotnetDomainsDnsSecurityCategoriesXml_11_0_2 `xml:"dns-security-categories>entry,omitempty"`
	Lists                 []BotnetDomainsListsXml_11_0_2                 `xml:"lists>entry,omitempty"`
	RtypeAction           *BotnetDomainsRtypeActionXml_11_0_2            `xml:"rtype-action,omitempty"`
	Sinkhole              *BotnetDomainsSinkholeXml_11_0_2               `xml:"sinkhole,omitempty"`
	ThreatException       []BotnetDomainsThreatExceptionXml_11_0_2       `xml:"threat-exception>entry,omitempty"`
	Whitelist             []BotnetDomainsWhitelistXml_11_0_2             `xml:"whitelist>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsDnsSecurityCategoriesXml_11_0_2 struct {
	Action        *string  `xml:"action,omitempty"`
	LogLevel      *string  `xml:"log-level,omitempty"`
	XMLName       xml.Name `xml:"entry"`
	Name          string   `xml:"name,attr"`
	PacketCapture *string  `xml:"packet-capture,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsListsXml_11_0_2 struct {
	Action        *BotnetDomainsListsActionXml_11_0_2 `xml:"action,omitempty"`
	XMLName       xml.Name                            `xml:"entry"`
	Name          string                              `xml:"name,attr"`
	PacketCapture *string                             `xml:"packet-capture,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsListsActionXml_11_0_2 struct {
	Alert    *BotnetDomainsListsActionAlertXml_11_0_2    `xml:"alert,omitempty"`
	Allow    *BotnetDomainsListsActionAllowXml_11_0_2    `xml:"allow,omitempty"`
	Block    *BotnetDomainsListsActionBlockXml_11_0_2    `xml:"block,omitempty"`
	Sinkhole *BotnetDomainsListsActionSinkholeXml_11_0_2 `xml:"sinkhole,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsListsActionAlertXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsListsActionAllowXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsListsActionBlockXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsListsActionSinkholeXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsRtypeActionXml_11_0_2 struct {
	Any   *string `xml:"any,omitempty"`
	Https *string `xml:"https,omitempty"`
	Svcb  *string `xml:"svcb,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsSinkholeXml_11_0_2 struct {
	Ipv4Address *string `xml:"ipv4-address,omitempty"`
	Ipv6Address *string `xml:"ipv6-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsThreatExceptionXml_11_0_2 struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type BotnetDomainsWhitelistXml_11_0_2 struct {
	Description *string  `xml:"description,omitempty"`
	XMLName     xml.Name `xml:"entry"`
	Name        string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type MicaEngineSpywareEnabledXml_11_0_2 struct {
	InlinePolicyAction *string  `xml:"inline-policy-action,omitempty"`
	XMLName            xml.Name `xml:"entry"`
	Name               string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type RulesXml_11_0_2 struct {
	Action        *RulesActionXml_11_0_2 `xml:"action,omitempty"`
	Category      *string                `xml:"category,omitempty"`
	XMLName       xml.Name               `xml:"entry"`
	Name          string                 `xml:"name,attr"`
	PacketCapture *string                `xml:"packet-capture,omitempty"`
	Severity      *util.MemberType       `xml:"severity,omitempty"`
	ThreatName    *string                `xml:"threat-name,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RulesActionXml_11_0_2 struct {
	Alert       *RulesActionAlertXml_11_0_2       `xml:"alert,omitempty"`
	Allow       *RulesActionAllowXml_11_0_2       `xml:"allow,omitempty"`
	BlockIp     *RulesActionBlockIpXml_11_0_2     `xml:"block-ip,omitempty"`
	Default     *RulesActionDefaultXml_11_0_2     `xml:"default,omitempty"`
	Drop        *RulesActionDropXml_11_0_2        `xml:"drop,omitempty"`
	ResetBoth   *RulesActionResetBothXml_11_0_2   `xml:"reset-both,omitempty"`
	ResetClient *RulesActionResetClientXml_11_0_2 `xml:"reset-client,omitempty"`
	ResetServer *RulesActionResetServerXml_11_0_2 `xml:"reset-server,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RulesActionAlertXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type RulesActionAllowXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type RulesActionBlockIpXml_11_0_2 struct {
	Duration *int64  `xml:"duration,omitempty"`
	TrackBy  *string `xml:"track-by,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RulesActionDefaultXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type RulesActionDropXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type RulesActionResetBothXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type RulesActionResetClientXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type RulesActionResetServerXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionXml_11_0_2 struct {
	Action        *ThreatExceptionActionXml_11_0_2    `xml:"action,omitempty"`
	ExemptIp      []ThreatExceptionExemptIpXml_11_0_2 `xml:"exempt-ip>entry,omitempty"`
	XMLName       xml.Name                            `xml:"entry"`
	Name          string                              `xml:"name,attr"`
	PacketCapture *string                             `xml:"packet-capture,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionXml_11_0_2 struct {
	Alert       *ThreatExceptionActionAlertXml_11_0_2       `xml:"alert,omitempty"`
	Allow       *ThreatExceptionActionAllowXml_11_0_2       `xml:"allow,omitempty"`
	BlockIp     *ThreatExceptionActionBlockIpXml_11_0_2     `xml:"block-ip,omitempty"`
	Default     *ThreatExceptionActionDefaultXml_11_0_2     `xml:"default,omitempty"`
	Drop        *ThreatExceptionActionDropXml_11_0_2        `xml:"drop,omitempty"`
	ResetBoth   *ThreatExceptionActionResetBothXml_11_0_2   `xml:"reset-both,omitempty"`
	ResetClient *ThreatExceptionActionResetClientXml_11_0_2 `xml:"reset-client,omitempty"`
	ResetServer *ThreatExceptionActionResetServerXml_11_0_2 `xml:"reset-server,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionAlertXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionAllowXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionBlockIpXml_11_0_2 struct {
	Duration *int64  `xml:"duration,omitempty"`
	TrackBy  *string `xml:"track-by,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionDefaultXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionDropXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionResetBothXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionResetClientXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionActionResetServerXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type ThreatExceptionExemptIpXml_11_0_2 struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "botnet_domains" || v == "BotnetDomains" {
		return e.BotnetDomains, nil
	}
	if v == "cloud_inline_analysis" || v == "CloudInlineAnalysis" {
		return e.CloudInlineAnalysis, nil
	}
	if v == "description" || v == "Description" {
		return e.Description, nil
	}
	if v == "disable_override" || v == "DisableOverride" {
		return e.DisableOverride, nil
	}
	if v == "inline_exception_edl_url" || v == "InlineExceptionEdlUrl" {
		return e.InlineExceptionEdlUrl, nil
	}
	if v == "inline_exception_edl_url|LENGTH" || v == "InlineExceptionEdlUrl|LENGTH" {
		return int64(len(e.InlineExceptionEdlUrl)), nil
	}
	if v == "inline_exception_ip_address" || v == "InlineExceptionIpAddress" {
		return e.InlineExceptionIpAddress, nil
	}
	if v == "inline_exception_ip_address|LENGTH" || v == "InlineExceptionIpAddress|LENGTH" {
		return int64(len(e.InlineExceptionIpAddress)), nil
	}
	if v == "mica_engine_spyware_enabled" || v == "MicaEngineSpywareEnabled" {
		return e.MicaEngineSpywareEnabled, nil
	}
	if v == "mica_engine_spyware_enabled|LENGTH" || v == "MicaEngineSpywareEnabled|LENGTH" {
		return int64(len(e.MicaEngineSpywareEnabled)), nil
	}
	if v == "rules" || v == "Rules" {
		return e.Rules, nil
	}
	if v == "rules|LENGTH" || v == "Rules|LENGTH" {
		return int64(len(e.Rules)), nil
	}
	if v == "threat_exception" || v == "ThreatException" {
		return e.ThreatException, nil
	}
	if v == "threat_exception|LENGTH" || v == "ThreatException|LENGTH" {
		return int64(len(e.ThreatException)), nil
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
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	var nestedBotnetDomains *BotnetDomainsXml
	if o.BotnetDomains != nil {
		nestedBotnetDomains = &BotnetDomainsXml{}
		if _, ok := o.Misc["BotnetDomains"]; ok {
			nestedBotnetDomains.Misc = o.Misc["BotnetDomains"]
		}
		if o.BotnetDomains.Sinkhole != nil {
			nestedBotnetDomains.Sinkhole = &BotnetDomainsSinkholeXml{}
			if _, ok := o.Misc["BotnetDomainsSinkhole"]; ok {
				nestedBotnetDomains.Sinkhole.Misc = o.Misc["BotnetDomainsSinkhole"]
			}
			if o.BotnetDomains.Sinkhole.Ipv4Address != nil {
				nestedBotnetDomains.Sinkhole.Ipv4Address = o.BotnetDomains.Sinkhole.Ipv4Address
			}
			if o.BotnetDomains.Sinkhole.Ipv6Address != nil {
				nestedBotnetDomains.Sinkhole.Ipv6Address = o.BotnetDomains.Sinkhole.Ipv6Address
			}
		}
		if o.BotnetDomains.ThreatException != nil {
			nestedBotnetDomains.ThreatException = []BotnetDomainsThreatExceptionXml{}
			for _, oBotnetDomainsThreatException := range o.BotnetDomains.ThreatException {
				nestedBotnetDomainsThreatException := BotnetDomainsThreatExceptionXml{}
				if _, ok := o.Misc["BotnetDomainsThreatException"]; ok {
					nestedBotnetDomainsThreatException.Misc = o.Misc["BotnetDomainsThreatException"]
				}
				if oBotnetDomainsThreatException.Name != "" {
					nestedBotnetDomainsThreatException.Name = oBotnetDomainsThreatException.Name
				}
				nestedBotnetDomains.ThreatException = append(nestedBotnetDomains.ThreatException, nestedBotnetDomainsThreatException)
			}
		}
		if o.BotnetDomains.Whitelist != nil {
			nestedBotnetDomains.Whitelist = []BotnetDomainsWhitelistXml{}
			for _, oBotnetDomainsWhitelist := range o.BotnetDomains.Whitelist {
				nestedBotnetDomainsWhitelist := BotnetDomainsWhitelistXml{}
				if _, ok := o.Misc["BotnetDomainsWhitelist"]; ok {
					nestedBotnetDomainsWhitelist.Misc = o.Misc["BotnetDomainsWhitelist"]
				}
				if oBotnetDomainsWhitelist.Description != nil {
					nestedBotnetDomainsWhitelist.Description = oBotnetDomainsWhitelist.Description
				}
				if oBotnetDomainsWhitelist.Name != "" {
					nestedBotnetDomainsWhitelist.Name = oBotnetDomainsWhitelist.Name
				}
				nestedBotnetDomains.Whitelist = append(nestedBotnetDomains.Whitelist, nestedBotnetDomainsWhitelist)
			}
		}
		if o.BotnetDomains.RtypeAction != nil {
			nestedBotnetDomains.RtypeAction = &BotnetDomainsRtypeActionXml{}
			if _, ok := o.Misc["BotnetDomainsRtypeAction"]; ok {
				nestedBotnetDomains.RtypeAction.Misc = o.Misc["BotnetDomainsRtypeAction"]
			}
			if o.BotnetDomains.RtypeAction.Any != nil {
				nestedBotnetDomains.RtypeAction.Any = o.BotnetDomains.RtypeAction.Any
			}
			if o.BotnetDomains.RtypeAction.Https != nil {
				nestedBotnetDomains.RtypeAction.Https = o.BotnetDomains.RtypeAction.Https
			}
			if o.BotnetDomains.RtypeAction.Svcb != nil {
				nestedBotnetDomains.RtypeAction.Svcb = o.BotnetDomains.RtypeAction.Svcb
			}
		}
		if o.BotnetDomains.DnsSecurityCategories != nil {
			nestedBotnetDomains.DnsSecurityCategories = []BotnetDomainsDnsSecurityCategoriesXml{}
			for _, oBotnetDomainsDnsSecurityCategories := range o.BotnetDomains.DnsSecurityCategories {
				nestedBotnetDomainsDnsSecurityCategories := BotnetDomainsDnsSecurityCategoriesXml{}
				if _, ok := o.Misc["BotnetDomainsDnsSecurityCategories"]; ok {
					nestedBotnetDomainsDnsSecurityCategories.Misc = o.Misc["BotnetDomainsDnsSecurityCategories"]
				}
				if oBotnetDomainsDnsSecurityCategories.Action != nil {
					nestedBotnetDomainsDnsSecurityCategories.Action = oBotnetDomainsDnsSecurityCategories.Action
				}
				if oBotnetDomainsDnsSecurityCategories.LogLevel != nil {
					nestedBotnetDomainsDnsSecurityCategories.LogLevel = oBotnetDomainsDnsSecurityCategories.LogLevel
				}
				if oBotnetDomainsDnsSecurityCategories.PacketCapture != nil {
					nestedBotnetDomainsDnsSecurityCategories.PacketCapture = oBotnetDomainsDnsSecurityCategories.PacketCapture
				}
				if oBotnetDomainsDnsSecurityCategories.Name != "" {
					nestedBotnetDomainsDnsSecurityCategories.Name = oBotnetDomainsDnsSecurityCategories.Name
				}
				nestedBotnetDomains.DnsSecurityCategories = append(nestedBotnetDomains.DnsSecurityCategories, nestedBotnetDomainsDnsSecurityCategories)
			}
		}
		if o.BotnetDomains.Lists != nil {
			nestedBotnetDomains.Lists = []BotnetDomainsListsXml{}
			for _, oBotnetDomainsLists := range o.BotnetDomains.Lists {
				nestedBotnetDomainsLists := BotnetDomainsListsXml{}
				if _, ok := o.Misc["BotnetDomainsLists"]; ok {
					nestedBotnetDomainsLists.Misc = o.Misc["BotnetDomainsLists"]
				}
				if oBotnetDomainsLists.Action != nil {
					nestedBotnetDomainsLists.Action = &BotnetDomainsListsActionXml{}
					if _, ok := o.Misc["BotnetDomainsListsAction"]; ok {
						nestedBotnetDomainsLists.Action.Misc = o.Misc["BotnetDomainsListsAction"]
					}
					if oBotnetDomainsLists.Action.Alert != nil {
						nestedBotnetDomainsLists.Action.Alert = &BotnetDomainsListsActionAlertXml{}
						if _, ok := o.Misc["BotnetDomainsListsActionAlert"]; ok {
							nestedBotnetDomainsLists.Action.Alert.Misc = o.Misc["BotnetDomainsListsActionAlert"]
						}
					}
					if oBotnetDomainsLists.Action.Allow != nil {
						nestedBotnetDomainsLists.Action.Allow = &BotnetDomainsListsActionAllowXml{}
						if _, ok := o.Misc["BotnetDomainsListsActionAllow"]; ok {
							nestedBotnetDomainsLists.Action.Allow.Misc = o.Misc["BotnetDomainsListsActionAllow"]
						}
					}
					if oBotnetDomainsLists.Action.Block != nil {
						nestedBotnetDomainsLists.Action.Block = &BotnetDomainsListsActionBlockXml{}
						if _, ok := o.Misc["BotnetDomainsListsActionBlock"]; ok {
							nestedBotnetDomainsLists.Action.Block.Misc = o.Misc["BotnetDomainsListsActionBlock"]
						}
					}
					if oBotnetDomainsLists.Action.Sinkhole != nil {
						nestedBotnetDomainsLists.Action.Sinkhole = &BotnetDomainsListsActionSinkholeXml{}
						if _, ok := o.Misc["BotnetDomainsListsActionSinkhole"]; ok {
							nestedBotnetDomainsLists.Action.Sinkhole.Misc = o.Misc["BotnetDomainsListsActionSinkhole"]
						}
					}
				}
				if oBotnetDomainsLists.PacketCapture != nil {
					nestedBotnetDomainsLists.PacketCapture = oBotnetDomainsLists.PacketCapture
				}
				if oBotnetDomainsLists.Name != "" {
					nestedBotnetDomainsLists.Name = oBotnetDomainsLists.Name
				}
				nestedBotnetDomains.Lists = append(nestedBotnetDomains.Lists, nestedBotnetDomainsLists)
			}
		}
	}
	entry.BotnetDomains = nestedBotnetDomains

	entry.CloudInlineAnalysis = util.YesNo(o.CloudInlineAnalysis, nil)
	entry.Description = o.Description
	entry.DisableOverride = o.DisableOverride
	entry.InlineExceptionEdlUrl = util.StrToMem(o.InlineExceptionEdlUrl)
	entry.InlineExceptionIpAddress = util.StrToMem(o.InlineExceptionIpAddress)
	var nestedMicaEngineSpywareEnabledCol []MicaEngineSpywareEnabledXml
	if o.MicaEngineSpywareEnabled != nil {
		nestedMicaEngineSpywareEnabledCol = []MicaEngineSpywareEnabledXml{}
		for _, oMicaEngineSpywareEnabled := range o.MicaEngineSpywareEnabled {
			nestedMicaEngineSpywareEnabled := MicaEngineSpywareEnabledXml{}
			if _, ok := o.Misc["MicaEngineSpywareEnabled"]; ok {
				nestedMicaEngineSpywareEnabled.Misc = o.Misc["MicaEngineSpywareEnabled"]
			}
			if oMicaEngineSpywareEnabled.InlinePolicyAction != nil {
				nestedMicaEngineSpywareEnabled.InlinePolicyAction = oMicaEngineSpywareEnabled.InlinePolicyAction
			}
			if oMicaEngineSpywareEnabled.Name != "" {
				nestedMicaEngineSpywareEnabled.Name = oMicaEngineSpywareEnabled.Name
			}
			nestedMicaEngineSpywareEnabledCol = append(nestedMicaEngineSpywareEnabledCol, nestedMicaEngineSpywareEnabled)
		}
		entry.MicaEngineSpywareEnabled = nestedMicaEngineSpywareEnabledCol
	}

	var nestedRulesCol []RulesXml
	if o.Rules != nil {
		nestedRulesCol = []RulesXml{}
		for _, oRules := range o.Rules {
			nestedRules := RulesXml{}
			if _, ok := o.Misc["Rules"]; ok {
				nestedRules.Misc = o.Misc["Rules"]
			}
			if oRules.Name != "" {
				nestedRules.Name = oRules.Name
			}
			if oRules.ThreatName != nil {
				nestedRules.ThreatName = oRules.ThreatName
			}
			if oRules.Category != nil {
				nestedRules.Category = oRules.Category
			}
			if oRules.PacketCapture != nil {
				nestedRules.PacketCapture = oRules.PacketCapture
			}
			if oRules.Severity != nil {
				nestedRules.Severity = util.StrToMem(oRules.Severity)
			}
			if oRules.Action != nil {
				nestedRules.Action = &RulesActionXml{}
				if _, ok := o.Misc["RulesAction"]; ok {
					nestedRules.Action.Misc = o.Misc["RulesAction"]
				}
				if oRules.Action.BlockIp != nil {
					nestedRules.Action.BlockIp = &RulesActionBlockIpXml{}
					if _, ok := o.Misc["RulesActionBlockIp"]; ok {
						nestedRules.Action.BlockIp.Misc = o.Misc["RulesActionBlockIp"]
					}
					if oRules.Action.BlockIp.TrackBy != nil {
						nestedRules.Action.BlockIp.TrackBy = oRules.Action.BlockIp.TrackBy
					}
					if oRules.Action.BlockIp.Duration != nil {
						nestedRules.Action.BlockIp.Duration = oRules.Action.BlockIp.Duration
					}
				}
				if oRules.Action.Default != nil {
					nestedRules.Action.Default = &RulesActionDefaultXml{}
					if _, ok := o.Misc["RulesActionDefault"]; ok {
						nestedRules.Action.Default.Misc = o.Misc["RulesActionDefault"]
					}
				}
				if oRules.Action.Allow != nil {
					nestedRules.Action.Allow = &RulesActionAllowXml{}
					if _, ok := o.Misc["RulesActionAllow"]; ok {
						nestedRules.Action.Allow.Misc = o.Misc["RulesActionAllow"]
					}
				}
				if oRules.Action.Alert != nil {
					nestedRules.Action.Alert = &RulesActionAlertXml{}
					if _, ok := o.Misc["RulesActionAlert"]; ok {
						nestedRules.Action.Alert.Misc = o.Misc["RulesActionAlert"]
					}
				}
				if oRules.Action.Drop != nil {
					nestedRules.Action.Drop = &RulesActionDropXml{}
					if _, ok := o.Misc["RulesActionDrop"]; ok {
						nestedRules.Action.Drop.Misc = o.Misc["RulesActionDrop"]
					}
				}
				if oRules.Action.ResetClient != nil {
					nestedRules.Action.ResetClient = &RulesActionResetClientXml{}
					if _, ok := o.Misc["RulesActionResetClient"]; ok {
						nestedRules.Action.ResetClient.Misc = o.Misc["RulesActionResetClient"]
					}
				}
				if oRules.Action.ResetServer != nil {
					nestedRules.Action.ResetServer = &RulesActionResetServerXml{}
					if _, ok := o.Misc["RulesActionResetServer"]; ok {
						nestedRules.Action.ResetServer.Misc = o.Misc["RulesActionResetServer"]
					}
				}
				if oRules.Action.ResetBoth != nil {
					nestedRules.Action.ResetBoth = &RulesActionResetBothXml{}
					if _, ok := o.Misc["RulesActionResetBoth"]; ok {
						nestedRules.Action.ResetBoth.Misc = o.Misc["RulesActionResetBoth"]
					}
				}
			}
			nestedRulesCol = append(nestedRulesCol, nestedRules)
		}
		entry.Rules = nestedRulesCol
	}

	var nestedThreatExceptionCol []ThreatExceptionXml
	if o.ThreatException != nil {
		nestedThreatExceptionCol = []ThreatExceptionXml{}
		for _, oThreatException := range o.ThreatException {
			nestedThreatException := ThreatExceptionXml{}
			if _, ok := o.Misc["ThreatException"]; ok {
				nestedThreatException.Misc = o.Misc["ThreatException"]
			}
			if oThreatException.PacketCapture != nil {
				nestedThreatException.PacketCapture = oThreatException.PacketCapture
			}
			if oThreatException.Action != nil {
				nestedThreatException.Action = &ThreatExceptionActionXml{}
				if _, ok := o.Misc["ThreatExceptionAction"]; ok {
					nestedThreatException.Action.Misc = o.Misc["ThreatExceptionAction"]
				}
				if oThreatException.Action.Allow != nil {
					nestedThreatException.Action.Allow = &ThreatExceptionActionAllowXml{}
					if _, ok := o.Misc["ThreatExceptionActionAllow"]; ok {
						nestedThreatException.Action.Allow.Misc = o.Misc["ThreatExceptionActionAllow"]
					}
				}
				if oThreatException.Action.Alert != nil {
					nestedThreatException.Action.Alert = &ThreatExceptionActionAlertXml{}
					if _, ok := o.Misc["ThreatExceptionActionAlert"]; ok {
						nestedThreatException.Action.Alert.Misc = o.Misc["ThreatExceptionActionAlert"]
					}
				}
				if oThreatException.Action.Drop != nil {
					nestedThreatException.Action.Drop = &ThreatExceptionActionDropXml{}
					if _, ok := o.Misc["ThreatExceptionActionDrop"]; ok {
						nestedThreatException.Action.Drop.Misc = o.Misc["ThreatExceptionActionDrop"]
					}
				}
				if oThreatException.Action.ResetBoth != nil {
					nestedThreatException.Action.ResetBoth = &ThreatExceptionActionResetBothXml{}
					if _, ok := o.Misc["ThreatExceptionActionResetBoth"]; ok {
						nestedThreatException.Action.ResetBoth.Misc = o.Misc["ThreatExceptionActionResetBoth"]
					}
				}
				if oThreatException.Action.ResetClient != nil {
					nestedThreatException.Action.ResetClient = &ThreatExceptionActionResetClientXml{}
					if _, ok := o.Misc["ThreatExceptionActionResetClient"]; ok {
						nestedThreatException.Action.ResetClient.Misc = o.Misc["ThreatExceptionActionResetClient"]
					}
				}
				if oThreatException.Action.ResetServer != nil {
					nestedThreatException.Action.ResetServer = &ThreatExceptionActionResetServerXml{}
					if _, ok := o.Misc["ThreatExceptionActionResetServer"]; ok {
						nestedThreatException.Action.ResetServer.Misc = o.Misc["ThreatExceptionActionResetServer"]
					}
				}
				if oThreatException.Action.BlockIp != nil {
					nestedThreatException.Action.BlockIp = &ThreatExceptionActionBlockIpXml{}
					if _, ok := o.Misc["ThreatExceptionActionBlockIp"]; ok {
						nestedThreatException.Action.BlockIp.Misc = o.Misc["ThreatExceptionActionBlockIp"]
					}
					if oThreatException.Action.BlockIp.TrackBy != nil {
						nestedThreatException.Action.BlockIp.TrackBy = oThreatException.Action.BlockIp.TrackBy
					}
					if oThreatException.Action.BlockIp.Duration != nil {
						nestedThreatException.Action.BlockIp.Duration = oThreatException.Action.BlockIp.Duration
					}
				}
				if oThreatException.Action.Default != nil {
					nestedThreatException.Action.Default = &ThreatExceptionActionDefaultXml{}
					if _, ok := o.Misc["ThreatExceptionActionDefault"]; ok {
						nestedThreatException.Action.Default.Misc = o.Misc["ThreatExceptionActionDefault"]
					}
				}
			}
			if oThreatException.ExemptIp != nil {
				nestedThreatException.ExemptIp = []ThreatExceptionExemptIpXml{}
				for _, oThreatExceptionExemptIp := range oThreatException.ExemptIp {
					nestedThreatExceptionExemptIp := ThreatExceptionExemptIpXml{}
					if _, ok := o.Misc["ThreatExceptionExemptIp"]; ok {
						nestedThreatExceptionExemptIp.Misc = o.Misc["ThreatExceptionExemptIp"]
					}
					if oThreatExceptionExemptIp.Name != "" {
						nestedThreatExceptionExemptIp.Name = oThreatExceptionExemptIp.Name
					}
					nestedThreatException.ExemptIp = append(nestedThreatException.ExemptIp, nestedThreatExceptionExemptIp)
				}
			}
			if oThreatException.Name != "" {
				nestedThreatException.Name = oThreatException.Name
			}
			nestedThreatExceptionCol = append(nestedThreatExceptionCol, nestedThreatException)
		}
		entry.ThreatException = nestedThreatExceptionCol
	}

	entry.Misc = o.Misc["Entry"]

	return entry, nil
}

func specifyEntry_11_0_2(o *Entry) (any, error) {
	entry := entryXml_11_0_2{}
	entry.Name = o.Name
	var nestedBotnetDomains *BotnetDomainsXml_11_0_2
	if o.BotnetDomains != nil {
		nestedBotnetDomains = &BotnetDomainsXml_11_0_2{}
		if _, ok := o.Misc["BotnetDomains"]; ok {
			nestedBotnetDomains.Misc = o.Misc["BotnetDomains"]
		}
		if o.BotnetDomains.DnsSecurityCategories != nil {
			nestedBotnetDomains.DnsSecurityCategories = []BotnetDomainsDnsSecurityCategoriesXml_11_0_2{}
			for _, oBotnetDomainsDnsSecurityCategories := range o.BotnetDomains.DnsSecurityCategories {
				nestedBotnetDomainsDnsSecurityCategories := BotnetDomainsDnsSecurityCategoriesXml_11_0_2{}
				if _, ok := o.Misc["BotnetDomainsDnsSecurityCategories"]; ok {
					nestedBotnetDomainsDnsSecurityCategories.Misc = o.Misc["BotnetDomainsDnsSecurityCategories"]
				}
				if oBotnetDomainsDnsSecurityCategories.Name != "" {
					nestedBotnetDomainsDnsSecurityCategories.Name = oBotnetDomainsDnsSecurityCategories.Name
				}
				if oBotnetDomainsDnsSecurityCategories.Action != nil {
					nestedBotnetDomainsDnsSecurityCategories.Action = oBotnetDomainsDnsSecurityCategories.Action
				}
				if oBotnetDomainsDnsSecurityCategories.LogLevel != nil {
					nestedBotnetDomainsDnsSecurityCategories.LogLevel = oBotnetDomainsDnsSecurityCategories.LogLevel
				}
				if oBotnetDomainsDnsSecurityCategories.PacketCapture != nil {
					nestedBotnetDomainsDnsSecurityCategories.PacketCapture = oBotnetDomainsDnsSecurityCategories.PacketCapture
				}
				nestedBotnetDomains.DnsSecurityCategories = append(nestedBotnetDomains.DnsSecurityCategories, nestedBotnetDomainsDnsSecurityCategories)
			}
		}
		if o.BotnetDomains.Lists != nil {
			nestedBotnetDomains.Lists = []BotnetDomainsListsXml_11_0_2{}
			for _, oBotnetDomainsLists := range o.BotnetDomains.Lists {
				nestedBotnetDomainsLists := BotnetDomainsListsXml_11_0_2{}
				if _, ok := o.Misc["BotnetDomainsLists"]; ok {
					nestedBotnetDomainsLists.Misc = o.Misc["BotnetDomainsLists"]
				}
				if oBotnetDomainsLists.Action != nil {
					nestedBotnetDomainsLists.Action = &BotnetDomainsListsActionXml_11_0_2{}
					if _, ok := o.Misc["BotnetDomainsListsAction"]; ok {
						nestedBotnetDomainsLists.Action.Misc = o.Misc["BotnetDomainsListsAction"]
					}
					if oBotnetDomainsLists.Action.Alert != nil {
						nestedBotnetDomainsLists.Action.Alert = &BotnetDomainsListsActionAlertXml_11_0_2{}
						if _, ok := o.Misc["BotnetDomainsListsActionAlert"]; ok {
							nestedBotnetDomainsLists.Action.Alert.Misc = o.Misc["BotnetDomainsListsActionAlert"]
						}
					}
					if oBotnetDomainsLists.Action.Allow != nil {
						nestedBotnetDomainsLists.Action.Allow = &BotnetDomainsListsActionAllowXml_11_0_2{}
						if _, ok := o.Misc["BotnetDomainsListsActionAllow"]; ok {
							nestedBotnetDomainsLists.Action.Allow.Misc = o.Misc["BotnetDomainsListsActionAllow"]
						}
					}
					if oBotnetDomainsLists.Action.Block != nil {
						nestedBotnetDomainsLists.Action.Block = &BotnetDomainsListsActionBlockXml_11_0_2{}
						if _, ok := o.Misc["BotnetDomainsListsActionBlock"]; ok {
							nestedBotnetDomainsLists.Action.Block.Misc = o.Misc["BotnetDomainsListsActionBlock"]
						}
					}
					if oBotnetDomainsLists.Action.Sinkhole != nil {
						nestedBotnetDomainsLists.Action.Sinkhole = &BotnetDomainsListsActionSinkholeXml_11_0_2{}
						if _, ok := o.Misc["BotnetDomainsListsActionSinkhole"]; ok {
							nestedBotnetDomainsLists.Action.Sinkhole.Misc = o.Misc["BotnetDomainsListsActionSinkhole"]
						}
					}
				}
				if oBotnetDomainsLists.PacketCapture != nil {
					nestedBotnetDomainsLists.PacketCapture = oBotnetDomainsLists.PacketCapture
				}
				if oBotnetDomainsLists.Name != "" {
					nestedBotnetDomainsLists.Name = oBotnetDomainsLists.Name
				}
				nestedBotnetDomains.Lists = append(nestedBotnetDomains.Lists, nestedBotnetDomainsLists)
			}
		}
		if o.BotnetDomains.Sinkhole != nil {
			nestedBotnetDomains.Sinkhole = &BotnetDomainsSinkholeXml_11_0_2{}
			if _, ok := o.Misc["BotnetDomainsSinkhole"]; ok {
				nestedBotnetDomains.Sinkhole.Misc = o.Misc["BotnetDomainsSinkhole"]
			}
			if o.BotnetDomains.Sinkhole.Ipv4Address != nil {
				nestedBotnetDomains.Sinkhole.Ipv4Address = o.BotnetDomains.Sinkhole.Ipv4Address
			}
			if o.BotnetDomains.Sinkhole.Ipv6Address != nil {
				nestedBotnetDomains.Sinkhole.Ipv6Address = o.BotnetDomains.Sinkhole.Ipv6Address
			}
		}
		if o.BotnetDomains.ThreatException != nil {
			nestedBotnetDomains.ThreatException = []BotnetDomainsThreatExceptionXml_11_0_2{}
			for _, oBotnetDomainsThreatException := range o.BotnetDomains.ThreatException {
				nestedBotnetDomainsThreatException := BotnetDomainsThreatExceptionXml_11_0_2{}
				if _, ok := o.Misc["BotnetDomainsThreatException"]; ok {
					nestedBotnetDomainsThreatException.Misc = o.Misc["BotnetDomainsThreatException"]
				}
				if oBotnetDomainsThreatException.Name != "" {
					nestedBotnetDomainsThreatException.Name = oBotnetDomainsThreatException.Name
				}
				nestedBotnetDomains.ThreatException = append(nestedBotnetDomains.ThreatException, nestedBotnetDomainsThreatException)
			}
		}
		if o.BotnetDomains.Whitelist != nil {
			nestedBotnetDomains.Whitelist = []BotnetDomainsWhitelistXml_11_0_2{}
			for _, oBotnetDomainsWhitelist := range o.BotnetDomains.Whitelist {
				nestedBotnetDomainsWhitelist := BotnetDomainsWhitelistXml_11_0_2{}
				if _, ok := o.Misc["BotnetDomainsWhitelist"]; ok {
					nestedBotnetDomainsWhitelist.Misc = o.Misc["BotnetDomainsWhitelist"]
				}
				if oBotnetDomainsWhitelist.Description != nil {
					nestedBotnetDomainsWhitelist.Description = oBotnetDomainsWhitelist.Description
				}
				if oBotnetDomainsWhitelist.Name != "" {
					nestedBotnetDomainsWhitelist.Name = oBotnetDomainsWhitelist.Name
				}
				nestedBotnetDomains.Whitelist = append(nestedBotnetDomains.Whitelist, nestedBotnetDomainsWhitelist)
			}
		}
		if o.BotnetDomains.RtypeAction != nil {
			nestedBotnetDomains.RtypeAction = &BotnetDomainsRtypeActionXml_11_0_2{}
			if _, ok := o.Misc["BotnetDomainsRtypeAction"]; ok {
				nestedBotnetDomains.RtypeAction.Misc = o.Misc["BotnetDomainsRtypeAction"]
			}
			if o.BotnetDomains.RtypeAction.Any != nil {
				nestedBotnetDomains.RtypeAction.Any = o.BotnetDomains.RtypeAction.Any
			}
			if o.BotnetDomains.RtypeAction.Https != nil {
				nestedBotnetDomains.RtypeAction.Https = o.BotnetDomains.RtypeAction.Https
			}
			if o.BotnetDomains.RtypeAction.Svcb != nil {
				nestedBotnetDomains.RtypeAction.Svcb = o.BotnetDomains.RtypeAction.Svcb
			}
		}
	}
	entry.BotnetDomains = nestedBotnetDomains

	entry.CloudInlineAnalysis = util.YesNo(o.CloudInlineAnalysis, nil)
	entry.Description = o.Description
	entry.DisableOverride = o.DisableOverride
	entry.InlineExceptionEdlUrl = util.StrToMem(o.InlineExceptionEdlUrl)
	entry.InlineExceptionIpAddress = util.StrToMem(o.InlineExceptionIpAddress)
	var nestedMicaEngineSpywareEnabledCol []MicaEngineSpywareEnabledXml_11_0_2
	if o.MicaEngineSpywareEnabled != nil {
		nestedMicaEngineSpywareEnabledCol = []MicaEngineSpywareEnabledXml_11_0_2{}
		for _, oMicaEngineSpywareEnabled := range o.MicaEngineSpywareEnabled {
			nestedMicaEngineSpywareEnabled := MicaEngineSpywareEnabledXml_11_0_2{}
			if _, ok := o.Misc["MicaEngineSpywareEnabled"]; ok {
				nestedMicaEngineSpywareEnabled.Misc = o.Misc["MicaEngineSpywareEnabled"]
			}
			if oMicaEngineSpywareEnabled.InlinePolicyAction != nil {
				nestedMicaEngineSpywareEnabled.InlinePolicyAction = oMicaEngineSpywareEnabled.InlinePolicyAction
			}
			if oMicaEngineSpywareEnabled.Name != "" {
				nestedMicaEngineSpywareEnabled.Name = oMicaEngineSpywareEnabled.Name
			}
			nestedMicaEngineSpywareEnabledCol = append(nestedMicaEngineSpywareEnabledCol, nestedMicaEngineSpywareEnabled)
		}
		entry.MicaEngineSpywareEnabled = nestedMicaEngineSpywareEnabledCol
	}

	var nestedRulesCol []RulesXml_11_0_2
	if o.Rules != nil {
		nestedRulesCol = []RulesXml_11_0_2{}
		for _, oRules := range o.Rules {
			nestedRules := RulesXml_11_0_2{}
			if _, ok := o.Misc["Rules"]; ok {
				nestedRules.Misc = o.Misc["Rules"]
			}
			if oRules.PacketCapture != nil {
				nestedRules.PacketCapture = oRules.PacketCapture
			}
			if oRules.Severity != nil {
				nestedRules.Severity = util.StrToMem(oRules.Severity)
			}
			if oRules.Action != nil {
				nestedRules.Action = &RulesActionXml_11_0_2{}
				if _, ok := o.Misc["RulesAction"]; ok {
					nestedRules.Action.Misc = o.Misc["RulesAction"]
				}
				if oRules.Action.Alert != nil {
					nestedRules.Action.Alert = &RulesActionAlertXml_11_0_2{}
					if _, ok := o.Misc["RulesActionAlert"]; ok {
						nestedRules.Action.Alert.Misc = o.Misc["RulesActionAlert"]
					}
				}
				if oRules.Action.Drop != nil {
					nestedRules.Action.Drop = &RulesActionDropXml_11_0_2{}
					if _, ok := o.Misc["RulesActionDrop"]; ok {
						nestedRules.Action.Drop.Misc = o.Misc["RulesActionDrop"]
					}
				}
				if oRules.Action.ResetClient != nil {
					nestedRules.Action.ResetClient = &RulesActionResetClientXml_11_0_2{}
					if _, ok := o.Misc["RulesActionResetClient"]; ok {
						nestedRules.Action.ResetClient.Misc = o.Misc["RulesActionResetClient"]
					}
				}
				if oRules.Action.ResetServer != nil {
					nestedRules.Action.ResetServer = &RulesActionResetServerXml_11_0_2{}
					if _, ok := o.Misc["RulesActionResetServer"]; ok {
						nestedRules.Action.ResetServer.Misc = o.Misc["RulesActionResetServer"]
					}
				}
				if oRules.Action.ResetBoth != nil {
					nestedRules.Action.ResetBoth = &RulesActionResetBothXml_11_0_2{}
					if _, ok := o.Misc["RulesActionResetBoth"]; ok {
						nestedRules.Action.ResetBoth.Misc = o.Misc["RulesActionResetBoth"]
					}
				}
				if oRules.Action.BlockIp != nil {
					nestedRules.Action.BlockIp = &RulesActionBlockIpXml_11_0_2{}
					if _, ok := o.Misc["RulesActionBlockIp"]; ok {
						nestedRules.Action.BlockIp.Misc = o.Misc["RulesActionBlockIp"]
					}
					if oRules.Action.BlockIp.TrackBy != nil {
						nestedRules.Action.BlockIp.TrackBy = oRules.Action.BlockIp.TrackBy
					}
					if oRules.Action.BlockIp.Duration != nil {
						nestedRules.Action.BlockIp.Duration = oRules.Action.BlockIp.Duration
					}
				}
				if oRules.Action.Default != nil {
					nestedRules.Action.Default = &RulesActionDefaultXml_11_0_2{}
					if _, ok := o.Misc["RulesActionDefault"]; ok {
						nestedRules.Action.Default.Misc = o.Misc["RulesActionDefault"]
					}
				}
				if oRules.Action.Allow != nil {
					nestedRules.Action.Allow = &RulesActionAllowXml_11_0_2{}
					if _, ok := o.Misc["RulesActionAllow"]; ok {
						nestedRules.Action.Allow.Misc = o.Misc["RulesActionAllow"]
					}
				}
			}
			if oRules.Name != "" {
				nestedRules.Name = oRules.Name
			}
			if oRules.ThreatName != nil {
				nestedRules.ThreatName = oRules.ThreatName
			}
			if oRules.Category != nil {
				nestedRules.Category = oRules.Category
			}
			nestedRulesCol = append(nestedRulesCol, nestedRules)
		}
		entry.Rules = nestedRulesCol
	}

	var nestedThreatExceptionCol []ThreatExceptionXml_11_0_2
	if o.ThreatException != nil {
		nestedThreatExceptionCol = []ThreatExceptionXml_11_0_2{}
		for _, oThreatException := range o.ThreatException {
			nestedThreatException := ThreatExceptionXml_11_0_2{}
			if _, ok := o.Misc["ThreatException"]; ok {
				nestedThreatException.Misc = o.Misc["ThreatException"]
			}
			if oThreatException.ExemptIp != nil {
				nestedThreatException.ExemptIp = []ThreatExceptionExemptIpXml_11_0_2{}
				for _, oThreatExceptionExemptIp := range oThreatException.ExemptIp {
					nestedThreatExceptionExemptIp := ThreatExceptionExemptIpXml_11_0_2{}
					if _, ok := o.Misc["ThreatExceptionExemptIp"]; ok {
						nestedThreatExceptionExemptIp.Misc = o.Misc["ThreatExceptionExemptIp"]
					}
					if oThreatExceptionExemptIp.Name != "" {
						nestedThreatExceptionExemptIp.Name = oThreatExceptionExemptIp.Name
					}
					nestedThreatException.ExemptIp = append(nestedThreatException.ExemptIp, nestedThreatExceptionExemptIp)
				}
			}
			if oThreatException.Name != "" {
				nestedThreatException.Name = oThreatException.Name
			}
			if oThreatException.PacketCapture != nil {
				nestedThreatException.PacketCapture = oThreatException.PacketCapture
			}
			if oThreatException.Action != nil {
				nestedThreatException.Action = &ThreatExceptionActionXml_11_0_2{}
				if _, ok := o.Misc["ThreatExceptionAction"]; ok {
					nestedThreatException.Action.Misc = o.Misc["ThreatExceptionAction"]
				}
				if oThreatException.Action.BlockIp != nil {
					nestedThreatException.Action.BlockIp = &ThreatExceptionActionBlockIpXml_11_0_2{}
					if _, ok := o.Misc["ThreatExceptionActionBlockIp"]; ok {
						nestedThreatException.Action.BlockIp.Misc = o.Misc["ThreatExceptionActionBlockIp"]
					}
					if oThreatException.Action.BlockIp.TrackBy != nil {
						nestedThreatException.Action.BlockIp.TrackBy = oThreatException.Action.BlockIp.TrackBy
					}
					if oThreatException.Action.BlockIp.Duration != nil {
						nestedThreatException.Action.BlockIp.Duration = oThreatException.Action.BlockIp.Duration
					}
				}
				if oThreatException.Action.Default != nil {
					nestedThreatException.Action.Default = &ThreatExceptionActionDefaultXml_11_0_2{}
					if _, ok := o.Misc["ThreatExceptionActionDefault"]; ok {
						nestedThreatException.Action.Default.Misc = o.Misc["ThreatExceptionActionDefault"]
					}
				}
				if oThreatException.Action.Allow != nil {
					nestedThreatException.Action.Allow = &ThreatExceptionActionAllowXml_11_0_2{}
					if _, ok := o.Misc["ThreatExceptionActionAllow"]; ok {
						nestedThreatException.Action.Allow.Misc = o.Misc["ThreatExceptionActionAllow"]
					}
				}
				if oThreatException.Action.Alert != nil {
					nestedThreatException.Action.Alert = &ThreatExceptionActionAlertXml_11_0_2{}
					if _, ok := o.Misc["ThreatExceptionActionAlert"]; ok {
						nestedThreatException.Action.Alert.Misc = o.Misc["ThreatExceptionActionAlert"]
					}
				}
				if oThreatException.Action.Drop != nil {
					nestedThreatException.Action.Drop = &ThreatExceptionActionDropXml_11_0_2{}
					if _, ok := o.Misc["ThreatExceptionActionDrop"]; ok {
						nestedThreatException.Action.Drop.Misc = o.Misc["ThreatExceptionActionDrop"]
					}
				}
				if oThreatException.Action.ResetBoth != nil {
					nestedThreatException.Action.ResetBoth = &ThreatExceptionActionResetBothXml_11_0_2{}
					if _, ok := o.Misc["ThreatExceptionActionResetBoth"]; ok {
						nestedThreatException.Action.ResetBoth.Misc = o.Misc["ThreatExceptionActionResetBoth"]
					}
				}
				if oThreatException.Action.ResetClient != nil {
					nestedThreatException.Action.ResetClient = &ThreatExceptionActionResetClientXml_11_0_2{}
					if _, ok := o.Misc["ThreatExceptionActionResetClient"]; ok {
						nestedThreatException.Action.ResetClient.Misc = o.Misc["ThreatExceptionActionResetClient"]
					}
				}
				if oThreatException.Action.ResetServer != nil {
					nestedThreatException.Action.ResetServer = &ThreatExceptionActionResetServerXml_11_0_2{}
					if _, ok := o.Misc["ThreatExceptionActionResetServer"]; ok {
						nestedThreatException.Action.ResetServer.Misc = o.Misc["ThreatExceptionActionResetServer"]
					}
				}
			}
			nestedThreatExceptionCol = append(nestedThreatExceptionCol, nestedThreatException)
		}
		entry.ThreatException = nestedThreatExceptionCol
	}

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
		var nestedBotnetDomains *BotnetDomains
		if o.BotnetDomains != nil {
			nestedBotnetDomains = &BotnetDomains{}
			if o.BotnetDomains.Misc != nil {
				entry.Misc["BotnetDomains"] = o.BotnetDomains.Misc
			}
			if o.BotnetDomains.Lists != nil {
				nestedBotnetDomains.Lists = []BotnetDomainsLists{}
				for _, oBotnetDomainsLists := range o.BotnetDomains.Lists {
					nestedBotnetDomainsLists := BotnetDomainsLists{}
					if oBotnetDomainsLists.Misc != nil {
						entry.Misc["BotnetDomainsLists"] = oBotnetDomainsLists.Misc
					}
					if oBotnetDomainsLists.Name != "" {
						nestedBotnetDomainsLists.Name = oBotnetDomainsLists.Name
					}
					if oBotnetDomainsLists.Action != nil {
						nestedBotnetDomainsLists.Action = &BotnetDomainsListsAction{}
						if oBotnetDomainsLists.Action.Misc != nil {
							entry.Misc["BotnetDomainsListsAction"] = oBotnetDomainsLists.Action.Misc
						}
						if oBotnetDomainsLists.Action.Alert != nil {
							nestedBotnetDomainsLists.Action.Alert = &BotnetDomainsListsActionAlert{}
							if oBotnetDomainsLists.Action.Alert.Misc != nil {
								entry.Misc["BotnetDomainsListsActionAlert"] = oBotnetDomainsLists.Action.Alert.Misc
							}
						}
						if oBotnetDomainsLists.Action.Allow != nil {
							nestedBotnetDomainsLists.Action.Allow = &BotnetDomainsListsActionAllow{}
							if oBotnetDomainsLists.Action.Allow.Misc != nil {
								entry.Misc["BotnetDomainsListsActionAllow"] = oBotnetDomainsLists.Action.Allow.Misc
							}
						}
						if oBotnetDomainsLists.Action.Block != nil {
							nestedBotnetDomainsLists.Action.Block = &BotnetDomainsListsActionBlock{}
							if oBotnetDomainsLists.Action.Block.Misc != nil {
								entry.Misc["BotnetDomainsListsActionBlock"] = oBotnetDomainsLists.Action.Block.Misc
							}
						}
						if oBotnetDomainsLists.Action.Sinkhole != nil {
							nestedBotnetDomainsLists.Action.Sinkhole = &BotnetDomainsListsActionSinkhole{}
							if oBotnetDomainsLists.Action.Sinkhole.Misc != nil {
								entry.Misc["BotnetDomainsListsActionSinkhole"] = oBotnetDomainsLists.Action.Sinkhole.Misc
							}
						}
					}
					if oBotnetDomainsLists.PacketCapture != nil {
						nestedBotnetDomainsLists.PacketCapture = oBotnetDomainsLists.PacketCapture
					}
					nestedBotnetDomains.Lists = append(nestedBotnetDomains.Lists, nestedBotnetDomainsLists)
				}
			}
			if o.BotnetDomains.Sinkhole != nil {
				nestedBotnetDomains.Sinkhole = &BotnetDomainsSinkhole{}
				if o.BotnetDomains.Sinkhole.Misc != nil {
					entry.Misc["BotnetDomainsSinkhole"] = o.BotnetDomains.Sinkhole.Misc
				}
				if o.BotnetDomains.Sinkhole.Ipv4Address != nil {
					nestedBotnetDomains.Sinkhole.Ipv4Address = o.BotnetDomains.Sinkhole.Ipv4Address
				}
				if o.BotnetDomains.Sinkhole.Ipv6Address != nil {
					nestedBotnetDomains.Sinkhole.Ipv6Address = o.BotnetDomains.Sinkhole.Ipv6Address
				}
			}
			if o.BotnetDomains.ThreatException != nil {
				nestedBotnetDomains.ThreatException = []BotnetDomainsThreatException{}
				for _, oBotnetDomainsThreatException := range o.BotnetDomains.ThreatException {
					nestedBotnetDomainsThreatException := BotnetDomainsThreatException{}
					if oBotnetDomainsThreatException.Misc != nil {
						entry.Misc["BotnetDomainsThreatException"] = oBotnetDomainsThreatException.Misc
					}
					if oBotnetDomainsThreatException.Name != "" {
						nestedBotnetDomainsThreatException.Name = oBotnetDomainsThreatException.Name
					}
					nestedBotnetDomains.ThreatException = append(nestedBotnetDomains.ThreatException, nestedBotnetDomainsThreatException)
				}
			}
			if o.BotnetDomains.Whitelist != nil {
				nestedBotnetDomains.Whitelist = []BotnetDomainsWhitelist{}
				for _, oBotnetDomainsWhitelist := range o.BotnetDomains.Whitelist {
					nestedBotnetDomainsWhitelist := BotnetDomainsWhitelist{}
					if oBotnetDomainsWhitelist.Misc != nil {
						entry.Misc["BotnetDomainsWhitelist"] = oBotnetDomainsWhitelist.Misc
					}
					if oBotnetDomainsWhitelist.Description != nil {
						nestedBotnetDomainsWhitelist.Description = oBotnetDomainsWhitelist.Description
					}
					if oBotnetDomainsWhitelist.Name != "" {
						nestedBotnetDomainsWhitelist.Name = oBotnetDomainsWhitelist.Name
					}
					nestedBotnetDomains.Whitelist = append(nestedBotnetDomains.Whitelist, nestedBotnetDomainsWhitelist)
				}
			}
			if o.BotnetDomains.RtypeAction != nil {
				nestedBotnetDomains.RtypeAction = &BotnetDomainsRtypeAction{}
				if o.BotnetDomains.RtypeAction.Misc != nil {
					entry.Misc["BotnetDomainsRtypeAction"] = o.BotnetDomains.RtypeAction.Misc
				}
				if o.BotnetDomains.RtypeAction.Https != nil {
					nestedBotnetDomains.RtypeAction.Https = o.BotnetDomains.RtypeAction.Https
				}
				if o.BotnetDomains.RtypeAction.Svcb != nil {
					nestedBotnetDomains.RtypeAction.Svcb = o.BotnetDomains.RtypeAction.Svcb
				}
				if o.BotnetDomains.RtypeAction.Any != nil {
					nestedBotnetDomains.RtypeAction.Any = o.BotnetDomains.RtypeAction.Any
				}
			}
			if o.BotnetDomains.DnsSecurityCategories != nil {
				nestedBotnetDomains.DnsSecurityCategories = []BotnetDomainsDnsSecurityCategories{}
				for _, oBotnetDomainsDnsSecurityCategories := range o.BotnetDomains.DnsSecurityCategories {
					nestedBotnetDomainsDnsSecurityCategories := BotnetDomainsDnsSecurityCategories{}
					if oBotnetDomainsDnsSecurityCategories.Misc != nil {
						entry.Misc["BotnetDomainsDnsSecurityCategories"] = oBotnetDomainsDnsSecurityCategories.Misc
					}
					if oBotnetDomainsDnsSecurityCategories.LogLevel != nil {
						nestedBotnetDomainsDnsSecurityCategories.LogLevel = oBotnetDomainsDnsSecurityCategories.LogLevel
					}
					if oBotnetDomainsDnsSecurityCategories.PacketCapture != nil {
						nestedBotnetDomainsDnsSecurityCategories.PacketCapture = oBotnetDomainsDnsSecurityCategories.PacketCapture
					}
					if oBotnetDomainsDnsSecurityCategories.Name != "" {
						nestedBotnetDomainsDnsSecurityCategories.Name = oBotnetDomainsDnsSecurityCategories.Name
					}
					if oBotnetDomainsDnsSecurityCategories.Action != nil {
						nestedBotnetDomainsDnsSecurityCategories.Action = oBotnetDomainsDnsSecurityCategories.Action
					}
					nestedBotnetDomains.DnsSecurityCategories = append(nestedBotnetDomains.DnsSecurityCategories, nestedBotnetDomainsDnsSecurityCategories)
				}
			}
		}
		entry.BotnetDomains = nestedBotnetDomains

		entry.CloudInlineAnalysis = util.AsBool(o.CloudInlineAnalysis, nil)
		entry.Description = o.Description
		entry.DisableOverride = o.DisableOverride
		entry.InlineExceptionEdlUrl = util.MemToStr(o.InlineExceptionEdlUrl)
		entry.InlineExceptionIpAddress = util.MemToStr(o.InlineExceptionIpAddress)
		var nestedMicaEngineSpywareEnabledCol []MicaEngineSpywareEnabled
		if o.MicaEngineSpywareEnabled != nil {
			nestedMicaEngineSpywareEnabledCol = []MicaEngineSpywareEnabled{}
			for _, oMicaEngineSpywareEnabled := range o.MicaEngineSpywareEnabled {
				nestedMicaEngineSpywareEnabled := MicaEngineSpywareEnabled{}
				if oMicaEngineSpywareEnabled.Misc != nil {
					entry.Misc["MicaEngineSpywareEnabled"] = oMicaEngineSpywareEnabled.Misc
				}
				if oMicaEngineSpywareEnabled.InlinePolicyAction != nil {
					nestedMicaEngineSpywareEnabled.InlinePolicyAction = oMicaEngineSpywareEnabled.InlinePolicyAction
				}
				if oMicaEngineSpywareEnabled.Name != "" {
					nestedMicaEngineSpywareEnabled.Name = oMicaEngineSpywareEnabled.Name
				}
				nestedMicaEngineSpywareEnabledCol = append(nestedMicaEngineSpywareEnabledCol, nestedMicaEngineSpywareEnabled)
			}
			entry.MicaEngineSpywareEnabled = nestedMicaEngineSpywareEnabledCol
		}

		var nestedRulesCol []Rules
		if o.Rules != nil {
			nestedRulesCol = []Rules{}
			for _, oRules := range o.Rules {
				nestedRules := Rules{}
				if oRules.Misc != nil {
					entry.Misc["Rules"] = oRules.Misc
				}
				if oRules.Category != nil {
					nestedRules.Category = oRules.Category
				}
				if oRules.PacketCapture != nil {
					nestedRules.PacketCapture = oRules.PacketCapture
				}
				if oRules.Severity != nil {
					nestedRules.Severity = util.MemToStr(oRules.Severity)
				}
				if oRules.Action != nil {
					nestedRules.Action = &RulesAction{}
					if oRules.Action.Misc != nil {
						entry.Misc["RulesAction"] = oRules.Action.Misc
					}
					if oRules.Action.Allow != nil {
						nestedRules.Action.Allow = &RulesActionAllow{}
						if oRules.Action.Allow.Misc != nil {
							entry.Misc["RulesActionAllow"] = oRules.Action.Allow.Misc
						}
					}
					if oRules.Action.Alert != nil {
						nestedRules.Action.Alert = &RulesActionAlert{}
						if oRules.Action.Alert.Misc != nil {
							entry.Misc["RulesActionAlert"] = oRules.Action.Alert.Misc
						}
					}
					if oRules.Action.Drop != nil {
						nestedRules.Action.Drop = &RulesActionDrop{}
						if oRules.Action.Drop.Misc != nil {
							entry.Misc["RulesActionDrop"] = oRules.Action.Drop.Misc
						}
					}
					if oRules.Action.ResetClient != nil {
						nestedRules.Action.ResetClient = &RulesActionResetClient{}
						if oRules.Action.ResetClient.Misc != nil {
							entry.Misc["RulesActionResetClient"] = oRules.Action.ResetClient.Misc
						}
					}
					if oRules.Action.ResetServer != nil {
						nestedRules.Action.ResetServer = &RulesActionResetServer{}
						if oRules.Action.ResetServer.Misc != nil {
							entry.Misc["RulesActionResetServer"] = oRules.Action.ResetServer.Misc
						}
					}
					if oRules.Action.ResetBoth != nil {
						nestedRules.Action.ResetBoth = &RulesActionResetBoth{}
						if oRules.Action.ResetBoth.Misc != nil {
							entry.Misc["RulesActionResetBoth"] = oRules.Action.ResetBoth.Misc
						}
					}
					if oRules.Action.BlockIp != nil {
						nestedRules.Action.BlockIp = &RulesActionBlockIp{}
						if oRules.Action.BlockIp.Misc != nil {
							entry.Misc["RulesActionBlockIp"] = oRules.Action.BlockIp.Misc
						}
						if oRules.Action.BlockIp.TrackBy != nil {
							nestedRules.Action.BlockIp.TrackBy = oRules.Action.BlockIp.TrackBy
						}
						if oRules.Action.BlockIp.Duration != nil {
							nestedRules.Action.BlockIp.Duration = oRules.Action.BlockIp.Duration
						}
					}
					if oRules.Action.Default != nil {
						nestedRules.Action.Default = &RulesActionDefault{}
						if oRules.Action.Default.Misc != nil {
							entry.Misc["RulesActionDefault"] = oRules.Action.Default.Misc
						}
					}
				}
				if oRules.Name != "" {
					nestedRules.Name = oRules.Name
				}
				if oRules.ThreatName != nil {
					nestedRules.ThreatName = oRules.ThreatName
				}
				nestedRulesCol = append(nestedRulesCol, nestedRules)
			}
			entry.Rules = nestedRulesCol
		}

		var nestedThreatExceptionCol []ThreatException
		if o.ThreatException != nil {
			nestedThreatExceptionCol = []ThreatException{}
			for _, oThreatException := range o.ThreatException {
				nestedThreatException := ThreatException{}
				if oThreatException.Misc != nil {
					entry.Misc["ThreatException"] = oThreatException.Misc
				}
				if oThreatException.Action != nil {
					nestedThreatException.Action = &ThreatExceptionAction{}
					if oThreatException.Action.Misc != nil {
						entry.Misc["ThreatExceptionAction"] = oThreatException.Action.Misc
					}
					if oThreatException.Action.BlockIp != nil {
						nestedThreatException.Action.BlockIp = &ThreatExceptionActionBlockIp{}
						if oThreatException.Action.BlockIp.Misc != nil {
							entry.Misc["ThreatExceptionActionBlockIp"] = oThreatException.Action.BlockIp.Misc
						}
						if oThreatException.Action.BlockIp.TrackBy != nil {
							nestedThreatException.Action.BlockIp.TrackBy = oThreatException.Action.BlockIp.TrackBy
						}
						if oThreatException.Action.BlockIp.Duration != nil {
							nestedThreatException.Action.BlockIp.Duration = oThreatException.Action.BlockIp.Duration
						}
					}
					if oThreatException.Action.Default != nil {
						nestedThreatException.Action.Default = &ThreatExceptionActionDefault{}
						if oThreatException.Action.Default.Misc != nil {
							entry.Misc["ThreatExceptionActionDefault"] = oThreatException.Action.Default.Misc
						}
					}
					if oThreatException.Action.Allow != nil {
						nestedThreatException.Action.Allow = &ThreatExceptionActionAllow{}
						if oThreatException.Action.Allow.Misc != nil {
							entry.Misc["ThreatExceptionActionAllow"] = oThreatException.Action.Allow.Misc
						}
					}
					if oThreatException.Action.Alert != nil {
						nestedThreatException.Action.Alert = &ThreatExceptionActionAlert{}
						if oThreatException.Action.Alert.Misc != nil {
							entry.Misc["ThreatExceptionActionAlert"] = oThreatException.Action.Alert.Misc
						}
					}
					if oThreatException.Action.Drop != nil {
						nestedThreatException.Action.Drop = &ThreatExceptionActionDrop{}
						if oThreatException.Action.Drop.Misc != nil {
							entry.Misc["ThreatExceptionActionDrop"] = oThreatException.Action.Drop.Misc
						}
					}
					if oThreatException.Action.ResetBoth != nil {
						nestedThreatException.Action.ResetBoth = &ThreatExceptionActionResetBoth{}
						if oThreatException.Action.ResetBoth.Misc != nil {
							entry.Misc["ThreatExceptionActionResetBoth"] = oThreatException.Action.ResetBoth.Misc
						}
					}
					if oThreatException.Action.ResetClient != nil {
						nestedThreatException.Action.ResetClient = &ThreatExceptionActionResetClient{}
						if oThreatException.Action.ResetClient.Misc != nil {
							entry.Misc["ThreatExceptionActionResetClient"] = oThreatException.Action.ResetClient.Misc
						}
					}
					if oThreatException.Action.ResetServer != nil {
						nestedThreatException.Action.ResetServer = &ThreatExceptionActionResetServer{}
						if oThreatException.Action.ResetServer.Misc != nil {
							entry.Misc["ThreatExceptionActionResetServer"] = oThreatException.Action.ResetServer.Misc
						}
					}
				}
				if oThreatException.ExemptIp != nil {
					nestedThreatException.ExemptIp = []ThreatExceptionExemptIp{}
					for _, oThreatExceptionExemptIp := range oThreatException.ExemptIp {
						nestedThreatExceptionExemptIp := ThreatExceptionExemptIp{}
						if oThreatExceptionExemptIp.Misc != nil {
							entry.Misc["ThreatExceptionExemptIp"] = oThreatExceptionExemptIp.Misc
						}
						if oThreatExceptionExemptIp.Name != "" {
							nestedThreatExceptionExemptIp.Name = oThreatExceptionExemptIp.Name
						}
						nestedThreatException.ExemptIp = append(nestedThreatException.ExemptIp, nestedThreatExceptionExemptIp)
					}
				}
				if oThreatException.Name != "" {
					nestedThreatException.Name = oThreatException.Name
				}
				if oThreatException.PacketCapture != nil {
					nestedThreatException.PacketCapture = oThreatException.PacketCapture
				}
				nestedThreatExceptionCol = append(nestedThreatExceptionCol, nestedThreatException)
			}
			entry.ThreatException = nestedThreatExceptionCol
		}

		entry.Misc["Entry"] = o.Misc

		entryList = append(entryList, entry)
	}

	return entryList, nil
}
func (c *entryXmlContainer_11_0_2) Normalize() ([]*Entry, error) {
	entryList := make([]*Entry, 0, len(c.Answer))
	for _, o := range c.Answer {
		entry := &Entry{
			Misc: make(map[string][]generic.Xml),
		}
		entry.Name = o.Name
		var nestedBotnetDomains *BotnetDomains
		if o.BotnetDomains != nil {
			nestedBotnetDomains = &BotnetDomains{}
			if o.BotnetDomains.Misc != nil {
				entry.Misc["BotnetDomains"] = o.BotnetDomains.Misc
			}
			if o.BotnetDomains.DnsSecurityCategories != nil {
				nestedBotnetDomains.DnsSecurityCategories = []BotnetDomainsDnsSecurityCategories{}
				for _, oBotnetDomainsDnsSecurityCategories := range o.BotnetDomains.DnsSecurityCategories {
					nestedBotnetDomainsDnsSecurityCategories := BotnetDomainsDnsSecurityCategories{}
					if oBotnetDomainsDnsSecurityCategories.Misc != nil {
						entry.Misc["BotnetDomainsDnsSecurityCategories"] = oBotnetDomainsDnsSecurityCategories.Misc
					}
					if oBotnetDomainsDnsSecurityCategories.Action != nil {
						nestedBotnetDomainsDnsSecurityCategories.Action = oBotnetDomainsDnsSecurityCategories.Action
					}
					if oBotnetDomainsDnsSecurityCategories.LogLevel != nil {
						nestedBotnetDomainsDnsSecurityCategories.LogLevel = oBotnetDomainsDnsSecurityCategories.LogLevel
					}
					if oBotnetDomainsDnsSecurityCategories.PacketCapture != nil {
						nestedBotnetDomainsDnsSecurityCategories.PacketCapture = oBotnetDomainsDnsSecurityCategories.PacketCapture
					}
					if oBotnetDomainsDnsSecurityCategories.Name != "" {
						nestedBotnetDomainsDnsSecurityCategories.Name = oBotnetDomainsDnsSecurityCategories.Name
					}
					nestedBotnetDomains.DnsSecurityCategories = append(nestedBotnetDomains.DnsSecurityCategories, nestedBotnetDomainsDnsSecurityCategories)
				}
			}
			if o.BotnetDomains.Lists != nil {
				nestedBotnetDomains.Lists = []BotnetDomainsLists{}
				for _, oBotnetDomainsLists := range o.BotnetDomains.Lists {
					nestedBotnetDomainsLists := BotnetDomainsLists{}
					if oBotnetDomainsLists.Misc != nil {
						entry.Misc["BotnetDomainsLists"] = oBotnetDomainsLists.Misc
					}
					if oBotnetDomainsLists.Action != nil {
						nestedBotnetDomainsLists.Action = &BotnetDomainsListsAction{}
						if oBotnetDomainsLists.Action.Misc != nil {
							entry.Misc["BotnetDomainsListsAction"] = oBotnetDomainsLists.Action.Misc
						}
						if oBotnetDomainsLists.Action.Sinkhole != nil {
							nestedBotnetDomainsLists.Action.Sinkhole = &BotnetDomainsListsActionSinkhole{}
							if oBotnetDomainsLists.Action.Sinkhole.Misc != nil {
								entry.Misc["BotnetDomainsListsActionSinkhole"] = oBotnetDomainsLists.Action.Sinkhole.Misc
							}
						}
						if oBotnetDomainsLists.Action.Alert != nil {
							nestedBotnetDomainsLists.Action.Alert = &BotnetDomainsListsActionAlert{}
							if oBotnetDomainsLists.Action.Alert.Misc != nil {
								entry.Misc["BotnetDomainsListsActionAlert"] = oBotnetDomainsLists.Action.Alert.Misc
							}
						}
						if oBotnetDomainsLists.Action.Allow != nil {
							nestedBotnetDomainsLists.Action.Allow = &BotnetDomainsListsActionAllow{}
							if oBotnetDomainsLists.Action.Allow.Misc != nil {
								entry.Misc["BotnetDomainsListsActionAllow"] = oBotnetDomainsLists.Action.Allow.Misc
							}
						}
						if oBotnetDomainsLists.Action.Block != nil {
							nestedBotnetDomainsLists.Action.Block = &BotnetDomainsListsActionBlock{}
							if oBotnetDomainsLists.Action.Block.Misc != nil {
								entry.Misc["BotnetDomainsListsActionBlock"] = oBotnetDomainsLists.Action.Block.Misc
							}
						}
					}
					if oBotnetDomainsLists.PacketCapture != nil {
						nestedBotnetDomainsLists.PacketCapture = oBotnetDomainsLists.PacketCapture
					}
					if oBotnetDomainsLists.Name != "" {
						nestedBotnetDomainsLists.Name = oBotnetDomainsLists.Name
					}
					nestedBotnetDomains.Lists = append(nestedBotnetDomains.Lists, nestedBotnetDomainsLists)
				}
			}
			if o.BotnetDomains.Sinkhole != nil {
				nestedBotnetDomains.Sinkhole = &BotnetDomainsSinkhole{}
				if o.BotnetDomains.Sinkhole.Misc != nil {
					entry.Misc["BotnetDomainsSinkhole"] = o.BotnetDomains.Sinkhole.Misc
				}
				if o.BotnetDomains.Sinkhole.Ipv4Address != nil {
					nestedBotnetDomains.Sinkhole.Ipv4Address = o.BotnetDomains.Sinkhole.Ipv4Address
				}
				if o.BotnetDomains.Sinkhole.Ipv6Address != nil {
					nestedBotnetDomains.Sinkhole.Ipv6Address = o.BotnetDomains.Sinkhole.Ipv6Address
				}
			}
			if o.BotnetDomains.ThreatException != nil {
				nestedBotnetDomains.ThreatException = []BotnetDomainsThreatException{}
				for _, oBotnetDomainsThreatException := range o.BotnetDomains.ThreatException {
					nestedBotnetDomainsThreatException := BotnetDomainsThreatException{}
					if oBotnetDomainsThreatException.Misc != nil {
						entry.Misc["BotnetDomainsThreatException"] = oBotnetDomainsThreatException.Misc
					}
					if oBotnetDomainsThreatException.Name != "" {
						nestedBotnetDomainsThreatException.Name = oBotnetDomainsThreatException.Name
					}
					nestedBotnetDomains.ThreatException = append(nestedBotnetDomains.ThreatException, nestedBotnetDomainsThreatException)
				}
			}
			if o.BotnetDomains.Whitelist != nil {
				nestedBotnetDomains.Whitelist = []BotnetDomainsWhitelist{}
				for _, oBotnetDomainsWhitelist := range o.BotnetDomains.Whitelist {
					nestedBotnetDomainsWhitelist := BotnetDomainsWhitelist{}
					if oBotnetDomainsWhitelist.Misc != nil {
						entry.Misc["BotnetDomainsWhitelist"] = oBotnetDomainsWhitelist.Misc
					}
					if oBotnetDomainsWhitelist.Description != nil {
						nestedBotnetDomainsWhitelist.Description = oBotnetDomainsWhitelist.Description
					}
					if oBotnetDomainsWhitelist.Name != "" {
						nestedBotnetDomainsWhitelist.Name = oBotnetDomainsWhitelist.Name
					}
					nestedBotnetDomains.Whitelist = append(nestedBotnetDomains.Whitelist, nestedBotnetDomainsWhitelist)
				}
			}
			if o.BotnetDomains.RtypeAction != nil {
				nestedBotnetDomains.RtypeAction = &BotnetDomainsRtypeAction{}
				if o.BotnetDomains.RtypeAction.Misc != nil {
					entry.Misc["BotnetDomainsRtypeAction"] = o.BotnetDomains.RtypeAction.Misc
				}
				if o.BotnetDomains.RtypeAction.Https != nil {
					nestedBotnetDomains.RtypeAction.Https = o.BotnetDomains.RtypeAction.Https
				}
				if o.BotnetDomains.RtypeAction.Svcb != nil {
					nestedBotnetDomains.RtypeAction.Svcb = o.BotnetDomains.RtypeAction.Svcb
				}
				if o.BotnetDomains.RtypeAction.Any != nil {
					nestedBotnetDomains.RtypeAction.Any = o.BotnetDomains.RtypeAction.Any
				}
			}
		}
		entry.BotnetDomains = nestedBotnetDomains

		entry.CloudInlineAnalysis = util.AsBool(o.CloudInlineAnalysis, nil)
		entry.Description = o.Description
		entry.DisableOverride = o.DisableOverride
		entry.InlineExceptionEdlUrl = util.MemToStr(o.InlineExceptionEdlUrl)
		entry.InlineExceptionIpAddress = util.MemToStr(o.InlineExceptionIpAddress)
		var nestedMicaEngineSpywareEnabledCol []MicaEngineSpywareEnabled
		if o.MicaEngineSpywareEnabled != nil {
			nestedMicaEngineSpywareEnabledCol = []MicaEngineSpywareEnabled{}
			for _, oMicaEngineSpywareEnabled := range o.MicaEngineSpywareEnabled {
				nestedMicaEngineSpywareEnabled := MicaEngineSpywareEnabled{}
				if oMicaEngineSpywareEnabled.Misc != nil {
					entry.Misc["MicaEngineSpywareEnabled"] = oMicaEngineSpywareEnabled.Misc
				}
				if oMicaEngineSpywareEnabled.Name != "" {
					nestedMicaEngineSpywareEnabled.Name = oMicaEngineSpywareEnabled.Name
				}
				if oMicaEngineSpywareEnabled.InlinePolicyAction != nil {
					nestedMicaEngineSpywareEnabled.InlinePolicyAction = oMicaEngineSpywareEnabled.InlinePolicyAction
				}
				nestedMicaEngineSpywareEnabledCol = append(nestedMicaEngineSpywareEnabledCol, nestedMicaEngineSpywareEnabled)
			}
			entry.MicaEngineSpywareEnabled = nestedMicaEngineSpywareEnabledCol
		}

		var nestedRulesCol []Rules
		if o.Rules != nil {
			nestedRulesCol = []Rules{}
			for _, oRules := range o.Rules {
				nestedRules := Rules{}
				if oRules.Misc != nil {
					entry.Misc["Rules"] = oRules.Misc
				}
				if oRules.ThreatName != nil {
					nestedRules.ThreatName = oRules.ThreatName
				}
				if oRules.Category != nil {
					nestedRules.Category = oRules.Category
				}
				if oRules.PacketCapture != nil {
					nestedRules.PacketCapture = oRules.PacketCapture
				}
				if oRules.Severity != nil {
					nestedRules.Severity = util.MemToStr(oRules.Severity)
				}
				if oRules.Action != nil {
					nestedRules.Action = &RulesAction{}
					if oRules.Action.Misc != nil {
						entry.Misc["RulesAction"] = oRules.Action.Misc
					}
					if oRules.Action.Default != nil {
						nestedRules.Action.Default = &RulesActionDefault{}
						if oRules.Action.Default.Misc != nil {
							entry.Misc["RulesActionDefault"] = oRules.Action.Default.Misc
						}
					}
					if oRules.Action.Allow != nil {
						nestedRules.Action.Allow = &RulesActionAllow{}
						if oRules.Action.Allow.Misc != nil {
							entry.Misc["RulesActionAllow"] = oRules.Action.Allow.Misc
						}
					}
					if oRules.Action.Alert != nil {
						nestedRules.Action.Alert = &RulesActionAlert{}
						if oRules.Action.Alert.Misc != nil {
							entry.Misc["RulesActionAlert"] = oRules.Action.Alert.Misc
						}
					}
					if oRules.Action.Drop != nil {
						nestedRules.Action.Drop = &RulesActionDrop{}
						if oRules.Action.Drop.Misc != nil {
							entry.Misc["RulesActionDrop"] = oRules.Action.Drop.Misc
						}
					}
					if oRules.Action.ResetClient != nil {
						nestedRules.Action.ResetClient = &RulesActionResetClient{}
						if oRules.Action.ResetClient.Misc != nil {
							entry.Misc["RulesActionResetClient"] = oRules.Action.ResetClient.Misc
						}
					}
					if oRules.Action.ResetServer != nil {
						nestedRules.Action.ResetServer = &RulesActionResetServer{}
						if oRules.Action.ResetServer.Misc != nil {
							entry.Misc["RulesActionResetServer"] = oRules.Action.ResetServer.Misc
						}
					}
					if oRules.Action.ResetBoth != nil {
						nestedRules.Action.ResetBoth = &RulesActionResetBoth{}
						if oRules.Action.ResetBoth.Misc != nil {
							entry.Misc["RulesActionResetBoth"] = oRules.Action.ResetBoth.Misc
						}
					}
					if oRules.Action.BlockIp != nil {
						nestedRules.Action.BlockIp = &RulesActionBlockIp{}
						if oRules.Action.BlockIp.Misc != nil {
							entry.Misc["RulesActionBlockIp"] = oRules.Action.BlockIp.Misc
						}
						if oRules.Action.BlockIp.Duration != nil {
							nestedRules.Action.BlockIp.Duration = oRules.Action.BlockIp.Duration
						}
						if oRules.Action.BlockIp.TrackBy != nil {
							nestedRules.Action.BlockIp.TrackBy = oRules.Action.BlockIp.TrackBy
						}
					}
				}
				if oRules.Name != "" {
					nestedRules.Name = oRules.Name
				}
				nestedRulesCol = append(nestedRulesCol, nestedRules)
			}
			entry.Rules = nestedRulesCol
		}

		var nestedThreatExceptionCol []ThreatException
		if o.ThreatException != nil {
			nestedThreatExceptionCol = []ThreatException{}
			for _, oThreatException := range o.ThreatException {
				nestedThreatException := ThreatException{}
				if oThreatException.Misc != nil {
					entry.Misc["ThreatException"] = oThreatException.Misc
				}
				if oThreatException.ExemptIp != nil {
					nestedThreatException.ExemptIp = []ThreatExceptionExemptIp{}
					for _, oThreatExceptionExemptIp := range oThreatException.ExemptIp {
						nestedThreatExceptionExemptIp := ThreatExceptionExemptIp{}
						if oThreatExceptionExemptIp.Misc != nil {
							entry.Misc["ThreatExceptionExemptIp"] = oThreatExceptionExemptIp.Misc
						}
						if oThreatExceptionExemptIp.Name != "" {
							nestedThreatExceptionExemptIp.Name = oThreatExceptionExemptIp.Name
						}
						nestedThreatException.ExemptIp = append(nestedThreatException.ExemptIp, nestedThreatExceptionExemptIp)
					}
				}
				if oThreatException.Name != "" {
					nestedThreatException.Name = oThreatException.Name
				}
				if oThreatException.PacketCapture != nil {
					nestedThreatException.PacketCapture = oThreatException.PacketCapture
				}
				if oThreatException.Action != nil {
					nestedThreatException.Action = &ThreatExceptionAction{}
					if oThreatException.Action.Misc != nil {
						entry.Misc["ThreatExceptionAction"] = oThreatException.Action.Misc
					}
					if oThreatException.Action.Drop != nil {
						nestedThreatException.Action.Drop = &ThreatExceptionActionDrop{}
						if oThreatException.Action.Drop.Misc != nil {
							entry.Misc["ThreatExceptionActionDrop"] = oThreatException.Action.Drop.Misc
						}
					}
					if oThreatException.Action.ResetBoth != nil {
						nestedThreatException.Action.ResetBoth = &ThreatExceptionActionResetBoth{}
						if oThreatException.Action.ResetBoth.Misc != nil {
							entry.Misc["ThreatExceptionActionResetBoth"] = oThreatException.Action.ResetBoth.Misc
						}
					}
					if oThreatException.Action.ResetClient != nil {
						nestedThreatException.Action.ResetClient = &ThreatExceptionActionResetClient{}
						if oThreatException.Action.ResetClient.Misc != nil {
							entry.Misc["ThreatExceptionActionResetClient"] = oThreatException.Action.ResetClient.Misc
						}
					}
					if oThreatException.Action.ResetServer != nil {
						nestedThreatException.Action.ResetServer = &ThreatExceptionActionResetServer{}
						if oThreatException.Action.ResetServer.Misc != nil {
							entry.Misc["ThreatExceptionActionResetServer"] = oThreatException.Action.ResetServer.Misc
						}
					}
					if oThreatException.Action.BlockIp != nil {
						nestedThreatException.Action.BlockIp = &ThreatExceptionActionBlockIp{}
						if oThreatException.Action.BlockIp.Misc != nil {
							entry.Misc["ThreatExceptionActionBlockIp"] = oThreatException.Action.BlockIp.Misc
						}
						if oThreatException.Action.BlockIp.TrackBy != nil {
							nestedThreatException.Action.BlockIp.TrackBy = oThreatException.Action.BlockIp.TrackBy
						}
						if oThreatException.Action.BlockIp.Duration != nil {
							nestedThreatException.Action.BlockIp.Duration = oThreatException.Action.BlockIp.Duration
						}
					}
					if oThreatException.Action.Default != nil {
						nestedThreatException.Action.Default = &ThreatExceptionActionDefault{}
						if oThreatException.Action.Default.Misc != nil {
							entry.Misc["ThreatExceptionActionDefault"] = oThreatException.Action.Default.Misc
						}
					}
					if oThreatException.Action.Allow != nil {
						nestedThreatException.Action.Allow = &ThreatExceptionActionAllow{}
						if oThreatException.Action.Allow.Misc != nil {
							entry.Misc["ThreatExceptionActionAllow"] = oThreatException.Action.Allow.Misc
						}
					}
					if oThreatException.Action.Alert != nil {
						nestedThreatException.Action.Alert = &ThreatExceptionActionAlert{}
						if oThreatException.Action.Alert.Misc != nil {
							entry.Misc["ThreatExceptionActionAlert"] = oThreatException.Action.Alert.Misc
						}
					}
				}
				nestedThreatExceptionCol = append(nestedThreatExceptionCol, nestedThreatException)
			}
			entry.ThreatException = nestedThreatExceptionCol
		}

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
	if !matchBotnetDomains(a.BotnetDomains, b.BotnetDomains) {
		return false
	}
	if !util.BoolsMatch(a.CloudInlineAnalysis, b.CloudInlineAnalysis) {
		return false
	}
	if !util.StringsMatch(a.Description, b.Description) {
		return false
	}
	if !util.StringsMatch(a.DisableOverride, b.DisableOverride) {
		return false
	}
	if !util.OrderedListsMatch(a.InlineExceptionEdlUrl, b.InlineExceptionEdlUrl) {
		return false
	}
	if !util.OrderedListsMatch(a.InlineExceptionIpAddress, b.InlineExceptionIpAddress) {
		return false
	}
	if !matchMicaEngineSpywareEnabled(a.MicaEngineSpywareEnabled, b.MicaEngineSpywareEnabled) {
		return false
	}
	if !matchRules(a.Rules, b.Rules) {
		return false
	}
	if !matchThreatException(a.ThreatException, b.ThreatException) {
		return false
	}

	return true
}

func matchBotnetDomainsDnsSecurityCategories(a []BotnetDomainsDnsSecurityCategories, b []BotnetDomainsDnsSecurityCategories) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Action, b.Action) {
				return false
			}
			if !util.StringsMatch(a.LogLevel, b.LogLevel) {
				return false
			}
			if !util.StringsMatch(a.PacketCapture, b.PacketCapture) {
				return false
			}
		}
	}
	return true
}
func matchBotnetDomainsListsActionAlert(a *BotnetDomainsListsActionAlert, b *BotnetDomainsListsActionAlert) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchBotnetDomainsListsActionAllow(a *BotnetDomainsListsActionAllow, b *BotnetDomainsListsActionAllow) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchBotnetDomainsListsActionBlock(a *BotnetDomainsListsActionBlock, b *BotnetDomainsListsActionBlock) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchBotnetDomainsListsActionSinkhole(a *BotnetDomainsListsActionSinkhole, b *BotnetDomainsListsActionSinkhole) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchBotnetDomainsListsAction(a *BotnetDomainsListsAction, b *BotnetDomainsListsAction) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchBotnetDomainsListsActionSinkhole(a.Sinkhole, b.Sinkhole) {
		return false
	}
	if !matchBotnetDomainsListsActionAlert(a.Alert, b.Alert) {
		return false
	}
	if !matchBotnetDomainsListsActionAllow(a.Allow, b.Allow) {
		return false
	}
	if !matchBotnetDomainsListsActionBlock(a.Block, b.Block) {
		return false
	}
	return true
}
func matchBotnetDomainsLists(a []BotnetDomainsLists, b []BotnetDomainsLists) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !matchBotnetDomainsListsAction(a.Action, b.Action) {
				return false
			}
			if !util.StringsMatch(a.PacketCapture, b.PacketCapture) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchBotnetDomainsSinkhole(a *BotnetDomainsSinkhole, b *BotnetDomainsSinkhole) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Ipv4Address, b.Ipv4Address) {
		return false
	}
	if !util.StringsMatch(a.Ipv6Address, b.Ipv6Address) {
		return false
	}
	return true
}
func matchBotnetDomainsThreatException(a []BotnetDomainsThreatException, b []BotnetDomainsThreatException) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchBotnetDomainsWhitelist(a []BotnetDomainsWhitelist, b []BotnetDomainsWhitelist) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Description, b.Description) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchBotnetDomainsRtypeAction(a *BotnetDomainsRtypeAction, b *BotnetDomainsRtypeAction) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Any, b.Any) {
		return false
	}
	if !util.StringsMatch(a.Https, b.Https) {
		return false
	}
	if !util.StringsMatch(a.Svcb, b.Svcb) {
		return false
	}
	return true
}
func matchBotnetDomains(a *BotnetDomains, b *BotnetDomains) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchBotnetDomainsSinkhole(a.Sinkhole, b.Sinkhole) {
		return false
	}
	if !matchBotnetDomainsThreatException(a.ThreatException, b.ThreatException) {
		return false
	}
	if !matchBotnetDomainsWhitelist(a.Whitelist, b.Whitelist) {
		return false
	}
	if !matchBotnetDomainsRtypeAction(a.RtypeAction, b.RtypeAction) {
		return false
	}
	if !matchBotnetDomainsDnsSecurityCategories(a.DnsSecurityCategories, b.DnsSecurityCategories) {
		return false
	}
	if !matchBotnetDomainsLists(a.Lists, b.Lists) {
		return false
	}
	return true
}
func matchThreatExceptionExemptIp(a []ThreatExceptionExemptIp, b []ThreatExceptionExemptIp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchThreatExceptionActionDefault(a *ThreatExceptionActionDefault, b *ThreatExceptionActionDefault) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchThreatExceptionActionAllow(a *ThreatExceptionActionAllow, b *ThreatExceptionActionAllow) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchThreatExceptionActionAlert(a *ThreatExceptionActionAlert, b *ThreatExceptionActionAlert) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchThreatExceptionActionDrop(a *ThreatExceptionActionDrop, b *ThreatExceptionActionDrop) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchThreatExceptionActionResetBoth(a *ThreatExceptionActionResetBoth, b *ThreatExceptionActionResetBoth) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchThreatExceptionActionResetClient(a *ThreatExceptionActionResetClient, b *ThreatExceptionActionResetClient) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchThreatExceptionActionResetServer(a *ThreatExceptionActionResetServer, b *ThreatExceptionActionResetServer) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchThreatExceptionActionBlockIp(a *ThreatExceptionActionBlockIp, b *ThreatExceptionActionBlockIp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.TrackBy, b.TrackBy) {
		return false
	}
	if !util.Ints64Match(a.Duration, b.Duration) {
		return false
	}
	return true
}
func matchThreatExceptionAction(a *ThreatExceptionAction, b *ThreatExceptionAction) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchThreatExceptionActionDefault(a.Default, b.Default) {
		return false
	}
	if !matchThreatExceptionActionAllow(a.Allow, b.Allow) {
		return false
	}
	if !matchThreatExceptionActionAlert(a.Alert, b.Alert) {
		return false
	}
	if !matchThreatExceptionActionDrop(a.Drop, b.Drop) {
		return false
	}
	if !matchThreatExceptionActionResetBoth(a.ResetBoth, b.ResetBoth) {
		return false
	}
	if !matchThreatExceptionActionResetClient(a.ResetClient, b.ResetClient) {
		return false
	}
	if !matchThreatExceptionActionResetServer(a.ResetServer, b.ResetServer) {
		return false
	}
	if !matchThreatExceptionActionBlockIp(a.BlockIp, b.BlockIp) {
		return false
	}
	return true
}
func matchThreatException(a []ThreatException, b []ThreatException) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.PacketCapture, b.PacketCapture) {
				return false
			}
			if !matchThreatExceptionAction(a.Action, b.Action) {
				return false
			}
			if !matchThreatExceptionExemptIp(a.ExemptIp, b.ExemptIp) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchMicaEngineSpywareEnabled(a []MicaEngineSpywareEnabled, b []MicaEngineSpywareEnabled) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.InlinePolicyAction, b.InlinePolicyAction) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchRulesActionBlockIp(a *RulesActionBlockIp, b *RulesActionBlockIp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.TrackBy, b.TrackBy) {
		return false
	}
	if !util.Ints64Match(a.Duration, b.Duration) {
		return false
	}
	return true
}
func matchRulesActionDefault(a *RulesActionDefault, b *RulesActionDefault) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRulesActionAllow(a *RulesActionAllow, b *RulesActionAllow) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRulesActionAlert(a *RulesActionAlert, b *RulesActionAlert) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRulesActionDrop(a *RulesActionDrop, b *RulesActionDrop) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRulesActionResetClient(a *RulesActionResetClient, b *RulesActionResetClient) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRulesActionResetServer(a *RulesActionResetServer, b *RulesActionResetServer) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRulesActionResetBoth(a *RulesActionResetBoth, b *RulesActionResetBoth) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchRulesAction(a *RulesAction, b *RulesAction) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRulesActionDrop(a.Drop, b.Drop) {
		return false
	}
	if !matchRulesActionResetClient(a.ResetClient, b.ResetClient) {
		return false
	}
	if !matchRulesActionResetServer(a.ResetServer, b.ResetServer) {
		return false
	}
	if !matchRulesActionResetBoth(a.ResetBoth, b.ResetBoth) {
		return false
	}
	if !matchRulesActionBlockIp(a.BlockIp, b.BlockIp) {
		return false
	}
	if !matchRulesActionDefault(a.Default, b.Default) {
		return false
	}
	if !matchRulesActionAllow(a.Allow, b.Allow) {
		return false
	}
	if !matchRulesActionAlert(a.Alert, b.Alert) {
		return false
	}
	return true
}
func matchRules(a []Rules, b []Rules) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Category, b.Category) {
				return false
			}
			if !util.StringsMatch(a.PacketCapture, b.PacketCapture) {
				return false
			}
			if !util.OrderedListsMatch(a.Severity, b.Severity) {
				return false
			}
			if !matchRulesAction(a.Action, b.Action) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.ThreatName, b.ThreatName) {
				return false
			}
		}
	}
	return true
}

func (o *Entry) EntryName() string {
	return o.Name
}

func (o *Entry) SetEntryName(name string) {
	o.Name = name
}
