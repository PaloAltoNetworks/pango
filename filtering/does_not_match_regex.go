package filtering

import (
	"regexp"

	"github.com/PaloAltoNetworks/pango/errors"
)

var (
	_ Matcher = &DoesNotMatchRegex{}
)

type DoesNotMatchRegex struct {
	Path  string
	Value string

	Operator *Operator
}

func (o *DoesNotMatchRegex) Matches(f Fielder) (bool, error) {
	x, err := f.Field(o.Path)
	if err != nil {
		return false, err
	}

	re, err := regexp.Compile(o.Value)
	if err != nil {
		return false, err
	}

	switch v := x.(type) {
	case *string:
		return !re.Match([]byte(*v)), nil
	case string:
		return !re.Match([]byte(v)), nil
	case []string:
		for _, str := range v {
			if re.Match([]byte(str)) {
				return false, nil
			}
		}

		return true, nil
	}

	return false, errors.UnsupportedFilterTypeError
}

func (o *DoesNotMatchRegex) GetOperator() *Operator { return o.Operator }
