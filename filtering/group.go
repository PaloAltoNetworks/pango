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

		// First check sets the group's result so far.
		if num == 0 {
			ma, err := m.Matches(f)
			if err != nil {
				return false, err
			}
			ans = ma
			continue
		}

		// Subsequent checks update the result.
		switch {
		case op.And:
			// Don't bother running the check if the result is already false.
			if !ans {
				continue
			}

			ma, err := m.Matches(f)
			if err != nil {
				return false, err
			}
			ans = ma
		case op.Or:
			// Run the check only if the result is currently false.
			if ans {
				continue
			}

			ma, err := m.Matches(f)
			if err != nil {
				return false, err
			}
			ans = ma
		case op.Xor:
			// We always have to run the check here.
			ma, err := m.Matches(f)
			if err != nil {
				return false, err
			}
			ans = ans != ma
		default:
			return false, errors.InvalidFilterError
		}
	}

	if o.NegateGroup {
		ans = !ans
	}

	return ans, nil
}

func (o *Group) GetOperator() *Operator { return o.Operator }
