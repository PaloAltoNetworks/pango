package filtering

type testFielderResponse struct {
	Value any
	Error error
}

type testFielder struct {
	Answers []testFielderResponse

	index int
}

func (o *testFielder) Field(_ string) (any, error) {
	ans := o.Answers[o.index]
	o.index++
	return ans.Value, ans.Error
}
