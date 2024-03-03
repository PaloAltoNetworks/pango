package filtering

import (
	"strings"

	"github.com/PaloAltoNetworks/pango/errors"
)

var (
	_ Matcher = &StartsWith{}
)

type StartsWith struct {
	Path  string
	Value string

	Operator *Operator
}

func (o *StartsWith) Matches(f Fielder) (bool, error) {
	x, err := f.Field(o.Path)
	if err != nil {
		return false, err
	}

	switch v := x.(type) {
	case *string:
		if v == nil {
			return false, nil
		}

		return strings.HasPrefix(*v, o.Value), nil
	case string:
		return strings.HasPrefix(v, o.Value), nil
	case []string:
		if len(v) == 0 {
			return false, nil
		}

		return v[0] == o.Value, nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *StartsWith) GetOperator() *Operator { return o.Operator }
