package filtering

import (
	"strings"

	"github.com/PaloAltoNetworks/pango/errors"
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
	case []string:
		if len(v) == 0 {
			return false, nil
		}

		return v[len(v)-1] == o.Value, nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *EndsWith) GetOperator() *Operator { return o.Operator }
