package movement

import (
	"errors"
	"fmt"
	"log/slog"
	"slices"
	"strings"
)

var _ = slog.LevelDebug

type ActionWhereType string

const (
	ActionWhereFirst  ActionWhereType = "top"
	ActionWhereLast   ActionWhereType = "bottom"
	ActionWhereBefore ActionWhereType = "before"
	ActionWhereAfter  ActionWhereType = "after"
)

type Movable interface {
	EntryName() string
}

type MoveAction struct {
	Movable     Movable
	Where       ActionWhereType
	Destination Movable
}

type Position interface {
	Move(entries []Movable, existing []Movable) ([]MoveAction, error)
	GetExpected(entries []Movable, existing []Movable) ([]Movable, error)
	IsDirectly() bool
	Where() ActionWhereType
	PivotEntryName() string
}

type PositionFirst struct{}

func (o PositionFirst) IsDirectly() bool {
	return false
}

func (o PositionFirst) Where() ActionWhereType {
	return ActionWhereFirst
}

func (o PositionFirst) PivotEntryName() string {
	return ""
}

type PositionLast struct{}

func (o PositionLast) IsDirectly() bool {
	return false
}

func (o PositionLast) Where() ActionWhereType {
	return ActionWhereLast
}

func (o PositionLast) PivotEntryName() string {
	return ""
}

type PositionBefore struct {
	Directly bool
	Pivot    string
}

func (o PositionBefore) IsDirectly() bool {
	return o.Directly
}

func (o PositionBefore) Where() ActionWhereType {
	return ActionWhereBefore
}

func (o PositionBefore) PivotEntryName() string {
	return o.Pivot
}

type PositionAfter struct {
	Directly bool
	Pivot    string
}

func (o PositionAfter) IsDirectly() bool {
	return o.Directly
}

func (o PositionAfter) Where() ActionWhereType {
	return ActionWhereAfter
}

func (o PositionAfter) PivotEntryName() string {
	return o.Pivot
}

type entryWithIdx[E Movable] struct {
	Entry E
	Idx   int
}

func entriesByName[E Movable](entries []E) map[string]entryWithIdx[E] {
	entriesIdxMap := make(map[string]entryWithIdx[E], len(entries))
	for idx, elt := range entries {
		entriesIdxMap[elt.EntryName()] = entryWithIdx[E]{
			Entry: elt,
			Idx:   idx,
		}
	}
	return entriesIdxMap
}

func removeEntriesFromExisting(entries []Movable, filterFn func(entry Movable) bool) []Movable {
	entryNames := make(map[string]bool, len(entries))
	for _, elt := range entries {
		entryNames[elt.EntryName()] = true
	}

	filtered := make([]Movable, len(entries))
	copy(filtered, entries)

	filtered = slices.DeleteFunc(filtered, filterFn)

	return filtered
}

func findPivotIdx(entries []Movable, pivot string) (int, Movable) {
	var pivotEntry Movable
	pivotIdx := slices.IndexFunc(entries, func(entry Movable) bool {
		if entry.EntryName() == pivot {
			pivotEntry = entry
			return true
		}

		return false
	})

	return pivotIdx, pivotEntry
}

var (
	errNoMovements          = errors.New("no movements needed")
	ErrSlicesNotEqualLength = errors.New("existing and expected slices length mismatch")
	ErrPivotInEntries       = errors.New("pivot element found in the entries slice")
	ErrPivotNotInExisting   = errors.New("pivot element not foudn in the existing slice")
	ErrInvalidMovementPlan  = errors.New("created movement plan is invalid")
)

// PositionBefore and PositionAfter are similar enough that we can generate expected sequences
// for both using the same code and some conditionals based on the given movement.
func getPivotMovement(entries []Movable, existing []Movable, pivot string, direct bool, movement ActionWhereType) ([]Movable, error) {
	existingIdxMap := entriesByName(existing)

	entriesPivotIdx, _ := findPivotIdx(entries, pivot)
	if entriesPivotIdx != -1 {
		return nil, ErrPivotInEntries
	}

	existingPivotIdx, _ := findPivotIdx(existing, pivot)
	if existingPivotIdx == -1 {
		return nil, ErrPivotNotInExisting
	}

	if !direct {
		movementRequired := false
		entriesLen := len(entries)
	loop:
		for i := 0; i < entriesLen; i++ {
			existingEntryIdx := existingIdxMap[entries[i].EntryName()].Idx
			// For any given entry in the list of entries to move check if the entry
			// index is at or after pivot point index, which will require movement
			// set to be generated.

			// Then check if the entries to be moved have the same order in the existing
			// slice, and if not require a movement set to be generated.
			switch movement {
			case ActionWhereBefore:
				if existingEntryIdx >= existingPivotIdx {
					movementRequired = true
					break
				}

				if i == 0 {
					continue
				}

				if existingIdxMap[entries[i-1].EntryName()].Idx >= existingEntryIdx {
					movementRequired = true
					break loop

				}
			case ActionWhereAfter:
				if existingEntryIdx <= existingPivotIdx {
					movementRequired = true
					break
				}

				if i == len(entries)-1 {
					continue
				}

				if existingIdxMap[entries[i+1].EntryName()].Idx < existingEntryIdx {
					movementRequired = true
					break loop

				}

			}
		}

		if !movementRequired {
			return nil, errNoMovements
		}
	}

	expected := make([]Movable, len(existing))

	entriesIdxMap := entriesByName(entries)

	filtered := removeEntriesFromExisting(existing, func(entry Movable) bool {
		_, ok := entriesIdxMap[entry.EntryName()]
		return ok
	})

	filteredPivotIdx, pivotEntry := findPivotIdx(filtered, pivot)

	switch movement {
	case ActionWhereBefore:
		expectedIdx := 0

		if direct {
			// If this is a direct move, all entries preceeding the pivot are moved
			// to the beginning of the list
			for ; expectedIdx < filteredPivotIdx; expectedIdx++ {
				expected[expectedIdx] = filtered[expectedIdx]
			}
		} else {
			for _, elt := range filtered {
				filteredInExisting, _ := existingIdxMap[elt.EntryName()]
				firstEntryInExisting, _ := existingIdxMap[entries[0].EntryName()]
				if filteredInExisting.Idx < firstEntryInExisting.Idx && filteredInExisting.Idx < filteredPivotIdx {
					filtered = filtered[1:]
					expected[expectedIdx] = filteredInExisting.Entry
					expectedIdx += 1
				}
			}

			filteredPivotIdx, _ = findPivotIdx(filtered, pivot)
		}

		for _, elt := range entries {
			expected[expectedIdx] = elt
			expectedIdx++
		}

		filteredLen := len(filtered)

		if !direct {
			filteredIdx := 0
			for ; filteredIdx < filteredPivotIdx; expectedIdx++ {
				expected[expectedIdx] = filtered[filteredIdx]
				filteredIdx += 1
			}
		}

		expected[expectedIdx] = pivotEntry
		expectedIdx++

		for i := filteredPivotIdx + 1; i < filteredLen; i++ {
			expected[expectedIdx] = filtered[i]
			expectedIdx++
		}

	case ActionWhereAfter:
		expectedIdx := 0
		for ; expectedIdx < filteredPivotIdx+1; expectedIdx++ {
			expected[expectedIdx] = filtered[expectedIdx]
		}

		if direct {
			for _, elt := range entries {
				expected[expectedIdx] = elt
				expectedIdx++
			}

			filteredLen := len(filtered)
			for i := filteredPivotIdx + 1; i < filteredLen; i++ {
				expected[expectedIdx] = filtered[i]
				expectedIdx++
			}
		} else {
			filteredLen := len(filtered)
			for i := filteredPivotIdx + 1; i < filteredLen; i++ {
				expected[expectedIdx] = filtered[i]
				expectedIdx++
			}

			for _, elt := range entries {
				expected[expectedIdx] = elt
				expectedIdx++
			}

		}
	}

	return expected, nil
}

func (o PositionAfter) GetExpected(entries []Movable, existing []Movable) ([]Movable, error) {
	return getPivotMovement(entries, existing, o.Pivot, o.Directly, ActionWhereAfter)
}

func (o PositionAfter) Move(entries []Movable, existing []Movable) ([]MoveAction, error) {
	expected, err := o.GetExpected(entries, existing)
	if err != nil {
		if errors.Is(err, errNoMovements) {
			return nil, nil
		}
		return nil, err
	}

	actions, err := GenerateMovements(existing, expected, entries, ActionWhereAfter, o.Pivot, o.Directly)
	if err != nil {
		return nil, err
	}

	return OptimizeMovements(existing, expected, entries, actions, o), nil
}

func (o PositionBefore) GetExpected(entries []Movable, existing []Movable) ([]Movable, error) {
	return getPivotMovement(entries, existing, o.Pivot, o.Directly, ActionWhereBefore)
}

func (o PositionBefore) Move(entries []Movable, existing []Movable) ([]MoveAction, error) {
	expected, err := o.GetExpected(entries, existing)
	if err != nil {
		if errors.Is(err, errNoMovements) {
			return nil, nil
		}
		return nil, err
	}

	actions, err := GenerateMovements(existing, expected, entries, ActionWhereBefore, o.Pivot, o.Directly)
	if err != nil {
		return nil, err
	}

	return OptimizeMovements(existing, expected, entries, actions, o), nil
}

type Entry struct {
	Element  Movable
	Expected int
	Existing int
}

type sequencePosition struct {
	Start int
	End   int
}

func updateSimulatedIdxMap[E Movable](idxMap *map[string]entryWithIdx[E], moved Movable, startingIdx int, targetIdx int) {
	for name, entry := range *idxMap {
		if name == moved.EntryName() {
			continue
		}

		idx := entry.Idx

		if startingIdx > targetIdx && idx >= targetIdx {
			entry.Idx = idx + 1
			(*idxMap)[name] = entry
		} else if startingIdx < targetIdx && idx >= startingIdx && idx <= targetIdx {
			entry.Idx = idx - 1
			(*idxMap)[name] = entry
		}
	}
}

func OptimizeMovements(existing []Movable, expected []Movable, entries []Movable, actions []MoveAction, position Position) []MoveAction {
	simulated := make([]Movable, len(existing))
	copy(simulated, existing)
	simulatedIdxMap := entriesByName(simulated)

	var optimized []MoveAction
	for _, action := range actions {
		currentIdx := simulatedIdxMap[action.Movable.EntryName()].Idx

		var targetIdx int
		switch action.Where {
		case ActionWhereFirst:
			targetIdx = 0
		case ActionWhereLast:
			targetIdx = len(simulated) - 1
		case ActionWhereBefore:
			targetIdx = simulatedIdxMap[action.Destination.EntryName()].Idx
		case ActionWhereAfter:
			targetIdx = simulatedIdxMap[action.Destination.EntryName()].Idx + 1
		}

		if targetIdx != currentIdx {
			optimized = append(optimized, action)
			entry := simulatedIdxMap[action.Movable.EntryName()]
			entry.Idx = targetIdx
			simulatedIdxMap[action.Movable.EntryName()] = entry
			updateSimulatedIdxMap(&simulatedIdxMap, action.Movable, currentIdx, targetIdx)
		}
	}

	slog.Debug("OptimizeMovements()", "optimized", movementsAsDebug(optimized))

	return optimized
}

func entryWithIdxAsDebug[E Movable](entry entryWithIdx[E]) string {
	return fmt.Sprintf("{Entry:%s Idx:%d}", entry.Entry.EntryName(), entry.Idx)
}

func movablesAsDebug(movables []Movable) string {
	var movablesFormatted []string
	for _, elt := range movables {
		var entryName string
		if elt == nil {
			entryName = "<MISSING>"
		} else {
			entryName = elt.EntryName()
		}
		movablesFormatted = append(movablesFormatted, entryName)
	}

	return fmt.Sprintf("[%s]", strings.Join(movablesFormatted, " "))
}

func movementsAsDebug(movements []MoveAction) string {
	var movementsFormatted []string
	for _, elt := range movements {
		var destination string
		if elt.Destination != nil {
			destination = elt.Destination.EntryName()
		}
		movementsFormatted = append(movementsFormatted, fmt.Sprintf("{Movable:%s Where:%s Destination:%s}", elt.Movable.EntryName(), elt.Where, destination))
	}

	return fmt.Sprintf("[%s]", strings.Join(movementsFormatted, " "))
}

func GenerateMovements(existing []Movable, expected []Movable, entries []Movable, movement ActionWhereType, pivot string, directly bool) ([]MoveAction, error) {
	if len(existing) != len(expected) {
		slog.Error("GenerateMovements()", "len(existing)", len(existing), "len(expected)", len(expected))
		return nil, ErrSlicesNotEqualLength
	}

	// Early optimization to skip generation of movements if existing list already matches expected one
	var needsSorting bool
	for idx := range existing {
		if existing[idx].EntryName() != expected[idx].EntryName() {
			needsSorting = true
		}
	}

	if !needsSorting {
		return nil, nil
	}

	entriesIdxMap := entriesByName(entries)
	existingIdxMap := entriesByName(existing)
	expectedIdxMap := entriesByName(expected)

	var movements []MoveAction
	var previous Movable
	for _, elt := range entries {
		slog.Debug("GenerateMovements()", "elt", elt, "existing", existingIdxMap[elt.EntryName()], "expected", expectedIdxMap[elt.EntryName()])

		if previous != nil {
			movements = append(movements, MoveAction{
				Movable:     elt,
				Destination: previous,
				Where:       ActionWhereAfter,
			})
			previous = elt
			continue
		}
		if expectedIdxMap[elt.EntryName()].Idx == 0 {
			movements = append(movements, MoveAction{
				Movable:     elt,
				Destination: nil,
				Where:       ActionWhereFirst,
			})
			previous = elt
		} else if expectedIdxMap[elt.EntryName()].Idx == len(expectedIdxMap)-1 {
			movements = append(movements, MoveAction{
				Movable:     elt,
				Destination: nil,
				Where:       ActionWhereLast,
			})
			previous = elt
		} else {
			var where ActionWhereType
			var pivot Movable

			switch movement {
			case ActionWhereLast:
				where = ActionWhereLast
			case ActionWhereAfter:
				pivot = expected[expectedIdxMap[elt.EntryName()].Idx-1]
				where = ActionWhereAfter
			case ActionWhereFirst:
				pivot = existing[0]
				where = ActionWhereBefore
			case ActionWhereBefore:
				eltExpectedIdx := expectedIdxMap[elt.EntryName()].Idx
				pivot = expected[eltExpectedIdx+1]
				where = ActionWhereBefore
				// When entries are to be put directly before the pivot point, if previous was nil (we
				// are processing the first element in entries set) and selected pivot is part of the
				// entries set, we need to find the actual pivot, i.e. element of the expected list
				// that directly follows all elements from the entries set.
				if _, ok := entriesIdxMap[pivot.EntryName()]; ok {
					if directly {
						// If the move is direct, he actual pivot for the move is the element that
						// follows all elements from the existing set.
						pivotIdx := eltExpectedIdx + len(entries)
						if pivotIdx >= len(expected) {
							// This should never happen as by definition there is at least
							// element (pivot point) at the end of the expected slice.
							return nil, ErrInvalidMovementPlan
						}
						pivot = expected[pivotIdx]
					} else {
						// Otherwise, if the move is not direct we use the element preceeding the
						// first element of the entries list as a pivot and change where action
						// to ActionWhereAfter.
						pivotIdx := eltExpectedIdx - 1
						if pivotIdx < 0 {
							return nil, ErrInvalidMovementPlan
						} else {
							pivot = expected[pivotIdx]
							where = ActionWhereAfter
						}
					}
				}
			}

			movements = append(movements, MoveAction{
				Movable:     elt,
				Destination: pivot,
				Where:       where,
			})
		}

		previous = elt
	}

	slog.Debug("GenerateMovements()", "movements", movementsAsDebug(movements))
	return movements, nil
}

func (o PositionFirst) GetExpected(entries []Movable, existing []Movable) ([]Movable, error) {
	entriesIdxMap := entriesByName(entries)

	filtered := removeEntriesFromExisting(existing, func(entry Movable) bool {
		_, ok := entriesIdxMap[entry.EntryName()]
		return ok
	})

	expected := append(entries, filtered...)

	return expected, nil
}

func (o PositionFirst) Move(entries []Movable, existing []Movable) ([]MoveAction, error) {
	expected, err := o.GetExpected(entries, existing)
	if err != nil {
		return nil, err
	}

	slog.Error("PositionFirst.Move()", "len(expected)", len(expected), "len(existing)", len(existing))

	actions, err := GenerateMovements(existing, expected, entries, ActionWhereFirst, "", false)
	if err != nil {
		return nil, err
	}

	return OptimizeMovements(existing, expected, entries, actions, o), nil
}

func (o PositionLast) GetExpected(entries []Movable, existing []Movable) ([]Movable, error) {
	entriesIdxMap := entriesByName(entries)

	filtered := removeEntriesFromExisting(existing, func(entry Movable) bool {
		_, ok := entriesIdxMap[entry.EntryName()]
		return ok
	})

	expected := append(filtered, entries...)

	return expected, nil
}

func (o PositionLast) Move(entries []Movable, existing []Movable) ([]MoveAction, error) {
	expected, err := o.GetExpected(entries, existing)
	if err != nil {
		return nil, err
	}

	actions, err := GenerateMovements(existing, expected, entries, ActionWhereLast, "", false)
	if err != nil {
		slog.Debug("PositionLast()", "err", err)
		return nil, err
	}
	return OptimizeMovements(existing, expected, entries, actions, o), nil
}

type Movement struct {
	Entries  []Movable
	Position Position
}

func MoveGroups[E Movable](existing []Movable, movements []Movement) ([]MoveAction, error) {
	expected := existing
	for idx := range len(movements) - 1 {
		position := movements[idx].Position
		entries := movements[idx].Entries
		result, err := position.GetExpected(entries, expected)
		if err != nil {
			if !errors.Is(err, errNoMovements) {
				return nil, err
			}
			continue
		}
		expected = result
	}

	entries := movements[len(movements)-1].Entries
	position := movements[len(movements)-1].Position
	return position.Move(entries, expected)
}

func MoveGroup[E Movable](position Position, entries []E, existing []E) ([]MoveAction, error) {
	var movableEntries []Movable
	for _, elt := range entries {
		slog.Warn("MoveGroup", "entry.EntryName()", elt.EntryName())
		movableEntries = append(movableEntries, elt)
	}
	var movableExisting []Movable
	for _, elt := range existing {
		slog.Warn("MoveGroup", "existing.EntryName()", elt.EntryName())
		movableExisting = append(movableExisting, elt)
	}
	return position.Move(movableEntries, movableExisting)
}

type Move struct {
	Position Position
	Existing []Movable
}
