package filtering

import (
	"github.com/PaloAltoNetworks/pango/errors"
)

var (
	_ Matcher = &Group{}
)

type Group struct {
	NegateGroup bool

	Matchers []Matcher
	Operator *Operator
}

func (o *Group) Matches(f Fielder) (bool, error) {
	if o.Matchers == nil {
		return false, errors.InvalidFilterError
	}

	var ans bool

	for num, m := range o.Matchers {
		op := m.GetOperator()
		if num == 0 && op != nil {
			return false, errors.InvalidFilterError
		} else if num != 0 && op == nil {
			return false, errors.InvalidFilterError
		}

		ma, err := m.Matches(f)
		if err != nil {
			return false, err
		}

		if op == nil {
			ans = ma
		} else {
			switch {
			case op.And:
				ans = ans && ma
			case op.Or:
				ans = ans || ma
			case op.Xor:
				ans = ans != ma
			default:
				return false, errors.InvalidFilterError
			}
		}
	}

	if o.NegateGroup {
		ans = !ans
	}

	return ans, nil
}

func (o *Group) GetOperator() *Operator { return o.Operator }
