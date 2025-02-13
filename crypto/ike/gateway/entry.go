package gateway

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
	Suffix = []string{}
)

type Entry struct {
	Name           string
	Authentication *Authentication
	Comment        *string
	Disabled       *bool
	Ipv6           *bool
	LocalAddress   *LocalAddress
	LocalId        *LocalId
	PeerAddress    *PeerAddress
	PeerId         *PeerId
	Protocol       *Protocol
	ProtocolCommon *ProtocolCommon

	Misc map[string][]generic.Xml
}

type Authentication struct {
	Certificate  *AuthenticationCertificate
	PreSharedKey *AuthenticationPreSharedKey
}
type AuthenticationCertificate struct {
	AllowIdPayloadMismatch     *bool
	CertificateProfile         *string
	LocalCertificate           *AuthenticationCertificateLocalCertificate
	StrictValidationRevocation *bool
	UseManagementAsSource      *bool
}
type AuthenticationCertificateLocalCertificate struct {
	HashAndUrl *AuthenticationCertificateLocalCertificateHashAndUrl
	Name       *string
}
type AuthenticationCertificateLocalCertificateHashAndUrl struct {
	BaseUrl *string
	Enable  *bool
}
type AuthenticationPreSharedKey struct {
	Key *string
}
type LocalAddress struct {
	Interface  *string
	FloatingIp *string
	Ip         *string
}
type LocalId struct {
	Id   *string
	Type *string
}
type PeerAddress struct {
	Dynamic *PeerAddressDynamic
	Fqdn    *string
	Ip      *string
}
type PeerAddressDynamic struct {
}
type PeerId struct {
	Id       *string
	Matching *string
	Type     *string
}
type Protocol struct {
	Ikev1   *ProtocolIkev1
	Ikev2   *ProtocolIkev2
	Version *string
}
type ProtocolCommon struct {
	Fragmentation *ProtocolCommonFragmentation
	NatTraversal  *ProtocolCommonNatTraversal
	PassiveMode   *bool
}
type ProtocolCommonFragmentation struct {
	Enable *bool
}
type ProtocolCommonNatTraversal struct {
	Enable            *bool
	KeepAliveInterval *int64
	UdpChecksumEnable *bool
}
type ProtocolIkev1 struct {
	Dpd              *ProtocolIkev1Dpd
	ExchangeMode     *string
	IkeCryptoProfile *string
}
type ProtocolIkev1Dpd struct {
	Enable   *bool
	Interval *int64
	Retry    *int64
}
type ProtocolIkev2 struct {
	Dpd              *ProtocolIkev2Dpd
	IkeCryptoProfile *string
	RequireCookie    *bool
}
type ProtocolIkev2Dpd struct {
	Enable   *bool
	Interval *int64
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName        xml.Name           `xml:"entry"`
	Name           string             `xml:"name,attr"`
	Authentication *AuthenticationXml `xml:"authentication,omitempty"`
	Comment        *string            `xml:"comment,omitempty"`
	Disabled       *string            `xml:"disabled,omitempty"`
	Ipv6           *string            `xml:"ipv6,omitempty"`
	LocalAddress   *LocalAddressXml   `xml:"local-address,omitempty"`
	LocalId        *LocalIdXml        `xml:"local-id,omitempty"`
	PeerAddress    *PeerAddressXml    `xml:"peer-address,omitempty"`
	PeerId         *PeerIdXml         `xml:"peer-id,omitempty"`
	Protocol       *ProtocolXml       `xml:"protocol,omitempty"`
	ProtocolCommon *ProtocolCommonXml `xml:"protocol-common,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AuthenticationXml struct {
	Certificate  *AuthenticationCertificateXml  `xml:"certificate,omitempty"`
	PreSharedKey *AuthenticationPreSharedKeyXml `xml:"pre-shared-key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AuthenticationCertificateXml struct {
	AllowIdPayloadMismatch     *string                                       `xml:"allow-id-payload-mismatch,omitempty"`
	CertificateProfile         *string                                       `xml:"certificate-profile,omitempty"`
	LocalCertificate           *AuthenticationCertificateLocalCertificateXml `xml:"local-certificate,omitempty"`
	StrictValidationRevocation *string                                       `xml:"strict-validation-revocation,omitempty"`
	UseManagementAsSource      *string                                       `xml:"use-management-as-source,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AuthenticationCertificateLocalCertificateXml struct {
	HashAndUrl *AuthenticationCertificateLocalCertificateHashAndUrlXml `xml:"hash-and-url,omitempty"`
	XMLName    xml.Name                                                `xml:"entry"`
	Name       *string                                                 `xml:"name,attr,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AuthenticationCertificateLocalCertificateHashAndUrlXml struct {
	BaseUrl *string `xml:"base-url,omitempty"`
	Enable  *string `xml:"enable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type AuthenticationPreSharedKeyXml struct {
	Key *string `xml:"key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type LocalAddressXml struct {
	Interface  *string `xml:"interface,omitempty"`
	FloatingIp *string `xml:"floating-ip,omitempty"`
	Ip         *string `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type LocalIdXml struct {
	Id   *string `xml:"id,omitempty"`
	Type *string `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type PeerAddressXml struct {
	Dynamic *PeerAddressDynamicXml `xml:"dynamic,omitempty"`
	Fqdn    *string                `xml:"fqdn,omitempty"`
	Ip      *string                `xml:"ip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type PeerAddressDynamicXml struct {
	Misc []generic.Xml `xml:",any"`
}
type PeerIdXml struct {
	Id       *string `xml:"id,omitempty"`
	Matching *string `xml:"matching,omitempty"`
	Type     *string `xml:"type,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolXml struct {
	Ikev1   *ProtocolIkev1Xml `xml:"ikev1,omitempty"`
	Ikev2   *ProtocolIkev2Xml `xml:"ikev2,omitempty"`
	Version *string           `xml:"version,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolCommonXml struct {
	Fragmentation *ProtocolCommonFragmentationXml `xml:"fragmentation,omitempty"`
	NatTraversal  *ProtocolCommonNatTraversalXml  `xml:"nat-traversal,omitempty"`
	PassiveMode   *string                         `xml:"passive-mode,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolCommonFragmentationXml struct {
	Enable *string `xml:"enable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolCommonNatTraversalXml struct {
	Enable            *string `xml:"enable,omitempty"`
	KeepAliveInterval *int64  `xml:"keep-alive-interval,omitempty"`
	UdpChecksumEnable *string `xml:"udp-checksum-enable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolIkev1Xml struct {
	Dpd              *ProtocolIkev1DpdXml `xml:"dpd,omitempty"`
	ExchangeMode     *string              `xml:"exchange-mode,omitempty"`
	IkeCryptoProfile *string              `xml:"ike-crypto-profile,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolIkev1DpdXml struct {
	Enable   *string `xml:"enable,omitempty"`
	Interval *int64  `xml:"interval,omitempty"`
	Retry    *int64  `xml:"retry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolIkev2Xml struct {
	Dpd              *ProtocolIkev2DpdXml `xml:"dpd,omitempty"`
	IkeCryptoProfile *string              `xml:"ike-crypto-profile,omitempty"`
	RequireCookie    *string              `xml:"require-cookie,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolIkev2DpdXml struct {
	Enable   *string `xml:"enable,omitempty"`
	Interval *int64  `xml:"interval,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "authentication" || v == "Authentication" {
		return e.Authentication, nil
	}
	if v == "comment" || v == "Comment" {
		return e.Comment, nil
	}
	if v == "disabled" || v == "Disabled" {
		return e.Disabled, nil
	}
	if v == "ipv6" || v == "Ipv6" {
		return e.Ipv6, nil
	}
	if v == "local_address" || v == "LocalAddress" {
		return e.LocalAddress, nil
	}
	if v == "local_id" || v == "LocalId" {
		return e.LocalId, nil
	}
	if v == "peer_address" || v == "PeerAddress" {
		return e.PeerAddress, nil
	}
	if v == "peer_id" || v == "PeerId" {
		return e.PeerId, nil
	}
	if v == "protocol" || v == "Protocol" {
		return e.Protocol, nil
	}
	if v == "protocol_common" || v == "ProtocolCommon" {
		return e.ProtocolCommon, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {

	return specifyEntry, &entryXmlContainer{}, nil
}
func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}
	entry.Name = o.Name
	var nestedAuthentication *AuthenticationXml
	if o.Authentication != nil {
		nestedAuthentication = &AuthenticationXml{}
		if _, ok := o.Misc["Authentication"]; ok {
			nestedAuthentication.Misc = o.Misc["Authentication"]
		}
		if o.Authentication.Certificate != nil {
			nestedAuthentication.Certificate = &AuthenticationCertificateXml{}
			if _, ok := o.Misc["AuthenticationCertificate"]; ok {
				nestedAuthentication.Certificate.Misc = o.Misc["AuthenticationCertificate"]
			}
			if o.Authentication.Certificate.UseManagementAsSource != nil {
				nestedAuthentication.Certificate.UseManagementAsSource = util.YesNo(o.Authentication.Certificate.UseManagementAsSource, nil)
			}
			if o.Authentication.Certificate.AllowIdPayloadMismatch != nil {
				nestedAuthentication.Certificate.AllowIdPayloadMismatch = util.YesNo(o.Authentication.Certificate.AllowIdPayloadMismatch, nil)
			}
			if o.Authentication.Certificate.CertificateProfile != nil {
				nestedAuthentication.Certificate.CertificateProfile = o.Authentication.Certificate.CertificateProfile
			}
			if o.Authentication.Certificate.LocalCertificate != nil {
				nestedAuthentication.Certificate.LocalCertificate = &AuthenticationCertificateLocalCertificateXml{}
				if _, ok := o.Misc["AuthenticationCertificateLocalCertificate"]; ok {
					nestedAuthentication.Certificate.LocalCertificate.Misc = o.Misc["AuthenticationCertificateLocalCertificate"]
				}
				if o.Authentication.Certificate.LocalCertificate.HashAndUrl != nil {
					nestedAuthentication.Certificate.LocalCertificate.HashAndUrl = &AuthenticationCertificateLocalCertificateHashAndUrlXml{}
					if _, ok := o.Misc["AuthenticationCertificateLocalCertificateHashAndUrl"]; ok {
						nestedAuthentication.Certificate.LocalCertificate.HashAndUrl.Misc = o.Misc["AuthenticationCertificateLocalCertificateHashAndUrl"]
					}
					if o.Authentication.Certificate.LocalCertificate.HashAndUrl.BaseUrl != nil {
						nestedAuthentication.Certificate.LocalCertificate.HashAndUrl.BaseUrl = o.Authentication.Certificate.LocalCertificate.HashAndUrl.BaseUrl
					}
					if o.Authentication.Certificate.LocalCertificate.HashAndUrl.Enable != nil {
						nestedAuthentication.Certificate.LocalCertificate.HashAndUrl.Enable = util.YesNo(o.Authentication.Certificate.LocalCertificate.HashAndUrl.Enable, nil)
					}
				}
				if o.Authentication.Certificate.LocalCertificate.Name != nil {
					nestedAuthentication.Certificate.LocalCertificate.Name = o.Authentication.Certificate.LocalCertificate.Name
				}
			}
			if o.Authentication.Certificate.StrictValidationRevocation != nil {
				nestedAuthentication.Certificate.StrictValidationRevocation = util.YesNo(o.Authentication.Certificate.StrictValidationRevocation, nil)
			}
		}
		if o.Authentication.PreSharedKey != nil {
			nestedAuthentication.PreSharedKey = &AuthenticationPreSharedKeyXml{}
			if _, ok := o.Misc["AuthenticationPreSharedKey"]; ok {
				nestedAuthentication.PreSharedKey.Misc = o.Misc["AuthenticationPreSharedKey"]
			}
			if o.Authentication.PreSharedKey.Key != nil {
				nestedAuthentication.PreSharedKey.Key = o.Authentication.PreSharedKey.Key
			}
		}
	}
	entry.Authentication = nestedAuthentication

	entry.Comment = o.Comment
	entry.Disabled = util.YesNo(o.Disabled, nil)
	entry.Ipv6 = util.YesNo(o.Ipv6, nil)
	var nestedLocalAddress *LocalAddressXml
	if o.LocalAddress != nil {
		nestedLocalAddress = &LocalAddressXml{}
		if _, ok := o.Misc["LocalAddress"]; ok {
			nestedLocalAddress.Misc = o.Misc["LocalAddress"]
		}
		if o.LocalAddress.Interface != nil {
			nestedLocalAddress.Interface = o.LocalAddress.Interface
		}
		if o.LocalAddress.FloatingIp != nil {
			nestedLocalAddress.FloatingIp = o.LocalAddress.FloatingIp
		}
		if o.LocalAddress.Ip != nil {
			nestedLocalAddress.Ip = o.LocalAddress.Ip
		}
	}
	entry.LocalAddress = nestedLocalAddress

	var nestedLocalId *LocalIdXml
	if o.LocalId != nil {
		nestedLocalId = &LocalIdXml{}
		if _, ok := o.Misc["LocalId"]; ok {
			nestedLocalId.Misc = o.Misc["LocalId"]
		}
		if o.LocalId.Id != nil {
			nestedLocalId.Id = o.LocalId.Id
		}
		if o.LocalId.Type != nil {
			nestedLocalId.Type = o.LocalId.Type
		}
	}
	entry.LocalId = nestedLocalId

	var nestedPeerAddress *PeerAddressXml
	if o.PeerAddress != nil {
		nestedPeerAddress = &PeerAddressXml{}
		if _, ok := o.Misc["PeerAddress"]; ok {
			nestedPeerAddress.Misc = o.Misc["PeerAddress"]
		}
		if o.PeerAddress.Dynamic != nil {
			nestedPeerAddress.Dynamic = &PeerAddressDynamicXml{}
			if _, ok := o.Misc["PeerAddressDynamic"]; ok {
				nestedPeerAddress.Dynamic.Misc = o.Misc["PeerAddressDynamic"]
			}
		}
		if o.PeerAddress.Fqdn != nil {
			nestedPeerAddress.Fqdn = o.PeerAddress.Fqdn
		}
		if o.PeerAddress.Ip != nil {
			nestedPeerAddress.Ip = o.PeerAddress.Ip
		}
	}
	entry.PeerAddress = nestedPeerAddress

	var nestedPeerId *PeerIdXml
	if o.PeerId != nil {
		nestedPeerId = &PeerIdXml{}
		if _, ok := o.Misc["PeerId"]; ok {
			nestedPeerId.Misc = o.Misc["PeerId"]
		}
		if o.PeerId.Id != nil {
			nestedPeerId.Id = o.PeerId.Id
		}
		if o.PeerId.Matching != nil {
			nestedPeerId.Matching = o.PeerId.Matching
		}
		if o.PeerId.Type != nil {
			nestedPeerId.Type = o.PeerId.Type
		}
	}
	entry.PeerId = nestedPeerId

	var nestedProtocol *ProtocolXml
	if o.Protocol != nil {
		nestedProtocol = &ProtocolXml{}
		if _, ok := o.Misc["Protocol"]; ok {
			nestedProtocol.Misc = o.Misc["Protocol"]
		}
		if o.Protocol.Ikev2 != nil {
			nestedProtocol.Ikev2 = &ProtocolIkev2Xml{}
			if _, ok := o.Misc["ProtocolIkev2"]; ok {
				nestedProtocol.Ikev2.Misc = o.Misc["ProtocolIkev2"]
			}
			if o.Protocol.Ikev2.Dpd != nil {
				nestedProtocol.Ikev2.Dpd = &ProtocolIkev2DpdXml{}
				if _, ok := o.Misc["ProtocolIkev2Dpd"]; ok {
					nestedProtocol.Ikev2.Dpd.Misc = o.Misc["ProtocolIkev2Dpd"]
				}
				if o.Protocol.Ikev2.Dpd.Enable != nil {
					nestedProtocol.Ikev2.Dpd.Enable = util.YesNo(o.Protocol.Ikev2.Dpd.Enable, nil)
				}
				if o.Protocol.Ikev2.Dpd.Interval != nil {
					nestedProtocol.Ikev2.Dpd.Interval = o.Protocol.Ikev2.Dpd.Interval
				}
			}
			if o.Protocol.Ikev2.IkeCryptoProfile != nil {
				nestedProtocol.Ikev2.IkeCryptoProfile = o.Protocol.Ikev2.IkeCryptoProfile
			}
			if o.Protocol.Ikev2.RequireCookie != nil {
				nestedProtocol.Ikev2.RequireCookie = util.YesNo(o.Protocol.Ikev2.RequireCookie, nil)
			}
		}
		if o.Protocol.Version != nil {
			nestedProtocol.Version = o.Protocol.Version
		}
		if o.Protocol.Ikev1 != nil {
			nestedProtocol.Ikev1 = &ProtocolIkev1Xml{}
			if _, ok := o.Misc["ProtocolIkev1"]; ok {
				nestedProtocol.Ikev1.Misc = o.Misc["ProtocolIkev1"]
			}
			if o.Protocol.Ikev1.Dpd != nil {
				nestedProtocol.Ikev1.Dpd = &ProtocolIkev1DpdXml{}
				if _, ok := o.Misc["ProtocolIkev1Dpd"]; ok {
					nestedProtocol.Ikev1.Dpd.Misc = o.Misc["ProtocolIkev1Dpd"]
				}
				if o.Protocol.Ikev1.Dpd.Enable != nil {
					nestedProtocol.Ikev1.Dpd.Enable = util.YesNo(o.Protocol.Ikev1.Dpd.Enable, nil)
				}
				if o.Protocol.Ikev1.Dpd.Interval != nil {
					nestedProtocol.Ikev1.Dpd.Interval = o.Protocol.Ikev1.Dpd.Interval
				}
				if o.Protocol.Ikev1.Dpd.Retry != nil {
					nestedProtocol.Ikev1.Dpd.Retry = o.Protocol.Ikev1.Dpd.Retry
				}
			}
			if o.Protocol.Ikev1.ExchangeMode != nil {
				nestedProtocol.Ikev1.ExchangeMode = o.Protocol.Ikev1.ExchangeMode
			}
			if o.Protocol.Ikev1.IkeCryptoProfile != nil {
				nestedProtocol.Ikev1.IkeCryptoProfile = o.Protocol.Ikev1.IkeCryptoProfile
			}
		}
	}
	entry.Protocol = nestedProtocol

	var nestedProtocolCommon *ProtocolCommonXml
	if o.ProtocolCommon != nil {
		nestedProtocolCommon = &ProtocolCommonXml{}
		if _, ok := o.Misc["ProtocolCommon"]; ok {
			nestedProtocolCommon.Misc = o.Misc["ProtocolCommon"]
		}
		if o.ProtocolCommon.Fragmentation != nil {
			nestedProtocolCommon.Fragmentation = &ProtocolCommonFragmentationXml{}
			if _, ok := o.Misc["ProtocolCommonFragmentation"]; ok {
				nestedProtocolCommon.Fragmentation.Misc = o.Misc["ProtocolCommonFragmentation"]
			}
			if o.ProtocolCommon.Fragmentation.Enable != nil {
				nestedProtocolCommon.Fragmentation.Enable = util.YesNo(o.ProtocolCommon.Fragmentation.Enable, nil)
			}
		}
		if o.ProtocolCommon.NatTraversal != nil {
			nestedProtocolCommon.NatTraversal = &ProtocolCommonNatTraversalXml{}
			if _, ok := o.Misc["ProtocolCommonNatTraversal"]; ok {
				nestedProtocolCommon.NatTraversal.Misc = o.Misc["ProtocolCommonNatTraversal"]
			}
			if o.ProtocolCommon.NatTraversal.UdpChecksumEnable != nil {
				nestedProtocolCommon.NatTraversal.UdpChecksumEnable = util.YesNo(o.ProtocolCommon.NatTraversal.UdpChecksumEnable, nil)
			}
			if o.ProtocolCommon.NatTraversal.Enable != nil {
				nestedProtocolCommon.NatTraversal.Enable = util.YesNo(o.ProtocolCommon.NatTraversal.Enable, nil)
			}
			if o.ProtocolCommon.NatTraversal.KeepAliveInterval != nil {
				nestedProtocolCommon.NatTraversal.KeepAliveInterval = o.ProtocolCommon.NatTraversal.KeepAliveInterval
			}
		}
		if o.ProtocolCommon.PassiveMode != nil {
			nestedProtocolCommon.PassiveMode = util.YesNo(o.ProtocolCommon.PassiveMode, nil)
		}
	}
	entry.ProtocolCommon = nestedProtocolCommon

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
		var nestedAuthentication *Authentication
		if o.Authentication != nil {
			nestedAuthentication = &Authentication{}
			if o.Authentication.Misc != nil {
				entry.Misc["Authentication"] = o.Authentication.Misc
			}
			if o.Authentication.Certificate != nil {
				nestedAuthentication.Certificate = &AuthenticationCertificate{}
				if o.Authentication.Certificate.Misc != nil {
					entry.Misc["AuthenticationCertificate"] = o.Authentication.Certificate.Misc
				}
				if o.Authentication.Certificate.AllowIdPayloadMismatch != nil {
					nestedAuthentication.Certificate.AllowIdPayloadMismatch = util.AsBool(o.Authentication.Certificate.AllowIdPayloadMismatch, nil)
				}
				if o.Authentication.Certificate.CertificateProfile != nil {
					nestedAuthentication.Certificate.CertificateProfile = o.Authentication.Certificate.CertificateProfile
				}
				if o.Authentication.Certificate.LocalCertificate != nil {
					nestedAuthentication.Certificate.LocalCertificate = &AuthenticationCertificateLocalCertificate{}
					if o.Authentication.Certificate.LocalCertificate.Misc != nil {
						entry.Misc["AuthenticationCertificateLocalCertificate"] = o.Authentication.Certificate.LocalCertificate.Misc
					}
					if o.Authentication.Certificate.LocalCertificate.HashAndUrl != nil {
						nestedAuthentication.Certificate.LocalCertificate.HashAndUrl = &AuthenticationCertificateLocalCertificateHashAndUrl{}
						if o.Authentication.Certificate.LocalCertificate.HashAndUrl.Misc != nil {
							entry.Misc["AuthenticationCertificateLocalCertificateHashAndUrl"] = o.Authentication.Certificate.LocalCertificate.HashAndUrl.Misc
						}
						if o.Authentication.Certificate.LocalCertificate.HashAndUrl.BaseUrl != nil {
							nestedAuthentication.Certificate.LocalCertificate.HashAndUrl.BaseUrl = o.Authentication.Certificate.LocalCertificate.HashAndUrl.BaseUrl
						}
						if o.Authentication.Certificate.LocalCertificate.HashAndUrl.Enable != nil {
							nestedAuthentication.Certificate.LocalCertificate.HashAndUrl.Enable = util.AsBool(o.Authentication.Certificate.LocalCertificate.HashAndUrl.Enable, nil)
						}
					}
					if o.Authentication.Certificate.LocalCertificate.Name != nil {
						nestedAuthentication.Certificate.LocalCertificate.Name = o.Authentication.Certificate.LocalCertificate.Name
					}
				}
				if o.Authentication.Certificate.StrictValidationRevocation != nil {
					nestedAuthentication.Certificate.StrictValidationRevocation = util.AsBool(o.Authentication.Certificate.StrictValidationRevocation, nil)
				}
				if o.Authentication.Certificate.UseManagementAsSource != nil {
					nestedAuthentication.Certificate.UseManagementAsSource = util.AsBool(o.Authentication.Certificate.UseManagementAsSource, nil)
				}
			}
			if o.Authentication.PreSharedKey != nil {
				nestedAuthentication.PreSharedKey = &AuthenticationPreSharedKey{}
				if o.Authentication.PreSharedKey.Misc != nil {
					entry.Misc["AuthenticationPreSharedKey"] = o.Authentication.PreSharedKey.Misc
				}
				if o.Authentication.PreSharedKey.Key != nil {
					nestedAuthentication.PreSharedKey.Key = o.Authentication.PreSharedKey.Key
				}
			}
		}
		entry.Authentication = nestedAuthentication

		entry.Comment = o.Comment
		entry.Disabled = util.AsBool(o.Disabled, nil)
		entry.Ipv6 = util.AsBool(o.Ipv6, nil)
		var nestedLocalAddress *LocalAddress
		if o.LocalAddress != nil {
			nestedLocalAddress = &LocalAddress{}
			if o.LocalAddress.Misc != nil {
				entry.Misc["LocalAddress"] = o.LocalAddress.Misc
			}
			if o.LocalAddress.Interface != nil {
				nestedLocalAddress.Interface = o.LocalAddress.Interface
			}
			if o.LocalAddress.FloatingIp != nil {
				nestedLocalAddress.FloatingIp = o.LocalAddress.FloatingIp
			}
			if o.LocalAddress.Ip != nil {
				nestedLocalAddress.Ip = o.LocalAddress.Ip
			}
		}
		entry.LocalAddress = nestedLocalAddress

		var nestedLocalId *LocalId
		if o.LocalId != nil {
			nestedLocalId = &LocalId{}
			if o.LocalId.Misc != nil {
				entry.Misc["LocalId"] = o.LocalId.Misc
			}
			if o.LocalId.Id != nil {
				nestedLocalId.Id = o.LocalId.Id
			}
			if o.LocalId.Type != nil {
				nestedLocalId.Type = o.LocalId.Type
			}
		}
		entry.LocalId = nestedLocalId

		var nestedPeerAddress *PeerAddress
		if o.PeerAddress != nil {
			nestedPeerAddress = &PeerAddress{}
			if o.PeerAddress.Misc != nil {
				entry.Misc["PeerAddress"] = o.PeerAddress.Misc
			}
			if o.PeerAddress.Dynamic != nil {
				nestedPeerAddress.Dynamic = &PeerAddressDynamic{}
				if o.PeerAddress.Dynamic.Misc != nil {
					entry.Misc["PeerAddressDynamic"] = o.PeerAddress.Dynamic.Misc
				}
			}
			if o.PeerAddress.Fqdn != nil {
				nestedPeerAddress.Fqdn = o.PeerAddress.Fqdn
			}
			if o.PeerAddress.Ip != nil {
				nestedPeerAddress.Ip = o.PeerAddress.Ip
			}
		}
		entry.PeerAddress = nestedPeerAddress

		var nestedPeerId *PeerId
		if o.PeerId != nil {
			nestedPeerId = &PeerId{}
			if o.PeerId.Misc != nil {
				entry.Misc["PeerId"] = o.PeerId.Misc
			}
			if o.PeerId.Id != nil {
				nestedPeerId.Id = o.PeerId.Id
			}
			if o.PeerId.Matching != nil {
				nestedPeerId.Matching = o.PeerId.Matching
			}
			if o.PeerId.Type != nil {
				nestedPeerId.Type = o.PeerId.Type
			}
		}
		entry.PeerId = nestedPeerId

		var nestedProtocol *Protocol
		if o.Protocol != nil {
			nestedProtocol = &Protocol{}
			if o.Protocol.Misc != nil {
				entry.Misc["Protocol"] = o.Protocol.Misc
			}
			if o.Protocol.Version != nil {
				nestedProtocol.Version = o.Protocol.Version
			}
			if o.Protocol.Ikev1 != nil {
				nestedProtocol.Ikev1 = &ProtocolIkev1{}
				if o.Protocol.Ikev1.Misc != nil {
					entry.Misc["ProtocolIkev1"] = o.Protocol.Ikev1.Misc
				}
				if o.Protocol.Ikev1.Dpd != nil {
					nestedProtocol.Ikev1.Dpd = &ProtocolIkev1Dpd{}
					if o.Protocol.Ikev1.Dpd.Misc != nil {
						entry.Misc["ProtocolIkev1Dpd"] = o.Protocol.Ikev1.Dpd.Misc
					}
					if o.Protocol.Ikev1.Dpd.Enable != nil {
						nestedProtocol.Ikev1.Dpd.Enable = util.AsBool(o.Protocol.Ikev1.Dpd.Enable, nil)
					}
					if o.Protocol.Ikev1.Dpd.Interval != nil {
						nestedProtocol.Ikev1.Dpd.Interval = o.Protocol.Ikev1.Dpd.Interval
					}
					if o.Protocol.Ikev1.Dpd.Retry != nil {
						nestedProtocol.Ikev1.Dpd.Retry = o.Protocol.Ikev1.Dpd.Retry
					}
				}
				if o.Protocol.Ikev1.ExchangeMode != nil {
					nestedProtocol.Ikev1.ExchangeMode = o.Protocol.Ikev1.ExchangeMode
				}
				if o.Protocol.Ikev1.IkeCryptoProfile != nil {
					nestedProtocol.Ikev1.IkeCryptoProfile = o.Protocol.Ikev1.IkeCryptoProfile
				}
			}
			if o.Protocol.Ikev2 != nil {
				nestedProtocol.Ikev2 = &ProtocolIkev2{}
				if o.Protocol.Ikev2.Misc != nil {
					entry.Misc["ProtocolIkev2"] = o.Protocol.Ikev2.Misc
				}
				if o.Protocol.Ikev2.Dpd != nil {
					nestedProtocol.Ikev2.Dpd = &ProtocolIkev2Dpd{}
					if o.Protocol.Ikev2.Dpd.Misc != nil {
						entry.Misc["ProtocolIkev2Dpd"] = o.Protocol.Ikev2.Dpd.Misc
					}
					if o.Protocol.Ikev2.Dpd.Enable != nil {
						nestedProtocol.Ikev2.Dpd.Enable = util.AsBool(o.Protocol.Ikev2.Dpd.Enable, nil)
					}
					if o.Protocol.Ikev2.Dpd.Interval != nil {
						nestedProtocol.Ikev2.Dpd.Interval = o.Protocol.Ikev2.Dpd.Interval
					}
				}
				if o.Protocol.Ikev2.IkeCryptoProfile != nil {
					nestedProtocol.Ikev2.IkeCryptoProfile = o.Protocol.Ikev2.IkeCryptoProfile
				}
				if o.Protocol.Ikev2.RequireCookie != nil {
					nestedProtocol.Ikev2.RequireCookie = util.AsBool(o.Protocol.Ikev2.RequireCookie, nil)
				}
			}
		}
		entry.Protocol = nestedProtocol

		var nestedProtocolCommon *ProtocolCommon
		if o.ProtocolCommon != nil {
			nestedProtocolCommon = &ProtocolCommon{}
			if o.ProtocolCommon.Misc != nil {
				entry.Misc["ProtocolCommon"] = o.ProtocolCommon.Misc
			}
			if o.ProtocolCommon.Fragmentation != nil {
				nestedProtocolCommon.Fragmentation = &ProtocolCommonFragmentation{}
				if o.ProtocolCommon.Fragmentation.Misc != nil {
					entry.Misc["ProtocolCommonFragmentation"] = o.ProtocolCommon.Fragmentation.Misc
				}
				if o.ProtocolCommon.Fragmentation.Enable != nil {
					nestedProtocolCommon.Fragmentation.Enable = util.AsBool(o.ProtocolCommon.Fragmentation.Enable, nil)
				}
			}
			if o.ProtocolCommon.NatTraversal != nil {
				nestedProtocolCommon.NatTraversal = &ProtocolCommonNatTraversal{}
				if o.ProtocolCommon.NatTraversal.Misc != nil {
					entry.Misc["ProtocolCommonNatTraversal"] = o.ProtocolCommon.NatTraversal.Misc
				}
				if o.ProtocolCommon.NatTraversal.Enable != nil {
					nestedProtocolCommon.NatTraversal.Enable = util.AsBool(o.ProtocolCommon.NatTraversal.Enable, nil)
				}
				if o.ProtocolCommon.NatTraversal.KeepAliveInterval != nil {
					nestedProtocolCommon.NatTraversal.KeepAliveInterval = o.ProtocolCommon.NatTraversal.KeepAliveInterval
				}
				if o.ProtocolCommon.NatTraversal.UdpChecksumEnable != nil {
					nestedProtocolCommon.NatTraversal.UdpChecksumEnable = util.AsBool(o.ProtocolCommon.NatTraversal.UdpChecksumEnable, nil)
				}
			}
			if o.ProtocolCommon.PassiveMode != nil {
				nestedProtocolCommon.PassiveMode = util.AsBool(o.ProtocolCommon.PassiveMode, nil)
			}
		}
		entry.ProtocolCommon = nestedProtocolCommon

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
	if !matchAuthentication(a.Authentication, b.Authentication) {
		return false
	}
	if !util.StringsMatch(a.Comment, b.Comment) {
		return false
	}
	if !util.BoolsMatch(a.Disabled, b.Disabled) {
		return false
	}
	if !util.BoolsMatch(a.Ipv6, b.Ipv6) {
		return false
	}
	if !matchLocalAddress(a.LocalAddress, b.LocalAddress) {
		return false
	}
	if !matchLocalId(a.LocalId, b.LocalId) {
		return false
	}
	if !matchPeerAddress(a.PeerAddress, b.PeerAddress) {
		return false
	}
	if !matchPeerId(a.PeerId, b.PeerId) {
		return false
	}
	if !matchProtocol(a.Protocol, b.Protocol) {
		return false
	}
	if !matchProtocolCommon(a.ProtocolCommon, b.ProtocolCommon) {
		return false
	}

	return true
}

func matchAuthenticationCertificateLocalCertificateHashAndUrl(a *AuthenticationCertificateLocalCertificateHashAndUrl, b *AuthenticationCertificateLocalCertificateHashAndUrl) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.BaseUrl, b.BaseUrl) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	return true
}
func matchAuthenticationCertificateLocalCertificate(a *AuthenticationCertificateLocalCertificate, b *AuthenticationCertificateLocalCertificate) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchAuthenticationCertificateLocalCertificateHashAndUrl(a.HashAndUrl, b.HashAndUrl) {
		return false
	}
	if !util.StringsMatch(a.Name, b.Name) {
		return false
	}
	return true
}
func matchAuthenticationCertificate(a *AuthenticationCertificate, b *AuthenticationCertificate) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.CertificateProfile, b.CertificateProfile) {
		return false
	}
	if !matchAuthenticationCertificateLocalCertificate(a.LocalCertificate, b.LocalCertificate) {
		return false
	}
	if !util.BoolsMatch(a.StrictValidationRevocation, b.StrictValidationRevocation) {
		return false
	}
	if !util.BoolsMatch(a.UseManagementAsSource, b.UseManagementAsSource) {
		return false
	}
	if !util.BoolsMatch(a.AllowIdPayloadMismatch, b.AllowIdPayloadMismatch) {
		return false
	}
	return true
}
func matchAuthenticationPreSharedKey(a *AuthenticationPreSharedKey, b *AuthenticationPreSharedKey) bool {
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
func matchAuthentication(a *Authentication, b *Authentication) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchAuthenticationCertificate(a.Certificate, b.Certificate) {
		return false
	}
	if !matchAuthenticationPreSharedKey(a.PreSharedKey, b.PreSharedKey) {
		return false
	}
	return true
}
func matchLocalId(a *LocalId, b *LocalId) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Id, b.Id) {
		return false
	}
	if !util.StringsMatch(a.Type, b.Type) {
		return false
	}
	return true
}
func matchProtocolIkev1Dpd(a *ProtocolIkev1Dpd, b *ProtocolIkev1Dpd) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.Ints64Match(a.Interval, b.Interval) {
		return false
	}
	if !util.Ints64Match(a.Retry, b.Retry) {
		return false
	}
	return true
}
func matchProtocolIkev1(a *ProtocolIkev1, b *ProtocolIkev1) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolIkev1Dpd(a.Dpd, b.Dpd) {
		return false
	}
	if !util.StringsMatch(a.ExchangeMode, b.ExchangeMode) {
		return false
	}
	if !util.StringsMatch(a.IkeCryptoProfile, b.IkeCryptoProfile) {
		return false
	}
	return true
}
func matchProtocolIkev2Dpd(a *ProtocolIkev2Dpd, b *ProtocolIkev2Dpd) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.Ints64Match(a.Interval, b.Interval) {
		return false
	}
	return true
}
func matchProtocolIkev2(a *ProtocolIkev2, b *ProtocolIkev2) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolIkev2Dpd(a.Dpd, b.Dpd) {
		return false
	}
	if !util.StringsMatch(a.IkeCryptoProfile, b.IkeCryptoProfile) {
		return false
	}
	if !util.BoolsMatch(a.RequireCookie, b.RequireCookie) {
		return false
	}
	return true
}
func matchProtocol(a *Protocol, b *Protocol) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Version, b.Version) {
		return false
	}
	if !matchProtocolIkev1(a.Ikev1, b.Ikev1) {
		return false
	}
	if !matchProtocolIkev2(a.Ikev2, b.Ikev2) {
		return false
	}
	return true
}
func matchProtocolCommonFragmentation(a *ProtocolCommonFragmentation, b *ProtocolCommonFragmentation) bool {
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
func matchProtocolCommonNatTraversal(a *ProtocolCommonNatTraversal, b *ProtocolCommonNatTraversal) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.KeepAliveInterval, b.KeepAliveInterval) {
		return false
	}
	if !util.BoolsMatch(a.UdpChecksumEnable, b.UdpChecksumEnable) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	return true
}
func matchProtocolCommon(a *ProtocolCommon, b *ProtocolCommon) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolCommonNatTraversal(a.NatTraversal, b.NatTraversal) {
		return false
	}
	if !util.BoolsMatch(a.PassiveMode, b.PassiveMode) {
		return false
	}
	if !matchProtocolCommonFragmentation(a.Fragmentation, b.Fragmentation) {
		return false
	}
	return true
}
func matchLocalAddress(a *LocalAddress, b *LocalAddress) bool {
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
func matchPeerAddressDynamic(a *PeerAddressDynamic, b *PeerAddressDynamic) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchPeerAddress(a *PeerAddress, b *PeerAddress) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchPeerAddressDynamic(a.Dynamic, b.Dynamic) {
		return false
	}
	if !util.StringsMatch(a.Fqdn, b.Fqdn) {
		return false
	}
	if !util.StringsMatch(a.Ip, b.Ip) {
		return false
	}
	return true
}
func matchPeerId(a *PeerId, b *PeerId) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Matching, b.Matching) {
		return false
	}
	if !util.StringsMatch(a.Type, b.Type) {
		return false
	}
	if !util.StringsMatch(a.Id, b.Id) {
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
