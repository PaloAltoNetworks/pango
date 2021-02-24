package mngtprof

type testCase struct {
	desc string
	conf Entry
}

func getTests() []testCase {
	return []testCase{
		{"test1", Entry{
			Name:                    "test1",
			Ping:                    true,
			Telnet:                  true,
			Ssh:                     true,
			Https:                   true,
			UseridSyslogListenerSsl: true,
		}},
		{"test2", Entry{
			Name:                    "test2",
			Http:                    true,
			HttpOcsp:                true,
			Snmp:                    true,
			ResponsePages:           true,
			UseridService:           true,
			UseridSyslogListenerUdp: true,
		}},
		{"test3", Entry{
			Name:         "test3",
			Ping:         true,
			PermittedIps: []string{"10.1.1.1", "10.2.2.2"},
		}},
	}
}
