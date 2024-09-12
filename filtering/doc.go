/*
Package filtering implements a way to do client-side filtering on a list of
normalized objects.

To use, start with the `Parse()` function:

	logic, err := filtering.Parse(`name contains "foo" and description is-nil`, `"`)

Once the string has been parsed, the logic can be applied to any normalized object that
implements the `Fielder` interface:

	ok, err := logic.Matches(obj)

# Groupings

Groupings via parenthesis is supported. A grouping can be negated by giving a grouping
a bang prefix:

	// These two are equivalent.
	_, _ = filtering.Parse(`description is-nil`, `"`)
	_, _ = filtering.Parse(`!(description is-not-nil)`, `"`)

# Operators

The following operators are supported:

  - and: `&&` or `and`
  - or: `||` or `or`
  - xor: `xor`

# Comparisons

The following checks are supported:

  - `is-nil`: True when a field is a pointer and that pointer is undefined. This check
    does not expect a value to be specified.
  - `is-not-nil`: True when a field is a pointer and that pointer is defined. This check
    does not expect a value to be specified.
  - `==` | `equals` | `is`: True when a field matches the given value. A field that is a
    pointer and that pointer is undefined will return False.
  - `!=` | `not-equal-to` | `is-not`: True when a field does not match the given value. A field
    that is a nil pointer will return True.
  - `<` | `less-than`: (int64 / float64) True when the field is less than the value. A
    field that is a nil pointer will return False.
  - `<=` | `less-than-or-equal-to`: (int64 / float64) True when the field is less than or equal
    to the value.  A field that is a nil pointer will return False.
  - `>` | `greater-than`: (int64 / float64) True when the field is greater than the value. A
    field that is a nil pointer will return False.
  - `>=` | `greater-than-or-equal-to`: (int64 / float64) True when the field is greater than
    or equal to the value.  A field that is a nil pointer will return False.
  - `contains`: For strings, this is treated as a substring match, and will be False if the
    field is nil.  For a list of strings, this is treated as an exact across the string slice.
  - `does-not-contain`: For strings, this returns True if the substring is not present in the
    field, and will be True if the field is a nil pointer.  For a list of strings, this is treated
    as an exact match across all strings in the slice, and will return True if none of the strings
    exactly match the given value.
  - `starts-with`: (string only) Returns True if the field has the given prefix. A nil pointer
    to a string returns False.
  - `does-not-start-with`: (string only) Returns True if the field does not have the given
    prefix. A nil pointer to a string returns True.
  - `ends-with`: (string only) Returns True if the field has the given suffix. A nil pointer
    to a string returns False.
  - `does-not-end-with`: (string only) Returns True if the field does not have the given suffix. A
    nil pointer to a string returns True.
  - `matches-regex`: (string / []string) For strings, searches the string for the given
    regex; a field that is a pointer to a nil string returns False.  For string slices, the
    regex is checked against all strings in the slice, and returns True if any of them
    match the regex.
  - `does-not-match-regex`: (string / slice of strings) For a string, returns True if the field
    does not match the given regex; a field that is a pointer to a nil string returns True. For
    string slices, the regex is checked against all strings in the slice, and returns True if
    none of them match the regex.

# Field Selection

To select a field, use dotted notation as appropriate:

	// Returns "Name" from the base object.
	"name"

	// Returns "Bar" in the "Foo" object.
	"Foo.Bar"

Names can be specified using the camel case name (for users using pango directly) or the
camel case name (for users using the panos Terraform provider).
*/
package filtering
