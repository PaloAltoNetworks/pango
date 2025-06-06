package filtering

import (
	"strings"

	"github.com/PaloAltoNetworks/pango/v2/errors"
)

var (
	_ Matcher = &EndsWith{}
)

type EndsWith struct {
	Path  string
	Value string

	Operator *Operator
}

func (o *EndsWith) Matches(f Fielder) (bool, error) {
	x, err := f.Field(o.Path)
	if err != nil {
		return false, err
	}

	switch v := x.(type) {
	case *string:
		if v == nil {
			return false, nil
		}

		return strings.HasSuffix(*v, o.Value), nil
	case string:
		return strings.HasSuffix(v, o.Value), nil
	case []*string:
		for _, str := range v {
			if str == nil {
				continue
			}
			if strings.HasSuffix(*str, o.Value) {
				return true, nil
			}
		}

		return false, nil
	case []string:
		for _, str := range v {
			if strings.HasSuffix(str, o.Value) {
				return true, nil
			}
		}

		return false, nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *EndsWith) GetOperator() *Operator { return o.Operator }
