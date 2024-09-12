package util

import "github.com/PaloAltoNetworks/pango/version"

type ILocation interface {
	XpathPrefix(version.Number) ([]string, error)
}
