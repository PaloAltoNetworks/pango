package filtering

import (
	"github.com/PaloAltoNetworks/pango/v2/errors"
)

var (
	_ Matcher = &IsNotNil{}
)

type IsNotNil struct {
	Path string

	Operator *Operator
}

func (o *IsNotNil) Matches(f Fielder) (bool, error) {
	x, err := f.Field(o.Path)
	if err != nil {
		return false, err
	}

	switch v := x.(type) {
	case *bool:
		return v != nil, nil
	case bool:
		return true, nil
	case *int64:
		return v != nil, nil
	case int64:
		return true, nil
	case *float64:
		return v != nil, nil
	case float64:
		return true, nil
	case *string:
		return v != nil, nil
	case string:
		return true, nil
	case []string:
		return len(v) != 0, nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *IsNotNil) GetOperator() *Operator { return o.Operator }
