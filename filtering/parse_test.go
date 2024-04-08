package filtering

import (
	"testing"
)

func TestParseEscapeWorks(t *testing.T) {
	g, err := Parse(`foo contains "bar\"baz"`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*Contains)
	if !ok {
		t.Fatalf("g.Matchers[0] type: %T", g.Matchers[0])
	}
	if v.Path != "foo" {
		t.Fatalf("path is %q not \"foo\"", v.Path)
	}
	if v.Value != `bar\"baz` {
		t.Fatalf("value is unexpected: %q", v.Value)
	}
}

func TestParseRegexWithAlternativeQuoteWorks(t *testing.T) {
	g, err := Parse(`foo matches-regex @one"two .*)@`, "@")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*MatchesRegex)
	if !ok {
		t.Fatalf("g.Matchers[0] type: %T", g.Matchers[0])
	}
	if v.Path != "foo" {
		t.Fatalf("path is %q not \"foo\"", v.Path)
	}
	if v.Value != `one"two .*)` {
		t.Fatalf("value is unexpected: %q", v.Value)
	}
}

func TestParseIsNil(t *testing.T) {
	g, err := Parse("foo is-nil", `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*IsNil)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "foo" {
		t.Fatalf("path is %q not \"foo\"", v.Path)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseIsNotNil(t *testing.T) {
	g, err := Parse("bar is-not-nil", `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*IsNotNil)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "bar" {
		t.Fatalf("path is %q not \"bar\"", v.Path)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseEqualsWithDoubleEqualsAndDoubleQuotes(t *testing.T) {
	g, err := Parse(`wutang == "clan"`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*Equals)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "wutang" {
		t.Fatalf("path is %q not \"wutang\"", v.Path)
	}
	if v.Value != "clan" {
		t.Fatalf("value is %q not \"clan\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseEqualsWithEqualsWordAndAlternativeQuotes(t *testing.T) {
	g, err := Parse(`wutang equals ^clan^`, `^`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*Equals)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "wutang" {
		t.Fatalf("path is %q not \"wutang\"", v.Path)
	}
	if v.Value != "clan" {
		t.Fatalf("value is %q not \"clan\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseEqualsWithEqualsWordAndNoQuotes(t *testing.T) {
	g, err := Parse(`bond equals 7`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*Equals)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "bond" {
		t.Fatalf("path is %q not \"bond\"", v.Path)
	}
	if v.Value != "7" {
		t.Fatalf("value is %q not \"7\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseIsNotEqualWithBangEquals(t *testing.T) {
	g, err := Parse(`persona != 3r`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*NotEqualTo)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3r" {
		t.Fatalf("value is %q not \"3r\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseIsNotEqualWithWords(t *testing.T) {
	g, err := Parse(`persona not-equal-to 3r`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*NotEqualTo)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3r" {
		t.Fatalf("value is %q not \"3r\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseLessThanSymbol(t *testing.T) {
	g, err := Parse(`persona < 3`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*LessThan)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3" {
		t.Fatalf("value is %q not \"3\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseLessThanWords(t *testing.T) {
	g, err := Parse(`persona less-than 3`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*LessThan)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3" {
		t.Fatalf("value is %q not \"3\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseLessThanOrEqualToSymbol(t *testing.T) {
	g, err := Parse(`persona <= 3`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*LessThanOrEqualTo)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3" {
		t.Fatalf("value is %q not \"3\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseLessThanOrEqualToWords(t *testing.T) {
	g, err := Parse(`persona less-than-or-equal-to 3`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*LessThanOrEqualTo)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3" {
		t.Fatalf("value is %q not \"3\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseGreaterThanSymbol(t *testing.T) {
	g, err := Parse(`persona > 3`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*GreaterThan)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3" {
		t.Fatalf("value is %q not \"3\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseGreaterThanWords(t *testing.T) {
	g, err := Parse(`persona greater-than 3`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*GreaterThan)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3" {
		t.Fatalf("value is %q not \"3\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseGreaterThanOrEqualToSymbol(t *testing.T) {
	g, err := Parse(`persona >= 3`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*GreaterThanOrEqualTo)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3" {
		t.Fatalf("value is %q not \"3\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseGreaterThanOrEqualToWords(t *testing.T) {
	g, err := Parse(`persona greater-than-or-equal-to 3`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*GreaterThanOrEqualTo)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3" {
		t.Fatalf("value is %q not \"3\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseContains(t *testing.T) {
	g, err := Parse(`persona contains 3r`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*Contains)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3r" {
		t.Fatalf("value is %q not \"3r\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseDoesNotContain(t *testing.T) {
	g, err := Parse(`persona does-not-contain 3r`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*DoesNotContain)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3r" {
		t.Fatalf("value is %q not \"3r\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseStartsWith(t *testing.T) {
	g, err := Parse(`persona starts-with 3r`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*StartsWith)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3r" {
		t.Fatalf("value is %q not \"3r\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseDoesNotStartWith(t *testing.T) {
	g, err := Parse(`persona does-not-start-with 3r`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*DoesNotStartWith)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3r" {
		t.Fatalf("value is %q not \"3r\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseEndsWith(t *testing.T) {
	g, err := Parse(`persona ends-with 3r`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*EndsWith)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3r" {
		t.Fatalf("value is %q not \"3r\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseDoesNotEndWith(t *testing.T) {
	g, err := Parse(`persona does-not-end-with 3r`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*DoesNotEndWith)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "persona" {
		t.Fatalf("path is %q not \"persona\"", v.Path)
	}
	if v.Value != "3r" {
		t.Fatalf("value is %q not \"3r\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseMatchesRegex(t *testing.T) {
	g, err := Parse(`alphabet matches-regex "ab*c"`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*MatchesRegex)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "alphabet" {
		t.Fatalf("path is %q not \"alphabet\"", v.Path)
	}
	if v.Value != "ab*c" {
		t.Fatalf("value is %q not \"ab*c\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseDoesNotMatchRegex(t *testing.T) {
	g, err := Parse(`alphabet does-not-match-regex ab*c`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}
	v, ok := g.Matchers[0].(*DoesNotMatchRegex)
	if !ok {
		t.Fatalf("g.Matcher[0] type: %T", g.Matchers[0])
	}
	if v.Path != "alphabet" {
		t.Fatalf("path is %q not \"alphabet\"", v.Path)
	}
	if v.Value != "ab*c" {
		t.Fatalf("value is %q not \"ab*c\"", v.Value)
	}
	if v.Operator != nil {
		t.Fatalf("operator is not nil")
	}
}

func TestParseAndConditions(t *testing.T) {
	g, err := Parse(`game == shadowrun and slang is chummer && platform is snes`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 3 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}

	v1, ok := g.Matchers[0].(*Equals)
	if !ok {
		t.Fatalf("g.Matchers[0] type: %T", g.Matchers[0])
	}
	if v1.Path != "game" {
		t.Fatalf("path is %q not \"game\"", v1.Path)
	}
	if v1.Value != "shadowrun" {
		t.Fatalf("value is %q not \"shadowrun\"", v1.Value)
	}
	if v1.Operator != nil {
		t.Fatalf("operator is not nil")
	}

	v2, ok := g.Matchers[1].(*Equals)
	if !ok {
		t.Fatalf("g.Matchers[1] type: %T", g.Matchers[1])
	}
	if v2.Path != "slang" {
		t.Fatalf("path is %q not \"slang\"", v2.Path)
	}
	if v2.Value != "chummer" {
		t.Fatalf("value is %q not \"chummer\"", v2.Value)
	}
	if v2.Operator == nil {
		t.Fatalf("operator is not nil")
	}
	if !v2.Operator.And {
		t.Fatalf("expected 'and' operator: %#v", v2.Operator)
	}

	v3, ok := g.Matchers[2].(*Equals)
	if !ok {
		t.Fatalf("g.Matchers[2] type: %T", g.Matchers[2])
	}
	if v3.Path != "platform" {
		t.Fatalf("path is %q not \"platform\"", v3.Path)
	}
	if v3.Value != "snes" {
		t.Fatalf("value is %q not \"snes\"", v3.Value)
	}
	if v3.Operator == nil {
		t.Fatalf("operator is not nil")
	}
	if !v3.Operator.And {
		t.Fatalf("expected 'and' operator: %#v", v3.Operator)
	}
}

func TestParseOrConditions(t *testing.T) {
	g, err := Parse(`game == shadowrun or slang is chummer || platform is snes`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 3 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}

	v1, ok := g.Matchers[0].(*Equals)
	if !ok {
		t.Fatalf("g.Matchers[0] type: %T", g.Matchers[0])
	}
	if v1.Path != "game" {
		t.Fatalf("path is %q not \"game\"", v1.Path)
	}
	if v1.Value != "shadowrun" {
		t.Fatalf("value is %q not \"shadowrun\"", v1.Value)
	}
	if v1.Operator != nil {
		t.Fatalf("operator is not nil")
	}

	v2, ok := g.Matchers[1].(*Equals)
	if !ok {
		t.Fatalf("g.Matchers[1] type: %T", g.Matchers[1])
	}
	if v2.Path != "slang" {
		t.Fatalf("path is %q not \"slang\"", v2.Path)
	}
	if v2.Value != "chummer" {
		t.Fatalf("value is %q not \"chummer\"", v2.Value)
	}
	if v2.Operator == nil {
		t.Fatalf("operator is not nil")
	}
	if !v2.Operator.Or {
		t.Fatalf("expected 'or' operator: %#v", v2.Operator)
	}

	v3, ok := g.Matchers[2].(*Equals)
	if !ok {
		t.Fatalf("g.Matchers[2] type: %T", g.Matchers[2])
	}
	if v3.Path != "platform" {
		t.Fatalf("path is %q not \"platform\"", v3.Path)
	}
	if v3.Value != "snes" {
		t.Fatalf("value is %q not \"snes\"", v3.Value)
	}
	if v3.Operator == nil {
		t.Fatalf("operator is not nil")
	}
	if !v3.Operator.Or {
		t.Fatalf("expected 'or' operator: %#v", v3.Operator)
	}
}

func TestParseXorCondition(t *testing.T) {
	g, err := Parse(`game == shadowrun xor slang is chummer`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 2 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}

	v1, ok := g.Matchers[0].(*Equals)
	if !ok {
		t.Fatalf("g.Matchers[0] type: %T", g.Matchers[0])
	}
	if v1.Path != "game" {
		t.Fatalf("path is %q not \"game\"", v1.Path)
	}
	if v1.Value != "shadowrun" {
		t.Fatalf("value is %q not \"shadowrun\"", v1.Value)
	}
	if v1.Operator != nil {
		t.Fatalf("operator is not nil")
	}

	v2, ok := g.Matchers[1].(*Equals)
	if !ok {
		t.Fatalf("g.Matchers[1] type: %T", g.Matchers[1])
	}
	if v2.Path != "slang" {
		t.Fatalf("path is %q not \"slang\"", v2.Path)
	}
	if v2.Value != "chummer" {
		t.Fatalf("value is %q not \"chummer\"", v2.Value)
	}
	if v2.Operator == nil {
		t.Fatalf("operator is not nil")
	}
	if !v2.Operator.Xor {
		t.Fatalf("expected 'or' operator: %#v", v2.Operator)
	}
}

func TestParseSingleGrouping(t *testing.T) {
	g, err := Parse(`(arcana == "The Fool")`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d | %#v", len(g.Matchers), g)
	}

	g1, ok := g.Matchers[0].(*Group)
	if !ok {
		t.Fatalf("g1.Matchers[0] type: %T", g.Matchers[0])
	}
	if g1.NegateGroup {
		t.Fatalf("g1.Matchers[0] is negated")
	}
	if len(g1.Matchers) != 1 {
		t.Fatalf("len g1 matchers: %d", len(g1.Matchers))
	}

	v1, ok := g1.Matchers[0].(*Equals)
	if !ok {
		t.Fatalf("g1 type: %T", g1.Matchers[0])
	}
	if v1.Path != "arcana" {
		t.Fatalf("v1.Path is %q not \"arcana\"", v1.Path)
	}
	if v1.Value != "The Fool" {
		t.Fatalf("v1.Value is %q not \"The Fool\"", v1.Value)
	}
	if v1.Operator != nil {
		t.Fatalf("v1 has operator: %#v", v1.Operator)
	}
}

func TestParseDoubleNestedGroupNegatedInner(t *testing.T) {
	g, err := Parse(`(!(arcana == "The Fool"))`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d | %#v", len(g.Matchers), g)
	}

	g1, ok := g.Matchers[0].(*Group)
	if !ok {
		t.Fatalf("g.Matchers[0] type: %T", g.Matchers[0])
	}
	if g1.NegateGroup {
		t.Fatalf("g1.Matchers[0] is negated")
	}
	if len(g1.Matchers) != 1 {
		t.Fatalf("len g1 matchers: %d", len(g1.Matchers))
	}

	g2, ok := g1.Matchers[0].(*Group)
	if !ok {
		t.Fatalf("g1.Matchers type: %T", g1.Matchers[0])
	}
	if !g2.NegateGroup {
		t.Fatalf("g2 is not negated")
	}
	if len(g2.Matchers) != 1 {
		t.Fatalf("len g2.Matchers: %d", len(g2.Matchers))
	}

	v1, ok := g2.Matchers[0].(*Equals)
	if !ok {
		t.Fatalf("g1 type: %T", g1.Matchers[0])
	}
	if v1.Path != "arcana" {
		t.Fatalf("v1.Path is %q not \"arcana\"", v1.Path)
	}
	if v1.Value != "The Fool" {
		t.Fatalf("v1.Value is %q not \"The Fool\"", v1.Value)
	}
	if v1.Operator != nil {
		t.Fatalf("v1 has operator: %#v", v1.Operator)
	}
}

func TestParseDoubleNestedGroupNegatedOuter(t *testing.T) {
	g, err := Parse(`!((arcana == "The Fool"))`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(g.Matchers) != 1 {
		t.Fatalf("len matchers: %d | %#v", len(g.Matchers), g)
	}

	g1, ok := g.Matchers[0].(*Group)
	if !ok {
		t.Fatalf("g.Matchers[0] type: %T", g.Matchers[0])
	}
	if !g1.NegateGroup {
		t.Fatalf("g1.Matchers[0] is not negated")
	}
	if len(g1.Matchers) != 1 {
		t.Fatalf("len g1 matchers: %d", len(g1.Matchers))
	}

	g2, ok := g1.Matchers[0].(*Group)
	if !ok {
		t.Fatalf("g1.Matchers type: %T", g1.Matchers[0])
	}
	if g2.NegateGroup {
		t.Fatalf("g2 is negated")
	}
	if len(g2.Matchers) != 1 {
		t.Fatalf("len g2.Matchers: %d", len(g2.Matchers))
	}

	v1, ok := g2.Matchers[0].(*Equals)
	if !ok {
		t.Fatalf("g1 type: %T", g1.Matchers[0])
	}
	if v1.Path != "arcana" {
		t.Fatalf("v1.Path is %q not \"arcana\"", v1.Path)
	}
	if v1.Value != "The Fool" {
		t.Fatalf("v1.Value is %q not \"The Fool\"", v1.Value)
	}
	if v1.Operator != nil {
		t.Fatalf("v1 has operator: %#v", v1.Operator)
	}
}

func TestParseFirstOfThreeIsGroup(t *testing.T) {
	g, err := Parse(`(foo is-nil) and bar is-nil and baz is-nil`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if len(g.Matchers) != 3 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}

	if _, ok := g.Matchers[0].(*Group); !ok {
		t.Fatalf("first is not group: %T", g.Matchers[0])
	}

	if _, ok := g.Matchers[1].(*IsNil); !ok {
		t.Fatalf("second not is-nil: %T", g.Matchers[1])
	}

	if _, ok := g.Matchers[2].(*IsNil); !ok {
		t.Fatalf("third not is-nil: %T", g.Matchers[2])
	}
}

func TestParseSecondOfThreeIsGroup(t *testing.T) {
	g, err := Parse(`foo is-nil and (bar is-nil) and baz is-nil`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if len(g.Matchers) != 3 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}

	if _, ok := g.Matchers[0].(*IsNil); !ok {
		t.Fatalf("first not is-nil: %T", g.Matchers[1])
	}

	if _, ok := g.Matchers[1].(*Group); !ok {
		t.Fatalf("second is not group: %T", g.Matchers[0])
	}

	if _, ok := g.Matchers[2].(*IsNil); !ok {
		t.Fatalf("third not is-nil: %T", g.Matchers[2])
	}
}

func TestParseThirdOfThreeIsGroup(t *testing.T) {
	g, err := Parse(`foo is-nil and bar is-nil and !(baz is-nil)`, `"`)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if len(g.Matchers) != 3 {
		t.Fatalf("len matchers: %d", len(g.Matchers))
	}

	if _, ok := g.Matchers[0].(*IsNil); !ok {
		t.Fatalf("first not is-nil: %T", g.Matchers[1])
	}

	if _, ok := g.Matchers[1].(*IsNil); !ok {
		t.Fatalf("second not is-nil: %T", g.Matchers[0])
	}

	if _, ok := g.Matchers[2].(*Group); !ok {
		t.Fatalf("third not group: %T", g.Matchers[2])
	}
}

func TestParseReturnsError(t *testing.T) {
	checks := []string{
		"(foo is-nil",
		"bar contains \"jack\")",
		"unknown operator here",
		"(((foo is-nil) is-nil))",
		`quote contains "jack`,
		"this is-not-nil bar",
		"foo contains o and",
		"foobar",
		"(",
		")",
	}

	for _, s := range checks {
		if _, err := Parse(s, `"`); err == nil {
			t.Fatalf("invalid filter string was ok: %q", s)
		}
	}
}

func TestParseWithInvalidQuoteReturnsError(t *testing.T) {
	quotes := []string{
		"|",
		"&",
		">",
		"<",
		"=",
		"(",
		")",
		" ",
		`\`,
		"!",
		".",
		"-",
		"_",
	}

	okstr := "foo > 5"
	for _, quote := range quotes {
		if _, err := Parse(okstr, quote); err == nil {
			t.Fatalf("invalid quote was ok: %q", quote)
		}
	}
}

func TestParseEmptyString(t *testing.T) {
	resp, err := Parse("", `"`)
	if err != nil {
		t.Fatalf("parse empty string error: %s", err)
	} else if resp != nil {
		t.Fatalf("parse empty string has non-nil group")
	}
}
