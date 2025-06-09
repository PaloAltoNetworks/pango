package filtering

import (
	"strings"

	"github.com/PaloAltoNetworks/pango/v2/errors"
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
	case []*string:
		for _, val := range v {
			if val == nil {
				continue
			}
			if strings.HasPrefix(*val, o.Value) {
				return true, nil
			}
		}

		return false, nil
	case []string:
		for _, val := range v {
			if strings.HasPrefix(val, o.Value) {
				return true, nil
			}
		}

		return false, nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *StartsWith) GetOperator() *Operator { return o.Operator }
