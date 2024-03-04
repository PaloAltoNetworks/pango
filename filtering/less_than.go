package filtering

import (
	"strconv"

	"github.com/PaloAltoNetworks/pango/errors"
)

var (
	_ Matcher = &LessThan{}
)

type LessThan struct {
	Path  string
	Value string

	Operator *Operator
}

func (o *LessThan) Matches(f Fielder) (bool, error) {
	x, err := f.Field(o.Path)
	if err != nil {
		return false, err
	}

	switch v := x.(type) {
	case *int64:
		if v == nil {
			return false, nil
		}

		vi, err := strconv.ParseInt(o.Value, 10, 64)
		if err != nil {
			return false, err
		}

		return *v < vi, nil
	case int64:
		vi, err := strconv.ParseInt(o.Value, 10, 64)
		if err != nil {
			return false, err
		}

		return v < vi, nil
	case []*int64:
		vi, err := strconv.ParseInt(o.Value, 10, 64)
		if err != nil {
			return false, err
		}

		for _, val := range v {
			if val == nil {
				continue
			}
			if *val < vi {
				return true, nil
			}
		}

		return false, nil
	case []int64:
		vi, err := strconv.ParseInt(o.Value, 10, 64)
		if err != nil {
			return false, err
		}

		for _, val := range v {
			if val < vi {
				return true, nil
			}
		}

		return false, nil
	case *float64:
		if v == nil {
			return false, nil
		}

		vf, err := strconv.ParseFloat(o.Value, 64)
		if err != nil {
			return false, err
		}

		return *v < vf, nil
	case float64:
		vf, err := strconv.ParseFloat(o.Value, 64)
		if err != nil {
			return false, err
		}

		return v < vf, nil
	case []*float64:
		vf, err := strconv.ParseFloat(o.Value, 64)
		if err != nil {
			return false, err
		}

		for _, val := range v {
			if val == nil {
				continue
			}
			if *val < vf {
				return true, nil
			}
		}

		return false, nil
	case []float64:
		vf, err := strconv.ParseFloat(o.Value, 64)
		if err != nil {
			return false, err
		}

		for _, val := range v {
			if val < vf {
				return true, nil
			}
		}

		return false, nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *LessThan) GetOperator() *Operator { return o.Operator }
