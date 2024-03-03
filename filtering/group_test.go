package filtering

import (
	"testing"
)

func TestSingleGroup(t *testing.T) {
	g := &Group{
		Matchers: []Matcher{
			&Equals{
				Value: "7",
			},
		},
	}

	f := &testFielder{
		Answers: []testFielderResponse{
			{Value: int64(7)},
		},
	}

	ans, err := g.Matches(f)
	if err != nil {
		t.Fatalf("err from matches: %s", err)
	}
	if !ans {
		t.Fatalf("ans was false")
	}
}

func TestNegateGroup(t *testing.T) {
	g := &Group{
		NegateGroup: true,
		Matchers: []Matcher{
			&Equals{
				Value: "7",
			},
		},
	}

	f := &testFielder{
		Answers: []testFielderResponse{
			{Value: int64(7)},
		},
	}

	ans, err := g.Matches(f)
	if err != nil {
		t.Fatalf("err from matches: %s", err)
	}
	if ans {
		t.Fatalf("ans was true")
	}
}

func TestGroupWithinGroup(t *testing.T) {
	g := &Group{
		Matchers: []Matcher{
			&Group{
				Matchers: []Matcher{
					&EndsWith{
						Value: "e",
					},
				},
			},
		},
	}

	f := &testFielder{
		Answers: []testFielderResponse{
			{Value: "this is fine"},
		},
	}

	ans, err := g.Matches(f)
	if err != nil {
		t.Fatalf("err from matches: %s", err)
	}
	if !ans {
		t.Fatalf("ans was false")
	}
}

func TestNegateGroupWithinGroup(t *testing.T) {
	g := &Group{
		Matchers: []Matcher{
			&Group{
				NegateGroup: true,
				Matchers: []Matcher{
					&EndsWith{
						Value: "e",
					},
				},
			},
		},
	}

	f := &testFielder{
		Answers: []testFielderResponse{
			{Value: "this is fine"},
		},
	}

	ans, err := g.Matches(f)
	if err != nil {
		t.Fatalf("err from matches: %s", err)
	}
	if ans {
		t.Fatalf("ans was true")
	}
}

func TestGroupWithTwoChecksAndOperatorReturnsTrue(t *testing.T) {
	g := &Group{
		Matchers: []Matcher{
			&LessThan{
				Value: "10",
			},
			&GreaterThan{
				Operator: &Operator{
					And: true,
				},
				Value: "0",
			},
		},
	}

	f := &testFielder{
		Answers: []testFielderResponse{
			{Value: int64(5)},
			{Value: float64(5)},
		},
	}

	ans, err := g.Matches(f)
	if err != nil {
		t.Fatalf("err in matches: %s", err)
	}
	if !ans {
		t.Fatalf("ans was false")
	}
}

func TestGroupWithTwoChecksOrOperatorReturnsTrueSecondIsTrue(t *testing.T) {
	g := &Group{
		Matchers: []Matcher{
			&LessThan{
				Value: "10",
			},
			&GreaterThan{
				Operator: &Operator{
					Or: true,
				},
				Value: "0",
			},
		},
	}

	f := &testFielder{
		Answers: []testFielderResponse{
			{Value: int64(11)},
			{Value: float64(5)},
		},
	}

	ans, err := g.Matches(f)
	if err != nil {
		t.Fatalf("err in matches: %s", err)
	}
	if !ans {
		t.Fatalf("ans was false")
	}
}

func TestGroupWithTwoChecksOrOperatorReturnsTrueFirstIsTrue(t *testing.T) {
	g := &Group{
		Matchers: []Matcher{
			&LessThan{
				Value: "10",
			},
			&GreaterThan{
				Operator: &Operator{
					Or: true,
				},
				Value: "0",
			},
		},
	}

	f := &testFielder{
		Answers: []testFielderResponse{
			{Value: int64(0)},
			{Value: float64(0)},
		},
	}

	ans, err := g.Matches(f)
	if err != nil {
		t.Fatalf("err in matches: %s", err)
	}
	if !ans {
		t.Fatalf("ans was false")
	}
}

func TestGroupWithTwoChecksXorOperatorReturnsTrue(t *testing.T) {
	var v *string
	g := &Group{
		Matchers: []Matcher{
			&IsNil{},
			&GreaterThan{
				Operator: &Operator{
					Xor: true,
				},
				Value: "0",
			},
		},
	}

	f := &testFielder{
		Answers: []testFielderResponse{
			{Value: v},
			{Value: float64(0)},
		},
	}

	ans, err := g.Matches(f)
	if err != nil {
		t.Fatalf("err in matches: %s", err)
	}
	if !ans {
		t.Fatalf("ans was false")
	}
}

func TestGroupWithTwoChecksXorOperatorReturnsFalse(t *testing.T) {
	var v *int64
	g := &Group{
		Matchers: []Matcher{
			&IsNil{},
			&GreaterThan{
				Operator: &Operator{
					Xor: true,
				},
				Value: "0",
			},
		},
	}

	f := &testFielder{
		Answers: []testFielderResponse{
			{Value: v},
			{Value: float64(1)},
		},
	}

	ans, err := g.Matches(f)
	if err != nil {
		t.Fatalf("err in matches: %s", err)
	}
	if ans {
		t.Fatalf("ans was true")
	}
}
