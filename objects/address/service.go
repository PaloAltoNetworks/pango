package address

import (
	"context"
	"fmt"

	"github.com/PaloAltoNetworks/pango/errors"
	"github.com/PaloAltoNetworks/pango/filtering"
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
	path, err := loc.Xpath(vn, entry.Name)
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

	path, err := loc.Xpath(vn, name)
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

	path, err := loc.Xpath(vn, name)
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
		path, err := loc.Xpath(vn, oldName)
		if err != nil {
			return nil, err
		}

		old, err = s.Read(ctx, loc, oldName, "get")

		updates.Add(&xmlapi.Config{
			Action:  "rename",
			Xpath:   util.AsXpath(path),
			NewName: entry.Name,
		})
	} else {
		old, err = s.Read(ctx, loc, entry.Name, "get")
	}
	if err != nil {
		return nil, err
	}

	if !SpecMatches(&entry, old) {
		// Action needed: edit.
		path, err := loc.Xpath(vn, entry.Name)
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

	path, err := loc.Xpath(vn, name)
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

	path, err := loc.Xpath(vn, "")
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

	path, err := loc.Xpath(vn, "")
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

// ConfigureGroup performs all necessary set / edit / delete commands to ensure that the
// objects are configured as specified.
func (s *Service) ConfigureGroup(ctx context.Context, loc Location, entries []Entry, prevNames []string) ([]Entry, error) {
	var err error

	vn := s.client.Versioning()
	updates := xmlapi.NewMultiConfig(len(prevNames) + len(entries))
	specifier, _, err := Versioning(vn)
	if err != nil {
		return nil, err
	}

	curObjs, err := s.List(ctx, loc, "get", "", "")
	if err != nil {
		return nil, err
	}

	//unfound := make([]Entry, 0, len(entries))

	// Determine set vs edit for desired objects.
	for _, entry := range entries {
		var found bool
		for _, live := range curObjs {
			if entry.Name == live.Name {
				found = true
				if !SpecMatches(&entry, &live) {
					path, err := loc.Xpath(vn, entry.Name)
					if err != nil {
						return nil, err
					}

					// Copy over the misc stuff.
					entry.CopyMiscFrom(&live)

					elm, err := specifier(entry)
					if err != nil {
						return nil, err
					}

					updates.Add(&xmlapi.Config{
						Action:  "edit",
						Xpath:   util.AsXpath(path),
						Element: elm,
					})
				}
				break
			}
		}

		if !found {
			path, err := loc.Xpath(vn, entry.Name)
			if err != nil {
				return nil, err
			}

			elm, err := specifier(entry)
			if err != nil {
				return nil, err
			}

			updates.Add(&xmlapi.Config{
				Action:  "set",
				Xpath:   util.AsXpath(path),
				Element: elm,
			})
		}
	}

	// Determine which old objects need to be removed.
	if len(prevNames) != 0 {
		for _, name := range prevNames {
			var found bool
			for _, entry := range entries {
				if entry.Name == name {
					found = true
					break
				}
			}

			if !found {
				path, err := loc.Xpath(vn, name)
				if err != nil {
					return nil, err
				}

				updates.Add(&xmlapi.Config{
					Action: "delete",
					Xpath:  util.AsXpath(path),
				})
			}
		}
	}

	// Perform the multi-config if there was stuff to do.
	if len(updates.Operations) != 0 {
		if _, _, _, err = s.client.MultiConfig(ctx, updates, false, nil); err != nil {
			return nil, err
		}
	}

	// Get the live version of the entries passed in.
	curObjs, err = s.List(ctx, loc, "get", "", "")
	if err != nil {
		return nil, err
	}

	ans := make([]Entry, 0, len(entries))
	for _, entry := range entries {
		for _, live := range curObjs {
			if entry.Name == live.Name {
				ans = append(ans, live)
				break
			}
		}
	}

	// Done.
	return ans, nil
}
