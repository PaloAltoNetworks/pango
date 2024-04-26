package service

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
	return s.doRead(ctx, loc, name, action, true)
}

// ReadFromConfig returns the given config object from the loaded XML config.
//
// Requires that client.LoadPanosConfig() has been invoked.
func (s *Service) ReadFromConfig(ctx context.Context, loc Location, name string) (*Entry, error) {
	return s.doRead(ctx, loc, name, "", false)
}

func (s *Service) doRead(ctx context.Context, loc Location, name, action string, fromPanos bool) (*Entry, error) {
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

	if fromPanos {
		cmd := &xmlapi.Config{
			Action: action,
			Xpath:  util.AsXpath(path),
			Target: s.client.GetTarget(),
		}

		if _, _, err := s.client.Communicate(ctx, cmd, true, normalizer); err != nil {
			// action=show returns empty config like this
			if err.Error() == "No such node" && action == "show" {
				return nil, errors.ObjectNotFound()
			}
			return nil, err
		}
	} else {
		if _, err := s.client.ReadFromConfig(ctx, path, true, normalizer); err != nil {
			return nil, err
		}
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
func (s *Service) Delete(ctx context.Context, loc Location, names ...string) error {
	if len(names) == 0 {
		return nil
	} else {
		for _, name := range names {
			if name == "" {
				return errors.NameNotSpecifiedError
			}
		}
	}

	vn := s.client.Versioning()

	updates := xmlapi.NewMultiConfig(len(names))
	for _, name := range names {
		path, err := loc.Xpath(vn, name)
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

// List returns a list of service objects using the given action.
//
// Param action should be either "get" or "show".
//
// Params filter and quote are for client side filtering.
func (s *Service) List(ctx context.Context, loc Location, action, filter, quote string) ([]Entry, error) {
	return s.doList(ctx, loc, action, filter, quote, true)
}

// ListFromConfig returns a list of objects at the given location.
//
// Requires that client.LoadPanosConfig() has been invoked.
//
// Params filter and quote are for client side filtering.
func (s *Service) ListFromConfig(ctx context.Context, loc Location, filter, quote string) ([]Entry, error) {
	return s.doList(ctx, loc, "", filter, quote, false)
}

func (s *Service) doList(ctx context.Context, loc Location, action, filter, quote string, fromPanos bool) ([]Entry, error) {
	logic, err := filtering.Parse(filter, quote)
	if err != nil {
		return nil, err
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

	if fromPanos {
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
	} else {
		path = path[:len(path)-1]

		if _, err = s.client.ReadFromConfig(ctx, path, false, normalizer); err != nil {
			return nil, err
		}
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
