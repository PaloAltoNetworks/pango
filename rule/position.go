package rule

import (
	"fmt"

	"github.com/PaloAltoNetworks/pango/errors"
)

type Position struct {
	First           *bool   `json:"first,omitempty"`
	Last            *bool   `json:"last,omitempty"`
	SomewhereBefore *string `json:"somewhere_before,omitempty"`
	DirectlyBefore  *string `json:"directly_before,omitempty"`
	SomewhereAfter  *string `json:"somewhere_after,omitempty"`
	DirectlyAfter   *string `json:"directly_after,omitempty"`
}

func (o *Position) IsValid(removeEverythingElse bool) error {
	count := 0

	if o.First != nil && *o.First {
		count++
	}

	if o.Last != nil && *o.Last {
		count++
	}

	if o.SomewhereBefore != nil {
		if removeEverythingElse {
			return errors.RelativePositionWithRemoveEverythingElseError
		}
		if *o.SomewhereBefore != "" {
			count++
		}
	}

	if o.DirectlyBefore != nil {
		if removeEverythingElse {
			return errors.RelativePositionWithRemoveEverythingElseError
		}
		if *o.DirectlyBefore != "" {
			count++
		}
	}

	if o.SomewhereAfter != nil {
		if removeEverythingElse {
			return errors.RelativePositionWithRemoveEverythingElseError
		}
		if *o.SomewhereAfter != "" {
			count++
		}
	}

	if o.DirectlyAfter != nil {
		if removeEverythingElse {
			return errors.RelativePositionWithRemoveEverythingElseError
		}
		if *o.DirectlyAfter != "" {
			count++
		}
	}

	if count > 1 {
		return fmt.Errorf("multiple positions specified: only one should be specified")
	}

	return nil
}
