package addressfamily

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
	suffix = []string{"network", "routing-profile", "bgp", "address-family-profile", "$name"}
)

type Entry struct {
	Name           string
	Ipv4           *Ipv4
	Ipv6           *Ipv6
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4 struct {
	Multicast      *Ipv4Multicast
	Unicast        *Ipv4Unicast
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4Multicast struct {
	AddPath                    *Ipv4MulticastAddPath
	AllowasIn                  *Ipv4MulticastAllowasIn
	AsOverride                 *bool
	DefaultOriginate           *bool
	DefaultOriginateMap        *string
	Enable                     *bool
	MaximumPrefix              *Ipv4MulticastMaximumPrefix
	NextHop                    *Ipv4MulticastNextHop
	Orf                        *Ipv4MulticastOrf
	RemovePrivateAS            *Ipv4MulticastRemovePrivateAS
	RouteReflectorClient       *bool
	SendCommunity              *Ipv4MulticastSendCommunity
	SoftReconfigWithStoredInfo *bool
	Misc                       []generic.Xml
	MiscAttributes             []xml.Attr
}
type Ipv4MulticastAddPath struct {
	TxAllPaths      *bool
	TxBestpathPerAS *bool
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type Ipv4MulticastAllowasIn struct {
	Occurrence     *int64
	Origin         *Ipv4MulticastAllowasInOrigin
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastAllowasInOrigin struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastMaximumPrefix struct {
	Action         *Ipv4MulticastMaximumPrefixAction
	NumPrefixes    *int64
	Threshold      *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastMaximumPrefixAction struct {
	Restart        *Ipv4MulticastMaximumPrefixActionRestart
	WarningOnly    *Ipv4MulticastMaximumPrefixActionWarningOnly
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastMaximumPrefixActionRestart struct {
	Interval       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastMaximumPrefixActionWarningOnly struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastNextHop struct {
	Self           *Ipv4MulticastNextHopSelf
	SelfForce      *Ipv4MulticastNextHopSelfForce
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastNextHopSelf struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastNextHopSelfForce struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastOrf struct {
	OrfPrefixList  *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastRemovePrivateAS struct {
	All            *Ipv4MulticastRemovePrivateASAll
	ReplaceAS      *Ipv4MulticastRemovePrivateASReplaceAS
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastRemovePrivateASAll struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastRemovePrivateASReplaceAS struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastSendCommunity struct {
	All            *Ipv4MulticastSendCommunityAll
	Both           *Ipv4MulticastSendCommunityBoth
	Extended       *Ipv4MulticastSendCommunityExtended
	Large          *Ipv4MulticastSendCommunityLarge
	Standard       *Ipv4MulticastSendCommunityStandard
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastSendCommunityAll struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastSendCommunityBoth struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastSendCommunityExtended struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastSendCommunityLarge struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4MulticastSendCommunityStandard struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4Unicast struct {
	AddPath                    *Ipv4UnicastAddPath
	AllowasIn                  *Ipv4UnicastAllowasIn
	AsOverride                 *bool
	DefaultOriginate           *bool
	DefaultOriginateMap        *string
	Enable                     *bool
	MaximumPrefix              *Ipv4UnicastMaximumPrefix
	NextHop                    *Ipv4UnicastNextHop
	Orf                        *Ipv4UnicastOrf
	RemovePrivateAS            *Ipv4UnicastRemovePrivateAS
	RouteReflectorClient       *bool
	SendCommunity              *Ipv4UnicastSendCommunity
	SoftReconfigWithStoredInfo *bool
	Misc                       []generic.Xml
	MiscAttributes             []xml.Attr
}
type Ipv4UnicastAddPath struct {
	TxAllPaths      *bool
	TxBestpathPerAS *bool
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type Ipv4UnicastAllowasIn struct {
	Occurrence     *int64
	Origin         *Ipv4UnicastAllowasInOrigin
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastAllowasInOrigin struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastMaximumPrefix struct {
	Action         *Ipv4UnicastMaximumPrefixAction
	NumPrefixes    *int64
	Threshold      *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastMaximumPrefixAction struct {
	Restart        *Ipv4UnicastMaximumPrefixActionRestart
	WarningOnly    *Ipv4UnicastMaximumPrefixActionWarningOnly
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastMaximumPrefixActionRestart struct {
	Interval       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastMaximumPrefixActionWarningOnly struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastNextHop struct {
	Self           *Ipv4UnicastNextHopSelf
	SelfForce      *Ipv4UnicastNextHopSelfForce
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastNextHopSelf struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastNextHopSelfForce struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastOrf struct {
	OrfPrefixList  *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastRemovePrivateAS struct {
	All            *Ipv4UnicastRemovePrivateASAll
	ReplaceAS      *Ipv4UnicastRemovePrivateASReplaceAS
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastRemovePrivateASAll struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastRemovePrivateASReplaceAS struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastSendCommunity struct {
	All            *Ipv4UnicastSendCommunityAll
	Both           *Ipv4UnicastSendCommunityBoth
	Extended       *Ipv4UnicastSendCommunityExtended
	Large          *Ipv4UnicastSendCommunityLarge
	Standard       *Ipv4UnicastSendCommunityStandard
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastSendCommunityAll struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastSendCommunityBoth struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastSendCommunityExtended struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastSendCommunityLarge struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv4UnicastSendCommunityStandard struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6 struct {
	Unicast        *Ipv6Unicast
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6Unicast struct {
	AddPath                    *Ipv6UnicastAddPath
	AllowasIn                  *Ipv6UnicastAllowasIn
	AsOverride                 *bool
	DefaultOriginate           *bool
	DefaultOriginateMap        *string
	Enable                     *bool
	MaximumPrefix              *Ipv6UnicastMaximumPrefix
	NextHop                    *Ipv6UnicastNextHop
	Orf                        *Ipv6UnicastOrf
	RemovePrivateAS            *Ipv6UnicastRemovePrivateAS
	RouteReflectorClient       *bool
	SendCommunity              *Ipv6UnicastSendCommunity
	SoftReconfigWithStoredInfo *bool
	Misc                       []generic.Xml
	MiscAttributes             []xml.Attr
}
type Ipv6UnicastAddPath struct {
	TxAllPaths      *bool
	TxBestpathPerAS *bool
	Misc            []generic.Xml
	MiscAttributes  []xml.Attr
}
type Ipv6UnicastAllowasIn struct {
	Occurrence     *int64
	Origin         *Ipv6UnicastAllowasInOrigin
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastAllowasInOrigin struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastMaximumPrefix struct {
	Action         *Ipv6UnicastMaximumPrefixAction
	NumPrefixes    *int64
	Threshold      *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastMaximumPrefixAction struct {
	Restart        *Ipv6UnicastMaximumPrefixActionRestart
	WarningOnly    *Ipv6UnicastMaximumPrefixActionWarningOnly
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastMaximumPrefixActionRestart struct {
	Interval       *int64
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastMaximumPrefixActionWarningOnly struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastNextHop struct {
	Self           *Ipv6UnicastNextHopSelf
	SelfForce      *Ipv6UnicastNextHopSelfForce
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastNextHopSelf struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastNextHopSelfForce struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastOrf struct {
	OrfPrefixList  *string
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastRemovePrivateAS struct {
	All            *Ipv6UnicastRemovePrivateASAll
	ReplaceAS      *Ipv6UnicastRemovePrivateASReplaceAS
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastRemovePrivateASAll struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastRemovePrivateASReplaceAS struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastSendCommunity struct {
	All            *Ipv6UnicastSendCommunityAll
	Both           *Ipv6UnicastSendCommunityBoth
	Extended       *Ipv6UnicastSendCommunityExtended
	Large          *Ipv6UnicastSendCommunityLarge
	Standard       *Ipv6UnicastSendCommunityStandard
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastSendCommunityAll struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastSendCommunityBoth struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastSendCommunityExtended struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastSendCommunityLarge struct {
	Misc           []generic.Xml
	MiscAttributes []xml.Attr
}
type Ipv6UnicastSendCommunityStandard struct {
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
	XMLName        xml.Name      `xml:"entry"`
	Name           string        `xml:"name,attr"`
	Ipv4           *ipv4Xml      `xml:"ipv4,omitempty"`
	Ipv6           *ipv6Xml      `xml:"ipv6,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4Xml struct {
	Multicast      *ipv4MulticastXml `xml:"multicast,omitempty"`
	Unicast        *ipv4UnicastXml   `xml:"unicast,omitempty"`
	Misc           []generic.Xml     `xml:",any"`
	MiscAttributes []xml.Attr        `xml:",any,attr"`
}
type ipv4MulticastXml struct {
	AddPath                    *ipv4MulticastAddPathXml         `xml:"add-path,omitempty"`
	AllowasIn                  *ipv4MulticastAllowasInXml       `xml:"allowas-in,omitempty"`
	AsOverride                 *string                          `xml:"as-override,omitempty"`
	DefaultOriginate           *string                          `xml:"default-originate,omitempty"`
	DefaultOriginateMap        *string                          `xml:"default-originate-map,omitempty"`
	Enable                     *string                          `xml:"enable,omitempty"`
	MaximumPrefix              *ipv4MulticastMaximumPrefixXml   `xml:"maximum-prefix,omitempty"`
	NextHop                    *ipv4MulticastNextHopXml         `xml:"next-hop,omitempty"`
	Orf                        *ipv4MulticastOrfXml             `xml:"orf,omitempty"`
	RemovePrivateAS            *ipv4MulticastRemovePrivateASXml `xml:"remove-private-AS,omitempty"`
	RouteReflectorClient       *string                          `xml:"route-reflector-client,omitempty"`
	SendCommunity              *ipv4MulticastSendCommunityXml   `xml:"send-community,omitempty"`
	SoftReconfigWithStoredInfo *string                          `xml:"soft-reconfig-with-stored-info,omitempty"`
	Misc                       []generic.Xml                    `xml:",any"`
	MiscAttributes             []xml.Attr                       `xml:",any,attr"`
}
type ipv4MulticastAddPathXml struct {
	TxAllPaths      *string       `xml:"tx-all-paths,omitempty"`
	TxBestpathPerAS *string       `xml:"tx-bestpath-per-AS,omitempty"`
	Misc            []generic.Xml `xml:",any"`
	MiscAttributes  []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastAllowasInXml struct {
	Occurrence     *int64                           `xml:"occurrence,omitempty"`
	Origin         *ipv4MulticastAllowasInOriginXml `xml:"origin,omitempty"`
	Misc           []generic.Xml                    `xml:",any"`
	MiscAttributes []xml.Attr                       `xml:",any,attr"`
}
type ipv4MulticastAllowasInOriginXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastMaximumPrefixXml struct {
	Action         *ipv4MulticastMaximumPrefixActionXml `xml:"action,omitempty"`
	NumPrefixes    *int64                               `xml:"num_prefixes,omitempty"`
	Threshold      *int64                               `xml:"threshold,omitempty"`
	Misc           []generic.Xml                        `xml:",any"`
	MiscAttributes []xml.Attr                           `xml:",any,attr"`
}
type ipv4MulticastMaximumPrefixActionXml struct {
	Restart        *ipv4MulticastMaximumPrefixActionRestartXml     `xml:"restart,omitempty"`
	WarningOnly    *ipv4MulticastMaximumPrefixActionWarningOnlyXml `xml:"warning-only,omitempty"`
	Misc           []generic.Xml                                   `xml:",any"`
	MiscAttributes []xml.Attr                                      `xml:",any,attr"`
}
type ipv4MulticastMaximumPrefixActionRestartXml struct {
	Interval       *int64        `xml:"interval,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastMaximumPrefixActionWarningOnlyXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastNextHopXml struct {
	Self           *ipv4MulticastNextHopSelfXml      `xml:"self,omitempty"`
	SelfForce      *ipv4MulticastNextHopSelfForceXml `xml:"self-force,omitempty"`
	Misc           []generic.Xml                     `xml:",any"`
	MiscAttributes []xml.Attr                        `xml:",any,attr"`
}
type ipv4MulticastNextHopSelfXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastNextHopSelfForceXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastOrfXml struct {
	OrfPrefixList  *string       `xml:"orf-prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastRemovePrivateASXml struct {
	All            *ipv4MulticastRemovePrivateASAllXml       `xml:"all,omitempty"`
	ReplaceAS      *ipv4MulticastRemovePrivateASReplaceASXml `xml:"replace-AS,omitempty"`
	Misc           []generic.Xml                             `xml:",any"`
	MiscAttributes []xml.Attr                                `xml:",any,attr"`
}
type ipv4MulticastRemovePrivateASAllXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastRemovePrivateASReplaceASXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastSendCommunityXml struct {
	All            *ipv4MulticastSendCommunityAllXml      `xml:"all,omitempty"`
	Both           *ipv4MulticastSendCommunityBothXml     `xml:"both,omitempty"`
	Extended       *ipv4MulticastSendCommunityExtendedXml `xml:"extended,omitempty"`
	Large          *ipv4MulticastSendCommunityLargeXml    `xml:"large,omitempty"`
	Standard       *ipv4MulticastSendCommunityStandardXml `xml:"standard,omitempty"`
	Misc           []generic.Xml                          `xml:",any"`
	MiscAttributes []xml.Attr                             `xml:",any,attr"`
}
type ipv4MulticastSendCommunityAllXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastSendCommunityBothXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastSendCommunityExtendedXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastSendCommunityLargeXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4MulticastSendCommunityStandardXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastXml struct {
	AddPath                    *ipv4UnicastAddPathXml         `xml:"add-path,omitempty"`
	AllowasIn                  *ipv4UnicastAllowasInXml       `xml:"allowas-in,omitempty"`
	AsOverride                 *string                        `xml:"as-override,omitempty"`
	DefaultOriginate           *string                        `xml:"default-originate,omitempty"`
	DefaultOriginateMap        *string                        `xml:"default-originate-map,omitempty"`
	Enable                     *string                        `xml:"enable,omitempty"`
	MaximumPrefix              *ipv4UnicastMaximumPrefixXml   `xml:"maximum-prefix,omitempty"`
	NextHop                    *ipv4UnicastNextHopXml         `xml:"next-hop,omitempty"`
	Orf                        *ipv4UnicastOrfXml             `xml:"orf,omitempty"`
	RemovePrivateAS            *ipv4UnicastRemovePrivateASXml `xml:"remove-private-AS,omitempty"`
	RouteReflectorClient       *string                        `xml:"route-reflector-client,omitempty"`
	SendCommunity              *ipv4UnicastSendCommunityXml   `xml:"send-community,omitempty"`
	SoftReconfigWithStoredInfo *string                        `xml:"soft-reconfig-with-stored-info,omitempty"`
	Misc                       []generic.Xml                  `xml:",any"`
	MiscAttributes             []xml.Attr                     `xml:",any,attr"`
}
type ipv4UnicastAddPathXml struct {
	TxAllPaths      *string       `xml:"tx-all-paths,omitempty"`
	TxBestpathPerAS *string       `xml:"tx-bestpath-per-AS,omitempty"`
	Misc            []generic.Xml `xml:",any"`
	MiscAttributes  []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastAllowasInXml struct {
	Occurrence     *int64                         `xml:"occurrence,omitempty"`
	Origin         *ipv4UnicastAllowasInOriginXml `xml:"origin,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type ipv4UnicastAllowasInOriginXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastMaximumPrefixXml struct {
	Action         *ipv4UnicastMaximumPrefixActionXml `xml:"action,omitempty"`
	NumPrefixes    *int64                             `xml:"num_prefixes,omitempty"`
	Threshold      *int64                             `xml:"threshold,omitempty"`
	Misc           []generic.Xml                      `xml:",any"`
	MiscAttributes []xml.Attr                         `xml:",any,attr"`
}
type ipv4UnicastMaximumPrefixActionXml struct {
	Restart        *ipv4UnicastMaximumPrefixActionRestartXml     `xml:"restart,omitempty"`
	WarningOnly    *ipv4UnicastMaximumPrefixActionWarningOnlyXml `xml:"warning-only,omitempty"`
	Misc           []generic.Xml                                 `xml:",any"`
	MiscAttributes []xml.Attr                                    `xml:",any,attr"`
}
type ipv4UnicastMaximumPrefixActionRestartXml struct {
	Interval       *int64        `xml:"interval,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastMaximumPrefixActionWarningOnlyXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastNextHopXml struct {
	Self           *ipv4UnicastNextHopSelfXml      `xml:"self,omitempty"`
	SelfForce      *ipv4UnicastNextHopSelfForceXml `xml:"self-force,omitempty"`
	Misc           []generic.Xml                   `xml:",any"`
	MiscAttributes []xml.Attr                      `xml:",any,attr"`
}
type ipv4UnicastNextHopSelfXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastNextHopSelfForceXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastOrfXml struct {
	OrfPrefixList  *string       `xml:"orf-prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastRemovePrivateASXml struct {
	All            *ipv4UnicastRemovePrivateASAllXml       `xml:"all,omitempty"`
	ReplaceAS      *ipv4UnicastRemovePrivateASReplaceASXml `xml:"replace-AS,omitempty"`
	Misc           []generic.Xml                           `xml:",any"`
	MiscAttributes []xml.Attr                              `xml:",any,attr"`
}
type ipv4UnicastRemovePrivateASAllXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastRemovePrivateASReplaceASXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastSendCommunityXml struct {
	All            *ipv4UnicastSendCommunityAllXml      `xml:"all,omitempty"`
	Both           *ipv4UnicastSendCommunityBothXml     `xml:"both,omitempty"`
	Extended       *ipv4UnicastSendCommunityExtendedXml `xml:"extended,omitempty"`
	Large          *ipv4UnicastSendCommunityLargeXml    `xml:"large,omitempty"`
	Standard       *ipv4UnicastSendCommunityStandardXml `xml:"standard,omitempty"`
	Misc           []generic.Xml                        `xml:",any"`
	MiscAttributes []xml.Attr                           `xml:",any,attr"`
}
type ipv4UnicastSendCommunityAllXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastSendCommunityBothXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastSendCommunityExtendedXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastSendCommunityLargeXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv4UnicastSendCommunityStandardXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6Xml struct {
	Unicast        *ipv6UnicastXml `xml:"unicast,omitempty"`
	Misc           []generic.Xml   `xml:",any"`
	MiscAttributes []xml.Attr      `xml:",any,attr"`
}
type ipv6UnicastXml struct {
	AddPath                    *ipv6UnicastAddPathXml         `xml:"add-path,omitempty"`
	AllowasIn                  *ipv6UnicastAllowasInXml       `xml:"allowas-in,omitempty"`
	AsOverride                 *string                        `xml:"as-override,omitempty"`
	DefaultOriginate           *string                        `xml:"default-originate,omitempty"`
	DefaultOriginateMap        *string                        `xml:"default-originate-map,omitempty"`
	Enable                     *string                        `xml:"enable,omitempty"`
	MaximumPrefix              *ipv6UnicastMaximumPrefixXml   `xml:"maximum-prefix,omitempty"`
	NextHop                    *ipv6UnicastNextHopXml         `xml:"next-hop,omitempty"`
	Orf                        *ipv6UnicastOrfXml             `xml:"orf,omitempty"`
	RemovePrivateAS            *ipv6UnicastRemovePrivateASXml `xml:"remove-private-AS,omitempty"`
	RouteReflectorClient       *string                        `xml:"route-reflector-client,omitempty"`
	SendCommunity              *ipv6UnicastSendCommunityXml   `xml:"send-community,omitempty"`
	SoftReconfigWithStoredInfo *string                        `xml:"soft-reconfig-with-stored-info,omitempty"`
	Misc                       []generic.Xml                  `xml:",any"`
	MiscAttributes             []xml.Attr                     `xml:",any,attr"`
}
type ipv6UnicastAddPathXml struct {
	TxAllPaths      *string       `xml:"tx-all-paths,omitempty"`
	TxBestpathPerAS *string       `xml:"tx-bestpath-per-AS,omitempty"`
	Misc            []generic.Xml `xml:",any"`
	MiscAttributes  []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastAllowasInXml struct {
	Occurrence     *int64                         `xml:"occurrence,omitempty"`
	Origin         *ipv6UnicastAllowasInOriginXml `xml:"origin,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
	MiscAttributes []xml.Attr                     `xml:",any,attr"`
}
type ipv6UnicastAllowasInOriginXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastMaximumPrefixXml struct {
	Action         *ipv6UnicastMaximumPrefixActionXml `xml:"action,omitempty"`
	NumPrefixes    *int64                             `xml:"num_prefixes,omitempty"`
	Threshold      *int64                             `xml:"threshold,omitempty"`
	Misc           []generic.Xml                      `xml:",any"`
	MiscAttributes []xml.Attr                         `xml:",any,attr"`
}
type ipv6UnicastMaximumPrefixActionXml struct {
	Restart        *ipv6UnicastMaximumPrefixActionRestartXml     `xml:"restart,omitempty"`
	WarningOnly    *ipv6UnicastMaximumPrefixActionWarningOnlyXml `xml:"warning-only,omitempty"`
	Misc           []generic.Xml                                 `xml:",any"`
	MiscAttributes []xml.Attr                                    `xml:",any,attr"`
}
type ipv6UnicastMaximumPrefixActionRestartXml struct {
	Interval       *int64        `xml:"interval,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastMaximumPrefixActionWarningOnlyXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastNextHopXml struct {
	Self           *ipv6UnicastNextHopSelfXml      `xml:"self,omitempty"`
	SelfForce      *ipv6UnicastNextHopSelfForceXml `xml:"self-force,omitempty"`
	Misc           []generic.Xml                   `xml:",any"`
	MiscAttributes []xml.Attr                      `xml:",any,attr"`
}
type ipv6UnicastNextHopSelfXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastNextHopSelfForceXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastOrfXml struct {
	OrfPrefixList  *string       `xml:"orf-prefix-list,omitempty"`
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastRemovePrivateASXml struct {
	All            *ipv6UnicastRemovePrivateASAllXml       `xml:"all,omitempty"`
	ReplaceAS      *ipv6UnicastRemovePrivateASReplaceASXml `xml:"replace-AS,omitempty"`
	Misc           []generic.Xml                           `xml:",any"`
	MiscAttributes []xml.Attr                              `xml:",any,attr"`
}
type ipv6UnicastRemovePrivateASAllXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastRemovePrivateASReplaceASXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastSendCommunityXml struct {
	All            *ipv6UnicastSendCommunityAllXml      `xml:"all,omitempty"`
	Both           *ipv6UnicastSendCommunityBothXml     `xml:"both,omitempty"`
	Extended       *ipv6UnicastSendCommunityExtendedXml `xml:"extended,omitempty"`
	Large          *ipv6UnicastSendCommunityLargeXml    `xml:"large,omitempty"`
	Standard       *ipv6UnicastSendCommunityStandardXml `xml:"standard,omitempty"`
	Misc           []generic.Xml                        `xml:",any"`
	MiscAttributes []xml.Attr                           `xml:",any,attr"`
}
type ipv6UnicastSendCommunityAllXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastSendCommunityBothXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastSendCommunityExtendedXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastSendCommunityLargeXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}
type ipv6UnicastSendCommunityStandardXml struct {
	Misc           []generic.Xml `xml:",any"`
	MiscAttributes []xml.Attr    `xml:",any,attr"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	if s.Ipv4 != nil {
		var obj ipv4Xml
		obj.MarshalFromObject(*s.Ipv4)
		o.Ipv4 = &obj
	}
	if s.Ipv6 != nil {
		var obj ipv6Xml
		obj.MarshalFromObject(*s.Ipv6)
		o.Ipv6 = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var ipv4Val *Ipv4
	if o.Ipv4 != nil {
		obj, err := o.Ipv4.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv4Val = obj
	}
	var ipv6Val *Ipv6
	if o.Ipv6 != nil {
		obj, err := o.Ipv6.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipv6Val = obj
	}

	result := &Entry{
		Name:           o.Name,
		Ipv4:           ipv4Val,
		Ipv6:           ipv6Val,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4Xml) MarshalFromObject(s Ipv4) {
	if s.Multicast != nil {
		var obj ipv4MulticastXml
		obj.MarshalFromObject(*s.Multicast)
		o.Multicast = &obj
	}
	if s.Unicast != nil {
		var obj ipv4UnicastXml
		obj.MarshalFromObject(*s.Unicast)
		o.Unicast = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4Xml) UnmarshalToObject() (*Ipv4, error) {
	var multicastVal *Ipv4Multicast
	if o.Multicast != nil {
		obj, err := o.Multicast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		multicastVal = obj
	}
	var unicastVal *Ipv4Unicast
	if o.Unicast != nil {
		obj, err := o.Unicast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		unicastVal = obj
	}

	result := &Ipv4{
		Multicast:      multicastVal,
		Unicast:        unicastVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastXml) MarshalFromObject(s Ipv4Multicast) {
	if s.AddPath != nil {
		var obj ipv4MulticastAddPathXml
		obj.MarshalFromObject(*s.AddPath)
		o.AddPath = &obj
	}
	if s.AllowasIn != nil {
		var obj ipv4MulticastAllowasInXml
		obj.MarshalFromObject(*s.AllowasIn)
		o.AllowasIn = &obj
	}
	o.AsOverride = util.YesNo(s.AsOverride, nil)
	o.DefaultOriginate = util.YesNo(s.DefaultOriginate, nil)
	o.DefaultOriginateMap = s.DefaultOriginateMap
	o.Enable = util.YesNo(s.Enable, nil)
	if s.MaximumPrefix != nil {
		var obj ipv4MulticastMaximumPrefixXml
		obj.MarshalFromObject(*s.MaximumPrefix)
		o.MaximumPrefix = &obj
	}
	if s.NextHop != nil {
		var obj ipv4MulticastNextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	if s.Orf != nil {
		var obj ipv4MulticastOrfXml
		obj.MarshalFromObject(*s.Orf)
		o.Orf = &obj
	}
	if s.RemovePrivateAS != nil {
		var obj ipv4MulticastRemovePrivateASXml
		obj.MarshalFromObject(*s.RemovePrivateAS)
		o.RemovePrivateAS = &obj
	}
	o.RouteReflectorClient = util.YesNo(s.RouteReflectorClient, nil)
	if s.SendCommunity != nil {
		var obj ipv4MulticastSendCommunityXml
		obj.MarshalFromObject(*s.SendCommunity)
		o.SendCommunity = &obj
	}
	o.SoftReconfigWithStoredInfo = util.YesNo(s.SoftReconfigWithStoredInfo, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastXml) UnmarshalToObject() (*Ipv4Multicast, error) {
	var addPathVal *Ipv4MulticastAddPath
	if o.AddPath != nil {
		obj, err := o.AddPath.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addPathVal = obj
	}
	var allowasInVal *Ipv4MulticastAllowasIn
	if o.AllowasIn != nil {
		obj, err := o.AllowasIn.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allowasInVal = obj
	}
	var maximumPrefixVal *Ipv4MulticastMaximumPrefix
	if o.MaximumPrefix != nil {
		obj, err := o.MaximumPrefix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		maximumPrefixVal = obj
	}
	var nextHopVal *Ipv4MulticastNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}
	var orfVal *Ipv4MulticastOrf
	if o.Orf != nil {
		obj, err := o.Orf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		orfVal = obj
	}
	var removePrivateASVal *Ipv4MulticastRemovePrivateAS
	if o.RemovePrivateAS != nil {
		obj, err := o.RemovePrivateAS.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		removePrivateASVal = obj
	}
	var sendCommunityVal *Ipv4MulticastSendCommunity
	if o.SendCommunity != nil {
		obj, err := o.SendCommunity.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sendCommunityVal = obj
	}

	result := &Ipv4Multicast{
		AddPath:                    addPathVal,
		AllowasIn:                  allowasInVal,
		AsOverride:                 util.AsBool(o.AsOverride, nil),
		DefaultOriginate:           util.AsBool(o.DefaultOriginate, nil),
		DefaultOriginateMap:        o.DefaultOriginateMap,
		Enable:                     util.AsBool(o.Enable, nil),
		MaximumPrefix:              maximumPrefixVal,
		NextHop:                    nextHopVal,
		Orf:                        orfVal,
		RemovePrivateAS:            removePrivateASVal,
		RouteReflectorClient:       util.AsBool(o.RouteReflectorClient, nil),
		SendCommunity:              sendCommunityVal,
		SoftReconfigWithStoredInfo: util.AsBool(o.SoftReconfigWithStoredInfo, nil),
		Misc:                       o.Misc,
		MiscAttributes:             o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastAddPathXml) MarshalFromObject(s Ipv4MulticastAddPath) {
	o.TxAllPaths = util.YesNo(s.TxAllPaths, nil)
	o.TxBestpathPerAS = util.YesNo(s.TxBestpathPerAS, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastAddPathXml) UnmarshalToObject() (*Ipv4MulticastAddPath, error) {

	result := &Ipv4MulticastAddPath{
		TxAllPaths:      util.AsBool(o.TxAllPaths, nil),
		TxBestpathPerAS: util.AsBool(o.TxBestpathPerAS, nil),
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastAllowasInXml) MarshalFromObject(s Ipv4MulticastAllowasIn) {
	o.Occurrence = s.Occurrence
	if s.Origin != nil {
		var obj ipv4MulticastAllowasInOriginXml
		obj.MarshalFromObject(*s.Origin)
		o.Origin = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastAllowasInXml) UnmarshalToObject() (*Ipv4MulticastAllowasIn, error) {
	var originVal *Ipv4MulticastAllowasInOrigin
	if o.Origin != nil {
		obj, err := o.Origin.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		originVal = obj
	}

	result := &Ipv4MulticastAllowasIn{
		Occurrence:     o.Occurrence,
		Origin:         originVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastAllowasInOriginXml) MarshalFromObject(s Ipv4MulticastAllowasInOrigin) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastAllowasInOriginXml) UnmarshalToObject() (*Ipv4MulticastAllowasInOrigin, error) {

	result := &Ipv4MulticastAllowasInOrigin{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastMaximumPrefixXml) MarshalFromObject(s Ipv4MulticastMaximumPrefix) {
	if s.Action != nil {
		var obj ipv4MulticastMaximumPrefixActionXml
		obj.MarshalFromObject(*s.Action)
		o.Action = &obj
	}
	o.NumPrefixes = s.NumPrefixes
	o.Threshold = s.Threshold
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastMaximumPrefixXml) UnmarshalToObject() (*Ipv4MulticastMaximumPrefix, error) {
	var actionVal *Ipv4MulticastMaximumPrefixAction
	if o.Action != nil {
		obj, err := o.Action.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		actionVal = obj
	}

	result := &Ipv4MulticastMaximumPrefix{
		Action:         actionVal,
		NumPrefixes:    o.NumPrefixes,
		Threshold:      o.Threshold,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastMaximumPrefixActionXml) MarshalFromObject(s Ipv4MulticastMaximumPrefixAction) {
	if s.Restart != nil {
		var obj ipv4MulticastMaximumPrefixActionRestartXml
		obj.MarshalFromObject(*s.Restart)
		o.Restart = &obj
	}
	if s.WarningOnly != nil {
		var obj ipv4MulticastMaximumPrefixActionWarningOnlyXml
		obj.MarshalFromObject(*s.WarningOnly)
		o.WarningOnly = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastMaximumPrefixActionXml) UnmarshalToObject() (*Ipv4MulticastMaximumPrefixAction, error) {
	var restartVal *Ipv4MulticastMaximumPrefixActionRestart
	if o.Restart != nil {
		obj, err := o.Restart.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		restartVal = obj
	}
	var warningOnlyVal *Ipv4MulticastMaximumPrefixActionWarningOnly
	if o.WarningOnly != nil {
		obj, err := o.WarningOnly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		warningOnlyVal = obj
	}

	result := &Ipv4MulticastMaximumPrefixAction{
		Restart:        restartVal,
		WarningOnly:    warningOnlyVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastMaximumPrefixActionRestartXml) MarshalFromObject(s Ipv4MulticastMaximumPrefixActionRestart) {
	o.Interval = s.Interval
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastMaximumPrefixActionRestartXml) UnmarshalToObject() (*Ipv4MulticastMaximumPrefixActionRestart, error) {

	result := &Ipv4MulticastMaximumPrefixActionRestart{
		Interval:       o.Interval,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastMaximumPrefixActionWarningOnlyXml) MarshalFromObject(s Ipv4MulticastMaximumPrefixActionWarningOnly) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastMaximumPrefixActionWarningOnlyXml) UnmarshalToObject() (*Ipv4MulticastMaximumPrefixActionWarningOnly, error) {

	result := &Ipv4MulticastMaximumPrefixActionWarningOnly{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastNextHopXml) MarshalFromObject(s Ipv4MulticastNextHop) {
	if s.Self != nil {
		var obj ipv4MulticastNextHopSelfXml
		obj.MarshalFromObject(*s.Self)
		o.Self = &obj
	}
	if s.SelfForce != nil {
		var obj ipv4MulticastNextHopSelfForceXml
		obj.MarshalFromObject(*s.SelfForce)
		o.SelfForce = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastNextHopXml) UnmarshalToObject() (*Ipv4MulticastNextHop, error) {
	var selfVal *Ipv4MulticastNextHopSelf
	if o.Self != nil {
		obj, err := o.Self.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		selfVal = obj
	}
	var selfForceVal *Ipv4MulticastNextHopSelfForce
	if o.SelfForce != nil {
		obj, err := o.SelfForce.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		selfForceVal = obj
	}

	result := &Ipv4MulticastNextHop{
		Self:           selfVal,
		SelfForce:      selfForceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastNextHopSelfXml) MarshalFromObject(s Ipv4MulticastNextHopSelf) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastNextHopSelfXml) UnmarshalToObject() (*Ipv4MulticastNextHopSelf, error) {

	result := &Ipv4MulticastNextHopSelf{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastNextHopSelfForceXml) MarshalFromObject(s Ipv4MulticastNextHopSelfForce) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastNextHopSelfForceXml) UnmarshalToObject() (*Ipv4MulticastNextHopSelfForce, error) {

	result := &Ipv4MulticastNextHopSelfForce{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastOrfXml) MarshalFromObject(s Ipv4MulticastOrf) {
	o.OrfPrefixList = s.OrfPrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastOrfXml) UnmarshalToObject() (*Ipv4MulticastOrf, error) {

	result := &Ipv4MulticastOrf{
		OrfPrefixList:  o.OrfPrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastRemovePrivateASXml) MarshalFromObject(s Ipv4MulticastRemovePrivateAS) {
	if s.All != nil {
		var obj ipv4MulticastRemovePrivateASAllXml
		obj.MarshalFromObject(*s.All)
		o.All = &obj
	}
	if s.ReplaceAS != nil {
		var obj ipv4MulticastRemovePrivateASReplaceASXml
		obj.MarshalFromObject(*s.ReplaceAS)
		o.ReplaceAS = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastRemovePrivateASXml) UnmarshalToObject() (*Ipv4MulticastRemovePrivateAS, error) {
	var allVal *Ipv4MulticastRemovePrivateASAll
	if o.All != nil {
		obj, err := o.All.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allVal = obj
	}
	var replaceASVal *Ipv4MulticastRemovePrivateASReplaceAS
	if o.ReplaceAS != nil {
		obj, err := o.ReplaceAS.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		replaceASVal = obj
	}

	result := &Ipv4MulticastRemovePrivateAS{
		All:            allVal,
		ReplaceAS:      replaceASVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastRemovePrivateASAllXml) MarshalFromObject(s Ipv4MulticastRemovePrivateASAll) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastRemovePrivateASAllXml) UnmarshalToObject() (*Ipv4MulticastRemovePrivateASAll, error) {

	result := &Ipv4MulticastRemovePrivateASAll{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastRemovePrivateASReplaceASXml) MarshalFromObject(s Ipv4MulticastRemovePrivateASReplaceAS) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastRemovePrivateASReplaceASXml) UnmarshalToObject() (*Ipv4MulticastRemovePrivateASReplaceAS, error) {

	result := &Ipv4MulticastRemovePrivateASReplaceAS{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastSendCommunityXml) MarshalFromObject(s Ipv4MulticastSendCommunity) {
	if s.All != nil {
		var obj ipv4MulticastSendCommunityAllXml
		obj.MarshalFromObject(*s.All)
		o.All = &obj
	}
	if s.Both != nil {
		var obj ipv4MulticastSendCommunityBothXml
		obj.MarshalFromObject(*s.Both)
		o.Both = &obj
	}
	if s.Extended != nil {
		var obj ipv4MulticastSendCommunityExtendedXml
		obj.MarshalFromObject(*s.Extended)
		o.Extended = &obj
	}
	if s.Large != nil {
		var obj ipv4MulticastSendCommunityLargeXml
		obj.MarshalFromObject(*s.Large)
		o.Large = &obj
	}
	if s.Standard != nil {
		var obj ipv4MulticastSendCommunityStandardXml
		obj.MarshalFromObject(*s.Standard)
		o.Standard = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastSendCommunityXml) UnmarshalToObject() (*Ipv4MulticastSendCommunity, error) {
	var allVal *Ipv4MulticastSendCommunityAll
	if o.All != nil {
		obj, err := o.All.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allVal = obj
	}
	var bothVal *Ipv4MulticastSendCommunityBoth
	if o.Both != nil {
		obj, err := o.Both.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bothVal = obj
	}
	var extendedVal *Ipv4MulticastSendCommunityExtended
	if o.Extended != nil {
		obj, err := o.Extended.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		extendedVal = obj
	}
	var largeVal *Ipv4MulticastSendCommunityLarge
	if o.Large != nil {
		obj, err := o.Large.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		largeVal = obj
	}
	var standardVal *Ipv4MulticastSendCommunityStandard
	if o.Standard != nil {
		obj, err := o.Standard.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		standardVal = obj
	}

	result := &Ipv4MulticastSendCommunity{
		All:            allVal,
		Both:           bothVal,
		Extended:       extendedVal,
		Large:          largeVal,
		Standard:       standardVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastSendCommunityAllXml) MarshalFromObject(s Ipv4MulticastSendCommunityAll) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastSendCommunityAllXml) UnmarshalToObject() (*Ipv4MulticastSendCommunityAll, error) {

	result := &Ipv4MulticastSendCommunityAll{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastSendCommunityBothXml) MarshalFromObject(s Ipv4MulticastSendCommunityBoth) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastSendCommunityBothXml) UnmarshalToObject() (*Ipv4MulticastSendCommunityBoth, error) {

	result := &Ipv4MulticastSendCommunityBoth{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastSendCommunityExtendedXml) MarshalFromObject(s Ipv4MulticastSendCommunityExtended) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastSendCommunityExtendedXml) UnmarshalToObject() (*Ipv4MulticastSendCommunityExtended, error) {

	result := &Ipv4MulticastSendCommunityExtended{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastSendCommunityLargeXml) MarshalFromObject(s Ipv4MulticastSendCommunityLarge) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastSendCommunityLargeXml) UnmarshalToObject() (*Ipv4MulticastSendCommunityLarge, error) {

	result := &Ipv4MulticastSendCommunityLarge{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4MulticastSendCommunityStandardXml) MarshalFromObject(s Ipv4MulticastSendCommunityStandard) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4MulticastSendCommunityStandardXml) UnmarshalToObject() (*Ipv4MulticastSendCommunityStandard, error) {

	result := &Ipv4MulticastSendCommunityStandard{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastXml) MarshalFromObject(s Ipv4Unicast) {
	if s.AddPath != nil {
		var obj ipv4UnicastAddPathXml
		obj.MarshalFromObject(*s.AddPath)
		o.AddPath = &obj
	}
	if s.AllowasIn != nil {
		var obj ipv4UnicastAllowasInXml
		obj.MarshalFromObject(*s.AllowasIn)
		o.AllowasIn = &obj
	}
	o.AsOverride = util.YesNo(s.AsOverride, nil)
	o.DefaultOriginate = util.YesNo(s.DefaultOriginate, nil)
	o.DefaultOriginateMap = s.DefaultOriginateMap
	o.Enable = util.YesNo(s.Enable, nil)
	if s.MaximumPrefix != nil {
		var obj ipv4UnicastMaximumPrefixXml
		obj.MarshalFromObject(*s.MaximumPrefix)
		o.MaximumPrefix = &obj
	}
	if s.NextHop != nil {
		var obj ipv4UnicastNextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	if s.Orf != nil {
		var obj ipv4UnicastOrfXml
		obj.MarshalFromObject(*s.Orf)
		o.Orf = &obj
	}
	if s.RemovePrivateAS != nil {
		var obj ipv4UnicastRemovePrivateASXml
		obj.MarshalFromObject(*s.RemovePrivateAS)
		o.RemovePrivateAS = &obj
	}
	o.RouteReflectorClient = util.YesNo(s.RouteReflectorClient, nil)
	if s.SendCommunity != nil {
		var obj ipv4UnicastSendCommunityXml
		obj.MarshalFromObject(*s.SendCommunity)
		o.SendCommunity = &obj
	}
	o.SoftReconfigWithStoredInfo = util.YesNo(s.SoftReconfigWithStoredInfo, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastXml) UnmarshalToObject() (*Ipv4Unicast, error) {
	var addPathVal *Ipv4UnicastAddPath
	if o.AddPath != nil {
		obj, err := o.AddPath.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addPathVal = obj
	}
	var allowasInVal *Ipv4UnicastAllowasIn
	if o.AllowasIn != nil {
		obj, err := o.AllowasIn.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allowasInVal = obj
	}
	var maximumPrefixVal *Ipv4UnicastMaximumPrefix
	if o.MaximumPrefix != nil {
		obj, err := o.MaximumPrefix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		maximumPrefixVal = obj
	}
	var nextHopVal *Ipv4UnicastNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}
	var orfVal *Ipv4UnicastOrf
	if o.Orf != nil {
		obj, err := o.Orf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		orfVal = obj
	}
	var removePrivateASVal *Ipv4UnicastRemovePrivateAS
	if o.RemovePrivateAS != nil {
		obj, err := o.RemovePrivateAS.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		removePrivateASVal = obj
	}
	var sendCommunityVal *Ipv4UnicastSendCommunity
	if o.SendCommunity != nil {
		obj, err := o.SendCommunity.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sendCommunityVal = obj
	}

	result := &Ipv4Unicast{
		AddPath:                    addPathVal,
		AllowasIn:                  allowasInVal,
		AsOverride:                 util.AsBool(o.AsOverride, nil),
		DefaultOriginate:           util.AsBool(o.DefaultOriginate, nil),
		DefaultOriginateMap:        o.DefaultOriginateMap,
		Enable:                     util.AsBool(o.Enable, nil),
		MaximumPrefix:              maximumPrefixVal,
		NextHop:                    nextHopVal,
		Orf:                        orfVal,
		RemovePrivateAS:            removePrivateASVal,
		RouteReflectorClient:       util.AsBool(o.RouteReflectorClient, nil),
		SendCommunity:              sendCommunityVal,
		SoftReconfigWithStoredInfo: util.AsBool(o.SoftReconfigWithStoredInfo, nil),
		Misc:                       o.Misc,
		MiscAttributes:             o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastAddPathXml) MarshalFromObject(s Ipv4UnicastAddPath) {
	o.TxAllPaths = util.YesNo(s.TxAllPaths, nil)
	o.TxBestpathPerAS = util.YesNo(s.TxBestpathPerAS, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastAddPathXml) UnmarshalToObject() (*Ipv4UnicastAddPath, error) {

	result := &Ipv4UnicastAddPath{
		TxAllPaths:      util.AsBool(o.TxAllPaths, nil),
		TxBestpathPerAS: util.AsBool(o.TxBestpathPerAS, nil),
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastAllowasInXml) MarshalFromObject(s Ipv4UnicastAllowasIn) {
	o.Occurrence = s.Occurrence
	if s.Origin != nil {
		var obj ipv4UnicastAllowasInOriginXml
		obj.MarshalFromObject(*s.Origin)
		o.Origin = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastAllowasInXml) UnmarshalToObject() (*Ipv4UnicastAllowasIn, error) {
	var originVal *Ipv4UnicastAllowasInOrigin
	if o.Origin != nil {
		obj, err := o.Origin.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		originVal = obj
	}

	result := &Ipv4UnicastAllowasIn{
		Occurrence:     o.Occurrence,
		Origin:         originVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastAllowasInOriginXml) MarshalFromObject(s Ipv4UnicastAllowasInOrigin) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastAllowasInOriginXml) UnmarshalToObject() (*Ipv4UnicastAllowasInOrigin, error) {

	result := &Ipv4UnicastAllowasInOrigin{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastMaximumPrefixXml) MarshalFromObject(s Ipv4UnicastMaximumPrefix) {
	if s.Action != nil {
		var obj ipv4UnicastMaximumPrefixActionXml
		obj.MarshalFromObject(*s.Action)
		o.Action = &obj
	}
	o.NumPrefixes = s.NumPrefixes
	o.Threshold = s.Threshold
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastMaximumPrefixXml) UnmarshalToObject() (*Ipv4UnicastMaximumPrefix, error) {
	var actionVal *Ipv4UnicastMaximumPrefixAction
	if o.Action != nil {
		obj, err := o.Action.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		actionVal = obj
	}

	result := &Ipv4UnicastMaximumPrefix{
		Action:         actionVal,
		NumPrefixes:    o.NumPrefixes,
		Threshold:      o.Threshold,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastMaximumPrefixActionXml) MarshalFromObject(s Ipv4UnicastMaximumPrefixAction) {
	if s.Restart != nil {
		var obj ipv4UnicastMaximumPrefixActionRestartXml
		obj.MarshalFromObject(*s.Restart)
		o.Restart = &obj
	}
	if s.WarningOnly != nil {
		var obj ipv4UnicastMaximumPrefixActionWarningOnlyXml
		obj.MarshalFromObject(*s.WarningOnly)
		o.WarningOnly = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastMaximumPrefixActionXml) UnmarshalToObject() (*Ipv4UnicastMaximumPrefixAction, error) {
	var restartVal *Ipv4UnicastMaximumPrefixActionRestart
	if o.Restart != nil {
		obj, err := o.Restart.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		restartVal = obj
	}
	var warningOnlyVal *Ipv4UnicastMaximumPrefixActionWarningOnly
	if o.WarningOnly != nil {
		obj, err := o.WarningOnly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		warningOnlyVal = obj
	}

	result := &Ipv4UnicastMaximumPrefixAction{
		Restart:        restartVal,
		WarningOnly:    warningOnlyVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastMaximumPrefixActionRestartXml) MarshalFromObject(s Ipv4UnicastMaximumPrefixActionRestart) {
	o.Interval = s.Interval
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastMaximumPrefixActionRestartXml) UnmarshalToObject() (*Ipv4UnicastMaximumPrefixActionRestart, error) {

	result := &Ipv4UnicastMaximumPrefixActionRestart{
		Interval:       o.Interval,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastMaximumPrefixActionWarningOnlyXml) MarshalFromObject(s Ipv4UnicastMaximumPrefixActionWarningOnly) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastMaximumPrefixActionWarningOnlyXml) UnmarshalToObject() (*Ipv4UnicastMaximumPrefixActionWarningOnly, error) {

	result := &Ipv4UnicastMaximumPrefixActionWarningOnly{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastNextHopXml) MarshalFromObject(s Ipv4UnicastNextHop) {
	if s.Self != nil {
		var obj ipv4UnicastNextHopSelfXml
		obj.MarshalFromObject(*s.Self)
		o.Self = &obj
	}
	if s.SelfForce != nil {
		var obj ipv4UnicastNextHopSelfForceXml
		obj.MarshalFromObject(*s.SelfForce)
		o.SelfForce = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastNextHopXml) UnmarshalToObject() (*Ipv4UnicastNextHop, error) {
	var selfVal *Ipv4UnicastNextHopSelf
	if o.Self != nil {
		obj, err := o.Self.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		selfVal = obj
	}
	var selfForceVal *Ipv4UnicastNextHopSelfForce
	if o.SelfForce != nil {
		obj, err := o.SelfForce.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		selfForceVal = obj
	}

	result := &Ipv4UnicastNextHop{
		Self:           selfVal,
		SelfForce:      selfForceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastNextHopSelfXml) MarshalFromObject(s Ipv4UnicastNextHopSelf) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastNextHopSelfXml) UnmarshalToObject() (*Ipv4UnicastNextHopSelf, error) {

	result := &Ipv4UnicastNextHopSelf{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastNextHopSelfForceXml) MarshalFromObject(s Ipv4UnicastNextHopSelfForce) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastNextHopSelfForceXml) UnmarshalToObject() (*Ipv4UnicastNextHopSelfForce, error) {

	result := &Ipv4UnicastNextHopSelfForce{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastOrfXml) MarshalFromObject(s Ipv4UnicastOrf) {
	o.OrfPrefixList = s.OrfPrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastOrfXml) UnmarshalToObject() (*Ipv4UnicastOrf, error) {

	result := &Ipv4UnicastOrf{
		OrfPrefixList:  o.OrfPrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastRemovePrivateASXml) MarshalFromObject(s Ipv4UnicastRemovePrivateAS) {
	if s.All != nil {
		var obj ipv4UnicastRemovePrivateASAllXml
		obj.MarshalFromObject(*s.All)
		o.All = &obj
	}
	if s.ReplaceAS != nil {
		var obj ipv4UnicastRemovePrivateASReplaceASXml
		obj.MarshalFromObject(*s.ReplaceAS)
		o.ReplaceAS = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastRemovePrivateASXml) UnmarshalToObject() (*Ipv4UnicastRemovePrivateAS, error) {
	var allVal *Ipv4UnicastRemovePrivateASAll
	if o.All != nil {
		obj, err := o.All.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allVal = obj
	}
	var replaceASVal *Ipv4UnicastRemovePrivateASReplaceAS
	if o.ReplaceAS != nil {
		obj, err := o.ReplaceAS.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		replaceASVal = obj
	}

	result := &Ipv4UnicastRemovePrivateAS{
		All:            allVal,
		ReplaceAS:      replaceASVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastRemovePrivateASAllXml) MarshalFromObject(s Ipv4UnicastRemovePrivateASAll) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastRemovePrivateASAllXml) UnmarshalToObject() (*Ipv4UnicastRemovePrivateASAll, error) {

	result := &Ipv4UnicastRemovePrivateASAll{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastRemovePrivateASReplaceASXml) MarshalFromObject(s Ipv4UnicastRemovePrivateASReplaceAS) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastRemovePrivateASReplaceASXml) UnmarshalToObject() (*Ipv4UnicastRemovePrivateASReplaceAS, error) {

	result := &Ipv4UnicastRemovePrivateASReplaceAS{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastSendCommunityXml) MarshalFromObject(s Ipv4UnicastSendCommunity) {
	if s.All != nil {
		var obj ipv4UnicastSendCommunityAllXml
		obj.MarshalFromObject(*s.All)
		o.All = &obj
	}
	if s.Both != nil {
		var obj ipv4UnicastSendCommunityBothXml
		obj.MarshalFromObject(*s.Both)
		o.Both = &obj
	}
	if s.Extended != nil {
		var obj ipv4UnicastSendCommunityExtendedXml
		obj.MarshalFromObject(*s.Extended)
		o.Extended = &obj
	}
	if s.Large != nil {
		var obj ipv4UnicastSendCommunityLargeXml
		obj.MarshalFromObject(*s.Large)
		o.Large = &obj
	}
	if s.Standard != nil {
		var obj ipv4UnicastSendCommunityStandardXml
		obj.MarshalFromObject(*s.Standard)
		o.Standard = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastSendCommunityXml) UnmarshalToObject() (*Ipv4UnicastSendCommunity, error) {
	var allVal *Ipv4UnicastSendCommunityAll
	if o.All != nil {
		obj, err := o.All.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allVal = obj
	}
	var bothVal *Ipv4UnicastSendCommunityBoth
	if o.Both != nil {
		obj, err := o.Both.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bothVal = obj
	}
	var extendedVal *Ipv4UnicastSendCommunityExtended
	if o.Extended != nil {
		obj, err := o.Extended.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		extendedVal = obj
	}
	var largeVal *Ipv4UnicastSendCommunityLarge
	if o.Large != nil {
		obj, err := o.Large.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		largeVal = obj
	}
	var standardVal *Ipv4UnicastSendCommunityStandard
	if o.Standard != nil {
		obj, err := o.Standard.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		standardVal = obj
	}

	result := &Ipv4UnicastSendCommunity{
		All:            allVal,
		Both:           bothVal,
		Extended:       extendedVal,
		Large:          largeVal,
		Standard:       standardVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastSendCommunityAllXml) MarshalFromObject(s Ipv4UnicastSendCommunityAll) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastSendCommunityAllXml) UnmarshalToObject() (*Ipv4UnicastSendCommunityAll, error) {

	result := &Ipv4UnicastSendCommunityAll{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastSendCommunityBothXml) MarshalFromObject(s Ipv4UnicastSendCommunityBoth) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastSendCommunityBothXml) UnmarshalToObject() (*Ipv4UnicastSendCommunityBoth, error) {

	result := &Ipv4UnicastSendCommunityBoth{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastSendCommunityExtendedXml) MarshalFromObject(s Ipv4UnicastSendCommunityExtended) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastSendCommunityExtendedXml) UnmarshalToObject() (*Ipv4UnicastSendCommunityExtended, error) {

	result := &Ipv4UnicastSendCommunityExtended{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastSendCommunityLargeXml) MarshalFromObject(s Ipv4UnicastSendCommunityLarge) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastSendCommunityLargeXml) UnmarshalToObject() (*Ipv4UnicastSendCommunityLarge, error) {

	result := &Ipv4UnicastSendCommunityLarge{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv4UnicastSendCommunityStandardXml) MarshalFromObject(s Ipv4UnicastSendCommunityStandard) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv4UnicastSendCommunityStandardXml) UnmarshalToObject() (*Ipv4UnicastSendCommunityStandard, error) {

	result := &Ipv4UnicastSendCommunityStandard{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6Xml) MarshalFromObject(s Ipv6) {
	if s.Unicast != nil {
		var obj ipv6UnicastXml
		obj.MarshalFromObject(*s.Unicast)
		o.Unicast = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6Xml) UnmarshalToObject() (*Ipv6, error) {
	var unicastVal *Ipv6Unicast
	if o.Unicast != nil {
		obj, err := o.Unicast.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		unicastVal = obj
	}

	result := &Ipv6{
		Unicast:        unicastVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastXml) MarshalFromObject(s Ipv6Unicast) {
	if s.AddPath != nil {
		var obj ipv6UnicastAddPathXml
		obj.MarshalFromObject(*s.AddPath)
		o.AddPath = &obj
	}
	if s.AllowasIn != nil {
		var obj ipv6UnicastAllowasInXml
		obj.MarshalFromObject(*s.AllowasIn)
		o.AllowasIn = &obj
	}
	o.AsOverride = util.YesNo(s.AsOverride, nil)
	o.DefaultOriginate = util.YesNo(s.DefaultOriginate, nil)
	o.DefaultOriginateMap = s.DefaultOriginateMap
	o.Enable = util.YesNo(s.Enable, nil)
	if s.MaximumPrefix != nil {
		var obj ipv6UnicastMaximumPrefixXml
		obj.MarshalFromObject(*s.MaximumPrefix)
		o.MaximumPrefix = &obj
	}
	if s.NextHop != nil {
		var obj ipv6UnicastNextHopXml
		obj.MarshalFromObject(*s.NextHop)
		o.NextHop = &obj
	}
	if s.Orf != nil {
		var obj ipv6UnicastOrfXml
		obj.MarshalFromObject(*s.Orf)
		o.Orf = &obj
	}
	if s.RemovePrivateAS != nil {
		var obj ipv6UnicastRemovePrivateASXml
		obj.MarshalFromObject(*s.RemovePrivateAS)
		o.RemovePrivateAS = &obj
	}
	o.RouteReflectorClient = util.YesNo(s.RouteReflectorClient, nil)
	if s.SendCommunity != nil {
		var obj ipv6UnicastSendCommunityXml
		obj.MarshalFromObject(*s.SendCommunity)
		o.SendCommunity = &obj
	}
	o.SoftReconfigWithStoredInfo = util.YesNo(s.SoftReconfigWithStoredInfo, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastXml) UnmarshalToObject() (*Ipv6Unicast, error) {
	var addPathVal *Ipv6UnicastAddPath
	if o.AddPath != nil {
		obj, err := o.AddPath.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		addPathVal = obj
	}
	var allowasInVal *Ipv6UnicastAllowasIn
	if o.AllowasIn != nil {
		obj, err := o.AllowasIn.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allowasInVal = obj
	}
	var maximumPrefixVal *Ipv6UnicastMaximumPrefix
	if o.MaximumPrefix != nil {
		obj, err := o.MaximumPrefix.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		maximumPrefixVal = obj
	}
	var nextHopVal *Ipv6UnicastNextHop
	if o.NextHop != nil {
		obj, err := o.NextHop.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		nextHopVal = obj
	}
	var orfVal *Ipv6UnicastOrf
	if o.Orf != nil {
		obj, err := o.Orf.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		orfVal = obj
	}
	var removePrivateASVal *Ipv6UnicastRemovePrivateAS
	if o.RemovePrivateAS != nil {
		obj, err := o.RemovePrivateAS.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		removePrivateASVal = obj
	}
	var sendCommunityVal *Ipv6UnicastSendCommunity
	if o.SendCommunity != nil {
		obj, err := o.SendCommunity.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sendCommunityVal = obj
	}

	result := &Ipv6Unicast{
		AddPath:                    addPathVal,
		AllowasIn:                  allowasInVal,
		AsOverride:                 util.AsBool(o.AsOverride, nil),
		DefaultOriginate:           util.AsBool(o.DefaultOriginate, nil),
		DefaultOriginateMap:        o.DefaultOriginateMap,
		Enable:                     util.AsBool(o.Enable, nil),
		MaximumPrefix:              maximumPrefixVal,
		NextHop:                    nextHopVal,
		Orf:                        orfVal,
		RemovePrivateAS:            removePrivateASVal,
		RouteReflectorClient:       util.AsBool(o.RouteReflectorClient, nil),
		SendCommunity:              sendCommunityVal,
		SoftReconfigWithStoredInfo: util.AsBool(o.SoftReconfigWithStoredInfo, nil),
		Misc:                       o.Misc,
		MiscAttributes:             o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastAddPathXml) MarshalFromObject(s Ipv6UnicastAddPath) {
	o.TxAllPaths = util.YesNo(s.TxAllPaths, nil)
	o.TxBestpathPerAS = util.YesNo(s.TxBestpathPerAS, nil)
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastAddPathXml) UnmarshalToObject() (*Ipv6UnicastAddPath, error) {

	result := &Ipv6UnicastAddPath{
		TxAllPaths:      util.AsBool(o.TxAllPaths, nil),
		TxBestpathPerAS: util.AsBool(o.TxBestpathPerAS, nil),
		Misc:            o.Misc,
		MiscAttributes:  o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastAllowasInXml) MarshalFromObject(s Ipv6UnicastAllowasIn) {
	o.Occurrence = s.Occurrence
	if s.Origin != nil {
		var obj ipv6UnicastAllowasInOriginXml
		obj.MarshalFromObject(*s.Origin)
		o.Origin = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastAllowasInXml) UnmarshalToObject() (*Ipv6UnicastAllowasIn, error) {
	var originVal *Ipv6UnicastAllowasInOrigin
	if o.Origin != nil {
		obj, err := o.Origin.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		originVal = obj
	}

	result := &Ipv6UnicastAllowasIn{
		Occurrence:     o.Occurrence,
		Origin:         originVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastAllowasInOriginXml) MarshalFromObject(s Ipv6UnicastAllowasInOrigin) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastAllowasInOriginXml) UnmarshalToObject() (*Ipv6UnicastAllowasInOrigin, error) {

	result := &Ipv6UnicastAllowasInOrigin{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastMaximumPrefixXml) MarshalFromObject(s Ipv6UnicastMaximumPrefix) {
	if s.Action != nil {
		var obj ipv6UnicastMaximumPrefixActionXml
		obj.MarshalFromObject(*s.Action)
		o.Action = &obj
	}
	o.NumPrefixes = s.NumPrefixes
	o.Threshold = s.Threshold
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastMaximumPrefixXml) UnmarshalToObject() (*Ipv6UnicastMaximumPrefix, error) {
	var actionVal *Ipv6UnicastMaximumPrefixAction
	if o.Action != nil {
		obj, err := o.Action.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		actionVal = obj
	}

	result := &Ipv6UnicastMaximumPrefix{
		Action:         actionVal,
		NumPrefixes:    o.NumPrefixes,
		Threshold:      o.Threshold,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastMaximumPrefixActionXml) MarshalFromObject(s Ipv6UnicastMaximumPrefixAction) {
	if s.Restart != nil {
		var obj ipv6UnicastMaximumPrefixActionRestartXml
		obj.MarshalFromObject(*s.Restart)
		o.Restart = &obj
	}
	if s.WarningOnly != nil {
		var obj ipv6UnicastMaximumPrefixActionWarningOnlyXml
		obj.MarshalFromObject(*s.WarningOnly)
		o.WarningOnly = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastMaximumPrefixActionXml) UnmarshalToObject() (*Ipv6UnicastMaximumPrefixAction, error) {
	var restartVal *Ipv6UnicastMaximumPrefixActionRestart
	if o.Restart != nil {
		obj, err := o.Restart.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		restartVal = obj
	}
	var warningOnlyVal *Ipv6UnicastMaximumPrefixActionWarningOnly
	if o.WarningOnly != nil {
		obj, err := o.WarningOnly.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		warningOnlyVal = obj
	}

	result := &Ipv6UnicastMaximumPrefixAction{
		Restart:        restartVal,
		WarningOnly:    warningOnlyVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastMaximumPrefixActionRestartXml) MarshalFromObject(s Ipv6UnicastMaximumPrefixActionRestart) {
	o.Interval = s.Interval
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastMaximumPrefixActionRestartXml) UnmarshalToObject() (*Ipv6UnicastMaximumPrefixActionRestart, error) {

	result := &Ipv6UnicastMaximumPrefixActionRestart{
		Interval:       o.Interval,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastMaximumPrefixActionWarningOnlyXml) MarshalFromObject(s Ipv6UnicastMaximumPrefixActionWarningOnly) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastMaximumPrefixActionWarningOnlyXml) UnmarshalToObject() (*Ipv6UnicastMaximumPrefixActionWarningOnly, error) {

	result := &Ipv6UnicastMaximumPrefixActionWarningOnly{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastNextHopXml) MarshalFromObject(s Ipv6UnicastNextHop) {
	if s.Self != nil {
		var obj ipv6UnicastNextHopSelfXml
		obj.MarshalFromObject(*s.Self)
		o.Self = &obj
	}
	if s.SelfForce != nil {
		var obj ipv6UnicastNextHopSelfForceXml
		obj.MarshalFromObject(*s.SelfForce)
		o.SelfForce = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastNextHopXml) UnmarshalToObject() (*Ipv6UnicastNextHop, error) {
	var selfVal *Ipv6UnicastNextHopSelf
	if o.Self != nil {
		obj, err := o.Self.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		selfVal = obj
	}
	var selfForceVal *Ipv6UnicastNextHopSelfForce
	if o.SelfForce != nil {
		obj, err := o.SelfForce.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		selfForceVal = obj
	}

	result := &Ipv6UnicastNextHop{
		Self:           selfVal,
		SelfForce:      selfForceVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastNextHopSelfXml) MarshalFromObject(s Ipv6UnicastNextHopSelf) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastNextHopSelfXml) UnmarshalToObject() (*Ipv6UnicastNextHopSelf, error) {

	result := &Ipv6UnicastNextHopSelf{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastNextHopSelfForceXml) MarshalFromObject(s Ipv6UnicastNextHopSelfForce) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastNextHopSelfForceXml) UnmarshalToObject() (*Ipv6UnicastNextHopSelfForce, error) {

	result := &Ipv6UnicastNextHopSelfForce{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastOrfXml) MarshalFromObject(s Ipv6UnicastOrf) {
	o.OrfPrefixList = s.OrfPrefixList
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastOrfXml) UnmarshalToObject() (*Ipv6UnicastOrf, error) {

	result := &Ipv6UnicastOrf{
		OrfPrefixList:  o.OrfPrefixList,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastRemovePrivateASXml) MarshalFromObject(s Ipv6UnicastRemovePrivateAS) {
	if s.All != nil {
		var obj ipv6UnicastRemovePrivateASAllXml
		obj.MarshalFromObject(*s.All)
		o.All = &obj
	}
	if s.ReplaceAS != nil {
		var obj ipv6UnicastRemovePrivateASReplaceASXml
		obj.MarshalFromObject(*s.ReplaceAS)
		o.ReplaceAS = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastRemovePrivateASXml) UnmarshalToObject() (*Ipv6UnicastRemovePrivateAS, error) {
	var allVal *Ipv6UnicastRemovePrivateASAll
	if o.All != nil {
		obj, err := o.All.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allVal = obj
	}
	var replaceASVal *Ipv6UnicastRemovePrivateASReplaceAS
	if o.ReplaceAS != nil {
		obj, err := o.ReplaceAS.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		replaceASVal = obj
	}

	result := &Ipv6UnicastRemovePrivateAS{
		All:            allVal,
		ReplaceAS:      replaceASVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastRemovePrivateASAllXml) MarshalFromObject(s Ipv6UnicastRemovePrivateASAll) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastRemovePrivateASAllXml) UnmarshalToObject() (*Ipv6UnicastRemovePrivateASAll, error) {

	result := &Ipv6UnicastRemovePrivateASAll{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastRemovePrivateASReplaceASXml) MarshalFromObject(s Ipv6UnicastRemovePrivateASReplaceAS) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastRemovePrivateASReplaceASXml) UnmarshalToObject() (*Ipv6UnicastRemovePrivateASReplaceAS, error) {

	result := &Ipv6UnicastRemovePrivateASReplaceAS{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastSendCommunityXml) MarshalFromObject(s Ipv6UnicastSendCommunity) {
	if s.All != nil {
		var obj ipv6UnicastSendCommunityAllXml
		obj.MarshalFromObject(*s.All)
		o.All = &obj
	}
	if s.Both != nil {
		var obj ipv6UnicastSendCommunityBothXml
		obj.MarshalFromObject(*s.Both)
		o.Both = &obj
	}
	if s.Extended != nil {
		var obj ipv6UnicastSendCommunityExtendedXml
		obj.MarshalFromObject(*s.Extended)
		o.Extended = &obj
	}
	if s.Large != nil {
		var obj ipv6UnicastSendCommunityLargeXml
		obj.MarshalFromObject(*s.Large)
		o.Large = &obj
	}
	if s.Standard != nil {
		var obj ipv6UnicastSendCommunityStandardXml
		obj.MarshalFromObject(*s.Standard)
		o.Standard = &obj
	}
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastSendCommunityXml) UnmarshalToObject() (*Ipv6UnicastSendCommunity, error) {
	var allVal *Ipv6UnicastSendCommunityAll
	if o.All != nil {
		obj, err := o.All.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		allVal = obj
	}
	var bothVal *Ipv6UnicastSendCommunityBoth
	if o.Both != nil {
		obj, err := o.Both.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		bothVal = obj
	}
	var extendedVal *Ipv6UnicastSendCommunityExtended
	if o.Extended != nil {
		obj, err := o.Extended.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		extendedVal = obj
	}
	var largeVal *Ipv6UnicastSendCommunityLarge
	if o.Large != nil {
		obj, err := o.Large.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		largeVal = obj
	}
	var standardVal *Ipv6UnicastSendCommunityStandard
	if o.Standard != nil {
		obj, err := o.Standard.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		standardVal = obj
	}

	result := &Ipv6UnicastSendCommunity{
		All:            allVal,
		Both:           bothVal,
		Extended:       extendedVal,
		Large:          largeVal,
		Standard:       standardVal,
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastSendCommunityAllXml) MarshalFromObject(s Ipv6UnicastSendCommunityAll) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastSendCommunityAllXml) UnmarshalToObject() (*Ipv6UnicastSendCommunityAll, error) {

	result := &Ipv6UnicastSendCommunityAll{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastSendCommunityBothXml) MarshalFromObject(s Ipv6UnicastSendCommunityBoth) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastSendCommunityBothXml) UnmarshalToObject() (*Ipv6UnicastSendCommunityBoth, error) {

	result := &Ipv6UnicastSendCommunityBoth{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastSendCommunityExtendedXml) MarshalFromObject(s Ipv6UnicastSendCommunityExtended) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastSendCommunityExtendedXml) UnmarshalToObject() (*Ipv6UnicastSendCommunityExtended, error) {

	result := &Ipv6UnicastSendCommunityExtended{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastSendCommunityLargeXml) MarshalFromObject(s Ipv6UnicastSendCommunityLarge) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastSendCommunityLargeXml) UnmarshalToObject() (*Ipv6UnicastSendCommunityLarge, error) {

	result := &Ipv6UnicastSendCommunityLarge{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}
func (o *ipv6UnicastSendCommunityStandardXml) MarshalFromObject(s Ipv6UnicastSendCommunityStandard) {
	o.Misc = s.Misc
	o.MiscAttributes = s.MiscAttributes
}

func (o ipv6UnicastSendCommunityStandardXml) UnmarshalToObject() (*Ipv6UnicastSendCommunityStandard, error) {

	result := &Ipv6UnicastSendCommunityStandard{
		Misc:           o.Misc,
		MiscAttributes: o.MiscAttributes,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "ipv4" || v == "Ipv4" {
		return e.Ipv4, nil
	}
	if v == "ipv6" || v == "Ipv6" {
		return e.Ipv6, nil
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
	if !o.Ipv4.matches(other.Ipv4) {
		return false
	}
	if !o.Ipv6.matches(other.Ipv6) {
		return false
	}

	return true
}

func (o *Ipv4) matches(other *Ipv4) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Multicast.matches(other.Multicast) {
		return false
	}
	if !o.Unicast.matches(other.Unicast) {
		return false
	}

	return true
}

func (o *Ipv4Multicast) matches(other *Ipv4Multicast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.AddPath.matches(other.AddPath) {
		return false
	}
	if !o.AllowasIn.matches(other.AllowasIn) {
		return false
	}
	if !util.BoolsMatch(o.AsOverride, other.AsOverride) {
		return false
	}
	if !util.BoolsMatch(o.DefaultOriginate, other.DefaultOriginate) {
		return false
	}
	if !util.StringsMatch(o.DefaultOriginateMap, other.DefaultOriginateMap) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.MaximumPrefix.matches(other.MaximumPrefix) {
		return false
	}
	if !o.NextHop.matches(other.NextHop) {
		return false
	}
	if !o.Orf.matches(other.Orf) {
		return false
	}
	if !o.RemovePrivateAS.matches(other.RemovePrivateAS) {
		return false
	}
	if !util.BoolsMatch(o.RouteReflectorClient, other.RouteReflectorClient) {
		return false
	}
	if !o.SendCommunity.matches(other.SendCommunity) {
		return false
	}
	if !util.BoolsMatch(o.SoftReconfigWithStoredInfo, other.SoftReconfigWithStoredInfo) {
		return false
	}

	return true
}

func (o *Ipv4MulticastAddPath) matches(other *Ipv4MulticastAddPath) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.TxAllPaths, other.TxAllPaths) {
		return false
	}
	if !util.BoolsMatch(o.TxBestpathPerAS, other.TxBestpathPerAS) {
		return false
	}

	return true
}

func (o *Ipv4MulticastAllowasIn) matches(other *Ipv4MulticastAllowasIn) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Occurrence, other.Occurrence) {
		return false
	}
	if !o.Origin.matches(other.Origin) {
		return false
	}

	return true
}

func (o *Ipv4MulticastAllowasInOrigin) matches(other *Ipv4MulticastAllowasInOrigin) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4MulticastMaximumPrefix) matches(other *Ipv4MulticastMaximumPrefix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Action.matches(other.Action) {
		return false
	}
	if !util.Ints64Match(o.NumPrefixes, other.NumPrefixes) {
		return false
	}
	if !util.Ints64Match(o.Threshold, other.Threshold) {
		return false
	}

	return true
}

func (o *Ipv4MulticastMaximumPrefixAction) matches(other *Ipv4MulticastMaximumPrefixAction) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Restart.matches(other.Restart) {
		return false
	}
	if !o.WarningOnly.matches(other.WarningOnly) {
		return false
	}

	return true
}

func (o *Ipv4MulticastMaximumPrefixActionRestart) matches(other *Ipv4MulticastMaximumPrefixActionRestart) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Interval, other.Interval) {
		return false
	}

	return true
}

func (o *Ipv4MulticastMaximumPrefixActionWarningOnly) matches(other *Ipv4MulticastMaximumPrefixActionWarningOnly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4MulticastNextHop) matches(other *Ipv4MulticastNextHop) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Self.matches(other.Self) {
		return false
	}
	if !o.SelfForce.matches(other.SelfForce) {
		return false
	}

	return true
}

func (o *Ipv4MulticastNextHopSelf) matches(other *Ipv4MulticastNextHopSelf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4MulticastNextHopSelfForce) matches(other *Ipv4MulticastNextHopSelfForce) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4MulticastOrf) matches(other *Ipv4MulticastOrf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.OrfPrefixList, other.OrfPrefixList) {
		return false
	}

	return true
}

func (o *Ipv4MulticastRemovePrivateAS) matches(other *Ipv4MulticastRemovePrivateAS) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.All.matches(other.All) {
		return false
	}
	if !o.ReplaceAS.matches(other.ReplaceAS) {
		return false
	}

	return true
}

func (o *Ipv4MulticastRemovePrivateASAll) matches(other *Ipv4MulticastRemovePrivateASAll) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4MulticastRemovePrivateASReplaceAS) matches(other *Ipv4MulticastRemovePrivateASReplaceAS) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4MulticastSendCommunity) matches(other *Ipv4MulticastSendCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.All.matches(other.All) {
		return false
	}
	if !o.Both.matches(other.Both) {
		return false
	}
	if !o.Extended.matches(other.Extended) {
		return false
	}
	if !o.Large.matches(other.Large) {
		return false
	}
	if !o.Standard.matches(other.Standard) {
		return false
	}

	return true
}

func (o *Ipv4MulticastSendCommunityAll) matches(other *Ipv4MulticastSendCommunityAll) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4MulticastSendCommunityBoth) matches(other *Ipv4MulticastSendCommunityBoth) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4MulticastSendCommunityExtended) matches(other *Ipv4MulticastSendCommunityExtended) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4MulticastSendCommunityLarge) matches(other *Ipv4MulticastSendCommunityLarge) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4MulticastSendCommunityStandard) matches(other *Ipv4MulticastSendCommunityStandard) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4Unicast) matches(other *Ipv4Unicast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.AddPath.matches(other.AddPath) {
		return false
	}
	if !o.AllowasIn.matches(other.AllowasIn) {
		return false
	}
	if !util.BoolsMatch(o.AsOverride, other.AsOverride) {
		return false
	}
	if !util.BoolsMatch(o.DefaultOriginate, other.DefaultOriginate) {
		return false
	}
	if !util.StringsMatch(o.DefaultOriginateMap, other.DefaultOriginateMap) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.MaximumPrefix.matches(other.MaximumPrefix) {
		return false
	}
	if !o.NextHop.matches(other.NextHop) {
		return false
	}
	if !o.Orf.matches(other.Orf) {
		return false
	}
	if !o.RemovePrivateAS.matches(other.RemovePrivateAS) {
		return false
	}
	if !util.BoolsMatch(o.RouteReflectorClient, other.RouteReflectorClient) {
		return false
	}
	if !o.SendCommunity.matches(other.SendCommunity) {
		return false
	}
	if !util.BoolsMatch(o.SoftReconfigWithStoredInfo, other.SoftReconfigWithStoredInfo) {
		return false
	}

	return true
}

func (o *Ipv4UnicastAddPath) matches(other *Ipv4UnicastAddPath) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.TxAllPaths, other.TxAllPaths) {
		return false
	}
	if !util.BoolsMatch(o.TxBestpathPerAS, other.TxBestpathPerAS) {
		return false
	}

	return true
}

func (o *Ipv4UnicastAllowasIn) matches(other *Ipv4UnicastAllowasIn) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Occurrence, other.Occurrence) {
		return false
	}
	if !o.Origin.matches(other.Origin) {
		return false
	}

	return true
}

func (o *Ipv4UnicastAllowasInOrigin) matches(other *Ipv4UnicastAllowasInOrigin) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4UnicastMaximumPrefix) matches(other *Ipv4UnicastMaximumPrefix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Action.matches(other.Action) {
		return false
	}
	if !util.Ints64Match(o.NumPrefixes, other.NumPrefixes) {
		return false
	}
	if !util.Ints64Match(o.Threshold, other.Threshold) {
		return false
	}

	return true
}

func (o *Ipv4UnicastMaximumPrefixAction) matches(other *Ipv4UnicastMaximumPrefixAction) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Restart.matches(other.Restart) {
		return false
	}
	if !o.WarningOnly.matches(other.WarningOnly) {
		return false
	}

	return true
}

func (o *Ipv4UnicastMaximumPrefixActionRestart) matches(other *Ipv4UnicastMaximumPrefixActionRestart) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Interval, other.Interval) {
		return false
	}

	return true
}

func (o *Ipv4UnicastMaximumPrefixActionWarningOnly) matches(other *Ipv4UnicastMaximumPrefixActionWarningOnly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4UnicastNextHop) matches(other *Ipv4UnicastNextHop) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Self.matches(other.Self) {
		return false
	}
	if !o.SelfForce.matches(other.SelfForce) {
		return false
	}

	return true
}

func (o *Ipv4UnicastNextHopSelf) matches(other *Ipv4UnicastNextHopSelf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4UnicastNextHopSelfForce) matches(other *Ipv4UnicastNextHopSelfForce) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4UnicastOrf) matches(other *Ipv4UnicastOrf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.OrfPrefixList, other.OrfPrefixList) {
		return false
	}

	return true
}

func (o *Ipv4UnicastRemovePrivateAS) matches(other *Ipv4UnicastRemovePrivateAS) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.All.matches(other.All) {
		return false
	}
	if !o.ReplaceAS.matches(other.ReplaceAS) {
		return false
	}

	return true
}

func (o *Ipv4UnicastRemovePrivateASAll) matches(other *Ipv4UnicastRemovePrivateASAll) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4UnicastRemovePrivateASReplaceAS) matches(other *Ipv4UnicastRemovePrivateASReplaceAS) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4UnicastSendCommunity) matches(other *Ipv4UnicastSendCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.All.matches(other.All) {
		return false
	}
	if !o.Both.matches(other.Both) {
		return false
	}
	if !o.Extended.matches(other.Extended) {
		return false
	}
	if !o.Large.matches(other.Large) {
		return false
	}
	if !o.Standard.matches(other.Standard) {
		return false
	}

	return true
}

func (o *Ipv4UnicastSendCommunityAll) matches(other *Ipv4UnicastSendCommunityAll) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4UnicastSendCommunityBoth) matches(other *Ipv4UnicastSendCommunityBoth) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4UnicastSendCommunityExtended) matches(other *Ipv4UnicastSendCommunityExtended) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4UnicastSendCommunityLarge) matches(other *Ipv4UnicastSendCommunityLarge) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv4UnicastSendCommunityStandard) matches(other *Ipv4UnicastSendCommunityStandard) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6) matches(other *Ipv6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Unicast.matches(other.Unicast) {
		return false
	}

	return true
}

func (o *Ipv6Unicast) matches(other *Ipv6Unicast) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.AddPath.matches(other.AddPath) {
		return false
	}
	if !o.AllowasIn.matches(other.AllowasIn) {
		return false
	}
	if !util.BoolsMatch(o.AsOverride, other.AsOverride) {
		return false
	}
	if !util.BoolsMatch(o.DefaultOriginate, other.DefaultOriginate) {
		return false
	}
	if !util.StringsMatch(o.DefaultOriginateMap, other.DefaultOriginateMap) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !o.MaximumPrefix.matches(other.MaximumPrefix) {
		return false
	}
	if !o.NextHop.matches(other.NextHop) {
		return false
	}
	if !o.Orf.matches(other.Orf) {
		return false
	}
	if !o.RemovePrivateAS.matches(other.RemovePrivateAS) {
		return false
	}
	if !util.BoolsMatch(o.RouteReflectorClient, other.RouteReflectorClient) {
		return false
	}
	if !o.SendCommunity.matches(other.SendCommunity) {
		return false
	}
	if !util.BoolsMatch(o.SoftReconfigWithStoredInfo, other.SoftReconfigWithStoredInfo) {
		return false
	}

	return true
}

func (o *Ipv6UnicastAddPath) matches(other *Ipv6UnicastAddPath) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.TxAllPaths, other.TxAllPaths) {
		return false
	}
	if !util.BoolsMatch(o.TxBestpathPerAS, other.TxBestpathPerAS) {
		return false
	}

	return true
}

func (o *Ipv6UnicastAllowasIn) matches(other *Ipv6UnicastAllowasIn) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Occurrence, other.Occurrence) {
		return false
	}
	if !o.Origin.matches(other.Origin) {
		return false
	}

	return true
}

func (o *Ipv6UnicastAllowasInOrigin) matches(other *Ipv6UnicastAllowasInOrigin) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6UnicastMaximumPrefix) matches(other *Ipv6UnicastMaximumPrefix) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Action.matches(other.Action) {
		return false
	}
	if !util.Ints64Match(o.NumPrefixes, other.NumPrefixes) {
		return false
	}
	if !util.Ints64Match(o.Threshold, other.Threshold) {
		return false
	}

	return true
}

func (o *Ipv6UnicastMaximumPrefixAction) matches(other *Ipv6UnicastMaximumPrefixAction) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Restart.matches(other.Restart) {
		return false
	}
	if !o.WarningOnly.matches(other.WarningOnly) {
		return false
	}

	return true
}

func (o *Ipv6UnicastMaximumPrefixActionRestart) matches(other *Ipv6UnicastMaximumPrefixActionRestart) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Interval, other.Interval) {
		return false
	}

	return true
}

func (o *Ipv6UnicastMaximumPrefixActionWarningOnly) matches(other *Ipv6UnicastMaximumPrefixActionWarningOnly) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6UnicastNextHop) matches(other *Ipv6UnicastNextHop) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Self.matches(other.Self) {
		return false
	}
	if !o.SelfForce.matches(other.SelfForce) {
		return false
	}

	return true
}

func (o *Ipv6UnicastNextHopSelf) matches(other *Ipv6UnicastNextHopSelf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6UnicastNextHopSelfForce) matches(other *Ipv6UnicastNextHopSelfForce) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6UnicastOrf) matches(other *Ipv6UnicastOrf) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.OrfPrefixList, other.OrfPrefixList) {
		return false
	}

	return true
}

func (o *Ipv6UnicastRemovePrivateAS) matches(other *Ipv6UnicastRemovePrivateAS) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.All.matches(other.All) {
		return false
	}
	if !o.ReplaceAS.matches(other.ReplaceAS) {
		return false
	}

	return true
}

func (o *Ipv6UnicastRemovePrivateASAll) matches(other *Ipv6UnicastRemovePrivateASAll) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6UnicastRemovePrivateASReplaceAS) matches(other *Ipv6UnicastRemovePrivateASReplaceAS) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6UnicastSendCommunity) matches(other *Ipv6UnicastSendCommunity) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.All.matches(other.All) {
		return false
	}
	if !o.Both.matches(other.Both) {
		return false
	}
	if !o.Extended.matches(other.Extended) {
		return false
	}
	if !o.Large.matches(other.Large) {
		return false
	}
	if !o.Standard.matches(other.Standard) {
		return false
	}

	return true
}

func (o *Ipv6UnicastSendCommunityAll) matches(other *Ipv6UnicastSendCommunityAll) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6UnicastSendCommunityBoth) matches(other *Ipv6UnicastSendCommunityBoth) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6UnicastSendCommunityExtended) matches(other *Ipv6UnicastSendCommunityExtended) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6UnicastSendCommunityLarge) matches(other *Ipv6UnicastSendCommunityLarge) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *Ipv6UnicastSendCommunityStandard) matches(other *Ipv6UnicastSendCommunityStandard) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
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
