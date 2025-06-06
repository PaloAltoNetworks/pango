package ipsec

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
	suffix = []string{"network", "tunnel", "ipsec", "$name"}
)

type Entry struct {
	Name                   string
	AntiReplay             *bool
	AntiReplayWindow       *string
	Comment                *string
	CopyFlowLabel          *bool
	CopyTos                *bool
	Disabled               *bool
	EnableGreEncapsulation *bool
	IpsecMode              *string
	Ipv6                   *bool
	TunnelInterface        *string
	TunnelMonitor          *TunnelMonitor
	AutoKey                *AutoKey
	GlobalProtectSatellite *GlobalProtectSatellite
	ManualKey              *ManualKey
	Misc                   []generic.Xml
}
type TunnelMonitor struct {
	DestinationIp        *string
	Enable               *bool
	ProxyId              *string
	TunnelMonitorProfile *string
	Misc                 []generic.Xml
}
type AutoKey struct {
	IkeGateway         []AutoKeyIkeGateway
	IpsecCryptoProfile *string
	ProxyId            []AutoKeyProxyId
	ProxyIdV6          []AutoKeyProxyIdV6
	Misc               []generic.Xml
}
type AutoKeyIkeGateway struct {
	Name string
	Misc []generic.Xml
}
type AutoKeyProxyId struct {
	Name     string
	Local    *string
	Remote   *string
	Protocol *AutoKeyProxyIdProtocol
	Misc     []generic.Xml
}
type AutoKeyProxyIdProtocol struct {
	Number *int64
	Any    *AutoKeyProxyIdProtocolAny
	Tcp    *AutoKeyProxyIdProtocolTcp
	Udp    *AutoKeyProxyIdProtocolUdp
	Misc   []generic.Xml
}
type AutoKeyProxyIdProtocolAny struct {
	Misc []generic.Xml
}
type AutoKeyProxyIdProtocolTcp struct {
	LocalPort  *int64
	RemotePort *int64
	Misc       []generic.Xml
}
type AutoKeyProxyIdProtocolUdp struct {
	LocalPort  *int64
	RemotePort *int64
	Misc       []generic.Xml
}
type AutoKeyProxyIdV6 struct {
	Name     string
	Local    *string
	Remote   *string
	Protocol *AutoKeyProxyIdV6Protocol
	Misc     []generic.Xml
}
type AutoKeyProxyIdV6Protocol struct {
	Number *int64
	Any    *AutoKeyProxyIdV6ProtocolAny
	Tcp    *AutoKeyProxyIdV6ProtocolTcp
	Udp    *AutoKeyProxyIdV6ProtocolUdp
	Misc   []generic.Xml
}
type AutoKeyProxyIdV6ProtocolAny struct {
	Misc []generic.Xml
}
type AutoKeyProxyIdV6ProtocolTcp struct {
	LocalPort  *int64
	RemotePort *int64
	Misc       []generic.Xml
}
type AutoKeyProxyIdV6ProtocolUdp struct {
	LocalPort  *int64
	RemotePort *int64
	Misc       []generic.Xml
}
type GlobalProtectSatellite struct {
	ExternalCa             *GlobalProtectSatelliteExternalCa
	Ipv6Preferred          *bool
	LocalAddress           *GlobalProtectSatelliteLocalAddress
	PortalAddress          *string
	PublishConnectedRoutes *GlobalProtectSatellitePublishConnectedRoutes
	PublishRoutes          []string
	Misc                   []generic.Xml
}
type GlobalProtectSatelliteExternalCa struct {
	CertificateProfile *string
	LocalCertificate   *string
	Misc               []generic.Xml
}
type GlobalProtectSatelliteLocalAddress struct {
	Interface  *string
	FloatingIp *GlobalProtectSatelliteLocalAddressFloatingIp
	Ip         *GlobalProtectSatelliteLocalAddressIp
	Misc       []generic.Xml
}
type GlobalProtectSatelliteLocalAddressFloatingIp struct {
	Ipv4 *string
	Ipv6 *string
	Misc []generic.Xml
}
type GlobalProtectSatelliteLocalAddressIp struct {
	Ipv4 *string
	Ipv6 *string
	Misc []generic.Xml
}
type GlobalProtectSatellitePublishConnectedRoutes struct {
	Enable *bool
	Misc   []generic.Xml
}
type ManualKey struct {
	LocalAddress *ManualKeyLocalAddress
	LocalSpi     *string
	PeerAddress  *ManualKeyPeerAddress
	RemoteSpi    *string
	Ah           *ManualKeyAh
	Esp          *ManualKeyEsp
	Misc         []generic.Xml
}
type ManualKeyLocalAddress struct {
	Interface  *string
	FloatingIp *string
	Ip         *string
	Misc       []generic.Xml
}
type ManualKeyPeerAddress struct {
	Ip   *string
	Misc []generic.Xml
}
type ManualKeyAh struct {
	Md5    *ManualKeyAhMd5
	Sha1   *ManualKeyAhSha1
	Sha256 *ManualKeyAhSha256
	Sha384 *ManualKeyAhSha384
	Sha512 *ManualKeyAhSha512
	Misc   []generic.Xml
}
type ManualKeyAhMd5 struct {
	Key  *string
	Misc []generic.Xml
}
type ManualKeyAhSha1 struct {
	Key  *string
	Misc []generic.Xml
}
type ManualKeyAhSha256 struct {
	Key  *string
	Misc []generic.Xml
}
type ManualKeyAhSha384 struct {
	Key  *string
	Misc []generic.Xml
}
type ManualKeyAhSha512 struct {
	Key  *string
	Misc []generic.Xml
}
type ManualKeyEsp struct {
	Authentication *ManualKeyEspAuthentication
	Encryption     *ManualKeyEspEncryption
	Misc           []generic.Xml
}
type ManualKeyEspAuthentication struct {
	Md5    *ManualKeyEspAuthenticationMd5
	None   *ManualKeyEspAuthenticationNone
	Sha1   *ManualKeyEspAuthenticationSha1
	Sha256 *ManualKeyEspAuthenticationSha256
	Sha384 *ManualKeyEspAuthenticationSha384
	Sha512 *ManualKeyEspAuthenticationSha512
	Misc   []generic.Xml
}
type ManualKeyEspAuthenticationMd5 struct {
	Key  *string
	Misc []generic.Xml
}
type ManualKeyEspAuthenticationNone struct {
	Misc []generic.Xml
}
type ManualKeyEspAuthenticationSha1 struct {
	Key  *string
	Misc []generic.Xml
}
type ManualKeyEspAuthenticationSha256 struct {
	Key  *string
	Misc []generic.Xml
}
type ManualKeyEspAuthenticationSha384 struct {
	Key  *string
	Misc []generic.Xml
}
type ManualKeyEspAuthenticationSha512 struct {
	Key  *string
	Misc []generic.Xml
}
type ManualKeyEspEncryption struct {
	Algorithm *string
	Key       *string
	Misc      []generic.Xml
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
	XMLName                xml.Name                   `xml:"entry"`
	Name                   string                     `xml:"name,attr"`
	AntiReplay             *string                    `xml:"anti-replay,omitempty"`
	AntiReplayWindow       *string                    `xml:"anti-replay-window,omitempty"`
	Comment                *string                    `xml:"comment,omitempty"`
	CopyFlowLabel          *string                    `xml:"copy-flow-label,omitempty"`
	CopyTos                *string                    `xml:"copy-tos,omitempty"`
	Disabled               *string                    `xml:"disabled,omitempty"`
	EnableGreEncapsulation *string                    `xml:"enable-gre-encapsulation,omitempty"`
	IpsecMode              *string                    `xml:"ipsec-mode,omitempty"`
	Ipv6                   *string                    `xml:"ipv6,omitempty"`
	TunnelInterface        *string                    `xml:"tunnel-interface,omitempty"`
	TunnelMonitor          *tunnelMonitorXml          `xml:"tunnel-monitor,omitempty"`
	AutoKey                *autoKeyXml                `xml:"auto-key,omitempty"`
	GlobalProtectSatellite *globalProtectSatelliteXml `xml:"global-protect-satellite,omitempty"`
	ManualKey              *manualKeyXml              `xml:"manual-key,omitempty"`
	Misc                   []generic.Xml              `xml:",any"`
}
type tunnelMonitorXml struct {
	DestinationIp        *string       `xml:"destination-ip,omitempty"`
	Enable               *string       `xml:"enable,omitempty"`
	ProxyId              *string       `xml:"proxy-id,omitempty"`
	TunnelMonitorProfile *string       `xml:"tunnel-monitor-profile,omitempty"`
	Misc                 []generic.Xml `xml:",any"`
}
type autoKeyXml struct {
	IkeGateway         *autoKeyIkeGatewayContainerXml `xml:"ike-gateway,omitempty"`
	IpsecCryptoProfile *string                        `xml:"ipsec-crypto-profile,omitempty"`
	ProxyId            *autoKeyProxyIdContainerXml    `xml:"proxy-id,omitempty"`
	ProxyIdV6          *autoKeyProxyIdV6ContainerXml  `xml:"proxy-id-v6,omitempty"`
	Misc               []generic.Xml                  `xml:",any"`
}
type autoKeyIkeGatewayContainerXml struct {
	Entries []autoKeyIkeGatewayXml `xml:"entry"`
}
type autoKeyIkeGatewayXml struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}
type autoKeyProxyIdContainerXml struct {
	Entries []autoKeyProxyIdXml `xml:"entry"`
}
type autoKeyProxyIdXml struct {
	XMLName  xml.Name                   `xml:"entry"`
	Name     string                     `xml:"name,attr"`
	Local    *string                    `xml:"local,omitempty"`
	Remote   *string                    `xml:"remote,omitempty"`
	Protocol *autoKeyProxyIdProtocolXml `xml:"protocol,omitempty"`
	Misc     []generic.Xml              `xml:",any"`
}
type autoKeyProxyIdProtocolXml struct {
	Number *int64                        `xml:"number,omitempty"`
	Any    *autoKeyProxyIdProtocolAnyXml `xml:"any,omitempty"`
	Tcp    *autoKeyProxyIdProtocolTcpXml `xml:"tcp,omitempty"`
	Udp    *autoKeyProxyIdProtocolUdpXml `xml:"udp,omitempty"`
	Misc   []generic.Xml                 `xml:",any"`
}
type autoKeyProxyIdProtocolAnyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type autoKeyProxyIdProtocolTcpXml struct {
	LocalPort  *int64        `xml:"local-port,omitempty"`
	RemotePort *int64        `xml:"remote-port,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type autoKeyProxyIdProtocolUdpXml struct {
	LocalPort  *int64        `xml:"local-port,omitempty"`
	RemotePort *int64        `xml:"remote-port,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type autoKeyProxyIdV6ContainerXml struct {
	Entries []autoKeyProxyIdV6Xml `xml:"entry"`
}
type autoKeyProxyIdV6Xml struct {
	XMLName  xml.Name                     `xml:"entry"`
	Name     string                       `xml:"name,attr"`
	Local    *string                      `xml:"local,omitempty"`
	Remote   *string                      `xml:"remote,omitempty"`
	Protocol *autoKeyProxyIdV6ProtocolXml `xml:"protocol,omitempty"`
	Misc     []generic.Xml                `xml:",any"`
}
type autoKeyProxyIdV6ProtocolXml struct {
	Number *int64                          `xml:"number,omitempty"`
	Any    *autoKeyProxyIdV6ProtocolAnyXml `xml:"any,omitempty"`
	Tcp    *autoKeyProxyIdV6ProtocolTcpXml `xml:"tcp,omitempty"`
	Udp    *autoKeyProxyIdV6ProtocolUdpXml `xml:"udp,omitempty"`
	Misc   []generic.Xml                   `xml:",any"`
}
type autoKeyProxyIdV6ProtocolAnyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type autoKeyProxyIdV6ProtocolTcpXml struct {
	LocalPort  *int64        `xml:"local-port,omitempty"`
	RemotePort *int64        `xml:"remote-port,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type autoKeyProxyIdV6ProtocolUdpXml struct {
	LocalPort  *int64        `xml:"local-port,omitempty"`
	RemotePort *int64        `xml:"remote-port,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type globalProtectSatelliteXml struct {
	ExternalCa             *globalProtectSatelliteExternalCaXml             `xml:"external-ca,omitempty"`
	Ipv6Preferred          *string                                          `xml:"ipv6-preferred,omitempty"`
	LocalAddress           *globalProtectSatelliteLocalAddressXml           `xml:"local-address,omitempty"`
	PortalAddress          *string                                          `xml:"portal-address,omitempty"`
	PublishConnectedRoutes *globalProtectSatellitePublishConnectedRoutesXml `xml:"publish-connected-routes,omitempty"`
	PublishRoutes          *util.MemberType                                 `xml:"publish-routes,omitempty"`
	Misc                   []generic.Xml                                    `xml:",any"`
}
type globalProtectSatelliteExternalCaXml struct {
	CertificateProfile *string       `xml:"certificate-profile,omitempty"`
	LocalCertificate   *string       `xml:"local-certificate,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type globalProtectSatelliteLocalAddressXml struct {
	Interface  *string                                          `xml:"interface,omitempty"`
	FloatingIp *globalProtectSatelliteLocalAddressFloatingIpXml `xml:"floating-ip,omitempty"`
	Ip         *globalProtectSatelliteLocalAddressIpXml         `xml:"ip,omitempty"`
	Misc       []generic.Xml                                    `xml:",any"`
}
type globalProtectSatelliteLocalAddressFloatingIpXml struct {
	Ipv4 *string       `xml:"ipv4,omitempty"`
	Ipv6 *string       `xml:"ipv6,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type globalProtectSatelliteLocalAddressIpXml struct {
	Ipv4 *string       `xml:"ipv4,omitempty"`
	Ipv6 *string       `xml:"ipv6,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type globalProtectSatellitePublishConnectedRoutesXml struct {
	Enable *string       `xml:"enable,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type manualKeyXml struct {
	LocalAddress *manualKeyLocalAddressXml `xml:"local-address,omitempty"`
	LocalSpi     *string                   `xml:"local-spi,omitempty"`
	PeerAddress  *manualKeyPeerAddressXml  `xml:"peer-address,omitempty"`
	RemoteSpi    *string                   `xml:"remote-spi,omitempty"`
	Ah           *manualKeyAhXml           `xml:"ah,omitempty"`
	Esp          *manualKeyEspXml          `xml:"esp,omitempty"`
	Misc         []generic.Xml             `xml:",any"`
}
type manualKeyLocalAddressXml struct {
	Interface  *string       `xml:"interface,omitempty"`
	FloatingIp *string       `xml:"floating-ip,omitempty"`
	Ip         *string       `xml:"ip,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type manualKeyPeerAddressXml struct {
	Ip   *string       `xml:"ip,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyAhXml struct {
	Md5    *manualKeyAhMd5Xml    `xml:"md5,omitempty"`
	Sha1   *manualKeyAhSha1Xml   `xml:"sha1,omitempty"`
	Sha256 *manualKeyAhSha256Xml `xml:"sha256,omitempty"`
	Sha384 *manualKeyAhSha384Xml `xml:"sha384,omitempty"`
	Sha512 *manualKeyAhSha512Xml `xml:"sha512,omitempty"`
	Misc   []generic.Xml         `xml:",any"`
}
type manualKeyAhMd5Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyAhSha1Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyAhSha256Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyAhSha384Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyAhSha512Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyEspXml struct {
	Authentication *manualKeyEspAuthenticationXml `xml:"authentication,omitempty"`
	Encryption     *manualKeyEspEncryptionXml     `xml:"encryption,omitempty"`
	Misc           []generic.Xml                  `xml:",any"`
}
type manualKeyEspAuthenticationXml struct {
	Md5    *manualKeyEspAuthenticationMd5Xml    `xml:"md5,omitempty"`
	None   *manualKeyEspAuthenticationNoneXml   `xml:"none,omitempty"`
	Sha1   *manualKeyEspAuthenticationSha1Xml   `xml:"sha1,omitempty"`
	Sha256 *manualKeyEspAuthenticationSha256Xml `xml:"sha256,omitempty"`
	Sha384 *manualKeyEspAuthenticationSha384Xml `xml:"sha384,omitempty"`
	Sha512 *manualKeyEspAuthenticationSha512Xml `xml:"sha512,omitempty"`
	Misc   []generic.Xml                        `xml:",any"`
}
type manualKeyEspAuthenticationMd5Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyEspAuthenticationNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type manualKeyEspAuthenticationSha1Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyEspAuthenticationSha256Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyEspAuthenticationSha384Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyEspAuthenticationSha512Xml struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyEspEncryptionXml struct {
	Algorithm *string       `xml:"algorithm,omitempty"`
	Key       *string       `xml:"key,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}
type entryXml_11_0_2 struct {
	XMLName                xml.Name                          `xml:"entry"`
	Name                   string                            `xml:"name,attr"`
	AntiReplay             *string                           `xml:"anti-replay,omitempty"`
	AntiReplayWindow       *string                           `xml:"anti-replay-window,omitempty"`
	Comment                *string                           `xml:"comment,omitempty"`
	CopyFlowLabel          *string                           `xml:"copy-flow-label,omitempty"`
	CopyTos                *string                           `xml:"copy-tos,omitempty"`
	Disabled               *string                           `xml:"disabled,omitempty"`
	EnableGreEncapsulation *string                           `xml:"enable-gre-encapsulation,omitempty"`
	IpsecMode              *string                           `xml:"ipsec-mode,omitempty"`
	Ipv6                   *string                           `xml:"ipv6,omitempty"`
	TunnelInterface        *string                           `xml:"tunnel-interface,omitempty"`
	TunnelMonitor          *tunnelMonitorXml_11_0_2          `xml:"tunnel-monitor,omitempty"`
	AutoKey                *autoKeyXml_11_0_2                `xml:"auto-key,omitempty"`
	GlobalProtectSatellite *globalProtectSatelliteXml_11_0_2 `xml:"global-protect-satellite,omitempty"`
	ManualKey              *manualKeyXml_11_0_2              `xml:"manual-key,omitempty"`
	Misc                   []generic.Xml                     `xml:",any"`
}
type tunnelMonitorXml_11_0_2 struct {
	DestinationIp        *string       `xml:"destination-ip,omitempty"`
	Enable               *string       `xml:"enable,omitempty"`
	ProxyId              *string       `xml:"proxy-id,omitempty"`
	TunnelMonitorProfile *string       `xml:"tunnel-monitor-profile,omitempty"`
	Misc                 []generic.Xml `xml:",any"`
}
type autoKeyXml_11_0_2 struct {
	IkeGateway         *autoKeyIkeGatewayContainerXml_11_0_2 `xml:"ike-gateway,omitempty"`
	IpsecCryptoProfile *string                               `xml:"ipsec-crypto-profile,omitempty"`
	ProxyId            *autoKeyProxyIdContainerXml_11_0_2    `xml:"proxy-id,omitempty"`
	ProxyIdV6          *autoKeyProxyIdV6ContainerXml_11_0_2  `xml:"proxy-id-v6,omitempty"`
	Misc               []generic.Xml                         `xml:",any"`
}
type autoKeyIkeGatewayContainerXml_11_0_2 struct {
	Entries []autoKeyIkeGatewayXml_11_0_2 `xml:"entry"`
}
type autoKeyIkeGatewayXml_11_0_2 struct {
	XMLName xml.Name      `xml:"entry"`
	Name    string        `xml:"name,attr"`
	Misc    []generic.Xml `xml:",any"`
}
type autoKeyProxyIdContainerXml_11_0_2 struct {
	Entries []autoKeyProxyIdXml_11_0_2 `xml:"entry"`
}
type autoKeyProxyIdXml_11_0_2 struct {
	XMLName  xml.Name                          `xml:"entry"`
	Name     string                            `xml:"name,attr"`
	Local    *string                           `xml:"local,omitempty"`
	Remote   *string                           `xml:"remote,omitempty"`
	Protocol *autoKeyProxyIdProtocolXml_11_0_2 `xml:"protocol,omitempty"`
	Misc     []generic.Xml                     `xml:",any"`
}
type autoKeyProxyIdProtocolXml_11_0_2 struct {
	Number *int64                               `xml:"number,omitempty"`
	Any    *autoKeyProxyIdProtocolAnyXml_11_0_2 `xml:"any,omitempty"`
	Tcp    *autoKeyProxyIdProtocolTcpXml_11_0_2 `xml:"tcp,omitempty"`
	Udp    *autoKeyProxyIdProtocolUdpXml_11_0_2 `xml:"udp,omitempty"`
	Misc   []generic.Xml                        `xml:",any"`
}
type autoKeyProxyIdProtocolAnyXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type autoKeyProxyIdProtocolTcpXml_11_0_2 struct {
	LocalPort  *int64        `xml:"local-port,omitempty"`
	RemotePort *int64        `xml:"remote-port,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type autoKeyProxyIdProtocolUdpXml_11_0_2 struct {
	LocalPort  *int64        `xml:"local-port,omitempty"`
	RemotePort *int64        `xml:"remote-port,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type autoKeyProxyIdV6ContainerXml_11_0_2 struct {
	Entries []autoKeyProxyIdV6Xml_11_0_2 `xml:"entry"`
}
type autoKeyProxyIdV6Xml_11_0_2 struct {
	XMLName  xml.Name                            `xml:"entry"`
	Name     string                              `xml:"name,attr"`
	Local    *string                             `xml:"local,omitempty"`
	Remote   *string                             `xml:"remote,omitempty"`
	Protocol *autoKeyProxyIdV6ProtocolXml_11_0_2 `xml:"protocol,omitempty"`
	Misc     []generic.Xml                       `xml:",any"`
}
type autoKeyProxyIdV6ProtocolXml_11_0_2 struct {
	Number *int64                                 `xml:"number,omitempty"`
	Any    *autoKeyProxyIdV6ProtocolAnyXml_11_0_2 `xml:"any,omitempty"`
	Tcp    *autoKeyProxyIdV6ProtocolTcpXml_11_0_2 `xml:"tcp,omitempty"`
	Udp    *autoKeyProxyIdV6ProtocolUdpXml_11_0_2 `xml:"udp,omitempty"`
	Misc   []generic.Xml                          `xml:",any"`
}
type autoKeyProxyIdV6ProtocolAnyXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type autoKeyProxyIdV6ProtocolTcpXml_11_0_2 struct {
	LocalPort  *int64        `xml:"local-port,omitempty"`
	RemotePort *int64        `xml:"remote-port,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type autoKeyProxyIdV6ProtocolUdpXml_11_0_2 struct {
	LocalPort  *int64        `xml:"local-port,omitempty"`
	RemotePort *int64        `xml:"remote-port,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type globalProtectSatelliteXml_11_0_2 struct {
	ExternalCa             *globalProtectSatelliteExternalCaXml_11_0_2             `xml:"external-ca,omitempty"`
	Ipv6Preferred          *string                                                 `xml:"ipv6-preferred,omitempty"`
	LocalAddress           *globalProtectSatelliteLocalAddressXml_11_0_2           `xml:"local-address,omitempty"`
	PortalAddress          *string                                                 `xml:"portal-address,omitempty"`
	PublishConnectedRoutes *globalProtectSatellitePublishConnectedRoutesXml_11_0_2 `xml:"publish-connected-routes,omitempty"`
	PublishRoutes          *util.MemberType                                        `xml:"publish-routes,omitempty"`
	Misc                   []generic.Xml                                           `xml:",any"`
}
type globalProtectSatelliteExternalCaXml_11_0_2 struct {
	CertificateProfile *string       `xml:"certificate-profile,omitempty"`
	LocalCertificate   *string       `xml:"local-certificate,omitempty"`
	Misc               []generic.Xml `xml:",any"`
}
type globalProtectSatelliteLocalAddressXml_11_0_2 struct {
	Interface  *string                                                 `xml:"interface,omitempty"`
	FloatingIp *globalProtectSatelliteLocalAddressFloatingIpXml_11_0_2 `xml:"floating-ip,omitempty"`
	Ip         *globalProtectSatelliteLocalAddressIpXml_11_0_2         `xml:"ip,omitempty"`
	Misc       []generic.Xml                                           `xml:",any"`
}
type globalProtectSatelliteLocalAddressFloatingIpXml_11_0_2 struct {
	Ipv4 *string       `xml:"ipv4,omitempty"`
	Ipv6 *string       `xml:"ipv6,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type globalProtectSatelliteLocalAddressIpXml_11_0_2 struct {
	Ipv4 *string       `xml:"ipv4,omitempty"`
	Ipv6 *string       `xml:"ipv6,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type globalProtectSatellitePublishConnectedRoutesXml_11_0_2 struct {
	Enable *string       `xml:"enable,omitempty"`
	Misc   []generic.Xml `xml:",any"`
}
type manualKeyXml_11_0_2 struct {
	LocalAddress *manualKeyLocalAddressXml_11_0_2 `xml:"local-address,omitempty"`
	LocalSpi     *string                          `xml:"local-spi,omitempty"`
	PeerAddress  *manualKeyPeerAddressXml_11_0_2  `xml:"peer-address,omitempty"`
	RemoteSpi    *string                          `xml:"remote-spi,omitempty"`
	Ah           *manualKeyAhXml_11_0_2           `xml:"ah,omitempty"`
	Esp          *manualKeyEspXml_11_0_2          `xml:"esp,omitempty"`
	Misc         []generic.Xml                    `xml:",any"`
}
type manualKeyLocalAddressXml_11_0_2 struct {
	Interface  *string       `xml:"interface,omitempty"`
	FloatingIp *string       `xml:"floating-ip,omitempty"`
	Ip         *string       `xml:"ip,omitempty"`
	Misc       []generic.Xml `xml:",any"`
}
type manualKeyPeerAddressXml_11_0_2 struct {
	Ip   *string       `xml:"ip,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyAhXml_11_0_2 struct {
	Md5    *manualKeyAhMd5Xml_11_0_2    `xml:"md5,omitempty"`
	Sha1   *manualKeyAhSha1Xml_11_0_2   `xml:"sha1,omitempty"`
	Sha256 *manualKeyAhSha256Xml_11_0_2 `xml:"sha256,omitempty"`
	Sha384 *manualKeyAhSha384Xml_11_0_2 `xml:"sha384,omitempty"`
	Sha512 *manualKeyAhSha512Xml_11_0_2 `xml:"sha512,omitempty"`
	Misc   []generic.Xml                `xml:",any"`
}
type manualKeyAhMd5Xml_11_0_2 struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyAhSha1Xml_11_0_2 struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyAhSha256Xml_11_0_2 struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyAhSha384Xml_11_0_2 struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyAhSha512Xml_11_0_2 struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyEspXml_11_0_2 struct {
	Authentication *manualKeyEspAuthenticationXml_11_0_2 `xml:"authentication,omitempty"`
	Encryption     *manualKeyEspEncryptionXml_11_0_2     `xml:"encryption,omitempty"`
	Misc           []generic.Xml                         `xml:",any"`
}
type manualKeyEspAuthenticationXml_11_0_2 struct {
	Md5    *manualKeyEspAuthenticationMd5Xml_11_0_2    `xml:"md5,omitempty"`
	None   *manualKeyEspAuthenticationNoneXml_11_0_2   `xml:"none,omitempty"`
	Sha1   *manualKeyEspAuthenticationSha1Xml_11_0_2   `xml:"sha1,omitempty"`
	Sha256 *manualKeyEspAuthenticationSha256Xml_11_0_2 `xml:"sha256,omitempty"`
	Sha384 *manualKeyEspAuthenticationSha384Xml_11_0_2 `xml:"sha384,omitempty"`
	Sha512 *manualKeyEspAuthenticationSha512Xml_11_0_2 `xml:"sha512,omitempty"`
	Misc   []generic.Xml                               `xml:",any"`
}
type manualKeyEspAuthenticationMd5Xml_11_0_2 struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyEspAuthenticationNoneXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type manualKeyEspAuthenticationSha1Xml_11_0_2 struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyEspAuthenticationSha256Xml_11_0_2 struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyEspAuthenticationSha384Xml_11_0_2 struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyEspAuthenticationSha512Xml_11_0_2 struct {
	Key  *string       `xml:"key,omitempty"`
	Misc []generic.Xml `xml:",any"`
}
type manualKeyEspEncryptionXml_11_0_2 struct {
	Algorithm *string       `xml:"algorithm,omitempty"`
	Key       *string       `xml:"key,omitempty"`
	Misc      []generic.Xml `xml:",any"`
}

func (o *entryXml) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.AntiReplay = util.YesNo(s.AntiReplay, nil)
	o.AntiReplayWindow = s.AntiReplayWindow
	o.Comment = s.Comment
	o.CopyFlowLabel = util.YesNo(s.CopyFlowLabel, nil)
	o.CopyTos = util.YesNo(s.CopyTos, nil)
	o.Disabled = util.YesNo(s.Disabled, nil)
	o.EnableGreEncapsulation = util.YesNo(s.EnableGreEncapsulation, nil)
	o.IpsecMode = s.IpsecMode
	o.Ipv6 = util.YesNo(s.Ipv6, nil)
	o.TunnelInterface = s.TunnelInterface
	if s.TunnelMonitor != nil {
		var obj tunnelMonitorXml
		obj.MarshalFromObject(*s.TunnelMonitor)
		o.TunnelMonitor = &obj
	}
	if s.AutoKey != nil {
		var obj autoKeyXml
		obj.MarshalFromObject(*s.AutoKey)
		o.AutoKey = &obj
	}
	if s.GlobalProtectSatellite != nil {
		var obj globalProtectSatelliteXml
		obj.MarshalFromObject(*s.GlobalProtectSatellite)
		o.GlobalProtectSatellite = &obj
	}
	if s.ManualKey != nil {
		var obj manualKeyXml
		obj.MarshalFromObject(*s.ManualKey)
		o.ManualKey = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml) UnmarshalToObject() (*Entry, error) {
	var tunnelMonitorVal *TunnelMonitor
	if o.TunnelMonitor != nil {
		obj, err := o.TunnelMonitor.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		tunnelMonitorVal = obj
	}
	var autoKeyVal *AutoKey
	if o.AutoKey != nil {
		obj, err := o.AutoKey.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		autoKeyVal = obj
	}
	var globalProtectSatelliteVal *GlobalProtectSatellite
	if o.GlobalProtectSatellite != nil {
		obj, err := o.GlobalProtectSatellite.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalProtectSatelliteVal = obj
	}
	var manualKeyVal *ManualKey
	if o.ManualKey != nil {
		obj, err := o.ManualKey.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualKeyVal = obj
	}

	result := &Entry{
		Name:                   o.Name,
		AntiReplay:             util.AsBool(o.AntiReplay, nil),
		AntiReplayWindow:       o.AntiReplayWindow,
		Comment:                o.Comment,
		CopyFlowLabel:          util.AsBool(o.CopyFlowLabel, nil),
		CopyTos:                util.AsBool(o.CopyTos, nil),
		Disabled:               util.AsBool(o.Disabled, nil),
		EnableGreEncapsulation: util.AsBool(o.EnableGreEncapsulation, nil),
		IpsecMode:              o.IpsecMode,
		Ipv6:                   util.AsBool(o.Ipv6, nil),
		TunnelInterface:        o.TunnelInterface,
		TunnelMonitor:          tunnelMonitorVal,
		AutoKey:                autoKeyVal,
		GlobalProtectSatellite: globalProtectSatelliteVal,
		ManualKey:              manualKeyVal,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *tunnelMonitorXml) MarshalFromObject(s TunnelMonitor) {
	o.DestinationIp = s.DestinationIp
	o.Enable = util.YesNo(s.Enable, nil)
	o.ProxyId = s.ProxyId
	o.TunnelMonitorProfile = s.TunnelMonitorProfile
	o.Misc = s.Misc
}

func (o tunnelMonitorXml) UnmarshalToObject() (*TunnelMonitor, error) {

	result := &TunnelMonitor{
		DestinationIp:        o.DestinationIp,
		Enable:               util.AsBool(o.Enable, nil),
		ProxyId:              o.ProxyId,
		TunnelMonitorProfile: o.TunnelMonitorProfile,
		Misc:                 o.Misc,
	}
	return result, nil
}
func (o *autoKeyXml) MarshalFromObject(s AutoKey) {
	if s.IkeGateway != nil {
		var objs []autoKeyIkeGatewayXml
		for _, elt := range s.IkeGateway {
			var obj autoKeyIkeGatewayXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.IkeGateway = &autoKeyIkeGatewayContainerXml{Entries: objs}
	}
	o.IpsecCryptoProfile = s.IpsecCryptoProfile
	if s.ProxyId != nil {
		var objs []autoKeyProxyIdXml
		for _, elt := range s.ProxyId {
			var obj autoKeyProxyIdXml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ProxyId = &autoKeyProxyIdContainerXml{Entries: objs}
	}
	if s.ProxyIdV6 != nil {
		var objs []autoKeyProxyIdV6Xml
		for _, elt := range s.ProxyIdV6 {
			var obj autoKeyProxyIdV6Xml
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ProxyIdV6 = &autoKeyProxyIdV6ContainerXml{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o autoKeyXml) UnmarshalToObject() (*AutoKey, error) {
	var ikeGatewayVal []AutoKeyIkeGateway
	if o.IkeGateway != nil {
		for _, elt := range o.IkeGateway.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			ikeGatewayVal = append(ikeGatewayVal, *obj)
		}
	}
	var proxyIdVal []AutoKeyProxyId
	if o.ProxyId != nil {
		for _, elt := range o.ProxyId.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			proxyIdVal = append(proxyIdVal, *obj)
		}
	}
	var proxyIdV6Val []AutoKeyProxyIdV6
	if o.ProxyIdV6 != nil {
		for _, elt := range o.ProxyIdV6.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			proxyIdV6Val = append(proxyIdV6Val, *obj)
		}
	}

	result := &AutoKey{
		IkeGateway:         ikeGatewayVal,
		IpsecCryptoProfile: o.IpsecCryptoProfile,
		ProxyId:            proxyIdVal,
		ProxyIdV6:          proxyIdV6Val,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *autoKeyIkeGatewayXml) MarshalFromObject(s AutoKeyIkeGateway) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o autoKeyIkeGatewayXml) UnmarshalToObject() (*AutoKeyIkeGateway, error) {

	result := &AutoKeyIkeGateway{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdXml) MarshalFromObject(s AutoKeyProxyId) {
	o.Name = s.Name
	o.Local = s.Local
	o.Remote = s.Remote
	if s.Protocol != nil {
		var obj autoKeyProxyIdProtocolXml
		obj.MarshalFromObject(*s.Protocol)
		o.Protocol = &obj
	}
	o.Misc = s.Misc
}

func (o autoKeyProxyIdXml) UnmarshalToObject() (*AutoKeyProxyId, error) {
	var protocolVal *AutoKeyProxyIdProtocol
	if o.Protocol != nil {
		obj, err := o.Protocol.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		protocolVal = obj
	}

	result := &AutoKeyProxyId{
		Name:     o.Name,
		Local:    o.Local,
		Remote:   o.Remote,
		Protocol: protocolVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdProtocolXml) MarshalFromObject(s AutoKeyProxyIdProtocol) {
	o.Number = s.Number
	if s.Any != nil {
		var obj autoKeyProxyIdProtocolAnyXml
		obj.MarshalFromObject(*s.Any)
		o.Any = &obj
	}
	if s.Tcp != nil {
		var obj autoKeyProxyIdProtocolTcpXml
		obj.MarshalFromObject(*s.Tcp)
		o.Tcp = &obj
	}
	if s.Udp != nil {
		var obj autoKeyProxyIdProtocolUdpXml
		obj.MarshalFromObject(*s.Udp)
		o.Udp = &obj
	}
	o.Misc = s.Misc
}

func (o autoKeyProxyIdProtocolXml) UnmarshalToObject() (*AutoKeyProxyIdProtocol, error) {
	var anyVal *AutoKeyProxyIdProtocolAny
	if o.Any != nil {
		obj, err := o.Any.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		anyVal = obj
	}
	var tcpVal *AutoKeyProxyIdProtocolTcp
	if o.Tcp != nil {
		obj, err := o.Tcp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		tcpVal = obj
	}
	var udpVal *AutoKeyProxyIdProtocolUdp
	if o.Udp != nil {
		obj, err := o.Udp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		udpVal = obj
	}

	result := &AutoKeyProxyIdProtocol{
		Number: o.Number,
		Any:    anyVal,
		Tcp:    tcpVal,
		Udp:    udpVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdProtocolAnyXml) MarshalFromObject(s AutoKeyProxyIdProtocolAny) {
	o.Misc = s.Misc
}

func (o autoKeyProxyIdProtocolAnyXml) UnmarshalToObject() (*AutoKeyProxyIdProtocolAny, error) {

	result := &AutoKeyProxyIdProtocolAny{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdProtocolTcpXml) MarshalFromObject(s AutoKeyProxyIdProtocolTcp) {
	o.LocalPort = s.LocalPort
	o.RemotePort = s.RemotePort
	o.Misc = s.Misc
}

func (o autoKeyProxyIdProtocolTcpXml) UnmarshalToObject() (*AutoKeyProxyIdProtocolTcp, error) {

	result := &AutoKeyProxyIdProtocolTcp{
		LocalPort:  o.LocalPort,
		RemotePort: o.RemotePort,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdProtocolUdpXml) MarshalFromObject(s AutoKeyProxyIdProtocolUdp) {
	o.LocalPort = s.LocalPort
	o.RemotePort = s.RemotePort
	o.Misc = s.Misc
}

func (o autoKeyProxyIdProtocolUdpXml) UnmarshalToObject() (*AutoKeyProxyIdProtocolUdp, error) {

	result := &AutoKeyProxyIdProtocolUdp{
		LocalPort:  o.LocalPort,
		RemotePort: o.RemotePort,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdV6Xml) MarshalFromObject(s AutoKeyProxyIdV6) {
	o.Name = s.Name
	o.Local = s.Local
	o.Remote = s.Remote
	if s.Protocol != nil {
		var obj autoKeyProxyIdV6ProtocolXml
		obj.MarshalFromObject(*s.Protocol)
		o.Protocol = &obj
	}
	o.Misc = s.Misc
}

func (o autoKeyProxyIdV6Xml) UnmarshalToObject() (*AutoKeyProxyIdV6, error) {
	var protocolVal *AutoKeyProxyIdV6Protocol
	if o.Protocol != nil {
		obj, err := o.Protocol.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		protocolVal = obj
	}

	result := &AutoKeyProxyIdV6{
		Name:     o.Name,
		Local:    o.Local,
		Remote:   o.Remote,
		Protocol: protocolVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdV6ProtocolXml) MarshalFromObject(s AutoKeyProxyIdV6Protocol) {
	o.Number = s.Number
	if s.Any != nil {
		var obj autoKeyProxyIdV6ProtocolAnyXml
		obj.MarshalFromObject(*s.Any)
		o.Any = &obj
	}
	if s.Tcp != nil {
		var obj autoKeyProxyIdV6ProtocolTcpXml
		obj.MarshalFromObject(*s.Tcp)
		o.Tcp = &obj
	}
	if s.Udp != nil {
		var obj autoKeyProxyIdV6ProtocolUdpXml
		obj.MarshalFromObject(*s.Udp)
		o.Udp = &obj
	}
	o.Misc = s.Misc
}

func (o autoKeyProxyIdV6ProtocolXml) UnmarshalToObject() (*AutoKeyProxyIdV6Protocol, error) {
	var anyVal *AutoKeyProxyIdV6ProtocolAny
	if o.Any != nil {
		obj, err := o.Any.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		anyVal = obj
	}
	var tcpVal *AutoKeyProxyIdV6ProtocolTcp
	if o.Tcp != nil {
		obj, err := o.Tcp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		tcpVal = obj
	}
	var udpVal *AutoKeyProxyIdV6ProtocolUdp
	if o.Udp != nil {
		obj, err := o.Udp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		udpVal = obj
	}

	result := &AutoKeyProxyIdV6Protocol{
		Number: o.Number,
		Any:    anyVal,
		Tcp:    tcpVal,
		Udp:    udpVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdV6ProtocolAnyXml) MarshalFromObject(s AutoKeyProxyIdV6ProtocolAny) {
	o.Misc = s.Misc
}

func (o autoKeyProxyIdV6ProtocolAnyXml) UnmarshalToObject() (*AutoKeyProxyIdV6ProtocolAny, error) {

	result := &AutoKeyProxyIdV6ProtocolAny{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdV6ProtocolTcpXml) MarshalFromObject(s AutoKeyProxyIdV6ProtocolTcp) {
	o.LocalPort = s.LocalPort
	o.RemotePort = s.RemotePort
	o.Misc = s.Misc
}

func (o autoKeyProxyIdV6ProtocolTcpXml) UnmarshalToObject() (*AutoKeyProxyIdV6ProtocolTcp, error) {

	result := &AutoKeyProxyIdV6ProtocolTcp{
		LocalPort:  o.LocalPort,
		RemotePort: o.RemotePort,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdV6ProtocolUdpXml) MarshalFromObject(s AutoKeyProxyIdV6ProtocolUdp) {
	o.LocalPort = s.LocalPort
	o.RemotePort = s.RemotePort
	o.Misc = s.Misc
}

func (o autoKeyProxyIdV6ProtocolUdpXml) UnmarshalToObject() (*AutoKeyProxyIdV6ProtocolUdp, error) {

	result := &AutoKeyProxyIdV6ProtocolUdp{
		LocalPort:  o.LocalPort,
		RemotePort: o.RemotePort,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *globalProtectSatelliteXml) MarshalFromObject(s GlobalProtectSatellite) {
	if s.ExternalCa != nil {
		var obj globalProtectSatelliteExternalCaXml
		obj.MarshalFromObject(*s.ExternalCa)
		o.ExternalCa = &obj
	}
	o.Ipv6Preferred = util.YesNo(s.Ipv6Preferred, nil)
	if s.LocalAddress != nil {
		var obj globalProtectSatelliteLocalAddressXml
		obj.MarshalFromObject(*s.LocalAddress)
		o.LocalAddress = &obj
	}
	o.PortalAddress = s.PortalAddress
	if s.PublishConnectedRoutes != nil {
		var obj globalProtectSatellitePublishConnectedRoutesXml
		obj.MarshalFromObject(*s.PublishConnectedRoutes)
		o.PublishConnectedRoutes = &obj
	}
	if s.PublishRoutes != nil {
		o.PublishRoutes = util.StrToMem(s.PublishRoutes)
	}
	o.Misc = s.Misc
}

func (o globalProtectSatelliteXml) UnmarshalToObject() (*GlobalProtectSatellite, error) {
	var externalCaVal *GlobalProtectSatelliteExternalCa
	if o.ExternalCa != nil {
		obj, err := o.ExternalCa.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		externalCaVal = obj
	}
	var localAddressVal *GlobalProtectSatelliteLocalAddress
	if o.LocalAddress != nil {
		obj, err := o.LocalAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localAddressVal = obj
	}
	var publishConnectedRoutesVal *GlobalProtectSatellitePublishConnectedRoutes
	if o.PublishConnectedRoutes != nil {
		obj, err := o.PublishConnectedRoutes.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		publishConnectedRoutesVal = obj
	}
	var publishRoutesVal []string
	if o.PublishRoutes != nil {
		publishRoutesVal = util.MemToStr(o.PublishRoutes)
	}

	result := &GlobalProtectSatellite{
		ExternalCa:             externalCaVal,
		Ipv6Preferred:          util.AsBool(o.Ipv6Preferred, nil),
		LocalAddress:           localAddressVal,
		PortalAddress:          o.PortalAddress,
		PublishConnectedRoutes: publishConnectedRoutesVal,
		PublishRoutes:          publishRoutesVal,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *globalProtectSatelliteExternalCaXml) MarshalFromObject(s GlobalProtectSatelliteExternalCa) {
	o.CertificateProfile = s.CertificateProfile
	o.LocalCertificate = s.LocalCertificate
	o.Misc = s.Misc
}

func (o globalProtectSatelliteExternalCaXml) UnmarshalToObject() (*GlobalProtectSatelliteExternalCa, error) {

	result := &GlobalProtectSatelliteExternalCa{
		CertificateProfile: o.CertificateProfile,
		LocalCertificate:   o.LocalCertificate,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *globalProtectSatelliteLocalAddressXml) MarshalFromObject(s GlobalProtectSatelliteLocalAddress) {
	o.Interface = s.Interface
	if s.FloatingIp != nil {
		var obj globalProtectSatelliteLocalAddressFloatingIpXml
		obj.MarshalFromObject(*s.FloatingIp)
		o.FloatingIp = &obj
	}
	if s.Ip != nil {
		var obj globalProtectSatelliteLocalAddressIpXml
		obj.MarshalFromObject(*s.Ip)
		o.Ip = &obj
	}
	o.Misc = s.Misc
}

func (o globalProtectSatelliteLocalAddressXml) UnmarshalToObject() (*GlobalProtectSatelliteLocalAddress, error) {
	var floatingIpVal *GlobalProtectSatelliteLocalAddressFloatingIp
	if o.FloatingIp != nil {
		obj, err := o.FloatingIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		floatingIpVal = obj
	}
	var ipVal *GlobalProtectSatelliteLocalAddressIp
	if o.Ip != nil {
		obj, err := o.Ip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipVal = obj
	}

	result := &GlobalProtectSatelliteLocalAddress{
		Interface:  o.Interface,
		FloatingIp: floatingIpVal,
		Ip:         ipVal,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *globalProtectSatelliteLocalAddressFloatingIpXml) MarshalFromObject(s GlobalProtectSatelliteLocalAddressFloatingIp) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
}

func (o globalProtectSatelliteLocalAddressFloatingIpXml) UnmarshalToObject() (*GlobalProtectSatelliteLocalAddressFloatingIp, error) {

	result := &GlobalProtectSatelliteLocalAddressFloatingIp{
		Ipv4: o.Ipv4,
		Ipv6: o.Ipv6,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *globalProtectSatelliteLocalAddressIpXml) MarshalFromObject(s GlobalProtectSatelliteLocalAddressIp) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
}

func (o globalProtectSatelliteLocalAddressIpXml) UnmarshalToObject() (*GlobalProtectSatelliteLocalAddressIp, error) {

	result := &GlobalProtectSatelliteLocalAddressIp{
		Ipv4: o.Ipv4,
		Ipv6: o.Ipv6,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *globalProtectSatellitePublishConnectedRoutesXml) MarshalFromObject(s GlobalProtectSatellitePublishConnectedRoutes) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Misc = s.Misc
}

func (o globalProtectSatellitePublishConnectedRoutesXml) UnmarshalToObject() (*GlobalProtectSatellitePublishConnectedRoutes, error) {

	result := &GlobalProtectSatellitePublishConnectedRoutes{
		Enable: util.AsBool(o.Enable, nil),
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *manualKeyXml) MarshalFromObject(s ManualKey) {
	if s.LocalAddress != nil {
		var obj manualKeyLocalAddressXml
		obj.MarshalFromObject(*s.LocalAddress)
		o.LocalAddress = &obj
	}
	o.LocalSpi = s.LocalSpi
	if s.PeerAddress != nil {
		var obj manualKeyPeerAddressXml
		obj.MarshalFromObject(*s.PeerAddress)
		o.PeerAddress = &obj
	}
	o.RemoteSpi = s.RemoteSpi
	if s.Ah != nil {
		var obj manualKeyAhXml
		obj.MarshalFromObject(*s.Ah)
		o.Ah = &obj
	}
	if s.Esp != nil {
		var obj manualKeyEspXml
		obj.MarshalFromObject(*s.Esp)
		o.Esp = &obj
	}
	o.Misc = s.Misc
}

func (o manualKeyXml) UnmarshalToObject() (*ManualKey, error) {
	var localAddressVal *ManualKeyLocalAddress
	if o.LocalAddress != nil {
		obj, err := o.LocalAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localAddressVal = obj
	}
	var peerAddressVal *ManualKeyPeerAddress
	if o.PeerAddress != nil {
		obj, err := o.PeerAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		peerAddressVal = obj
	}
	var ahVal *ManualKeyAh
	if o.Ah != nil {
		obj, err := o.Ah.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ahVal = obj
	}
	var espVal *ManualKeyEsp
	if o.Esp != nil {
		obj, err := o.Esp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		espVal = obj
	}

	result := &ManualKey{
		LocalAddress: localAddressVal,
		LocalSpi:     o.LocalSpi,
		PeerAddress:  peerAddressVal,
		RemoteSpi:    o.RemoteSpi,
		Ah:           ahVal,
		Esp:          espVal,
		Misc:         o.Misc,
	}
	return result, nil
}
func (o *manualKeyLocalAddressXml) MarshalFromObject(s ManualKeyLocalAddress) {
	o.Interface = s.Interface
	o.FloatingIp = s.FloatingIp
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o manualKeyLocalAddressXml) UnmarshalToObject() (*ManualKeyLocalAddress, error) {

	result := &ManualKeyLocalAddress{
		Interface:  o.Interface,
		FloatingIp: o.FloatingIp,
		Ip:         o.Ip,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *manualKeyPeerAddressXml) MarshalFromObject(s ManualKeyPeerAddress) {
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o manualKeyPeerAddressXml) UnmarshalToObject() (*ManualKeyPeerAddress, error) {

	result := &ManualKeyPeerAddress{
		Ip:   o.Ip,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyAhXml) MarshalFromObject(s ManualKeyAh) {
	if s.Md5 != nil {
		var obj manualKeyAhMd5Xml
		obj.MarshalFromObject(*s.Md5)
		o.Md5 = &obj
	}
	if s.Sha1 != nil {
		var obj manualKeyAhSha1Xml
		obj.MarshalFromObject(*s.Sha1)
		o.Sha1 = &obj
	}
	if s.Sha256 != nil {
		var obj manualKeyAhSha256Xml
		obj.MarshalFromObject(*s.Sha256)
		o.Sha256 = &obj
	}
	if s.Sha384 != nil {
		var obj manualKeyAhSha384Xml
		obj.MarshalFromObject(*s.Sha384)
		o.Sha384 = &obj
	}
	if s.Sha512 != nil {
		var obj manualKeyAhSha512Xml
		obj.MarshalFromObject(*s.Sha512)
		o.Sha512 = &obj
	}
	o.Misc = s.Misc
}

func (o manualKeyAhXml) UnmarshalToObject() (*ManualKeyAh, error) {
	var md5Val *ManualKeyAhMd5
	if o.Md5 != nil {
		obj, err := o.Md5.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		md5Val = obj
	}
	var sha1Val *ManualKeyAhSha1
	if o.Sha1 != nil {
		obj, err := o.Sha1.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha1Val = obj
	}
	var sha256Val *ManualKeyAhSha256
	if o.Sha256 != nil {
		obj, err := o.Sha256.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha256Val = obj
	}
	var sha384Val *ManualKeyAhSha384
	if o.Sha384 != nil {
		obj, err := o.Sha384.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha384Val = obj
	}
	var sha512Val *ManualKeyAhSha512
	if o.Sha512 != nil {
		obj, err := o.Sha512.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha512Val = obj
	}

	result := &ManualKeyAh{
		Md5:    md5Val,
		Sha1:   sha1Val,
		Sha256: sha256Val,
		Sha384: sha384Val,
		Sha512: sha512Val,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *manualKeyAhMd5Xml) MarshalFromObject(s ManualKeyAhMd5) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyAhMd5Xml) UnmarshalToObject() (*ManualKeyAhMd5, error) {

	result := &ManualKeyAhMd5{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyAhSha1Xml) MarshalFromObject(s ManualKeyAhSha1) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyAhSha1Xml) UnmarshalToObject() (*ManualKeyAhSha1, error) {

	result := &ManualKeyAhSha1{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyAhSha256Xml) MarshalFromObject(s ManualKeyAhSha256) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyAhSha256Xml) UnmarshalToObject() (*ManualKeyAhSha256, error) {

	result := &ManualKeyAhSha256{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyAhSha384Xml) MarshalFromObject(s ManualKeyAhSha384) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyAhSha384Xml) UnmarshalToObject() (*ManualKeyAhSha384, error) {

	result := &ManualKeyAhSha384{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyAhSha512Xml) MarshalFromObject(s ManualKeyAhSha512) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyAhSha512Xml) UnmarshalToObject() (*ManualKeyAhSha512, error) {

	result := &ManualKeyAhSha512{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspXml) MarshalFromObject(s ManualKeyEsp) {
	if s.Authentication != nil {
		var obj manualKeyEspAuthenticationXml
		obj.MarshalFromObject(*s.Authentication)
		o.Authentication = &obj
	}
	if s.Encryption != nil {
		var obj manualKeyEspEncryptionXml
		obj.MarshalFromObject(*s.Encryption)
		o.Encryption = &obj
	}
	o.Misc = s.Misc
}

func (o manualKeyEspXml) UnmarshalToObject() (*ManualKeyEsp, error) {
	var authenticationVal *ManualKeyEspAuthentication
	if o.Authentication != nil {
		obj, err := o.Authentication.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authenticationVal = obj
	}
	var encryptionVal *ManualKeyEspEncryption
	if o.Encryption != nil {
		obj, err := o.Encryption.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		encryptionVal = obj
	}

	result := &ManualKeyEsp{
		Authentication: authenticationVal,
		Encryption:     encryptionVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspAuthenticationXml) MarshalFromObject(s ManualKeyEspAuthentication) {
	if s.Md5 != nil {
		var obj manualKeyEspAuthenticationMd5Xml
		obj.MarshalFromObject(*s.Md5)
		o.Md5 = &obj
	}
	if s.None != nil {
		var obj manualKeyEspAuthenticationNoneXml
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.Sha1 != nil {
		var obj manualKeyEspAuthenticationSha1Xml
		obj.MarshalFromObject(*s.Sha1)
		o.Sha1 = &obj
	}
	if s.Sha256 != nil {
		var obj manualKeyEspAuthenticationSha256Xml
		obj.MarshalFromObject(*s.Sha256)
		o.Sha256 = &obj
	}
	if s.Sha384 != nil {
		var obj manualKeyEspAuthenticationSha384Xml
		obj.MarshalFromObject(*s.Sha384)
		o.Sha384 = &obj
	}
	if s.Sha512 != nil {
		var obj manualKeyEspAuthenticationSha512Xml
		obj.MarshalFromObject(*s.Sha512)
		o.Sha512 = &obj
	}
	o.Misc = s.Misc
}

func (o manualKeyEspAuthenticationXml) UnmarshalToObject() (*ManualKeyEspAuthentication, error) {
	var md5Val *ManualKeyEspAuthenticationMd5
	if o.Md5 != nil {
		obj, err := o.Md5.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		md5Val = obj
	}
	var noneVal *ManualKeyEspAuthenticationNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var sha1Val *ManualKeyEspAuthenticationSha1
	if o.Sha1 != nil {
		obj, err := o.Sha1.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha1Val = obj
	}
	var sha256Val *ManualKeyEspAuthenticationSha256
	if o.Sha256 != nil {
		obj, err := o.Sha256.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha256Val = obj
	}
	var sha384Val *ManualKeyEspAuthenticationSha384
	if o.Sha384 != nil {
		obj, err := o.Sha384.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha384Val = obj
	}
	var sha512Val *ManualKeyEspAuthenticationSha512
	if o.Sha512 != nil {
		obj, err := o.Sha512.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha512Val = obj
	}

	result := &ManualKeyEspAuthentication{
		Md5:    md5Val,
		None:   noneVal,
		Sha1:   sha1Val,
		Sha256: sha256Val,
		Sha384: sha384Val,
		Sha512: sha512Val,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspAuthenticationMd5Xml) MarshalFromObject(s ManualKeyEspAuthenticationMd5) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyEspAuthenticationMd5Xml) UnmarshalToObject() (*ManualKeyEspAuthenticationMd5, error) {

	result := &ManualKeyEspAuthenticationMd5{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspAuthenticationNoneXml) MarshalFromObject(s ManualKeyEspAuthenticationNone) {
	o.Misc = s.Misc
}

func (o manualKeyEspAuthenticationNoneXml) UnmarshalToObject() (*ManualKeyEspAuthenticationNone, error) {

	result := &ManualKeyEspAuthenticationNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspAuthenticationSha1Xml) MarshalFromObject(s ManualKeyEspAuthenticationSha1) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyEspAuthenticationSha1Xml) UnmarshalToObject() (*ManualKeyEspAuthenticationSha1, error) {

	result := &ManualKeyEspAuthenticationSha1{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspAuthenticationSha256Xml) MarshalFromObject(s ManualKeyEspAuthenticationSha256) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyEspAuthenticationSha256Xml) UnmarshalToObject() (*ManualKeyEspAuthenticationSha256, error) {

	result := &ManualKeyEspAuthenticationSha256{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspAuthenticationSha384Xml) MarshalFromObject(s ManualKeyEspAuthenticationSha384) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyEspAuthenticationSha384Xml) UnmarshalToObject() (*ManualKeyEspAuthenticationSha384, error) {

	result := &ManualKeyEspAuthenticationSha384{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspAuthenticationSha512Xml) MarshalFromObject(s ManualKeyEspAuthenticationSha512) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyEspAuthenticationSha512Xml) UnmarshalToObject() (*ManualKeyEspAuthenticationSha512, error) {

	result := &ManualKeyEspAuthenticationSha512{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspEncryptionXml) MarshalFromObject(s ManualKeyEspEncryption) {
	o.Algorithm = s.Algorithm
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyEspEncryptionXml) UnmarshalToObject() (*ManualKeyEspEncryption, error) {

	result := &ManualKeyEspEncryption{
		Algorithm: o.Algorithm,
		Key:       o.Key,
		Misc:      o.Misc,
	}
	return result, nil
}
func (o *entryXml_11_0_2) MarshalFromObject(s Entry) {
	o.Name = s.Name
	o.AntiReplay = util.YesNo(s.AntiReplay, nil)
	o.AntiReplayWindow = s.AntiReplayWindow
	o.Comment = s.Comment
	o.CopyFlowLabel = util.YesNo(s.CopyFlowLabel, nil)
	o.CopyTos = util.YesNo(s.CopyTos, nil)
	o.Disabled = util.YesNo(s.Disabled, nil)
	o.EnableGreEncapsulation = util.YesNo(s.EnableGreEncapsulation, nil)
	o.IpsecMode = s.IpsecMode
	o.Ipv6 = util.YesNo(s.Ipv6, nil)
	o.TunnelInterface = s.TunnelInterface
	if s.TunnelMonitor != nil {
		var obj tunnelMonitorXml_11_0_2
		obj.MarshalFromObject(*s.TunnelMonitor)
		o.TunnelMonitor = &obj
	}
	if s.AutoKey != nil {
		var obj autoKeyXml_11_0_2
		obj.MarshalFromObject(*s.AutoKey)
		o.AutoKey = &obj
	}
	if s.GlobalProtectSatellite != nil {
		var obj globalProtectSatelliteXml_11_0_2
		obj.MarshalFromObject(*s.GlobalProtectSatellite)
		o.GlobalProtectSatellite = &obj
	}
	if s.ManualKey != nil {
		var obj manualKeyXml_11_0_2
		obj.MarshalFromObject(*s.ManualKey)
		o.ManualKey = &obj
	}
	o.Misc = s.Misc
}

func (o entryXml_11_0_2) UnmarshalToObject() (*Entry, error) {
	var tunnelMonitorVal *TunnelMonitor
	if o.TunnelMonitor != nil {
		obj, err := o.TunnelMonitor.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		tunnelMonitorVal = obj
	}
	var autoKeyVal *AutoKey
	if o.AutoKey != nil {
		obj, err := o.AutoKey.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		autoKeyVal = obj
	}
	var globalProtectSatelliteVal *GlobalProtectSatellite
	if o.GlobalProtectSatellite != nil {
		obj, err := o.GlobalProtectSatellite.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		globalProtectSatelliteVal = obj
	}
	var manualKeyVal *ManualKey
	if o.ManualKey != nil {
		obj, err := o.ManualKey.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		manualKeyVal = obj
	}

	result := &Entry{
		Name:                   o.Name,
		AntiReplay:             util.AsBool(o.AntiReplay, nil),
		AntiReplayWindow:       o.AntiReplayWindow,
		Comment:                o.Comment,
		CopyFlowLabel:          util.AsBool(o.CopyFlowLabel, nil),
		CopyTos:                util.AsBool(o.CopyTos, nil),
		Disabled:               util.AsBool(o.Disabled, nil),
		EnableGreEncapsulation: util.AsBool(o.EnableGreEncapsulation, nil),
		IpsecMode:              o.IpsecMode,
		Ipv6:                   util.AsBool(o.Ipv6, nil),
		TunnelInterface:        o.TunnelInterface,
		TunnelMonitor:          tunnelMonitorVal,
		AutoKey:                autoKeyVal,
		GlobalProtectSatellite: globalProtectSatelliteVal,
		ManualKey:              manualKeyVal,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *tunnelMonitorXml_11_0_2) MarshalFromObject(s TunnelMonitor) {
	o.DestinationIp = s.DestinationIp
	o.Enable = util.YesNo(s.Enable, nil)
	o.ProxyId = s.ProxyId
	o.TunnelMonitorProfile = s.TunnelMonitorProfile
	o.Misc = s.Misc
}

func (o tunnelMonitorXml_11_0_2) UnmarshalToObject() (*TunnelMonitor, error) {

	result := &TunnelMonitor{
		DestinationIp:        o.DestinationIp,
		Enable:               util.AsBool(o.Enable, nil),
		ProxyId:              o.ProxyId,
		TunnelMonitorProfile: o.TunnelMonitorProfile,
		Misc:                 o.Misc,
	}
	return result, nil
}
func (o *autoKeyXml_11_0_2) MarshalFromObject(s AutoKey) {
	if s.IkeGateway != nil {
		var objs []autoKeyIkeGatewayXml_11_0_2
		for _, elt := range s.IkeGateway {
			var obj autoKeyIkeGatewayXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.IkeGateway = &autoKeyIkeGatewayContainerXml_11_0_2{Entries: objs}
	}
	o.IpsecCryptoProfile = s.IpsecCryptoProfile
	if s.ProxyId != nil {
		var objs []autoKeyProxyIdXml_11_0_2
		for _, elt := range s.ProxyId {
			var obj autoKeyProxyIdXml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ProxyId = &autoKeyProxyIdContainerXml_11_0_2{Entries: objs}
	}
	if s.ProxyIdV6 != nil {
		var objs []autoKeyProxyIdV6Xml_11_0_2
		for _, elt := range s.ProxyIdV6 {
			var obj autoKeyProxyIdV6Xml_11_0_2
			obj.MarshalFromObject(elt)
			objs = append(objs, obj)
		}
		o.ProxyIdV6 = &autoKeyProxyIdV6ContainerXml_11_0_2{Entries: objs}
	}
	o.Misc = s.Misc
}

func (o autoKeyXml_11_0_2) UnmarshalToObject() (*AutoKey, error) {
	var ikeGatewayVal []AutoKeyIkeGateway
	if o.IkeGateway != nil {
		for _, elt := range o.IkeGateway.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			ikeGatewayVal = append(ikeGatewayVal, *obj)
		}
	}
	var proxyIdVal []AutoKeyProxyId
	if o.ProxyId != nil {
		for _, elt := range o.ProxyId.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			proxyIdVal = append(proxyIdVal, *obj)
		}
	}
	var proxyIdV6Val []AutoKeyProxyIdV6
	if o.ProxyIdV6 != nil {
		for _, elt := range o.ProxyIdV6.Entries {
			obj, err := elt.UnmarshalToObject()
			if err != nil {
				return nil, err
			}
			proxyIdV6Val = append(proxyIdV6Val, *obj)
		}
	}

	result := &AutoKey{
		IkeGateway:         ikeGatewayVal,
		IpsecCryptoProfile: o.IpsecCryptoProfile,
		ProxyId:            proxyIdVal,
		ProxyIdV6:          proxyIdV6Val,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *autoKeyIkeGatewayXml_11_0_2) MarshalFromObject(s AutoKeyIkeGateway) {
	o.Name = s.Name
	o.Misc = s.Misc
}

func (o autoKeyIkeGatewayXml_11_0_2) UnmarshalToObject() (*AutoKeyIkeGateway, error) {

	result := &AutoKeyIkeGateway{
		Name: o.Name,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdXml_11_0_2) MarshalFromObject(s AutoKeyProxyId) {
	o.Name = s.Name
	o.Local = s.Local
	o.Remote = s.Remote
	if s.Protocol != nil {
		var obj autoKeyProxyIdProtocolXml_11_0_2
		obj.MarshalFromObject(*s.Protocol)
		o.Protocol = &obj
	}
	o.Misc = s.Misc
}

func (o autoKeyProxyIdXml_11_0_2) UnmarshalToObject() (*AutoKeyProxyId, error) {
	var protocolVal *AutoKeyProxyIdProtocol
	if o.Protocol != nil {
		obj, err := o.Protocol.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		protocolVal = obj
	}

	result := &AutoKeyProxyId{
		Name:     o.Name,
		Local:    o.Local,
		Remote:   o.Remote,
		Protocol: protocolVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdProtocolXml_11_0_2) MarshalFromObject(s AutoKeyProxyIdProtocol) {
	o.Number = s.Number
	if s.Any != nil {
		var obj autoKeyProxyIdProtocolAnyXml_11_0_2
		obj.MarshalFromObject(*s.Any)
		o.Any = &obj
	}
	if s.Tcp != nil {
		var obj autoKeyProxyIdProtocolTcpXml_11_0_2
		obj.MarshalFromObject(*s.Tcp)
		o.Tcp = &obj
	}
	if s.Udp != nil {
		var obj autoKeyProxyIdProtocolUdpXml_11_0_2
		obj.MarshalFromObject(*s.Udp)
		o.Udp = &obj
	}
	o.Misc = s.Misc
}

func (o autoKeyProxyIdProtocolXml_11_0_2) UnmarshalToObject() (*AutoKeyProxyIdProtocol, error) {
	var anyVal *AutoKeyProxyIdProtocolAny
	if o.Any != nil {
		obj, err := o.Any.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		anyVal = obj
	}
	var tcpVal *AutoKeyProxyIdProtocolTcp
	if o.Tcp != nil {
		obj, err := o.Tcp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		tcpVal = obj
	}
	var udpVal *AutoKeyProxyIdProtocolUdp
	if o.Udp != nil {
		obj, err := o.Udp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		udpVal = obj
	}

	result := &AutoKeyProxyIdProtocol{
		Number: o.Number,
		Any:    anyVal,
		Tcp:    tcpVal,
		Udp:    udpVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdProtocolAnyXml_11_0_2) MarshalFromObject(s AutoKeyProxyIdProtocolAny) {
	o.Misc = s.Misc
}

func (o autoKeyProxyIdProtocolAnyXml_11_0_2) UnmarshalToObject() (*AutoKeyProxyIdProtocolAny, error) {

	result := &AutoKeyProxyIdProtocolAny{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdProtocolTcpXml_11_0_2) MarshalFromObject(s AutoKeyProxyIdProtocolTcp) {
	o.LocalPort = s.LocalPort
	o.RemotePort = s.RemotePort
	o.Misc = s.Misc
}

func (o autoKeyProxyIdProtocolTcpXml_11_0_2) UnmarshalToObject() (*AutoKeyProxyIdProtocolTcp, error) {

	result := &AutoKeyProxyIdProtocolTcp{
		LocalPort:  o.LocalPort,
		RemotePort: o.RemotePort,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdProtocolUdpXml_11_0_2) MarshalFromObject(s AutoKeyProxyIdProtocolUdp) {
	o.LocalPort = s.LocalPort
	o.RemotePort = s.RemotePort
	o.Misc = s.Misc
}

func (o autoKeyProxyIdProtocolUdpXml_11_0_2) UnmarshalToObject() (*AutoKeyProxyIdProtocolUdp, error) {

	result := &AutoKeyProxyIdProtocolUdp{
		LocalPort:  o.LocalPort,
		RemotePort: o.RemotePort,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdV6Xml_11_0_2) MarshalFromObject(s AutoKeyProxyIdV6) {
	o.Name = s.Name
	o.Local = s.Local
	o.Remote = s.Remote
	if s.Protocol != nil {
		var obj autoKeyProxyIdV6ProtocolXml_11_0_2
		obj.MarshalFromObject(*s.Protocol)
		o.Protocol = &obj
	}
	o.Misc = s.Misc
}

func (o autoKeyProxyIdV6Xml_11_0_2) UnmarshalToObject() (*AutoKeyProxyIdV6, error) {
	var protocolVal *AutoKeyProxyIdV6Protocol
	if o.Protocol != nil {
		obj, err := o.Protocol.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		protocolVal = obj
	}

	result := &AutoKeyProxyIdV6{
		Name:     o.Name,
		Local:    o.Local,
		Remote:   o.Remote,
		Protocol: protocolVal,
		Misc:     o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdV6ProtocolXml_11_0_2) MarshalFromObject(s AutoKeyProxyIdV6Protocol) {
	o.Number = s.Number
	if s.Any != nil {
		var obj autoKeyProxyIdV6ProtocolAnyXml_11_0_2
		obj.MarshalFromObject(*s.Any)
		o.Any = &obj
	}
	if s.Tcp != nil {
		var obj autoKeyProxyIdV6ProtocolTcpXml_11_0_2
		obj.MarshalFromObject(*s.Tcp)
		o.Tcp = &obj
	}
	if s.Udp != nil {
		var obj autoKeyProxyIdV6ProtocolUdpXml_11_0_2
		obj.MarshalFromObject(*s.Udp)
		o.Udp = &obj
	}
	o.Misc = s.Misc
}

func (o autoKeyProxyIdV6ProtocolXml_11_0_2) UnmarshalToObject() (*AutoKeyProxyIdV6Protocol, error) {
	var anyVal *AutoKeyProxyIdV6ProtocolAny
	if o.Any != nil {
		obj, err := o.Any.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		anyVal = obj
	}
	var tcpVal *AutoKeyProxyIdV6ProtocolTcp
	if o.Tcp != nil {
		obj, err := o.Tcp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		tcpVal = obj
	}
	var udpVal *AutoKeyProxyIdV6ProtocolUdp
	if o.Udp != nil {
		obj, err := o.Udp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		udpVal = obj
	}

	result := &AutoKeyProxyIdV6Protocol{
		Number: o.Number,
		Any:    anyVal,
		Tcp:    tcpVal,
		Udp:    udpVal,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdV6ProtocolAnyXml_11_0_2) MarshalFromObject(s AutoKeyProxyIdV6ProtocolAny) {
	o.Misc = s.Misc
}

func (o autoKeyProxyIdV6ProtocolAnyXml_11_0_2) UnmarshalToObject() (*AutoKeyProxyIdV6ProtocolAny, error) {

	result := &AutoKeyProxyIdV6ProtocolAny{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdV6ProtocolTcpXml_11_0_2) MarshalFromObject(s AutoKeyProxyIdV6ProtocolTcp) {
	o.LocalPort = s.LocalPort
	o.RemotePort = s.RemotePort
	o.Misc = s.Misc
}

func (o autoKeyProxyIdV6ProtocolTcpXml_11_0_2) UnmarshalToObject() (*AutoKeyProxyIdV6ProtocolTcp, error) {

	result := &AutoKeyProxyIdV6ProtocolTcp{
		LocalPort:  o.LocalPort,
		RemotePort: o.RemotePort,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *autoKeyProxyIdV6ProtocolUdpXml_11_0_2) MarshalFromObject(s AutoKeyProxyIdV6ProtocolUdp) {
	o.LocalPort = s.LocalPort
	o.RemotePort = s.RemotePort
	o.Misc = s.Misc
}

func (o autoKeyProxyIdV6ProtocolUdpXml_11_0_2) UnmarshalToObject() (*AutoKeyProxyIdV6ProtocolUdp, error) {

	result := &AutoKeyProxyIdV6ProtocolUdp{
		LocalPort:  o.LocalPort,
		RemotePort: o.RemotePort,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *globalProtectSatelliteXml_11_0_2) MarshalFromObject(s GlobalProtectSatellite) {
	if s.ExternalCa != nil {
		var obj globalProtectSatelliteExternalCaXml_11_0_2
		obj.MarshalFromObject(*s.ExternalCa)
		o.ExternalCa = &obj
	}
	o.Ipv6Preferred = util.YesNo(s.Ipv6Preferred, nil)
	if s.LocalAddress != nil {
		var obj globalProtectSatelliteLocalAddressXml_11_0_2
		obj.MarshalFromObject(*s.LocalAddress)
		o.LocalAddress = &obj
	}
	o.PortalAddress = s.PortalAddress
	if s.PublishConnectedRoutes != nil {
		var obj globalProtectSatellitePublishConnectedRoutesXml_11_0_2
		obj.MarshalFromObject(*s.PublishConnectedRoutes)
		o.PublishConnectedRoutes = &obj
	}
	if s.PublishRoutes != nil {
		o.PublishRoutes = util.StrToMem(s.PublishRoutes)
	}
	o.Misc = s.Misc
}

func (o globalProtectSatelliteXml_11_0_2) UnmarshalToObject() (*GlobalProtectSatellite, error) {
	var externalCaVal *GlobalProtectSatelliteExternalCa
	if o.ExternalCa != nil {
		obj, err := o.ExternalCa.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		externalCaVal = obj
	}
	var localAddressVal *GlobalProtectSatelliteLocalAddress
	if o.LocalAddress != nil {
		obj, err := o.LocalAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localAddressVal = obj
	}
	var publishConnectedRoutesVal *GlobalProtectSatellitePublishConnectedRoutes
	if o.PublishConnectedRoutes != nil {
		obj, err := o.PublishConnectedRoutes.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		publishConnectedRoutesVal = obj
	}
	var publishRoutesVal []string
	if o.PublishRoutes != nil {
		publishRoutesVal = util.MemToStr(o.PublishRoutes)
	}

	result := &GlobalProtectSatellite{
		ExternalCa:             externalCaVal,
		Ipv6Preferred:          util.AsBool(o.Ipv6Preferred, nil),
		LocalAddress:           localAddressVal,
		PortalAddress:          o.PortalAddress,
		PublishConnectedRoutes: publishConnectedRoutesVal,
		PublishRoutes:          publishRoutesVal,
		Misc:                   o.Misc,
	}
	return result, nil
}
func (o *globalProtectSatelliteExternalCaXml_11_0_2) MarshalFromObject(s GlobalProtectSatelliteExternalCa) {
	o.CertificateProfile = s.CertificateProfile
	o.LocalCertificate = s.LocalCertificate
	o.Misc = s.Misc
}

func (o globalProtectSatelliteExternalCaXml_11_0_2) UnmarshalToObject() (*GlobalProtectSatelliteExternalCa, error) {

	result := &GlobalProtectSatelliteExternalCa{
		CertificateProfile: o.CertificateProfile,
		LocalCertificate:   o.LocalCertificate,
		Misc:               o.Misc,
	}
	return result, nil
}
func (o *globalProtectSatelliteLocalAddressXml_11_0_2) MarshalFromObject(s GlobalProtectSatelliteLocalAddress) {
	o.Interface = s.Interface
	if s.FloatingIp != nil {
		var obj globalProtectSatelliteLocalAddressFloatingIpXml_11_0_2
		obj.MarshalFromObject(*s.FloatingIp)
		o.FloatingIp = &obj
	}
	if s.Ip != nil {
		var obj globalProtectSatelliteLocalAddressIpXml_11_0_2
		obj.MarshalFromObject(*s.Ip)
		o.Ip = &obj
	}
	o.Misc = s.Misc
}

func (o globalProtectSatelliteLocalAddressXml_11_0_2) UnmarshalToObject() (*GlobalProtectSatelliteLocalAddress, error) {
	var floatingIpVal *GlobalProtectSatelliteLocalAddressFloatingIp
	if o.FloatingIp != nil {
		obj, err := o.FloatingIp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		floatingIpVal = obj
	}
	var ipVal *GlobalProtectSatelliteLocalAddressIp
	if o.Ip != nil {
		obj, err := o.Ip.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ipVal = obj
	}

	result := &GlobalProtectSatelliteLocalAddress{
		Interface:  o.Interface,
		FloatingIp: floatingIpVal,
		Ip:         ipVal,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *globalProtectSatelliteLocalAddressFloatingIpXml_11_0_2) MarshalFromObject(s GlobalProtectSatelliteLocalAddressFloatingIp) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
}

func (o globalProtectSatelliteLocalAddressFloatingIpXml_11_0_2) UnmarshalToObject() (*GlobalProtectSatelliteLocalAddressFloatingIp, error) {

	result := &GlobalProtectSatelliteLocalAddressFloatingIp{
		Ipv4: o.Ipv4,
		Ipv6: o.Ipv6,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *globalProtectSatelliteLocalAddressIpXml_11_0_2) MarshalFromObject(s GlobalProtectSatelliteLocalAddressIp) {
	o.Ipv4 = s.Ipv4
	o.Ipv6 = s.Ipv6
	o.Misc = s.Misc
}

func (o globalProtectSatelliteLocalAddressIpXml_11_0_2) UnmarshalToObject() (*GlobalProtectSatelliteLocalAddressIp, error) {

	result := &GlobalProtectSatelliteLocalAddressIp{
		Ipv4: o.Ipv4,
		Ipv6: o.Ipv6,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *globalProtectSatellitePublishConnectedRoutesXml_11_0_2) MarshalFromObject(s GlobalProtectSatellitePublishConnectedRoutes) {
	o.Enable = util.YesNo(s.Enable, nil)
	o.Misc = s.Misc
}

func (o globalProtectSatellitePublishConnectedRoutesXml_11_0_2) UnmarshalToObject() (*GlobalProtectSatellitePublishConnectedRoutes, error) {

	result := &GlobalProtectSatellitePublishConnectedRoutes{
		Enable: util.AsBool(o.Enable, nil),
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *manualKeyXml_11_0_2) MarshalFromObject(s ManualKey) {
	if s.LocalAddress != nil {
		var obj manualKeyLocalAddressXml_11_0_2
		obj.MarshalFromObject(*s.LocalAddress)
		o.LocalAddress = &obj
	}
	o.LocalSpi = s.LocalSpi
	if s.PeerAddress != nil {
		var obj manualKeyPeerAddressXml_11_0_2
		obj.MarshalFromObject(*s.PeerAddress)
		o.PeerAddress = &obj
	}
	o.RemoteSpi = s.RemoteSpi
	if s.Ah != nil {
		var obj manualKeyAhXml_11_0_2
		obj.MarshalFromObject(*s.Ah)
		o.Ah = &obj
	}
	if s.Esp != nil {
		var obj manualKeyEspXml_11_0_2
		obj.MarshalFromObject(*s.Esp)
		o.Esp = &obj
	}
	o.Misc = s.Misc
}

func (o manualKeyXml_11_0_2) UnmarshalToObject() (*ManualKey, error) {
	var localAddressVal *ManualKeyLocalAddress
	if o.LocalAddress != nil {
		obj, err := o.LocalAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		localAddressVal = obj
	}
	var peerAddressVal *ManualKeyPeerAddress
	if o.PeerAddress != nil {
		obj, err := o.PeerAddress.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		peerAddressVal = obj
	}
	var ahVal *ManualKeyAh
	if o.Ah != nil {
		obj, err := o.Ah.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		ahVal = obj
	}
	var espVal *ManualKeyEsp
	if o.Esp != nil {
		obj, err := o.Esp.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		espVal = obj
	}

	result := &ManualKey{
		LocalAddress: localAddressVal,
		LocalSpi:     o.LocalSpi,
		PeerAddress:  peerAddressVal,
		RemoteSpi:    o.RemoteSpi,
		Ah:           ahVal,
		Esp:          espVal,
		Misc:         o.Misc,
	}
	return result, nil
}
func (o *manualKeyLocalAddressXml_11_0_2) MarshalFromObject(s ManualKeyLocalAddress) {
	o.Interface = s.Interface
	o.FloatingIp = s.FloatingIp
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o manualKeyLocalAddressXml_11_0_2) UnmarshalToObject() (*ManualKeyLocalAddress, error) {

	result := &ManualKeyLocalAddress{
		Interface:  o.Interface,
		FloatingIp: o.FloatingIp,
		Ip:         o.Ip,
		Misc:       o.Misc,
	}
	return result, nil
}
func (o *manualKeyPeerAddressXml_11_0_2) MarshalFromObject(s ManualKeyPeerAddress) {
	o.Ip = s.Ip
	o.Misc = s.Misc
}

func (o manualKeyPeerAddressXml_11_0_2) UnmarshalToObject() (*ManualKeyPeerAddress, error) {

	result := &ManualKeyPeerAddress{
		Ip:   o.Ip,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyAhXml_11_0_2) MarshalFromObject(s ManualKeyAh) {
	if s.Md5 != nil {
		var obj manualKeyAhMd5Xml_11_0_2
		obj.MarshalFromObject(*s.Md5)
		o.Md5 = &obj
	}
	if s.Sha1 != nil {
		var obj manualKeyAhSha1Xml_11_0_2
		obj.MarshalFromObject(*s.Sha1)
		o.Sha1 = &obj
	}
	if s.Sha256 != nil {
		var obj manualKeyAhSha256Xml_11_0_2
		obj.MarshalFromObject(*s.Sha256)
		o.Sha256 = &obj
	}
	if s.Sha384 != nil {
		var obj manualKeyAhSha384Xml_11_0_2
		obj.MarshalFromObject(*s.Sha384)
		o.Sha384 = &obj
	}
	if s.Sha512 != nil {
		var obj manualKeyAhSha512Xml_11_0_2
		obj.MarshalFromObject(*s.Sha512)
		o.Sha512 = &obj
	}
	o.Misc = s.Misc
}

func (o manualKeyAhXml_11_0_2) UnmarshalToObject() (*ManualKeyAh, error) {
	var md5Val *ManualKeyAhMd5
	if o.Md5 != nil {
		obj, err := o.Md5.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		md5Val = obj
	}
	var sha1Val *ManualKeyAhSha1
	if o.Sha1 != nil {
		obj, err := o.Sha1.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha1Val = obj
	}
	var sha256Val *ManualKeyAhSha256
	if o.Sha256 != nil {
		obj, err := o.Sha256.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha256Val = obj
	}
	var sha384Val *ManualKeyAhSha384
	if o.Sha384 != nil {
		obj, err := o.Sha384.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha384Val = obj
	}
	var sha512Val *ManualKeyAhSha512
	if o.Sha512 != nil {
		obj, err := o.Sha512.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha512Val = obj
	}

	result := &ManualKeyAh{
		Md5:    md5Val,
		Sha1:   sha1Val,
		Sha256: sha256Val,
		Sha384: sha384Val,
		Sha512: sha512Val,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *manualKeyAhMd5Xml_11_0_2) MarshalFromObject(s ManualKeyAhMd5) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyAhMd5Xml_11_0_2) UnmarshalToObject() (*ManualKeyAhMd5, error) {

	result := &ManualKeyAhMd5{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyAhSha1Xml_11_0_2) MarshalFromObject(s ManualKeyAhSha1) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyAhSha1Xml_11_0_2) UnmarshalToObject() (*ManualKeyAhSha1, error) {

	result := &ManualKeyAhSha1{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyAhSha256Xml_11_0_2) MarshalFromObject(s ManualKeyAhSha256) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyAhSha256Xml_11_0_2) UnmarshalToObject() (*ManualKeyAhSha256, error) {

	result := &ManualKeyAhSha256{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyAhSha384Xml_11_0_2) MarshalFromObject(s ManualKeyAhSha384) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyAhSha384Xml_11_0_2) UnmarshalToObject() (*ManualKeyAhSha384, error) {

	result := &ManualKeyAhSha384{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyAhSha512Xml_11_0_2) MarshalFromObject(s ManualKeyAhSha512) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyAhSha512Xml_11_0_2) UnmarshalToObject() (*ManualKeyAhSha512, error) {

	result := &ManualKeyAhSha512{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspXml_11_0_2) MarshalFromObject(s ManualKeyEsp) {
	if s.Authentication != nil {
		var obj manualKeyEspAuthenticationXml_11_0_2
		obj.MarshalFromObject(*s.Authentication)
		o.Authentication = &obj
	}
	if s.Encryption != nil {
		var obj manualKeyEspEncryptionXml_11_0_2
		obj.MarshalFromObject(*s.Encryption)
		o.Encryption = &obj
	}
	o.Misc = s.Misc
}

func (o manualKeyEspXml_11_0_2) UnmarshalToObject() (*ManualKeyEsp, error) {
	var authenticationVal *ManualKeyEspAuthentication
	if o.Authentication != nil {
		obj, err := o.Authentication.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		authenticationVal = obj
	}
	var encryptionVal *ManualKeyEspEncryption
	if o.Encryption != nil {
		obj, err := o.Encryption.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		encryptionVal = obj
	}

	result := &ManualKeyEsp{
		Authentication: authenticationVal,
		Encryption:     encryptionVal,
		Misc:           o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspAuthenticationXml_11_0_2) MarshalFromObject(s ManualKeyEspAuthentication) {
	if s.Md5 != nil {
		var obj manualKeyEspAuthenticationMd5Xml_11_0_2
		obj.MarshalFromObject(*s.Md5)
		o.Md5 = &obj
	}
	if s.None != nil {
		var obj manualKeyEspAuthenticationNoneXml_11_0_2
		obj.MarshalFromObject(*s.None)
		o.None = &obj
	}
	if s.Sha1 != nil {
		var obj manualKeyEspAuthenticationSha1Xml_11_0_2
		obj.MarshalFromObject(*s.Sha1)
		o.Sha1 = &obj
	}
	if s.Sha256 != nil {
		var obj manualKeyEspAuthenticationSha256Xml_11_0_2
		obj.MarshalFromObject(*s.Sha256)
		o.Sha256 = &obj
	}
	if s.Sha384 != nil {
		var obj manualKeyEspAuthenticationSha384Xml_11_0_2
		obj.MarshalFromObject(*s.Sha384)
		o.Sha384 = &obj
	}
	if s.Sha512 != nil {
		var obj manualKeyEspAuthenticationSha512Xml_11_0_2
		obj.MarshalFromObject(*s.Sha512)
		o.Sha512 = &obj
	}
	o.Misc = s.Misc
}

func (o manualKeyEspAuthenticationXml_11_0_2) UnmarshalToObject() (*ManualKeyEspAuthentication, error) {
	var md5Val *ManualKeyEspAuthenticationMd5
	if o.Md5 != nil {
		obj, err := o.Md5.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		md5Val = obj
	}
	var noneVal *ManualKeyEspAuthenticationNone
	if o.None != nil {
		obj, err := o.None.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		noneVal = obj
	}
	var sha1Val *ManualKeyEspAuthenticationSha1
	if o.Sha1 != nil {
		obj, err := o.Sha1.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha1Val = obj
	}
	var sha256Val *ManualKeyEspAuthenticationSha256
	if o.Sha256 != nil {
		obj, err := o.Sha256.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha256Val = obj
	}
	var sha384Val *ManualKeyEspAuthenticationSha384
	if o.Sha384 != nil {
		obj, err := o.Sha384.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha384Val = obj
	}
	var sha512Val *ManualKeyEspAuthenticationSha512
	if o.Sha512 != nil {
		obj, err := o.Sha512.UnmarshalToObject()
		if err != nil {
			return nil, err
		}
		sha512Val = obj
	}

	result := &ManualKeyEspAuthentication{
		Md5:    md5Val,
		None:   noneVal,
		Sha1:   sha1Val,
		Sha256: sha256Val,
		Sha384: sha384Val,
		Sha512: sha512Val,
		Misc:   o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspAuthenticationMd5Xml_11_0_2) MarshalFromObject(s ManualKeyEspAuthenticationMd5) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyEspAuthenticationMd5Xml_11_0_2) UnmarshalToObject() (*ManualKeyEspAuthenticationMd5, error) {

	result := &ManualKeyEspAuthenticationMd5{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspAuthenticationNoneXml_11_0_2) MarshalFromObject(s ManualKeyEspAuthenticationNone) {
	o.Misc = s.Misc
}

func (o manualKeyEspAuthenticationNoneXml_11_0_2) UnmarshalToObject() (*ManualKeyEspAuthenticationNone, error) {

	result := &ManualKeyEspAuthenticationNone{
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspAuthenticationSha1Xml_11_0_2) MarshalFromObject(s ManualKeyEspAuthenticationSha1) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyEspAuthenticationSha1Xml_11_0_2) UnmarshalToObject() (*ManualKeyEspAuthenticationSha1, error) {

	result := &ManualKeyEspAuthenticationSha1{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspAuthenticationSha256Xml_11_0_2) MarshalFromObject(s ManualKeyEspAuthenticationSha256) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyEspAuthenticationSha256Xml_11_0_2) UnmarshalToObject() (*ManualKeyEspAuthenticationSha256, error) {

	result := &ManualKeyEspAuthenticationSha256{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspAuthenticationSha384Xml_11_0_2) MarshalFromObject(s ManualKeyEspAuthenticationSha384) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyEspAuthenticationSha384Xml_11_0_2) UnmarshalToObject() (*ManualKeyEspAuthenticationSha384, error) {

	result := &ManualKeyEspAuthenticationSha384{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspAuthenticationSha512Xml_11_0_2) MarshalFromObject(s ManualKeyEspAuthenticationSha512) {
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyEspAuthenticationSha512Xml_11_0_2) UnmarshalToObject() (*ManualKeyEspAuthenticationSha512, error) {

	result := &ManualKeyEspAuthenticationSha512{
		Key:  o.Key,
		Misc: o.Misc,
	}
	return result, nil
}
func (o *manualKeyEspEncryptionXml_11_0_2) MarshalFromObject(s ManualKeyEspEncryption) {
	o.Algorithm = s.Algorithm
	o.Key = s.Key
	o.Misc = s.Misc
}

func (o manualKeyEspEncryptionXml_11_0_2) UnmarshalToObject() (*ManualKeyEspEncryption, error) {

	result := &ManualKeyEspEncryption{
		Algorithm: o.Algorithm,
		Key:       o.Key,
		Misc:      o.Misc,
	}
	return result, nil
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "anti_replay" || v == "AntiReplay" {
		return e.AntiReplay, nil
	}
	if v == "anti_replay_window" || v == "AntiReplayWindow" {
		return e.AntiReplayWindow, nil
	}
	if v == "comment" || v == "Comment" {
		return e.Comment, nil
	}
	if v == "copy_flow_label" || v == "CopyFlowLabel" {
		return e.CopyFlowLabel, nil
	}
	if v == "copy_tos" || v == "CopyTos" {
		return e.CopyTos, nil
	}
	if v == "disabled" || v == "Disabled" {
		return e.Disabled, nil
	}
	if v == "enable_gre_encapsulation" || v == "EnableGreEncapsulation" {
		return e.EnableGreEncapsulation, nil
	}
	if v == "ipsec_mode" || v == "IpsecMode" {
		return e.IpsecMode, nil
	}
	if v == "ipv6" || v == "Ipv6" {
		return e.Ipv6, nil
	}
	if v == "tunnel_interface" || v == "TunnelInterface" {
		return e.TunnelInterface, nil
	}
	if v == "tunnel_monitor" || v == "TunnelMonitor" {
		return e.TunnelMonitor, nil
	}
	if v == "auto_key" || v == "AutoKey" {
		return e.AutoKey, nil
	}
	if v == "global_protect_satellite" || v == "GlobalProtectSatellite" {
		return e.GlobalProtectSatellite, nil
	}
	if v == "manual_key" || v == "ManualKey" {
		return e.ManualKey, nil
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
	if !util.BoolsMatch(o.AntiReplay, other.AntiReplay) {
		return false
	}
	if !util.StringsMatch(o.AntiReplayWindow, other.AntiReplayWindow) {
		return false
	}
	if !util.StringsMatch(o.Comment, other.Comment) {
		return false
	}
	if !util.BoolsMatch(o.CopyFlowLabel, other.CopyFlowLabel) {
		return false
	}
	if !util.BoolsMatch(o.CopyTos, other.CopyTos) {
		return false
	}
	if !util.BoolsMatch(o.Disabled, other.Disabled) {
		return false
	}
	if !util.BoolsMatch(o.EnableGreEncapsulation, other.EnableGreEncapsulation) {
		return false
	}
	if !util.StringsMatch(o.IpsecMode, other.IpsecMode) {
		return false
	}
	if !util.BoolsMatch(o.Ipv6, other.Ipv6) {
		return false
	}
	if !util.StringsMatch(o.TunnelInterface, other.TunnelInterface) {
		return false
	}
	if !o.TunnelMonitor.matches(other.TunnelMonitor) {
		return false
	}
	if !o.AutoKey.matches(other.AutoKey) {
		return false
	}
	if !o.GlobalProtectSatellite.matches(other.GlobalProtectSatellite) {
		return false
	}
	if !o.ManualKey.matches(other.ManualKey) {
		return false
	}

	return true
}

func (o *TunnelMonitor) matches(other *TunnelMonitor) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.DestinationIp, other.DestinationIp) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}
	if !util.StringsMatch(o.ProxyId, other.ProxyId) {
		return false
	}
	if !util.StringsMatch(o.TunnelMonitorProfile, other.TunnelMonitorProfile) {
		return false
	}

	return true
}

func (o *AutoKey) matches(other *AutoKey) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if len(o.IkeGateway) != len(other.IkeGateway) {
		return false
	}
	for idx := range o.IkeGateway {
		if !o.IkeGateway[idx].matches(&other.IkeGateway[idx]) {
			return false
		}
	}
	if !util.StringsMatch(o.IpsecCryptoProfile, other.IpsecCryptoProfile) {
		return false
	}
	if len(o.ProxyId) != len(other.ProxyId) {
		return false
	}
	for idx := range o.ProxyId {
		if !o.ProxyId[idx].matches(&other.ProxyId[idx]) {
			return false
		}
	}
	if len(o.ProxyIdV6) != len(other.ProxyIdV6) {
		return false
	}
	for idx := range o.ProxyIdV6 {
		if !o.ProxyIdV6[idx].matches(&other.ProxyIdV6[idx]) {
			return false
		}
	}

	return true
}

func (o *AutoKeyIkeGateway) matches(other *AutoKeyIkeGateway) bool {
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

func (o *AutoKeyProxyId) matches(other *AutoKeyProxyId) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Local, other.Local) {
		return false
	}
	if !util.StringsMatch(o.Remote, other.Remote) {
		return false
	}
	if !o.Protocol.matches(other.Protocol) {
		return false
	}

	return true
}

func (o *AutoKeyProxyIdProtocol) matches(other *AutoKeyProxyIdProtocol) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Number, other.Number) {
		return false
	}
	if !o.Any.matches(other.Any) {
		return false
	}
	if !o.Tcp.matches(other.Tcp) {
		return false
	}
	if !o.Udp.matches(other.Udp) {
		return false
	}

	return true
}

func (o *AutoKeyProxyIdProtocolAny) matches(other *AutoKeyProxyIdProtocolAny) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *AutoKeyProxyIdProtocolTcp) matches(other *AutoKeyProxyIdProtocolTcp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.LocalPort, other.LocalPort) {
		return false
	}
	if !util.Ints64Match(o.RemotePort, other.RemotePort) {
		return false
	}

	return true
}

func (o *AutoKeyProxyIdProtocolUdp) matches(other *AutoKeyProxyIdProtocolUdp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.LocalPort, other.LocalPort) {
		return false
	}
	if !util.Ints64Match(o.RemotePort, other.RemotePort) {
		return false
	}

	return true
}

func (o *AutoKeyProxyIdV6) matches(other *AutoKeyProxyIdV6) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if o.Name != other.Name {
		return false
	}
	if !util.StringsMatch(o.Local, other.Local) {
		return false
	}
	if !util.StringsMatch(o.Remote, other.Remote) {
		return false
	}
	if !o.Protocol.matches(other.Protocol) {
		return false
	}

	return true
}

func (o *AutoKeyProxyIdV6Protocol) matches(other *AutoKeyProxyIdV6Protocol) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.Number, other.Number) {
		return false
	}
	if !o.Any.matches(other.Any) {
		return false
	}
	if !o.Tcp.matches(other.Tcp) {
		return false
	}
	if !o.Udp.matches(other.Udp) {
		return false
	}

	return true
}

func (o *AutoKeyProxyIdV6ProtocolAny) matches(other *AutoKeyProxyIdV6ProtocolAny) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *AutoKeyProxyIdV6ProtocolTcp) matches(other *AutoKeyProxyIdV6ProtocolTcp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.LocalPort, other.LocalPort) {
		return false
	}
	if !util.Ints64Match(o.RemotePort, other.RemotePort) {
		return false
	}

	return true
}

func (o *AutoKeyProxyIdV6ProtocolUdp) matches(other *AutoKeyProxyIdV6ProtocolUdp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.Ints64Match(o.LocalPort, other.LocalPort) {
		return false
	}
	if !util.Ints64Match(o.RemotePort, other.RemotePort) {
		return false
	}

	return true
}

func (o *GlobalProtectSatellite) matches(other *GlobalProtectSatellite) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.ExternalCa.matches(other.ExternalCa) {
		return false
	}
	if !util.BoolsMatch(o.Ipv6Preferred, other.Ipv6Preferred) {
		return false
	}
	if !o.LocalAddress.matches(other.LocalAddress) {
		return false
	}
	if !util.StringsMatch(o.PortalAddress, other.PortalAddress) {
		return false
	}
	if !o.PublishConnectedRoutes.matches(other.PublishConnectedRoutes) {
		return false
	}
	if !util.OrderedListsMatch[string](o.PublishRoutes, other.PublishRoutes) {
		return false
	}

	return true
}

func (o *GlobalProtectSatelliteExternalCa) matches(other *GlobalProtectSatelliteExternalCa) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.CertificateProfile, other.CertificateProfile) {
		return false
	}
	if !util.StringsMatch(o.LocalCertificate, other.LocalCertificate) {
		return false
	}

	return true
}

func (o *GlobalProtectSatelliteLocalAddress) matches(other *GlobalProtectSatelliteLocalAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !o.FloatingIp.matches(other.FloatingIp) {
		return false
	}
	if !o.Ip.matches(other.Ip) {
		return false
	}

	return true
}

func (o *GlobalProtectSatelliteLocalAddressFloatingIp) matches(other *GlobalProtectSatelliteLocalAddressFloatingIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ipv4, other.Ipv4) {
		return false
	}
	if !util.StringsMatch(o.Ipv6, other.Ipv6) {
		return false
	}

	return true
}

func (o *GlobalProtectSatelliteLocalAddressIp) matches(other *GlobalProtectSatelliteLocalAddressIp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ipv4, other.Ipv4) {
		return false
	}
	if !util.StringsMatch(o.Ipv6, other.Ipv6) {
		return false
	}

	return true
}

func (o *GlobalProtectSatellitePublishConnectedRoutes) matches(other *GlobalProtectSatellitePublishConnectedRoutes) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.BoolsMatch(o.Enable, other.Enable) {
		return false
	}

	return true
}

func (o *ManualKey) matches(other *ManualKey) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.LocalAddress.matches(other.LocalAddress) {
		return false
	}
	if !util.StringsMatch(o.LocalSpi, other.LocalSpi) {
		return false
	}
	if !o.PeerAddress.matches(other.PeerAddress) {
		return false
	}
	if !util.StringsMatch(o.RemoteSpi, other.RemoteSpi) {
		return false
	}
	if !o.Ah.matches(other.Ah) {
		return false
	}
	if !o.Esp.matches(other.Esp) {
		return false
	}

	return true
}

func (o *ManualKeyLocalAddress) matches(other *ManualKeyLocalAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Interface, other.Interface) {
		return false
	}
	if !util.StringsMatch(o.FloatingIp, other.FloatingIp) {
		return false
	}
	if !util.StringsMatch(o.Ip, other.Ip) {
		return false
	}

	return true
}

func (o *ManualKeyPeerAddress) matches(other *ManualKeyPeerAddress) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Ip, other.Ip) {
		return false
	}

	return true
}

func (o *ManualKeyAh) matches(other *ManualKeyAh) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Md5.matches(other.Md5) {
		return false
	}
	if !o.Sha1.matches(other.Sha1) {
		return false
	}
	if !o.Sha256.matches(other.Sha256) {
		return false
	}
	if !o.Sha384.matches(other.Sha384) {
		return false
	}
	if !o.Sha512.matches(other.Sha512) {
		return false
	}

	return true
}

func (o *ManualKeyAhMd5) matches(other *ManualKeyAhMd5) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ManualKeyAhSha1) matches(other *ManualKeyAhSha1) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ManualKeyAhSha256) matches(other *ManualKeyAhSha256) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ManualKeyAhSha384) matches(other *ManualKeyAhSha384) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ManualKeyAhSha512) matches(other *ManualKeyAhSha512) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ManualKeyEsp) matches(other *ManualKeyEsp) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Authentication.matches(other.Authentication) {
		return false
	}
	if !o.Encryption.matches(other.Encryption) {
		return false
	}

	return true
}

func (o *ManualKeyEspAuthentication) matches(other *ManualKeyEspAuthentication) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !o.Md5.matches(other.Md5) {
		return false
	}
	if !o.None.matches(other.None) {
		return false
	}
	if !o.Sha1.matches(other.Sha1) {
		return false
	}
	if !o.Sha256.matches(other.Sha256) {
		return false
	}
	if !o.Sha384.matches(other.Sha384) {
		return false
	}
	if !o.Sha512.matches(other.Sha512) {
		return false
	}

	return true
}

func (o *ManualKeyEspAuthenticationMd5) matches(other *ManualKeyEspAuthenticationMd5) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ManualKeyEspAuthenticationNone) matches(other *ManualKeyEspAuthenticationNone) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}

	return true
}

func (o *ManualKeyEspAuthenticationSha1) matches(other *ManualKeyEspAuthenticationSha1) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ManualKeyEspAuthenticationSha256) matches(other *ManualKeyEspAuthenticationSha256) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ManualKeyEspAuthenticationSha384) matches(other *ManualKeyEspAuthenticationSha384) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ManualKeyEspAuthenticationSha512) matches(other *ManualKeyEspAuthenticationSha512) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
		return false
	}

	return true
}

func (o *ManualKeyEspEncryption) matches(other *ManualKeyEspEncryption) bool {
	if o == nil && other == nil {
		return true
	}

	if (o == nil && other != nil) || (o != nil && other == nil) {
		return false
	}
	if !util.StringsMatch(o.Algorithm, other.Algorithm) {
		return false
	}
	if !util.StringsMatch(o.Key, other.Key) {
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
