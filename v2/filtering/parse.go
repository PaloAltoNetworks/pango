package filtering

import (
	"fmt"
	"io"
	"strings"

	"github.com/PaloAltoNetworks/pango/v2/errors"
)

func Parse(s string, quote string) (*Group, error) {
	if s == "" {
		return nil, nil
	} else if len(quote) != 1 {
		return nil, fmt.Errorf("quote character should be one character")
	} else if quote == "&" || quote == "|" || quote == "(" || quote == ")" || quote == " " || quote == `\` || quote == "!" || quote == "." || quote == "<" || quote == ">" || quote == "=" || quote == "-" || quote == "_" {
		return nil, fmt.Errorf("quote character cannot be a reserved character")
	}

	var ch rune
	var err error
	var op *Operator
	var path string
	var token strings.Builder
	token.Grow(50)

	var quoteRune rune
	for _, quoteRune = range quote {
		break
	}

	groups := make([]*Group, 0, 10)
	groups = append(groups, &Group{
		Matchers: make([]Matcher, 0, 10),
	})

	// So, we are going to cheat a bit here and add a trailing space to the end of the
	// filter passed in.  This means that we shouldn't have to add any trailing handling
	// outside of the main "for" loop.
	r := strings.NewReader(s + " ")

	for true {
		token.Reset()

		if err = skipSpaces(r); err != nil {
			break
		}

		ch, _, err = r.ReadRune()
		if err != nil {
			break
		}

		// New group: regular.
		if ch == '(' {
			group := &Group{
				Matchers: make([]Matcher, 0, 10),
				Operator: op,
			}
			cur := groups[len(groups)-1]
			cur.Matchers = append(cur.Matchers, group)
			groups = append(groups, group)

			op = nil
			continue
		}

		// Negation can be either negating a group or "!=".
		if ch == '!' {
			ch, _, err = r.ReadRune()
			if err != nil {
				if err == io.EOF {
					return nil, errors.InvalidFilterError
				}
				return nil, err
			}
			if ch == '(' {
				group := &Group{
					NegateGroup: true,
					Matchers:    make([]Matcher, 0, 10),
					Operator:    op,
				}
				cur := groups[len(groups)-1]
				cur.Matchers = append(cur.Matchers, group)
				groups = append(groups, group)
				op = nil
				continue
			}

			if err = r.UnreadRune(); err != nil {
				return nil, err
			}
			ch = '!'
		}

		// End group.
		if ch == ')' {
			if len(groups) == 1 {
				return nil, errors.InvalidFilterError
			}
			groups = groups[:len(groups)-1]
			continue
		}

		// Get the next token.
		token.Reset()
		if ch == quoteRune {
			prevIsEscape := false
			for true {
				ch, _, err = r.ReadRune()
				if err != nil {
					if err == io.EOF {
						return nil, errors.InvalidFilterError
					}
					return nil, err
				}
				if ch == quoteRune && !prevIsEscape {
					break
				}
				token.WriteRune(ch)
				prevIsEscape = ch == '\\'
			}
		} else {
			token.WriteRune(ch)
			for true {
				ch, _, err = r.ReadRune()
				if err != nil {
					if err == io.EOF {
						return nil, errors.InvalidFilterError
					}
					return nil, err
				}
				if ch == ' ' || ch == ')' {
					if err = r.UnreadRune(); err != nil {
						return nil, err
					}
					break
				}
				token.WriteRune(ch)
			}
		}

		// At this point, we have a token that needs to be interpretted.
		if len(groups[len(groups)-1].Matchers) != 0 && op == nil {
			switch token.String() {
			case "&&", "and":
				op = &Operator{And: true}
			case "||", "or":
				op = &Operator{Or: true}
			case "xor":
				op = &Operator{Xor: true}
			default:
				return nil, errors.InvalidFilterError
			}
		} else if path == "" {
			path = token.String()
		} else {
			group := groups[len(groups)-1]
			switch token.String() {
			case "is-nil":
				group.Matchers = append(group.Matchers, &IsNil{
					Path:     path,
					Operator: op,
				})
			case "is-not-nil":
				group.Matchers = append(group.Matchers, &IsNotNil{
					Path:     path,
					Operator: op,
				})
			case "==", "equals", "is":
				value, err := getValue(r, quoteRune)
				if err != nil {
					return nil, err
				}
				group.Matchers = append(group.Matchers, &Equals{
					Path:     path,
					Value:    value,
					Operator: op,
				})
			case "!=", "not-equal-to", "is-not":
				value, err := getValue(r, quoteRune)
				if err != nil {
					return nil, err
				}
				group.Matchers = append(group.Matchers, &NotEqualTo{
					Path:     path,
					Value:    value,
					Operator: op,
				})
			case "<", "less-than":
				value, err := getValue(r, quoteRune)
				if err != nil {
					return nil, err
				}
				group.Matchers = append(group.Matchers, &LessThan{
					Path:     path,
					Value:    value,
					Operator: op,
				})
			case "<=", "less-than-or-equal-to":
				value, err := getValue(r, quoteRune)
				if err != nil {
					return nil, err
				}
				group.Matchers = append(group.Matchers, &LessThanOrEqualTo{
					Path:     path,
					Value:    value,
					Operator: op,
				})
			case ">", "greater-than":
				value, err := getValue(r, quoteRune)
				if err != nil {
					return nil, err
				}
				group.Matchers = append(group.Matchers, &GreaterThan{
					Path:     path,
					Value:    value,
					Operator: op,
				})
			case ">=", "greater-than-or-equal-to":
				value, err := getValue(r, quoteRune)
				if err != nil {
					return nil, err
				}
				group.Matchers = append(group.Matchers, &GreaterThanOrEqualTo{
					Path:     path,
					Value:    value,
					Operator: op,
				})
			case "contains":
				value, err := getValue(r, quoteRune)
				if err != nil {
					return nil, err
				}
				group.Matchers = append(group.Matchers, &Contains{
					Path:     path,
					Value:    value,
					Operator: op,
				})
			case "does-not-contain":
				value, err := getValue(r, quoteRune)
				if err != nil {
					return nil, err
				}
				group.Matchers = append(group.Matchers, &DoesNotContain{
					Path:     path,
					Value:    value,
					Operator: op,
				})
			case "starts-with":
				value, err := getValue(r, quoteRune)
				if err != nil {
					return nil, err
				}
				group.Matchers = append(group.Matchers, &StartsWith{
					Path:     path,
					Value:    value,
					Operator: op,
				})
			case "does-not-start-with":
				value, err := getValue(r, quoteRune)
				if err != nil {
					return nil, err
				}
				group.Matchers = append(group.Matchers, &DoesNotStartWith{
					Path:     path,
					Value:    value,
					Operator: op,
				})
			case "ends-with":
				value, err := getValue(r, quoteRune)
				if err != nil {
					return nil, err
				}
				group.Matchers = append(group.Matchers, &EndsWith{
					Path:     path,
					Value:    value,
					Operator: op,
				})
			case "does-not-end-with":
				value, err := getValue(r, quoteRune)
				if err != nil {
					return nil, err
				}
				group.Matchers = append(group.Matchers, &DoesNotEndWith{
					Path:     path,
					Value:    value,
					Operator: op,
				})
			case "matches-regex":
				value, err := getValue(r, quoteRune)
				if err != nil {
					return nil, err
				}
				group.Matchers = append(group.Matchers, &MatchesRegex{
					Path:     path,
					Value:    value,
					Operator: op,
				})
			case "does-not-match-regex":
				value, err := getValue(r, quoteRune)
				if err != nil {
					return nil, err
				}
				group.Matchers = append(group.Matchers, &DoesNotMatchRegex{
					Path:     path,
					Value:    value,
					Operator: op,
				})
			default:
				return nil, fmt.Errorf("unknown operator: %q", token.String())
			}

			op = nil
			path = ""
		}
	}

	if err != nil && err != io.EOF {
		return nil, err
	}

	if len(groups) != 1 || op != nil || path != "" {
		return nil, errors.InvalidFilterError
	}

	return groups[0], nil
}

func getValue(r *strings.Reader, quoteRune rune) (string, error) {
	var ch rune
	var err error
	var token strings.Builder
	token.Grow(30)

	if err = skipSpaces(r); err != nil {
		return "", err
	}

	ch, _, err = r.ReadRune()
	if err != nil {
		return "", err
	}

	if ch == quoteRune {
		prevIsEscape := false
		for true {
			ch, _, err = r.ReadRune()
			if err != nil {
				return token.String(), err
			}
			if ch == quoteRune && !prevIsEscape {
				break
			}
			token.WriteRune(ch)
			prevIsEscape = ch == '\\'
		}
	} else {
		token.WriteRune(ch)
		for true {
			ch, _, err = r.ReadRune()
			if err != nil {
				return token.String(), err
			}
			if ch == ' ' {
				if err = r.UnreadRune(); err != nil {
					return token.String(), err
				}
				break
			}
			token.WriteRune(ch)
		}
	}

	return token.String(), nil
}

func skipSpaces(r *strings.Reader) error {
	var ch rune
	var err error

	for true {
		ch, _, err = r.ReadRune()
		if err != nil {
			return err
		}
		if ch != ' ' && ch != '\t' {
			err = r.UnreadRune()
			break
		}
	}

	return err
}
