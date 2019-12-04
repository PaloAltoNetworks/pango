package userid

import (
    "testing"

    "github.com/PaloAltoNetworks/pango/testdata"
)


func TestGetGroups(t *testing.T) {
    testCases := []struct{
        desc string
        groups []string
        body string
    }{
        {"no groups", nil, `<response status="success"><result><![CDATA[

Total: 0
* : Custom Group

]]></result></response>`},
        {"with groups", []string{"malicious_users", "foo users"}, `<response status="success"><result><![CDATA[
malicious_users 
foo users 

Total: 2
* : Custom Group

]]></result></response>`},
    }

    mc := &testdata.MockClient{}
    mc.Resp = make([]testdata.Response, 0, len(testCases))
    u := &UserId{}
    u.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            i := mc.Called
            mc.Resp = append(mc.Resp, testdata.Response{[]byte(tc.body), nil})
            ans, err := u.GetGroups("", "vsys1")
            if err != nil || mc.Called == i {
                t.Errorf("Failed basic checks")
            } else if len(ans) != len(tc.groups) {
                t.Errorf("Mismatched groups: expected %#v, got %#v", tc.groups, ans)
            } else {
                for j := range ans {
                    if ans[j] != tc.groups[j] {
                        t.Errorf("Mismatched group (%d): expected %s, got %s", i, tc.groups[j], ans[j])
                    }
                }
            }
        })
    }
}

func TestGetGroupMembers(t *testing.T) {
    testCases := []struct{
        desc string
        users []string
        body string
    }{
        {"no members", nil, `<response status="success"><result><![CDATA[

source type: xmlapi
Group type: Dynamic


]]></result></response>`},
        {"with members", []string{"london", "elektricity"}, `<response status="success"><result><![CDATA[

source type: xmlapi

[1     ] london
[2     ] elektricity

]]></result></response>`},
    }

    mc := &testdata.MockClient{}
    mc.Resp = make([]testdata.Response, 0, len(testCases))
    u := &UserId{}
    u.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            i := mc.Called
            mc.Resp = append(mc.Resp, testdata.Response{[]byte(tc.body), nil})
            ans, err := u.GetGroupMembers("foo", "vsys1")
            if err != nil || mc.Called == i {
                t.Errorf("Failed basic checks")
            } else if len(ans) != len(tc.users) {
                t.Errorf("Mismatched groups: expected %#v, got %#v", tc.users, ans)
            } else {
                for j := range ans {
                    if ans[j] != tc.users[j] {
                        t.Errorf("Mismatched group (%d): expected %s, got %s", i, tc.users[j], ans[j])
                    }
                }
            }
        })
    }
}
