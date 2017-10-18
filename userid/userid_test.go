package userid

import (
    "strings"
    "testing"

    "github.com/PaloAltoNetworks/xapi/testdata"
    "github.com/PaloAltoNetworks/xapi/version"
)


func TestRun(t *testing.T) {
    testCases := []struct{
        logins, logouts map[string] string
        reg, unreg map[string] []string
        vsys, desc string
    }{
        {map[string] string{"john": "1.2.3.4"}, nil, nil, nil, "", "login and empty vsys"},
        {nil, map[string] string{"jack": "2.3.4.5"}, nil, nil, "vsys2", "logout and vsys2"},
        {nil, nil, map[string] []string{"3.4.5.6": []string{"one", "two"}}, nil, "vsys3", "register and vsys3"},
        {nil, nil, nil, map[string] []string{"4.5.6.7": []string{"three", "four"}}, "vsys4", "unregister and vsys4"},
    }
    mc := &testdata.MockClient{Resp: []testdata.Response{
        testdata.Response{[]byte(""), nil},
        testdata.Response{[]byte(""), nil},
        testdata.Response{[]byte(""), nil},
        testdata.Response{[]byte(""), nil},
    }}
    u := &UserId{}
    u.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            i := mc.Called
            err := u.Run(tc.logins, tc.logouts, tc.reg, tc.unreg, tc.vsys)
            if err != nil || mc.Called == i {
                t.Errorf("Failed basic checks")
            } else {
                if tc.vsys == "" {
                    if mc.Vsys != "vsys1" {
                        t.Errorf("Vsys is %s, not vsys1", mc.Vsys)
                    }
                } else if mc.Vsys != tc.vsys {
                        t.Errorf("Vsys is %s, not %s", mc.Vsys, tc.vsys)
                }
            }
        })
    }
}

func TestRegistered(t *testing.T) {
    testCases := []struct{
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
            _, err := u.Registered("", "", "vsys1")
            if err != nil || mc.Called == i || mc.Vsys != "vsys1" || strings.Index(mc.Elm, tc.r) == -1 {
                t.Fail()
            }
        })
    }
}
