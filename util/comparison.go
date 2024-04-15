package util

func UnorderedListsMatch(a, b []string) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else if len(a) != len(b) {
		return false
	}

	for _, x := range a {
		var found bool
		for _, y := range b {
			if x == y {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func OrderedListsMatch(a, b []string) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TargetsMatch(a, b map[string][]string) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else if len(a) != len(b) {
		return false
	}

	for key := range a {
		if !UnorderedListsMatch(a[key], b[key]) {
			return false
		}
	}

	return true
}

func StringsMatch(a, b *string) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}

	return *a == *b
}

func BoolsMatch(a, b *bool) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}

	return *a == *b
}

func FloatsMatch(a, b *float64) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}

	return *a == *b
}

func IntsMatch(a, b *int64) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	}

	return *a == *b
}

func AnysMatch(a, b any) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return true
}
