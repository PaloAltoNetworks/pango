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
	case []*string:
		for _, str := range v {
			if str == nil {
				continue
			}
			if strings.HasSuffix(*str, o.Value) {
				return false, nil
			}
		}

		return true, nil
	case []string:
		for _, str := range v {
			if strings.HasSuffix(str, o.Value) {
				return false, nil
			}
		}

		return true, nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *DoesNotEndWith) GetOperator() *Operator { return o.Operator }
