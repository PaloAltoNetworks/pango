package orcond

import (
	"github.com/PaloAltoNetworks/pango/version"
)

type tc struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []tc {
	return []tc{
		{"v1 pattern match no qual", version.Number{7, 1, 0, ""}, Entry{
			Name:     "v1",
			Operator: OperatorPatternMatch,
			Context:  "context",
			Pattern:  "myregexhere",
		}},
		{"v1 pattern match with qual", version.Number{7, 1, 0, ""}, Entry{
			Name:     "v1",
			Operator: OperatorPatternMatch,
			Context:  "context",
			Pattern:  "myregexhere",
			Qualifiers: map[string]string{
				"qual1": "value1",
				"qual2": "value2",
			},
		}},
		{"v1 greater than no qual", version.Number{7, 1, 0, ""}, Entry{
			Name:     "v1",
			Operator: OperatorGreaterThan,
			Context:  "context",
			Value:    "gt value",
		}},
		{"v1 greater than with qual", version.Number{7, 1, 0, ""}, Entry{
			Name:     "v1",
			Operator: OperatorGreaterThan,
			Context:  "context",
			Value:    "gt value",
			Qualifiers: map[string]string{
				"qual1": "value1",
				"qual2": "value2",
			},
		}},
		{"v1 less than no qual", version.Number{7, 1, 0, ""}, Entry{
			Name:     "v1",
			Operator: OperatorLessThan,
			Context:  "context",
			Value:    "lt value",
		}},
		{"v1 less than with qual", version.Number{7, 1, 0, ""}, Entry{
			Name:     "v1",
			Operator: OperatorLessThan,
			Context:  "context",
			Value:    "lt value",
			Qualifiers: map[string]string{
				"qual1": "value1",
				"qual2": "value2",
			},
		}},
		{"v1 equal to no extras", version.Number{7, 1, 0, ""}, Entry{
			Name:     "v1",
			Operator: OperatorEqualTo,
			Context:  "context",
			Value:    "my value",
		}},
		{"v1 equal to with extras", version.Number{7, 1, 0, ""}, Entry{
			Name:     "v1",
			Operator: OperatorEqualTo,
			Context:  "context",
			Value:    "my value",
			Position: "position",
			Mask:     "0a0b",
		}},
	}
}
