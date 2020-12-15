package group

type testCase struct {
	desc string
	conf Entry
}

func getTests() []testCase {
	return []testCase{
		{"ebgp check one", Entry{
			Name:                    "ebgpOne",
			Enable:                  true,
			AggregatedConfedAsPath:  false,
			SoftResetWithStoredInfo: true,
			Type:                    TypeEbgp,
			ExportNextHop:           NextHopResolve,
			ImportNextHop:           NextHopUsePeer,
			RemovePrivateAs:         false,
		}},
		{"ebgp check two with raw info", Entry{
			Name:                    "ebgpTwo",
			Enable:                  false,
			AggregatedConfedAsPath:  true,
			SoftResetWithStoredInfo: false,
			Type:                    TypeEbgp,
			ExportNextHop:           NextHopUseSelf,
			ImportNextHop:           NextHopOriginal,
			RemovePrivateAs:         true,
			raw: map[string]string{
				"peer": "peer info",
			},
		}},
		{"ebgp confed", Entry{
			Name:                    "ebgpConfed",
			Enable:                  true,
			AggregatedConfedAsPath:  true,
			SoftResetWithStoredInfo: true,
			Type:                    TypeEbgpConfed,
			ExportNextHop:           NextHopOriginal,
		}},
		{"ibgp", Entry{
			Name:                    "ibgp",
			Enable:                  true,
			AggregatedConfedAsPath:  true,
			SoftResetWithStoredInfo: true,
			Type:                    TypeIbgp,
			ExportNextHop:           NextHopOriginal,
		}},
		{"ibgp confed", Entry{
			Name:                    "ibgpConfed",
			Enable:                  true,
			AggregatedConfedAsPath:  true,
			SoftResetWithStoredInfo: true,
			Type:                    TypeIbgpConfed,
			ExportNextHop:           NextHopOriginal,
		}},
	}
}
