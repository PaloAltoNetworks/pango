package util

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/PaloAltoNetworks/pango/plugin"
	"github.com/PaloAltoNetworks/pango/version"
	"github.com/PaloAltoNetworks/pango/xmlapi"
)

// PangoClient interface.
type PangoClient interface {
	// Basics.
	Versioning() version.Number
	Plugins() []plugin.Info
	GetTarget() string
	IsPanorama() (bool, error)
	IsFirewall() (bool, error)

	// Local inspection mode functions.
	ReadFromConfig(context.Context, []string, bool, any) ([]byte, error)

	// Communication functions.
	StartJob(context.Context, PangoCommand) (uint, []byte, *http.Response, error)
	Communicate(context.Context, PangoCommand, bool, any) ([]byte, *http.Response, error)

	// Polling functions.
	WaitForJob(context.Context, uint, time.Duration, any) error
	WaitForLogs(context.Context, uint, time.Duration, any) ([]byte, error)

	// Specialized communication functions around specific XPI API commands.
	MultiConfig(context.Context, *xmlapi.MultiConfig, bool, url.Values) ([]byte, *http.Response, *xmlapi.MultiConfigResponse, error)
	ImportFile(context.Context, *xmlapi.Import, string, string, string, bool, any) ([]byte, *http.Response, error)
	ExportFile(context.Context, *xmlapi.Export, any) (string, []byte, error)

	// Operational functions in use by one or more resources / data sources / namespaces.
	RequestPasswordHash(context.Context, string) (string, error)
	GetTechSupportFile(context.Context) (string, []byte, error)
}
