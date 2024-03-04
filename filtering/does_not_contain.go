package filtering

import (
	"strings"

	"github.com/PaloAltoNetworks/pango/errors"
)

var (
	_ Matcher = &DoesNotContain{}
)

type DoesNotContain struct {
	Path  string
	Value string

	Operator *Operator
}

func (o *DoesNotContain) Matches(f Fielder) (bool, error) {
	x, err := f.Field(o.Path)
	if err != nil {
		return false, err
	}

	switch v := x.(type) {
	case *string:
		if v == nil {
			return true, nil
		}

		return !strings.Contains(*v, o.Value), nil
	case string:
		return !strings.Contains(v, o.Value), nil
	case []*string:
		for _, str := range v {
			if str == nil {
				continue
			}

			if strings.Contains(*str, o.Value) {
				return false, nil
			}
		}

		return true, nil
	case []string:
		for _, str := range v {
			if !strings.Contains(str, o.Value) {
				return false, nil
			}
		}

		return true, nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *DoesNotContain) GetOperator() *Operator { return o.Operator }
