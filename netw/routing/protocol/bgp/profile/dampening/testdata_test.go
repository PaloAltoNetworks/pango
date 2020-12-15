package dampening

type testCase struct {
	desc string
	conf Entry
}

func getTests() []testCase {
	return []testCase{
		{"with enable", Entry{
			Name:                     "with enable",
			Enable:                   true,
			Cutoff:                   1.25,
			Reuse:                    0.5,
			MaxHoldTime:              900,
			DecayHalfLifeReachable:   300,
			DecayHalfLifeUnreachable: 900,
		}},
		{"without enable", Entry{
			Name:                     "without enable",
			Enable:                   false,
			Cutoff:                   2,
			Reuse:                    500,
			MaxHoldTime:              600,
			DecayHalfLifeReachable:   700,
			DecayHalfLifeUnreachable: 800,
		}},
	}
}
