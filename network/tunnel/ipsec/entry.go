package ipsec

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
	Suffix = []string{"network", "tunnel", "ipsec"}
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

	Misc map[string][]generic.Xml
}

type AutoKey struct {
	IkeGateway         []AutoKeyIkeGateway
	IpsecCryptoProfile *string
	ProxyId            []AutoKeyProxyId
	ProxyIdV6          []AutoKeyProxyIdV6
}
type AutoKeyIkeGateway struct {
	Name string
}
type AutoKeyProxyId struct {
	Local    *string
	Name     string
	Protocol *AutoKeyProxyIdProtocol
	Remote   *string
}
type AutoKeyProxyIdProtocol struct {
	Any    *AutoKeyProxyIdProtocolAny
	Number *int64
	Tcp    *AutoKeyProxyIdProtocolTcp
	Udp    *AutoKeyProxyIdProtocolUdp
}
type AutoKeyProxyIdProtocolAny struct {
}
type AutoKeyProxyIdProtocolTcp struct {
	LocalPort  *int64
	RemotePort *int64
}
type AutoKeyProxyIdProtocolUdp struct {
	LocalPort  *int64
	RemotePort *int64
}
type AutoKeyProxyIdV6 struct {
	Local    *string
	Name     string
	Protocol *AutoKeyProxyIdV6Protocol
	Remote   *string
}
type AutoKeyProxyIdV6Protocol struct {
	Any    *AutoKeyProxyIdV6ProtocolAny
	Number *int64
	Tcp    *AutoKeyProxyIdV6ProtocolTcp
	Udp    *AutoKeyProxyIdV6ProtocolUdp
}
type AutoKeyProxyIdV6ProtocolAny struct {
}
type AutoKeyProxyIdV6ProtocolTcp struct {
	LocalPort  *int64
	RemotePort *int64
}
type AutoKeyProxyIdV6ProtocolUdp struct {
	LocalPort  *int64
	RemotePort *int64
}
type GlobalProtectSatellite struct {
	ExternalCa             *GlobalProtectSatelliteExternalCa
	Ipv6Preferred          *bool
	LocalAddress           *GlobalProtectSatelliteLocalAddress
	PortalAddress          *string
	PublishConnectedRoutes *GlobalProtectSatellitePublishConnectedRoutes
	PublishRoutes          []string
}
type GlobalProtectSatelliteExternalCa struct {
	CertificateProfile *string
	LocalCertificate   *string
}
type GlobalProtectSatelliteLocalAddress struct {
	Interface  *string
	FloatingIp *GlobalProtectSatelliteLocalAddressFloatingIp
	Ip         *GlobalProtectSatelliteLocalAddressIp
}
type GlobalProtectSatelliteLocalAddressFloatingIp struct {
	Ipv4 *string
	Ipv6 *string
}
type GlobalProtectSatelliteLocalAddressIp struct {
	Ipv4 *string
	Ipv6 *string
}
type GlobalProtectSatellitePublishConnectedRoutes struct {
	Enable *bool
}
type ManualKey struct {
	LocalAddress *ManualKeyLocalAddress
	LocalSpi     *string
	PeerAddress  *ManualKeyPeerAddress
	RemoteSpi    *string
	Ah           *ManualKeyAh
	Esp          *ManualKeyEsp
}
type ManualKeyAh struct {
	Md5    *ManualKeyAhMd5
	Sha1   *ManualKeyAhSha1
	Sha256 *ManualKeyAhSha256
	Sha384 *ManualKeyAhSha384
	Sha512 *ManualKeyAhSha512
}
type ManualKeyAhMd5 struct {
	Key *string
}
type ManualKeyAhSha1 struct {
	Key *string
}
type ManualKeyAhSha256 struct {
	Key *string
}
type ManualKeyAhSha384 struct {
	Key *string
}
type ManualKeyAhSha512 struct {
	Key *string
}
type ManualKeyEsp struct {
	Authentication *ManualKeyEspAuthentication
	Encryption     *ManualKeyEspEncryption
}
type ManualKeyEspAuthentication struct {
	Md5    *ManualKeyEspAuthenticationMd5
	None   *ManualKeyEspAuthenticationNone
	Sha1   *ManualKeyEspAuthenticationSha1
	Sha256 *ManualKeyEspAuthenticationSha256
	Sha384 *ManualKeyEspAuthenticationSha384
	Sha512 *ManualKeyEspAuthenticationSha512
}
type ManualKeyEspAuthenticationMd5 struct {
	Key *string
}
type ManualKeyEspAuthenticationNone struct {
}
type ManualKeyEspAuthenticationSha1 struct {
	Key *string
}
type ManualKeyEspAuthenticationSha256 struct {
	Key *string
}
type ManualKeyEspAuthenticationSha384 struct {
	Key *string
}
type ManualKeyEspAuthenticationSha512 struct {
	Key *string
}
type ManualKeyEspEncryption struct {
	Algorithm *string
	Key       *string
}
type ManualKeyLocalAddress struct {
	Interface  *string
	FloatingIp *string
	Ip         *string
}
type ManualKeyPeerAddress struct {
	Ip *string
}
type TunnelMonitor struct {
	DestinationIp        *string
	Enable               *bool
	ProxyId              *string
	TunnelMonitorProfile *string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXmlContainer_11_0_2 struct {
	Answer []entryXml_11_0_2 `xml:"entry"`
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
	TunnelMonitor          *TunnelMonitorXml          `xml:"tunnel-monitor,omitempty"`
	AutoKey                *AutoKeyXml                `xml:"auto-key,omitempty"`
	GlobalProtectSatellite *GlobalProtectSatelliteXml `xml:"global-protect-satellite,omitempty"`
	ManualKey              *ManualKeyXml              `xml:"manual-key,omitempty"`

	Misc []generic.Xml `xml:",any"`
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
	TunnelMonitor          *TunnelMonitorXml_11_0_2          `xml:"tunnel-monitor,omitempty"`
	AutoKey                *AutoKeyXml_11_0_2                `xml:"auto-key,omitempty"`
	GlobalProtectSatellite *GlobalProtectSatelliteXml_11_0_2 `xml:"global-protect-satellite,omitempty"`
	ManualKey              *ManualKeyXml_11_0_2              `xml:"manual-key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyXml struct {
	IkeGateway         []AutoKeyIkeGatewayXml `xml:"ike-gateway>entry,omitempty"`
	IpsecCryptoProfile *string                `xml:"ipsec-crypto-profile,omitempty"`
	ProxyId            []AutoKeyProxyIdXml    `xml:"proxy-id>entry,omitempty"`
	ProxyIdV6          []AutoKeyProxyIdV6Xml  `xml:"proxy-id-v6>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyIkeGatewayXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdXml struct {
	Local    *string                    `xml:"local,omitempty"`
	XMLName  xml.Name                   `xml:"entry"`
	Name     string                     `xml:"name,attr"`
	Protocol *AutoKeyProxyIdProtocolXml `xml:"protocol,omitempty"`
	Remote   *string                    `xml:"remote,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdProtocolXml struct {
	Any    *AutoKeyProxyIdProtocolAnyXml `xml:"any,omitempty"`
	Number *int64                        `xml:"number,omitempty"`
	Tcp    *AutoKeyProxyIdProtocolTcpXml `xml:"tcp,omitempty"`
	Udp    *AutoKeyProxyIdProtocolUdpXml `xml:"udp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdProtocolAnyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdProtocolTcpXml struct {
	LocalPort  *int64 `xml:"local-port,omitempty"`
	RemotePort *int64 `xml:"remote-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdProtocolUdpXml struct {
	LocalPort  *int64 `xml:"local-port,omitempty"`
	RemotePort *int64 `xml:"remote-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdV6Xml struct {
	Local    *string                      `xml:"local,omitempty"`
	XMLName  xml.Name                     `xml:"entry"`
	Name     string                       `xml:"name,attr"`
	Protocol *AutoKeyProxyIdV6ProtocolXml `xml:"protocol,omitempty"`
	Remote   *string                      `xml:"remote,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdV6ProtocolXml struct {
	Any    *AutoKeyProxyIdV6ProtocolAnyXml `xml:"any,omitempty"`
	Number *int64                          `xml:"number,omitempty"`
	Tcp    *AutoKeyProxyIdV6ProtocolTcpXml `xml:"tcp,omitempty"`
	Udp    *AutoKeyProxyIdV6ProtocolUdpXml `xml:"udp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdV6ProtocolAnyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdV6ProtocolTcpXml struct {
	LocalPort  *int64 `xml:"local-port,omitempty"`
	RemotePort *int64 `xml:"remote-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdV6ProtocolUdpXml struct {
	LocalPort  *int64 `xml:"local-port,omitempty"`
	RemotePort *int64 `xml:"remote-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type GlobalProtectSatelliteXml struct {
	ExternalCa             *GlobalProtectSatelliteExternalCaXml             `xml:"external-ca,omitempty"`
	Ipv6Preferred          *string                                          `xml:"ipv6-preferred,omitempty"`
	LocalAddress           *GlobalProtectSatelliteLocalAddressXml           `xml:"local-address,omitempty"`
	PortalAddress          *string                                          `xml:"portal-address,omitempty"`
	PublishConnectedRoutes *GlobalProtectSatellitePublishConnectedRoutesXml `xml:"publish-connected-routes,omitempty"`
	PublishRoutes          *util.MemberType                                 `xml:"publish-routes,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type GlobalProtectSatelliteExternalCaXml struct {
	CertificateProfile *string `xml:"certificate-profile,omitempty"`
	LocalCertificate   *string `xml:"local-certificate,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type GlobalProtectSatelliteLocalAddressXml struct {
	Interface  *string                                          `xml:"interface,omitempty"`
	FloatingIp *GlobalProtectSatelliteLocalAddressFloatingIpXml `xml:"floating-ip,omitempty"`
	Ip         *GlobalProtectSatelliteLocalAddressIpXml         `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type GlobalProtectSatelliteLocalAddressFloatingIpXml struct {
	Ipv4 *string `xml:"ipv4,omitempty"`
	Ipv6 *string `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type GlobalProtectSatelliteLocalAddressIpXml struct {
	Ipv4 *string `xml:"ipv4,omitempty"`
	Ipv6 *string `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type GlobalProtectSatellitePublishConnectedRoutesXml struct {
	Enable *string `xml:"enable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyXml struct {
	LocalAddress *ManualKeyLocalAddressXml `xml:"local-address,omitempty"`
	LocalSpi     *string                   `xml:"local-spi,omitempty"`
	PeerAddress  *ManualKeyPeerAddressXml  `xml:"peer-address,omitempty"`
	RemoteSpi    *string                   `xml:"remote-spi,omitempty"`
	Ah           *ManualKeyAhXml           `xml:"ah,omitempty"`
	Esp          *ManualKeyEspXml          `xml:"esp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyAhXml struct {
	Md5    *ManualKeyAhMd5Xml    `xml:"md5,omitempty"`
	Sha1   *ManualKeyAhSha1Xml   `xml:"sha1,omitempty"`
	Sha256 *ManualKeyAhSha256Xml `xml:"sha256,omitempty"`
	Sha384 *ManualKeyAhSha384Xml `xml:"sha384,omitempty"`
	Sha512 *ManualKeyAhSha512Xml `xml:"sha512,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyAhMd5Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyAhSha1Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyAhSha256Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyAhSha384Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyAhSha512Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspXml struct {
	Authentication *ManualKeyEspAuthenticationXml `xml:"authentication,omitempty"`
	Encryption     *ManualKeyEspEncryptionXml     `xml:"encryption,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspAuthenticationXml struct {
	Md5    *ManualKeyEspAuthenticationMd5Xml    `xml:"md5,omitempty"`
	None   *ManualKeyEspAuthenticationNoneXml   `xml:"none,omitempty"`
	Sha1   *ManualKeyEspAuthenticationSha1Xml   `xml:"sha1,omitempty"`
	Sha256 *ManualKeyEspAuthenticationSha256Xml `xml:"sha256,omitempty"`
	Sha384 *ManualKeyEspAuthenticationSha384Xml `xml:"sha384,omitempty"`
	Sha512 *ManualKeyEspAuthenticationSha512Xml `xml:"sha512,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspAuthenticationMd5Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspAuthenticationNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspAuthenticationSha1Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspAuthenticationSha256Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspAuthenticationSha384Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspAuthenticationSha512Xml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspEncryptionXml struct {
	Algorithm *string `xml:"algorithm,omitempty"`
	Key       *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyLocalAddressXml struct {
	Interface  *string `xml:"interface,omitempty"`
	FloatingIp *string `xml:"floating-ip,omitempty"`
	Ip         *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyPeerAddressXml struct {
	Ip *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TunnelMonitorXml struct {
	DestinationIp        *string `xml:"destination-ip,omitempty"`
	Enable               *string `xml:"enable,omitempty"`
	ProxyId              *string `xml:"proxy-id,omitempty"`
	TunnelMonitorProfile *string `xml:"tunnel-monitor-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyXml_11_0_2 struct {
	IkeGateway         []AutoKeyIkeGatewayXml_11_0_2 `xml:"ike-gateway>entry,omitempty"`
	IpsecCryptoProfile *string                       `xml:"ipsec-crypto-profile,omitempty"`
	ProxyId            []AutoKeyProxyIdXml_11_0_2    `xml:"proxy-id>entry,omitempty"`
	ProxyIdV6          []AutoKeyProxyIdV6Xml_11_0_2  `xml:"proxy-id-v6>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyIkeGatewayXml_11_0_2 struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdXml_11_0_2 struct {
	Local    *string                           `xml:"local,omitempty"`
	XMLName  xml.Name                          `xml:"entry"`
	Name     string                            `xml:"name,attr"`
	Protocol *AutoKeyProxyIdProtocolXml_11_0_2 `xml:"protocol,omitempty"`
	Remote   *string                           `xml:"remote,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdProtocolXml_11_0_2 struct {
	Any    *AutoKeyProxyIdProtocolAnyXml_11_0_2 `xml:"any,omitempty"`
	Number *int64                               `xml:"number,omitempty"`
	Tcp    *AutoKeyProxyIdProtocolTcpXml_11_0_2 `xml:"tcp,omitempty"`
	Udp    *AutoKeyProxyIdProtocolUdpXml_11_0_2 `xml:"udp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdProtocolAnyXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdProtocolTcpXml_11_0_2 struct {
	LocalPort  *int64 `xml:"local-port,omitempty"`
	RemotePort *int64 `xml:"remote-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdProtocolUdpXml_11_0_2 struct {
	LocalPort  *int64 `xml:"local-port,omitempty"`
	RemotePort *int64 `xml:"remote-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdV6Xml_11_0_2 struct {
	Local    *string                             `xml:"local,omitempty"`
	XMLName  xml.Name                            `xml:"entry"`
	Name     string                              `xml:"name,attr"`
	Protocol *AutoKeyProxyIdV6ProtocolXml_11_0_2 `xml:"protocol,omitempty"`
	Remote   *string                             `xml:"remote,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdV6ProtocolXml_11_0_2 struct {
	Any    *AutoKeyProxyIdV6ProtocolAnyXml_11_0_2 `xml:"any,omitempty"`
	Number *int64                                 `xml:"number,omitempty"`
	Tcp    *AutoKeyProxyIdV6ProtocolTcpXml_11_0_2 `xml:"tcp,omitempty"`
	Udp    *AutoKeyProxyIdV6ProtocolUdpXml_11_0_2 `xml:"udp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdV6ProtocolAnyXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdV6ProtocolTcpXml_11_0_2 struct {
	LocalPort  *int64 `xml:"local-port,omitempty"`
	RemotePort *int64 `xml:"remote-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AutoKeyProxyIdV6ProtocolUdpXml_11_0_2 struct {
	LocalPort  *int64 `xml:"local-port,omitempty"`
	RemotePort *int64 `xml:"remote-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type GlobalProtectSatelliteXml_11_0_2 struct {
	ExternalCa             *GlobalProtectSatelliteExternalCaXml_11_0_2             `xml:"external-ca,omitempty"`
	Ipv6Preferred          *string                                                 `xml:"ipv6-preferred,omitempty"`
	LocalAddress           *GlobalProtectSatelliteLocalAddressXml_11_0_2           `xml:"local-address,omitempty"`
	PortalAddress          *string                                                 `xml:"portal-address,omitempty"`
	PublishConnectedRoutes *GlobalProtectSatellitePublishConnectedRoutesXml_11_0_2 `xml:"publish-connected-routes,omitempty"`
	PublishRoutes          *util.MemberType                                        `xml:"publish-routes,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type GlobalProtectSatelliteExternalCaXml_11_0_2 struct {
	CertificateProfile *string `xml:"certificate-profile,omitempty"`
	LocalCertificate   *string `xml:"local-certificate,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type GlobalProtectSatelliteLocalAddressXml_11_0_2 struct {
	Interface  *string                                                 `xml:"interface,omitempty"`
	FloatingIp *GlobalProtectSatelliteLocalAddressFloatingIpXml_11_0_2 `xml:"floating-ip,omitempty"`
	Ip         *GlobalProtectSatelliteLocalAddressIpXml_11_0_2         `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type GlobalProtectSatelliteLocalAddressFloatingIpXml_11_0_2 struct {
	Ipv4 *string `xml:"ipv4,omitempty"`
	Ipv6 *string `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type GlobalProtectSatelliteLocalAddressIpXml_11_0_2 struct {
	Ipv4 *string `xml:"ipv4,omitempty"`
	Ipv6 *string `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type GlobalProtectSatellitePublishConnectedRoutesXml_11_0_2 struct {
	Enable *string `xml:"enable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyXml_11_0_2 struct {
	LocalAddress *ManualKeyLocalAddressXml_11_0_2 `xml:"local-address,omitempty"`
	LocalSpi     *string                          `xml:"local-spi,omitempty"`
	PeerAddress  *ManualKeyPeerAddressXml_11_0_2  `xml:"peer-address,omitempty"`
	RemoteSpi    *string                          `xml:"remote-spi,omitempty"`
	Ah           *ManualKeyAhXml_11_0_2           `xml:"ah,omitempty"`
	Esp          *ManualKeyEspXml_11_0_2          `xml:"esp,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyAhXml_11_0_2 struct {
	Md5    *ManualKeyAhMd5Xml_11_0_2    `xml:"md5,omitempty"`
	Sha1   *ManualKeyAhSha1Xml_11_0_2   `xml:"sha1,omitempty"`
	Sha256 *ManualKeyAhSha256Xml_11_0_2 `xml:"sha256,omitempty"`
	Sha384 *ManualKeyAhSha384Xml_11_0_2 `xml:"sha384,omitempty"`
	Sha512 *ManualKeyAhSha512Xml_11_0_2 `xml:"sha512,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyAhMd5Xml_11_0_2 struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyAhSha1Xml_11_0_2 struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyAhSha256Xml_11_0_2 struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyAhSha384Xml_11_0_2 struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyAhSha512Xml_11_0_2 struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspXml_11_0_2 struct {
	Authentication *ManualKeyEspAuthenticationXml_11_0_2 `xml:"authentication,omitempty"`
	Encryption     *ManualKeyEspEncryptionXml_11_0_2     `xml:"encryption,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspAuthenticationXml_11_0_2 struct {
	Md5    *ManualKeyEspAuthenticationMd5Xml_11_0_2    `xml:"md5,omitempty"`
	None   *ManualKeyEspAuthenticationNoneXml_11_0_2   `xml:"none,omitempty"`
	Sha1   *ManualKeyEspAuthenticationSha1Xml_11_0_2   `xml:"sha1,omitempty"`
	Sha256 *ManualKeyEspAuthenticationSha256Xml_11_0_2 `xml:"sha256,omitempty"`
	Sha384 *ManualKeyEspAuthenticationSha384Xml_11_0_2 `xml:"sha384,omitempty"`
	Sha512 *ManualKeyEspAuthenticationSha512Xml_11_0_2 `xml:"sha512,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspAuthenticationMd5Xml_11_0_2 struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspAuthenticationNoneXml_11_0_2 struct {
	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspAuthenticationSha1Xml_11_0_2 struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspAuthenticationSha256Xml_11_0_2 struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspAuthenticationSha384Xml_11_0_2 struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspAuthenticationSha512Xml_11_0_2 struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyEspEncryptionXml_11_0_2 struct {
	Algorithm *string `xml:"algorithm,omitempty"`
	Key       *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyLocalAddressXml_11_0_2 struct {
	Interface  *string `xml:"interface,omitempty"`
	FloatingIp *string `xml:"floating-ip,omitempty"`
	Ip         *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ManualKeyPeerAddressXml_11_0_2 struct {
	Ip *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type TunnelMonitorXml_11_0_2 struct {
	DestinationIp        *string `xml:"destination-ip,omitempty"`
	Enable               *string `xml:"enable,omitempty"`
	ProxyId              *string `xml:"proxy-id,omitempty"`
	TunnelMonitorProfile *string `xml:"tunnel-monitor-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
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
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	entry.AntiReplay = util.YesNo(o.AntiReplay, nil)
	entry.AntiReplayWindow = o.AntiReplayWindow
	entry.Comment = o.Comment
	entry.CopyFlowLabel = util.YesNo(o.CopyFlowLabel, nil)
	entry.CopyTos = util.YesNo(o.CopyTos, nil)
	entry.Disabled = util.YesNo(o.Disabled, nil)
	entry.EnableGreEncapsulation = util.YesNo(o.EnableGreEncapsulation, nil)
	entry.IpsecMode = o.IpsecMode
	entry.Ipv6 = util.YesNo(o.Ipv6, nil)
	entry.TunnelInterface = o.TunnelInterface
	var nestedTunnelMonitor *TunnelMonitorXml
	if o.TunnelMonitor != nil {
		nestedTunnelMonitor = &TunnelMonitorXml{}
		if _, ok := o.Misc["TunnelMonitor"]; ok {
			nestedTunnelMonitor.Misc = o.Misc["TunnelMonitor"]
		}
		if o.TunnelMonitor.DestinationIp != nil {
			nestedTunnelMonitor.DestinationIp = o.TunnelMonitor.DestinationIp
		}
		if o.TunnelMonitor.Enable != nil {
			nestedTunnelMonitor.Enable = util.YesNo(o.TunnelMonitor.Enable, nil)
		}
		if o.TunnelMonitor.ProxyId != nil {
			nestedTunnelMonitor.ProxyId = o.TunnelMonitor.ProxyId
		}
		if o.TunnelMonitor.TunnelMonitorProfile != nil {
			nestedTunnelMonitor.TunnelMonitorProfile = o.TunnelMonitor.TunnelMonitorProfile
		}
	}
	entry.TunnelMonitor = nestedTunnelMonitor

	var nestedAutoKey *AutoKeyXml
	if o.AutoKey != nil {
		nestedAutoKey = &AutoKeyXml{}
		if _, ok := o.Misc["AutoKey"]; ok {
			nestedAutoKey.Misc = o.Misc["AutoKey"]
		}
		if o.AutoKey.IkeGateway != nil {
			nestedAutoKey.IkeGateway = []AutoKeyIkeGatewayXml{}
			for _, oAutoKeyIkeGateway := range o.AutoKey.IkeGateway {
				nestedAutoKeyIkeGateway := AutoKeyIkeGatewayXml{}
				if _, ok := o.Misc["AutoKeyIkeGateway"]; ok {
					nestedAutoKeyIkeGateway.Misc = o.Misc["AutoKeyIkeGateway"]
				}
				if oAutoKeyIkeGateway.Name != "" {
					nestedAutoKeyIkeGateway.Name = oAutoKeyIkeGateway.Name
				}
				nestedAutoKey.IkeGateway = append(nestedAutoKey.IkeGateway, nestedAutoKeyIkeGateway)
			}
		}
		if o.AutoKey.IpsecCryptoProfile != nil {
			nestedAutoKey.IpsecCryptoProfile = o.AutoKey.IpsecCryptoProfile
		}
		if o.AutoKey.ProxyId != nil {
			nestedAutoKey.ProxyId = []AutoKeyProxyIdXml{}
			for _, oAutoKeyProxyId := range o.AutoKey.ProxyId {
				nestedAutoKeyProxyId := AutoKeyProxyIdXml{}
				if _, ok := o.Misc["AutoKeyProxyId"]; ok {
					nestedAutoKeyProxyId.Misc = o.Misc["AutoKeyProxyId"]
				}
				if oAutoKeyProxyId.Local != nil {
					nestedAutoKeyProxyId.Local = oAutoKeyProxyId.Local
				}
				if oAutoKeyProxyId.Remote != nil {
					nestedAutoKeyProxyId.Remote = oAutoKeyProxyId.Remote
				}
				if oAutoKeyProxyId.Protocol != nil {
					nestedAutoKeyProxyId.Protocol = &AutoKeyProxyIdProtocolXml{}
					if _, ok := o.Misc["AutoKeyProxyIdProtocol"]; ok {
						nestedAutoKeyProxyId.Protocol.Misc = o.Misc["AutoKeyProxyIdProtocol"]
					}
					if oAutoKeyProxyId.Protocol.Any != nil {
						nestedAutoKeyProxyId.Protocol.Any = &AutoKeyProxyIdProtocolAnyXml{}
						if _, ok := o.Misc["AutoKeyProxyIdProtocolAny"]; ok {
							nestedAutoKeyProxyId.Protocol.Any.Misc = o.Misc["AutoKeyProxyIdProtocolAny"]
						}
					}
					if oAutoKeyProxyId.Protocol.Tcp != nil {
						nestedAutoKeyProxyId.Protocol.Tcp = &AutoKeyProxyIdProtocolTcpXml{}
						if _, ok := o.Misc["AutoKeyProxyIdProtocolTcp"]; ok {
							nestedAutoKeyProxyId.Protocol.Tcp.Misc = o.Misc["AutoKeyProxyIdProtocolTcp"]
						}
						if oAutoKeyProxyId.Protocol.Tcp.LocalPort != nil {
							nestedAutoKeyProxyId.Protocol.Tcp.LocalPort = oAutoKeyProxyId.Protocol.Tcp.LocalPort
						}
						if oAutoKeyProxyId.Protocol.Tcp.RemotePort != nil {
							nestedAutoKeyProxyId.Protocol.Tcp.RemotePort = oAutoKeyProxyId.Protocol.Tcp.RemotePort
						}
					}
					if oAutoKeyProxyId.Protocol.Udp != nil {
						nestedAutoKeyProxyId.Protocol.Udp = &AutoKeyProxyIdProtocolUdpXml{}
						if _, ok := o.Misc["AutoKeyProxyIdProtocolUdp"]; ok {
							nestedAutoKeyProxyId.Protocol.Udp.Misc = o.Misc["AutoKeyProxyIdProtocolUdp"]
						}
						if oAutoKeyProxyId.Protocol.Udp.LocalPort != nil {
							nestedAutoKeyProxyId.Protocol.Udp.LocalPort = oAutoKeyProxyId.Protocol.Udp.LocalPort
						}
						if oAutoKeyProxyId.Protocol.Udp.RemotePort != nil {
							nestedAutoKeyProxyId.Protocol.Udp.RemotePort = oAutoKeyProxyId.Protocol.Udp.RemotePort
						}
					}
					if oAutoKeyProxyId.Protocol.Number != nil {
						nestedAutoKeyProxyId.Protocol.Number = oAutoKeyProxyId.Protocol.Number
					}
				}
				if oAutoKeyProxyId.Name != "" {
					nestedAutoKeyProxyId.Name = oAutoKeyProxyId.Name
				}
				nestedAutoKey.ProxyId = append(nestedAutoKey.ProxyId, nestedAutoKeyProxyId)
			}
		}
		if o.AutoKey.ProxyIdV6 != nil {
			nestedAutoKey.ProxyIdV6 = []AutoKeyProxyIdV6Xml{}
			for _, oAutoKeyProxyIdV6 := range o.AutoKey.ProxyIdV6 {
				nestedAutoKeyProxyIdV6 := AutoKeyProxyIdV6Xml{}
				if _, ok := o.Misc["AutoKeyProxyIdV6"]; ok {
					nestedAutoKeyProxyIdV6.Misc = o.Misc["AutoKeyProxyIdV6"]
				}
				if oAutoKeyProxyIdV6.Local != nil {
					nestedAutoKeyProxyIdV6.Local = oAutoKeyProxyIdV6.Local
				}
				if oAutoKeyProxyIdV6.Remote != nil {
					nestedAutoKeyProxyIdV6.Remote = oAutoKeyProxyIdV6.Remote
				}
				if oAutoKeyProxyIdV6.Protocol != nil {
					nestedAutoKeyProxyIdV6.Protocol = &AutoKeyProxyIdV6ProtocolXml{}
					if _, ok := o.Misc["AutoKeyProxyIdV6Protocol"]; ok {
						nestedAutoKeyProxyIdV6.Protocol.Misc = o.Misc["AutoKeyProxyIdV6Protocol"]
					}
					if oAutoKeyProxyIdV6.Protocol.Number != nil {
						nestedAutoKeyProxyIdV6.Protocol.Number = oAutoKeyProxyIdV6.Protocol.Number
					}
					if oAutoKeyProxyIdV6.Protocol.Any != nil {
						nestedAutoKeyProxyIdV6.Protocol.Any = &AutoKeyProxyIdV6ProtocolAnyXml{}
						if _, ok := o.Misc["AutoKeyProxyIdV6ProtocolAny"]; ok {
							nestedAutoKeyProxyIdV6.Protocol.Any.Misc = o.Misc["AutoKeyProxyIdV6ProtocolAny"]
						}
					}
					if oAutoKeyProxyIdV6.Protocol.Tcp != nil {
						nestedAutoKeyProxyIdV6.Protocol.Tcp = &AutoKeyProxyIdV6ProtocolTcpXml{}
						if _, ok := o.Misc["AutoKeyProxyIdV6ProtocolTcp"]; ok {
							nestedAutoKeyProxyIdV6.Protocol.Tcp.Misc = o.Misc["AutoKeyProxyIdV6ProtocolTcp"]
						}
						if oAutoKeyProxyIdV6.Protocol.Tcp.LocalPort != nil {
							nestedAutoKeyProxyIdV6.Protocol.Tcp.LocalPort = oAutoKeyProxyIdV6.Protocol.Tcp.LocalPort
						}
						if oAutoKeyProxyIdV6.Protocol.Tcp.RemotePort != nil {
							nestedAutoKeyProxyIdV6.Protocol.Tcp.RemotePort = oAutoKeyProxyIdV6.Protocol.Tcp.RemotePort
						}
					}
					if oAutoKeyProxyIdV6.Protocol.Udp != nil {
						nestedAutoKeyProxyIdV6.Protocol.Udp = &AutoKeyProxyIdV6ProtocolUdpXml{}
						if _, ok := o.Misc["AutoKeyProxyIdV6ProtocolUdp"]; ok {
							nestedAutoKeyProxyIdV6.Protocol.Udp.Misc = o.Misc["AutoKeyProxyIdV6ProtocolUdp"]
						}
						if oAutoKeyProxyIdV6.Protocol.Udp.LocalPort != nil {
							nestedAutoKeyProxyIdV6.Protocol.Udp.LocalPort = oAutoKeyProxyIdV6.Protocol.Udp.LocalPort
						}
						if oAutoKeyProxyIdV6.Protocol.Udp.RemotePort != nil {
							nestedAutoKeyProxyIdV6.Protocol.Udp.RemotePort = oAutoKeyProxyIdV6.Protocol.Udp.RemotePort
						}
					}
				}
				if oAutoKeyProxyIdV6.Name != "" {
					nestedAutoKeyProxyIdV6.Name = oAutoKeyProxyIdV6.Name
				}
				nestedAutoKey.ProxyIdV6 = append(nestedAutoKey.ProxyIdV6, nestedAutoKeyProxyIdV6)
			}
		}
	}
	entry.AutoKey = nestedAutoKey

	var nestedGlobalProtectSatellite *GlobalProtectSatelliteXml
	if o.GlobalProtectSatellite != nil {
		nestedGlobalProtectSatellite = &GlobalProtectSatelliteXml{}
		if _, ok := o.Misc["GlobalProtectSatellite"]; ok {
			nestedGlobalProtectSatellite.Misc = o.Misc["GlobalProtectSatellite"]
		}
		if o.GlobalProtectSatellite.Ipv6Preferred != nil {
			nestedGlobalProtectSatellite.Ipv6Preferred = util.YesNo(o.GlobalProtectSatellite.Ipv6Preferred, nil)
		}
		if o.GlobalProtectSatellite.LocalAddress != nil {
			nestedGlobalProtectSatellite.LocalAddress = &GlobalProtectSatelliteLocalAddressXml{}
			if _, ok := o.Misc["GlobalProtectSatelliteLocalAddress"]; ok {
				nestedGlobalProtectSatellite.LocalAddress.Misc = o.Misc["GlobalProtectSatelliteLocalAddress"]
			}
			if o.GlobalProtectSatellite.LocalAddress.Interface != nil {
				nestedGlobalProtectSatellite.LocalAddress.Interface = o.GlobalProtectSatellite.LocalAddress.Interface
			}
			if o.GlobalProtectSatellite.LocalAddress.FloatingIp != nil {
				nestedGlobalProtectSatellite.LocalAddress.FloatingIp = &GlobalProtectSatelliteLocalAddressFloatingIpXml{}
				if _, ok := o.Misc["GlobalProtectSatelliteLocalAddressFloatingIp"]; ok {
					nestedGlobalProtectSatellite.LocalAddress.FloatingIp.Misc = o.Misc["GlobalProtectSatelliteLocalAddressFloatingIp"]
				}
				if o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv4 != nil {
					nestedGlobalProtectSatellite.LocalAddress.FloatingIp.Ipv4 = o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv4
				}
				if o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv6 != nil {
					nestedGlobalProtectSatellite.LocalAddress.FloatingIp.Ipv6 = o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv6
				}
			}
			if o.GlobalProtectSatellite.LocalAddress.Ip != nil {
				nestedGlobalProtectSatellite.LocalAddress.Ip = &GlobalProtectSatelliteLocalAddressIpXml{}
				if _, ok := o.Misc["GlobalProtectSatelliteLocalAddressIp"]; ok {
					nestedGlobalProtectSatellite.LocalAddress.Ip.Misc = o.Misc["GlobalProtectSatelliteLocalAddressIp"]
				}
				if o.GlobalProtectSatellite.LocalAddress.Ip.Ipv6 != nil {
					nestedGlobalProtectSatellite.LocalAddress.Ip.Ipv6 = o.GlobalProtectSatellite.LocalAddress.Ip.Ipv6
				}
				if o.GlobalProtectSatellite.LocalAddress.Ip.Ipv4 != nil {
					nestedGlobalProtectSatellite.LocalAddress.Ip.Ipv4 = o.GlobalProtectSatellite.LocalAddress.Ip.Ipv4
				}
			}
		}
		if o.GlobalProtectSatellite.PortalAddress != nil {
			nestedGlobalProtectSatellite.PortalAddress = o.GlobalProtectSatellite.PortalAddress
		}
		if o.GlobalProtectSatellite.PublishConnectedRoutes != nil {
			nestedGlobalProtectSatellite.PublishConnectedRoutes = &GlobalProtectSatellitePublishConnectedRoutesXml{}
			if _, ok := o.Misc["GlobalProtectSatellitePublishConnectedRoutes"]; ok {
				nestedGlobalProtectSatellite.PublishConnectedRoutes.Misc = o.Misc["GlobalProtectSatellitePublishConnectedRoutes"]
			}
			if o.GlobalProtectSatellite.PublishConnectedRoutes.Enable != nil {
				nestedGlobalProtectSatellite.PublishConnectedRoutes.Enable = util.YesNo(o.GlobalProtectSatellite.PublishConnectedRoutes.Enable, nil)
			}
		}
		if o.GlobalProtectSatellite.PublishRoutes != nil {
			nestedGlobalProtectSatellite.PublishRoutes = util.StrToMem(o.GlobalProtectSatellite.PublishRoutes)
		}
		if o.GlobalProtectSatellite.ExternalCa != nil {
			nestedGlobalProtectSatellite.ExternalCa = &GlobalProtectSatelliteExternalCaXml{}
			if _, ok := o.Misc["GlobalProtectSatelliteExternalCa"]; ok {
				nestedGlobalProtectSatellite.ExternalCa.Misc = o.Misc["GlobalProtectSatelliteExternalCa"]
			}
			if o.GlobalProtectSatellite.ExternalCa.CertificateProfile != nil {
				nestedGlobalProtectSatellite.ExternalCa.CertificateProfile = o.GlobalProtectSatellite.ExternalCa.CertificateProfile
			}
			if o.GlobalProtectSatellite.ExternalCa.LocalCertificate != nil {
				nestedGlobalProtectSatellite.ExternalCa.LocalCertificate = o.GlobalProtectSatellite.ExternalCa.LocalCertificate
			}
		}
	}
	entry.GlobalProtectSatellite = nestedGlobalProtectSatellite

	var nestedManualKey *ManualKeyXml
	if o.ManualKey != nil {
		nestedManualKey = &ManualKeyXml{}
		if _, ok := o.Misc["ManualKey"]; ok {
			nestedManualKey.Misc = o.Misc["ManualKey"]
		}
		if o.ManualKey.LocalSpi != nil {
			nestedManualKey.LocalSpi = o.ManualKey.LocalSpi
		}
		if o.ManualKey.PeerAddress != nil {
			nestedManualKey.PeerAddress = &ManualKeyPeerAddressXml{}
			if _, ok := o.Misc["ManualKeyPeerAddress"]; ok {
				nestedManualKey.PeerAddress.Misc = o.Misc["ManualKeyPeerAddress"]
			}
			if o.ManualKey.PeerAddress.Ip != nil {
				nestedManualKey.PeerAddress.Ip = o.ManualKey.PeerAddress.Ip
			}
		}
		if o.ManualKey.RemoteSpi != nil {
			nestedManualKey.RemoteSpi = o.ManualKey.RemoteSpi
		}
		if o.ManualKey.LocalAddress != nil {
			nestedManualKey.LocalAddress = &ManualKeyLocalAddressXml{}
			if _, ok := o.Misc["ManualKeyLocalAddress"]; ok {
				nestedManualKey.LocalAddress.Misc = o.Misc["ManualKeyLocalAddress"]
			}
			if o.ManualKey.LocalAddress.Interface != nil {
				nestedManualKey.LocalAddress.Interface = o.ManualKey.LocalAddress.Interface
			}
			if o.ManualKey.LocalAddress.FloatingIp != nil {
				nestedManualKey.LocalAddress.FloatingIp = o.ManualKey.LocalAddress.FloatingIp
			}
			if o.ManualKey.LocalAddress.Ip != nil {
				nestedManualKey.LocalAddress.Ip = o.ManualKey.LocalAddress.Ip
			}
		}
		if o.ManualKey.Ah != nil {
			nestedManualKey.Ah = &ManualKeyAhXml{}
			if _, ok := o.Misc["ManualKeyAh"]; ok {
				nestedManualKey.Ah.Misc = o.Misc["ManualKeyAh"]
			}
			if o.ManualKey.Ah.Sha1 != nil {
				nestedManualKey.Ah.Sha1 = &ManualKeyAhSha1Xml{}
				if _, ok := o.Misc["ManualKeyAhSha1"]; ok {
					nestedManualKey.Ah.Sha1.Misc = o.Misc["ManualKeyAhSha1"]
				}
				if o.ManualKey.Ah.Sha1.Key != nil {
					nestedManualKey.Ah.Sha1.Key = o.ManualKey.Ah.Sha1.Key
				}
			}
			if o.ManualKey.Ah.Sha256 != nil {
				nestedManualKey.Ah.Sha256 = &ManualKeyAhSha256Xml{}
				if _, ok := o.Misc["ManualKeyAhSha256"]; ok {
					nestedManualKey.Ah.Sha256.Misc = o.Misc["ManualKeyAhSha256"]
				}
				if o.ManualKey.Ah.Sha256.Key != nil {
					nestedManualKey.Ah.Sha256.Key = o.ManualKey.Ah.Sha256.Key
				}
			}
			if o.ManualKey.Ah.Sha384 != nil {
				nestedManualKey.Ah.Sha384 = &ManualKeyAhSha384Xml{}
				if _, ok := o.Misc["ManualKeyAhSha384"]; ok {
					nestedManualKey.Ah.Sha384.Misc = o.Misc["ManualKeyAhSha384"]
				}
				if o.ManualKey.Ah.Sha384.Key != nil {
					nestedManualKey.Ah.Sha384.Key = o.ManualKey.Ah.Sha384.Key
				}
			}
			if o.ManualKey.Ah.Sha512 != nil {
				nestedManualKey.Ah.Sha512 = &ManualKeyAhSha512Xml{}
				if _, ok := o.Misc["ManualKeyAhSha512"]; ok {
					nestedManualKey.Ah.Sha512.Misc = o.Misc["ManualKeyAhSha512"]
				}
				if o.ManualKey.Ah.Sha512.Key != nil {
					nestedManualKey.Ah.Sha512.Key = o.ManualKey.Ah.Sha512.Key
				}
			}
			if o.ManualKey.Ah.Md5 != nil {
				nestedManualKey.Ah.Md5 = &ManualKeyAhMd5Xml{}
				if _, ok := o.Misc["ManualKeyAhMd5"]; ok {
					nestedManualKey.Ah.Md5.Misc = o.Misc["ManualKeyAhMd5"]
				}
				if o.ManualKey.Ah.Md5.Key != nil {
					nestedManualKey.Ah.Md5.Key = o.ManualKey.Ah.Md5.Key
				}
			}
		}
		if o.ManualKey.Esp != nil {
			nestedManualKey.Esp = &ManualKeyEspXml{}
			if _, ok := o.Misc["ManualKeyEsp"]; ok {
				nestedManualKey.Esp.Misc = o.Misc["ManualKeyEsp"]
			}
			if o.ManualKey.Esp.Authentication != nil {
				nestedManualKey.Esp.Authentication = &ManualKeyEspAuthenticationXml{}
				if _, ok := o.Misc["ManualKeyEspAuthentication"]; ok {
					nestedManualKey.Esp.Authentication.Misc = o.Misc["ManualKeyEspAuthentication"]
				}
				if o.ManualKey.Esp.Authentication.Sha384 != nil {
					nestedManualKey.Esp.Authentication.Sha384 = &ManualKeyEspAuthenticationSha384Xml{}
					if _, ok := o.Misc["ManualKeyEspAuthenticationSha384"]; ok {
						nestedManualKey.Esp.Authentication.Sha384.Misc = o.Misc["ManualKeyEspAuthenticationSha384"]
					}
					if o.ManualKey.Esp.Authentication.Sha384.Key != nil {
						nestedManualKey.Esp.Authentication.Sha384.Key = o.ManualKey.Esp.Authentication.Sha384.Key
					}
				}
				if o.ManualKey.Esp.Authentication.Sha512 != nil {
					nestedManualKey.Esp.Authentication.Sha512 = &ManualKeyEspAuthenticationSha512Xml{}
					if _, ok := o.Misc["ManualKeyEspAuthenticationSha512"]; ok {
						nestedManualKey.Esp.Authentication.Sha512.Misc = o.Misc["ManualKeyEspAuthenticationSha512"]
					}
					if o.ManualKey.Esp.Authentication.Sha512.Key != nil {
						nestedManualKey.Esp.Authentication.Sha512.Key = o.ManualKey.Esp.Authentication.Sha512.Key
					}
				}
				if o.ManualKey.Esp.Authentication.Md5 != nil {
					nestedManualKey.Esp.Authentication.Md5 = &ManualKeyEspAuthenticationMd5Xml{}
					if _, ok := o.Misc["ManualKeyEspAuthenticationMd5"]; ok {
						nestedManualKey.Esp.Authentication.Md5.Misc = o.Misc["ManualKeyEspAuthenticationMd5"]
					}
					if o.ManualKey.Esp.Authentication.Md5.Key != nil {
						nestedManualKey.Esp.Authentication.Md5.Key = o.ManualKey.Esp.Authentication.Md5.Key
					}
				}
				if o.ManualKey.Esp.Authentication.None != nil {
					nestedManualKey.Esp.Authentication.None = &ManualKeyEspAuthenticationNoneXml{}
					if _, ok := o.Misc["ManualKeyEspAuthenticationNone"]; ok {
						nestedManualKey.Esp.Authentication.None.Misc = o.Misc["ManualKeyEspAuthenticationNone"]
					}
				}
				if o.ManualKey.Esp.Authentication.Sha1 != nil {
					nestedManualKey.Esp.Authentication.Sha1 = &ManualKeyEspAuthenticationSha1Xml{}
					if _, ok := o.Misc["ManualKeyEspAuthenticationSha1"]; ok {
						nestedManualKey.Esp.Authentication.Sha1.Misc = o.Misc["ManualKeyEspAuthenticationSha1"]
					}
					if o.ManualKey.Esp.Authentication.Sha1.Key != nil {
						nestedManualKey.Esp.Authentication.Sha1.Key = o.ManualKey.Esp.Authentication.Sha1.Key
					}
				}
				if o.ManualKey.Esp.Authentication.Sha256 != nil {
					nestedManualKey.Esp.Authentication.Sha256 = &ManualKeyEspAuthenticationSha256Xml{}
					if _, ok := o.Misc["ManualKeyEspAuthenticationSha256"]; ok {
						nestedManualKey.Esp.Authentication.Sha256.Misc = o.Misc["ManualKeyEspAuthenticationSha256"]
					}
					if o.ManualKey.Esp.Authentication.Sha256.Key != nil {
						nestedManualKey.Esp.Authentication.Sha256.Key = o.ManualKey.Esp.Authentication.Sha256.Key
					}
				}
			}
			if o.ManualKey.Esp.Encryption != nil {
				nestedManualKey.Esp.Encryption = &ManualKeyEspEncryptionXml{}
				if _, ok := o.Misc["ManualKeyEspEncryption"]; ok {
					nestedManualKey.Esp.Encryption.Misc = o.Misc["ManualKeyEspEncryption"]
				}
				if o.ManualKey.Esp.Encryption.Algorithm != nil {
					nestedManualKey.Esp.Encryption.Algorithm = o.ManualKey.Esp.Encryption.Algorithm
				}
				if o.ManualKey.Esp.Encryption.Key != nil {
					nestedManualKey.Esp.Encryption.Key = o.ManualKey.Esp.Encryption.Key
				}
			}
		}
	}
	entry.ManualKey = nestedManualKey

	entry.Misc = o.Misc["Entry"]

	return entry, nil
}

func specifyEntry_11_0_2(o *Entry) (any, error) {
	entry := entryXml_11_0_2{}
	entry.Name = o.Name
	entry.AntiReplay = util.YesNo(o.AntiReplay, nil)
	entry.AntiReplayWindow = o.AntiReplayWindow
	entry.Comment = o.Comment
	entry.CopyFlowLabel = util.YesNo(o.CopyFlowLabel, nil)
	entry.CopyTos = util.YesNo(o.CopyTos, nil)
	entry.Disabled = util.YesNo(o.Disabled, nil)
	entry.EnableGreEncapsulation = util.YesNo(o.EnableGreEncapsulation, nil)
	entry.IpsecMode = o.IpsecMode
	entry.Ipv6 = util.YesNo(o.Ipv6, nil)
	entry.TunnelInterface = o.TunnelInterface
	var nestedTunnelMonitor *TunnelMonitorXml_11_0_2
	if o.TunnelMonitor != nil {
		nestedTunnelMonitor = &TunnelMonitorXml_11_0_2{}
		if _, ok := o.Misc["TunnelMonitor"]; ok {
			nestedTunnelMonitor.Misc = o.Misc["TunnelMonitor"]
		}
		if o.TunnelMonitor.DestinationIp != nil {
			nestedTunnelMonitor.DestinationIp = o.TunnelMonitor.DestinationIp
		}
		if o.TunnelMonitor.Enable != nil {
			nestedTunnelMonitor.Enable = util.YesNo(o.TunnelMonitor.Enable, nil)
		}
		if o.TunnelMonitor.ProxyId != nil {
			nestedTunnelMonitor.ProxyId = o.TunnelMonitor.ProxyId
		}
		if o.TunnelMonitor.TunnelMonitorProfile != nil {
			nestedTunnelMonitor.TunnelMonitorProfile = o.TunnelMonitor.TunnelMonitorProfile
		}
	}
	entry.TunnelMonitor = nestedTunnelMonitor

	var nestedAutoKey *AutoKeyXml_11_0_2
	if o.AutoKey != nil {
		nestedAutoKey = &AutoKeyXml_11_0_2{}
		if _, ok := o.Misc["AutoKey"]; ok {
			nestedAutoKey.Misc = o.Misc["AutoKey"]
		}
		if o.AutoKey.ProxyId != nil {
			nestedAutoKey.ProxyId = []AutoKeyProxyIdXml_11_0_2{}
			for _, oAutoKeyProxyId := range o.AutoKey.ProxyId {
				nestedAutoKeyProxyId := AutoKeyProxyIdXml_11_0_2{}
				if _, ok := o.Misc["AutoKeyProxyId"]; ok {
					nestedAutoKeyProxyId.Misc = o.Misc["AutoKeyProxyId"]
				}
				if oAutoKeyProxyId.Remote != nil {
					nestedAutoKeyProxyId.Remote = oAutoKeyProxyId.Remote
				}
				if oAutoKeyProxyId.Protocol != nil {
					nestedAutoKeyProxyId.Protocol = &AutoKeyProxyIdProtocolXml_11_0_2{}
					if _, ok := o.Misc["AutoKeyProxyIdProtocol"]; ok {
						nestedAutoKeyProxyId.Protocol.Misc = o.Misc["AutoKeyProxyIdProtocol"]
					}
					if oAutoKeyProxyId.Protocol.Number != nil {
						nestedAutoKeyProxyId.Protocol.Number = oAutoKeyProxyId.Protocol.Number
					}
					if oAutoKeyProxyId.Protocol.Any != nil {
						nestedAutoKeyProxyId.Protocol.Any = &AutoKeyProxyIdProtocolAnyXml_11_0_2{}
						if _, ok := o.Misc["AutoKeyProxyIdProtocolAny"]; ok {
							nestedAutoKeyProxyId.Protocol.Any.Misc = o.Misc["AutoKeyProxyIdProtocolAny"]
						}
					}
					if oAutoKeyProxyId.Protocol.Tcp != nil {
						nestedAutoKeyProxyId.Protocol.Tcp = &AutoKeyProxyIdProtocolTcpXml_11_0_2{}
						if _, ok := o.Misc["AutoKeyProxyIdProtocolTcp"]; ok {
							nestedAutoKeyProxyId.Protocol.Tcp.Misc = o.Misc["AutoKeyProxyIdProtocolTcp"]
						}
						if oAutoKeyProxyId.Protocol.Tcp.LocalPort != nil {
							nestedAutoKeyProxyId.Protocol.Tcp.LocalPort = oAutoKeyProxyId.Protocol.Tcp.LocalPort
						}
						if oAutoKeyProxyId.Protocol.Tcp.RemotePort != nil {
							nestedAutoKeyProxyId.Protocol.Tcp.RemotePort = oAutoKeyProxyId.Protocol.Tcp.RemotePort
						}
					}
					if oAutoKeyProxyId.Protocol.Udp != nil {
						nestedAutoKeyProxyId.Protocol.Udp = &AutoKeyProxyIdProtocolUdpXml_11_0_2{}
						if _, ok := o.Misc["AutoKeyProxyIdProtocolUdp"]; ok {
							nestedAutoKeyProxyId.Protocol.Udp.Misc = o.Misc["AutoKeyProxyIdProtocolUdp"]
						}
						if oAutoKeyProxyId.Protocol.Udp.RemotePort != nil {
							nestedAutoKeyProxyId.Protocol.Udp.RemotePort = oAutoKeyProxyId.Protocol.Udp.RemotePort
						}
						if oAutoKeyProxyId.Protocol.Udp.LocalPort != nil {
							nestedAutoKeyProxyId.Protocol.Udp.LocalPort = oAutoKeyProxyId.Protocol.Udp.LocalPort
						}
					}
				}
				if oAutoKeyProxyId.Name != "" {
					nestedAutoKeyProxyId.Name = oAutoKeyProxyId.Name
				}
				if oAutoKeyProxyId.Local != nil {
					nestedAutoKeyProxyId.Local = oAutoKeyProxyId.Local
				}
				nestedAutoKey.ProxyId = append(nestedAutoKey.ProxyId, nestedAutoKeyProxyId)
			}
		}
		if o.AutoKey.ProxyIdV6 != nil {
			nestedAutoKey.ProxyIdV6 = []AutoKeyProxyIdV6Xml_11_0_2{}
			for _, oAutoKeyProxyIdV6 := range o.AutoKey.ProxyIdV6 {
				nestedAutoKeyProxyIdV6 := AutoKeyProxyIdV6Xml_11_0_2{}
				if _, ok := o.Misc["AutoKeyProxyIdV6"]; ok {
					nestedAutoKeyProxyIdV6.Misc = o.Misc["AutoKeyProxyIdV6"]
				}
				if oAutoKeyProxyIdV6.Protocol != nil {
					nestedAutoKeyProxyIdV6.Protocol = &AutoKeyProxyIdV6ProtocolXml_11_0_2{}
					if _, ok := o.Misc["AutoKeyProxyIdV6Protocol"]; ok {
						nestedAutoKeyProxyIdV6.Protocol.Misc = o.Misc["AutoKeyProxyIdV6Protocol"]
					}
					if oAutoKeyProxyIdV6.Protocol.Number != nil {
						nestedAutoKeyProxyIdV6.Protocol.Number = oAutoKeyProxyIdV6.Protocol.Number
					}
					if oAutoKeyProxyIdV6.Protocol.Any != nil {
						nestedAutoKeyProxyIdV6.Protocol.Any = &AutoKeyProxyIdV6ProtocolAnyXml_11_0_2{}
						if _, ok := o.Misc["AutoKeyProxyIdV6ProtocolAny"]; ok {
							nestedAutoKeyProxyIdV6.Protocol.Any.Misc = o.Misc["AutoKeyProxyIdV6ProtocolAny"]
						}
					}
					if oAutoKeyProxyIdV6.Protocol.Tcp != nil {
						nestedAutoKeyProxyIdV6.Protocol.Tcp = &AutoKeyProxyIdV6ProtocolTcpXml_11_0_2{}
						if _, ok := o.Misc["AutoKeyProxyIdV6ProtocolTcp"]; ok {
							nestedAutoKeyProxyIdV6.Protocol.Tcp.Misc = o.Misc["AutoKeyProxyIdV6ProtocolTcp"]
						}
						if oAutoKeyProxyIdV6.Protocol.Tcp.LocalPort != nil {
							nestedAutoKeyProxyIdV6.Protocol.Tcp.LocalPort = oAutoKeyProxyIdV6.Protocol.Tcp.LocalPort
						}
						if oAutoKeyProxyIdV6.Protocol.Tcp.RemotePort != nil {
							nestedAutoKeyProxyIdV6.Protocol.Tcp.RemotePort = oAutoKeyProxyIdV6.Protocol.Tcp.RemotePort
						}
					}
					if oAutoKeyProxyIdV6.Protocol.Udp != nil {
						nestedAutoKeyProxyIdV6.Protocol.Udp = &AutoKeyProxyIdV6ProtocolUdpXml_11_0_2{}
						if _, ok := o.Misc["AutoKeyProxyIdV6ProtocolUdp"]; ok {
							nestedAutoKeyProxyIdV6.Protocol.Udp.Misc = o.Misc["AutoKeyProxyIdV6ProtocolUdp"]
						}
						if oAutoKeyProxyIdV6.Protocol.Udp.LocalPort != nil {
							nestedAutoKeyProxyIdV6.Protocol.Udp.LocalPort = oAutoKeyProxyIdV6.Protocol.Udp.LocalPort
						}
						if oAutoKeyProxyIdV6.Protocol.Udp.RemotePort != nil {
							nestedAutoKeyProxyIdV6.Protocol.Udp.RemotePort = oAutoKeyProxyIdV6.Protocol.Udp.RemotePort
						}
					}
				}
				if oAutoKeyProxyIdV6.Name != "" {
					nestedAutoKeyProxyIdV6.Name = oAutoKeyProxyIdV6.Name
				}
				if oAutoKeyProxyIdV6.Local != nil {
					nestedAutoKeyProxyIdV6.Local = oAutoKeyProxyIdV6.Local
				}
				if oAutoKeyProxyIdV6.Remote != nil {
					nestedAutoKeyProxyIdV6.Remote = oAutoKeyProxyIdV6.Remote
				}
				nestedAutoKey.ProxyIdV6 = append(nestedAutoKey.ProxyIdV6, nestedAutoKeyProxyIdV6)
			}
		}
		if o.AutoKey.IkeGateway != nil {
			nestedAutoKey.IkeGateway = []AutoKeyIkeGatewayXml_11_0_2{}
			for _, oAutoKeyIkeGateway := range o.AutoKey.IkeGateway {
				nestedAutoKeyIkeGateway := AutoKeyIkeGatewayXml_11_0_2{}
				if _, ok := o.Misc["AutoKeyIkeGateway"]; ok {
					nestedAutoKeyIkeGateway.Misc = o.Misc["AutoKeyIkeGateway"]
				}
				if oAutoKeyIkeGateway.Name != "" {
					nestedAutoKeyIkeGateway.Name = oAutoKeyIkeGateway.Name
				}
				nestedAutoKey.IkeGateway = append(nestedAutoKey.IkeGateway, nestedAutoKeyIkeGateway)
			}
		}
		if o.AutoKey.IpsecCryptoProfile != nil {
			nestedAutoKey.IpsecCryptoProfile = o.AutoKey.IpsecCryptoProfile
		}
	}
	entry.AutoKey = nestedAutoKey

	var nestedGlobalProtectSatellite *GlobalProtectSatelliteXml_11_0_2
	if o.GlobalProtectSatellite != nil {
		nestedGlobalProtectSatellite = &GlobalProtectSatelliteXml_11_0_2{}
		if _, ok := o.Misc["GlobalProtectSatellite"]; ok {
			nestedGlobalProtectSatellite.Misc = o.Misc["GlobalProtectSatellite"]
		}
		if o.GlobalProtectSatellite.ExternalCa != nil {
			nestedGlobalProtectSatellite.ExternalCa = &GlobalProtectSatelliteExternalCaXml_11_0_2{}
			if _, ok := o.Misc["GlobalProtectSatelliteExternalCa"]; ok {
				nestedGlobalProtectSatellite.ExternalCa.Misc = o.Misc["GlobalProtectSatelliteExternalCa"]
			}
			if o.GlobalProtectSatellite.ExternalCa.CertificateProfile != nil {
				nestedGlobalProtectSatellite.ExternalCa.CertificateProfile = o.GlobalProtectSatellite.ExternalCa.CertificateProfile
			}
			if o.GlobalProtectSatellite.ExternalCa.LocalCertificate != nil {
				nestedGlobalProtectSatellite.ExternalCa.LocalCertificate = o.GlobalProtectSatellite.ExternalCa.LocalCertificate
			}
		}
		if o.GlobalProtectSatellite.Ipv6Preferred != nil {
			nestedGlobalProtectSatellite.Ipv6Preferred = util.YesNo(o.GlobalProtectSatellite.Ipv6Preferred, nil)
		}
		if o.GlobalProtectSatellite.LocalAddress != nil {
			nestedGlobalProtectSatellite.LocalAddress = &GlobalProtectSatelliteLocalAddressXml_11_0_2{}
			if _, ok := o.Misc["GlobalProtectSatelliteLocalAddress"]; ok {
				nestedGlobalProtectSatellite.LocalAddress.Misc = o.Misc["GlobalProtectSatelliteLocalAddress"]
			}
			if o.GlobalProtectSatellite.LocalAddress.Interface != nil {
				nestedGlobalProtectSatellite.LocalAddress.Interface = o.GlobalProtectSatellite.LocalAddress.Interface
			}
			if o.GlobalProtectSatellite.LocalAddress.FloatingIp != nil {
				nestedGlobalProtectSatellite.LocalAddress.FloatingIp = &GlobalProtectSatelliteLocalAddressFloatingIpXml_11_0_2{}
				if _, ok := o.Misc["GlobalProtectSatelliteLocalAddressFloatingIp"]; ok {
					nestedGlobalProtectSatellite.LocalAddress.FloatingIp.Misc = o.Misc["GlobalProtectSatelliteLocalAddressFloatingIp"]
				}
				if o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv4 != nil {
					nestedGlobalProtectSatellite.LocalAddress.FloatingIp.Ipv4 = o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv4
				}
				if o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv6 != nil {
					nestedGlobalProtectSatellite.LocalAddress.FloatingIp.Ipv6 = o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv6
				}
			}
			if o.GlobalProtectSatellite.LocalAddress.Ip != nil {
				nestedGlobalProtectSatellite.LocalAddress.Ip = &GlobalProtectSatelliteLocalAddressIpXml_11_0_2{}
				if _, ok := o.Misc["GlobalProtectSatelliteLocalAddressIp"]; ok {
					nestedGlobalProtectSatellite.LocalAddress.Ip.Misc = o.Misc["GlobalProtectSatelliteLocalAddressIp"]
				}
				if o.GlobalProtectSatellite.LocalAddress.Ip.Ipv4 != nil {
					nestedGlobalProtectSatellite.LocalAddress.Ip.Ipv4 = o.GlobalProtectSatellite.LocalAddress.Ip.Ipv4
				}
				if o.GlobalProtectSatellite.LocalAddress.Ip.Ipv6 != nil {
					nestedGlobalProtectSatellite.LocalAddress.Ip.Ipv6 = o.GlobalProtectSatellite.LocalAddress.Ip.Ipv6
				}
			}
		}
		if o.GlobalProtectSatellite.PortalAddress != nil {
			nestedGlobalProtectSatellite.PortalAddress = o.GlobalProtectSatellite.PortalAddress
		}
		if o.GlobalProtectSatellite.PublishConnectedRoutes != nil {
			nestedGlobalProtectSatellite.PublishConnectedRoutes = &GlobalProtectSatellitePublishConnectedRoutesXml_11_0_2{}
			if _, ok := o.Misc["GlobalProtectSatellitePublishConnectedRoutes"]; ok {
				nestedGlobalProtectSatellite.PublishConnectedRoutes.Misc = o.Misc["GlobalProtectSatellitePublishConnectedRoutes"]
			}
			if o.GlobalProtectSatellite.PublishConnectedRoutes.Enable != nil {
				nestedGlobalProtectSatellite.PublishConnectedRoutes.Enable = util.YesNo(o.GlobalProtectSatellite.PublishConnectedRoutes.Enable, nil)
			}
		}
		if o.GlobalProtectSatellite.PublishRoutes != nil {
			nestedGlobalProtectSatellite.PublishRoutes = util.StrToMem(o.GlobalProtectSatellite.PublishRoutes)
		}
	}
	entry.GlobalProtectSatellite = nestedGlobalProtectSatellite

	var nestedManualKey *ManualKeyXml_11_0_2
	if o.ManualKey != nil {
		nestedManualKey = &ManualKeyXml_11_0_2{}
		if _, ok := o.Misc["ManualKey"]; ok {
			nestedManualKey.Misc = o.Misc["ManualKey"]
		}
		if o.ManualKey.LocalAddress != nil {
			nestedManualKey.LocalAddress = &ManualKeyLocalAddressXml_11_0_2{}
			if _, ok := o.Misc["ManualKeyLocalAddress"]; ok {
				nestedManualKey.LocalAddress.Misc = o.Misc["ManualKeyLocalAddress"]
			}
			if o.ManualKey.LocalAddress.Interface != nil {
				nestedManualKey.LocalAddress.Interface = o.ManualKey.LocalAddress.Interface
			}
			if o.ManualKey.LocalAddress.FloatingIp != nil {
				nestedManualKey.LocalAddress.FloatingIp = o.ManualKey.LocalAddress.FloatingIp
			}
			if o.ManualKey.LocalAddress.Ip != nil {
				nestedManualKey.LocalAddress.Ip = o.ManualKey.LocalAddress.Ip
			}
		}
		if o.ManualKey.LocalSpi != nil {
			nestedManualKey.LocalSpi = o.ManualKey.LocalSpi
		}
		if o.ManualKey.PeerAddress != nil {
			nestedManualKey.PeerAddress = &ManualKeyPeerAddressXml_11_0_2{}
			if _, ok := o.Misc["ManualKeyPeerAddress"]; ok {
				nestedManualKey.PeerAddress.Misc = o.Misc["ManualKeyPeerAddress"]
			}
			if o.ManualKey.PeerAddress.Ip != nil {
				nestedManualKey.PeerAddress.Ip = o.ManualKey.PeerAddress.Ip
			}
		}
		if o.ManualKey.RemoteSpi != nil {
			nestedManualKey.RemoteSpi = o.ManualKey.RemoteSpi
		}
		if o.ManualKey.Ah != nil {
			nestedManualKey.Ah = &ManualKeyAhXml_11_0_2{}
			if _, ok := o.Misc["ManualKeyAh"]; ok {
				nestedManualKey.Ah.Misc = o.Misc["ManualKeyAh"]
			}
			if o.ManualKey.Ah.Sha1 != nil {
				nestedManualKey.Ah.Sha1 = &ManualKeyAhSha1Xml_11_0_2{}
				if _, ok := o.Misc["ManualKeyAhSha1"]; ok {
					nestedManualKey.Ah.Sha1.Misc = o.Misc["ManualKeyAhSha1"]
				}
				if o.ManualKey.Ah.Sha1.Key != nil {
					nestedManualKey.Ah.Sha1.Key = o.ManualKey.Ah.Sha1.Key
				}
			}
			if o.ManualKey.Ah.Sha256 != nil {
				nestedManualKey.Ah.Sha256 = &ManualKeyAhSha256Xml_11_0_2{}
				if _, ok := o.Misc["ManualKeyAhSha256"]; ok {
					nestedManualKey.Ah.Sha256.Misc = o.Misc["ManualKeyAhSha256"]
				}
				if o.ManualKey.Ah.Sha256.Key != nil {
					nestedManualKey.Ah.Sha256.Key = o.ManualKey.Ah.Sha256.Key
				}
			}
			if o.ManualKey.Ah.Sha384 != nil {
				nestedManualKey.Ah.Sha384 = &ManualKeyAhSha384Xml_11_0_2{}
				if _, ok := o.Misc["ManualKeyAhSha384"]; ok {
					nestedManualKey.Ah.Sha384.Misc = o.Misc["ManualKeyAhSha384"]
				}
				if o.ManualKey.Ah.Sha384.Key != nil {
					nestedManualKey.Ah.Sha384.Key = o.ManualKey.Ah.Sha384.Key
				}
			}
			if o.ManualKey.Ah.Sha512 != nil {
				nestedManualKey.Ah.Sha512 = &ManualKeyAhSha512Xml_11_0_2{}
				if _, ok := o.Misc["ManualKeyAhSha512"]; ok {
					nestedManualKey.Ah.Sha512.Misc = o.Misc["ManualKeyAhSha512"]
				}
				if o.ManualKey.Ah.Sha512.Key != nil {
					nestedManualKey.Ah.Sha512.Key = o.ManualKey.Ah.Sha512.Key
				}
			}
			if o.ManualKey.Ah.Md5 != nil {
				nestedManualKey.Ah.Md5 = &ManualKeyAhMd5Xml_11_0_2{}
				if _, ok := o.Misc["ManualKeyAhMd5"]; ok {
					nestedManualKey.Ah.Md5.Misc = o.Misc["ManualKeyAhMd5"]
				}
				if o.ManualKey.Ah.Md5.Key != nil {
					nestedManualKey.Ah.Md5.Key = o.ManualKey.Ah.Md5.Key
				}
			}
		}
		if o.ManualKey.Esp != nil {
			nestedManualKey.Esp = &ManualKeyEspXml_11_0_2{}
			if _, ok := o.Misc["ManualKeyEsp"]; ok {
				nestedManualKey.Esp.Misc = o.Misc["ManualKeyEsp"]
			}
			if o.ManualKey.Esp.Authentication != nil {
				nestedManualKey.Esp.Authentication = &ManualKeyEspAuthenticationXml_11_0_2{}
				if _, ok := o.Misc["ManualKeyEspAuthentication"]; ok {
					nestedManualKey.Esp.Authentication.Misc = o.Misc["ManualKeyEspAuthentication"]
				}
				if o.ManualKey.Esp.Authentication.Sha1 != nil {
					nestedManualKey.Esp.Authentication.Sha1 = &ManualKeyEspAuthenticationSha1Xml_11_0_2{}
					if _, ok := o.Misc["ManualKeyEspAuthenticationSha1"]; ok {
						nestedManualKey.Esp.Authentication.Sha1.Misc = o.Misc["ManualKeyEspAuthenticationSha1"]
					}
					if o.ManualKey.Esp.Authentication.Sha1.Key != nil {
						nestedManualKey.Esp.Authentication.Sha1.Key = o.ManualKey.Esp.Authentication.Sha1.Key
					}
				}
				if o.ManualKey.Esp.Authentication.Sha256 != nil {
					nestedManualKey.Esp.Authentication.Sha256 = &ManualKeyEspAuthenticationSha256Xml_11_0_2{}
					if _, ok := o.Misc["ManualKeyEspAuthenticationSha256"]; ok {
						nestedManualKey.Esp.Authentication.Sha256.Misc = o.Misc["ManualKeyEspAuthenticationSha256"]
					}
					if o.ManualKey.Esp.Authentication.Sha256.Key != nil {
						nestedManualKey.Esp.Authentication.Sha256.Key = o.ManualKey.Esp.Authentication.Sha256.Key
					}
				}
				if o.ManualKey.Esp.Authentication.Sha384 != nil {
					nestedManualKey.Esp.Authentication.Sha384 = &ManualKeyEspAuthenticationSha384Xml_11_0_2{}
					if _, ok := o.Misc["ManualKeyEspAuthenticationSha384"]; ok {
						nestedManualKey.Esp.Authentication.Sha384.Misc = o.Misc["ManualKeyEspAuthenticationSha384"]
					}
					if o.ManualKey.Esp.Authentication.Sha384.Key != nil {
						nestedManualKey.Esp.Authentication.Sha384.Key = o.ManualKey.Esp.Authentication.Sha384.Key
					}
				}
				if o.ManualKey.Esp.Authentication.Sha512 != nil {
					nestedManualKey.Esp.Authentication.Sha512 = &ManualKeyEspAuthenticationSha512Xml_11_0_2{}
					if _, ok := o.Misc["ManualKeyEspAuthenticationSha512"]; ok {
						nestedManualKey.Esp.Authentication.Sha512.Misc = o.Misc["ManualKeyEspAuthenticationSha512"]
					}
					if o.ManualKey.Esp.Authentication.Sha512.Key != nil {
						nestedManualKey.Esp.Authentication.Sha512.Key = o.ManualKey.Esp.Authentication.Sha512.Key
					}
				}
				if o.ManualKey.Esp.Authentication.Md5 != nil {
					nestedManualKey.Esp.Authentication.Md5 = &ManualKeyEspAuthenticationMd5Xml_11_0_2{}
					if _, ok := o.Misc["ManualKeyEspAuthenticationMd5"]; ok {
						nestedManualKey.Esp.Authentication.Md5.Misc = o.Misc["ManualKeyEspAuthenticationMd5"]
					}
					if o.ManualKey.Esp.Authentication.Md5.Key != nil {
						nestedManualKey.Esp.Authentication.Md5.Key = o.ManualKey.Esp.Authentication.Md5.Key
					}
				}
				if o.ManualKey.Esp.Authentication.None != nil {
					nestedManualKey.Esp.Authentication.None = &ManualKeyEspAuthenticationNoneXml_11_0_2{}
					if _, ok := o.Misc["ManualKeyEspAuthenticationNone"]; ok {
						nestedManualKey.Esp.Authentication.None.Misc = o.Misc["ManualKeyEspAuthenticationNone"]
					}
				}
			}
			if o.ManualKey.Esp.Encryption != nil {
				nestedManualKey.Esp.Encryption = &ManualKeyEspEncryptionXml_11_0_2{}
				if _, ok := o.Misc["ManualKeyEspEncryption"]; ok {
					nestedManualKey.Esp.Encryption.Misc = o.Misc["ManualKeyEspEncryption"]
				}
				if o.ManualKey.Esp.Encryption.Algorithm != nil {
					nestedManualKey.Esp.Encryption.Algorithm = o.ManualKey.Esp.Encryption.Algorithm
				}
				if o.ManualKey.Esp.Encryption.Key != nil {
					nestedManualKey.Esp.Encryption.Key = o.ManualKey.Esp.Encryption.Key
				}
			}
		}
	}
	entry.ManualKey = nestedManualKey

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
		entry.AntiReplay = util.AsBool(o.AntiReplay, nil)
		entry.AntiReplayWindow = o.AntiReplayWindow
		entry.Comment = o.Comment
		entry.CopyFlowLabel = util.AsBool(o.CopyFlowLabel, nil)
		entry.CopyTos = util.AsBool(o.CopyTos, nil)
		entry.Disabled = util.AsBool(o.Disabled, nil)
		entry.EnableGreEncapsulation = util.AsBool(o.EnableGreEncapsulation, nil)
		entry.IpsecMode = o.IpsecMode
		entry.Ipv6 = util.AsBool(o.Ipv6, nil)
		entry.TunnelInterface = o.TunnelInterface
		var nestedTunnelMonitor *TunnelMonitor
		if o.TunnelMonitor != nil {
			nestedTunnelMonitor = &TunnelMonitor{}
			if o.TunnelMonitor.Misc != nil {
				entry.Misc["TunnelMonitor"] = o.TunnelMonitor.Misc
			}
			if o.TunnelMonitor.DestinationIp != nil {
				nestedTunnelMonitor.DestinationIp = o.TunnelMonitor.DestinationIp
			}
			if o.TunnelMonitor.Enable != nil {
				nestedTunnelMonitor.Enable = util.AsBool(o.TunnelMonitor.Enable, nil)
			}
			if o.TunnelMonitor.ProxyId != nil {
				nestedTunnelMonitor.ProxyId = o.TunnelMonitor.ProxyId
			}
			if o.TunnelMonitor.TunnelMonitorProfile != nil {
				nestedTunnelMonitor.TunnelMonitorProfile = o.TunnelMonitor.TunnelMonitorProfile
			}
		}
		entry.TunnelMonitor = nestedTunnelMonitor

		var nestedAutoKey *AutoKey
		if o.AutoKey != nil {
			nestedAutoKey = &AutoKey{}
			if o.AutoKey.Misc != nil {
				entry.Misc["AutoKey"] = o.AutoKey.Misc
			}
			if o.AutoKey.IkeGateway != nil {
				nestedAutoKey.IkeGateway = []AutoKeyIkeGateway{}
				for _, oAutoKeyIkeGateway := range o.AutoKey.IkeGateway {
					nestedAutoKeyIkeGateway := AutoKeyIkeGateway{}
					if oAutoKeyIkeGateway.Misc != nil {
						entry.Misc["AutoKeyIkeGateway"] = oAutoKeyIkeGateway.Misc
					}
					if oAutoKeyIkeGateway.Name != "" {
						nestedAutoKeyIkeGateway.Name = oAutoKeyIkeGateway.Name
					}
					nestedAutoKey.IkeGateway = append(nestedAutoKey.IkeGateway, nestedAutoKeyIkeGateway)
				}
			}
			if o.AutoKey.IpsecCryptoProfile != nil {
				nestedAutoKey.IpsecCryptoProfile = o.AutoKey.IpsecCryptoProfile
			}
			if o.AutoKey.ProxyId != nil {
				nestedAutoKey.ProxyId = []AutoKeyProxyId{}
				for _, oAutoKeyProxyId := range o.AutoKey.ProxyId {
					nestedAutoKeyProxyId := AutoKeyProxyId{}
					if oAutoKeyProxyId.Misc != nil {
						entry.Misc["AutoKeyProxyId"] = oAutoKeyProxyId.Misc
					}
					if oAutoKeyProxyId.Local != nil {
						nestedAutoKeyProxyId.Local = oAutoKeyProxyId.Local
					}
					if oAutoKeyProxyId.Remote != nil {
						nestedAutoKeyProxyId.Remote = oAutoKeyProxyId.Remote
					}
					if oAutoKeyProxyId.Protocol != nil {
						nestedAutoKeyProxyId.Protocol = &AutoKeyProxyIdProtocol{}
						if oAutoKeyProxyId.Protocol.Misc != nil {
							entry.Misc["AutoKeyProxyIdProtocol"] = oAutoKeyProxyId.Protocol.Misc
						}
						if oAutoKeyProxyId.Protocol.Number != nil {
							nestedAutoKeyProxyId.Protocol.Number = oAutoKeyProxyId.Protocol.Number
						}
						if oAutoKeyProxyId.Protocol.Any != nil {
							nestedAutoKeyProxyId.Protocol.Any = &AutoKeyProxyIdProtocolAny{}
							if oAutoKeyProxyId.Protocol.Any.Misc != nil {
								entry.Misc["AutoKeyProxyIdProtocolAny"] = oAutoKeyProxyId.Protocol.Any.Misc
							}
						}
						if oAutoKeyProxyId.Protocol.Tcp != nil {
							nestedAutoKeyProxyId.Protocol.Tcp = &AutoKeyProxyIdProtocolTcp{}
							if oAutoKeyProxyId.Protocol.Tcp.Misc != nil {
								entry.Misc["AutoKeyProxyIdProtocolTcp"] = oAutoKeyProxyId.Protocol.Tcp.Misc
							}
							if oAutoKeyProxyId.Protocol.Tcp.LocalPort != nil {
								nestedAutoKeyProxyId.Protocol.Tcp.LocalPort = oAutoKeyProxyId.Protocol.Tcp.LocalPort
							}
							if oAutoKeyProxyId.Protocol.Tcp.RemotePort != nil {
								nestedAutoKeyProxyId.Protocol.Tcp.RemotePort = oAutoKeyProxyId.Protocol.Tcp.RemotePort
							}
						}
						if oAutoKeyProxyId.Protocol.Udp != nil {
							nestedAutoKeyProxyId.Protocol.Udp = &AutoKeyProxyIdProtocolUdp{}
							if oAutoKeyProxyId.Protocol.Udp.Misc != nil {
								entry.Misc["AutoKeyProxyIdProtocolUdp"] = oAutoKeyProxyId.Protocol.Udp.Misc
							}
							if oAutoKeyProxyId.Protocol.Udp.RemotePort != nil {
								nestedAutoKeyProxyId.Protocol.Udp.RemotePort = oAutoKeyProxyId.Protocol.Udp.RemotePort
							}
							if oAutoKeyProxyId.Protocol.Udp.LocalPort != nil {
								nestedAutoKeyProxyId.Protocol.Udp.LocalPort = oAutoKeyProxyId.Protocol.Udp.LocalPort
							}
						}
					}
					if oAutoKeyProxyId.Name != "" {
						nestedAutoKeyProxyId.Name = oAutoKeyProxyId.Name
					}
					nestedAutoKey.ProxyId = append(nestedAutoKey.ProxyId, nestedAutoKeyProxyId)
				}
			}
			if o.AutoKey.ProxyIdV6 != nil {
				nestedAutoKey.ProxyIdV6 = []AutoKeyProxyIdV6{}
				for _, oAutoKeyProxyIdV6 := range o.AutoKey.ProxyIdV6 {
					nestedAutoKeyProxyIdV6 := AutoKeyProxyIdV6{}
					if oAutoKeyProxyIdV6.Misc != nil {
						entry.Misc["AutoKeyProxyIdV6"] = oAutoKeyProxyIdV6.Misc
					}
					if oAutoKeyProxyIdV6.Local != nil {
						nestedAutoKeyProxyIdV6.Local = oAutoKeyProxyIdV6.Local
					}
					if oAutoKeyProxyIdV6.Remote != nil {
						nestedAutoKeyProxyIdV6.Remote = oAutoKeyProxyIdV6.Remote
					}
					if oAutoKeyProxyIdV6.Protocol != nil {
						nestedAutoKeyProxyIdV6.Protocol = &AutoKeyProxyIdV6Protocol{}
						if oAutoKeyProxyIdV6.Protocol.Misc != nil {
							entry.Misc["AutoKeyProxyIdV6Protocol"] = oAutoKeyProxyIdV6.Protocol.Misc
						}
						if oAutoKeyProxyIdV6.Protocol.Udp != nil {
							nestedAutoKeyProxyIdV6.Protocol.Udp = &AutoKeyProxyIdV6ProtocolUdp{}
							if oAutoKeyProxyIdV6.Protocol.Udp.Misc != nil {
								entry.Misc["AutoKeyProxyIdV6ProtocolUdp"] = oAutoKeyProxyIdV6.Protocol.Udp.Misc
							}
							if oAutoKeyProxyIdV6.Protocol.Udp.LocalPort != nil {
								nestedAutoKeyProxyIdV6.Protocol.Udp.LocalPort = oAutoKeyProxyIdV6.Protocol.Udp.LocalPort
							}
							if oAutoKeyProxyIdV6.Protocol.Udp.RemotePort != nil {
								nestedAutoKeyProxyIdV6.Protocol.Udp.RemotePort = oAutoKeyProxyIdV6.Protocol.Udp.RemotePort
							}
						}
						if oAutoKeyProxyIdV6.Protocol.Number != nil {
							nestedAutoKeyProxyIdV6.Protocol.Number = oAutoKeyProxyIdV6.Protocol.Number
						}
						if oAutoKeyProxyIdV6.Protocol.Any != nil {
							nestedAutoKeyProxyIdV6.Protocol.Any = &AutoKeyProxyIdV6ProtocolAny{}
							if oAutoKeyProxyIdV6.Protocol.Any.Misc != nil {
								entry.Misc["AutoKeyProxyIdV6ProtocolAny"] = oAutoKeyProxyIdV6.Protocol.Any.Misc
							}
						}
						if oAutoKeyProxyIdV6.Protocol.Tcp != nil {
							nestedAutoKeyProxyIdV6.Protocol.Tcp = &AutoKeyProxyIdV6ProtocolTcp{}
							if oAutoKeyProxyIdV6.Protocol.Tcp.Misc != nil {
								entry.Misc["AutoKeyProxyIdV6ProtocolTcp"] = oAutoKeyProxyIdV6.Protocol.Tcp.Misc
							}
							if oAutoKeyProxyIdV6.Protocol.Tcp.LocalPort != nil {
								nestedAutoKeyProxyIdV6.Protocol.Tcp.LocalPort = oAutoKeyProxyIdV6.Protocol.Tcp.LocalPort
							}
							if oAutoKeyProxyIdV6.Protocol.Tcp.RemotePort != nil {
								nestedAutoKeyProxyIdV6.Protocol.Tcp.RemotePort = oAutoKeyProxyIdV6.Protocol.Tcp.RemotePort
							}
						}
					}
					if oAutoKeyProxyIdV6.Name != "" {
						nestedAutoKeyProxyIdV6.Name = oAutoKeyProxyIdV6.Name
					}
					nestedAutoKey.ProxyIdV6 = append(nestedAutoKey.ProxyIdV6, nestedAutoKeyProxyIdV6)
				}
			}
		}
		entry.AutoKey = nestedAutoKey

		var nestedGlobalProtectSatellite *GlobalProtectSatellite
		if o.GlobalProtectSatellite != nil {
			nestedGlobalProtectSatellite = &GlobalProtectSatellite{}
			if o.GlobalProtectSatellite.Misc != nil {
				entry.Misc["GlobalProtectSatellite"] = o.GlobalProtectSatellite.Misc
			}
			if o.GlobalProtectSatellite.PublishRoutes != nil {
				nestedGlobalProtectSatellite.PublishRoutes = util.MemToStr(o.GlobalProtectSatellite.PublishRoutes)
			}
			if o.GlobalProtectSatellite.ExternalCa != nil {
				nestedGlobalProtectSatellite.ExternalCa = &GlobalProtectSatelliteExternalCa{}
				if o.GlobalProtectSatellite.ExternalCa.Misc != nil {
					entry.Misc["GlobalProtectSatelliteExternalCa"] = o.GlobalProtectSatellite.ExternalCa.Misc
				}
				if o.GlobalProtectSatellite.ExternalCa.CertificateProfile != nil {
					nestedGlobalProtectSatellite.ExternalCa.CertificateProfile = o.GlobalProtectSatellite.ExternalCa.CertificateProfile
				}
				if o.GlobalProtectSatellite.ExternalCa.LocalCertificate != nil {
					nestedGlobalProtectSatellite.ExternalCa.LocalCertificate = o.GlobalProtectSatellite.ExternalCa.LocalCertificate
				}
			}
			if o.GlobalProtectSatellite.Ipv6Preferred != nil {
				nestedGlobalProtectSatellite.Ipv6Preferred = util.AsBool(o.GlobalProtectSatellite.Ipv6Preferred, nil)
			}
			if o.GlobalProtectSatellite.LocalAddress != nil {
				nestedGlobalProtectSatellite.LocalAddress = &GlobalProtectSatelliteLocalAddress{}
				if o.GlobalProtectSatellite.LocalAddress.Misc != nil {
					entry.Misc["GlobalProtectSatelliteLocalAddress"] = o.GlobalProtectSatellite.LocalAddress.Misc
				}
				if o.GlobalProtectSatellite.LocalAddress.Interface != nil {
					nestedGlobalProtectSatellite.LocalAddress.Interface = o.GlobalProtectSatellite.LocalAddress.Interface
				}
				if o.GlobalProtectSatellite.LocalAddress.FloatingIp != nil {
					nestedGlobalProtectSatellite.LocalAddress.FloatingIp = &GlobalProtectSatelliteLocalAddressFloatingIp{}
					if o.GlobalProtectSatellite.LocalAddress.FloatingIp.Misc != nil {
						entry.Misc["GlobalProtectSatelliteLocalAddressFloatingIp"] = o.GlobalProtectSatellite.LocalAddress.FloatingIp.Misc
					}
					if o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv6 != nil {
						nestedGlobalProtectSatellite.LocalAddress.FloatingIp.Ipv6 = o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv6
					}
					if o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv4 != nil {
						nestedGlobalProtectSatellite.LocalAddress.FloatingIp.Ipv4 = o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv4
					}
				}
				if o.GlobalProtectSatellite.LocalAddress.Ip != nil {
					nestedGlobalProtectSatellite.LocalAddress.Ip = &GlobalProtectSatelliteLocalAddressIp{}
					if o.GlobalProtectSatellite.LocalAddress.Ip.Misc != nil {
						entry.Misc["GlobalProtectSatelliteLocalAddressIp"] = o.GlobalProtectSatellite.LocalAddress.Ip.Misc
					}
					if o.GlobalProtectSatellite.LocalAddress.Ip.Ipv4 != nil {
						nestedGlobalProtectSatellite.LocalAddress.Ip.Ipv4 = o.GlobalProtectSatellite.LocalAddress.Ip.Ipv4
					}
					if o.GlobalProtectSatellite.LocalAddress.Ip.Ipv6 != nil {
						nestedGlobalProtectSatellite.LocalAddress.Ip.Ipv6 = o.GlobalProtectSatellite.LocalAddress.Ip.Ipv6
					}
				}
			}
			if o.GlobalProtectSatellite.PortalAddress != nil {
				nestedGlobalProtectSatellite.PortalAddress = o.GlobalProtectSatellite.PortalAddress
			}
			if o.GlobalProtectSatellite.PublishConnectedRoutes != nil {
				nestedGlobalProtectSatellite.PublishConnectedRoutes = &GlobalProtectSatellitePublishConnectedRoutes{}
				if o.GlobalProtectSatellite.PublishConnectedRoutes.Misc != nil {
					entry.Misc["GlobalProtectSatellitePublishConnectedRoutes"] = o.GlobalProtectSatellite.PublishConnectedRoutes.Misc
				}
				if o.GlobalProtectSatellite.PublishConnectedRoutes.Enable != nil {
					nestedGlobalProtectSatellite.PublishConnectedRoutes.Enable = util.AsBool(o.GlobalProtectSatellite.PublishConnectedRoutes.Enable, nil)
				}
			}
		}
		entry.GlobalProtectSatellite = nestedGlobalProtectSatellite

		var nestedManualKey *ManualKey
		if o.ManualKey != nil {
			nestedManualKey = &ManualKey{}
			if o.ManualKey.Misc != nil {
				entry.Misc["ManualKey"] = o.ManualKey.Misc
			}
			if o.ManualKey.LocalAddress != nil {
				nestedManualKey.LocalAddress = &ManualKeyLocalAddress{}
				if o.ManualKey.LocalAddress.Misc != nil {
					entry.Misc["ManualKeyLocalAddress"] = o.ManualKey.LocalAddress.Misc
				}
				if o.ManualKey.LocalAddress.Interface != nil {
					nestedManualKey.LocalAddress.Interface = o.ManualKey.LocalAddress.Interface
				}
				if o.ManualKey.LocalAddress.FloatingIp != nil {
					nestedManualKey.LocalAddress.FloatingIp = o.ManualKey.LocalAddress.FloatingIp
				}
				if o.ManualKey.LocalAddress.Ip != nil {
					nestedManualKey.LocalAddress.Ip = o.ManualKey.LocalAddress.Ip
				}
			}
			if o.ManualKey.LocalSpi != nil {
				nestedManualKey.LocalSpi = o.ManualKey.LocalSpi
			}
			if o.ManualKey.PeerAddress != nil {
				nestedManualKey.PeerAddress = &ManualKeyPeerAddress{}
				if o.ManualKey.PeerAddress.Misc != nil {
					entry.Misc["ManualKeyPeerAddress"] = o.ManualKey.PeerAddress.Misc
				}
				if o.ManualKey.PeerAddress.Ip != nil {
					nestedManualKey.PeerAddress.Ip = o.ManualKey.PeerAddress.Ip
				}
			}
			if o.ManualKey.RemoteSpi != nil {
				nestedManualKey.RemoteSpi = o.ManualKey.RemoteSpi
			}
			if o.ManualKey.Ah != nil {
				nestedManualKey.Ah = &ManualKeyAh{}
				if o.ManualKey.Ah.Misc != nil {
					entry.Misc["ManualKeyAh"] = o.ManualKey.Ah.Misc
				}
				if o.ManualKey.Ah.Md5 != nil {
					nestedManualKey.Ah.Md5 = &ManualKeyAhMd5{}
					if o.ManualKey.Ah.Md5.Misc != nil {
						entry.Misc["ManualKeyAhMd5"] = o.ManualKey.Ah.Md5.Misc
					}
					if o.ManualKey.Ah.Md5.Key != nil {
						nestedManualKey.Ah.Md5.Key = o.ManualKey.Ah.Md5.Key
					}
				}
				if o.ManualKey.Ah.Sha1 != nil {
					nestedManualKey.Ah.Sha1 = &ManualKeyAhSha1{}
					if o.ManualKey.Ah.Sha1.Misc != nil {
						entry.Misc["ManualKeyAhSha1"] = o.ManualKey.Ah.Sha1.Misc
					}
					if o.ManualKey.Ah.Sha1.Key != nil {
						nestedManualKey.Ah.Sha1.Key = o.ManualKey.Ah.Sha1.Key
					}
				}
				if o.ManualKey.Ah.Sha256 != nil {
					nestedManualKey.Ah.Sha256 = &ManualKeyAhSha256{}
					if o.ManualKey.Ah.Sha256.Misc != nil {
						entry.Misc["ManualKeyAhSha256"] = o.ManualKey.Ah.Sha256.Misc
					}
					if o.ManualKey.Ah.Sha256.Key != nil {
						nestedManualKey.Ah.Sha256.Key = o.ManualKey.Ah.Sha256.Key
					}
				}
				if o.ManualKey.Ah.Sha384 != nil {
					nestedManualKey.Ah.Sha384 = &ManualKeyAhSha384{}
					if o.ManualKey.Ah.Sha384.Misc != nil {
						entry.Misc["ManualKeyAhSha384"] = o.ManualKey.Ah.Sha384.Misc
					}
					if o.ManualKey.Ah.Sha384.Key != nil {
						nestedManualKey.Ah.Sha384.Key = o.ManualKey.Ah.Sha384.Key
					}
				}
				if o.ManualKey.Ah.Sha512 != nil {
					nestedManualKey.Ah.Sha512 = &ManualKeyAhSha512{}
					if o.ManualKey.Ah.Sha512.Misc != nil {
						entry.Misc["ManualKeyAhSha512"] = o.ManualKey.Ah.Sha512.Misc
					}
					if o.ManualKey.Ah.Sha512.Key != nil {
						nestedManualKey.Ah.Sha512.Key = o.ManualKey.Ah.Sha512.Key
					}
				}
			}
			if o.ManualKey.Esp != nil {
				nestedManualKey.Esp = &ManualKeyEsp{}
				if o.ManualKey.Esp.Misc != nil {
					entry.Misc["ManualKeyEsp"] = o.ManualKey.Esp.Misc
				}
				if o.ManualKey.Esp.Authentication != nil {
					nestedManualKey.Esp.Authentication = &ManualKeyEspAuthentication{}
					if o.ManualKey.Esp.Authentication.Misc != nil {
						entry.Misc["ManualKeyEspAuthentication"] = o.ManualKey.Esp.Authentication.Misc
					}
					if o.ManualKey.Esp.Authentication.Sha512 != nil {
						nestedManualKey.Esp.Authentication.Sha512 = &ManualKeyEspAuthenticationSha512{}
						if o.ManualKey.Esp.Authentication.Sha512.Misc != nil {
							entry.Misc["ManualKeyEspAuthenticationSha512"] = o.ManualKey.Esp.Authentication.Sha512.Misc
						}
						if o.ManualKey.Esp.Authentication.Sha512.Key != nil {
							nestedManualKey.Esp.Authentication.Sha512.Key = o.ManualKey.Esp.Authentication.Sha512.Key
						}
					}
					if o.ManualKey.Esp.Authentication.Md5 != nil {
						nestedManualKey.Esp.Authentication.Md5 = &ManualKeyEspAuthenticationMd5{}
						if o.ManualKey.Esp.Authentication.Md5.Misc != nil {
							entry.Misc["ManualKeyEspAuthenticationMd5"] = o.ManualKey.Esp.Authentication.Md5.Misc
						}
						if o.ManualKey.Esp.Authentication.Md5.Key != nil {
							nestedManualKey.Esp.Authentication.Md5.Key = o.ManualKey.Esp.Authentication.Md5.Key
						}
					}
					if o.ManualKey.Esp.Authentication.None != nil {
						nestedManualKey.Esp.Authentication.None = &ManualKeyEspAuthenticationNone{}
						if o.ManualKey.Esp.Authentication.None.Misc != nil {
							entry.Misc["ManualKeyEspAuthenticationNone"] = o.ManualKey.Esp.Authentication.None.Misc
						}
					}
					if o.ManualKey.Esp.Authentication.Sha1 != nil {
						nestedManualKey.Esp.Authentication.Sha1 = &ManualKeyEspAuthenticationSha1{}
						if o.ManualKey.Esp.Authentication.Sha1.Misc != nil {
							entry.Misc["ManualKeyEspAuthenticationSha1"] = o.ManualKey.Esp.Authentication.Sha1.Misc
						}
						if o.ManualKey.Esp.Authentication.Sha1.Key != nil {
							nestedManualKey.Esp.Authentication.Sha1.Key = o.ManualKey.Esp.Authentication.Sha1.Key
						}
					}
					if o.ManualKey.Esp.Authentication.Sha256 != nil {
						nestedManualKey.Esp.Authentication.Sha256 = &ManualKeyEspAuthenticationSha256{}
						if o.ManualKey.Esp.Authentication.Sha256.Misc != nil {
							entry.Misc["ManualKeyEspAuthenticationSha256"] = o.ManualKey.Esp.Authentication.Sha256.Misc
						}
						if o.ManualKey.Esp.Authentication.Sha256.Key != nil {
							nestedManualKey.Esp.Authentication.Sha256.Key = o.ManualKey.Esp.Authentication.Sha256.Key
						}
					}
					if o.ManualKey.Esp.Authentication.Sha384 != nil {
						nestedManualKey.Esp.Authentication.Sha384 = &ManualKeyEspAuthenticationSha384{}
						if o.ManualKey.Esp.Authentication.Sha384.Misc != nil {
							entry.Misc["ManualKeyEspAuthenticationSha384"] = o.ManualKey.Esp.Authentication.Sha384.Misc
						}
						if o.ManualKey.Esp.Authentication.Sha384.Key != nil {
							nestedManualKey.Esp.Authentication.Sha384.Key = o.ManualKey.Esp.Authentication.Sha384.Key
						}
					}
				}
				if o.ManualKey.Esp.Encryption != nil {
					nestedManualKey.Esp.Encryption = &ManualKeyEspEncryption{}
					if o.ManualKey.Esp.Encryption.Misc != nil {
						entry.Misc["ManualKeyEspEncryption"] = o.ManualKey.Esp.Encryption.Misc
					}
					if o.ManualKey.Esp.Encryption.Algorithm != nil {
						nestedManualKey.Esp.Encryption.Algorithm = o.ManualKey.Esp.Encryption.Algorithm
					}
					if o.ManualKey.Esp.Encryption.Key != nil {
						nestedManualKey.Esp.Encryption.Key = o.ManualKey.Esp.Encryption.Key
					}
				}
			}
		}
		entry.ManualKey = nestedManualKey

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
		entry.AntiReplay = util.AsBool(o.AntiReplay, nil)
		entry.AntiReplayWindow = o.AntiReplayWindow
		entry.Comment = o.Comment
		entry.CopyFlowLabel = util.AsBool(o.CopyFlowLabel, nil)
		entry.CopyTos = util.AsBool(o.CopyTos, nil)
		entry.Disabled = util.AsBool(o.Disabled, nil)
		entry.EnableGreEncapsulation = util.AsBool(o.EnableGreEncapsulation, nil)
		entry.IpsecMode = o.IpsecMode
		entry.Ipv6 = util.AsBool(o.Ipv6, nil)
		entry.TunnelInterface = o.TunnelInterface
		var nestedTunnelMonitor *TunnelMonitor
		if o.TunnelMonitor != nil {
			nestedTunnelMonitor = &TunnelMonitor{}
			if o.TunnelMonitor.Misc != nil {
				entry.Misc["TunnelMonitor"] = o.TunnelMonitor.Misc
			}
			if o.TunnelMonitor.DestinationIp != nil {
				nestedTunnelMonitor.DestinationIp = o.TunnelMonitor.DestinationIp
			}
			if o.TunnelMonitor.Enable != nil {
				nestedTunnelMonitor.Enable = util.AsBool(o.TunnelMonitor.Enable, nil)
			}
			if o.TunnelMonitor.ProxyId != nil {
				nestedTunnelMonitor.ProxyId = o.TunnelMonitor.ProxyId
			}
			if o.TunnelMonitor.TunnelMonitorProfile != nil {
				nestedTunnelMonitor.TunnelMonitorProfile = o.TunnelMonitor.TunnelMonitorProfile
			}
		}
		entry.TunnelMonitor = nestedTunnelMonitor

		var nestedAutoKey *AutoKey
		if o.AutoKey != nil {
			nestedAutoKey = &AutoKey{}
			if o.AutoKey.Misc != nil {
				entry.Misc["AutoKey"] = o.AutoKey.Misc
			}
			if o.AutoKey.IkeGateway != nil {
				nestedAutoKey.IkeGateway = []AutoKeyIkeGateway{}
				for _, oAutoKeyIkeGateway := range o.AutoKey.IkeGateway {
					nestedAutoKeyIkeGateway := AutoKeyIkeGateway{}
					if oAutoKeyIkeGateway.Misc != nil {
						entry.Misc["AutoKeyIkeGateway"] = oAutoKeyIkeGateway.Misc
					}
					if oAutoKeyIkeGateway.Name != "" {
						nestedAutoKeyIkeGateway.Name = oAutoKeyIkeGateway.Name
					}
					nestedAutoKey.IkeGateway = append(nestedAutoKey.IkeGateway, nestedAutoKeyIkeGateway)
				}
			}
			if o.AutoKey.IpsecCryptoProfile != nil {
				nestedAutoKey.IpsecCryptoProfile = o.AutoKey.IpsecCryptoProfile
			}
			if o.AutoKey.ProxyId != nil {
				nestedAutoKey.ProxyId = []AutoKeyProxyId{}
				for _, oAutoKeyProxyId := range o.AutoKey.ProxyId {
					nestedAutoKeyProxyId := AutoKeyProxyId{}
					if oAutoKeyProxyId.Misc != nil {
						entry.Misc["AutoKeyProxyId"] = oAutoKeyProxyId.Misc
					}
					if oAutoKeyProxyId.Local != nil {
						nestedAutoKeyProxyId.Local = oAutoKeyProxyId.Local
					}
					if oAutoKeyProxyId.Remote != nil {
						nestedAutoKeyProxyId.Remote = oAutoKeyProxyId.Remote
					}
					if oAutoKeyProxyId.Protocol != nil {
						nestedAutoKeyProxyId.Protocol = &AutoKeyProxyIdProtocol{}
						if oAutoKeyProxyId.Protocol.Misc != nil {
							entry.Misc["AutoKeyProxyIdProtocol"] = oAutoKeyProxyId.Protocol.Misc
						}
						if oAutoKeyProxyId.Protocol.Any != nil {
							nestedAutoKeyProxyId.Protocol.Any = &AutoKeyProxyIdProtocolAny{}
							if oAutoKeyProxyId.Protocol.Any.Misc != nil {
								entry.Misc["AutoKeyProxyIdProtocolAny"] = oAutoKeyProxyId.Protocol.Any.Misc
							}
						}
						if oAutoKeyProxyId.Protocol.Tcp != nil {
							nestedAutoKeyProxyId.Protocol.Tcp = &AutoKeyProxyIdProtocolTcp{}
							if oAutoKeyProxyId.Protocol.Tcp.Misc != nil {
								entry.Misc["AutoKeyProxyIdProtocolTcp"] = oAutoKeyProxyId.Protocol.Tcp.Misc
							}
							if oAutoKeyProxyId.Protocol.Tcp.LocalPort != nil {
								nestedAutoKeyProxyId.Protocol.Tcp.LocalPort = oAutoKeyProxyId.Protocol.Tcp.LocalPort
							}
							if oAutoKeyProxyId.Protocol.Tcp.RemotePort != nil {
								nestedAutoKeyProxyId.Protocol.Tcp.RemotePort = oAutoKeyProxyId.Protocol.Tcp.RemotePort
							}
						}
						if oAutoKeyProxyId.Protocol.Udp != nil {
							nestedAutoKeyProxyId.Protocol.Udp = &AutoKeyProxyIdProtocolUdp{}
							if oAutoKeyProxyId.Protocol.Udp.Misc != nil {
								entry.Misc["AutoKeyProxyIdProtocolUdp"] = oAutoKeyProxyId.Protocol.Udp.Misc
							}
							if oAutoKeyProxyId.Protocol.Udp.LocalPort != nil {
								nestedAutoKeyProxyId.Protocol.Udp.LocalPort = oAutoKeyProxyId.Protocol.Udp.LocalPort
							}
							if oAutoKeyProxyId.Protocol.Udp.RemotePort != nil {
								nestedAutoKeyProxyId.Protocol.Udp.RemotePort = oAutoKeyProxyId.Protocol.Udp.RemotePort
							}
						}
						if oAutoKeyProxyId.Protocol.Number != nil {
							nestedAutoKeyProxyId.Protocol.Number = oAutoKeyProxyId.Protocol.Number
						}
					}
					if oAutoKeyProxyId.Name != "" {
						nestedAutoKeyProxyId.Name = oAutoKeyProxyId.Name
					}
					nestedAutoKey.ProxyId = append(nestedAutoKey.ProxyId, nestedAutoKeyProxyId)
				}
			}
			if o.AutoKey.ProxyIdV6 != nil {
				nestedAutoKey.ProxyIdV6 = []AutoKeyProxyIdV6{}
				for _, oAutoKeyProxyIdV6 := range o.AutoKey.ProxyIdV6 {
					nestedAutoKeyProxyIdV6 := AutoKeyProxyIdV6{}
					if oAutoKeyProxyIdV6.Misc != nil {
						entry.Misc["AutoKeyProxyIdV6"] = oAutoKeyProxyIdV6.Misc
					}
					if oAutoKeyProxyIdV6.Remote != nil {
						nestedAutoKeyProxyIdV6.Remote = oAutoKeyProxyIdV6.Remote
					}
					if oAutoKeyProxyIdV6.Protocol != nil {
						nestedAutoKeyProxyIdV6.Protocol = &AutoKeyProxyIdV6Protocol{}
						if oAutoKeyProxyIdV6.Protocol.Misc != nil {
							entry.Misc["AutoKeyProxyIdV6Protocol"] = oAutoKeyProxyIdV6.Protocol.Misc
						}
						if oAutoKeyProxyIdV6.Protocol.Number != nil {
							nestedAutoKeyProxyIdV6.Protocol.Number = oAutoKeyProxyIdV6.Protocol.Number
						}
						if oAutoKeyProxyIdV6.Protocol.Any != nil {
							nestedAutoKeyProxyIdV6.Protocol.Any = &AutoKeyProxyIdV6ProtocolAny{}
							if oAutoKeyProxyIdV6.Protocol.Any.Misc != nil {
								entry.Misc["AutoKeyProxyIdV6ProtocolAny"] = oAutoKeyProxyIdV6.Protocol.Any.Misc
							}
						}
						if oAutoKeyProxyIdV6.Protocol.Tcp != nil {
							nestedAutoKeyProxyIdV6.Protocol.Tcp = &AutoKeyProxyIdV6ProtocolTcp{}
							if oAutoKeyProxyIdV6.Protocol.Tcp.Misc != nil {
								entry.Misc["AutoKeyProxyIdV6ProtocolTcp"] = oAutoKeyProxyIdV6.Protocol.Tcp.Misc
							}
							if oAutoKeyProxyIdV6.Protocol.Tcp.LocalPort != nil {
								nestedAutoKeyProxyIdV6.Protocol.Tcp.LocalPort = oAutoKeyProxyIdV6.Protocol.Tcp.LocalPort
							}
							if oAutoKeyProxyIdV6.Protocol.Tcp.RemotePort != nil {
								nestedAutoKeyProxyIdV6.Protocol.Tcp.RemotePort = oAutoKeyProxyIdV6.Protocol.Tcp.RemotePort
							}
						}
						if oAutoKeyProxyIdV6.Protocol.Udp != nil {
							nestedAutoKeyProxyIdV6.Protocol.Udp = &AutoKeyProxyIdV6ProtocolUdp{}
							if oAutoKeyProxyIdV6.Protocol.Udp.Misc != nil {
								entry.Misc["AutoKeyProxyIdV6ProtocolUdp"] = oAutoKeyProxyIdV6.Protocol.Udp.Misc
							}
							if oAutoKeyProxyIdV6.Protocol.Udp.LocalPort != nil {
								nestedAutoKeyProxyIdV6.Protocol.Udp.LocalPort = oAutoKeyProxyIdV6.Protocol.Udp.LocalPort
							}
							if oAutoKeyProxyIdV6.Protocol.Udp.RemotePort != nil {
								nestedAutoKeyProxyIdV6.Protocol.Udp.RemotePort = oAutoKeyProxyIdV6.Protocol.Udp.RemotePort
							}
						}
					}
					if oAutoKeyProxyIdV6.Name != "" {
						nestedAutoKeyProxyIdV6.Name = oAutoKeyProxyIdV6.Name
					}
					if oAutoKeyProxyIdV6.Local != nil {
						nestedAutoKeyProxyIdV6.Local = oAutoKeyProxyIdV6.Local
					}
					nestedAutoKey.ProxyIdV6 = append(nestedAutoKey.ProxyIdV6, nestedAutoKeyProxyIdV6)
				}
			}
		}
		entry.AutoKey = nestedAutoKey

		var nestedGlobalProtectSatellite *GlobalProtectSatellite
		if o.GlobalProtectSatellite != nil {
			nestedGlobalProtectSatellite = &GlobalProtectSatellite{}
			if o.GlobalProtectSatellite.Misc != nil {
				entry.Misc["GlobalProtectSatellite"] = o.GlobalProtectSatellite.Misc
			}
			if o.GlobalProtectSatellite.PortalAddress != nil {
				nestedGlobalProtectSatellite.PortalAddress = o.GlobalProtectSatellite.PortalAddress
			}
			if o.GlobalProtectSatellite.PublishConnectedRoutes != nil {
				nestedGlobalProtectSatellite.PublishConnectedRoutes = &GlobalProtectSatellitePublishConnectedRoutes{}
				if o.GlobalProtectSatellite.PublishConnectedRoutes.Misc != nil {
					entry.Misc["GlobalProtectSatellitePublishConnectedRoutes"] = o.GlobalProtectSatellite.PublishConnectedRoutes.Misc
				}
				if o.GlobalProtectSatellite.PublishConnectedRoutes.Enable != nil {
					nestedGlobalProtectSatellite.PublishConnectedRoutes.Enable = util.AsBool(o.GlobalProtectSatellite.PublishConnectedRoutes.Enable, nil)
				}
			}
			if o.GlobalProtectSatellite.PublishRoutes != nil {
				nestedGlobalProtectSatellite.PublishRoutes = util.MemToStr(o.GlobalProtectSatellite.PublishRoutes)
			}
			if o.GlobalProtectSatellite.ExternalCa != nil {
				nestedGlobalProtectSatellite.ExternalCa = &GlobalProtectSatelliteExternalCa{}
				if o.GlobalProtectSatellite.ExternalCa.Misc != nil {
					entry.Misc["GlobalProtectSatelliteExternalCa"] = o.GlobalProtectSatellite.ExternalCa.Misc
				}
				if o.GlobalProtectSatellite.ExternalCa.CertificateProfile != nil {
					nestedGlobalProtectSatellite.ExternalCa.CertificateProfile = o.GlobalProtectSatellite.ExternalCa.CertificateProfile
				}
				if o.GlobalProtectSatellite.ExternalCa.LocalCertificate != nil {
					nestedGlobalProtectSatellite.ExternalCa.LocalCertificate = o.GlobalProtectSatellite.ExternalCa.LocalCertificate
				}
			}
			if o.GlobalProtectSatellite.Ipv6Preferred != nil {
				nestedGlobalProtectSatellite.Ipv6Preferred = util.AsBool(o.GlobalProtectSatellite.Ipv6Preferred, nil)
			}
			if o.GlobalProtectSatellite.LocalAddress != nil {
				nestedGlobalProtectSatellite.LocalAddress = &GlobalProtectSatelliteLocalAddress{}
				if o.GlobalProtectSatellite.LocalAddress.Misc != nil {
					entry.Misc["GlobalProtectSatelliteLocalAddress"] = o.GlobalProtectSatellite.LocalAddress.Misc
				}
				if o.GlobalProtectSatellite.LocalAddress.Interface != nil {
					nestedGlobalProtectSatellite.LocalAddress.Interface = o.GlobalProtectSatellite.LocalAddress.Interface
				}
				if o.GlobalProtectSatellite.LocalAddress.FloatingIp != nil {
					nestedGlobalProtectSatellite.LocalAddress.FloatingIp = &GlobalProtectSatelliteLocalAddressFloatingIp{}
					if o.GlobalProtectSatellite.LocalAddress.FloatingIp.Misc != nil {
						entry.Misc["GlobalProtectSatelliteLocalAddressFloatingIp"] = o.GlobalProtectSatellite.LocalAddress.FloatingIp.Misc
					}
					if o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv4 != nil {
						nestedGlobalProtectSatellite.LocalAddress.FloatingIp.Ipv4 = o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv4
					}
					if o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv6 != nil {
						nestedGlobalProtectSatellite.LocalAddress.FloatingIp.Ipv6 = o.GlobalProtectSatellite.LocalAddress.FloatingIp.Ipv6
					}
				}
				if o.GlobalProtectSatellite.LocalAddress.Ip != nil {
					nestedGlobalProtectSatellite.LocalAddress.Ip = &GlobalProtectSatelliteLocalAddressIp{}
					if o.GlobalProtectSatellite.LocalAddress.Ip.Misc != nil {
						entry.Misc["GlobalProtectSatelliteLocalAddressIp"] = o.GlobalProtectSatellite.LocalAddress.Ip.Misc
					}
					if o.GlobalProtectSatellite.LocalAddress.Ip.Ipv4 != nil {
						nestedGlobalProtectSatellite.LocalAddress.Ip.Ipv4 = o.GlobalProtectSatellite.LocalAddress.Ip.Ipv4
					}
					if o.GlobalProtectSatellite.LocalAddress.Ip.Ipv6 != nil {
						nestedGlobalProtectSatellite.LocalAddress.Ip.Ipv6 = o.GlobalProtectSatellite.LocalAddress.Ip.Ipv6
					}
				}
			}
		}
		entry.GlobalProtectSatellite = nestedGlobalProtectSatellite

		var nestedManualKey *ManualKey
		if o.ManualKey != nil {
			nestedManualKey = &ManualKey{}
			if o.ManualKey.Misc != nil {
				entry.Misc["ManualKey"] = o.ManualKey.Misc
			}
			if o.ManualKey.LocalAddress != nil {
				nestedManualKey.LocalAddress = &ManualKeyLocalAddress{}
				if o.ManualKey.LocalAddress.Misc != nil {
					entry.Misc["ManualKeyLocalAddress"] = o.ManualKey.LocalAddress.Misc
				}
				if o.ManualKey.LocalAddress.Interface != nil {
					nestedManualKey.LocalAddress.Interface = o.ManualKey.LocalAddress.Interface
				}
				if o.ManualKey.LocalAddress.FloatingIp != nil {
					nestedManualKey.LocalAddress.FloatingIp = o.ManualKey.LocalAddress.FloatingIp
				}
				if o.ManualKey.LocalAddress.Ip != nil {
					nestedManualKey.LocalAddress.Ip = o.ManualKey.LocalAddress.Ip
				}
			}
			if o.ManualKey.LocalSpi != nil {
				nestedManualKey.LocalSpi = o.ManualKey.LocalSpi
			}
			if o.ManualKey.PeerAddress != nil {
				nestedManualKey.PeerAddress = &ManualKeyPeerAddress{}
				if o.ManualKey.PeerAddress.Misc != nil {
					entry.Misc["ManualKeyPeerAddress"] = o.ManualKey.PeerAddress.Misc
				}
				if o.ManualKey.PeerAddress.Ip != nil {
					nestedManualKey.PeerAddress.Ip = o.ManualKey.PeerAddress.Ip
				}
			}
			if o.ManualKey.RemoteSpi != nil {
				nestedManualKey.RemoteSpi = o.ManualKey.RemoteSpi
			}
			if o.ManualKey.Ah != nil {
				nestedManualKey.Ah = &ManualKeyAh{}
				if o.ManualKey.Ah.Misc != nil {
					entry.Misc["ManualKeyAh"] = o.ManualKey.Ah.Misc
				}
				if o.ManualKey.Ah.Md5 != nil {
					nestedManualKey.Ah.Md5 = &ManualKeyAhMd5{}
					if o.ManualKey.Ah.Md5.Misc != nil {
						entry.Misc["ManualKeyAhMd5"] = o.ManualKey.Ah.Md5.Misc
					}
					if o.ManualKey.Ah.Md5.Key != nil {
						nestedManualKey.Ah.Md5.Key = o.ManualKey.Ah.Md5.Key
					}
				}
				if o.ManualKey.Ah.Sha1 != nil {
					nestedManualKey.Ah.Sha1 = &ManualKeyAhSha1{}
					if o.ManualKey.Ah.Sha1.Misc != nil {
						entry.Misc["ManualKeyAhSha1"] = o.ManualKey.Ah.Sha1.Misc
					}
					if o.ManualKey.Ah.Sha1.Key != nil {
						nestedManualKey.Ah.Sha1.Key = o.ManualKey.Ah.Sha1.Key
					}
				}
				if o.ManualKey.Ah.Sha256 != nil {
					nestedManualKey.Ah.Sha256 = &ManualKeyAhSha256{}
					if o.ManualKey.Ah.Sha256.Misc != nil {
						entry.Misc["ManualKeyAhSha256"] = o.ManualKey.Ah.Sha256.Misc
					}
					if o.ManualKey.Ah.Sha256.Key != nil {
						nestedManualKey.Ah.Sha256.Key = o.ManualKey.Ah.Sha256.Key
					}
				}
				if o.ManualKey.Ah.Sha384 != nil {
					nestedManualKey.Ah.Sha384 = &ManualKeyAhSha384{}
					if o.ManualKey.Ah.Sha384.Misc != nil {
						entry.Misc["ManualKeyAhSha384"] = o.ManualKey.Ah.Sha384.Misc
					}
					if o.ManualKey.Ah.Sha384.Key != nil {
						nestedManualKey.Ah.Sha384.Key = o.ManualKey.Ah.Sha384.Key
					}
				}
				if o.ManualKey.Ah.Sha512 != nil {
					nestedManualKey.Ah.Sha512 = &ManualKeyAhSha512{}
					if o.ManualKey.Ah.Sha512.Misc != nil {
						entry.Misc["ManualKeyAhSha512"] = o.ManualKey.Ah.Sha512.Misc
					}
					if o.ManualKey.Ah.Sha512.Key != nil {
						nestedManualKey.Ah.Sha512.Key = o.ManualKey.Ah.Sha512.Key
					}
				}
			}
			if o.ManualKey.Esp != nil {
				nestedManualKey.Esp = &ManualKeyEsp{}
				if o.ManualKey.Esp.Misc != nil {
					entry.Misc["ManualKeyEsp"] = o.ManualKey.Esp.Misc
				}
				if o.ManualKey.Esp.Authentication != nil {
					nestedManualKey.Esp.Authentication = &ManualKeyEspAuthentication{}
					if o.ManualKey.Esp.Authentication.Misc != nil {
						entry.Misc["ManualKeyEspAuthentication"] = o.ManualKey.Esp.Authentication.Misc
					}
					if o.ManualKey.Esp.Authentication.Md5 != nil {
						nestedManualKey.Esp.Authentication.Md5 = &ManualKeyEspAuthenticationMd5{}
						if o.ManualKey.Esp.Authentication.Md5.Misc != nil {
							entry.Misc["ManualKeyEspAuthenticationMd5"] = o.ManualKey.Esp.Authentication.Md5.Misc
						}
						if o.ManualKey.Esp.Authentication.Md5.Key != nil {
							nestedManualKey.Esp.Authentication.Md5.Key = o.ManualKey.Esp.Authentication.Md5.Key
						}
					}
					if o.ManualKey.Esp.Authentication.None != nil {
						nestedManualKey.Esp.Authentication.None = &ManualKeyEspAuthenticationNone{}
						if o.ManualKey.Esp.Authentication.None.Misc != nil {
							entry.Misc["ManualKeyEspAuthenticationNone"] = o.ManualKey.Esp.Authentication.None.Misc
						}
					}
					if o.ManualKey.Esp.Authentication.Sha1 != nil {
						nestedManualKey.Esp.Authentication.Sha1 = &ManualKeyEspAuthenticationSha1{}
						if o.ManualKey.Esp.Authentication.Sha1.Misc != nil {
							entry.Misc["ManualKeyEspAuthenticationSha1"] = o.ManualKey.Esp.Authentication.Sha1.Misc
						}
						if o.ManualKey.Esp.Authentication.Sha1.Key != nil {
							nestedManualKey.Esp.Authentication.Sha1.Key = o.ManualKey.Esp.Authentication.Sha1.Key
						}
					}
					if o.ManualKey.Esp.Authentication.Sha256 != nil {
						nestedManualKey.Esp.Authentication.Sha256 = &ManualKeyEspAuthenticationSha256{}
						if o.ManualKey.Esp.Authentication.Sha256.Misc != nil {
							entry.Misc["ManualKeyEspAuthenticationSha256"] = o.ManualKey.Esp.Authentication.Sha256.Misc
						}
						if o.ManualKey.Esp.Authentication.Sha256.Key != nil {
							nestedManualKey.Esp.Authentication.Sha256.Key = o.ManualKey.Esp.Authentication.Sha256.Key
						}
					}
					if o.ManualKey.Esp.Authentication.Sha384 != nil {
						nestedManualKey.Esp.Authentication.Sha384 = &ManualKeyEspAuthenticationSha384{}
						if o.ManualKey.Esp.Authentication.Sha384.Misc != nil {
							entry.Misc["ManualKeyEspAuthenticationSha384"] = o.ManualKey.Esp.Authentication.Sha384.Misc
						}
						if o.ManualKey.Esp.Authentication.Sha384.Key != nil {
							nestedManualKey.Esp.Authentication.Sha384.Key = o.ManualKey.Esp.Authentication.Sha384.Key
						}
					}
					if o.ManualKey.Esp.Authentication.Sha512 != nil {
						nestedManualKey.Esp.Authentication.Sha512 = &ManualKeyEspAuthenticationSha512{}
						if o.ManualKey.Esp.Authentication.Sha512.Misc != nil {
							entry.Misc["ManualKeyEspAuthenticationSha512"] = o.ManualKey.Esp.Authentication.Sha512.Misc
						}
						if o.ManualKey.Esp.Authentication.Sha512.Key != nil {
							nestedManualKey.Esp.Authentication.Sha512.Key = o.ManualKey.Esp.Authentication.Sha512.Key
						}
					}
				}
				if o.ManualKey.Esp.Encryption != nil {
					nestedManualKey.Esp.Encryption = &ManualKeyEspEncryption{}
					if o.ManualKey.Esp.Encryption.Misc != nil {
						entry.Misc["ManualKeyEspEncryption"] = o.ManualKey.Esp.Encryption.Misc
					}
					if o.ManualKey.Esp.Encryption.Algorithm != nil {
						nestedManualKey.Esp.Encryption.Algorithm = o.ManualKey.Esp.Encryption.Algorithm
					}
					if o.ManualKey.Esp.Encryption.Key != nil {
						nestedManualKey.Esp.Encryption.Key = o.ManualKey.Esp.Encryption.Key
					}
				}
			}
		}
		entry.ManualKey = nestedManualKey

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
	if !util.BoolsMatch(a.AntiReplay, b.AntiReplay) {
		return false
	}
	if !util.StringsMatch(a.AntiReplayWindow, b.AntiReplayWindow) {
		return false
	}
	if !util.StringsMatch(a.Comment, b.Comment) {
		return false
	}
	if !util.BoolsMatch(a.CopyFlowLabel, b.CopyFlowLabel) {
		return false
	}
	if !util.BoolsMatch(a.CopyTos, b.CopyTos) {
		return false
	}
	if !util.BoolsMatch(a.Disabled, b.Disabled) {
		return false
	}
	if !util.BoolsMatch(a.EnableGreEncapsulation, b.EnableGreEncapsulation) {
		return false
	}
	if !util.StringsMatch(a.IpsecMode, b.IpsecMode) {
		return false
	}
	if !util.BoolsMatch(a.Ipv6, b.Ipv6) {
		return false
	}
	if !util.StringsMatch(a.TunnelInterface, b.TunnelInterface) {
		return false
	}
	if !matchTunnelMonitor(a.TunnelMonitor, b.TunnelMonitor) {
		return false
	}
	if !matchAutoKey(a.AutoKey, b.AutoKey) {
		return false
	}
	if !matchGlobalProtectSatellite(a.GlobalProtectSatellite, b.GlobalProtectSatellite) {
		return false
	}
	if !matchManualKey(a.ManualKey, b.ManualKey) {
		return false
	}

	return true
}

func matchTunnelMonitor(a *TunnelMonitor, b *TunnelMonitor) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.DestinationIp, b.DestinationIp) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.StringsMatch(a.ProxyId, b.ProxyId) {
		return false
	}
	if !util.StringsMatch(a.TunnelMonitorProfile, b.TunnelMonitorProfile) {
		return false
	}
	return true
}
func matchManualKeyPeerAddress(a *ManualKeyPeerAddress, b *ManualKeyPeerAddress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Ip, b.Ip) {
		return false
	}
	return true
}
func matchManualKeyLocalAddress(a *ManualKeyLocalAddress, b *ManualKeyLocalAddress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Interface, b.Interface) {
		return false
	}
	if !util.StringsMatch(a.FloatingIp, b.FloatingIp) {
		return false
	}
	if !util.StringsMatch(a.Ip, b.Ip) {
		return false
	}
	return true
}
func matchManualKeyAhSha384(a *ManualKeyAhSha384, b *ManualKeyAhSha384) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchManualKeyAhSha512(a *ManualKeyAhSha512, b *ManualKeyAhSha512) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchManualKeyAhMd5(a *ManualKeyAhMd5, b *ManualKeyAhMd5) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchManualKeyAhSha1(a *ManualKeyAhSha1, b *ManualKeyAhSha1) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchManualKeyAhSha256(a *ManualKeyAhSha256, b *ManualKeyAhSha256) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchManualKeyAh(a *ManualKeyAh, b *ManualKeyAh) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchManualKeyAhMd5(a.Md5, b.Md5) {
		return false
	}
	if !matchManualKeyAhSha1(a.Sha1, b.Sha1) {
		return false
	}
	if !matchManualKeyAhSha256(a.Sha256, b.Sha256) {
		return false
	}
	if !matchManualKeyAhSha384(a.Sha384, b.Sha384) {
		return false
	}
	if !matchManualKeyAhSha512(a.Sha512, b.Sha512) {
		return false
	}
	return true
}
func matchManualKeyEspAuthenticationSha384(a *ManualKeyEspAuthenticationSha384, b *ManualKeyEspAuthenticationSha384) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchManualKeyEspAuthenticationSha512(a *ManualKeyEspAuthenticationSha512, b *ManualKeyEspAuthenticationSha512) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchManualKeyEspAuthenticationMd5(a *ManualKeyEspAuthenticationMd5, b *ManualKeyEspAuthenticationMd5) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchManualKeyEspAuthenticationNone(a *ManualKeyEspAuthenticationNone, b *ManualKeyEspAuthenticationNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchManualKeyEspAuthenticationSha1(a *ManualKeyEspAuthenticationSha1, b *ManualKeyEspAuthenticationSha1) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchManualKeyEspAuthenticationSha256(a *ManualKeyEspAuthenticationSha256, b *ManualKeyEspAuthenticationSha256) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchManualKeyEspAuthentication(a *ManualKeyEspAuthentication, b *ManualKeyEspAuthentication) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchManualKeyEspAuthenticationSha384(a.Sha384, b.Sha384) {
		return false
	}
	if !matchManualKeyEspAuthenticationSha512(a.Sha512, b.Sha512) {
		return false
	}
	if !matchManualKeyEspAuthenticationMd5(a.Md5, b.Md5) {
		return false
	}
	if !matchManualKeyEspAuthenticationNone(a.None, b.None) {
		return false
	}
	if !matchManualKeyEspAuthenticationSha1(a.Sha1, b.Sha1) {
		return false
	}
	if !matchManualKeyEspAuthenticationSha256(a.Sha256, b.Sha256) {
		return false
	}
	return true
}
func matchManualKeyEspEncryption(a *ManualKeyEspEncryption, b *ManualKeyEspEncryption) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Algorithm, b.Algorithm) {
		return false
	}
	if !util.StringsMatch(a.Key, b.Key) {
		return false
	}
	return true
}
func matchManualKeyEsp(a *ManualKeyEsp, b *ManualKeyEsp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchManualKeyEspAuthentication(a.Authentication, b.Authentication) {
		return false
	}
	if !matchManualKeyEspEncryption(a.Encryption, b.Encryption) {
		return false
	}
	return true
}
func matchManualKey(a *ManualKey, b *ManualKey) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchManualKeyLocalAddress(a.LocalAddress, b.LocalAddress) {
		return false
	}
	if !util.StringsMatch(a.LocalSpi, b.LocalSpi) {
		return false
	}
	if !matchManualKeyPeerAddress(a.PeerAddress, b.PeerAddress) {
		return false
	}
	if !util.StringsMatch(a.RemoteSpi, b.RemoteSpi) {
		return false
	}
	if !matchManualKeyAh(a.Ah, b.Ah) {
		return false
	}
	if !matchManualKeyEsp(a.Esp, b.Esp) {
		return false
	}
	return true
}
func matchAutoKeyIkeGateway(a []AutoKeyIkeGateway, b []AutoKeyIkeGateway) bool {
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
func matchAutoKeyProxyIdProtocolUdp(a *AutoKeyProxyIdProtocolUdp, b *AutoKeyProxyIdProtocolUdp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.LocalPort, b.LocalPort) {
		return false
	}
	if !util.Ints64Match(a.RemotePort, b.RemotePort) {
		return false
	}
	return true
}
func matchAutoKeyProxyIdProtocolAny(a *AutoKeyProxyIdProtocolAny, b *AutoKeyProxyIdProtocolAny) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchAutoKeyProxyIdProtocolTcp(a *AutoKeyProxyIdProtocolTcp, b *AutoKeyProxyIdProtocolTcp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.RemotePort, b.RemotePort) {
		return false
	}
	if !util.Ints64Match(a.LocalPort, b.LocalPort) {
		return false
	}
	return true
}
func matchAutoKeyProxyIdProtocol(a *AutoKeyProxyIdProtocol, b *AutoKeyProxyIdProtocol) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Number, b.Number) {
		return false
	}
	if !matchAutoKeyProxyIdProtocolAny(a.Any, b.Any) {
		return false
	}
	if !matchAutoKeyProxyIdProtocolTcp(a.Tcp, b.Tcp) {
		return false
	}
	if !matchAutoKeyProxyIdProtocolUdp(a.Udp, b.Udp) {
		return false
	}
	return true
}
func matchAutoKeyProxyId(a []AutoKeyProxyId, b []AutoKeyProxyId) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Local, b.Local) {
				return false
			}
			if !util.StringsMatch(a.Remote, b.Remote) {
				return false
			}
			if !matchAutoKeyProxyIdProtocol(a.Protocol, b.Protocol) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchAutoKeyProxyIdV6ProtocolAny(a *AutoKeyProxyIdV6ProtocolAny, b *AutoKeyProxyIdV6ProtocolAny) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchAutoKeyProxyIdV6ProtocolTcp(a *AutoKeyProxyIdV6ProtocolTcp, b *AutoKeyProxyIdV6ProtocolTcp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.LocalPort, b.LocalPort) {
		return false
	}
	if !util.Ints64Match(a.RemotePort, b.RemotePort) {
		return false
	}
	return true
}
func matchAutoKeyProxyIdV6ProtocolUdp(a *AutoKeyProxyIdV6ProtocolUdp, b *AutoKeyProxyIdV6ProtocolUdp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.LocalPort, b.LocalPort) {
		return false
	}
	if !util.Ints64Match(a.RemotePort, b.RemotePort) {
		return false
	}
	return true
}
func matchAutoKeyProxyIdV6Protocol(a *AutoKeyProxyIdV6Protocol, b *AutoKeyProxyIdV6Protocol) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.Number, b.Number) {
		return false
	}
	if !matchAutoKeyProxyIdV6ProtocolAny(a.Any, b.Any) {
		return false
	}
	if !matchAutoKeyProxyIdV6ProtocolTcp(a.Tcp, b.Tcp) {
		return false
	}
	if !matchAutoKeyProxyIdV6ProtocolUdp(a.Udp, b.Udp) {
		return false
	}
	return true
}
func matchAutoKeyProxyIdV6(a []AutoKeyProxyIdV6, b []AutoKeyProxyIdV6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.Local, b.Local) {
				return false
			}
			if !util.StringsMatch(a.Remote, b.Remote) {
				return false
			}
			if !matchAutoKeyProxyIdV6Protocol(a.Protocol, b.Protocol) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchAutoKey(a *AutoKey, b *AutoKey) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchAutoKeyIkeGateway(a.IkeGateway, b.IkeGateway) {
		return false
	}
	if !util.StringsMatch(a.IpsecCryptoProfile, b.IpsecCryptoProfile) {
		return false
	}
	if !matchAutoKeyProxyId(a.ProxyId, b.ProxyId) {
		return false
	}
	if !matchAutoKeyProxyIdV6(a.ProxyIdV6, b.ProxyIdV6) {
		return false
	}
	return true
}
func matchGlobalProtectSatelliteExternalCa(a *GlobalProtectSatelliteExternalCa, b *GlobalProtectSatelliteExternalCa) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.LocalCertificate, b.LocalCertificate) {
		return false
	}
	if !util.StringsMatch(a.CertificateProfile, b.CertificateProfile) {
		return false
	}
	return true
}
func matchGlobalProtectSatelliteLocalAddressFloatingIp(a *GlobalProtectSatelliteLocalAddressFloatingIp, b *GlobalProtectSatelliteLocalAddressFloatingIp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Ipv4, b.Ipv4) {
		return false
	}
	if !util.StringsMatch(a.Ipv6, b.Ipv6) {
		return false
	}
	return true
}
func matchGlobalProtectSatelliteLocalAddressIp(a *GlobalProtectSatelliteLocalAddressIp, b *GlobalProtectSatelliteLocalAddressIp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Ipv4, b.Ipv4) {
		return false
	}
	if !util.StringsMatch(a.Ipv6, b.Ipv6) {
		return false
	}
	return true
}
func matchGlobalProtectSatelliteLocalAddress(a *GlobalProtectSatelliteLocalAddress, b *GlobalProtectSatelliteLocalAddress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Interface, b.Interface) {
		return false
	}
	if !matchGlobalProtectSatelliteLocalAddressIp(a.Ip, b.Ip) {
		return false
	}
	if !matchGlobalProtectSatelliteLocalAddressFloatingIp(a.FloatingIp, b.FloatingIp) {
		return false
	}
	return true
}
func matchGlobalProtectSatellitePublishConnectedRoutes(a *GlobalProtectSatellitePublishConnectedRoutes, b *GlobalProtectSatellitePublishConnectedRoutes) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	return true
}
func matchGlobalProtectSatellite(a *GlobalProtectSatellite, b *GlobalProtectSatellite) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Ipv6Preferred, b.Ipv6Preferred) {
		return false
	}
	if !matchGlobalProtectSatelliteLocalAddress(a.LocalAddress, b.LocalAddress) {
		return false
	}
	if !util.StringsMatch(a.PortalAddress, b.PortalAddress) {
		return false
	}
	if !matchGlobalProtectSatellitePublishConnectedRoutes(a.PublishConnectedRoutes, b.PublishConnectedRoutes) {
		return false
	}
	if !util.OrderedListsMatch(a.PublishRoutes, b.PublishRoutes) {
		return false
	}
	if !matchGlobalProtectSatelliteExternalCa(a.ExternalCa, b.ExternalCa) {
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
