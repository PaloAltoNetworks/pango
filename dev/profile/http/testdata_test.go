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
		{"v1 tag reg", version.Number{7, 1, 0, ""}, Entry{
			Name:            "foo",
			TagRegistration: true,
		}},
		{"v1 with servers", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Servers: []Server{
				{
					Name:       "t1",
					Address:    "http.example.com",
					Protocol:   ProtocolHttps,
					Port:       443,
					HttpMethod: "POST",
					Username:   "someuser",
					Password:   "somepasswd",
				},
			},
		}},
		{"v1 config", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Config: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v1 config with headers", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Config: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v1 config with params", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Config: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v1 system", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			System: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v1 system with headers", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			System: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v1 system with params", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			System: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v1 threat", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Threat: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v1 threat with headers", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Threat: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v1 threat with params", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Threat: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v1 traffic", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Traffic: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v1 traffic with headers", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Traffic: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v1 traffic with params", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Traffic: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v1 hip match", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			HipMatch: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v1 hip match with headers", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			HipMatch: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v1 hip match with params", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			HipMatch: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v1 url", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Url: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v1 url with headers", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Url: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v1 url with params", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Url: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v1 data", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Data: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v1 data with headers", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Data: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v1 data with params", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Data: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v1 wildfire", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Wildfire: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v1 wildfire with headers", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Wildfire: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v1 wildfire with params", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Wildfire: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v1 tunnel", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Tunnel: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v1 tunnel with headers", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Tunnel: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v1 tunnel with params", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Tunnel: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v1 user id", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			UserId: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v1 user id with headers", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			UserId: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v1 user id with params", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			UserId: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v1 gtp", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Gtp: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v1 gtp with headers", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Gtp: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v1 gtp with params", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Gtp: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v1 auth", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Auth: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v1 auth with headers", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Auth: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v1 auth with params", version.Number{7, 1, 0, ""}, Entry{
			Name: "foo",
			Auth: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},

		// v2
		{"v2 tag reg", version.Number{8, 1, 0, ""}, Entry{
			Name:            "foo",
			TagRegistration: true,
		}},
		{"v2 with servers", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Servers: []Server{
				{
					Name:       "t1",
					Address:    "http.example.com",
					Protocol:   ProtocolHttps,
					Port:       443,
					HttpMethod: "POST",
					Username:   "someuser",
					Password:   "somepasswd",
				},
			},
		}},
		{"v2 config", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Config: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v2 config with headers", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Config: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v2 config with params", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Config: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v2 system", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			System: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v2 system with headers", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			System: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v2 system with params", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			System: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v2 threat", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Threat: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v2 threat with headers", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Threat: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v2 threat with params", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Threat: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v2 traffic", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Traffic: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v2 traffic with headers", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Traffic: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v2 traffic with params", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Traffic: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v2 hip match", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			HipMatch: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v2 hip match with headers", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			HipMatch: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v2 hip match with params", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			HipMatch: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v2 url", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Url: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v2 url with headers", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Url: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v2 url with params", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Url: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v2 data", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Data: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v2 data with headers", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Data: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v2 data with params", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Data: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v2 wildfire", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Wildfire: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v2 wildfire with headers", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Wildfire: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v2 wildfire with params", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Wildfire: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v2 tunnel", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Tunnel: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v2 tunnel with headers", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Tunnel: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v2 tunnel with params", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Tunnel: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v2 user id", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			UserId: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v2 user id with headers", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			UserId: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v2 user id with params", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			UserId: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v2 gtp", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Gtp: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v2 gtp with headers", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Gtp: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v2 gtp with params", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Gtp: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v2 auth", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Auth: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v2 auth with headers", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Auth: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v2 auth with params", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Auth: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v2 sctp", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Sctp: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v2 sctp with headers", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Sctp: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v2 sctp with params", version.Number{8, 1, 0, ""}, Entry{
			Name: "foo",
			Sctp: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},

		// v3
		{"v3 tag reg", version.Number{9, 0, 0, ""}, Entry{
			Name:            "foo",
			TagRegistration: true,
		}},
		{"v3 with servers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Servers: []Server{
				{
					Name:               "t1",
					Address:            "http.example.com",
					Protocol:           ProtocolHttps,
					Port:               443,
					HttpMethod:         "POST",
					Username:           "someuser",
					Password:           "somepasswd",
					TlsVersion:         "1.2",
					CertificateProfile: "cert-profile",
				},
			},
		}},
		{"v3 config", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Config: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v3 config with headers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Config: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v3 config with params", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Config: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v3 system", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			System: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v3 system with headers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			System: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v3 system with params", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			System: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v3 threat", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Threat: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v3 threat with headers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Threat: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v3 threat with params", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Threat: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v3 traffic", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Traffic: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v3 traffic with headers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Traffic: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v3 traffic with params", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Traffic: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v3 hip match", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			HipMatch: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v3 hip match with headers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			HipMatch: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v3 hip match with params", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			HipMatch: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v3 url", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Url: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v3 url with headers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Url: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v3 url with params", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Url: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v3 data", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Data: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v3 data with headers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Data: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v3 data with params", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Data: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v3 wildfire", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Wildfire: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v3 wildfire with headers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Wildfire: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v3 wildfire with params", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Wildfire: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v3 tunnel", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Tunnel: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v3 tunnel with headers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Tunnel: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v3 tunnel with params", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Tunnel: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v3 user id", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			UserId: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v3 user id with headers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			UserId: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v3 user id with params", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			UserId: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v3 gtp", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Gtp: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v3 gtp with headers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Gtp: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v3 gtp with params", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Gtp: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v3 auth", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Auth: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v3 auth with headers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Auth: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v3 auth with params", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Auth: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v3 sctp", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Sctp: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v3 sctp with headers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Sctp: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v3 sctp with params", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Sctp: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
		{"v3 iptag", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Iptag: &PayloadFormat{
				Name:      "the name",
				UriFormat: "uri format",
				Payload:   "payload",
			},
		}},
		{"v3 iptag with headers", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Iptag: &PayloadFormat{
				Name: "the name",
				Headers: []Header{
					{
						Name:  "Content-Type",
						Value: "application/json",
					},
				},
			},
		}},
		{"v3 iptag with params", version.Number{9, 0, 0, ""}, Entry{
			Name: "foo",
			Iptag: &PayloadFormat{
				Name: "the name",
				Parameters: []Parameter{
					{
						Name:  "SplunkId",
						Value: "secret",
					},
				},
			},
		}},
	}
}
