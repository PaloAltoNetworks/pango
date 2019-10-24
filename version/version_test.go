package version

import (
    "fmt"
    "testing"
)

func TestNew(t *testing.T) {
    testCases := []struct{
        s string
        r bool
        a, b, c int
        d string
    }{
        {"1.2.3", false, 1, 2, 3, ""},
        {"1.2.3-h4", false, 1, 2, 3, "h4"},
        {"12.34.56-78h", false, 12, 34, 56, "78h"},
        {"a.2.3", true, 0, 0, 0, ""},
        {"1.b.3", true, 0, 0, 0, ""},
        {"1.2.c", true, 0, 0, 0, ""},
        {"1.2.3h4", true, 0, 0, 0, ""},
        {"9.0.3.xfr", false, 9, 0, 3, ""},
    }

    for _, tc := range testCases {
        t.Run(fmt.Sprintf("%s should error %t", tc.s, tc.r), func(t *testing.T) {
            v, err := New(tc.s)
            if (err != nil) != tc.r || v.Major != tc.a || v.Minor != tc.b || v.Patch != tc.c || v.Suffix != tc.d {
                t.Fail()
            }
        })
    }
}

func TestStringer(t *testing.T) {
    testCases := []struct{
        a, b, c int
        d, want string
    }{
        {1, 2, 3, "", "1.2.3"},
        {1, 2, 3, "h4", "1.2.3-h4"},
        {12, 34, 56, "h78", "12.34.56-h78"},
    }

    for _, tc := range testCases {
        t.Run(tc.want, func(t *testing.T) {
            v := Number{tc.a, tc.b, tc.c, tc.d}
            if v.String() != tc.want {
                t.Fail()
            }
        })
    }
}

func TestGte(t *testing.T) {
    testCases := []struct{
        a, b, c int
        r bool
    }{
        {1, 1, 1, true},
        {1, 1, 2, true},
        {1, 1, 3, true},
        {1, 2, 1, true},
        {1, 2, 2, true},
        {1, 2, 3, true},
        {1, 3, 1, true},
        {1, 3, 2, true},
        {1, 3, 3, true},
        {2, 1, 1, true},
        {2, 1, 2, true},
        {2, 1, 3, true},
        {2, 2, 1, true},
        {2, 2, 2, true},
        {2, 2, 3, false},
        {2, 3, 1, false},
        {2, 3, 2, false},
        {2, 3, 3, false},
        {3, 1, 1, false},
        {3, 1, 2, false},
        {3, 1, 3, false},
        {3, 2, 1, false},
        {3, 2, 2, false},
        {3, 2, 3, false},
        {3, 3, 1, false},
        {3, 3, 2, false},
        {3, 3, 3, false},
    }
    v1 := Number{2, 2, 2, ""}

    for _, tc := range testCases {
        t.Run(fmt.Sprintf("%s >= %d.%d.%d == %t", v1, tc.a, tc.b, tc.c, tc.r), func(t *testing.T) {
            r := v1.Gte(Number{tc.a, tc.b, tc.c, ""})
            if r != tc.r {
                t.Fail()
            }
        })
    }
}

func BenchmarkGteMajor(b *testing.B) {
    v1 := Number{5, 5, 5, ""}
    v2 := Number{6, 5, 5, ""}

    for i := 0; i < b.N; i++ {
        _ = v1.Gte(v2)
    }
}

func BenchmarkGtePatch(b *testing.B) {
    v1 := Number{5, 5, 5, ""}
    v2 := Number{5, 5, 6, ""}

    for i := 0; i < b.N; i++ {
        _ = v1.Gte(v2)
    }
}

func BenchmarkNew(b *testing.B) {
    s := "7.1.12-h4"

    for i := 0; i < b.N; i++ {
        _, _ = New(s)
    }
}
