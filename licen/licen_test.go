package licen

import (
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
