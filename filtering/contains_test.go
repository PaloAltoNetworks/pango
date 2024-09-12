package filtering

import (
	"testing"
)

func TestContainsStringPointer(t *testing.T) {
	obj := &Contains{
		Path:  "test",
		Value: "bar",
	}

	str := "foo bar baz"
	f := &testFielder{
		Answers: []testFielderResponse{
			{Value: &str},
		},
	}

	ans, err := obj.Matches(f)
	if err != nil {
		t.Fatalf("err from matches: %s", err)
	}
	if !ans {
		t.Fatalf("fail: %q contains %q", str, obj.Value)
	}
}

func TestContainsEmptyStringPointer(t *testing.T) {
	obj := &Contains{
		Path:  "test",
		Value: "bar",
	}

	str := ""
	f := &testFielder{
		Answers: []testFielderResponse{
			{Value: &str},
		},
	}

	ans, err := obj.Matches(f)
	if err != nil {
		t.Fatalf("err from matches: %s", err)
	}
	if ans {
		t.Fatalf("fail: string contains %q", obj.Value)
	}
}

func TestContainsString(t *testing.T) {
	obj := &Contains{
		Path:  "test",
		Value: " ",
	}

	f := &testFielder{
		Answers: []testFielderResponse{
			{Value: "hello world"},
		},
	}

	ans, err := obj.Matches(f)
	if err != nil {
		t.Fatalf("err in matches: %s", err)
	}
	if !ans {
		t.Fatalf("fail: no space found")
	}
}
