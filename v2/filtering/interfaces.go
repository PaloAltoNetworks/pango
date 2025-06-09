package filtering

type Fielder interface {
	Field(string) (any, error)
}

type Matcher interface {
	Matches(Fielder) (bool, error)
	GetOperator() *Operator
}
