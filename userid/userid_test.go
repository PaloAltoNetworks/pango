package userid

import (
	"testing"

	"github.com/PaloAltoNetworks/pango/testdata"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		desc string
		vsys string
		msg  *Message
	}{
		{"login and empty vsys", "", &Message{
			Logins: []Login{
				{User: "john", Ip: "1.2.3.4"},
			},
		}},
		{"logout and vsys2", "vsys2", &Message{
			Logouts: []Logout{
				{User: "jack", Ip: "2.3.4.5"},
			},
		}},
		{"register and vsys3", "vsys3", &Message{
			TagIps: []TagIp{
				{Ip: "3.4.5.6", Tags: []string{"one", "two"}},
			},
		}},
		{"unregister and vsys4", "vsys4", &Message{
			UntagIps: []UntagIp{
				{Ip: "4.5.6.7", Tags: []string{"three", "four"}},
			},
		}},
		{"group and vsys1", "vsys1", &Message{
			Groups: []Group{
				{Name: "mygroup", Users: []string{"john", "doe"}},
			},
		}},
		{"tag user and vsys2", "vsys2", &Message{
			TagUsers: []TagUser{
				{
					User: "mydomain\\jack",
					Tags: []UserTag{
						{Tag: "one"},
						{Tag: "two", Timeout: 60},
					},
				},
			},
		}},
		{"untag user and vsys3", "vsys3", &Message{
			UntagUsers: []UntagUser{
				{User: "mydomain\\jack", Tags: []string{"one", "two"}},
			},
		}},
	}

	mc := &testdata.MockClient{}
	mc.Resp = make([]testdata.Response, 0, len(testCases))
	u := &UserId{}
	u.Initialize(mc)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			i := mc.Called
			mc.Resp = append(mc.Resp, testdata.Response{[]byte(""), nil})
			err := u.Run(tc.msg, tc.vsys)
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
