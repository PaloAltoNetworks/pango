package ipv4

type testCase struct {
	desc string
	conf Entry
}

func getTests() []testCase {
	return []testCase{
		{"any proto", Entry{
			Name:        "test any",
			Local:       "local",
			Remote:      "remote",
			ProtocolAny: true,
		}},
		{"number proto", Entry{
			Name:           "test any",
			Local:          "local",
			Remote:         "remote",
			ProtocolNumber: 42,
		}},
		{"tcp proto", Entry{
			Name:              "test any",
			Local:             "local",
			Remote:            "remote",
			ProtocolTcpLocal:  1,
			ProtocolTcpRemote: 2,
		}},
		{"udp proto", Entry{
			Name:              "test any",
			Local:             "local",
			Remote:            "remote",
			ProtocolUdpLocal:  3,
			ProtocolUdpRemote: 4,
		}},
	}
}
