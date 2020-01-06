package userid

import (
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
)

func TestGetUserTags(t *testing.T) {
	testCases := []struct {
		desc string
		data map[string][]string
		body string
	}{
		{
			"multi user", map[string][]string{
				"mydomain\\user1": []string{"tag1", "tag2"},
				"mydomain\\user2": []string{"tag3", "tag4"},
			}, `<response status="success"><result><entry user="mydomain\user1">
<tag>
<member>tag1</member>
<member>tag2</member>
</tag>
</entry>
<entry user="mydomain\user2">
<tag>
<member>tag3</member>
<member>tag4</member>
</tag>
</entry>
<count>2</count>
</result></response>`,
		},
		{
			"single user", map[string][]string{
				"foo\\jack": []string{"one", "two", "three"},
			}, `
<response status="success"><result><entry user="foo\jack">
<tag>
<member>one</member>
<member>two</member>
<member>three</member>
</tag>
</entry>
<count>1</count>
</result></response>`,
		},
	}

	mc := &testdata.MockClient{}
	mc.Resp = make([]testdata.Response, 0, len(testCases))
	u := &UserId{}
	u.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			i := mc.Called
			mc.Resp = append(mc.Resp, testdata.Response{[]byte(tc.body), nil})
			ans, err := u.GetUserTags("", "vsys1")
			if err != nil || mc.Called == i {
				t.Errorf("Failed basic checks")
			} else if len(ans) != len(tc.data) {
				t.Errorf("Mismatched data: expected %#v, got %#v", tc.data, ans)
			} else {
				for user, eTags := range tc.data {
					var aTags []string
					aTags = ans[user]
					if len(aTags) != len(eTags) {
						t.Errorf("Mismatched tags for %s: %#v, but got %#v", user, eTags, aTags)
					} else {
						for ti := range eTags {
							if eTags[ti] != aTags[ti] {
								t.Errorf("Tag %d: wanted %s, got %s", ti, eTags[ti], aTags[ti])
								break
							}
						}
					}
				}
			}
		})
	}
}
