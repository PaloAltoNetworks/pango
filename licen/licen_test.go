package licen

import (
	"fmt"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
)

func TestCurrent(t *testing.T) {
	mc := &testdata.MockClient{Resp: []testdata.Response{
		testdata.Response{[]byte(testdata.LicenseXml), nil},
	}}
	l := &Licen{}
	l.Initialize(mc)

	ans, err := l.Current()
	if err != nil {
		t.Fail()
	}
	if len(ans) != 2 {
		t.Fail()
	}
}

func TestAuthCode(t *testing.T) {
	testCases := []struct {
		desc       string
		shouldFail bool
		resp       testdata.Response
	}{
		{"fail xml response", true, testdata.Response{
			[]byte("<response status=\"error\"><msg><line>Failed to install licenses. Invalid Auth Code: foo</line></msg></response>"),
			fmt.Errorf("Failed to install licenses. Invalid Auth Code: foo"),
		}},
		{"success non xml response", false, testdata.Response{
			[]byte("VM Device License installed. Restarting pan services."),
			fmt.Errorf("EOF"),
		}},
		{"other non xml response", true, testdata.Response{
			[]byte("foobar"),
			fmt.Errorf("EOF"),
		}},
	}

	ns := &Licen{}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mc := &testdata.MockClient{Resp: []testdata.Response{tc.resp}}
			ns.Initialize(mc)

			err := ns.Activate("foo")
			if tc.shouldFail && err == nil {
				t.Errorf("Expected failure, but got nil error back")
			} else if !tc.shouldFail && err != nil {
				t.Errorf("Expected success, but got error: %s", err)
			}
		})
	}
}

func BenchmarkCurrent(b *testing.B) {
	mc := &testdata.MockClient{Resp: []testdata.Response{
		testdata.Response{[]byte(testdata.LicenseXml), nil},
	}}
	l := &Licen{}
	l.Initialize(mc)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = l.Current()
	}
}
