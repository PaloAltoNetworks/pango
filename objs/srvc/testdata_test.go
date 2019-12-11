package srvc

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type tc struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []tc {
	return []tc{
		{"v1 tcp service no source port no tag", version.Number{7, 1, 0, ""}, Entry{
			Name:            "tcp1",
			Description:     "my service",
			Protocol:        "tcp",
			DestinationPort: "1234",
		}},
		{"v1 tcp service no source port with tag", version.Number{7, 1, 0, ""}, Entry{
			Name:            "tcp2",
			Description:     "my service",
			Protocol:        "tcp",
			DestinationPort: "1234",
			Tags:            []string{"tag1", "tag2"},
		}},
		{"v1 tcp service with source port no tag", version.Number{7, 1, 0, ""}, Entry{
			Name:            "tcp3",
			Description:     "my service",
			Protocol:        "tcp",
			SourcePort:      "1025",
			DestinationPort: "1234",
		}},
		{"v1 tcp service with source port with tag", version.Number{7, 1, 0, ""}, Entry{
			Name:            "tcp4",
			Description:     "my service",
			Protocol:        "tcp",
			SourcePort:      "1025",
			DestinationPort: "1234",
			Tags:            []string{"tag1", "tag2"},
		}},
		{"v1 udp service no source port no tag", version.Number{7, 1, 0, ""}, Entry{
			Name:            "udp1",
			Description:     "my service",
			Protocol:        "udp",
			DestinationPort: "1234",
		}},
		{"v1 udp service no source port with tag", version.Number{7, 1, 0, ""}, Entry{
			Name:            "udp2",
			Description:     "my service",
			Protocol:        "udp",
			DestinationPort: "1234",
			Tags:            []string{"tag1", "tag2"},
		}},
		{"v1 udp service with source port no tag", version.Number{7, 1, 0, ""}, Entry{
			Name:            "udp3",
			Description:     "my service",
			Protocol:        "udp",
			SourcePort:      "1025",
			DestinationPort: "1234",
		}},
		{"v1 udp service with source port with tag", version.Number{7, 1, 0, ""}, Entry{
			Name:            "udp4",
			Description:     "my service",
			Protocol:        "udp",
			SourcePort:      "1025",
			DestinationPort: "1234",
			Tags:            []string{"tag1", "tag2"},
		}},
		{"v2 tcp service no source port no tag", version.Number{8, 1, 0, ""}, Entry{
			Name:            "tcp1",
			Description:     "my service",
			Protocol:        ProtocolTcp,
			DestinationPort: "1234",
		}},
		{"v2 tcp service no source port with tag", version.Number{8, 1, 0, ""}, Entry{
			Name:            "tcp2",
			Description:     "my service",
			Protocol:        ProtocolTcp,
			DestinationPort: "1234",
			Tags:            []string{"tag1", "tag2"},
		}},
		{"v2 tcp service with source port no tag", version.Number{8, 1, 0, ""}, Entry{
			Name:            "tcp3",
			Description:     "my service",
			Protocol:        ProtocolTcp,
			SourcePort:      "1025",
			DestinationPort: "1234",
		}},
		{"v2 tcp service with source port with tag", version.Number{8, 1, 0, ""}, Entry{
			Name:            "tcp4",
			Description:     "my service",
			Protocol:        ProtocolTcp,
			SourcePort:      "1025",
			DestinationPort: "1234",
			Tags:            []string{"tag1", "tag2"},
		}},
		{"v2 tcp service with override", version.Number{8, 1, 0, ""}, Entry{
			Name:                      "tcp5",
			Description:               "my service",
			Protocol:                  ProtocolTcp,
			SourcePort:                "1025",
			DestinationPort:           "1234",
			Tags:                      []string{"tag1", "tag2"},
			OverrideSessionTimeout:    true,
			OverrideTimeout:           50,
			OverrideHalfClosedTimeout: 60,
			OverrideTimeWaitTimeout:   70,
		}},
		{"v2 udp service no source port no tag", version.Number{8, 1, 0, ""}, Entry{
			Name:            "udp1",
			Description:     "my service",
			Protocol:        ProtocolUdp,
			DestinationPort: "1234",
		}},
		{"v2 udp service no source port with tag", version.Number{8, 1, 0, ""}, Entry{
			Name:            "udp2",
			Description:     "my service",
			Protocol:        ProtocolUdp,
			DestinationPort: "1234",
			Tags:            []string{"tag1", "tag2"},
		}},
		{"v2 udp service with source port no tag", version.Number{8, 1, 0, ""}, Entry{
			Name:            "udp3",
			Description:     "my service",
			Protocol:        ProtocolUdp,
			SourcePort:      "1025",
			DestinationPort: "1234",
		}},
		{"v2 udp service with source port with tag", version.Number{8, 1, 0, ""}, Entry{
			Name:            "udp4",
			Description:     "my service",
			Protocol:        ProtocolUdp,
			SourcePort:      "1025",
			DestinationPort: "1234",
			Tags:            []string{"tag1", "tag2"},
		}},
		{"v2 udp service with override", version.Number{8, 1, 0, ""}, Entry{
			Name:                   "udp4",
			Description:            "my service",
			Protocol:               ProtocolUdp,
			SourcePort:             "1025",
			DestinationPort:        "1234",
			Tags:                   []string{"tag1", "tag2"},
			OverrideSessionTimeout: true,
			OverrideTimeout:        42,
		}},
		{"v2 sctp service", version.Number{8, 1, 0, ""}, Entry{
			Name:            "sctp1",
			Description:     "my service",
			Protocol:        ProtocolSctp,
			SourcePort:      "1234",
			DestinationPort: "5678",
			Tags:            []string{"my", "tags"},
		}},
	}
}
