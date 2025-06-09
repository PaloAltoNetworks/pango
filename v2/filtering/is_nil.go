package filtering

import (
	"github.com/PaloAltoNetworks/pango/v2/errors"
)

var (
	_ Matcher = &IsNil{}
)

type IsNil struct {
	Path string

	Operator *Operator
}

func (o *IsNil) Matches(f Fielder) (bool, error) {
	x, err := f.Field(o.Path)
	if err != nil {
		return false, err
	}

	switch v := x.(type) {
	case *bool:
		return v == nil, nil
	case bool:
		return false, nil
	case *int64:
		return v == nil, nil
	case int64:
		return false, nil
	case *float64:
		return v == nil, nil
	case float64:
		return false, nil
	case *string:
		return v == nil, nil
	case string:
		return false, nil
	case []string:
		return len(v) == 0, nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *IsNil) GetOperator() *Operator { return o.Operator }
