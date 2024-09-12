package virtual_router

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
	Suffix = []string{"network", "virtual-router"}
)

type Entry struct {
	Name                    string
	AdministrativeDistances *AdministrativeDistances
	Ecmp                    *Ecmp
	Interfaces              []string
	Protocol                *Protocol
	RoutingTable            *RoutingTable

	Misc map[string][]generic.Xml
}

type AdministrativeDistances struct {
	Ebgp       *int64
	Ibgp       *int64
	OspfExt    *int64
	OspfInt    *int64
	Ospfv3Ext  *int64
	Ospfv3Int  *int64
	Rip        *int64
	Static     *int64
	StaticIpv6 *int64
}
type Ecmp struct {
	Algorithm        *EcmpAlgorithm
	Enable           *bool
	MaxPaths         *int64
	StrictSourcePath *bool
	SymmetricReturn  *bool
}
type EcmpAlgorithm struct {
	BalancedRoundRobin *EcmpAlgorithmBalancedRoundRobin
	IpHash             *EcmpAlgorithmIpHash
	IpModulo           *EcmpAlgorithmIpModulo
	WeightedRoundRobin *EcmpAlgorithmWeightedRoundRobin
}
type EcmpAlgorithmBalancedRoundRobin struct {
}
type EcmpAlgorithmIpHash struct {
	HashSeed *int64
	SrcOnly  *bool
	UsePort  *bool
}
type EcmpAlgorithmIpModulo struct {
}
type EcmpAlgorithmWeightedRoundRobin struct {
	Interfaces []EcmpAlgorithmWeightedRoundRobinInterfaces
}
type EcmpAlgorithmWeightedRoundRobinInterfaces struct {
	Name   string
	Weight *int64
}
type Protocol struct {
	Bgp    *ProtocolBgp
	Ospf   *ProtocolOspf
	Ospfv3 *ProtocolOspfv3
	Rip    *ProtocolRip
}
type ProtocolBgp struct {
	Enable *bool
}
type ProtocolOspf struct {
	Enable *bool
}
type ProtocolOspfv3 struct {
	Enable *bool
}
type ProtocolRip struct {
	Enable *bool
}
type RoutingTable struct {
	Ip   *RoutingTableIp
	Ipv6 *RoutingTableIpv6
}
type RoutingTableIp struct {
	StaticRoutes []RoutingTableIpStaticRoutes
}
type RoutingTableIpStaticRoutes struct {
	AdminDist   *int64
	Destination *string
	Interface   *string
	Metric      *int64
	Name        string
	NextHop     *RoutingTableIpStaticRoutesNextHop
	RouteTable  *string
}
type RoutingTableIpStaticRoutesNextHop struct {
	Fqdn      *string
	IpAddress *string
	NextVr    *string
	Tunnel    *string
}
type RoutingTableIpv6 struct {
	StaticRoutes []RoutingTableIpv6StaticRoutes
}
type RoutingTableIpv6StaticRoutes struct {
	AdminDist   *int64
	Destination *string
	Interface   *string
	Metric      *int64
	Name        string
	NextHop     *RoutingTableIpv6StaticRoutesNextHop
	RouteTable  *string
}
type RoutingTableIpv6StaticRoutesNextHop struct {
	Fqdn        *string
	Ipv6Address *string
	NextVr      *string
	Tunnel      *string
}

type entryXmlContainer struct {
	Answer []entryXml `xml:"entry"`
}

type entryXml struct {
	XMLName                 xml.Name                    `xml:"entry"`
	Name                    string                      `xml:"name,attr"`
	AdministrativeDistances *AdministrativeDistancesXml `xml:"admin-dists,omitempty"`
	Ecmp                    *EcmpXml                    `xml:"ecmp,omitempty"`
	Interfaces              *util.MemberType            `xml:"interface,omitempty"`
	Protocol                *ProtocolXml                `xml:"protocol,omitempty"`
	RoutingTable            *RoutingTableXml            `xml:"routing-table,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

type AdministrativeDistancesXml struct {
	Ebgp       *int64 `xml:"ebgp,omitempty"`
	Ibgp       *int64 `xml:"ibgp,omitempty"`
	OspfExt    *int64 `xml:"ospf-ext,omitempty"`
	OspfInt    *int64 `xml:"ospf-int,omitempty"`
	Ospfv3Ext  *int64 `xml:"ospfv3-ext,omitempty"`
	Ospfv3Int  *int64 `xml:"ospfv3-int,omitempty"`
	Rip        *int64 `xml:"rip,omitempty"`
	Static     *int64 `xml:"static,omitempty"`
	StaticIpv6 *int64 `xml:"static-ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type EcmpXml struct {
	Algorithm        *EcmpAlgorithmXml `xml:"algorithm,omitempty"`
	Enable           *string           `xml:"enable,omitempty"`
	MaxPaths         *int64            `xml:"max-path,omitempty"`
	StrictSourcePath *string           `xml:"strict-source-path,omitempty"`
	SymmetricReturn  *string           `xml:"symmetric-return,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type EcmpAlgorithmXml struct {
	BalancedRoundRobin *EcmpAlgorithmBalancedRoundRobinXml `xml:"balanced-round-robin,omitempty"`
	IpHash             *EcmpAlgorithmIpHashXml             `xml:"ip-hash,omitempty"`
	IpModulo           *EcmpAlgorithmIpModuloXml           `xml:"ip-modulo,omitempty"`
	WeightedRoundRobin *EcmpAlgorithmWeightedRoundRobinXml `xml:"weighted-round-robin,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type EcmpAlgorithmBalancedRoundRobinXml struct {
	Misc []generic.Xml `xml:",any"`
}
type EcmpAlgorithmIpHashXml struct {
	HashSeed *int64  `xml:"hash-seed,omitempty"`
	SrcOnly  *string `xml:"src-only,omitempty"`
	UsePort  *string `xml:"use-port,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type EcmpAlgorithmIpModuloXml struct {
	Misc []generic.Xml `xml:",any"`
}
type EcmpAlgorithmWeightedRoundRobinXml struct {
	Interfaces []EcmpAlgorithmWeightedRoundRobinInterfacesXml `xml:"interface>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type EcmpAlgorithmWeightedRoundRobinInterfacesXml struct {
	XMLName xml.Name `xml:"entry"`
	Name    string   `xml:"name,attr"`
	Weight  *int64

	Misc []generic.Xml `xml:",any"`
}
type ProtocolXml struct {
	Bgp    *ProtocolBgpXml    `xml:"bgp,omitempty"`
	Ospf   *ProtocolOspfXml   `xml:"ospf,omitempty"`
	Ospfv3 *ProtocolOspfv3Xml `xml:"ospfv3,omitempty"`
	Rip    *ProtocolRipXml    `xml:"rip,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolBgpXml struct {
	Enable *string `xml:"enable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfXml struct {
	Enable *string `xml:"enable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolOspfv3Xml struct {
	Enable *string `xml:"enable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type ProtocolRipXml struct {
	Enable *string `xml:"enable,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableXml struct {
	Ip   *RoutingTableIpXml   `xml:"ip,omitempty"`
	Ipv6 *RoutingTableIpv6Xml `xml:"ipv6,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpXml struct {
	StaticRoutes []RoutingTableIpStaticRoutesXml `xml:"static-route>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpStaticRoutesXml struct {
	AdminDist   *int64                                `xml:"admin-dist,omitempty"`
	Destination *string                               `xml:"destination,omitempty"`
	Interface   *string                               `xml:"interface,omitempty"`
	Metric      *int64                                `xml:"metric,omitempty"`
	XMLName     xml.Name                              `xml:"entry"`
	Name        string                                `xml:"name,attr"`
	NextHop     *RoutingTableIpStaticRoutesNextHopXml `xml:"nexthop,omitempty"`
	RouteTable  *string                               `xml:"route-table,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpStaticRoutesNextHopXml struct {
	Fqdn      *string `xml:"fqdn,omitempty"`
	IpAddress *string `xml:"ip-address,omitempty"`
	NextVr    *string `xml:"next-vr,omitempty"`
	Tunnel    *string `xml:"tunnel,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6Xml struct {
	StaticRoutes []RoutingTableIpv6StaticRoutesXml `xml:"static-route>entry,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6StaticRoutesXml struct {
	AdminDist   *int64                                  `xml:"admin-dist,omitempty"`
	Destination *string                                 `xml:"destination,omitempty"`
	Interface   *string                                 `xml:"interface,omitempty"`
	Metric      *int64                                  `xml:"metric,omitempty"`
	XMLName     xml.Name                                `xml:"entry"`
	Name        string                                  `xml:"name,attr"`
	NextHop     *RoutingTableIpv6StaticRoutesNextHopXml `xml:"nexthop,omitempty"`
	RouteTable  *string                                 `xml:"route-table,omitempty"`

	Misc []generic.Xml `xml:",any"`
}
type RoutingTableIpv6StaticRoutesNextHopXml struct {
	Fqdn        *string `xml:"fqdn,omitempty"`
	Ipv6Address *string `xml:"ipv6-address,omitempty"`
	NextVr      *string `xml:"next-vr,omitempty"`
	Tunnel      *string `xml:"tunnel,omitempty"`

	Misc []generic.Xml `xml:",any"`
}

func (e *Entry) Field(v string) (any, error) {
	if v == "name" || v == "Name" {
		return e.Name, nil
	}
	if v == "administrative_distances" || v == "AdministrativeDistances" {
		return e.AdministrativeDistances, nil
	}
	if v == "ecmp" || v == "Ecmp" {
		return e.Ecmp, nil
	}
	if v == "interfaces" || v == "Interfaces" {
		return e.Interfaces, nil
	}
	if v == "interfaces|LENGTH" || v == "Interfaces|LENGTH" {
		return int64(len(e.Interfaces)), nil
	}
	if v == "protocol" || v == "Protocol" {
		return e.Protocol, nil
	}
	if v == "routing_table" || v == "RoutingTable" {
		return e.RoutingTable, nil
	}

	return nil, fmt.Errorf("unknown field")
}

func Versioning(vn version.Number) (Specifier, Normalizer, error) {
	return specifyEntry, &entryXmlContainer{}, nil
}

func specifyEntry(o *Entry) (any, error) {
	entry := entryXml{}

	entry.Name = o.Name
	var nestedAdministrativeDistances *AdministrativeDistancesXml
	if o.AdministrativeDistances != nil {
		nestedAdministrativeDistances = &AdministrativeDistancesXml{}
		if _, ok := o.Misc["AdministrativeDistances"]; ok {
			nestedAdministrativeDistances.Misc = o.Misc["AdministrativeDistances"]
		}
		if o.AdministrativeDistances.OspfExt != nil {
			nestedAdministrativeDistances.OspfExt = o.AdministrativeDistances.OspfExt
		}
		if o.AdministrativeDistances.Ospfv3Int != nil {
			nestedAdministrativeDistances.Ospfv3Int = o.AdministrativeDistances.Ospfv3Int
		}
		if o.AdministrativeDistances.Ospfv3Ext != nil {
			nestedAdministrativeDistances.Ospfv3Ext = o.AdministrativeDistances.Ospfv3Ext
		}
		if o.AdministrativeDistances.Ibgp != nil {
			nestedAdministrativeDistances.Ibgp = o.AdministrativeDistances.Ibgp
		}
		if o.AdministrativeDistances.Ebgp != nil {
			nestedAdministrativeDistances.Ebgp = o.AdministrativeDistances.Ebgp
		}
		if o.AdministrativeDistances.Rip != nil {
			nestedAdministrativeDistances.Rip = o.AdministrativeDistances.Rip
		}
		if o.AdministrativeDistances.StaticIpv6 != nil {
			nestedAdministrativeDistances.StaticIpv6 = o.AdministrativeDistances.StaticIpv6
		}
		if o.AdministrativeDistances.OspfInt != nil {
			nestedAdministrativeDistances.OspfInt = o.AdministrativeDistances.OspfInt
		}
		if o.AdministrativeDistances.Static != nil {
			nestedAdministrativeDistances.Static = o.AdministrativeDistances.Static
		}
	}
	entry.AdministrativeDistances = nestedAdministrativeDistances

	var nestedEcmp *EcmpXml
	if o.Ecmp != nil {
		nestedEcmp = &EcmpXml{}
		if _, ok := o.Misc["Ecmp"]; ok {
			nestedEcmp.Misc = o.Misc["Ecmp"]
		}
		if o.Ecmp.MaxPaths != nil {
			nestedEcmp.MaxPaths = o.Ecmp.MaxPaths
		}
		if o.Ecmp.Algorithm != nil {
			nestedEcmp.Algorithm = &EcmpAlgorithmXml{}
			if _, ok := o.Misc["EcmpAlgorithm"]; ok {
				nestedEcmp.Algorithm.Misc = o.Misc["EcmpAlgorithm"]
			}
			if o.Ecmp.Algorithm.WeightedRoundRobin != nil {
				nestedEcmp.Algorithm.WeightedRoundRobin = &EcmpAlgorithmWeightedRoundRobinXml{}
				if _, ok := o.Misc["EcmpAlgorithmWeightedRoundRobin"]; ok {
					nestedEcmp.Algorithm.WeightedRoundRobin.Misc = o.Misc["EcmpAlgorithmWeightedRoundRobin"]
				}
				if o.Ecmp.Algorithm.WeightedRoundRobin.Interfaces != nil {
					nestedEcmp.Algorithm.WeightedRoundRobin.Interfaces = []EcmpAlgorithmWeightedRoundRobinInterfacesXml{}
					for _, oEcmpAlgorithmWeightedRoundRobinInterfaces := range o.Ecmp.Algorithm.WeightedRoundRobin.Interfaces {
						nestedEcmpAlgorithmWeightedRoundRobinInterfaces := EcmpAlgorithmWeightedRoundRobinInterfacesXml{}
						if _, ok := o.Misc["EcmpAlgorithmWeightedRoundRobinInterfaces"]; ok {
							nestedEcmpAlgorithmWeightedRoundRobinInterfaces.Misc = o.Misc["EcmpAlgorithmWeightedRoundRobinInterfaces"]
						}
						if oEcmpAlgorithmWeightedRoundRobinInterfaces.Weight != nil {
							nestedEcmpAlgorithmWeightedRoundRobinInterfaces.Weight = oEcmpAlgorithmWeightedRoundRobinInterfaces.Weight
						}
						if oEcmpAlgorithmWeightedRoundRobinInterfaces.Name != "" {
							nestedEcmpAlgorithmWeightedRoundRobinInterfaces.Name = oEcmpAlgorithmWeightedRoundRobinInterfaces.Name
						}
						nestedEcmp.Algorithm.WeightedRoundRobin.Interfaces = append(nestedEcmp.Algorithm.WeightedRoundRobin.Interfaces, nestedEcmpAlgorithmWeightedRoundRobinInterfaces)
					}
				}
			}
			if o.Ecmp.Algorithm.BalancedRoundRobin != nil {
				nestedEcmp.Algorithm.BalancedRoundRobin = &EcmpAlgorithmBalancedRoundRobinXml{}
				if _, ok := o.Misc["EcmpAlgorithmBalancedRoundRobin"]; ok {
					nestedEcmp.Algorithm.BalancedRoundRobin.Misc = o.Misc["EcmpAlgorithmBalancedRoundRobin"]
				}
			}
			if o.Ecmp.Algorithm.IpModulo != nil {
				nestedEcmp.Algorithm.IpModulo = &EcmpAlgorithmIpModuloXml{}
				if _, ok := o.Misc["EcmpAlgorithmIpModulo"]; ok {
					nestedEcmp.Algorithm.IpModulo.Misc = o.Misc["EcmpAlgorithmIpModulo"]
				}
			}
			if o.Ecmp.Algorithm.IpHash != nil {
				nestedEcmp.Algorithm.IpHash = &EcmpAlgorithmIpHashXml{}
				if _, ok := o.Misc["EcmpAlgorithmIpHash"]; ok {
					nestedEcmp.Algorithm.IpHash.Misc = o.Misc["EcmpAlgorithmIpHash"]
				}
				if o.Ecmp.Algorithm.IpHash.SrcOnly != nil {
					nestedEcmp.Algorithm.IpHash.SrcOnly = util.YesNo(o.Ecmp.Algorithm.IpHash.SrcOnly, nil)
				}
				if o.Ecmp.Algorithm.IpHash.UsePort != nil {
					nestedEcmp.Algorithm.IpHash.UsePort = util.YesNo(o.Ecmp.Algorithm.IpHash.UsePort, nil)
				}
				if o.Ecmp.Algorithm.IpHash.HashSeed != nil {
					nestedEcmp.Algorithm.IpHash.HashSeed = o.Ecmp.Algorithm.IpHash.HashSeed
				}
			}
		}
		if o.Ecmp.Enable != nil {
			nestedEcmp.Enable = util.YesNo(o.Ecmp.Enable, nil)
		}
		if o.Ecmp.SymmetricReturn != nil {
			nestedEcmp.SymmetricReturn = util.YesNo(o.Ecmp.SymmetricReturn, nil)
		}
		if o.Ecmp.StrictSourcePath != nil {
			nestedEcmp.StrictSourcePath = util.YesNo(o.Ecmp.StrictSourcePath, nil)
		}
	}
	entry.Ecmp = nestedEcmp

	entry.Interfaces = util.StrToMem(o.Interfaces)
	var nestedProtocol *ProtocolXml
	if o.Protocol != nil {
		nestedProtocol = &ProtocolXml{}
		if _, ok := o.Misc["Protocol"]; ok {
			nestedProtocol.Misc = o.Misc["Protocol"]
		}
		if o.Protocol.Bgp != nil {
			nestedProtocol.Bgp = &ProtocolBgpXml{}
			if _, ok := o.Misc["ProtocolBgp"]; ok {
				nestedProtocol.Bgp.Misc = o.Misc["ProtocolBgp"]
			}
			if o.Protocol.Bgp.Enable != nil {
				nestedProtocol.Bgp.Enable = util.YesNo(o.Protocol.Bgp.Enable, nil)
			}
		}
		if o.Protocol.Rip != nil {
			nestedProtocol.Rip = &ProtocolRipXml{}
			if _, ok := o.Misc["ProtocolRip"]; ok {
				nestedProtocol.Rip.Misc = o.Misc["ProtocolRip"]
			}
			if o.Protocol.Rip.Enable != nil {
				nestedProtocol.Rip.Enable = util.YesNo(o.Protocol.Rip.Enable, nil)
			}
		}
		if o.Protocol.Ospf != nil {
			nestedProtocol.Ospf = &ProtocolOspfXml{}
			if _, ok := o.Misc["ProtocolOspf"]; ok {
				nestedProtocol.Ospf.Misc = o.Misc["ProtocolOspf"]
			}
			if o.Protocol.Ospf.Enable != nil {
				nestedProtocol.Ospf.Enable = util.YesNo(o.Protocol.Ospf.Enable, nil)
			}
		}
		if o.Protocol.Ospfv3 != nil {
			nestedProtocol.Ospfv3 = &ProtocolOspfv3Xml{}
			if _, ok := o.Misc["ProtocolOspfv3"]; ok {
				nestedProtocol.Ospfv3.Misc = o.Misc["ProtocolOspfv3"]
			}
			if o.Protocol.Ospfv3.Enable != nil {
				nestedProtocol.Ospfv3.Enable = util.YesNo(o.Protocol.Ospfv3.Enable, nil)
			}
		}
	}
	entry.Protocol = nestedProtocol

	var nestedRoutingTable *RoutingTableXml
	if o.RoutingTable != nil {
		nestedRoutingTable = &RoutingTableXml{}
		if _, ok := o.Misc["RoutingTable"]; ok {
			nestedRoutingTable.Misc = o.Misc["RoutingTable"]
		}
		if o.RoutingTable.Ip != nil {
			nestedRoutingTable.Ip = &RoutingTableIpXml{}
			if _, ok := o.Misc["RoutingTableIp"]; ok {
				nestedRoutingTable.Ip.Misc = o.Misc["RoutingTableIp"]
			}
			if o.RoutingTable.Ip.StaticRoutes != nil {
				nestedRoutingTable.Ip.StaticRoutes = []RoutingTableIpStaticRoutesXml{}
				for _, oRoutingTableIpStaticRoutes := range o.RoutingTable.Ip.StaticRoutes {
					nestedRoutingTableIpStaticRoutes := RoutingTableIpStaticRoutesXml{}
					if _, ok := o.Misc["RoutingTableIpStaticRoutes"]; ok {
						nestedRoutingTableIpStaticRoutes.Misc = o.Misc["RoutingTableIpStaticRoutes"]
					}
					if oRoutingTableIpStaticRoutes.RouteTable != nil {
						nestedRoutingTableIpStaticRoutes.RouteTable = oRoutingTableIpStaticRoutes.RouteTable
					}
					if oRoutingTableIpStaticRoutes.Name != "" {
						nestedRoutingTableIpStaticRoutes.Name = oRoutingTableIpStaticRoutes.Name
					}
					if oRoutingTableIpStaticRoutes.Destination != nil {
						nestedRoutingTableIpStaticRoutes.Destination = oRoutingTableIpStaticRoutes.Destination
					}
					if oRoutingTableIpStaticRoutes.Interface != nil {
						nestedRoutingTableIpStaticRoutes.Interface = oRoutingTableIpStaticRoutes.Interface
					}
					if oRoutingTableIpStaticRoutes.NextHop != nil {
						nestedRoutingTableIpStaticRoutes.NextHop = &RoutingTableIpStaticRoutesNextHopXml{}
						if _, ok := o.Misc["RoutingTableIpStaticRoutesNextHop"]; ok {
							nestedRoutingTableIpStaticRoutes.NextHop.Misc = o.Misc["RoutingTableIpStaticRoutesNextHop"]
						}
						if oRoutingTableIpStaticRoutes.NextHop.IpAddress != nil {
							nestedRoutingTableIpStaticRoutes.NextHop.IpAddress = oRoutingTableIpStaticRoutes.NextHop.IpAddress
						}
						if oRoutingTableIpStaticRoutes.NextHop.Fqdn != nil {
							nestedRoutingTableIpStaticRoutes.NextHop.Fqdn = oRoutingTableIpStaticRoutes.NextHop.Fqdn
						}
						if oRoutingTableIpStaticRoutes.NextHop.NextVr != nil {
							nestedRoutingTableIpStaticRoutes.NextHop.NextVr = oRoutingTableIpStaticRoutes.NextHop.NextVr
						}
						if oRoutingTableIpStaticRoutes.NextHop.Tunnel != nil {
							nestedRoutingTableIpStaticRoutes.NextHop.Tunnel = oRoutingTableIpStaticRoutes.NextHop.Tunnel
						}
					}
					if oRoutingTableIpStaticRoutes.AdminDist != nil {
						nestedRoutingTableIpStaticRoutes.AdminDist = oRoutingTableIpStaticRoutes.AdminDist
					}
					if oRoutingTableIpStaticRoutes.Metric != nil {
						nestedRoutingTableIpStaticRoutes.Metric = oRoutingTableIpStaticRoutes.Metric
					}
					nestedRoutingTable.Ip.StaticRoutes = append(nestedRoutingTable.Ip.StaticRoutes, nestedRoutingTableIpStaticRoutes)
				}
			}
		}
		if o.RoutingTable.Ipv6 != nil {
			nestedRoutingTable.Ipv6 = &RoutingTableIpv6Xml{}
			if _, ok := o.Misc["RoutingTableIpv6"]; ok {
				nestedRoutingTable.Ipv6.Misc = o.Misc["RoutingTableIpv6"]
			}
			if o.RoutingTable.Ipv6.StaticRoutes != nil {
				nestedRoutingTable.Ipv6.StaticRoutes = []RoutingTableIpv6StaticRoutesXml{}
				for _, oRoutingTableIpv6StaticRoutes := range o.RoutingTable.Ipv6.StaticRoutes {
					nestedRoutingTableIpv6StaticRoutes := RoutingTableIpv6StaticRoutesXml{}
					if _, ok := o.Misc["RoutingTableIpv6StaticRoutes"]; ok {
						nestedRoutingTableIpv6StaticRoutes.Misc = o.Misc["RoutingTableIpv6StaticRoutes"]
					}
					if oRoutingTableIpv6StaticRoutes.RouteTable != nil {
						nestedRoutingTableIpv6StaticRoutes.RouteTable = oRoutingTableIpv6StaticRoutes.RouteTable
					}
					if oRoutingTableIpv6StaticRoutes.Name != "" {
						nestedRoutingTableIpv6StaticRoutes.Name = oRoutingTableIpv6StaticRoutes.Name
					}
					if oRoutingTableIpv6StaticRoutes.Destination != nil {
						nestedRoutingTableIpv6StaticRoutes.Destination = oRoutingTableIpv6StaticRoutes.Destination
					}
					if oRoutingTableIpv6StaticRoutes.Interface != nil {
						nestedRoutingTableIpv6StaticRoutes.Interface = oRoutingTableIpv6StaticRoutes.Interface
					}
					if oRoutingTableIpv6StaticRoutes.NextHop != nil {
						nestedRoutingTableIpv6StaticRoutes.NextHop = &RoutingTableIpv6StaticRoutesNextHopXml{}
						if _, ok := o.Misc["RoutingTableIpv6StaticRoutesNextHop"]; ok {
							nestedRoutingTableIpv6StaticRoutes.NextHop.Misc = o.Misc["RoutingTableIpv6StaticRoutesNextHop"]
						}
						if oRoutingTableIpv6StaticRoutes.NextHop.Fqdn != nil {
							nestedRoutingTableIpv6StaticRoutes.NextHop.Fqdn = oRoutingTableIpv6StaticRoutes.NextHop.Fqdn
						}
						if oRoutingTableIpv6StaticRoutes.NextHop.NextVr != nil {
							nestedRoutingTableIpv6StaticRoutes.NextHop.NextVr = oRoutingTableIpv6StaticRoutes.NextHop.NextVr
						}
						if oRoutingTableIpv6StaticRoutes.NextHop.Tunnel != nil {
							nestedRoutingTableIpv6StaticRoutes.NextHop.Tunnel = oRoutingTableIpv6StaticRoutes.NextHop.Tunnel
						}
						if oRoutingTableIpv6StaticRoutes.NextHop.Ipv6Address != nil {
							nestedRoutingTableIpv6StaticRoutes.NextHop.Ipv6Address = oRoutingTableIpv6StaticRoutes.NextHop.Ipv6Address
						}
					}
					if oRoutingTableIpv6StaticRoutes.AdminDist != nil {
						nestedRoutingTableIpv6StaticRoutes.AdminDist = oRoutingTableIpv6StaticRoutes.AdminDist
					}
					if oRoutingTableIpv6StaticRoutes.Metric != nil {
						nestedRoutingTableIpv6StaticRoutes.Metric = oRoutingTableIpv6StaticRoutes.Metric
					}
					nestedRoutingTable.Ipv6.StaticRoutes = append(nestedRoutingTable.Ipv6.StaticRoutes, nestedRoutingTableIpv6StaticRoutes)
				}
			}
		}
	}
	entry.RoutingTable = nestedRoutingTable

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
		var nestedAdministrativeDistances *AdministrativeDistances
		if o.AdministrativeDistances != nil {
			nestedAdministrativeDistances = &AdministrativeDistances{}
			if o.AdministrativeDistances.Misc != nil {
				entry.Misc["AdministrativeDistances"] = o.AdministrativeDistances.Misc
			}
			if o.AdministrativeDistances.Ospfv3Ext != nil {
				nestedAdministrativeDistances.Ospfv3Ext = o.AdministrativeDistances.Ospfv3Ext
			}
			if o.AdministrativeDistances.Ibgp != nil {
				nestedAdministrativeDistances.Ibgp = o.AdministrativeDistances.Ibgp
			}
			if o.AdministrativeDistances.Ebgp != nil {
				nestedAdministrativeDistances.Ebgp = o.AdministrativeDistances.Ebgp
			}
			if o.AdministrativeDistances.Rip != nil {
				nestedAdministrativeDistances.Rip = o.AdministrativeDistances.Rip
			}
			if o.AdministrativeDistances.StaticIpv6 != nil {
				nestedAdministrativeDistances.StaticIpv6 = o.AdministrativeDistances.StaticIpv6
			}
			if o.AdministrativeDistances.OspfExt != nil {
				nestedAdministrativeDistances.OspfExt = o.AdministrativeDistances.OspfExt
			}
			if o.AdministrativeDistances.Ospfv3Int != nil {
				nestedAdministrativeDistances.Ospfv3Int = o.AdministrativeDistances.Ospfv3Int
			}
			if o.AdministrativeDistances.Static != nil {
				nestedAdministrativeDistances.Static = o.AdministrativeDistances.Static
			}
			if o.AdministrativeDistances.OspfInt != nil {
				nestedAdministrativeDistances.OspfInt = o.AdministrativeDistances.OspfInt
			}
		}
		entry.AdministrativeDistances = nestedAdministrativeDistances

		var nestedEcmp *Ecmp
		if o.Ecmp != nil {
			nestedEcmp = &Ecmp{}
			if o.Ecmp.Misc != nil {
				entry.Misc["Ecmp"] = o.Ecmp.Misc
			}
			if o.Ecmp.MaxPaths != nil {
				nestedEcmp.MaxPaths = o.Ecmp.MaxPaths
			}
			if o.Ecmp.Algorithm != nil {
				nestedEcmp.Algorithm = &EcmpAlgorithm{}
				if o.Ecmp.Algorithm.Misc != nil {
					entry.Misc["EcmpAlgorithm"] = o.Ecmp.Algorithm.Misc
				}
				if o.Ecmp.Algorithm.IpModulo != nil {
					nestedEcmp.Algorithm.IpModulo = &EcmpAlgorithmIpModulo{}
					if o.Ecmp.Algorithm.IpModulo.Misc != nil {
						entry.Misc["EcmpAlgorithmIpModulo"] = o.Ecmp.Algorithm.IpModulo.Misc
					}
				}
				if o.Ecmp.Algorithm.IpHash != nil {
					nestedEcmp.Algorithm.IpHash = &EcmpAlgorithmIpHash{}
					if o.Ecmp.Algorithm.IpHash.Misc != nil {
						entry.Misc["EcmpAlgorithmIpHash"] = o.Ecmp.Algorithm.IpHash.Misc
					}
					if o.Ecmp.Algorithm.IpHash.SrcOnly != nil {
						nestedEcmp.Algorithm.IpHash.SrcOnly = util.AsBool(o.Ecmp.Algorithm.IpHash.SrcOnly, nil)
					}
					if o.Ecmp.Algorithm.IpHash.UsePort != nil {
						nestedEcmp.Algorithm.IpHash.UsePort = util.AsBool(o.Ecmp.Algorithm.IpHash.UsePort, nil)
					}
					if o.Ecmp.Algorithm.IpHash.HashSeed != nil {
						nestedEcmp.Algorithm.IpHash.HashSeed = o.Ecmp.Algorithm.IpHash.HashSeed
					}
				}
				if o.Ecmp.Algorithm.WeightedRoundRobin != nil {
					nestedEcmp.Algorithm.WeightedRoundRobin = &EcmpAlgorithmWeightedRoundRobin{}
					if o.Ecmp.Algorithm.WeightedRoundRobin.Misc != nil {
						entry.Misc["EcmpAlgorithmWeightedRoundRobin"] = o.Ecmp.Algorithm.WeightedRoundRobin.Misc
					}
					if o.Ecmp.Algorithm.WeightedRoundRobin.Interfaces != nil {
						nestedEcmp.Algorithm.WeightedRoundRobin.Interfaces = []EcmpAlgorithmWeightedRoundRobinInterfaces{}
						for _, oEcmpAlgorithmWeightedRoundRobinInterfaces := range o.Ecmp.Algorithm.WeightedRoundRobin.Interfaces {
							nestedEcmpAlgorithmWeightedRoundRobinInterfaces := EcmpAlgorithmWeightedRoundRobinInterfaces{}
							if oEcmpAlgorithmWeightedRoundRobinInterfaces.Misc != nil {
								entry.Misc["EcmpAlgorithmWeightedRoundRobinInterfaces"] = oEcmpAlgorithmWeightedRoundRobinInterfaces.Misc
							}
							if oEcmpAlgorithmWeightedRoundRobinInterfaces.Name != "" {
								nestedEcmpAlgorithmWeightedRoundRobinInterfaces.Name = oEcmpAlgorithmWeightedRoundRobinInterfaces.Name
							}
							if oEcmpAlgorithmWeightedRoundRobinInterfaces.Weight != nil {
								nestedEcmpAlgorithmWeightedRoundRobinInterfaces.Weight = oEcmpAlgorithmWeightedRoundRobinInterfaces.Weight
							}
							nestedEcmp.Algorithm.WeightedRoundRobin.Interfaces = append(nestedEcmp.Algorithm.WeightedRoundRobin.Interfaces, nestedEcmpAlgorithmWeightedRoundRobinInterfaces)
						}
					}
				}
				if o.Ecmp.Algorithm.BalancedRoundRobin != nil {
					nestedEcmp.Algorithm.BalancedRoundRobin = &EcmpAlgorithmBalancedRoundRobin{}
					if o.Ecmp.Algorithm.BalancedRoundRobin.Misc != nil {
						entry.Misc["EcmpAlgorithmBalancedRoundRobin"] = o.Ecmp.Algorithm.BalancedRoundRobin.Misc
					}
				}
			}
			if o.Ecmp.Enable != nil {
				nestedEcmp.Enable = util.AsBool(o.Ecmp.Enable, nil)
			}
			if o.Ecmp.SymmetricReturn != nil {
				nestedEcmp.SymmetricReturn = util.AsBool(o.Ecmp.SymmetricReturn, nil)
			}
			if o.Ecmp.StrictSourcePath != nil {
				nestedEcmp.StrictSourcePath = util.AsBool(o.Ecmp.StrictSourcePath, nil)
			}
		}
		entry.Ecmp = nestedEcmp

		entry.Interfaces = util.MemToStr(o.Interfaces)
		var nestedProtocol *Protocol
		if o.Protocol != nil {
			nestedProtocol = &Protocol{}
			if o.Protocol.Misc != nil {
				entry.Misc["Protocol"] = o.Protocol.Misc
			}
			if o.Protocol.Rip != nil {
				nestedProtocol.Rip = &ProtocolRip{}
				if o.Protocol.Rip.Misc != nil {
					entry.Misc["ProtocolRip"] = o.Protocol.Rip.Misc
				}
				if o.Protocol.Rip.Enable != nil {
					nestedProtocol.Rip.Enable = util.AsBool(o.Protocol.Rip.Enable, nil)
				}
			}
			if o.Protocol.Ospf != nil {
				nestedProtocol.Ospf = &ProtocolOspf{}
				if o.Protocol.Ospf.Misc != nil {
					entry.Misc["ProtocolOspf"] = o.Protocol.Ospf.Misc
				}
				if o.Protocol.Ospf.Enable != nil {
					nestedProtocol.Ospf.Enable = util.AsBool(o.Protocol.Ospf.Enable, nil)
				}
			}
			if o.Protocol.Ospfv3 != nil {
				nestedProtocol.Ospfv3 = &ProtocolOspfv3{}
				if o.Protocol.Ospfv3.Misc != nil {
					entry.Misc["ProtocolOspfv3"] = o.Protocol.Ospfv3.Misc
				}
				if o.Protocol.Ospfv3.Enable != nil {
					nestedProtocol.Ospfv3.Enable = util.AsBool(o.Protocol.Ospfv3.Enable, nil)
				}
			}
			if o.Protocol.Bgp != nil {
				nestedProtocol.Bgp = &ProtocolBgp{}
				if o.Protocol.Bgp.Misc != nil {
					entry.Misc["ProtocolBgp"] = o.Protocol.Bgp.Misc
				}
				if o.Protocol.Bgp.Enable != nil {
					nestedProtocol.Bgp.Enable = util.AsBool(o.Protocol.Bgp.Enable, nil)
				}
			}
		}
		entry.Protocol = nestedProtocol

		var nestedRoutingTable *RoutingTable
		if o.RoutingTable != nil {
			nestedRoutingTable = &RoutingTable{}
			if o.RoutingTable.Misc != nil {
				entry.Misc["RoutingTable"] = o.RoutingTable.Misc
			}
			if o.RoutingTable.Ip != nil {
				nestedRoutingTable.Ip = &RoutingTableIp{}
				if o.RoutingTable.Ip.Misc != nil {
					entry.Misc["RoutingTableIp"] = o.RoutingTable.Ip.Misc
				}
				if o.RoutingTable.Ip.StaticRoutes != nil {
					nestedRoutingTable.Ip.StaticRoutes = []RoutingTableIpStaticRoutes{}
					for _, oRoutingTableIpStaticRoutes := range o.RoutingTable.Ip.StaticRoutes {
						nestedRoutingTableIpStaticRoutes := RoutingTableIpStaticRoutes{}
						if oRoutingTableIpStaticRoutes.Misc != nil {
							entry.Misc["RoutingTableIpStaticRoutes"] = oRoutingTableIpStaticRoutes.Misc
						}
						if oRoutingTableIpStaticRoutes.Name != "" {
							nestedRoutingTableIpStaticRoutes.Name = oRoutingTableIpStaticRoutes.Name
						}
						if oRoutingTableIpStaticRoutes.Destination != nil {
							nestedRoutingTableIpStaticRoutes.Destination = oRoutingTableIpStaticRoutes.Destination
						}
						if oRoutingTableIpStaticRoutes.Interface != nil {
							nestedRoutingTableIpStaticRoutes.Interface = oRoutingTableIpStaticRoutes.Interface
						}
						if oRoutingTableIpStaticRoutes.NextHop != nil {
							nestedRoutingTableIpStaticRoutes.NextHop = &RoutingTableIpStaticRoutesNextHop{}
							if oRoutingTableIpStaticRoutes.NextHop.Misc != nil {
								entry.Misc["RoutingTableIpStaticRoutesNextHop"] = oRoutingTableIpStaticRoutes.NextHop.Misc
							}
							if oRoutingTableIpStaticRoutes.NextHop.IpAddress != nil {
								nestedRoutingTableIpStaticRoutes.NextHop.IpAddress = oRoutingTableIpStaticRoutes.NextHop.IpAddress
							}
							if oRoutingTableIpStaticRoutes.NextHop.Fqdn != nil {
								nestedRoutingTableIpStaticRoutes.NextHop.Fqdn = oRoutingTableIpStaticRoutes.NextHop.Fqdn
							}
							if oRoutingTableIpStaticRoutes.NextHop.NextVr != nil {
								nestedRoutingTableIpStaticRoutes.NextHop.NextVr = oRoutingTableIpStaticRoutes.NextHop.NextVr
							}
							if oRoutingTableIpStaticRoutes.NextHop.Tunnel != nil {
								nestedRoutingTableIpStaticRoutes.NextHop.Tunnel = oRoutingTableIpStaticRoutes.NextHop.Tunnel
							}
						}
						if oRoutingTableIpStaticRoutes.AdminDist != nil {
							nestedRoutingTableIpStaticRoutes.AdminDist = oRoutingTableIpStaticRoutes.AdminDist
						}
						if oRoutingTableIpStaticRoutes.Metric != nil {
							nestedRoutingTableIpStaticRoutes.Metric = oRoutingTableIpStaticRoutes.Metric
						}
						if oRoutingTableIpStaticRoutes.RouteTable != nil {
							nestedRoutingTableIpStaticRoutes.RouteTable = oRoutingTableIpStaticRoutes.RouteTable
						}
						nestedRoutingTable.Ip.StaticRoutes = append(nestedRoutingTable.Ip.StaticRoutes, nestedRoutingTableIpStaticRoutes)
					}
				}
			}
			if o.RoutingTable.Ipv6 != nil {
				nestedRoutingTable.Ipv6 = &RoutingTableIpv6{}
				if o.RoutingTable.Ipv6.Misc != nil {
					entry.Misc["RoutingTableIpv6"] = o.RoutingTable.Ipv6.Misc
				}
				if o.RoutingTable.Ipv6.StaticRoutes != nil {
					nestedRoutingTable.Ipv6.StaticRoutes = []RoutingTableIpv6StaticRoutes{}
					for _, oRoutingTableIpv6StaticRoutes := range o.RoutingTable.Ipv6.StaticRoutes {
						nestedRoutingTableIpv6StaticRoutes := RoutingTableIpv6StaticRoutes{}
						if oRoutingTableIpv6StaticRoutes.Misc != nil {
							entry.Misc["RoutingTableIpv6StaticRoutes"] = oRoutingTableIpv6StaticRoutes.Misc
						}
						if oRoutingTableIpv6StaticRoutes.Destination != nil {
							nestedRoutingTableIpv6StaticRoutes.Destination = oRoutingTableIpv6StaticRoutes.Destination
						}
						if oRoutingTableIpv6StaticRoutes.Interface != nil {
							nestedRoutingTableIpv6StaticRoutes.Interface = oRoutingTableIpv6StaticRoutes.Interface
						}
						if oRoutingTableIpv6StaticRoutes.NextHop != nil {
							nestedRoutingTableIpv6StaticRoutes.NextHop = &RoutingTableIpv6StaticRoutesNextHop{}
							if oRoutingTableIpv6StaticRoutes.NextHop.Misc != nil {
								entry.Misc["RoutingTableIpv6StaticRoutesNextHop"] = oRoutingTableIpv6StaticRoutes.NextHop.Misc
							}
							if oRoutingTableIpv6StaticRoutes.NextHop.Ipv6Address != nil {
								nestedRoutingTableIpv6StaticRoutes.NextHop.Ipv6Address = oRoutingTableIpv6StaticRoutes.NextHop.Ipv6Address
							}
							if oRoutingTableIpv6StaticRoutes.NextHop.Fqdn != nil {
								nestedRoutingTableIpv6StaticRoutes.NextHop.Fqdn = oRoutingTableIpv6StaticRoutes.NextHop.Fqdn
							}
							if oRoutingTableIpv6StaticRoutes.NextHop.NextVr != nil {
								nestedRoutingTableIpv6StaticRoutes.NextHop.NextVr = oRoutingTableIpv6StaticRoutes.NextHop.NextVr
							}
							if oRoutingTableIpv6StaticRoutes.NextHop.Tunnel != nil {
								nestedRoutingTableIpv6StaticRoutes.NextHop.Tunnel = oRoutingTableIpv6StaticRoutes.NextHop.Tunnel
							}
						}
						if oRoutingTableIpv6StaticRoutes.AdminDist != nil {
							nestedRoutingTableIpv6StaticRoutes.AdminDist = oRoutingTableIpv6StaticRoutes.AdminDist
						}
						if oRoutingTableIpv6StaticRoutes.Metric != nil {
							nestedRoutingTableIpv6StaticRoutes.Metric = oRoutingTableIpv6StaticRoutes.Metric
						}
						if oRoutingTableIpv6StaticRoutes.RouteTable != nil {
							nestedRoutingTableIpv6StaticRoutes.RouteTable = oRoutingTableIpv6StaticRoutes.RouteTable
						}
						if oRoutingTableIpv6StaticRoutes.Name != "" {
							nestedRoutingTableIpv6StaticRoutes.Name = oRoutingTableIpv6StaticRoutes.Name
						}
						nestedRoutingTable.Ipv6.StaticRoutes = append(nestedRoutingTable.Ipv6.StaticRoutes, nestedRoutingTableIpv6StaticRoutes)
					}
				}
			}
		}
		entry.RoutingTable = nestedRoutingTable

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
	if !matchAdministrativeDistances(a.AdministrativeDistances, b.AdministrativeDistances) {
		return false
	}
	if !matchEcmp(a.Ecmp, b.Ecmp) {
		return false
	}
	if !util.OrderedListsMatch(a.Interfaces, b.Interfaces) {
		return false
	}
	if !matchProtocol(a.Protocol, b.Protocol) {
		return false
	}
	if !matchRoutingTable(a.RoutingTable, b.RoutingTable) {
		return false
	}

	return true
}

func matchRoutingTableIpStaticRoutesNextHop(a *RoutingTableIpStaticRoutesNextHop, b *RoutingTableIpStaticRoutesNextHop) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.IpAddress, b.IpAddress) {
		return false
	}
	if !util.StringsMatch(a.Fqdn, b.Fqdn) {
		return false
	}
	if !util.StringsMatch(a.NextVr, b.NextVr) {
		return false
	}
	if !util.StringsMatch(a.Tunnel, b.Tunnel) {
		return false
	}
	return true
}
func matchRoutingTableIpStaticRoutes(a []RoutingTableIpStaticRoutes, b []RoutingTableIpStaticRoutes) bool {
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
			if !util.StringsMatch(a.Destination, b.Destination) {
				return false
			}
			if !util.StringsMatch(a.Interface, b.Interface) {
				return false
			}
			if !matchRoutingTableIpStaticRoutesNextHop(a.NextHop, b.NextHop) {
				return false
			}
			if !util.Ints64Match(a.AdminDist, b.AdminDist) {
				return false
			}
			if !util.Ints64Match(a.Metric, b.Metric) {
				return false
			}
			if !util.StringsMatch(a.RouteTable, b.RouteTable) {
				return false
			}
		}
	}
	return true
}
func matchRoutingTableIp(a *RoutingTableIp, b *RoutingTableIp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoutingTableIpStaticRoutes(a.StaticRoutes, b.StaticRoutes) {
		return false
	}
	return true
}
func matchRoutingTableIpv6StaticRoutesNextHop(a *RoutingTableIpv6StaticRoutesNextHop, b *RoutingTableIpv6StaticRoutesNextHop) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.StringsMatch(a.Ipv6Address, b.Ipv6Address) {
		return false
	}
	if !util.StringsMatch(a.Fqdn, b.Fqdn) {
		return false
	}
	if !util.StringsMatch(a.NextVr, b.NextVr) {
		return false
	}
	if !util.StringsMatch(a.Tunnel, b.Tunnel) {
		return false
	}
	return true
}
func matchRoutingTableIpv6StaticRoutes(a []RoutingTableIpv6StaticRoutes, b []RoutingTableIpv6StaticRoutes) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.StringsMatch(a.RouteTable, b.RouteTable) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
			if !util.StringsMatch(a.Destination, b.Destination) {
				return false
			}
			if !util.StringsMatch(a.Interface, b.Interface) {
				return false
			}
			if !matchRoutingTableIpv6StaticRoutesNextHop(a.NextHop, b.NextHop) {
				return false
			}
			if !util.Ints64Match(a.AdminDist, b.AdminDist) {
				return false
			}
			if !util.Ints64Match(a.Metric, b.Metric) {
				return false
			}
		}
	}
	return true
}
func matchRoutingTableIpv6(a *RoutingTableIpv6, b *RoutingTableIpv6) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoutingTableIpv6StaticRoutes(a.StaticRoutes, b.StaticRoutes) {
		return false
	}
	return true
}
func matchRoutingTable(a *RoutingTable, b *RoutingTable) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchRoutingTableIp(a.Ip, b.Ip) {
		return false
	}
	if !matchRoutingTableIpv6(a.Ipv6, b.Ipv6) {
		return false
	}
	return true
}
func matchProtocolBgp(a *ProtocolBgp, b *ProtocolBgp) bool {
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
func matchProtocolRip(a *ProtocolRip, b *ProtocolRip) bool {
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
func matchProtocolOspf(a *ProtocolOspf, b *ProtocolOspf) bool {
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
func matchProtocolOspfv3(a *ProtocolOspfv3, b *ProtocolOspfv3) bool {
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
func matchProtocol(a *Protocol, b *Protocol) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchProtocolBgp(a.Bgp, b.Bgp) {
		return false
	}
	if !matchProtocolRip(a.Rip, b.Rip) {
		return false
	}
	if !matchProtocolOspf(a.Ospf, b.Ospf) {
		return false
	}
	if !matchProtocolOspfv3(a.Ospfv3, b.Ospfv3) {
		return false
	}
	return true
}
func matchEcmpAlgorithmWeightedRoundRobinInterfaces(a []EcmpAlgorithmWeightedRoundRobinInterfaces, b []EcmpAlgorithmWeightedRoundRobinInterfaces) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	for _, a := range a {
		for _, b := range b {
			if !util.Ints64Match(a.Weight, b.Weight) {
				return false
			}
			if !util.StringsEqual(a.Name, b.Name) {
				return false
			}
		}
	}
	return true
}
func matchEcmpAlgorithmWeightedRoundRobin(a *EcmpAlgorithmWeightedRoundRobin, b *EcmpAlgorithmWeightedRoundRobin) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchEcmpAlgorithmWeightedRoundRobinInterfaces(a.Interfaces, b.Interfaces) {
		return false
	}
	return true
}
func matchEcmpAlgorithmBalancedRoundRobin(a *EcmpAlgorithmBalancedRoundRobin, b *EcmpAlgorithmBalancedRoundRobin) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchEcmpAlgorithmIpModulo(a *EcmpAlgorithmIpModulo, b *EcmpAlgorithmIpModulo) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	return true
}
func matchEcmpAlgorithmIpHash(a *EcmpAlgorithmIpHash, b *EcmpAlgorithmIpHash) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.SrcOnly, b.SrcOnly) {
		return false
	}
	if !util.BoolsMatch(a.UsePort, b.UsePort) {
		return false
	}
	if !util.Ints64Match(a.HashSeed, b.HashSeed) {
		return false
	}
	return true
}
func matchEcmpAlgorithm(a *EcmpAlgorithm, b *EcmpAlgorithm) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !matchEcmpAlgorithmIpModulo(a.IpModulo, b.IpModulo) {
		return false
	}
	if !matchEcmpAlgorithmIpHash(a.IpHash, b.IpHash) {
		return false
	}
	if !matchEcmpAlgorithmWeightedRoundRobin(a.WeightedRoundRobin, b.WeightedRoundRobin) {
		return false
	}
	if !matchEcmpAlgorithmBalancedRoundRobin(a.BalancedRoundRobin, b.BalancedRoundRobin) {
		return false
	}
	return true
}
func matchEcmp(a *Ecmp, b *Ecmp) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.BoolsMatch(a.StrictSourcePath, b.StrictSourcePath) {
		return false
	}
	if !util.Ints64Match(a.MaxPaths, b.MaxPaths) {
		return false
	}
	if !matchEcmpAlgorithm(a.Algorithm, b.Algorithm) {
		return false
	}
	if !util.BoolsMatch(a.Enable, b.Enable) {
		return false
	}
	if !util.BoolsMatch(a.SymmetricReturn, b.SymmetricReturn) {
		return false
	}
	return true
}
func matchAdministrativeDistances(a *AdministrativeDistances, b *AdministrativeDistances) bool {
	if a == nil && b != nil || a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	}
	if !util.Ints64Match(a.OspfExt, b.OspfExt) {
		return false
	}
	if !util.Ints64Match(a.Ospfv3Int, b.Ospfv3Int) {
		return false
	}
	if !util.Ints64Match(a.Ospfv3Ext, b.Ospfv3Ext) {
		return false
	}
	if !util.Ints64Match(a.Ibgp, b.Ibgp) {
		return false
	}
	if !util.Ints64Match(a.Ebgp, b.Ebgp) {
		return false
	}
	if !util.Ints64Match(a.Rip, b.Rip) {
		return false
	}
	if !util.Ints64Match(a.StaticIpv6, b.StaticIpv6) {
		return false
	}
	if !util.Ints64Match(a.OspfInt, b.OspfInt) {
		return false
	}
	if !util.Ints64Match(a.Static, b.Static) {
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
