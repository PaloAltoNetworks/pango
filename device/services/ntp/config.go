package ntp

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/generic"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

type Config struct {
	NtpServers *NtpServers

	Misc map[string][]generic.Xml
}
type NtpServers struct {
	PrimaryNtpServer   *NtpServersPrimaryNtpServer
	SecondaryNtpServer *NtpServersSecondaryNtpServer
}
type NtpServersPrimaryNtpServer struct {
	AuthenticationType *NtpServersPrimaryNtpServerAuthenticationType
	NtpServerAddress   *string
}
type NtpServersPrimaryNtpServerAuthenticationType struct {
	Autokey      *NtpServersPrimaryNtpServerAuthenticationTypeAutokey
	None         *NtpServersPrimaryNtpServerAuthenticationTypeNone
	SymmetricKey *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
}
type NtpServersPrimaryNtpServerAuthenticationTypeAutokey struct {
}
type NtpServersPrimaryNtpServerAuthenticationTypeNone struct {
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey struct {
	Algorithm *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
	KeyId     *int64
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm struct {
	Md5  *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
	Sha1 *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5 struct {
	AuthenticationKey *string
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1 struct {
	AuthenticationKey *string
}
type NtpServersSecondaryNtpServer struct {
	AuthenticationType *NtpServersSecondaryNtpServerAuthenticationType
	NtpServerAddress   *string
}
type NtpServersSecondaryNtpServerAuthenticationType struct {
	Autokey      *NtpServersSecondaryNtpServerAuthenticationTypeAutokey
	None         *NtpServersSecondaryNtpServerAuthenticationTypeNone
	SymmetricKey *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey
}
type NtpServersSecondaryNtpServerAuthenticationTypeAutokey struct {
}
type NtpServersSecondaryNtpServerAuthenticationTypeNone struct {
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey struct {
	Algorithm *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm
	KeyId     *int64
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm struct {
	Md5  *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5
	Sha1 *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5 struct {
	AuthenticationKey *string
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1 struct {
	AuthenticationKey *string
}
type configXmlContainer struct {
	XMLName xml.Name    `xml:"result"`
	Answer  []configXml `xml:"system"`
}
type configXml struct {
	XMLName    xml.Name       `xml:"system"`
	NtpServers *NtpServersXml `xml:"ntp-servers,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersXml struct {
	PrimaryNtpServer   *NtpServersPrimaryNtpServerXml   `xml:"primary-ntp-server,omitempty"`
	SecondaryNtpServer *NtpServersSecondaryNtpServerXml `xml:"secondary-ntp-server,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersPrimaryNtpServerXml struct {
	AuthenticationType *NtpServersPrimaryNtpServerAuthenticationTypeXml `xml:"authentication-type,omitempty"`
	NtpServerAddress   *string                                          `xml:"ntp-server-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersPrimaryNtpServerAuthenticationTypeXml struct {
	Autokey      *NtpServersPrimaryNtpServerAuthenticationTypeAutokeyXml      `xml:"autokey,omitempty"`
	None         *NtpServersPrimaryNtpServerAuthenticationTypeNoneXml         `xml:"none,omitempty"`
	SymmetricKey *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyXml `xml:"symmetric-key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersPrimaryNtpServerAuthenticationTypeAutokeyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type NtpServersPrimaryNtpServerAuthenticationTypeNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyXml struct {
	Algorithm *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml `xml:"algorithm,omitempty"`
	KeyId     *int64                                                                `xml:"key-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml struct {
	Md5  *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml  `xml:"md5,omitempty"`
	Sha1 *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml `xml:"sha1,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml struct {
	AuthenticationKey *string `xml:"authentication-key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml struct {
	AuthenticationKey *string `xml:"authentication-key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersSecondaryNtpServerXml struct {
	AuthenticationType *NtpServersSecondaryNtpServerAuthenticationTypeXml `xml:"authentication-type,omitempty"`
	NtpServerAddress   *string                                            `xml:"ntp-server-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersSecondaryNtpServerAuthenticationTypeXml struct {
	Autokey      *NtpServersSecondaryNtpServerAuthenticationTypeAutokeyXml      `xml:"autokey,omitempty"`
	None         *NtpServersSecondaryNtpServerAuthenticationTypeNoneXml         `xml:"none,omitempty"`
	SymmetricKey *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyXml `xml:"symmetric-key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersSecondaryNtpServerAuthenticationTypeAutokeyXml struct {
	Misc []generic.Xml `xml:",any"`
}
type NtpServersSecondaryNtpServerAuthenticationTypeNoneXml struct {
	Misc []generic.Xml `xml:",any"`
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyXml struct {
	Algorithm *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml `xml:"algorithm,omitempty"`
	KeyId     *int64                                                                  `xml:"key-id,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml struct {
	Md5  *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml  `xml:"md5,omitempty"`
	Sha1 *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml `xml:"sha1,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml struct {
	AuthenticationKey *string `xml:"authentication-key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml struct {
	AuthenticationKey *string `xml:"authentication-key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	return specifyConfig, &configXmlContainer{}, nil
}
func specifyConfig(o *Config) (any, error) {
	config := configXml{}
	var nestedNtpServers *NtpServersXml
	if o.NtpServers != nil {
		nestedNtpServers = &NtpServersXml{}
		if _, ok := o.Misc["NtpServers"]; ok {
			nestedNtpServers.Misc = o.Misc["NtpServers"]
		}
		if o.NtpServers.PrimaryNtpServer != nil {
			nestedNtpServers.PrimaryNtpServer = &NtpServersPrimaryNtpServerXml{}
			if _, ok := o.Misc["NtpServersPrimaryNtpServer"]; ok {
				nestedNtpServers.PrimaryNtpServer.Misc = o.Misc["NtpServersPrimaryNtpServer"]
			}
			if o.NtpServers.PrimaryNtpServer.AuthenticationType != nil {
				nestedNtpServers.PrimaryNtpServer.AuthenticationType = &NtpServersPrimaryNtpServerAuthenticationTypeXml{}
				if _, ok := o.Misc["NtpServersPrimaryNtpServerAuthenticationType"]; ok {
					nestedNtpServers.PrimaryNtpServer.AuthenticationType.Misc = o.Misc["NtpServersPrimaryNtpServerAuthenticationType"]
				}
				if o.NtpServers.PrimaryNtpServer.AuthenticationType.Autokey != nil {
					nestedNtpServers.PrimaryNtpServer.AuthenticationType.Autokey = &NtpServersPrimaryNtpServerAuthenticationTypeAutokeyXml{}
					if _, ok := o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeAutokey"]; ok {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.Autokey.Misc = o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeAutokey"]
					}
				}
				if o.NtpServers.PrimaryNtpServer.AuthenticationType.None != nil {
					nestedNtpServers.PrimaryNtpServer.AuthenticationType.None = &NtpServersPrimaryNtpServerAuthenticationTypeNoneXml{}
					if _, ok := o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeNone"]; ok {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.None.Misc = o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeNone"]
					}
				}
				if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey != nil {
					nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey = &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyXml{}
					if _, ok := o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey"]; ok {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Misc = o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey"]
					}
					if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm != nil {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm = &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml{}
						if _, ok := o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm"]; ok {
							nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Misc = o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm"]
						}
						if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5 != nil {
							nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5 = &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml{}
							if _, ok := o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5"]; ok {
								nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.Misc = o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5"]
							}
							if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey != nil {
								nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey
							}
						}
						if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1 != nil {
							nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1 = &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml{}
							if _, ok := o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1"]; ok {
								nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.Misc = o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1"]
							}
							if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey != nil {
								nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey
							}
						}
					}
					if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.KeyId != nil {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.KeyId = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.KeyId
					}
				}
			}
			if o.NtpServers.PrimaryNtpServer.NtpServerAddress != nil {
				nestedNtpServers.PrimaryNtpServer.NtpServerAddress = o.NtpServers.PrimaryNtpServer.NtpServerAddress
			}
		}
		if o.NtpServers.SecondaryNtpServer != nil {
			nestedNtpServers.SecondaryNtpServer = &NtpServersSecondaryNtpServerXml{}
			if _, ok := o.Misc["NtpServersSecondaryNtpServer"]; ok {
				nestedNtpServers.SecondaryNtpServer.Misc = o.Misc["NtpServersSecondaryNtpServer"]
			}
			if o.NtpServers.SecondaryNtpServer.AuthenticationType != nil {
				nestedNtpServers.SecondaryNtpServer.AuthenticationType = &NtpServersSecondaryNtpServerAuthenticationTypeXml{}
				if _, ok := o.Misc["NtpServersSecondaryNtpServerAuthenticationType"]; ok {
					nestedNtpServers.SecondaryNtpServer.AuthenticationType.Misc = o.Misc["NtpServersSecondaryNtpServerAuthenticationType"]
				}
				if o.NtpServers.SecondaryNtpServer.AuthenticationType.Autokey != nil {
					nestedNtpServers.SecondaryNtpServer.AuthenticationType.Autokey = &NtpServersSecondaryNtpServerAuthenticationTypeAutokeyXml{}
					if _, ok := o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeAutokey"]; ok {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.Autokey.Misc = o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeAutokey"]
					}
				}
				if o.NtpServers.SecondaryNtpServer.AuthenticationType.None != nil {
					nestedNtpServers.SecondaryNtpServer.AuthenticationType.None = &NtpServersSecondaryNtpServerAuthenticationTypeNoneXml{}
					if _, ok := o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeNone"]; ok {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.None.Misc = o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeNone"]
					}
				}
				if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey != nil {
					nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey = &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyXml{}
					if _, ok := o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey"]; ok {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Misc = o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey"]
					}
					if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm != nil {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm = &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmXml{}
						if _, ok := o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm"]; ok {
							nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Misc = o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm"]
						}
						if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5 != nil {
							nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5 = &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5Xml{}
							if _, ok := o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5"]; ok {
								nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.Misc = o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5"]
							}
							if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey != nil {
								nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey
							}
						}
						if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1 != nil {
							nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1 = &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1Xml{}
							if _, ok := o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1"]; ok {
								nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.Misc = o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1"]
							}
							if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey != nil {
								nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey
							}
						}
					}
					if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.KeyId != nil {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.KeyId = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.KeyId
					}
				}
			}
			if o.NtpServers.SecondaryNtpServer.NtpServerAddress != nil {
				nestedNtpServers.SecondaryNtpServer.NtpServerAddress = o.NtpServers.SecondaryNtpServer.NtpServerAddress
			}
		}
	}
	config.NtpServers = nestedNtpServers

	config.Misc = o.Misc["Config"]

	return config, nil
}
func (c *configXmlContainer) Normalize() ([]*Config, error) {
	configList := make([]*Config, 0, len(c.Answer))
	for _, o := range c.Answer {
		config := &Config{
			Misc: make(map[string][]generic.Xml),
		}
		var nestedNtpServers *NtpServers
		if o.NtpServers != nil {
			nestedNtpServers = &NtpServers{}
			if o.NtpServers.Misc != nil {
				config.Misc["NtpServers"] = o.NtpServers.Misc
			}
			if o.NtpServers.PrimaryNtpServer != nil {
				nestedNtpServers.PrimaryNtpServer = &NtpServersPrimaryNtpServer{}
				if o.NtpServers.PrimaryNtpServer.Misc != nil {
					config.Misc["NtpServersPrimaryNtpServer"] = o.NtpServers.PrimaryNtpServer.Misc
				}
				if o.NtpServers.PrimaryNtpServer.AuthenticationType != nil {
					nestedNtpServers.PrimaryNtpServer.AuthenticationType = &NtpServersPrimaryNtpServerAuthenticationType{}
					if o.NtpServers.PrimaryNtpServer.AuthenticationType.Misc != nil {
						config.Misc["NtpServersPrimaryNtpServerAuthenticationType"] = o.NtpServers.PrimaryNtpServer.AuthenticationType.Misc
					}
					if o.NtpServers.PrimaryNtpServer.AuthenticationType.Autokey != nil {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.Autokey = &NtpServersPrimaryNtpServerAuthenticationTypeAutokey{}
						if o.NtpServers.PrimaryNtpServer.AuthenticationType.Autokey.Misc != nil {
							config.Misc["NtpServersPrimaryNtpServerAuthenticationTypeAutokey"] = o.NtpServers.PrimaryNtpServer.AuthenticationType.Autokey.Misc
						}
					}
					if o.NtpServers.PrimaryNtpServer.AuthenticationType.None != nil {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.None = &NtpServersPrimaryNtpServerAuthenticationTypeNone{}
						if o.NtpServers.PrimaryNtpServer.AuthenticationType.None.Misc != nil {
							config.Misc["NtpServersPrimaryNtpServerAuthenticationTypeNone"] = o.NtpServers.PrimaryNtpServer.AuthenticationType.None.Misc
						}
					}
					if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey != nil {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey = &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}
						if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Misc != nil {
							config.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey"] = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Misc
						}
						if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm != nil {
							nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm = &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}
							if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Misc != nil {
								config.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm"] = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Misc
							}
							if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5 != nil {
								nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5 = &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}
								if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.Misc != nil {
									config.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5"] = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.Misc
								}
								if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey != nil {
									nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey
								}
							}
							if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1 != nil {
								nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1 = &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1{}
								if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.Misc != nil {
									config.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1"] = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.Misc
								}
								if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey != nil {
									nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey
								}
							}
						}
						if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.KeyId != nil {
							nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.KeyId = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.KeyId
						}
					}
				}
				if o.NtpServers.PrimaryNtpServer.NtpServerAddress != nil {
					nestedNtpServers.PrimaryNtpServer.NtpServerAddress = o.NtpServers.PrimaryNtpServer.NtpServerAddress
				}
			}
			if o.NtpServers.SecondaryNtpServer != nil {
				nestedNtpServers.SecondaryNtpServer = &NtpServersSecondaryNtpServer{}
				if o.NtpServers.SecondaryNtpServer.Misc != nil {
					config.Misc["NtpServersSecondaryNtpServer"] = o.NtpServers.SecondaryNtpServer.Misc
				}
				if o.NtpServers.SecondaryNtpServer.AuthenticationType != nil {
					nestedNtpServers.SecondaryNtpServer.AuthenticationType = &NtpServersSecondaryNtpServerAuthenticationType{}
					if o.NtpServers.SecondaryNtpServer.AuthenticationType.Misc != nil {
						config.Misc["NtpServersSecondaryNtpServerAuthenticationType"] = o.NtpServers.SecondaryNtpServer.AuthenticationType.Misc
					}
					if o.NtpServers.SecondaryNtpServer.AuthenticationType.Autokey != nil {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.Autokey = &NtpServersSecondaryNtpServerAuthenticationTypeAutokey{}
						if o.NtpServers.SecondaryNtpServer.AuthenticationType.Autokey.Misc != nil {
							config.Misc["NtpServersSecondaryNtpServerAuthenticationTypeAutokey"] = o.NtpServers.SecondaryNtpServer.AuthenticationType.Autokey.Misc
						}
					}
					if o.NtpServers.SecondaryNtpServer.AuthenticationType.None != nil {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.None = &NtpServersSecondaryNtpServerAuthenticationTypeNone{}
						if o.NtpServers.SecondaryNtpServer.AuthenticationType.None.Misc != nil {
							config.Misc["NtpServersSecondaryNtpServerAuthenticationTypeNone"] = o.NtpServers.SecondaryNtpServer.AuthenticationType.None.Misc
						}
					}
					if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey != nil {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey = &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey{}
						if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Misc != nil {
							config.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey"] = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Misc
						}
						if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.KeyId != nil {
							nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.KeyId = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.KeyId
						}
						if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm != nil {
							nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm = &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm{}
							if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Misc != nil {
								config.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm"] = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Misc
							}
							if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5 != nil {
								nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5 = &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5{}
								if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.Misc != nil {
									config.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5"] = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.Misc
								}
								if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey != nil {
									nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Md5.AuthenticationKey
								}
							}
							if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1 != nil {
								nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1 = &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1{}
								if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.Misc != nil {
									config.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1"] = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.Misc
								}
								if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey != nil {
									nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Algorithm.Sha1.AuthenticationKey
								}
							}
						}
					}
				}
				if o.NtpServers.SecondaryNtpServer.NtpServerAddress != nil {
					nestedNtpServers.SecondaryNtpServer.NtpServerAddress = o.NtpServers.SecondaryNtpServer.NtpServerAddress
				}
			}
		}
		config.NtpServers = nestedNtpServers

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
	if !matchNtpServers(a.NtpServers, b.NtpServers) {
		return false
	}

	return true
}

func matchNtpServersPrimaryNtpServerAuthenticationTypeAutokey(a *NtpServersPrimaryNtpServerAuthenticationTypeAutokey, b *NtpServersPrimaryNtpServerAuthenticationTypeAutokey) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchNtpServersPrimaryNtpServerAuthenticationTypeNone(a *NtpServersPrimaryNtpServerAuthenticationTypeNone, b *NtpServersPrimaryNtpServerAuthenticationTypeNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5(a *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5, b *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.AuthenticationKey, b.AuthenticationKey) {
		return false
	}
	return true
}
func matchNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1(a *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1, b *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.AuthenticationKey, b.AuthenticationKey) {
		return false
	}
	return true
}
func matchNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm(a *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm, b *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5(a.Md5, b.Md5) {
		return false
	}
	if !matchNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1(a.Sha1, b.Sha1) {
		return false
	}
	return true
}
func matchNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey(a *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey, b *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm(a.Algorithm, b.Algorithm) {
		return false
	}
	if !util.Ints64Match(a.KeyId, b.KeyId) {
		return false
	}
	return true
}
func matchNtpServersPrimaryNtpServerAuthenticationType(a *NtpServersPrimaryNtpServerAuthenticationType, b *NtpServersPrimaryNtpServerAuthenticationType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchNtpServersPrimaryNtpServerAuthenticationTypeAutokey(a.Autokey, b.Autokey) {
		return false
	}
	if !matchNtpServersPrimaryNtpServerAuthenticationTypeNone(a.None, b.None) {
		return false
	}
	if !matchNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey(a.SymmetricKey, b.SymmetricKey) {
		return false
	}
	return true
}
func matchNtpServersPrimaryNtpServer(a *NtpServersPrimaryNtpServer, b *NtpServersPrimaryNtpServer) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchNtpServersPrimaryNtpServerAuthenticationType(a.AuthenticationType, b.AuthenticationType) {
		return false
	}
	if !util.StringsMatch(a.NtpServerAddress, b.NtpServerAddress) {
		return false
	}
	return true
}
func matchNtpServersSecondaryNtpServerAuthenticationTypeNone(a *NtpServersSecondaryNtpServerAuthenticationTypeNone, b *NtpServersSecondaryNtpServerAuthenticationTypeNone) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchNtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5(a *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5, b *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.AuthenticationKey, b.AuthenticationKey) {
		return false
	}
	return true
}
func matchNtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1(a *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1, b *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.AuthenticationKey, b.AuthenticationKey) {
		return false
	}
	return true
}
func matchNtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm(a *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm, b *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchNtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmMd5(a.Md5, b.Md5) {
		return false
	}
	if !matchNtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithmSha1(a.Sha1, b.Sha1) {
		return false
	}
	return true
}
func matchNtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey(a *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey, b *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchNtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyAlgorithm(a.Algorithm, b.Algorithm) {
		return false
	}
	if !util.Ints64Match(a.KeyId, b.KeyId) {
		return false
	}
	return true
}
func matchNtpServersSecondaryNtpServerAuthenticationTypeAutokey(a *NtpServersSecondaryNtpServerAuthenticationTypeAutokey, b *NtpServersSecondaryNtpServerAuthenticationTypeAutokey) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchNtpServersSecondaryNtpServerAuthenticationType(a *NtpServersSecondaryNtpServerAuthenticationType, b *NtpServersSecondaryNtpServerAuthenticationType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchNtpServersSecondaryNtpServerAuthenticationTypeAutokey(a.Autokey, b.Autokey) {
		return false
	}
	if !matchNtpServersSecondaryNtpServerAuthenticationTypeNone(a.None, b.None) {
		return false
	}
	if !matchNtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey(a.SymmetricKey, b.SymmetricKey) {
		return false
	}
	return true
}
func matchNtpServersSecondaryNtpServer(a *NtpServersSecondaryNtpServer, b *NtpServersSecondaryNtpServer) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchNtpServersSecondaryNtpServerAuthenticationType(a.AuthenticationType, b.AuthenticationType) {
		return false
	}
	if !util.StringsMatch(a.NtpServerAddress, b.NtpServerAddress) {
		return false
	}
	return true
}
func matchNtpServers(a *NtpServers, b *NtpServers) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchNtpServersPrimaryNtpServer(a.PrimaryNtpServer, b.PrimaryNtpServer) {
		return false
	}
	if !matchNtpServersSecondaryNtpServer(a.SecondaryNtpServer, b.SecondaryNtpServer) {
		return false
	}
	return true
}
