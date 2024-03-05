package util

import (
	"context"
	"net/http"
	"net/url"

	"github.com/PaloAltoNetworks/pango/plugin"
	"github.com/PaloAltoNetworks/pango/version"
	"github.com/PaloAltoNetworks/pango/xmlapi"
)

type PangoClient interface {
	Versioning() version.Number
	GetTarget() string
	Plugins() []plugin.Info
	MultiConfig(context.Context, *xmlapi.MultiConfig, bool, url.Values) ([]byte, *http.Response, *xmlapi.MultiConfigResponse, error)
	Communicate(context.Context, PangoCommand, bool, any) ([]byte, *http.Response, error)
	CommunicateFile(context.Context, string, string, string, url.Values, bool, any) ([]byte, *http.Response, error)
	ReadFromConfig(context.Context, []string, bool, any) ([]byte, error)
}
