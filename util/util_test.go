package util

import (
    "fmt"
    "testing"
)


func TestMemToStrNil(t *testing.T) {
    r := MemToStr(nil)
    if r != nil {
        t.Fail()
    }
}

func TestEntToStrNil(t *testing.T) {
    r := EntToStr(nil)
    if r != nil {
        t.Fail()
    }
}

func TestStrToMem(t *testing.T) {
    v := []string{"one", "two"}
    r := StrToMem(v)
    if r == nil {
        t.Fail()
    } else if len(v) != len(r.Members) {
        t.Fail()
    } else {
        for i := range v {
            if v[i] != r.Members[i].Value {
                t.Fail()
                break
            }
        }
    }
}

func TestStrToEnt(t *testing.T) {
    v := []string{"one", "two"}
    r := StrToEnt(v)
    if r == nil {
        t.Fail()
    } else if len(v) != len(r.Entries) {
        t.Fail()
    } else {
        for i := range v {
            if v[i] != r.Entries[i].Value {
                t.Fail()
                break
            }
        }
    }
}

func BenchmarkStrToMem(b *testing.B) {
    v := []string{"one", "two", "three", "four", "five"}
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        _ = StrToMem(v)
    }
}

func BenchmarkMemToStr(b *testing.B) {
    m := &MemberType{[]Member{
        {Value: "one"},
        {Value: "two"},
        {Value: "three"},
        {Value: "four"},
        {Value: "five"},
    }}
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        _ = MemToStr(m)
    }
}

func BenchmarkStrToEnt(b *testing.B) {
    v := []string{"one", "two", "three", "four", "five"}
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        _ = StrToEnt(v)
    }
}

func BenchmarkEntToStr(b *testing.B) {
    v := &EntryType{[]Entry{
        {Value: "one"},
        {Value: "two"},
        {Value: "three"},
        {Value: "four"},
        {Value: "five"},
    }}
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        _ = EntToStr(v)
    }
}

func BenchmarkAsXpath(b *testing.B) {
    p := []string{
        "config",
        "devices",
        AsEntryXpath([]string{"localhost.localdomain"}),
        "vsys",
        AsEntryXpath([]string{"vsys1"}),
        "import",
        "network",
    }
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        _ = AsXpath(p)
    }
}

func BenchmarkAsEntryXpathMultiple(b *testing.B) {
    v := []string{"one", "two", "three"}
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        _ = AsEntryXpath(v)
    }
}

func TestAsXpath(t *testing.T) {
    testCases := []struct{
        i interface{}
        r string
    }{
        {"/one/two", "/one/two"},
        {[]string{"one", "two"}, "/one/two"},
        {42, ""},
    }

    for _, tc := range testCases {
        t.Run(fmt.Sprintf("%v to %s", tc.i, tc.r), func(t *testing.T) {
            if AsXpath(tc.i) != tc.r {
                t.Fail()
            }
        })
    }
}

func TestAsEntryXpath(t *testing.T) {
    testCases := []struct{
        v []string
        r string
    }{
        {[]string{"one"}, "entry[@name='one']"},
        {[]string{"one", "two"}, "entry[@name='one' or @name='two']"},
        {nil, "entry"},
    }

    for _, tc := range testCases {
        t.Run(tc.r, func(t *testing.T) {
            if AsEntryXpath(tc.v) != tc.r {
                t.Fail()
            }
        })
    }
}

func TestAsMemberXpath(t *testing.T) {
    testCases := []struct{
        v []string
        r string
    }{
        {[]string{"one"}, "member[text()='one']"},
        {[]string{"one", "two"}, "member[text()='one' or text()='two']"},
        {nil, "member[]"},
    }

    for _, tc := range testCases {
        t.Run(tc.r, func(t *testing.T) {
            if AsMemberXpath(tc.v) != tc.r {
                t.Fail()
            }
        })
    }
}

func TestCleanRawXml(t *testing.T) {
    v := `<foo admin="admin" dirtyId="2" time="1234/05/06 07:08:09">hi</foo>`
    if CleanRawXml(v) != "<foo>hi</foo>" {
        t.Fail()
    }
}
