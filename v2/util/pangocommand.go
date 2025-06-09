package util

import (
	"net/url"
)

type PangoCommand interface {
	AsUrlValues() (url.Values, error)
}
