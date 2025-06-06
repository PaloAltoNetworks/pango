package spyware

import (
	"encoding/xml"
	"fmt"

	"github.com/PaloAltoNetworks/pango/v2/filtering"
	"github.com/PaloAltoNetworks/pango/v2/generic"
	"github.com/PaloAltoNetworks/pango/v2/util"
	"github.com/PaloAltoNetworks/pango/v2/version"
)

var (
	_ filtering.Fielder = &Entry{}
)

var (
	suffix = []string{"profiles", "spyware", "$name"}
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
	Misc                     []generic.Xml
}
type BotnetDomains struct {
	DnsSecurityCategories []BotnetDomainsDnsSecurityCategories
	Lists                 []BotnetDomainsLists
	RtypeAction           *BotnetDomainsRtypeAction
	Sinkhole              *BotnetDomainsSinkhole
	ThreatException       []BotnetDomainsThreatException
	Whitelist             []BotnetDomainsWhitelist
	Misc                  []generic.Xml
}
type BotnetDomainsDnsSecurityCategories struct {
	Name          string
	Action        *string
	LogLevel      *string
	PacketCapture *string
	Misc          []generic.Xml
}
type BotnetDomainsLists struct {
	Name          string
	Action        *BotnetDomainsListsAction
	PacketCapture *string
	Misc          []generic.Xml
}
type BotnetDomainsListsAction struct {
	Alert    *BotnetDomainsListsActionAlert
	Allow    *BotnetDomainsListsActionAllow
	Block    *BotnetDomainsListsActionBlock
	Sinkhole *BotnetDomainsListsActionSinkhole
	Misc     []generic.Xml
}
type BotnetDomainsListsActionAlert struct {
	Misc []generic.Xml
}
type BotnetDomainsListsActionAllow struct {
	Misc []generic.Xml
}
type BotnetDomainsListsActionBlock struct {
	Misc []generic.Xml
}
type BotnetDomainsListsActionSinkhole struct {
	Misc []generic.Xml
}
type BotnetDomainsRtypeAction struct {
	Any   *string
	Https *string
	Svcb  *string
	Misc  []generic.Xml
}
type BotnetDomainsSinkhole struct {
	Ipv4Address *string
	Ipv6Address *string
	Misc        []generic.Xml
}
type BotnetDomainsThreatException struct {
	Name string
	Misc []generic.Xml
}
type BotnetDomainsWhitelist struct {
	Name        string
	Description *string
	Misc        []generic.Xml
}
type MicaEngineSpywareEnabled struct {
	Name               string
	InlinePolicyAction *string
	Misc               []generic.Xml
}
type Rules struct {
	Name          string
	ThreatName    *string
	Category      *string
	PacketCapture *string
	Severity      []string
	Action        *RulesAction
	Misc          []generic.Xml
}
type RulesAction struct {
	Default     *RulesActionDefault
	Allow       *RulesActionAllow
	Alert       *RulesActionAlert
	Drop        *RulesActionDrop
	ResetClient *RulesActionResetClient
	ResetServer *RulesActionResetServer
	ResetBoth   *RulesActionResetBoth
	BlockIp     *RulesActionBlockIp
	Misc        []generic.Xml
}
type RulesActionDefault struct {
	Misc []generic.Xml
}
type RulesActionAllow struct {
	Misc []generic.Xml
}
type RulesActionAlert struct {
	Misc []generic.Xml
}
type RulesActionDrop struct {
	Misc []generic.Xml
}
type RulesActionResetClient struct {
	Misc []generic.Xml
}
type RulesActionResetServer struct {
	Misc []generic.Xml
}
type RulesActionResetBoth struct {
	Misc []generic.Xml
}
type RulesActionBlockIp struct {
	TrackBy  *string
	Duration *int64
	Misc     []generic.Xml
}
type ThreatException struct {
	Name          string
	PacketCapture *string
	Action        *ThreatExceptionAction
	ExemptIp      []ThreatExceptionExemptIp
	Misc          []generic.Xml
}
type ThreatExceptionAction struct {
	Default     *ThreatExceptionActionDefault
	Allow       *ThreatExceptionActionAllow
	Alert       *ThreatExceptionActionAlert
	Drop        *ThreatExceptionActionDrop
	ResetBoth   *ThreatExceptionActionResetBoth
	ResetClient *ThreatExceptionActionResetClient
	ResetServer *ThreatExceptionActionResetServer
	BlockIp     *ThreatExceptionActionBlockIp
	Misc        []generic.Xml
}
type ThreatExceptionActionDefault struct {
	Misc []generic.Xml
}
type ThreatExceptionActionAllow struct {
	Misc []generic.Xml
}
type ThreatExceptionActionAlert struct {
	Misc []generic.Xml
}
type ThreatExceptionActionDrop struct {
	Misc []generic.Xml
}
type ThreatExceptionActionResetBoth struct {
	Misc []generic.Xml
}
type ThreatExceptionActionResetClient struct {
	Misc []generic.Xml
}
type ThreatExceptionActionResetServer struct {
	Misc []generic.Xml
}
type ThreatExceptionActionBlockIp struct {
	TrackBy  *string
	Duration *int64
	Misc     []generic.Xml
}
type ThreatExceptionExemptIp struct {
	Name string
	Misc []generic.Xml
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
	XMLName                  xml.Name                              `xml:"entry"`
	Name                     string                                `xml:"name,attr"`
	BotnetDomains            *botnetDomainsXml                     `xml:"botnet-domains,omitempty"`
	CloudInlineAnalysis      *string                               `xml:"cloud-inline-analysis,omitempty"`
	Description              *string                               `xml:"description,omitempty"`
	DisableOverride          *string                               `xml:"disable-override,omitempty"`
	InlineExceptionEdlUrl    *util.MemberType                      `xml:"inline-exception-edl-url,omitempty"`
	InlineExceptionIpAddress *util.MemberType                      `xml:"inline-exception-ip-address,omitempty"`
	MicaEngineSpywareEnabled *micaEngineSpywareEnabledContainerXml `xml:"mica-engine-spyware-enabled,omitempty"`
	Rules                    *rulesContainerXml                    `xml:"rules,omitempty"`
	ThreatException          *threatExceptionContainerXml          `xml:"threat-exception,omitempty"`
	Misc                     []generic.Xml                         `xml:",any"`
}
type botnetDomainsXml struct {
	DnsSecurityCategories *botnetDomainsDnsSecurityCategoriesContainerXml `xml:"dns-security-categories,omitempty"`
	Lists                 *botnetDomainsListsContainerXml                 `xml:"lists,omitempty"`
	RtypeAction           *botnetDomainsRtypeActionXml                    `xml:"rtype-action,omitempty"`
	Sinkhole              *botnetDomainsSinkholeXml                       `xml:"sinkhole,omitempty"`
	ThreatException       *botnetDomainsThreatExceptionContainerXml       `xml:"threat-exception,omitempty"`
	Whitelist             *botnetDomainsWhitelistContainerXml             `xml:"whitelist,omitempty"`
	Misc                  []generic.Xml                                   `xml:",any"`
}
type botnetDomainsDnsSecurityCategoriesContainerXml struct {
	Entries []botnetDomainsDnsSecurityCategoriesXml `xml:"entry"`
}
type botnetDomainsDnsSecurityCategoriesXml struct {
	XMLName       xml.Name      `xml:"entry"`
	Name          string        `xml:"name,attr"`
	Action        *string       `xml:"action,omitempty"`
	LogLevel      *string       `xml:"log-level,omitempty"`
	PacketCapture *string       `xml:"packet-capture,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type botnetDomainsListsContainerXml struct {
	Entries []botnetDomainsListsXml `xml:"entry"`
}
type botnetDomainsListsXml struct {
	XMLName       xml.Name                     `xml:"entry"`
	Name          string                       `xml:"name,attr"`
	Action        *botnetDomainsListsActionXml `xml:"action,omitempty"`
	PacketCapture *string                      `xml:"packet-capture,omitempty"`
	Misc          []generic.Xml                `xml:",any"`
}
type botnetDomainsListsActionXml struct {
	Alert    *botnetDomainsListsActionAlertXml    `xml:"alert,omitempty"`
	Allow    *botnetDomainsListsActionAllowXml    `xml:"allow,omitempty"`
	Block    *botnetDomainsListsActionBlockXml    `xml:"block,omitempty"`
	Sinkhole *botnetDomainsListsActionSinkholeXml `xml:"sinkhole,omitempty"`
	Misc     []generic.Xml                        `xml:",any"`
}
type botnetDomainsListsActionAlertXml struct {
	Misc []generic.Xml `xml:",any"`
}
type botnetDomainsListsActionAllowXml struct {
	Misc []generic.Xml `xml:",any"`
}
type botnetDomainsListsActionBlockXml struct {
	Misc []generic.Xml `xml:",any"`
}
type botnetDomainsListsActionSinkholeXml struct {
	Misc []generic.Xml `xml:",any"`
}
type botnetDomainsRtypeActionXml struct {
	Any   *string       `xml:"any,omitempty"`
	Https *string       `xml:"https,omitempty"`
	Svcb  *string       `xml:"svcb,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type botnetDomainsSinkholeXml struct {
	Ipv4Address *string       `xml:"ipv4-address,omitempty"`
	Ipv6Address *string       `xml:"ipv6-address,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type botnetDomainsThreatExceptionContainerXml struct {
	Entries []botnetDomainsThreatExceptionXml `xml:"entry"`
}
type botnetDomainsThreatExceptionXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}
type botnetDomainsWhitelistContainerXml struct {
	Entries []botnetDomainsWhitelistXml `xml:"entry"`
}
type botnetDomainsWhitelistXml struct {
	XMLName     xml.Name      `xml:"entry"`
	Name        string        `xml:"name,attr"`
	Description *string       `xml:"description,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type micaEngineSpywareEnabledContainerXml struct {
	Entries []micaEngineSpywareEnabledXml `xml:"entry"`
}
type micaEngineSpywareEnabledXml struct {
	XMLName            xml.Name      `xml:"entry"`
	Name               string        `xml:"name,attr"`
	InlinePolicyAction *string       `xml:"inline-policy-action,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type rulesContainerXml struct {
	Entries []rulesXml `xml:"entry"`
}
type rulesXml struct {
	XMLName       xml.Name         `xml:"entry"`
	Name          string           `xml:"name,attr"`
	ThreatName    *string          `xml:"threat-name,omitempty"`
	Category      *string          `xml:"category,omitempty"`
	PacketCapture *string          `xml:"packet-capture,omitempty"`
	Severity      *util.MemberType `xml:"severity,omitempty"`
	Action        *rulesActionXml  `xml:"action,omitempty"`
	Misc          []generic.Xml    `xml:",any"`
}
type rulesActionXml struct {
	Default     *rulesActionDefaultXml     `xml:"default,omitempty"`
	Allow       *rulesActionAllowXml       `xml:"allow,omitempty"`
	Alert       *rulesActionAlertXml       `xml:"alert,omitempty"`
	Drop        *rulesActionDropXml        `xml:"drop,omitempty"`
	ResetClient *rulesActionResetClientXml `xml:"reset-client,omitempty"`
	ResetServer *rulesActionResetServerXml `xml:"reset-server,omitempty"`
	ResetBoth   *rulesActionResetBothXml   `xml:"reset-both,omitempty"`
	BlockIp     *rulesActionBlockIpXml     `xml:"block-ip,omitempty"`
	Misc        []generic.Xml              `xml:",any"`
}
type rulesActionDefaultXml struct {
	Misc []generic.Xml `xml:",any"`
}
type rulesActionAllowXml struct {
	Misc []generic.Xml `xml:",any"`
}
type rulesActionAlertXml struct {
	Misc []generic.Xml `xml:",any"`
}
type rulesActionDropXml struct {
	Misc []generic.Xml `xml:",any"`
}
type rulesActionResetClientXml struct {
	Misc []generic.Xml `xml:",any"`
}
type rulesActionResetServerXml struct {
	Misc []generic.Xml `xml:",any"`
}
type rulesActionResetBothXml struct {
	Misc []generic.Xml `xml:",any"`
}
type rulesActionBlockIpXml struct {
	TrackBy  *string       `xml:"track-by,omitempty"`
	Duration *int64        `xml:"duration,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type threatExceptionContainerXml struct {
	Entries []threatExceptionXml `xml:"entry"`
}
type threatExceptionXml struct {
	XMLName       xml.Name                             `xml:"entry"`
	Name          string                               `xml:"name,attr"`
	PacketCapture *string                              `xml:"packet-capture,omitempty"`
	Action        *threatExceptionActionXml            `xml:"action,omitempty"`
	ExemptIp      *threatExceptionExemptIpContainerXml `xml:"exempt-ip,omitempty"`
	Misc          []generic.Xml                        `xml:",any"`
}
type threatExceptionActionXml struct {
	Default     *threatExceptionActionDefaultXml     `xml:"default,omitempty"`
	Allow       *threatExceptionActionAllowXml       `xml:"allow,omitempty"`
	Alert       *threatExceptionActionAlertXml       `xml:"alert,omitempty"`
	Drop        *threatExceptionActionDropXml        `xml:"drop,omitempty"`
	ResetBoth   *threatExceptionActionResetBothXml   `xml:"reset-both,omitempty"`
	ResetClient *threatExceptionActionResetClientXml `xml:"reset-client,omitempty"`
	ResetServer *threatExceptionActionResetServerXml `xml:"reset-server,omitempty"`
	BlockIp     *threatExceptionActionBlockIpXml     `xml:"block-ip,omitempty"`
	Misc        []generic.Xml                        `xml:",any"`
}
type threatExceptionActionDefaultXml struct {
	Misc []generic.Xml `xml:",any"`
}
type threatExceptionActionAllowXml struct {
	Misc []generic.Xml `xml:",any"`
}
type threatExceptionActionAlertXml struct {
	Misc []generic.Xml `xml:",any"`
}
type threatExceptionActionDropXml struct {
	Misc []generic.Xml `xml:",any"`
}
type threatExceptionActionResetBothXml struct {
	Misc []generic.Xml `xml:",any"`
}
type threatExceptionActionResetClientXml struct {
	Misc []generic.Xml `xml:",any"`
}
type threatExceptionActionResetServerXml struct {
	Misc []generic.Xml `xml:",any"`
}
type threatExceptionActionBlockIpXml struct {
	TrackBy  *string       `xml:"track-by,omitempty"`
	Duration *int64        `xml:"duration,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type threatExceptionExemptIpContainerXml struct {
	Entries []threatExceptionExemptIpXml `xml:"entry"`
}
type threatExceptionExemptIpXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}
type entryXml_11_0_2 struct {
	XMLName                  xml.Name                                     `xml:"entry"`
	Name                     string                                       `xml:"name,attr"`
	BotnetDomains            *botnetDomainsXml_11_0_2                     `xml:"botnet-domains,omitempty"`
	CloudInlineAnalysis      *string                                      `xml:"cloud-inline-analysis,omitempty"`
	Description              *string                                      `xml:"description,omitempty"`
	DisableOverride          *string                                      `xml:"disable-override,omitempty"`
	InlineExceptionEdlUrl    *util.MemberType                             `xml:"inline-exception-edl-url,omitempty"`
	InlineExceptionIpAddress *util.MemberType                             `xml:"inline-exception-ip-address,omitempty"`
	MicaEngineSpywareEnabled *micaEngineSpywareEnabledContainerXml_11_0_2 `xml:"mica-engine-spyware-enabled,omitempty"`
	Rules                    *rulesContainerXml_11_0_2                    `xml:"rules,omitempty"`
	ThreatException          *threatExceptionContainerXml_11_0_2          `xml:"threat-exception,omitempty"`
	Misc                     []generic.Xml                                `xml:",any"`
}
type botnetDomainsXml_11_0_2 struct {
	DnsSecurityCategories *botnetDomainsDnsSecurityCategoriesContainerXml_11_0_2 `xml:"dns-security-categories,omitempty"`
	Lists                 *botnetDomainsListsContainerXml_11_0_2                 `xml:"lists,omitempty"`
	RtypeAction           *botnetDomainsRtypeActionXml_11_0_2                    `xml:"rtype-action,omitempty"`
	Sinkhole              *botnetDomainsSinkholeXml_11_0_2                       `xml:"sinkhole,omitempty"`
	ThreatException       *botnetDomainsThreatExceptionContainerXml_11_0_2       `xml:"threat-exception,omitempty"`
	Whitelist             *botnetDomainsWhitelistContainerXml_11_0_2             `xml:"whitelist,omitempty"`
	Misc                  []generic.Xml                                          `xml:",any"`
}
type botnetDomainsDnsSecurityCategoriesContainerXml_11_0_2 struct {
	Entries []botnetDomainsDnsSecurityCategoriesXml_11_0_2 `xml:"entry"`
}
type botnetDomainsDnsSecurityCategoriesXml_11_0_2 struct {
	XMLName       xml.Name      `xml:"entry"`
	Name          string        `xml:"name,attr"`
	Action        *string       `xml:"action,omitempty"`
	LogLevel      *string       `xml:"log-level,omitempty"`
	PacketCapture *string       `xml:"packet-capture,omitempty"`
	Misc          []generic.Xml `xml:",any"`
}
type botnetDomainsListsContainerXml_11_0_2 struct {
	Entries []botnetDomainsListsXml_11_0_2 `xml:"entry"`
}
type botnetDomainsListsXml_11_0_2 struct {
	XMLName       xml.Name                            `xml:"entry"`
	Name          string                              `xml:"name,attr"`
	Action        *botnetDomainsListsActionXml_11_0_2 `xml:"action,omitempty"`
	PacketCapture *string                             `xml:"packet-capture,omitempty"`
	Misc          []generic.Xml                       `xml:",any"`
}
type botnetDomainsListsActionXml_11_0_2 struct {
	Alert    *botnetDomainsListsActionAlertXml_11_0_2    `xml:"alert,omitempty"`
	Allow    *botnetDomainsListsActionAllowXml_11_0_2    `xml:"allow,omitempty"`
	Block    *botnetDomainsListsActionBlockXml_11_0_2    `xml:"block,omitempty"`
	Sinkhole *botnetDomainsListsActionSinkholeXml_11_0_2 `xml:"sinkhole,omitempty"`
	Misc     []generic.Xml                               `xml:",any"`
}
type botnetDomainsListsActionAlertXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type botnetDomainsListsActionAllowXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type botnetDomainsListsActionBlockXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type botnetDomainsListsActionSinkholeXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type botnetDomainsRtypeActionXml_11_0_2 struct {
	Any   *string       `xml:"any,omitempty"`
	Https *string       `xml:"https,omitempty"`
	Svcb  *string       `xml:"svcb,omitempty"`
	Misc  []generic.Xml `xml:",any"`
}
type botnetDomainsSinkholeXml_11_0_2 struct {
	Ipv4Address *string       `xml:"ipv4-address,omitempty"`
	Ipv6Address *string       `xml:"ipv6-address,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type botnetDomainsThreatExceptionContainerXml_11_0_2 struct {
	Entries []botnetDomainsThreatExceptionXml_11_0_2 `xml:"entry"`
}
type botnetDomainsThreatExceptionXml_11_0_2 struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}
type botnetDomainsWhitelistContainerXml_11_0_2 struct {
	Entries []botnetDomainsWhitelistXml_11_0_2 `xml:"entry"`
}
type botnetDomainsWhitelistXml_11_0_2 struct {
	XMLName     xml.Name      `xml:"entry"`
	Name        string        `xml:"name,attr"`
	Description *string       `xml:"description,omitempty"`
	Misc        []generic.Xml `xml:",any"`
}
type micaEngineSpywareEnabledContainerXml_11_0_2 struct {
	Entries []micaEngineSpywareEnabledXml_11_0_2 `xml:"entry"`
}
type micaEngineSpywareEnabledXml_11_0_2 struct {
	XMLName            xml.Name      `xml:"entry"`
	Name               string        `xml:"name,attr"`
	InlinePolicyAction *string       `xml:"inline-policy-action,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type rulesContainerXml_11_0_2 struct {
	Entries []rulesXml_11_0_2 `xml:"entry"`
}
type rulesXml_11_0_2 struct {
	XMLName       xml.Name               `xml:"entry"`
	Name          string                 `xml:"name,attr"`
	ThreatName    *string                `xml:"threat-name,omitempty"`
	Category      *string                `xml:"category,omitempty"`
	PacketCapture *string                `xml:"packet-capture,omitempty"`
	Severity      *util.MemberType       `xml:"severity,omitempty"`
	Action        *rulesActionXml_11_0_2 `xml:"action,omitempty"`
	Misc          []generic.Xml          `xml:",any"`
}
type rulesActionXml_11_0_2 struct {
	Default     *rulesActionDefaultXml_11_0_2     `xml:"default,omitempty"`
	Allow       *rulesActionAllowXml_11_0_2       `xml:"allow,omitempty"`
	Alert       *rulesActionAlertXml_11_0_2       `xml:"alert,omitempty"`
	Drop        *rulesActionDropXml_11_0_2        `xml:"drop,omitempty"`
	ResetClient *rulesActionResetClientXml_11_0_2 `xml:"reset-client,omitempty"`
	ResetServer *rulesActionResetServerXml_11_0_2 `xml:"reset-server,omitempty"`
	ResetBoth   *rulesActionResetBothXml_11_0_2   `xml:"reset-both,omitempty"`
	BlockIp     *rulesActionBlockIpXml_11_0_2     `xml:"block-ip,omitempty"`
	Misc        []generic.Xml                     `xml:",any"`
}
type rulesActionDefaultXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type rulesActionAllowXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type rulesActionAlertXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type rulesActionDropXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type rulesActionResetClientXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type rulesActionResetServerXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type rulesActionResetBothXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type rulesActionBlockIpXml_11_0_2 struct {
	TrackBy  *string       `xml:"track-by,omitempty"`
	Duration *int64        `xml:"duration,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type threatExceptionContainerXml_11_0_2 struct {
	Entries []threatExceptionXml_11_0_2 `xml:"entry"`
}
type threatExceptionXml_11_0_2 struct {
	XMLName       xml.Name                                    `xml:"entry"`
	Name          string                                      `xml:"name,attr"`
	PacketCapture *string                                     `xml:"packet-capture,omitempty"`
	Action        *threatExceptionActionXml_11_0_2            `xml:"action,omitempty"`
	ExemptIp      *threatExceptionExemptIpContainerXml_11_0_2 `xml:"exempt-ip,omitempty"`
	Misc          []generic.Xml                               `xml:",any"`
}
type threatExceptionActionXml_11_0_2 struct {
	Default     *threatExceptionActionDefaultXml_11_0_2     `xml:"default,omitempty"`
	Allow       *threatExceptionActionAllowXml_11_0_2       `xml:"allow,omitempty"`
	Alert       *threatExceptionActionAlertXml_11_0_2       `xml:"alert,omitempty"`
	Drop        *threatExceptionActionDropXml_11_0_2        `xml:"drop,omitempty"`
	ResetBoth   *threatExceptionActionResetBothXml_11_0_2   `xml:"reset-both,omitempty"`
	ResetClient *threatExceptionActionResetClientXml_11_0_2 `xml:"reset-client,omitempty"`
	ResetServer *threatExceptionActionResetServerXml_11_0_2 `xml:"reset-server,omitempty"`
	BlockIp     *threatExceptionActionBlockIpXml_11_0_2     `xml:"block-ip,omitempty"`
	Misc        []generic.Xml                               `xml:",any"`
}
type threatExceptionActionDefaultXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type threatExceptionActionAllowXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type threatExceptionActionAlertXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type threatExceptionActionDropXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type threatExceptionActionResetBothXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type threatExceptionActionResetClientXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type threatExceptionActionResetServerXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type threatExceptionActionBlockIpXml_11_0_2 struct {
	TrackBy  *string       `xml:"track-by,omitempty"`
	Duration *int64        `xml:"duration,omitempty"`
	Misc     []generic.Xml `xml:",any"`
}
type threatExceptionExemptIpContainerXml_11_0_2 struct {
	Entries []threatExceptionExemptIpXml_11_0_2 `xml:"entry"`
}
type threatExceptionExemptIpXml_11_0_2 struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.BotnetDomains != nil {
		var obj botnetDomainsXml
		obj.MarshalFromObject(*s.BotnetDomains)
		o.BotnetDomains = &obj
	}
	o.CloudInlineAnalysis = util.YesNo(s.CloudInlineAnalysis, nil)
	o.Description = s.Description
	o.DisableOverride = s.DisableOverride
	if s.InlineExceptionEdlUrl != nil {
		o.InlineExceptionEdlUrl = util.StrToMem(s.InlineExceptionEdlUrl)
	}
	if s.InlineExceptionIpAddress != nil {
		o.InlineExceptionIpAddress = util.StrToMem(s.InlineExceptionIpAddress)
	}
	if s.MicaEngineSpywareEnabled != nil {
		var objs []micaEngineSpywareEnabledXml
		for _, elt := range s.MicaEngineSpywareEnabled {
			var obj micaEngineSpywareEnabledXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.MicaEngineSpywareEnabled = &micaEngineSpywareEnabledContainerXml{Entries: objs}
	}
	if s.Rules != nil {
		var objs []rulesXml
		for _, elt := range s.Rules {
			var obj rulesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Rules = &rulesContainerXml{Entries: objs}
	}
	if s.ThreatException != nil {
		var objs []threatExceptionXml
		for _, elt := range s.ThreatException {
			var obj threatExceptionXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ThreatException = &threatExceptionContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var botnetDomainsVal *BotnetDomains
	if o.BotnetDomains != nil {
		obj, err := o.BotnetDomains.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		botnetDomainsVal = obj
	}
	var inlineExceptionEdlUrlVal []string
	if o.InlineExceptionEdlUrl != nil {
		inlineExceptionEdlUrlVal = util.MemToStr(o.InlineExceptionEdlUrl)
	}
	var inlineExceptionIpAddressVal []string
	if o.InlineExceptionIpAddress != nil {
		inlineExceptionIpAddressVal = util.MemToStr(o.InlineExceptionIpAddress)
	}
	var micaEngineSpywareEnabledVal []MicaEngineSpywareEnabled
	if o.MicaEngineSpywareEnabled != nil {
		for _, elt := range o.MicaEngineSpywareEnabled.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			micaEngineSpywareEnabledVal = append(micaEngineSpywareEnabledVal, *obj)
		}
	}
	var rulesVal []Rules
	if o.Rules != nil {
		for _, elt := range o.Rules.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			rulesVal = append(rulesVal, *obj)
		}
	}
	var threatExceptionVal []ThreatException
	if o.ThreatException != nil {
		for _, elt := range o.ThreatException.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			threatExceptionVal = append(threatExceptionVal, *obj)
		}
	}

	result := &Entry{
		Name:                     o.Name,
		BotnetDomains:            botnetDomainsVal,
		CloudInlineAnalysis:      util.AsBool(o.CloudInlineAnalysis, nil),
		Description:              o.Description,
		DisableOverride:          o.DisableOverride,
		InlineExceptionEdlUrl:    inlineExceptionEdlUrlVal,
		InlineExceptionIpAddress: inlineExceptionIpAddressVal,
		MicaEngineSpywareEnabled: micaEngineSpywareEnabledVal,
		Rules:                    rulesVal,
		ThreatException:          threatExceptionVal,
		Misc:                     o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsXml) MarshalFromObject(s BotnetDomains) {
	if s.DnsSecurityCategories != nil {
		var objs []botnetDomainsDnsSecurityCategoriesXml
		for _, elt := range s.DnsSecurityCategories {
			var obj botnetDomainsDnsSecurityCategoriesXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.DnsSecurityCategories = &botnetDomainsDnsSecurityCategoriesContainerXml{Entries: objs}
	}
	if s.Lists != nil {
		var objs []botnetDomainsListsXml
		for _, elt := range s.Lists {
			var obj botnetDomainsListsXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Lists = &botnetDomainsListsContainerXml{Entries: objs}
	}
	if s.RtypeAction != nil {
		var obj botnetDomainsRtypeActionXml
		obj.MarshalFromObject(*s.RtypeAction)
		o.RtypeAction = &obj
	}
	if s.Sinkhole != nil {
		var obj botnetDomainsSinkholeXml
		obj.MarshalFromObject(*s.Sinkhole)
		o.Sinkhole = &obj
	}
	if s.ThreatException != nil {
		var objs []botnetDomainsThreatExceptionXml
		for _, elt := range s.ThreatException {
			var obj botnetDomainsThreatExceptionXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ThreatException = &botnetDomainsThreatExceptionContainerXml{Entries: objs}
	}
	if s.Whitelist != nil {
		var objs []botnetDomainsWhitelistXml
		for _, elt := range s.Whitelist {
			var obj botnetDomainsWhitelistXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Whitelist = &botnetDomainsWhitelistContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o botnetDomainsXml) UnmarshalToObject() (*BotnetDomains, error) {
	var dnsSecurityCategoriesVal []BotnetDomainsDnsSecurityCategories
	if o.DnsSecurityCategories != nil {
		for _, elt := range o.DnsSecurityCategories.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			dnsSecurityCategoriesVal = append(dnsSecurityCategoriesVal, *obj)
		}
	}
	var listsVal []BotnetDomainsLists
	if o.Lists != nil {
		for _, elt := range o.Lists.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			listsVal = append(listsVal, *obj)
		}
	}
	var rtypeActionVal *BotnetDomainsRtypeAction
	if o.RtypeAction != nil {
		obj, err := o.RtypeAction.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		rtypeActionVal = obj
	}
	var sinkholeVal *BotnetDomainsSinkhole
	if o.Sinkhole != nil {
		obj, err := o.Sinkhole.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sinkholeVal = obj
	}
	var threatExceptionVal []BotnetDomainsThreatException
	if o.ThreatException != nil {
		for _, elt := range o.ThreatException.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			threatExceptionVal = append(threatExceptionVal, *obj)
		}
	}
	var whitelistVal []BotnetDomainsWhitelist
	if o.Whitelist != nil {
		for _, elt := range o.Whitelist.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			whitelistVal = append(whitelistVal, *obj)
		}
	}

	result := &BotnetDomains{
		DnsSecurityCategories: dnsSecurityCategoriesVal,
		Lists:                 listsVal,
		RtypeAction:           rtypeActionVal,
		Sinkhole:              sinkholeVal,
		ThreatException:       threatExceptionVal,
		Whitelist:             whitelistVal,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsDnsSecurityCategoriesXml) MarshalFromObject(s BotnetDomainsDnsSecurityCategories) {
	o.Name = s.Name
	o.Action = s.Action
	o.LogLevel = s.LogLevel
	o.PacketCapture = s.PacketCapture
	o.Misc = s.Misc
}

func (o botnetDomainsDnsSecurityCategoriesXml) UnmarshalToObject() (*BotnetDomainsDnsSecurityCategories, error) {

	result := &BotnetDomainsDnsSecurityCategories{
		Name:          o.Name,
		Action:        o.Action,
		LogLevel:      o.LogLevel,
		PacketCapture: o.PacketCapture,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsListsXml) MarshalFromObject(s BotnetDomainsLists) {
	o.Name = s.Name
	if s.Action != nil {
		var obj botnetDomainsListsActionXml
		obj.MarshalFromObject(*s.Action)
		o.Action = &obj
	}
	o.PacketCapture = s.PacketCapture
	o.Misc = s.Misc
}

func (o botnetDomainsListsXml) UnmarshalToObject() (*BotnetDomainsLists, error) {
	var actionVal *BotnetDomainsListsAction
	if o.Action != nil {
		obj, err := o.Action.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		actionVal = obj
	}

	result := &BotnetDomainsLists{
		Name:          o.Name,
		Action:        actionVal,
		PacketCapture: o.PacketCapture,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsListsActionXml) MarshalFromObject(s BotnetDomainsListsAction) {
	if s.Alert != nil {
		var obj botnetDomainsListsActionAlertXml
		obj.MarshalFromObject(*s.Alert)
		o.Alert = &obj
	}
	if s.Allow != nil {
		var obj botnetDomainsListsActionAllowXml
		obj.MarshalFromObject(*s.Allow)
		o.Allow = &obj
	}
	if s.Block != nil {
		var obj botnetDomainsListsActionBlockXml
		obj.MarshalFromObject(*s.Block)
		o.Block = &obj
	}
	if s.Sinkhole != nil {
		var obj botnetDomainsListsActionSinkholeXml
		obj.MarshalFromObject(*s.Sinkhole)
		o.Sinkhole = &obj
	}
	o.Misc = s.Misc
}

func (o botnetDomainsListsActionXml) UnmarshalToObject() (*BotnetDomainsListsAction, error) {
	var alertVal *BotnetDomainsListsActionAlert
	if o.Alert != nil {
		obj, err := o.Alert.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		alertVal = obj
	}
	var allowVal *BotnetDomainsListsActionAllow
	if o.Allow != nil {
		obj, err := o.Allow.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allowVal = obj
	}
	var blockVal *BotnetDomainsListsActionBlock
	if o.Block != nil {
		obj, err := o.Block.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		blockVal = obj
	}
	var sinkholeVal *BotnetDomainsListsActionSinkhole
	if o.Sinkhole != nil {
		obj, err := o.Sinkhole.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sinkholeVal = obj
	}

	result := &BotnetDomainsListsAction{
		Alert:    alertVal,
		Allow:    allowVal,
		Block:    blockVal,
		Sinkhole: sinkholeVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsListsActionAlertXml) MarshalFromObject(s BotnetDomainsListsActionAlert) {
	o.Misc = s.Misc
}

func (o botnetDomainsListsActionAlertXml) UnmarshalToObject() (*BotnetDomainsListsActionAlert, error) {

	result := &BotnetDomainsListsActionAlert{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsListsActionAllowXml) MarshalFromObject(s BotnetDomainsListsActionAllow) {
	o.Misc = s.Misc
}

func (o botnetDomainsListsActionAllowXml) UnmarshalToObject() (*BotnetDomainsListsActionAllow, error) {

	result := &BotnetDomainsListsActionAllow{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsListsActionBlockXml) MarshalFromObject(s BotnetDomainsListsActionBlock) {
	o.Misc = s.Misc
}

func (o botnetDomainsListsActionBlockXml) UnmarshalToObject() (*BotnetDomainsListsActionBlock, error) {

	result := &BotnetDomainsListsActionBlock{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsListsActionSinkholeXml) MarshalFromObject(s BotnetDomainsListsActionSinkhole) {
	o.Misc = s.Misc
}

func (o botnetDomainsListsActionSinkholeXml) UnmarshalToObject() (*BotnetDomainsListsActionSinkhole, error) {

	result := &BotnetDomainsListsActionSinkhole{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsRtypeActionXml) MarshalFromObject(s BotnetDomainsRtypeAction) {
	o.Any = s.Any
	o.Https = s.Https
	o.Svcb = s.Svcb
	o.Misc = s.Misc
}

func (o botnetDomainsRtypeActionXml) UnmarshalToObject() (*BotnetDomainsRtypeAction, error) {

	result := &BotnetDomainsRtypeAction{
		Any:   o.Any,
		Https: o.Https,
		Svcb:  o.Svcb,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsSinkholeXml) MarshalFromObject(s BotnetDomainsSinkhole) {
	o.Ipv4Address = s.Ipv4Address
	o.Ipv6Address = s.Ipv6Address
	o.Misc = s.Misc
}

func (o botnetDomainsSinkholeXml) UnmarshalToObject() (*BotnetDomainsSinkhole, error) {

	result := &BotnetDomainsSinkhole{
		Ipv4Address: o.Ipv4Address,
		Ipv6Address: o.Ipv6Address,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsThreatExceptionXml) MarshalFromObject(s BotnetDomainsThreatException) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o botnetDomainsThreatExceptionXml) UnmarshalToObject() (*BotnetDomainsThreatException, error) {

	result := &BotnetDomainsThreatException{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsWhitelistXml) MarshalFromObject(s BotnetDomainsWhitelist) {
	o.Name = s.Name
	o.Description = s.Description
	o.Misc = s.Misc
}

func (o botnetDomainsWhitelistXml) UnmarshalToObject() (*BotnetDomainsWhitelist, error) {

	result := &BotnetDomainsWhitelist{
		Name:        o.Name,
		Description: o.Description,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *micaEngineSpywareEnabledXml) MarshalFromObject(s MicaEngineSpywareEnabled) {
	o.Name = s.Name
	o.InlinePolicyAction = s.InlinePolicyAction
	o.Misc = s.Misc
}

func (o micaEngineSpywareEnabledXml) UnmarshalToObject() (*MicaEngineSpywareEnabled, error) {

	result := &MicaEngineSpywareEnabled{
		Name:               o.Name,
		InlinePolicyAction: o.InlinePolicyAction,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *rulesXml) MarshalFromObject(s Rules) {
	o.Name = s.Name
	o.ThreatName = s.ThreatName
	o.Category = s.Category
	o.PacketCapture = s.PacketCapture
	if s.Severity != nil {
		o.Severity = util.StrToMem(s.Severity)
	}
	if s.Action != nil {
		var obj rulesActionXml
		obj.MarshalFromObject(*s.Action)
		o.Action = &obj
	}
	o.Misc = s.Misc
}

func (o rulesXml) UnmarshalToObject() (*Rules, error) {
	var severityVal []string
	if o.Severity != nil {
		severityVal = util.MemToStr(o.Severity)
	}
	var actionVal *RulesAction
	if o.Action != nil {
		obj, err := o.Action.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		actionVal = obj
	}

	result := &Rules{
		Name:          o.Name,
		ThreatName:    o.ThreatName,
		Category:      o.Category,
		PacketCapture: o.PacketCapture,
		Severity:      severityVal,
		Action:        actionVal,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *rulesActionXml) MarshalFromObject(s RulesAction) {
	if s.Default != nil {
		var obj rulesActionDefaultXml
		obj.MarshalFromObject(*s.Default)
		o.Default = &obj
	}
	if s.Allow != nil {
		var obj rulesActionAllowXml
		obj.MarshalFromObject(*s.Allow)
		o.Allow = &obj
	}
	if s.Alert != nil {
		var obj rulesActionAlertXml
		obj.MarshalFromObject(*s.Alert)
		o.Alert = &obj
	}
	if s.Drop != nil {
		var obj rulesActionDropXml
		obj.MarshalFromObject(*s.Drop)
		o.Drop = &obj
	}
	if s.ResetClient != nil {
		var obj rulesActionResetClientXml
		obj.MarshalFromObject(*s.ResetClient)
		o.ResetClient = &obj
	}
	if s.ResetServer != nil {
		var obj rulesActionResetServerXml
		obj.MarshalFromObject(*s.ResetServer)
		o.ResetServer = &obj
	}
	if s.ResetBoth != nil {
		var obj rulesActionResetBothXml
		obj.MarshalFromObject(*s.ResetBoth)
		o.ResetBoth = &obj
	}
	if s.BlockIp != nil {
		var obj rulesActionBlockIpXml
		obj.MarshalFromObject(*s.BlockIp)
		o.BlockIp = &obj
	}
	o.Misc = s.Misc
}

func (o rulesActionXml) UnmarshalToObject() (*RulesAction, error) {
	var defaultVal *RulesActionDefault
	if o.Default != nil {
		obj, err := o.Default.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultVal = obj
	}
	var allowVal *RulesActionAllow
	if o.Allow != nil {
		obj, err := o.Allow.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allowVal = obj
	}
	var alertVal *RulesActionAlert
	if o.Alert != nil {
		obj, err := o.Alert.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		alertVal = obj
	}
	var dropVal *RulesActionDrop
	if o.Drop != nil {
		obj, err := o.Drop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dropVal = obj
	}
	var resetClientVal *RulesActionResetClient
	if o.ResetClient != nil {
		obj, err := o.ResetClient.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetClientVal = obj
	}
	var resetServerVal *RulesActionResetServer
	if o.ResetServer != nil {
		obj, err := o.ResetServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetServerVal = obj
	}
	var resetBothVal *RulesActionResetBoth
	if o.ResetBoth != nil {
		obj, err := o.ResetBoth.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetBothVal = obj
	}
	var blockIpVal *RulesActionBlockIp
	if o.BlockIp != nil {
		obj, err := o.BlockIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		blockIpVal = obj
	}

	result := &RulesAction{
		Default:     defaultVal,
		Allow:       allowVal,
		Alert:       alertVal,
		Drop:        dropVal,
		ResetClient: resetClientVal,
		ResetServer: resetServerVal,
		ResetBoth:   resetBothVal,
		BlockIp:     blockIpVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *rulesActionDefaultXml) MarshalFromObject(s RulesActionDefault) {
	o.Misc = s.Misc
}

func (o rulesActionDefaultXml) UnmarshalToObject() (*RulesActionDefault, error) {

	result := &RulesActionDefault{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *rulesActionAllowXml) MarshalFromObject(s RulesActionAllow) {
	o.Misc = s.Misc
}

func (o rulesActionAllowXml) UnmarshalToObject() (*RulesActionAllow, error) {

	result := &RulesActionAllow{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *rulesActionAlertXml) MarshalFromObject(s RulesActionAlert) {
	o.Misc = s.Misc
}

func (o rulesActionAlertXml) UnmarshalToObject() (*RulesActionAlert, error) {

	result := &RulesActionAlert{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *rulesActionDropXml) MarshalFromObject(s RulesActionDrop) {
	o.Misc = s.Misc
}

func (o rulesActionDropXml) UnmarshalToObject() (*RulesActionDrop, error) {

	result := &RulesActionDrop{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *rulesActionResetClientXml) MarshalFromObject(s RulesActionResetClient) {
	o.Misc = s.Misc
}

func (o rulesActionResetClientXml) UnmarshalToObject() (*RulesActionResetClient, error) {

	result := &RulesActionResetClient{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *rulesActionResetServerXml) MarshalFromObject(s RulesActionResetServer) {
	o.Misc = s.Misc
}

func (o rulesActionResetServerXml) UnmarshalToObject() (*RulesActionResetServer, error) {

	result := &RulesActionResetServer{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *rulesActionResetBothXml) MarshalFromObject(s RulesActionResetBoth) {
	o.Misc = s.Misc
}

func (o rulesActionResetBothXml) UnmarshalToObject() (*RulesActionResetBoth, error) {

	result := &RulesActionResetBoth{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *rulesActionBlockIpXml) MarshalFromObject(s RulesActionBlockIp) {
	o.TrackBy = s.TrackBy
	o.Duration = s.Duration
	o.Misc = s.Misc
}

func (o rulesActionBlockIpXml) UnmarshalToObject() (*RulesActionBlockIp, error) {

	result := &RulesActionBlockIp{
		TrackBy:  o.TrackBy,
		Duration: o.Duration,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *threatExceptionXml) MarshalFromObject(s ThreatException) {
	o.Name = s.Name
	o.PacketCapture = s.PacketCapture
	if s.Action != nil {
		var obj threatExceptionActionXml
		obj.MarshalFromObject(*s.Action)
		o.Action = &obj
	}
	if s.ExemptIp != nil {
		var objs []threatExceptionExemptIpXml
		for _, elt := range s.ExemptIp {
			var obj threatExceptionExemptIpXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ExemptIp = &threatExceptionExemptIpContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o threatExceptionXml) UnmarshalToObject() (*ThreatException, error) {
	var actionVal *ThreatExceptionAction
	if o.Action != nil {
		obj, err := o.Action.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		actionVal = obj
	}
	var exemptIpVal []ThreatExceptionExemptIp
	if o.ExemptIp != nil {
		for _, elt := range o.ExemptIp.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			exemptIpVal = append(exemptIpVal, *obj)
		}
	}

	result := &ThreatException{
		Name:          o.Name,
		PacketCapture: o.PacketCapture,
		Action:        actionVal,
		ExemptIp:      exemptIpVal,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionXml) MarshalFromObject(s ThreatExceptionAction) {
	if s.Default != nil {
		var obj threatExceptionActionDefaultXml
		obj.MarshalFromObject(*s.Default)
		o.Default = &obj
	}
	if s.Allow != nil {
		var obj threatExceptionActionAllowXml
		obj.MarshalFromObject(*s.Allow)
		o.Allow = &obj
	}
	if s.Alert != nil {
		var obj threatExceptionActionAlertXml
		obj.MarshalFromObject(*s.Alert)
		o.Alert = &obj
	}
	if s.Drop != nil {
		var obj threatExceptionActionDropXml
		obj.MarshalFromObject(*s.Drop)
		o.Drop = &obj
	}
	if s.ResetBoth != nil {
		var obj threatExceptionActionResetBothXml
		obj.MarshalFromObject(*s.ResetBoth)
		o.ResetBoth = &obj
	}
	if s.ResetClient != nil {
		var obj threatExceptionActionResetClientXml
		obj.MarshalFromObject(*s.ResetClient)
		o.ResetClient = &obj
	}
	if s.ResetServer != nil {
		var obj threatExceptionActionResetServerXml
		obj.MarshalFromObject(*s.ResetServer)
		o.ResetServer = &obj
	}
	if s.BlockIp != nil {
		var obj threatExceptionActionBlockIpXml
		obj.MarshalFromObject(*s.BlockIp)
		o.BlockIp = &obj
	}
	o.Misc = s.Misc
}

func (o threatExceptionActionXml) UnmarshalToObject() (*ThreatExceptionAction, error) {
	var defaultVal *ThreatExceptionActionDefault
	if o.Default != nil {
		obj, err := o.Default.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultVal = obj
	}
	var allowVal *ThreatExceptionActionAllow
	if o.Allow != nil {
		obj, err := o.Allow.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allowVal = obj
	}
	var alertVal *ThreatExceptionActionAlert
	if o.Alert != nil {
		obj, err := o.Alert.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		alertVal = obj
	}
	var dropVal *ThreatExceptionActionDrop
	if o.Drop != nil {
		obj, err := o.Drop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dropVal = obj
	}
	var resetBothVal *ThreatExceptionActionResetBoth
	if o.ResetBoth != nil {
		obj, err := o.ResetBoth.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetBothVal = obj
	}
	var resetClientVal *ThreatExceptionActionResetClient
	if o.ResetClient != nil {
		obj, err := o.ResetClient.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetClientVal = obj
	}
	var resetServerVal *ThreatExceptionActionResetServer
	if o.ResetServer != nil {
		obj, err := o.ResetServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetServerVal = obj
	}
	var blockIpVal *ThreatExceptionActionBlockIp
	if o.BlockIp != nil {
		obj, err := o.BlockIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		blockIpVal = obj
	}

	result := &ThreatExceptionAction{
		Default:     defaultVal,
		Allow:       allowVal,
		Alert:       alertVal,
		Drop:        dropVal,
		ResetBoth:   resetBothVal,
		ResetClient: resetClientVal,
		ResetServer: resetServerVal,
		BlockIp:     blockIpVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionDefaultXml) MarshalFromObject(s ThreatExceptionActionDefault) {
	o.Misc = s.Misc
}

func (o threatExceptionActionDefaultXml) UnmarshalToObject() (*ThreatExceptionActionDefault, error) {

	result := &ThreatExceptionActionDefault{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionAllowXml) MarshalFromObject(s ThreatExceptionActionAllow) {
	o.Misc = s.Misc
}

func (o threatExceptionActionAllowXml) UnmarshalToObject() (*ThreatExceptionActionAllow, error) {

	result := &ThreatExceptionActionAllow{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionAlertXml) MarshalFromObject(s ThreatExceptionActionAlert) {
	o.Misc = s.Misc
}

func (o threatExceptionActionAlertXml) UnmarshalToObject() (*ThreatExceptionActionAlert, error) {

	result := &ThreatExceptionActionAlert{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionDropXml) MarshalFromObject(s ThreatExceptionActionDrop) {
	o.Misc = s.Misc
}

func (o threatExceptionActionDropXml) UnmarshalToObject() (*ThreatExceptionActionDrop, error) {

	result := &ThreatExceptionActionDrop{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionResetBothXml) MarshalFromObject(s ThreatExceptionActionResetBoth) {
	o.Misc = s.Misc
}

func (o threatExceptionActionResetBothXml) UnmarshalToObject() (*ThreatExceptionActionResetBoth, error) {

	result := &ThreatExceptionActionResetBoth{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionResetClientXml) MarshalFromObject(s ThreatExceptionActionResetClient) {
	o.Misc = s.Misc
}

func (o threatExceptionActionResetClientXml) UnmarshalToObject() (*ThreatExceptionActionResetClient, error) {

	result := &ThreatExceptionActionResetClient{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionResetServerXml) MarshalFromObject(s ThreatExceptionActionResetServer) {
	o.Misc = s.Misc
}

func (o threatExceptionActionResetServerXml) UnmarshalToObject() (*ThreatExceptionActionResetServer, error) {

	result := &ThreatExceptionActionResetServer{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionBlockIpXml) MarshalFromObject(s ThreatExceptionActionBlockIp) {
	o.TrackBy = s.TrackBy
	o.Duration = s.Duration
	o.Misc = s.Misc
}

func (o threatExceptionActionBlockIpXml) UnmarshalToObject() (*ThreatExceptionActionBlockIp, error) {

	result := &ThreatExceptionActionBlockIp{
		TrackBy:  o.TrackBy,
		Duration: o.Duration,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *threatExceptionExemptIpXml) MarshalFromObject(s ThreatExceptionExemptIp) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o threatExceptionExemptIpXml) UnmarshalToObject() (*ThreatExceptionExemptIp, error) {

	result := &ThreatExceptionExemptIp{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *entryXml_11_0_2) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.BotnetDomains != nil {
		var obj botnetDomainsXml_11_0_2
		obj.MarshalFromObject(*s.BotnetDomains)
		o.BotnetDomains = &obj
	}
	o.CloudInlineAnalysis = util.YesNo(s.CloudInlineAnalysis, nil)
	o.Description = s.Description
	o.DisableOverride = s.DisableOverride
	if s.InlineExceptionEdlUrl != nil {
		o.InlineExceptionEdlUrl = util.StrToMem(s.InlineExceptionEdlUrl)
	}
	if s.InlineExceptionIpAddress != nil {
		o.InlineExceptionIpAddress = util.StrToMem(s.InlineExceptionIpAddress)
	}
	if s.MicaEngineSpywareEnabled != nil {
		var objs []micaEngineSpywareEnabledXml_11_0_2
		for _, elt := range s.MicaEngineSpywareEnabled {
			var obj micaEngineSpywareEnabledXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.MicaEngineSpywareEnabled = &micaEngineSpywareEnabledContainerXml_11_0_2{Entries: objs}
	}
	if s.Rules != nil {
		var objs []rulesXml_11_0_2
		for _, elt := range s.Rules {
			var obj rulesXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Rules = &rulesContainerXml_11_0_2{Entries: objs}
	}
	if s.ThreatException != nil {
		var objs []threatExceptionXml_11_0_2
		for _, elt := range s.ThreatException {
			var obj threatExceptionXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ThreatException = &threatExceptionContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o entryXml_11_0_2) UnmarshalToObject() (*Entry, error) {
	var botnetDomainsVal *BotnetDomains
	if o.BotnetDomains != nil {
		obj, err := o.BotnetDomains.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		botnetDomainsVal = obj
	}
	var inlineExceptionEdlUrlVal []string
	if o.InlineExceptionEdlUrl != nil {
		inlineExceptionEdlUrlVal = util.MemToStr(o.InlineExceptionEdlUrl)
	}
	var inlineExceptionIpAddressVal []string
	if o.InlineExceptionIpAddress != nil {
		inlineExceptionIpAddressVal = util.MemToStr(o.InlineExceptionIpAddress)
	}
	var micaEngineSpywareEnabledVal []MicaEngineSpywareEnabled
	if o.MicaEngineSpywareEnabled != nil {
		for _, elt := range o.MicaEngineSpywareEnabled.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			micaEngineSpywareEnabledVal = append(micaEngineSpywareEnabledVal, *obj)
		}
	}
	var rulesVal []Rules
	if o.Rules != nil {
		for _, elt := range o.Rules.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			rulesVal = append(rulesVal, *obj)
		}
	}
	var threatExceptionVal []ThreatException
	if o.ThreatException != nil {
		for _, elt := range o.ThreatException.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			threatExceptionVal = append(threatExceptionVal, *obj)
		}
	}

	result := &Entry{
		Name:                     o.Name,
		BotnetDomains:            botnetDomainsVal,
		CloudInlineAnalysis:      util.AsBool(o.CloudInlineAnalysis, nil),
		Description:              o.Description,
		DisableOverride:          o.DisableOverride,
		InlineExceptionEdlUrl:    inlineExceptionEdlUrlVal,
		InlineExceptionIpAddress: inlineExceptionIpAddressVal,
		MicaEngineSpywareEnabled: micaEngineSpywareEnabledVal,
		Rules:                    rulesVal,
		ThreatException:          threatExceptionVal,
		Misc:                     o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsXml_11_0_2) MarshalFromObject(s BotnetDomains) {
	if s.DnsSecurityCategories != nil {
		var objs []botnetDomainsDnsSecurityCategoriesXml_11_0_2
		for _, elt := range s.DnsSecurityCategories {
			var obj botnetDomainsDnsSecurityCategoriesXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.DnsSecurityCategories = &botnetDomainsDnsSecurityCategoriesContainerXml_11_0_2{Entries: objs}
	}
	if s.Lists != nil {
		var objs []botnetDomainsListsXml_11_0_2
		for _, elt := range s.Lists {
			var obj botnetDomainsListsXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Lists = &botnetDomainsListsContainerXml_11_0_2{Entries: objs}
	}
	if s.RtypeAction != nil {
		var obj botnetDomainsRtypeActionXml_11_0_2
		obj.MarshalFromObject(*s.RtypeAction)
		o.RtypeAction = &obj
	}
	if s.Sinkhole != nil {
		var obj botnetDomainsSinkholeXml_11_0_2
		obj.MarshalFromObject(*s.Sinkhole)
		o.Sinkhole = &obj
	}
	if s.ThreatException != nil {
		var objs []botnetDomainsThreatExceptionXml_11_0_2
		for _, elt := range s.ThreatException {
			var obj botnetDomainsThreatExceptionXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ThreatException = &botnetDomainsThreatExceptionContainerXml_11_0_2{Entries: objs}
	}
	if s.Whitelist != nil {
		var objs []botnetDomainsWhitelistXml_11_0_2
		for _, elt := range s.Whitelist {
			var obj botnetDomainsWhitelistXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.Whitelist = &botnetDomainsWhitelistContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o botnetDomainsXml_11_0_2) UnmarshalToObject() (*BotnetDomains, error) {
	var dnsSecurityCategoriesVal []BotnetDomainsDnsSecurityCategories
	if o.DnsSecurityCategories != nil {
		for _, elt := range o.DnsSecurityCategories.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			dnsSecurityCategoriesVal = append(dnsSecurityCategoriesVal, *obj)
		}
	}
	var listsVal []BotnetDomainsLists
	if o.Lists != nil {
		for _, elt := range o.Lists.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			listsVal = append(listsVal, *obj)
		}
	}
	var rtypeActionVal *BotnetDomainsRtypeAction
	if o.RtypeAction != nil {
		obj, err := o.RtypeAction.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		rtypeActionVal = obj
	}
	var sinkholeVal *BotnetDomainsSinkhole
	if o.Sinkhole != nil {
		obj, err := o.Sinkhole.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sinkholeVal = obj
	}
	var threatExceptionVal []BotnetDomainsThreatException
	if o.ThreatException != nil {
		for _, elt := range o.ThreatException.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			threatExceptionVal = append(threatExceptionVal, *obj)
		}
	}
	var whitelistVal []BotnetDomainsWhitelist
	if o.Whitelist != nil {
		for _, elt := range o.Whitelist.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			whitelistVal = append(whitelistVal, *obj)
		}
	}

	result := &BotnetDomains{
		DnsSecurityCategories: dnsSecurityCategoriesVal,
		Lists:                 listsVal,
		RtypeAction:           rtypeActionVal,
		Sinkhole:              sinkholeVal,
		ThreatException:       threatExceptionVal,
		Whitelist:             whitelistVal,
		Misc:                  o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsDnsSecurityCategoriesXml_11_0_2) MarshalFromObject(s BotnetDomainsDnsSecurityCategories) {
	o.Name = s.Name
	o.Action = s.Action
	o.LogLevel = s.LogLevel
	o.PacketCapture = s.PacketCapture
	o.Misc = s.Misc
}

func (o botnetDomainsDnsSecurityCategoriesXml_11_0_2) UnmarshalToObject() (*BotnetDomainsDnsSecurityCategories, error) {

	result := &BotnetDomainsDnsSecurityCategories{
		Name:          o.Name,
		Action:        o.Action,
		LogLevel:      o.LogLevel,
		PacketCapture: o.PacketCapture,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsListsXml_11_0_2) MarshalFromObject(s BotnetDomainsLists) {
	o.Name = s.Name
	if s.Action != nil {
		var obj botnetDomainsListsActionXml_11_0_2
		obj.MarshalFromObject(*s.Action)
		o.Action = &obj
	}
	o.PacketCapture = s.PacketCapture
	o.Misc = s.Misc
}

func (o botnetDomainsListsXml_11_0_2) UnmarshalToObject() (*BotnetDomainsLists, error) {
	var actionVal *BotnetDomainsListsAction
	if o.Action != nil {
		obj, err := o.Action.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		actionVal = obj
	}

	result := &BotnetDomainsLists{
		Name:          o.Name,
		Action:        actionVal,
		PacketCapture: o.PacketCapture,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsListsActionXml_11_0_2) MarshalFromObject(s BotnetDomainsListsAction) {
	if s.Alert != nil {
		var obj botnetDomainsListsActionAlertXml_11_0_2
		obj.MarshalFromObject(*s.Alert)
		o.Alert = &obj
	}
	if s.Allow != nil {
		var obj botnetDomainsListsActionAllowXml_11_0_2
		obj.MarshalFromObject(*s.Allow)
		o.Allow = &obj
	}
	if s.Block != nil {
		var obj botnetDomainsListsActionBlockXml_11_0_2
		obj.MarshalFromObject(*s.Block)
		o.Block = &obj
	}
	if s.Sinkhole != nil {
		var obj botnetDomainsListsActionSinkholeXml_11_0_2
		obj.MarshalFromObject(*s.Sinkhole)
		o.Sinkhole = &obj
	}
	o.Misc = s.Misc
}

func (o botnetDomainsListsActionXml_11_0_2) UnmarshalToObject() (*BotnetDomainsListsAction, error) {
	var alertVal *BotnetDomainsListsActionAlert
	if o.Alert != nil {
		obj, err := o.Alert.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		alertVal = obj
	}
	var allowVal *BotnetDomainsListsActionAllow
	if o.Allow != nil {
		obj, err := o.Allow.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allowVal = obj
	}
	var blockVal *BotnetDomainsListsActionBlock
	if o.Block != nil {
		obj, err := o.Block.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		blockVal = obj
	}
	var sinkholeVal *BotnetDomainsListsActionSinkhole
	if o.Sinkhole != nil {
		obj, err := o.Sinkhole.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sinkholeVal = obj
	}

	result := &BotnetDomainsListsAction{
		Alert:    alertVal,
		Allow:    allowVal,
		Block:    blockVal,
		Sinkhole: sinkholeVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsListsActionAlertXml_11_0_2) MarshalFromObject(s BotnetDomainsListsActionAlert) {
	o.Misc = s.Misc
}

func (o botnetDomainsListsActionAlertXml_11_0_2) UnmarshalToObject() (*BotnetDomainsListsActionAlert, error) {

	result := &BotnetDomainsListsActionAlert{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsListsActionAllowXml_11_0_2) MarshalFromObject(s BotnetDomainsListsActionAllow) {
	o.Misc = s.Misc
}

func (o botnetDomainsListsActionAllowXml_11_0_2) UnmarshalToObject() (*BotnetDomainsListsActionAllow, error) {

	result := &BotnetDomainsListsActionAllow{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsListsActionBlockXml_11_0_2) MarshalFromObject(s BotnetDomainsListsActionBlock) {
	o.Misc = s.Misc
}

func (o botnetDomainsListsActionBlockXml_11_0_2) UnmarshalToObject() (*BotnetDomainsListsActionBlock, error) {

	result := &BotnetDomainsListsActionBlock{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsListsActionSinkholeXml_11_0_2) MarshalFromObject(s BotnetDomainsListsActionSinkhole) {
	o.Misc = s.Misc
}

func (o botnetDomainsListsActionSinkholeXml_11_0_2) UnmarshalToObject() (*BotnetDomainsListsActionSinkhole, error) {

	result := &BotnetDomainsListsActionSinkhole{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsRtypeActionXml_11_0_2) MarshalFromObject(s BotnetDomainsRtypeAction) {
	o.Any = s.Any
	o.Https = s.Https
	o.Svcb = s.Svcb
	o.Misc = s.Misc
}

func (o botnetDomainsRtypeActionXml_11_0_2) UnmarshalToObject() (*BotnetDomainsRtypeAction, error) {

	result := &BotnetDomainsRtypeAction{
		Any:   o.Any,
		Https: o.Https,
		Svcb:  o.Svcb,
		Misc:  o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsSinkholeXml_11_0_2) MarshalFromObject(s BotnetDomainsSinkhole) {
	o.Ipv4Address = s.Ipv4Address
	o.Ipv6Address = s.Ipv6Address
	o.Misc = s.Misc
}

func (o botnetDomainsSinkholeXml_11_0_2) UnmarshalToObject() (*BotnetDomainsSinkhole, error) {

	result := &BotnetDomainsSinkhole{
		Ipv4Address: o.Ipv4Address,
		Ipv6Address: o.Ipv6Address,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsThreatExceptionXml_11_0_2) MarshalFromObject(s BotnetDomainsThreatException) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o botnetDomainsThreatExceptionXml_11_0_2) UnmarshalToObject() (*BotnetDomainsThreatException, error) {

	result := &BotnetDomainsThreatException{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *botnetDomainsWhitelistXml_11_0_2) MarshalFromObject(s BotnetDomainsWhitelist) {
	o.Name = s.Name
	o.Description = s.Description
	o.Misc = s.Misc
}

func (o botnetDomainsWhitelistXml_11_0_2) UnmarshalToObject() (*BotnetDomainsWhitelist, error) {

	result := &BotnetDomainsWhitelist{
		Name:        o.Name,
		Description: o.Description,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *micaEngineSpywareEnabledXml_11_0_2) MarshalFromObject(s MicaEngineSpywareEnabled) {
	o.Name = s.Name
	o.InlinePolicyAction = s.InlinePolicyAction
	o.Misc = s.Misc
}

func (o micaEngineSpywareEnabledXml_11_0_2) UnmarshalToObject() (*MicaEngineSpywareEnabled, error) {

	result := &MicaEngineSpywareEnabled{
		Name:               o.Name,
		InlinePolicyAction: o.InlinePolicyAction,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *rulesXml_11_0_2) MarshalFromObject(s Rules) {
	o.Name = s.Name
	o.ThreatName = s.ThreatName
	o.Category = s.Category
	o.PacketCapture = s.PacketCapture
	if s.Severity != nil {
		o.Severity = util.StrToMem(s.Severity)
	}
	if s.Action != nil {
		var obj rulesActionXml_11_0_2
		obj.MarshalFromObject(*s.Action)
		o.Action = &obj
	}
	o.Misc = s.Misc
}

func (o rulesXml_11_0_2) UnmarshalToObject() (*Rules, error) {
	var severityVal []string
	if o.Severity != nil {
		severityVal = util.MemToStr(o.Severity)
	}
	var actionVal *RulesAction
	if o.Action != nil {
		obj, err := o.Action.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		actionVal = obj
	}

	result := &Rules{
		Name:          o.Name,
		ThreatName:    o.ThreatName,
		Category:      o.Category,
		PacketCapture: o.PacketCapture,
		Severity:      severityVal,
		Action:        actionVal,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *rulesActionXml_11_0_2) MarshalFromObject(s RulesAction) {
	if s.Default != nil {
		var obj rulesActionDefaultXml_11_0_2
		obj.MarshalFromObject(*s.Default)
		o.Default = &obj
	}
	if s.Allow != nil {
		var obj rulesActionAllowXml_11_0_2
		obj.MarshalFromObject(*s.Allow)
		o.Allow = &obj
	}
	if s.Alert != nil {
		var obj rulesActionAlertXml_11_0_2
		obj.MarshalFromObject(*s.Alert)
		o.Alert = &obj
	}
	if s.Drop != nil {
		var obj rulesActionDropXml_11_0_2
		obj.MarshalFromObject(*s.Drop)
		o.Drop = &obj
	}
	if s.ResetClient != nil {
		var obj rulesActionResetClientXml_11_0_2
		obj.MarshalFromObject(*s.ResetClient)
		o.ResetClient = &obj
	}
	if s.ResetServer != nil {
		var obj rulesActionResetServerXml_11_0_2
		obj.MarshalFromObject(*s.ResetServer)
		o.ResetServer = &obj
	}
	if s.ResetBoth != nil {
		var obj rulesActionResetBothXml_11_0_2
		obj.MarshalFromObject(*s.ResetBoth)
		o.ResetBoth = &obj
	}
	if s.BlockIp != nil {
		var obj rulesActionBlockIpXml_11_0_2
		obj.MarshalFromObject(*s.BlockIp)
		o.BlockIp = &obj
	}
	o.Misc = s.Misc
}

func (o rulesActionXml_11_0_2) UnmarshalToObject() (*RulesAction, error) {
	var defaultVal *RulesActionDefault
	if o.Default != nil {
		obj, err := o.Default.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultVal = obj
	}
	var allowVal *RulesActionAllow
	if o.Allow != nil {
		obj, err := o.Allow.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allowVal = obj
	}
	var alertVal *RulesActionAlert
	if o.Alert != nil {
		obj, err := o.Alert.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		alertVal = obj
	}
	var dropVal *RulesActionDrop
	if o.Drop != nil {
		obj, err := o.Drop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dropVal = obj
	}
	var resetClientVal *RulesActionResetClient
	if o.ResetClient != nil {
		obj, err := o.ResetClient.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetClientVal = obj
	}
	var resetServerVal *RulesActionResetServer
	if o.ResetServer != nil {
		obj, err := o.ResetServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetServerVal = obj
	}
	var resetBothVal *RulesActionResetBoth
	if o.ResetBoth != nil {
		obj, err := o.ResetBoth.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetBothVal = obj
	}
	var blockIpVal *RulesActionBlockIp
	if o.BlockIp != nil {
		obj, err := o.BlockIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		blockIpVal = obj
	}

	result := &RulesAction{
		Default:     defaultVal,
		Allow:       allowVal,
		Alert:       alertVal,
		Drop:        dropVal,
		ResetClient: resetClientVal,
		ResetServer: resetServerVal,
		ResetBoth:   resetBothVal,
		BlockIp:     blockIpVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *rulesActionDefaultXml_11_0_2) MarshalFromObject(s RulesActionDefault) {
	o.Misc = s.Misc
}

func (o rulesActionDefaultXml_11_0_2) UnmarshalToObject() (*RulesActionDefault, error) {

	result := &RulesActionDefault{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *rulesActionAllowXml_11_0_2) MarshalFromObject(s RulesActionAllow) {
	o.Misc = s.Misc
}

func (o rulesActionAllowXml_11_0_2) UnmarshalToObject() (*RulesActionAllow, error) {

	result := &RulesActionAllow{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *rulesActionAlertXml_11_0_2) MarshalFromObject(s RulesActionAlert) {
	o.Misc = s.Misc
}

func (o rulesActionAlertXml_11_0_2) UnmarshalToObject() (*RulesActionAlert, error) {

	result := &RulesActionAlert{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *rulesActionDropXml_11_0_2) MarshalFromObject(s RulesActionDrop) {
	o.Misc = s.Misc
}

func (o rulesActionDropXml_11_0_2) UnmarshalToObject() (*RulesActionDrop, error) {

	result := &RulesActionDrop{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *rulesActionResetClientXml_11_0_2) MarshalFromObject(s RulesActionResetClient) {
	o.Misc = s.Misc
}

func (o rulesActionResetClientXml_11_0_2) UnmarshalToObject() (*RulesActionResetClient, error) {

	result := &RulesActionResetClient{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *rulesActionResetServerXml_11_0_2) MarshalFromObject(s RulesActionResetServer) {
	o.Misc = s.Misc
}

func (o rulesActionResetServerXml_11_0_2) UnmarshalToObject() (*RulesActionResetServer, error) {

	result := &RulesActionResetServer{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *rulesActionResetBothXml_11_0_2) MarshalFromObject(s RulesActionResetBoth) {
	o.Misc = s.Misc
}

func (o rulesActionResetBothXml_11_0_2) UnmarshalToObject() (*RulesActionResetBoth, error) {

	result := &RulesActionResetBoth{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *rulesActionBlockIpXml_11_0_2) MarshalFromObject(s RulesActionBlockIp) {
	o.TrackBy = s.TrackBy
	o.Duration = s.Duration
	o.Misc = s.Misc
}

func (o rulesActionBlockIpXml_11_0_2) UnmarshalToObject() (*RulesActionBlockIp, error) {

	result := &RulesActionBlockIp{
		TrackBy:  o.TrackBy,
		Duration: o.Duration,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *threatExceptionXml_11_0_2) MarshalFromObject(s ThreatException) {
	o.Name = s.Name
	o.PacketCapture = s.PacketCapture
	if s.Action != nil {
		var obj threatExceptionActionXml_11_0_2
		obj.MarshalFromObject(*s.Action)
		o.Action = &obj
	}
	if s.ExemptIp != nil {
		var objs []threatExceptionExemptIpXml_11_0_2
		for _, elt := range s.ExemptIp {
			var obj threatExceptionExemptIpXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ExemptIp = &threatExceptionExemptIpContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o threatExceptionXml_11_0_2) UnmarshalToObject() (*ThreatException, error) {
	var actionVal *ThreatExceptionAction
	if o.Action != nil {
		obj, err := o.Action.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		actionVal = obj
	}
	var exemptIpVal []ThreatExceptionExemptIp
	if o.ExemptIp != nil {
		for _, elt := range o.ExemptIp.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			exemptIpVal = append(exemptIpVal, *obj)
		}
	}

	result := &ThreatException{
		Name:          o.Name,
		PacketCapture: o.PacketCapture,
		Action:        actionVal,
		ExemptIp:      exemptIpVal,
		Misc:          o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionXml_11_0_2) MarshalFromObject(s ThreatExceptionAction) {
	if s.Default != nil {
		var obj threatExceptionActionDefaultXml_11_0_2
		obj.MarshalFromObject(*s.Default)
		o.Default = &obj
	}
	if s.Allow != nil {
		var obj threatExceptionActionAllowXml_11_0_2
		obj.MarshalFromObject(*s.Allow)
		o.Allow = &obj
	}
	if s.Alert != nil {
		var obj threatExceptionActionAlertXml_11_0_2
		obj.MarshalFromObject(*s.Alert)
		o.Alert = &obj
	}
	if s.Drop != nil {
		var obj threatExceptionActionDropXml_11_0_2
		obj.MarshalFromObject(*s.Drop)
		o.Drop = &obj
	}
	if s.ResetBoth != nil {
		var obj threatExceptionActionResetBothXml_11_0_2
		obj.MarshalFromObject(*s.ResetBoth)
		o.ResetBoth = &obj
	}
	if s.ResetClient != nil {
		var obj threatExceptionActionResetClientXml_11_0_2
		obj.MarshalFromObject(*s.ResetClient)
		o.ResetClient = &obj
	}
	if s.ResetServer != nil {
		var obj threatExceptionActionResetServerXml_11_0_2
		obj.MarshalFromObject(*s.ResetServer)
		o.ResetServer = &obj
	}
	if s.BlockIp != nil {
		var obj threatExceptionActionBlockIpXml_11_0_2
		obj.MarshalFromObject(*s.BlockIp)
		o.BlockIp = &obj
	}
	o.Misc = s.Misc
}

func (o threatExceptionActionXml_11_0_2) UnmarshalToObject() (*ThreatExceptionAction, error) {
	var defaultVal *ThreatExceptionActionDefault
	if o.Default != nil {
		obj, err := o.Default.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		defaultVal = obj
	}
	var allowVal *ThreatExceptionActionAllow
	if o.Allow != nil {
		obj, err := o.Allow.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allowVal = obj
	}
	var alertVal *ThreatExceptionActionAlert
	if o.Alert != nil {
		obj, err := o.Alert.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		alertVal = obj
	}
	var dropVal *ThreatExceptionActionDrop
	if o.Drop != nil {
		obj, err := o.Drop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		dropVal = obj
	}
	var resetBothVal *ThreatExceptionActionResetBoth
	if o.ResetBoth != nil {
		obj, err := o.ResetBoth.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetBothVal = obj
	}
	var resetClientVal *ThreatExceptionActionResetClient
	if o.ResetClient != nil {
		obj, err := o.ResetClient.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetClientVal = obj
	}
	var resetServerVal *ThreatExceptionActionResetServer
	if o.ResetServer != nil {
		obj, err := o.ResetServer.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		resetServerVal = obj
	}
	var blockIpVal *ThreatExceptionActionBlockIp
	if o.BlockIp != nil {
		obj, err := o.BlockIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		blockIpVal = obj
	}

	result := &ThreatExceptionAction{
		Default:     defaultVal,
		Allow:       allowVal,
		Alert:       alertVal,
		Drop:        dropVal,
		ResetBoth:   resetBothVal,
		ResetClient: resetClientVal,
		ResetServer: resetServerVal,
		BlockIp:     blockIpVal,
		Misc:        o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionDefaultXml_11_0_2) MarshalFromObject(s ThreatExceptionActionDefault) {
	o.Misc = s.Misc
}

func (o threatExceptionActionDefaultXml_11_0_2) UnmarshalToObject() (*ThreatExceptionActionDefault, error) {

	result := &ThreatExceptionActionDefault{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionAllowXml_11_0_2) MarshalFromObject(s ThreatExceptionActionAllow) {
	o.Misc = s.Misc
}

func (o threatExceptionActionAllowXml_11_0_2) UnmarshalToObject() (*ThreatExceptionActionAllow, error) {

	result := &ThreatExceptionActionAllow{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionAlertXml_11_0_2) MarshalFromObject(s ThreatExceptionActionAlert) {
	o.Misc = s.Misc
}

func (o threatExceptionActionAlertXml_11_0_2) UnmarshalToObject() (*ThreatExceptionActionAlert, error) {

	result := &ThreatExceptionActionAlert{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionDropXml_11_0_2) MarshalFromObject(s ThreatExceptionActionDrop) {
	o.Misc = s.Misc
}

func (o threatExceptionActionDropXml_11_0_2) UnmarshalToObject() (*ThreatExceptionActionDrop, error) {

	result := &ThreatExceptionActionDrop{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionResetBothXml_11_0_2) MarshalFromObject(s ThreatExceptionActionResetBoth) {
	o.Misc = s.Misc
}

func (o threatExceptionActionResetBothXml_11_0_2) UnmarshalToObject() (*ThreatExceptionActionResetBoth, error) {

	result := &ThreatExceptionActionResetBoth{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionResetClientXml_11_0_2) MarshalFromObject(s ThreatExceptionActionResetClient) {
	o.Misc = s.Misc
}

func (o threatExceptionActionResetClientXml_11_0_2) UnmarshalToObject() (*ThreatExceptionActionResetClient, error) {

	result := &ThreatExceptionActionResetClient{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionResetServerXml_11_0_2) MarshalFromObject(s ThreatExceptionActionResetServer) {
	o.Misc = s.Misc
}

func (o threatExceptionActionResetServerXml_11_0_2) UnmarshalToObject() (*ThreatExceptionActionResetServer, error) {

	result := &ThreatExceptionActionResetServer{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *threatExceptionActionBlockIpXml_11_0_2) MarshalFromObject(s ThreatExceptionActionBlockIp) {
	o.TrackBy = s.TrackBy
	o.Duration = s.Duration
	o.Misc = s.Misc
}

func (o threatExceptionActionBlockIpXml_11_0_2) UnmarshalToObject() (*ThreatExceptionActionBlockIp, error) {

	result := &ThreatExceptionActionBlockIp{
		TrackBy:  o.TrackBy,
		Duration: o.Duration,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *threatExceptionExemptIpXml_11_0_2) MarshalFromObject(s ThreatExceptionExemptIp) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o threatExceptionExemptIpXml_11_0_2) UnmarshalToObject() (*ThreatExceptionExemptIp, error) {

	result := &ThreatExceptionExemptIp{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
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
	if !o.BotnetDomains.matches(other.BotnetDomains) {
		return false
	}
	if !util.BoolsMatch(o.CloudInlineAnalysis, other.CloudInlineAnalysis) {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}
	if !util.StringsMatch(o.DisableOverride, other.DisableOverride) {
		return false
	}
	if !util.OrderedListsMatch[string](o.InlineExceptionEdlUrl, other.InlineExceptionEdlUrl) {
		return false
	}
	if !util.OrderedListsMatch[string](o.InlineExceptionIpAddress, other.InlineExceptionIpAddress) {
		return false
	}
	if len(o.MicaEngineSpywareEnabled) != len(other.MicaEngineSpywareEnabled) {
		return false
	}
	for idx := range o.MicaEngineSpywareEnabled {
		if !o.MicaEngineSpywareEnabled[idx].matches(&other.MicaEngineSpywareEnabled[idx]) {
			return false
		}
	}
	if len(o.Rules) != len(other.Rules) {
		return false
	}
	for idx := range o.Rules {
		if !o.Rules[idx].matches(&other.Rules[idx]) {
			return false
		}
	}
	if len(o.ThreatException) != len(other.ThreatException) {
		return false
	}
	for idx := range o.ThreatException {
		if !o.ThreatException[idx].matches(&other.ThreatException[idx]) {
			return false
		}
	}

	return true
}

func (o *BotnetDomains) matches(other *BotnetDomains) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.DnsSecurityCategories) != len(other.DnsSecurityCategories) {
		return false
	}
	for idx := range o.DnsSecurityCategories {
		if !o.DnsSecurityCategories[idx].matches(&other.DnsSecurityCategories[idx]) {
			return false
		}
	}
	if len(o.Lists) != len(other.Lists) {
		return false
	}
	for idx := range o.Lists {
		if !o.Lists[idx].matches(&other.Lists[idx]) {
			return false
		}
	}
	if !o.RtypeAction.matches(other.RtypeAction) {
		return false
	}
	if !o.Sinkhole.matches(other.Sinkhole) {
		return false
	}
	if len(o.ThreatException) != len(other.ThreatException) {
		return false
	}
	for idx := range o.ThreatException {
		if !o.ThreatException[idx].matches(&other.ThreatException[idx]) {
			return false
		}
	}
	if len(o.Whitelist) != len(other.Whitelist) {
		return false
	}
	for idx := range o.Whitelist {
		if !o.Whitelist[idx].matches(&other.Whitelist[idx]) {
			return false
		}
	}

	return true
}

func (o *BotnetDomainsDnsSecurityCategories) matches(other *BotnetDomainsDnsSecurityCategories) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Action, other.Action) {
		return false
	}
	if !util.StringsMatch(o.LogLevel, other.LogLevel) {
		return false
	}
	if !util.StringsMatch(o.PacketCapture, other.PacketCapture) {
		return false
	}

	return true
}

func (o *BotnetDomainsLists) matches(other *BotnetDomainsLists) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !o.Action.matches(other.Action) {
		return false
	}
	if !util.StringsMatch(o.PacketCapture, other.PacketCapture) {
		return false
	}

	return true
}

func (o *BotnetDomainsListsAction) matches(other *BotnetDomainsListsAction) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Alert.matches(other.Alert) {
		return false
	}
	if !o.Allow.matches(other.Allow) {
		return false
	}
	if !o.Block.matches(other.Block) {
		return false
	}
	if !o.Sinkhole.matches(other.Sinkhole) {
		return false
	}

	return true
}

func (o *BotnetDomainsListsActionAlert) matches(other *BotnetDomainsListsActionAlert) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *BotnetDomainsListsActionAllow) matches(other *BotnetDomainsListsActionAllow) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *BotnetDomainsListsActionBlock) matches(other *BotnetDomainsListsActionBlock) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *BotnetDomainsListsActionSinkhole) matches(other *BotnetDomainsListsActionSinkhole) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *BotnetDomainsRtypeAction) matches(other *BotnetDomainsRtypeAction) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Any, other.Any) {
		return false
	}
	if !util.StringsMatch(o.Https, other.Https) {
		return false
	}
	if !util.StringsMatch(o.Svcb, other.Svcb) {
		return false
	}

	return true
}

func (o *BotnetDomainsSinkhole) matches(other *BotnetDomainsSinkhole) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ipv4Address, other.Ipv4Address) {
		return false
	}
	if !util.StringsMatch(o.Ipv6Address, other.Ipv6Address) {
		return false
	}

	return true
}

func (o *BotnetDomainsThreatException) matches(other *BotnetDomainsThreatException) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}

	return true
}

func (o *BotnetDomainsWhitelist) matches(other *BotnetDomainsWhitelist) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Description, other.Description) {
		return false
	}

	return true
}

func (o *MicaEngineSpywareEnabled) matches(other *MicaEngineSpywareEnabled) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.InlinePolicyAction, other.InlinePolicyAction) {
		return false
	}

	return true
}

func (o *Rules) matches(other *Rules) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.ThreatName, other.ThreatName) {
		return false
	}
	if !util.StringsMatch(o.Category, other.Category) {
		return false
	}
	if !util.StringsMatch(o.PacketCapture, other.PacketCapture) {
		return false
	}
	if !util.OrderedListsMatch[string](o.Severity, other.Severity) {
		return false
	}
	if !o.Action.matches(other.Action) {
		return false
	}

	return true
}

func (o *RulesAction) matches(other *RulesAction) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Default.matches(other.Default) {
		return false
	}
	if !o.Allow.matches(other.Allow) {
		return false
	}
	if !o.Alert.matches(other.Alert) {
		return false
	}
	if !o.Drop.matches(other.Drop) {
		return false
	}
	if !o.ResetClient.matches(other.ResetClient) {
		return false
	}
	if !o.ResetServer.matches(other.ResetServer) {
		return false
	}
	if !o.ResetBoth.matches(other.ResetBoth) {
		return false
	}
	if !o.BlockIp.matches(other.BlockIp) {
		return false
	}

	return true
}

func (o *RulesActionDefault) matches(other *RulesActionDefault) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *RulesActionAllow) matches(other *RulesActionAllow) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *RulesActionAlert) matches(other *RulesActionAlert) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *RulesActionDrop) matches(other *RulesActionDrop) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *RulesActionResetClient) matches(other *RulesActionResetClient) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *RulesActionResetServer) matches(other *RulesActionResetServer) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *RulesActionResetBoth) matches(other *RulesActionResetBoth) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *RulesActionBlockIp) matches(other *RulesActionBlockIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.TrackBy, other.TrackBy) {
		return false
	}
	if !util.Ints64Match(o.Duration, other.Duration) {
		return false
	}

	return true
}

func (o *ThreatException) matches(other *ThreatException) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.PacketCapture, other.PacketCapture) {
		return false
	}
	if !o.Action.matches(other.Action) {
		return false
	}
	if len(o.ExemptIp) != len(other.ExemptIp) {
		return false
	}
	for idx := range o.ExemptIp {
		if !o.ExemptIp[idx].matches(&other.ExemptIp[idx]) {
			return false
		}
	}

	return true
}

func (o *ThreatExceptionAction) matches(other *ThreatExceptionAction) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Default.matches(other.Default) {
		return false
	}
	if !o.Allow.matches(other.Allow) {
		return false
	}
	if !o.Alert.matches(other.Alert) {
		return false
	}
	if !o.Drop.matches(other.Drop) {
		return false
	}
	if !o.ResetBoth.matches(other.ResetBoth) {
		return false
	}
	if !o.ResetClient.matches(other.ResetClient) {
		return false
	}
	if !o.ResetServer.matches(other.ResetServer) {
		return false
	}
	if !o.BlockIp.matches(other.BlockIp) {
		return false
	}

	return true
}

func (o *ThreatExceptionActionDefault) matches(other *ThreatExceptionActionDefault) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ThreatExceptionActionAllow) matches(other *ThreatExceptionActionAllow) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ThreatExceptionActionAlert) matches(other *ThreatExceptionActionAlert) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ThreatExceptionActionDrop) matches(other *ThreatExceptionActionDrop) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ThreatExceptionActionResetBoth) matches(other *ThreatExceptionActionResetBoth) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ThreatExceptionActionResetClient) matches(other *ThreatExceptionActionResetClient) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ThreatExceptionActionResetServer) matches(other *ThreatExceptionActionResetServer) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ThreatExceptionActionBlockIp) matches(other *ThreatExceptionActionBlockIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.TrackBy, other.TrackBy) {
		return false
	}
	if !util.Ints64Match(o.Duration, other.Duration) {
		return false
	}

	return true
}

func (o *ThreatExceptionExemptIp) matches(other *ThreatExceptionExemptIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
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
