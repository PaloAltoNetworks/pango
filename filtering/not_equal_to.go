package filtering

import (
	"strconv"

	"github.com/PaloAltoNetworks/pango/errors"
)

var (
	_ Matcher = &NotEqualTo{}
)

type NotEqualTo struct {
	Path  string
	Value string

	Operator *Operator
}

func (o *NotEqualTo) Matches(f Fielder) (bool, error) {
	x, err := f.Field(o.Path)
	if err != nil {
		return false, err
	}

	switch v := x.(type) {
	case *bool:
		if v == nil {
			return true, nil
		}

		bv, err := strconv.ParseBool(o.Value)
		if err != nil {
			return false, err
		}

		return *v != bv, nil
	case bool:
		bv, err := strconv.ParseBool(o.Value)
		if err != nil {
			return false, err
		}

		return v != bv, nil
	case *int64:
		if v == nil {
			return true, nil
		}

		iv, err := strconv.ParseInt(o.Value, 10, 64)
		if err != nil {
			return false, err
		}

		return *v != iv, nil
	case int64:
		iv, err := strconv.ParseInt(o.Value, 10, 64)
		if err != nil {
			return false, err
		}

		return v != iv, nil
	case *float64:
		if v == nil {
			return true, nil
		}

		fv, err := strconv.ParseFloat(o.Value, 64)
		if err != nil {
			return false, err
		}

		return *v != fv, nil
	case float64:
		fv, err := strconv.ParseFloat(o.Value, 64)
		if err != nil {
			return false, err
		}

		return v != fv, nil
	case *string:
		if v == nil {
			return true, nil
		}

		return *v != o.Value, nil
	case string:
		return v != o.Value, nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *NotEqualTo) GetOperator() *Operator { return o.Operator }
