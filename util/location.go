package util

import "github.com/PaloAltoNetworks/pango/v2/version"

type ILocation interface {
	XpathPrefix(version.Number) ([]string, error)
}
