package http

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
		{"v1 tag reg", version.Number{8, 0, 0, ""}, Entry{
			Name:            "foo",
			TagRegistration: true,
			raw: map[string]string{
				"srv": "http servers",
			},
		}},
		{"v1 config", version.Number{8, 0, 0, ""}, Entry{
			Name:            "t1",
			ConfigName:      "the name",
			ConfigUriFormat: "uri format",
			ConfigPayload:   "payload",
			raw: map[string]string{
				"configh": "headers",
				"configp": "params",
			},
		}},
		{"v1 system", version.Number{8, 0, 0, ""}, Entry{
			Name:            "t1",
			SystemName:      "the name",
			SystemUriFormat: "uri format",
			SystemPayload:   "payload",
			raw: map[string]string{
				"systemh": "headers",
				"systemp": "params",
			},
		}},
		{"v1 threat", version.Number{8, 0, 0, ""}, Entry{
			Name:            "t1",
			ThreatName:      "the name",
			ThreatUriFormat: "uri format",
			ThreatPayload:   "payload",
			raw: map[string]string{
				"threath": "headers",
				"threatp": "params",
			},
		}},
		{"v1 traffic", version.Number{8, 0, 0, ""}, Entry{
			Name:             "t1",
			TrafficName:      "the name",
			TrafficUriFormat: "uri format",
			TrafficPayload:   "payload",
			raw: map[string]string{
				"traffich": "headers",
				"trafficp": "params",
			},
		}},
		{"v1 hipmatch", version.Number{8, 0, 0, ""}, Entry{
			Name:              "t1",
			HipMatchName:      "the name",
			HipMatchUriFormat: "uri format",
			HipMatchPayload:   "payload",
			raw: map[string]string{
				"hipmatchh": "headers",
				"hipmatchp": "params",
			},
		}},
		{"v1 url", version.Number{8, 0, 0, ""}, Entry{
			Name:         "t1",
			UrlName:      "the name",
			UrlUriFormat: "uri format",
			UrlPayload:   "payload",
			raw: map[string]string{
				"urlh": "headers",
				"urlp": "params",
			},
		}},
		{"v1 data", version.Number{8, 0, 0, ""}, Entry{
			Name:          "t1",
			DataName:      "the name",
			DataUriFormat: "uri format",
			DataPayload:   "payload",
			raw: map[string]string{
				"datah": "headers",
				"datap": "params",
			},
		}},
		{"v1 wildfire", version.Number{8, 0, 0, ""}, Entry{
			Name:              "t1",
			WildfireName:      "the name",
			WildfireUriFormat: "uri format",
			WildfirePayload:   "payload",
			raw: map[string]string{
				"wildfireh": "headers",
				"wildfirep": "params",
			},
		}},
		{"v1 tunnel", version.Number{8, 0, 0, ""}, Entry{
			Name:            "t1",
			TunnelName:      "the name",
			TunnelUriFormat: "uri format",
			TunnelPayload:   "payload",
			raw: map[string]string{
				"tunnelh": "headers",
				"tunnelp": "params",
			},
		}},
		{"v1 userid", version.Number{8, 0, 0, ""}, Entry{
			Name:            "t1",
			UserIdName:      "the name",
			UserIdUriFormat: "uri format",
			UserIdPayload:   "payload",
			raw: map[string]string{
				"useridh": "headers",
				"useridp": "params",
			},
		}},
		{"v1 gtp", version.Number{8, 0, 0, ""}, Entry{
			Name:         "t1",
			GtpName:      "the name",
			GtpUriFormat: "uri format",
			GtpPayload:   "payload",
			raw: map[string]string{
				"gtph": "headers",
				"gtpp": "params",
			},
		}},
		{"v1 auth", version.Number{8, 0, 0, ""}, Entry{
			Name:          "t1",
			AuthName:      "the name",
			AuthUriFormat: "uri format",
			AuthPayload:   "payload",
			raw: map[string]string{
				"authh": "headers",
				"authp": "params",
			},
		}},
		{"v2 tag reg", version.Number{8, 1, 0, ""}, Entry{
			Name:            "foo",
			TagRegistration: true,
			raw: map[string]string{
				"srv": "http servers",
			},
		}},
		{"v2 config", version.Number{8, 1, 0, ""}, Entry{
			Name:            "t1",
			ConfigName:      "the name",
			ConfigUriFormat: "uri format",
			ConfigPayload:   "payload",
			raw: map[string]string{
				"configh": "headers",
				"configp": "params",
			},
		}},
		{"v2 system", version.Number{8, 1, 0, ""}, Entry{
			Name:            "t1",
			SystemName:      "the name",
			SystemUriFormat: "uri format",
			SystemPayload:   "payload",
			raw: map[string]string{
				"systemh": "headers",
				"systemp": "params",
			},
		}},
		{"v2 threat", version.Number{8, 1, 0, ""}, Entry{
			Name:            "t1",
			ThreatName:      "the name",
			ThreatUriFormat: "uri format",
			ThreatPayload:   "payload",
			raw: map[string]string{
				"threath": "headers",
				"threatp": "params",
			},
		}},
		{"v2 traffic", version.Number{8, 1, 0, ""}, Entry{
			Name:             "t1",
			TrafficName:      "the name",
			TrafficUriFormat: "uri format",
			TrafficPayload:   "payload",
			raw: map[string]string{
				"traffich": "headers",
				"trafficp": "params",
			},
		}},
		{"v2 hipmatch", version.Number{8, 1, 0, ""}, Entry{
			Name:              "t1",
			HipMatchName:      "the name",
			HipMatchUriFormat: "uri format",
			HipMatchPayload:   "payload",
			raw: map[string]string{
				"hipmatchh": "headers",
				"hipmatchp": "params",
			},
		}},
		{"v2 url", version.Number{8, 1, 0, ""}, Entry{
			Name:         "t1",
			UrlName:      "the name",
			UrlUriFormat: "uri format",
			UrlPayload:   "payload",
			raw: map[string]string{
				"urlh": "headers",
				"urlp": "params",
			},
		}},
		{"v2 data", version.Number{8, 1, 0, ""}, Entry{
			Name:          "t1",
			DataName:      "the name",
			DataUriFormat: "uri format",
			DataPayload:   "payload",
			raw: map[string]string{
				"datah": "headers",
				"datap": "params",
			},
		}},
		{"v2 wildfire", version.Number{8, 1, 0, ""}, Entry{
			Name:              "t1",
			WildfireName:      "the name",
			WildfireUriFormat: "uri format",
			WildfirePayload:   "payload",
			raw: map[string]string{
				"wildfireh": "headers",
				"wildfirep": "params",
			},
		}},
		{"v2 tunnel", version.Number{8, 1, 0, ""}, Entry{
			Name:            "t1",
			TunnelName:      "the name",
			TunnelUriFormat: "uri format",
			TunnelPayload:   "payload",
			raw: map[string]string{
				"tunnelh": "headers",
				"tunnelp": "params",
			},
		}},
		{"v2 userid", version.Number{8, 1, 0, ""}, Entry{
			Name:            "t1",
			UserIdName:      "the name",
			UserIdUriFormat: "uri format",
			UserIdPayload:   "payload",
			raw: map[string]string{
				"useridh": "headers",
				"useridp": "params",
			},
		}},
		{"v2 gtp", version.Number{8, 1, 0, ""}, Entry{
			Name:         "t1",
			GtpName:      "the name",
			GtpUriFormat: "uri format",
			GtpPayload:   "payload",
			raw: map[string]string{
				"gtph": "headers",
				"gtpp": "params",
			},
		}},
		{"v2 auth", version.Number{8, 1, 0, ""}, Entry{
			Name:          "t1",
			AuthName:      "the name",
			AuthUriFormat: "uri format",
			AuthPayload:   "payload",
			raw: map[string]string{
				"authh": "headers",
				"authp": "params",
			},
		}},
		{"v2 sctp", version.Number{8, 1, 0, ""}, Entry{
			Name:          "t1",
			SctpName:      "the name",
			SctpUriFormat: "uri format",
			SctpPayload:   "payload",
			raw: map[string]string{
				"sctph": "headers",
				"sctpp": "params",
			},
		}},
		{"v3 tag reg", version.Number{9, 0, 0, ""}, Entry{
			Name:            "foo",
			TagRegistration: true,
			raw: map[string]string{
				"srv": "http servers",
			},
		}},
		{"v3 config", version.Number{9, 0, 0, ""}, Entry{
			Name:            "t1",
			ConfigName:      "the name",
			ConfigUriFormat: "uri format",
			ConfigPayload:   "payload",
			raw: map[string]string{
				"configh": "headers",
				"configp": "params",
			},
		}},
		{"v3 system", version.Number{9, 0, 0, ""}, Entry{
			Name:            "t1",
			SystemName:      "the name",
			SystemUriFormat: "uri format",
			SystemPayload:   "payload",
			raw: map[string]string{
				"systemh": "headers",
				"systemp": "params",
			},
		}},
		{"v3 threat", version.Number{9, 0, 0, ""}, Entry{
			Name:            "t1",
			ThreatName:      "the name",
			ThreatUriFormat: "uri format",
			ThreatPayload:   "payload",
			raw: map[string]string{
				"threath": "headers",
				"threatp": "params",
			},
		}},
		{"v3 traffic", version.Number{9, 0, 0, ""}, Entry{
			Name:             "t1",
			TrafficName:      "the name",
			TrafficUriFormat: "uri format",
			TrafficPayload:   "payload",
			raw: map[string]string{
				"traffich": "headers",
				"trafficp": "params",
			},
		}},
		{"v3 hipmatch", version.Number{9, 0, 0, ""}, Entry{
			Name:              "t1",
			HipMatchName:      "the name",
			HipMatchUriFormat: "uri format",
			HipMatchPayload:   "payload",
			raw: map[string]string{
				"hipmatchh": "headers",
				"hipmatchp": "params",
			},
		}},
		{"v3 url", version.Number{9, 0, 0, ""}, Entry{
			Name:         "t1",
			UrlName:      "the name",
			UrlUriFormat: "uri format",
			UrlPayload:   "payload",
			raw: map[string]string{
				"urlh": "headers",
				"urlp": "params",
			},
		}},
		{"v3 data", version.Number{9, 0, 0, ""}, Entry{
			Name:          "t1",
			DataName:      "the name",
			DataUriFormat: "uri format",
			DataPayload:   "payload",
			raw: map[string]string{
				"datah": "headers",
				"datap": "params",
			},
		}},
		{"v3 wildfire", version.Number{9, 0, 0, ""}, Entry{
			Name:              "t1",
			WildfireName:      "the name",
			WildfireUriFormat: "uri format",
			WildfirePayload:   "payload",
			raw: map[string]string{
				"wildfireh": "headers",
				"wildfirep": "params",
			},
		}},
		{"v3 tunnel", version.Number{9, 0, 0, ""}, Entry{
			Name:            "t1",
			TunnelName:      "the name",
			TunnelUriFormat: "uri format",
			TunnelPayload:   "payload",
			raw: map[string]string{
				"tunnelh": "headers",
				"tunnelp": "params",
			},
		}},
		{"v3 userid", version.Number{9, 0, 0, ""}, Entry{
			Name:            "t1",
			UserIdName:      "the name",
			UserIdUriFormat: "uri format",
			UserIdPayload:   "payload",
			raw: map[string]string{
				"useridh": "headers",
				"useridp": "params",
			},
		}},
		{"v3 gtp", version.Number{9, 0, 0, ""}, Entry{
			Name:         "t1",
			GtpName:      "the name",
			GtpUriFormat: "uri format",
			GtpPayload:   "payload",
			raw: map[string]string{
				"gtph": "headers",
				"gtpp": "params",
			},
		}},
		{"v3 auth", version.Number{9, 0, 0, ""}, Entry{
			Name:          "t1",
			AuthName:      "the name",
			AuthUriFormat: "uri format",
			AuthPayload:   "payload",
			raw: map[string]string{
				"authh": "headers",
				"authp": "params",
			},
		}},
		{"v3 sctp", version.Number{9, 0, 0, ""}, Entry{
			Name:          "t1",
			SctpName:      "the name",
			SctpUriFormat: "uri format",
			SctpPayload:   "payload",
			raw: map[string]string{
				"sctph": "headers",
				"sctpp": "params",
			},
		}},
		{"v3 iptag", version.Number{9, 0, 0, ""}, Entry{
			Name:           "t1",
			IptagName:      "the name",
			IptagUriFormat: "uri format",
			IptagPayload:   "payload",
			raw: map[string]string{
				"iptagh": "headers",
				"iptagp": "params",
			},
		}},
	}
}
