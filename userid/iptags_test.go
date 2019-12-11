package userid

import (
	"strings"
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
	"github.com/PaloAltoNetworks/pango/version"
)

func TestGetIpTags(t *testing.T) {
	testCases := []struct {
		n version.Number
		r string
	}{
		{version.Number{6, 0, 0, ""}, "registered-address"},
		{version.Number{7, 0, 0, ""}, "registered-ip"},
	}
	mc := &testdata.MockClient{Resp: []testdata.Response{
		testdata.Response{[]byte(testdata.UserIdXml), nil},
		testdata.Response{[]byte(testdata.UserIdXml), nil},
	}}
	u := &UserId{}
	u.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.r, func(t *testing.T) {
			mc.Version = tc.n
			i := mc.Called
			_, err := u.GetIpTags("", "", "vsys1")
			if err != nil || mc.Called == i || mc.Vsys != "vsys1" || strings.Index(mc.Elm, tc.r) == -1 {
				t.Fail()
			}
		})
	}
}
