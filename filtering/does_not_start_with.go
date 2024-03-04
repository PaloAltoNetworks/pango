package filtering

import (
	"strings"

	"github.com/PaloAltoNetworks/pango/errors"
)

var (
	_ Matcher = &DoesNotStartWith{}
)

type DoesNotStartWith struct {
	Path  string
	Value string

	Operator *Operator
}

func (o *DoesNotStartWith) Matches(f Fielder) (bool, error) {
	x, err := f.Field(o.Path)
	if err != nil {
		return false, err
	}

	switch v := x.(type) {
	case *string:
		if v == nil {
			return true, nil
		}

		return !strings.HasPrefix(*v, o.Value), nil
	case string:
		return !strings.HasPrefix(v, o.Value), nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *DoesNotStartWith) GetOperator() *Operator { return o.Operator }
