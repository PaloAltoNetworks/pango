package util

// For reference: https://github.com/golang/go/issues/45624

// Bool returns a pointer to the given bool.
func Bool(v bool) *bool {
	return &v
}

// Int returns a pointer to the given int64.
func Int(v int64) *int64 {
	return &v
}

// Float returns a pointer to the given float64.
func Float(v float64) *float64 {
	return &v
}

// String returns a pointer to the given string.
func String(v string) *string {
	return &v
}
