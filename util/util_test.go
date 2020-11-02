package util

import (
	"bytes"
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
	testCases := []struct {
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
	testCases := []struct {
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
	testCases := []struct {
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

func TestStripPanosPackagingNoTag(t *testing.T) {
	expected := "<outer><inner/></outer>"
	input := fmt.Sprintf("<response><result>%s</result></response>", expected)

	ans := StripPanosPackaging([]byte(input), "")
	if !bytes.Equal([]byte(expected), ans) {
		t.Errorf("Expected %q, got %q", expected, ans)
	}
}

func TestStripPanosPackagingWithTag(t *testing.T) {
	expected := "<inner/>"
	input := fmt.Sprintf("<response><result><outer>%s</outer></result></response>", expected)

	ans := StripPanosPackaging([]byte(input), "outer")
	if !bytes.Equal([]byte(expected), ans) {
		t.Errorf("Expected %q, got %q", expected, ans)
	}
}

func TestStripPanosPackagingNoResult(t *testing.T) {
	input := `<response status="success" code="19"><result total-count="1" count="1">
  <interface admin="admin" dirtyId="52" time="2020/10/28 11:55:24"/>
</result></response>`

	ans := StripPanosPackaging([]byte(input), "interface")
	if len(ans) != 0 {
		t.Errorf("Expected empty string, got %q", ans)
	}
}
