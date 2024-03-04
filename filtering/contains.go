package filtering

import (
	"strings"

	"github.com/PaloAltoNetworks/pango/errors"
)

var (
	_ Matcher = &Contains{}
)

type Contains struct {
	Path  string
	Value string

	Operator *Operator
}

func (o *Contains) Matches(f Fielder) (bool, error) {
	x, err := f.Field(o.Path)
	if err != nil {
		return false, err
	}

	switch v := x.(type) {
	case *string:
		if v == nil {
			return false, nil
		}

		return strings.Contains(*v, o.Value), nil
	case string:
		return strings.Contains(v, o.Value), nil
	case []*string:
		for _, str := range v {
			if str == nil {
				continue
			}

			if strings.Contains(*str, o.Value) {
				return true, nil
			}
		}

		return false, nil
	case []string:
		for _, str := range v {
			if strings.Contains(str, o.Value) {
				return true, nil
			}
		}

		return false, nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *Contains) GetOperator() *Operator { return o.Operator }
