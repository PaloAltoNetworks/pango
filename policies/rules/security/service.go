package security

import (
	"context"
	"fmt"
	"log"

	"github.com/PaloAltoNetworks/pango/audit"
	"github.com/PaloAltoNetworks/pango/errors"
	"github.com/PaloAltoNetworks/pango/filtering"
	"github.com/PaloAltoNetworks/pango/rule"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/xmlapi"
)

type Service struct {
	client util.PangoClient
}

func NewService(client util.PangoClient) *Service {
	return &Service{
		client: client,
	}
}

// Create creates the given config object.
func (s *Service) Create(ctx context.Context, loc Location, entry Entry) (*Entry, error) {
	if entry.Name == "" {
		return nil, errors.NameNotSpecifiedError
	}

	vn := s.client.Versioning()

	// Get versioning stuff.
	specifier, _, err := Versioning(vn)
	if err != nil {
		return nil, err
	}

	// Get the xpath.
	path, err := loc.Xpath(vn, entry.Name, "")
	if err != nil {
		return nil, err
	}

	createSpec, err := specifier(entry)
	if err != nil {
		return nil, err
	}

	cmd := &xmlapi.Config{
		Action:  "set",
		Xpath:   util.AsXpath(path[:len(path)-1]),
		Element: createSpec,
		Target:  s.client.GetTarget(),
	}

	// Perform the set.
	if _, _, err = s.client.Communicate(ctx, cmd, false, nil); err != nil {
		return nil, err
	}

	// Return the Read results.
	return s.Read(ctx, loc, entry.Name, "get")
}

// Read returns the given config object, using the specified action.
//
// Param action should be either "get" or "show".
func (s *Service) Read(ctx context.Context, loc Location, name, action string) (*Entry, error) {
	if name == "" {
		return nil, errors.NameNotSpecifiedError
	}

	vn := s.client.Versioning()
	_, normalizer, err := Versioning(vn)
	if err != nil {
		return nil, err
	}

	path, err := loc.Xpath(vn, name, "")
	if err != nil {
		return nil, err
	}

	cmd := &xmlapi.Config{
		Action: action,
		Xpath:  util.AsXpath(path),
		Target: s.client.GetTarget(),
	}

	if _, _, err = s.client.Communicate(ctx, cmd, true, normalizer); err != nil {
		// action=show returns empty config like this
		if err.Error() == "No such node" && action == "show" {
			return nil, errors.ObjectNotFound()
		}
		return nil, err
	}

	list, err := normalizer.Normalize()
	if err != nil {
		return nil, err
	} else if len(list) != 1 {
		return nil, fmt.Errorf("expected to %q 1 entry, got %d", action, len(list))
	}

	return &list[0], nil
}

// ReadById returns the given config object for the given uuid, using the specified action.
//
// Param action should be either "get" or "show".
func (s *Service) ReadById(ctx context.Context, loc Location, uuid, action string) (*Entry, error) {
	if uuid == "" {
		return nil, errors.UuidNotSpecifiedError
	}

	vn := s.client.Versioning()
	_, normalizer, err := Versioning(vn)
	if err != nil {
		return nil, err
	}

	path, err := loc.Xpath(vn, "", uuid)
	if err != nil {
		return nil, err
	}

	cmd := &xmlapi.Config{
		Action: action,
		Xpath:  util.AsXpath(path),
		Target: s.client.GetTarget(),
	}

	if _, _, err = s.client.Communicate(ctx, cmd, true, normalizer); err != nil {
		// action=show returns empty config like this
		if err.Error() == "No such node" && action == "show" {
			return nil, errors.ObjectNotFound()
		}
		return nil, err
	}

	list, err := normalizer.Normalize()
	if err != nil {
		return nil, err
	} else if len(list) != 1 {
		return nil, fmt.Errorf("expected to %q 1 entry, got %d", action, len(list))
	}

	return &list[0], nil
}

// ReadFromConfig returns the given config object from the loaded XML config.
//
// Requires that client.LoadPanosConfig() has been invoked.
func (s *Service) ReadFromConfig(ctx context.Context, loc Location, name string) (*Entry, error) {
	if name == "" {
		return nil, errors.NameNotSpecifiedError
	}

	vn := s.client.Versioning()
	_, normalizer, err := Versioning(vn)
	if err != nil {
		return nil, err
	}

	path, err := loc.Xpath(vn, name, "")
	if err != nil {
		return nil, err
	}

	if _, err = s.client.ReadFromConfig(ctx, path, true, normalizer); err != nil {
		return nil, err
	}

	list, err := normalizer.Normalize()
	if err != nil {
		return nil, err
	} else if len(list) != 1 {
		return nil, fmt.Errorf("expected to find 1 entry, got %d", len(list))
	}

	return &list[0], nil
}

// ReadFromConfigById returns the given config object for the given UUID from the
// loaded XML config.
//
// Requires that client.LoadPanosConfig() has been invoked.
func (s *Service) ReadFromConfigById(ctx context.Context, loc Location, uuid string) (*Entry, error) {
	if uuid == "" {
		return nil, errors.NameNotSpecifiedError
	}

	vn := s.client.Versioning()
	_, normalizer, err := Versioning(vn)
	if err != nil {
		return nil, err
	}

	path, err := loc.Xpath(vn, "", uuid)
	if err != nil {
		return nil, err
	}

	if _, err = s.client.ReadFromConfig(ctx, path, true, normalizer); err != nil {
		return nil, err
	}

	list, err := normalizer.Normalize()
	if err != nil {
		return nil, err
	} else if len(list) != 1 {
		return nil, fmt.Errorf("expected to find 1 entry, got %d", len(list))
	}

	return &list[0], nil
}

// Update updates the given config object, then returns the result.
func (s *Service) Update(ctx context.Context, loc Location, entry Entry, oldName string) (*Entry, error) {
	if entry.Name == "" {
		return nil, errors.NameNotSpecifiedError
	}

	vn := s.client.Versioning()
	updates := xmlapi.NewMultiConfig(2)
	specifier, _, err := Versioning(vn)
	if err != nil {
		return nil, err
	}

	// Get the old config.
	var old *Entry
	if oldName != "" && oldName != entry.Name {
		// Action needed: rename.
		path, err := loc.Xpath(vn, oldName, "")
		if err != nil {
			return nil, err
		}

		old, err = s.Read(ctx, loc, oldName, "get")

		updates.Add(&xmlapi.Config{
			Action:  "rename",
			Xpath:   util.AsXpath(path),
			NewName: entry.Name,
			Target:  s.client.GetTarget(),
		})
	} else {
		old, err = s.Read(ctx, loc, entry.Name, "get")
	}
	if err != nil {
		return nil, err
	}

	if !SpecMatches(&entry, old) {
		// Action needed: edit.
		path, err := loc.Xpath(vn, entry.Name, "")
		if err != nil {
			return nil, err
		}

		// Copy over the misc stuff.
		entry.CopyMiscFrom(old)

		updateSpec, err := specifier(entry)
		if err != nil {
			return nil, err
		}

		updates.Add(&xmlapi.Config{
			Action:  "edit",
			Xpath:   util.AsXpath(path),
			Element: updateSpec,
			Target:  s.client.GetTarget(),
		})
	}

	// Do the updates we've built up.
	if len(updates.Operations) != 0 {
		if _, _, _, err = s.client.MultiConfig(ctx, updates, false, nil); err != nil {
			return nil, err
		}
	}

	// Return the read results.
	return s.Read(ctx, loc, entry.Name, "get")
}

// UpdateById updates the given config object by uuid, then returns the result.
func (s *Service) UpdateById(ctx context.Context, loc Location, entry Entry, uuid string) (*Entry, error) {
	if entry.Name == "" {
		return nil, errors.NameNotSpecifiedError
	} else if uuid == "" {
		return nil, errors.UuidNotSpecifiedError
	}

	vn := s.client.Versioning()
	updates := xmlapi.NewMultiConfig(2)
	specifier, _, err := Versioning(vn)
	if err != nil {
		return nil, err
	}

	// Get the old config.
	old, err := s.ReadById(ctx, loc, uuid, "get")
	if err != nil {
		return nil, err
	}

	if old.Name != entry.Name {
		// Action needed: rename.
		path, err := loc.Xpath(vn, old.Name, "")
		if err != nil {
			return nil, err
		}

		updates.Add(&xmlapi.Config{
			Action:  "rename",
			Xpath:   util.AsXpath(path),
			NewName: entry.Name,
			Target:  s.client.GetTarget(),
		})
	}

	if !SpecMatches(&entry, old) {
		// Action needed: edit.
		path, err := loc.Xpath(vn, entry.Name, "")
		if err != nil {
			return nil, err
		}

		// Copy over the misc stuff.
		entry.CopyMiscFrom(old)

		updateSpec, err := specifier(entry)
		if err != nil {
			return nil, err
		}

		updates.Add(&xmlapi.Config{
			Action:  "edit",
			Xpath:   util.AsXpath(path),
			Element: updateSpec,
			Target:  s.client.GetTarget(),
		})
	}

	// Do the updates we've built up.
	if len(updates.Operations) != 0 {
		if _, _, _, err = s.client.MultiConfig(ctx, updates, false, nil); err != nil {
			return nil, err
		}
	}

	// Return the read results.
	return s.Read(ctx, loc, entry.Name, "get")
}

// Delete deletes the given item.
func (s *Service) Delete(ctx context.Context, loc Location, name string) error {
	if name == "" {
		return errors.NameNotSpecifiedError
	}

	vn := s.client.Versioning()

	path, err := loc.Xpath(vn, name, "")
	if err != nil {
		return err
	}

	cmd := &xmlapi.Config{
		Action: "delete",
		Xpath:  util.AsXpath(path),
		Target: s.client.GetTarget(),
	}

	_, _, err = s.client.Communicate(ctx, cmd, false, nil)

	return err
}

// DeleteById deletes the given item using the uuid.
func (s *Service) DeleteById(ctx context.Context, loc Location, uuid string) error {
	if uuid == "" {
		return errors.UuidNotSpecifiedError
	}

	vn := s.client.Versioning()

	path, err := loc.Xpath(vn, "", uuid)
	if err != nil {
		return err
	}

	cmd := &xmlapi.Config{
		Action: "delete",
		Xpath:  util.AsXpath(path),
		Target: s.client.GetTarget(),
	}

	_, _, err = s.client.Communicate(ctx, cmd, false, nil)

	return err
}

// List returns a list of service objects using the given action.
//
// Param action should be either "get" or "show".
//
// Params filter and quote are for client side filtering.
func (s *Service) List(ctx context.Context, loc Location, action, filter, quote string) ([]Entry, error) {
	var err error

	var logic *filtering.Group
	if filter != "" {
		logic, err = filtering.Parse(filter, quote)
		if err != nil {
			return nil, err
		}
	}

	vn := s.client.Versioning()

	_, normalizer, err := Versioning(vn)
	if err != nil {
		return nil, err
	}

	path, err := loc.Xpath(vn, "", "")
	if err != nil {
		return nil, err
	}

	cmd := &xmlapi.Config{
		Action: action,
		Xpath:  util.AsXpath(path),
		Target: s.client.GetTarget(),
	}

	if _, _, err = s.client.Communicate(ctx, cmd, true, normalizer); err != nil {
		// action=show returns empty config like this, it is not an error.
		if err.Error() == "No such node" && action == "show" {
			return nil, nil
		}
		return nil, err
	}

	listing, err := normalizer.Normalize()
	if err != nil || logic == nil {
		return listing, err
	}

	filtered := make([]Entry, 0, len(listing))
	for _, x := range listing {
		ok, err := logic.Matches(&x)
		if err != nil {
			return nil, err
		}
		if ok {
			filtered = append(filtered, x)
		}
	}

	return filtered, nil
}

// ListFromConfig returns a list of objects at the given location.
//
// Requires that client.LoadPanosConfig() has been invoked.
//
// Params filter and quote are for client side filtering.
func (s *Service) ListFromConfig(ctx context.Context, loc Location, filter, quote string) ([]Entry, error) {
	var err error

	var logic *filtering.Group
	if filter != "" {
		logic, err = filtering.Parse(filter, quote)
		if err != nil {
			return nil, err
		}
	}

	vn := s.client.Versioning()

	_, normalizer, err := Versioning(vn)
	if err != nil {
		return nil, err
	}

	path, err := loc.Xpath(vn, "", "")
	if err != nil {
		return nil, err
	}
	path = path[:len(path)-1]

	if _, err = s.client.ReadFromConfig(ctx, path, false, normalizer); err != nil {
		return nil, err
	}

	listing, err := normalizer.Normalize()
	if err != nil || logic == nil {
		return listing, err
	}

	filtered := make([]Entry, 0, len(listing))
	for _, x := range listing {
		ok, err := logic.Matches(&x)
		if err != nil {
			return nil, err
		}
		if ok {
			filtered = append(filtered, x)
		}
	}

	return filtered, nil
}

// ConfigureGroup performs all necessary edit, rename, and delete commands to ensure that
// the objects are configured as specified.
//
// If removeEverythingElse is true, then any rule not in the entries param is deleted.
func (s *Service) ConfigureGroup(ctx context.Context, loc Location, position rule.Position, entries []Entry, uuids, auditComments map[string]string, removeEverythingElse bool) ([]Entry, map[string]string, bool, bool, error) {
	if len(entries) == 0 {
		return nil, nil, false, false, fmt.Errorf("no rules given")
	}

	var err error

	listing, err := s.List(ctx, loc, "get", "", "")
	if err != nil {
		return nil, nil, false, false, err
	}

	vn := s.client.Versioning()
	updates := xmlapi.NewMultiConfig(len(entries) + len(listing))
	comments := make([]audit.SetComment, 0, len(entries))
	specifier, _, err := Versioning(vn)
	if err != nil {
		return nil, nil, false, false, err
	}

	newOrRename := make([]Entry, 0, len(entries))
	uuidIsUsed := make(map[string]bool)

	// First pass on the rules given.
	for _, entry := range entries {
		var id string
		for uuid, name := range uuids {
			if name == entry.Name {
				id = uuid
				break
			}
		}

		// If the rule name does not have a previously managed UUID, then this is
		// either a brand new rule or a renamed rule.  In either case, save it for
		// later.
		if id == "" {
			newOrRename = append(newOrRename, entry)
			continue
		}

		// There was a UUID, so see if it still exists.
		for _, live := range listing {
			if live.Uuid != id {
				continue
			}

			// The UUID is still there, so continue using it.
			uuidIsUsed[id] = true

			// Check if the name needs to be updated due to out-of-band modification.
			if live.Name != entry.Name {
				path, err := loc.Xpath(vn, live.Name, "")
				if err != nil {
					return nil, nil, false, false, err
				}
				updates.Add(&xmlapi.Config{
					Action:  "rename",
					Xpath:   util.AsXpath(path),
					NewName: entry.Name,
					Target:  s.client.GetTarget(),
				})
			}

			// Check if the rule spec matches or not.
			if !SpecMatches(&entry, &live) {
				path, err := loc.Xpath(vn, entry.Name, "")
				if err != nil {
					return nil, nil, false, false, err
				}

				entry.CopyMiscFrom(&live)
				elm, err := specifier(entry)
				if err != nil {
					return nil, nil, false, false, err
				}

				updates.Add(&xmlapi.Config{
					Action:  "edit",
					Xpath:   util.AsXpath(path),
					Element: elm,
					Target:  s.client.GetTarget(),
				})

				if auditComments[entry.Name] != "" {
					comments = append(comments, audit.SetComment{
						Xpath:   util.AsXpath(path),
						Comment: auditComments[entry.Name],
					})
				}
			}

			break
		}

		if !uuidIsUsed[id] {
			newOrRename = append(newOrRename, entry)
		}
	}

	// At this point, we only have new or renamed rules to deal with.
	// A renamed rule will be a rule that exactly matches one and only one other
	// rule, and that rule also has a uuid that is not currently being used.
	for _, entry := range newOrRename {
		matches := make([]Entry, 0, len(listing))
		for _, live := range listing {
			if live.Uuid == "" || uuidIsUsed[live.Uuid] {
				continue
			}

			if SpecMatches(&entry, &live) {
				matches = append(matches, live)
			}
		}

		if len(matches) == 1 && !uuidIsUsed[matches[0].Uuid] {
			path, err := loc.Xpath(vn, matches[0].Name, "")
			if err != nil {
				return nil, nil, false, false, err
			}

			updates.Add(&xmlapi.Config{
				Action:  "rename",
				Xpath:   util.AsXpath(path),
				NewName: entry.Name,
				Target:  s.client.GetTarget(),
			})

			uuidIsUsed[matches[0].Uuid] = true
		} else {
			// This is a new rule.
			path, err := loc.Xpath(vn, entry.Name, "")
			if err != nil {
				return nil, nil, false, false, err
			}

			elm, err := specifier(entry)
			if err != nil {
				return nil, nil, false, false, err
			}

			updates.Add(&xmlapi.Config{
				Action:  "edit",
				Xpath:   util.AsXpath(path),
				Element: elm,
				Target:  s.client.GetTarget(),
			})

			if auditComments[entry.Name] != "" {
				comments = append(comments, audit.SetComment{
					Xpath:   util.AsXpath(path),
					Comment: auditComments[entry.Name],
				})
			}
		}
	}

	// Rule deletion.  We're either deleting everything not specified or old
	// rules that aren't in use anymore.
	if removeEverythingElse {
		for _, live := range listing {
			if live.Uuid != "" && !uuidIsUsed[live.Uuid] {
				path, err := loc.Xpath(vn, "", live.Uuid)
				if err != nil {
					return nil, nil, false, false, err
				}

				updates.Add(&xmlapi.Config{
					Action: "delete",
					Xpath:  util.AsXpath(path),
					Target: s.client.GetTarget(),
				})
			}
		}
	} else {
		// Delete unused security rules from the previous group.
		for uuid := range uuids {
			if !uuidIsUsed[uuid] {
				path, err := loc.Xpath(vn, "", uuid)
				if err != nil {
					return nil, nil, false, false, err
				}

				updates.Add(&xmlapi.Config{
					Action: "delete",
					Xpath:  util.AsXpath(path),
					Target: s.client.GetTarget(),
				})
			}
		}
	}

	// It's possible the rules were fine, so don't run an empty multi-config.
	if len(updates.Operations) > 0 {
		if false {
			log.Printf("%d operations, doing multi-config", len(updates.Operations))
		}
		if _, _, _, err = s.client.MultiConfig(ctx, updates, false, nil); err != nil {
			return nil, nil, false, false, err
		}

		// Update audit comments.
		if false {
			log.Printf("%d audit comments to be applied", len(comments))
		}
		for _, ac := range comments {
			cmd := &xmlapi.Op{
				Command: ac,
				Target:  s.client.GetTarget(),
			}
			if _, _, err = s.client.Communicate(ctx, cmd, false, nil); err != nil {
				return nil, nil, false, false, err
			}
		}

		// Get an updated listing in preparation to move all rules into place.
		listing, err = s.List(ctx, loc, "get", "", "")
		if err != nil {
			return nil, nil, false, false, err
		} else if len(listing) == 0 {
			return nil, nil, false, false, fmt.Errorf("no rules present before move operations")
		}
	}

	// Rule placement mapping.
	rp := make(map[string]int)
	for idx, live := range listing {
		rp[live.Name] = idx
	}

	// TODO: Now we need to verify positioning of the rules.
	updates = xmlapi.NewMultiConfig(len(entries))

	var ok, topDown bool
	var otherIndex int
	baseIndex := -1
	switch {
	case position.First != nil && *position.First:
		topDown = true
		target := entries[0]

		baseIndex, ok = rp[target.Name]
		if !ok {
			return nil, nil, false, false, fmt.Errorf("could not find rule %q for first positioning", target.Name)
		}

		if baseIndex != 0 {
			path, err := loc.Xpath(vn, target.Name, "")
			if err != nil {
				return nil, nil, false, false, err
			}

			for name, val := range rp {
				switch {
				case name == entries[0].Name:
					rp[name] = 0
				case val < baseIndex:
					rp[name] = val + 1
				}
			}

			updates.Add(&xmlapi.Config{
				Action: "move",
				Xpath:  util.AsXpath(path),
				Where:  "top",
				Target: s.client.GetTarget(),
			})

			baseIndex = 0
		}
	case position.Last != nil && *position.Last:
		target := entries[len(entries)-1]

		baseIndex, ok = rp[target.Name]
		if !ok {
			return nil, nil, false, false, fmt.Errorf("could not find rule %q for last positioning", target.Name)
		}

		if baseIndex != len(listing)-1 {
			path, err := loc.Xpath(vn, target.Name, "")
			if err != nil {
				return nil, nil, false, false, err
			}

			for name, val := range rp {
				switch {
				case name == target.Name:
					rp[name] = len(listing) - 1
				case val > baseIndex:
					rp[name] = val - 1
				}
			}

			updates.Add(&xmlapi.Config{
				Action: "move",
				Xpath:  util.AsXpath(path),
				Where:  "bottom",
				Target: s.client.GetTarget(),
			})

			baseIndex = len(listing) - 1
		}
	case position.SomewhereAfter != nil && *position.SomewhereAfter != "":
		topDown = true
		target := entries[0]

		baseIndex, ok = rp[target.Name]
		if !ok {
			return nil, nil, false, false, fmt.Errorf("could not find rule %q for initial positioning", target.Name)
		}

		otherIndex, ok = rp[*position.SomewhereAfter]
		if !ok {
			return nil, nil, false, false, fmt.Errorf("could not find referenced rule %q for initial positioning", *position.SomewhereAfter)
		}

		if baseIndex < otherIndex {
			path, err := loc.Xpath(vn, target.Name, "")
			if err != nil {
				return nil, nil, false, false, err
			}

			for name, val := range rp {
				switch {
				case name == target.Name:
					rp[name] = otherIndex
				case val > baseIndex && val <= otherIndex:
					rp[name] = otherIndex - 1
				}
			}

			updates.Add(&xmlapi.Config{
				Action:      "move",
				Xpath:       util.AsXpath(path),
				Where:       "after",
				Destination: *position.SomewhereAfter,
				Target:      s.client.GetTarget(),
			})

			baseIndex = otherIndex
		}
	case position.SomewhereBefore != nil && *position.SomewhereBefore != "":
		target := entries[len(entries)-1]

		baseIndex, ok = rp[target.Name]
		if !ok {
			return nil, nil, false, false, fmt.Errorf("could not find rule %q for initial positioning", target.Name)
		}

		otherIndex, ok = rp[*position.SomewhereBefore]
		if !ok {
			return nil, nil, false, false, fmt.Errorf("could not find referenced rule %q", *position.SomewhereBefore)
		}

		if baseIndex > otherIndex {
			path, err := loc.Xpath(vn, target.Name, "")
			if err != nil {
				return nil, nil, false, false, err
			}

			for name, val := range rp {
				switch {
				case name == target.Name:
					rp[name] = otherIndex
				case val < baseIndex && val >= otherIndex:
					rp[name] = val + 1
				}
			}

			updates.Add(&xmlapi.Config{
				Action:      "move",
				Xpath:       util.AsXpath(path),
				Where:       "before",
				Destination: *position.SomewhereBefore,
				Target:      s.client.GetTarget(),
			})

			baseIndex = otherIndex
		}
	case position.DirectlyAfter != nil && *position.DirectlyAfter != "":
		topDown = true
		target := entries[0]

		baseIndex, ok = rp[target.Name]
		if !ok {
			return nil, nil, false, false, fmt.Errorf("could not find rule %q for initial positioning", target.Name)
		}

		otherIndex, ok = rp[*position.DirectlyAfter]
		if !ok {
			return nil, nil, false, false, fmt.Errorf("could not find referenced rule %q for initial positioning", *position.DirectlyAfter)
		}

		if baseIndex != otherIndex+1 {
			path, err := loc.Xpath(vn, target.Name, "")
			if err != nil {
				return nil, nil, false, false, err
			}

			for name, val := range rp {
				switch {
				case name == target.Name:
					rp[name] = otherIndex
				case val > baseIndex && val <= otherIndex:
					rp[name] = otherIndex - 1
				}
			}

			updates.Add(&xmlapi.Config{
				Action:      "move",
				Xpath:       util.AsXpath(path),
				Where:       "after",
				Destination: *position.DirectlyAfter,
				Target:      s.client.GetTarget(),
			})

			baseIndex = otherIndex
		}
	case position.DirectlyBefore != nil && *position.DirectlyBefore != "":
		target := entries[len(entries)-1]

		baseIndex, ok = rp[target.Name]
		if !ok {
			return nil, nil, false, false, fmt.Errorf("could not find rule %q for initial positioning", target.Name)
		}

		otherIndex, ok = rp[*position.DirectlyBefore]
		if !ok {
			return nil, nil, false, false, fmt.Errorf("could not find referenced rule %q", *position.DirectlyBefore)
		}

		if baseIndex+1 != otherIndex {
			path, err := loc.Xpath(vn, target.Name, "")
			if err != nil {
				return nil, nil, false, false, err
			}

			for name, val := range rp {
				switch {
				case name == target.Name:
					rp[name] = otherIndex
				case val < baseIndex && val >= otherIndex:
					rp[name] = val + 1
				}
			}

			updates.Add(&xmlapi.Config{
				Action:      "move",
				Xpath:       util.AsXpath(path),
				Where:       "before",
				Destination: *position.DirectlyBefore,
				Target:      s.client.GetTarget(),
			})

			baseIndex = otherIndex
		}
	default:
		topDown = true
		target := entries[0]

		baseIndex, ok = rp[target.Name]
		if !ok {
			return nil, nil, false, false, fmt.Errorf("could not find rule %q for first positioning", target.Name)
		}

	}

	var prevName, where string
	if topDown {
		prevName = entries[0].Name
		where = "after"
	} else {
		prevName = entries[len(entries)-1].Name
		where = "before"
	}

	// Move the rest of the rules if necessary.
	for i := 1; i < len(entries); i++ {
		var target Entry
		var desiredIndex int
		if topDown {
			target = entries[i]
			desiredIndex = baseIndex + i
		} else {
			target = entries[len(entries)-1-i]
			desiredIndex = baseIndex - i
		}

		idx, ok := rp[target.Name]
		if !ok {
			return nil, nil, false, false, fmt.Errorf("rule %q not present", target.Name)
		}

		if idx != desiredIndex {
			path, err := loc.Xpath(vn, target.Name, "")
			if err != nil {
				return nil, nil, false, false, err
			}

			if idx < desiredIndex {
				for name, val := range rp {
					if val > idx && val <= desiredIndex {
						rp[name] = val - 1
					}
				}
			} else {
				for name, val := range rp {
					if val < idx && val >= desiredIndex {
						rp[name] = val + 1
					}
				}
			}
			rp[target.Name] = desiredIndex

			updates.Add(&xmlapi.Config{
				Action:      "move",
				Xpath:       util.AsXpath(path),
				Where:       where,
				Destination: prevName,
				Target:      s.client.GetTarget(),
			})
		}

		prevName = target.Name
	}

	if len(updates.Operations) > 0 {
		if false {
			log.Printf("%d operations, doing multi-config", len(updates.Operations))
		}
		if _, _, _, err = s.client.MultiConfig(ctx, updates, false, nil); err != nil {
			return nil, nil, false, false, err
		}
	}

	// Build out the uuid map and return ReadGroup.
	uuidList := make([]string, 0, len(entries))
	for _, entry := range entries {
		for _, live := range listing {
			if live.Name == entry.Name {
				uuidList = append(uuidList, live.Uuid)
				break
			}
		}
	}

	// Done.
	return s.ReadGroup(ctx, loc, position, uuidList)
}

// ReadGroup returns the config for the given rule UUIDs.
//
// The first boolean returned is true if the UUIDs are contiguous.
// The second boolean returned returns if the given position matches.
func (s *Service) ReadGroup(ctx context.Context, loc Location, position rule.Position, uuids []string) ([]Entry, map[string]string, bool, bool, error) {
	objs, err := s.List(ctx, loc, "get", "", "")
	if err != nil {
		return nil, nil, false, false, err
	}

	contiguous := true
	listing := make([]Entry, 0, len(uuids))

	firstIndex, prev := -1, -1
	uuidMap := make(map[string]string, len(uuids))
	for _, uuid := range uuids {
		for num, entry := range objs {
			if entry.Uuid == uuid {
				uuidMap[entry.Uuid] = entry.Name
				if contiguous && prev != -1 {
					contiguous = num == prev+1
				}
				listing = append(listing, entry)
				prev = num
				break
			}
		}

		if firstIndex == -1 && prev != -1 {
			firstIndex = prev
		}
	}

	var positionOk, found bool
	if len(objs) > 0 && len(listing) > 0 && firstIndex != -1 {
		switch {
		case position.First != nil && *position.First:
			found = true
			positionOk = objs[0].Uuid == listing[0].Uuid
		case position.Last != nil && *position.Last:
			found = true
			positionOk = contiguous && firstIndex+len(listing) == len(objs)
		case position.SomewhereBefore != nil && *position.SomewhereBefore != "":
			for otherIndex, entry := range objs {
				if entry.Name == *position.SomewhereBefore {
					found = true
					positionOk = firstIndex < otherIndex
					break
				}
			}
		case position.DirectlyBefore != nil && *position.DirectlyBefore != "":
			for otherIndex, entry := range objs {
				if entry.Name == *position.DirectlyBefore {
					found = true
					positionOk = firstIndex+len(listing) == otherIndex
					break
				}
			}
		case position.SomewhereAfter != nil && *position.SomewhereAfter != "":
			for otherIndex, entry := range objs {
				if entry.Name == *position.SomewhereAfter {
					found = true
					positionOk = firstIndex > otherIndex
					break
				}
			}
		case position.DirectlyAfter != nil && *position.DirectlyAfter != "":
			for otherIndex, entry := range objs {
				if entry.Name == *position.DirectlyAfter {
					found = true
					positionOk = firstIndex-1 == otherIndex
					break
				}
			}
		default:
			found, positionOk = true, true
		}

		if !found {
			return listing, uuidMap, contiguous, false, fmt.Errorf("referenced rule does not exist")
		}
	}

	return listing, uuidMap, contiguous, positionOk, nil
}

// DeleteGroup deletes the given uuids.
func (s *Service) DeleteGroup(ctx context.Context, loc Location, uuids map[string]string) error {
	if len(uuids) == 0 {
		return errors.UuidNotSpecifiedError
	}

	vn := s.client.Versioning()
	updates := xmlapi.NewMultiConfig(len(uuids))

	for uuid := range uuids {
		path, err := loc.Xpath(vn, "", uuid)
		if err != nil {
			return err
		}

		updates.Add(&xmlapi.Config{
			Action: "delete",
			Xpath:  util.AsXpath(path),
			Target: s.client.GetTarget(),
		})
	}

	_, _, _, err := s.client.MultiConfig(ctx, updates, false, nil)

	return err
}
