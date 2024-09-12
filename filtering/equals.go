package filtering

import (
	"strconv"

	"github.com/PaloAltoNetworks/pango/errors"
)

var (
	_ Matcher = &Equals{}
)

type Equals struct {
	Path  string
	Value string

	Operator *Operator
}

func (o *Equals) Matches(f Fielder) (bool, error) {
	x, err := f.Field(o.Path)
	if err != nil {
		return false, err
	}

	switch v := x.(type) {
	case *bool:
		if v == nil {
			return false, nil
		}

		bv, err := strconv.ParseBool(o.Value)
		if err != nil {
			return false, err
		}

		return *v == bv, nil
	case bool:
		bv, err := strconv.ParseBool(o.Value)
		if err != nil {
			return false, err
		}

		return v == bv, nil
	case []*bool:
		bv, err := strconv.ParseBool(o.Value)
		if err != nil {
			return false, err
		}

		for _, val := range v {
			if val == nil {
				continue
			}
			if *val == bv {
				return true, nil
			}
		}

		return false, nil
	case []bool:
		bv, err := strconv.ParseBool(o.Value)
		if err != nil {
			return false, err
		}

		for _, val := range v {
			if val == bv {
				return true, nil
			}
		}

		return false, nil
	case *int64:
		if v == nil {
			return false, nil
		}

		iv, err := strconv.ParseInt(o.Value, 10, 64)
		if err != nil {
			return false, err
		}

		return *v == iv, nil
	case int64:
		iv, err := strconv.ParseInt(o.Value, 10, 64)
		if err != nil {
			return false, err
		}

		return v == iv, nil
	case []*int64:
		iv, err := strconv.ParseInt(o.Value, 10, 64)
		if err != nil {
			return false, err
		}

		for _, val := range v {
			if val == nil {
				continue
			}
			if *val == iv {
				return true, nil
			}
		}

		return false, nil
	case []int64:
		iv, err := strconv.ParseInt(o.Value, 10, 64)
		if err != nil {
			return false, err
		}

		for _, val := range v {
			if val == iv {
				return true, nil
			}
		}

		return false, nil
	case *float64:
		if v == nil {
			return false, nil
		}

		fv, err := strconv.ParseFloat(o.Value, 64)
		if err != nil {
			return false, err
		}

		return *v == fv, nil
	case float64:
		fv, err := strconv.ParseFloat(o.Value, 64)
		if err != nil {
			return false, err
		}

		return v == fv, nil
	case []*float64:
		fv, err := strconv.ParseFloat(o.Value, 64)
		if err != nil {
			return false, err
		}

		for _, val := range v {
			if val == nil {
				continue
			}
			if *val == fv {
				return true, nil
			}
		}

		return false, nil
	case []float64:
		fv, err := strconv.ParseFloat(o.Value, 64)
		if err != nil {
			return false, err
		}

		for _, val := range v {
			if val == fv {
				return true, nil
			}
		}

		return false, nil
	case *string:
		if v == nil {
			return false, nil
		}

		return *v == o.Value, nil
	case string:
		return v == o.Value, nil
	case []*string:
		for _, val := range v {
			if val == nil {
				continue
			}
			if *val == o.Value {
				return true, nil
			}
		}

		return false, nil
	case []string:
		for _, val := range v {
			if val == o.Value {
				return true, nil
			}
		}

		return false, nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *Equals) GetOperator() *Operator { return o.Operator }
