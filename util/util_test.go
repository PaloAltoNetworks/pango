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
		AsEntryXpath("localhost.localdomain"),
		"vsys",
		AsEntryXpath("vsys1"),
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
		_ = AsEntryXpath(v...)
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
		// CWE-643: names containing single quotes must not be able to
		// break out of the predicate; they are emitted via concat().
		{[]string{"x' or '1'='1"}, `entry[@name=concat('x',"'",' or ',"'",'1',"'",'=',"'",'1')]`},
		{[]string{"a'b"}, `entry[@name=concat('a',"'",'b')]`},
		{[]string{"one", "x'y"}, `entry[@name='one' or @name=concat('x',"'",'y')]`},
	}

	for _, tc := range testCases {
		t.Run(tc.r, func(t *testing.T) {
			if AsEntryXpath(tc.v...) != tc.r {
				t.Errorf("AsEntryXpath(%q) = %q, want %q", tc.v, AsEntryXpath(tc.v...), tc.r)
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
		// CWE-643: member values with single quotes are emitted via concat().
		{[]string{"a'b"}, `member[text()=concat('a',"'",'b')]`},
	}

	for _, tc := range testCases {
		t.Run(tc.r, func(t *testing.T) {
			if AsMemberXpath(tc.v) != tc.r {
				t.Errorf("AsMemberXpath(%q) = %q, want %q", tc.v, AsMemberXpath(tc.v), tc.r)
			}
		})
	}
}

func TestAsUuidXpath(t *testing.T) {
	testCases := []struct {
		v string
		r string
	}{
		{"one", "entry[@uuid='one']"},
		// CWE-643: uuid values with single quotes are emitted via concat().
		{"a'b", `entry[@uuid=concat('a',"'",'b')]`},
	}

	for _, tc := range testCases {
		t.Run(tc.r, func(t *testing.T) {
			if AsUuidXpath(tc.v) != tc.r {
				t.Errorf("AsUuidXpath(%q) = %q, want %q", tc.v, AsUuidXpath(tc.v), tc.r)
			}
		})
	}
}

func TestXpathSafe(t *testing.T) {
	testCases := []struct {
		v string
		r string
	}{
		{"", "''"},
		{"plain", "'plain'"},
		{"a'b", `concat('a',"'",'b')`},
		{"'", `concat('',"'",'')`},
		{"a'b'c", `concat('a',"'",'b',"'",'c')`},
	}

	for _, tc := range testCases {
		t.Run(tc.v, func(t *testing.T) {
			if xpathSafe(tc.v) != tc.r {
				t.Errorf("xpathSafe(%q) = %q, want %q", tc.v, xpathSafe(tc.v), tc.r)
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

func TestEntryName(t *testing.T) {
	testCases := []struct {
		component string
		want      string
	}{
		{"entry[@name='one']", "one"},
		{"entry[@name='']", ""},
		{`entry[@name=concat('a',"'",'b')]`, "a'b"},
		{`entry[@name=concat('',"'",'')]`, "'"},
		{`entry[@name=concat('a',"'",'b',"'",'c')]`, "a'b'c"},
		{`entry[@name=concat('a,b',"'",'c')]`, "a,b'c"},
		// Not an entry-name predicate: returned unchanged.
		{"entry", "entry"},
	}

	for _, tc := range testCases {
		t.Run(tc.component, func(t *testing.T) {
			if EntryName(tc.component) != tc.want {
				t.Errorf("EntryName(%q) = %q, want %q", tc.component, EntryName(tc.component), tc.want)
			}
		})
	}
}

// TestEntryNameRoundTrip proves EntryName is the exact inverse of AsEntryXpath
// for arbitrary names, including those that trigger the concat() form.
func TestEntryNameRoundTrip(t *testing.T) {
	// Note: "" is excluded because AsEntryXpath("") is the bare "entry" listing
	// segment, not a name predicate.
	names := []string{
		"plain",
		"a'b",
		"'",
		"O'Brien",
		`x' or '1'='1`,
		"a,b'c",
		`has"double`,
	}

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			got := EntryName(AsEntryXpath(name))
			if got != name {
				t.Errorf("EntryName(AsEntryXpath(%q)) = %q, want %q", name, got, name)
			}
		})
	}
}
