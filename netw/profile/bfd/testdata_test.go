package bfd

type testCase struct {
	desc string
	conf Entry
}

func getTests() []testCase {
	return []testCase{
		{"test1", Entry{
			Name:                "test1",
			Mode:                ModeActive,
			MinimumTxInterval:   3,
			MinimumRxInterval:   4,
			DetectionMultiplier: 5,
			HoldTime:            6,
		}},
		{"test2", Entry{
			Name:         "test2",
			MinimumRxTtl: 7,
		}},
		{"test3", Entry{
			Name:     "test3",
			Mode:     ModePassive,
			HoldTime: 42,
		}},
	}
}
