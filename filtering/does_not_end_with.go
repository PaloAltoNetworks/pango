package filtering

import (
	"strings"

	"github.com/PaloAltoNetworks/pango/errors"
)

var (
	_ Matcher = &DoesNotEndWith{}
)

type DoesNotEndWith struct {
	Path  string
	Value string

	Operator *Operator
}

func (o *DoesNotEndWith) Matches(f Fielder) (bool, error) {
	x, err := f.Field(o.Path)
	if err != nil {
		return false, err
	}

	switch v := x.(type) {
	case *string:
		if v == nil {
			return true, nil
		}

		return !strings.HasSuffix(*v, o.Value), nil
	case string:
		return !strings.HasSuffix(v, o.Value), nil
	case []string:
		if len(v) == 0 {
			return true, nil
		}

		return v[len(v)-1] != o.Value, nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *DoesNotEndWith) GetOperator() *Operator { return o.Operator }
