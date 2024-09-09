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
	Autokey      *string
	None         *string
	SymmetricKey *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey struct {
	KeyId *int64
	Md5   *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyMd5
	Sha1  *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeySha1
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyMd5 struct {
	AuthenticationKey *string
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeySha1 struct {
	AuthenticationKey *string
}
type NtpServersSecondaryNtpServer struct {
	AuthenticationType *NtpServersSecondaryNtpServerAuthenticationType
	NtpServerAddress   *string
}
type NtpServersSecondaryNtpServerAuthenticationType struct {
	Autokey      *string
	None         *string
	SymmetricKey *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey struct {
	KeyId *int64
	Md5   *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyMd5
	Sha1  *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeySha1
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyMd5 struct {
	AuthenticationKey *string
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeySha1 struct {
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
	Autokey      *string                                                      `xml:"autokey,omitempty"`
	None         *string                                                      `xml:"none,omitempty"`
	SymmetricKey *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyXml `xml:"symmetric-key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyXml struct {
	KeyId *int64                                                           `xml:"key-id,omitempty"`
	Md5   *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyMd5Xml  `xml:"algorithm>md5,omitempty"`
	Sha1  *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeySha1Xml `xml:"algorithm>sha1,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyMd5Xml struct {
	AuthenticationKey *string `xml:"authentication-key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeySha1Xml struct {
	AuthenticationKey *string `xml:"authentication-key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersSecondaryNtpServerXml struct {
	AuthenticationType *NtpServersSecondaryNtpServerAuthenticationTypeXml `xml:"authentication-type,omitempty"`
	NtpServerAddress   *string                                            `xml:"ntp-server-address,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersSecondaryNtpServerAuthenticationTypeXml struct {
	Autokey      *string                                                        `xml:"autokey,omitempty"`
	None         *string                                                        `xml:"none,omitempty"`
	SymmetricKey *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyXml `xml:"symmetric-key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyXml struct {
	KeyId *int64                                                             `xml:"key-id,omitempty"`
	Md5   *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyMd5Xml  `xml:"algorithm>md5,omitempty"`
	Sha1  *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeySha1Xml `xml:"algorithm>sha1,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyMd5Xml struct {
	AuthenticationKey *string `xml:"authentication-key,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeySha1Xml struct {
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
			if o.NtpServers.PrimaryNtpServer.NtpServerAddress != nil {
				nestedNtpServers.PrimaryNtpServer.NtpServerAddress = o.NtpServers.PrimaryNtpServer.NtpServerAddress
			}
			if o.NtpServers.PrimaryNtpServer.AuthenticationType != nil {
				nestedNtpServers.PrimaryNtpServer.AuthenticationType = &NtpServersPrimaryNtpServerAuthenticationTypeXml{}
				if _, ok := o.Misc["NtpServersPrimaryNtpServerAuthenticationType"]; ok {
					nestedNtpServers.PrimaryNtpServer.AuthenticationType.Misc = o.Misc["NtpServersPrimaryNtpServerAuthenticationType"]
				}
				if o.NtpServers.PrimaryNtpServer.AuthenticationType.None != nil {
					nestedNtpServers.PrimaryNtpServer.AuthenticationType.None = o.NtpServers.PrimaryNtpServer.AuthenticationType.None
				}
				if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey != nil {
					nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey = &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyXml{}
					if _, ok := o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey"]; ok {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Misc = o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey"]
					}
					if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.KeyId != nil {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.KeyId = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.KeyId
					}
					if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Md5 != nil {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Md5 = &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyMd5Xml{}
						if _, ok := o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyMd5"]; ok {
							nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Md5.Misc = o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyMd5"]
						}
						if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Md5.AuthenticationKey != nil {
							nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Md5.AuthenticationKey = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Md5.AuthenticationKey
						}
					}
					if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Sha1 != nil {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Sha1 = &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeySha1Xml{}
						if _, ok := o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeySha1"]; ok {
							nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Sha1.Misc = o.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeySha1"]
						}
						if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Sha1.AuthenticationKey != nil {
							nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Sha1.AuthenticationKey = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Sha1.AuthenticationKey
						}
					}
				}
				if o.NtpServers.PrimaryNtpServer.AuthenticationType.Autokey != nil {
					nestedNtpServers.PrimaryNtpServer.AuthenticationType.Autokey = o.NtpServers.PrimaryNtpServer.AuthenticationType.Autokey
				}
			}
		}
		if o.NtpServers.SecondaryNtpServer != nil {
			nestedNtpServers.SecondaryNtpServer = &NtpServersSecondaryNtpServerXml{}
			if _, ok := o.Misc["NtpServersSecondaryNtpServer"]; ok {
				nestedNtpServers.SecondaryNtpServer.Misc = o.Misc["NtpServersSecondaryNtpServer"]
			}
			if o.NtpServers.SecondaryNtpServer.NtpServerAddress != nil {
				nestedNtpServers.SecondaryNtpServer.NtpServerAddress = o.NtpServers.SecondaryNtpServer.NtpServerAddress
			}
			if o.NtpServers.SecondaryNtpServer.AuthenticationType != nil {
				nestedNtpServers.SecondaryNtpServer.AuthenticationType = &NtpServersSecondaryNtpServerAuthenticationTypeXml{}
				if _, ok := o.Misc["NtpServersSecondaryNtpServerAuthenticationType"]; ok {
					nestedNtpServers.SecondaryNtpServer.AuthenticationType.Misc = o.Misc["NtpServersSecondaryNtpServerAuthenticationType"]
				}
				if o.NtpServers.SecondaryNtpServer.AuthenticationType.None != nil {
					nestedNtpServers.SecondaryNtpServer.AuthenticationType.None = o.NtpServers.SecondaryNtpServer.AuthenticationType.None
				}
				if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey != nil {
					nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey = &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyXml{}
					if _, ok := o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey"]; ok {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Misc = o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey"]
					}
					if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.KeyId != nil {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.KeyId = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.KeyId
					}
					if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Sha1 != nil {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Sha1 = &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeySha1Xml{}
						if _, ok := o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeySha1"]; ok {
							nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Sha1.Misc = o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeySha1"]
						}
						if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Sha1.AuthenticationKey != nil {
							nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Sha1.AuthenticationKey = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Sha1.AuthenticationKey
						}
					}
					if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Md5 != nil {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Md5 = &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyMd5Xml{}
						if _, ok := o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyMd5"]; ok {
							nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Md5.Misc = o.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyMd5"]
						}
						if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Md5.AuthenticationKey != nil {
							nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Md5.AuthenticationKey = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Md5.AuthenticationKey
						}
					}
				}
				if o.NtpServers.SecondaryNtpServer.AuthenticationType.Autokey != nil {
					nestedNtpServers.SecondaryNtpServer.AuthenticationType.Autokey = o.NtpServers.SecondaryNtpServer.AuthenticationType.Autokey
				}
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
				if o.NtpServers.PrimaryNtpServer.NtpServerAddress != nil {
					nestedNtpServers.PrimaryNtpServer.NtpServerAddress = o.NtpServers.PrimaryNtpServer.NtpServerAddress
				}
				if o.NtpServers.PrimaryNtpServer.AuthenticationType != nil {
					nestedNtpServers.PrimaryNtpServer.AuthenticationType = &NtpServersPrimaryNtpServerAuthenticationType{}
					if o.NtpServers.PrimaryNtpServer.AuthenticationType.Misc != nil {
						config.Misc["NtpServersPrimaryNtpServerAuthenticationType"] = o.NtpServers.PrimaryNtpServer.AuthenticationType.Misc
					}
					if o.NtpServers.PrimaryNtpServer.AuthenticationType.None != nil {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.None = o.NtpServers.PrimaryNtpServer.AuthenticationType.None
					}
					if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey != nil {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey = &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey{}
						if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Misc != nil {
							config.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey"] = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Misc
						}
						if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.KeyId != nil {
							nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.KeyId = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.KeyId
						}
						if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Sha1 != nil {
							nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Sha1 = &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeySha1{}
							if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Sha1.Misc != nil {
								config.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeySha1"] = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Sha1.Misc
							}
							if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Sha1.AuthenticationKey != nil {
								nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Sha1.AuthenticationKey = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Sha1.AuthenticationKey
							}
						}
						if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Md5 != nil {
							nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Md5 = &NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyMd5{}
							if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Md5.Misc != nil {
								config.Misc["NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyMd5"] = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Md5.Misc
							}
							if o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Md5.AuthenticationKey != nil {
								nestedNtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Md5.AuthenticationKey = o.NtpServers.PrimaryNtpServer.AuthenticationType.SymmetricKey.Md5.AuthenticationKey
							}
						}
					}
					if o.NtpServers.PrimaryNtpServer.AuthenticationType.Autokey != nil {
						nestedNtpServers.PrimaryNtpServer.AuthenticationType.Autokey = o.NtpServers.PrimaryNtpServer.AuthenticationType.Autokey
					}
				}
			}
			if o.NtpServers.SecondaryNtpServer != nil {
				nestedNtpServers.SecondaryNtpServer = &NtpServersSecondaryNtpServer{}
				if o.NtpServers.SecondaryNtpServer.Misc != nil {
					config.Misc["NtpServersSecondaryNtpServer"] = o.NtpServers.SecondaryNtpServer.Misc
				}
				if o.NtpServers.SecondaryNtpServer.NtpServerAddress != nil {
					nestedNtpServers.SecondaryNtpServer.NtpServerAddress = o.NtpServers.SecondaryNtpServer.NtpServerAddress
				}
				if o.NtpServers.SecondaryNtpServer.AuthenticationType != nil {
					nestedNtpServers.SecondaryNtpServer.AuthenticationType = &NtpServersSecondaryNtpServerAuthenticationType{}
					if o.NtpServers.SecondaryNtpServer.AuthenticationType.Misc != nil {
						config.Misc["NtpServersSecondaryNtpServerAuthenticationType"] = o.NtpServers.SecondaryNtpServer.AuthenticationType.Misc
					}
					if o.NtpServers.SecondaryNtpServer.AuthenticationType.None != nil {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.None = o.NtpServers.SecondaryNtpServer.AuthenticationType.None
					}
					if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey != nil {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey = &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey{}
						if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Misc != nil {
							config.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey"] = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Misc
						}
						if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.KeyId != nil {
							nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.KeyId = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.KeyId
						}
						if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Md5 != nil {
							nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Md5 = &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyMd5{}
							if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Md5.Misc != nil {
								config.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyMd5"] = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Md5.Misc
							}
							if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Md5.AuthenticationKey != nil {
								nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Md5.AuthenticationKey = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Md5.AuthenticationKey
							}
						}
						if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Sha1 != nil {
							nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Sha1 = &NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeySha1{}
							if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Sha1.Misc != nil {
								config.Misc["NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeySha1"] = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Sha1.Misc
							}
							if o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Sha1.AuthenticationKey != nil {
								nestedNtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Sha1.AuthenticationKey = o.NtpServers.SecondaryNtpServer.AuthenticationType.SymmetricKey.Sha1.AuthenticationKey
							}
						}
					}
					if o.NtpServers.SecondaryNtpServer.AuthenticationType.Autokey != nil {
						nestedNtpServers.SecondaryNtpServer.AuthenticationType.Autokey = o.NtpServers.SecondaryNtpServer.AuthenticationType.Autokey
					}
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

func matchNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeySha1(a *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeySha1, b *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeySha1) bool {
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
func matchNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyMd5(a *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyMd5, b *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyMd5) bool {
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
func matchNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey(a *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey, b *NtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.KeyId, b.KeyId) {
		return false
	}
	if !matchNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeySha1(a.Sha1, b.Sha1) {
		return false
	}
	if !matchNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKeyMd5(a.Md5, b.Md5) {
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
	if !util.StringsMatch(a.None, b.None) {
		return false
	}
	if !matchNtpServersPrimaryNtpServerAuthenticationTypeSymmetricKey(a.SymmetricKey, b.SymmetricKey) {
		return false
	}
	if !util.StringsMatch(a.Autokey, b.Autokey) {
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
	if !util.StringsMatch(a.NtpServerAddress, b.NtpServerAddress) {
		return false
	}
	if !matchNtpServersPrimaryNtpServerAuthenticationType(a.AuthenticationType, b.AuthenticationType) {
		return false
	}
	return true
}
func matchNtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyMd5(a *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyMd5, b *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyMd5) bool {
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
func matchNtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeySha1(a *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeySha1, b *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeySha1) bool {
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
func matchNtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey(a *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey, b *NtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.KeyId, b.KeyId) {
		return false
	}
	if !matchNtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeyMd5(a.Md5, b.Md5) {
		return false
	}
	if !matchNtpServersSecondaryNtpServerAuthenticationTypeSymmetricKeySha1(a.Sha1, b.Sha1) {
		return false
	}
	return true
}
func matchNtpServersSecondaryNtpServerAuthenticationType(a *NtpServersSecondaryNtpServerAuthenticationType, b *NtpServersSecondaryNtpServerAuthenticationType) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.None, b.None) {
		return false
	}
	if !matchNtpServersSecondaryNtpServerAuthenticationTypeSymmetricKey(a.SymmetricKey, b.SymmetricKey) {
		return false
	}
	if !util.StringsMatch(a.Autokey, b.Autokey) {
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
	if !util.StringsMatch(a.NtpServerAddress, b.NtpServerAddress) {
		return false
	}
	if !matchNtpServersSecondaryNtpServerAuthenticationType(a.AuthenticationType, b.AuthenticationType) {
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
