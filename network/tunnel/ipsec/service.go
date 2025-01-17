package ipsec

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

// Create adds new item, then returns the result.
func (s *Service) Create(ctx context.Context, loc Location, entry *Entry) (*Entry, error) {
	if entry.Name == "" {
		return nil, errors.NameNotSpecifiedError
	}

	vn := s.client.Versioning()

	specifier, _, err := Versioning(vn)
	if err != nil {
		return nil, err
	}
	path, err := loc.XpathWithEntryName(vn, entry.Name)
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

	if _, _, err = s.client.Communicate(ctx, cmd, false, nil); err != nil {
		return nil, err
	}
	return s.Read(ctx, loc, entry.Name, "get")
}

// Read returns the given config object, using the specified action ("get" or "show").
func (s *Service) Read(ctx context.Context, loc Location, name, action string) (*Entry, error) {
	return s.read(ctx, loc, name, action, false)
}

// ReadFromConfig returns the given config object from the loaded XML config.
// Requires that client.LoadPanosConfig() has been invoked.
func (s *Service) ReadFromConfig(ctx context.Context, loc Location, name string) (*Entry, error) {
	return s.read(ctx, loc, name, "", true)
}

func (s *Service) read(ctx context.Context, loc Location, value, action string, usePanosConfig bool) (*Entry, error) {
	if value == "" {
		return nil, errors.NameNotSpecifiedError
	}
	vn := s.client.Versioning()
	_, normalizer, err := Versioning(vn)
	if err != nil {
		return nil, err
	}
	var path []string
	path, err = loc.XpathWithEntryName(vn, value)
	if err != nil {
		return nil, err
	}

	if usePanosConfig {
		if _, err = s.client.ReadFromConfig(ctx, path, true, normalizer); err != nil {
			return nil, err
		}
	} else {
		cmd := &xmlapi.Config{
			Action: action,
			Xpath:  util.AsXpath(path),
			Target: s.client.GetTarget(),
		}

		if _, _, err = s.client.Communicate(ctx, cmd, true, normalizer); err != nil {
			if err.Error() == "No such node" && action == "show" {
				return nil, errors.ObjectNotFound()
			}
			return nil, err
		}
	}

	list, err := normalizer.Normalize()
	if err != nil {
		return nil, err
	} else if len(list) != 1 {
		return nil, fmt.Errorf("expected to %q 1 entry, got %d", action, len(list))
	}

	return list[0], nil
}

// Update updates the given config object, then returns the result.
func (s *Service) Update(ctx context.Context, loc Location, entry *Entry, name string) (*Entry, error) {
	return s.update(ctx, loc, entry, name)
}
func (s *Service) update(ctx context.Context, loc Location, entry *Entry, value string) (*Entry, error) {
	if entry.Name == "" {
		return nil, errors.NameNotSpecifiedError
	}

	vn := s.client.Versioning()
	updates := xmlapi.NewMultiConfig(2)
	specifier, _, err := Versioning(vn)
	if err != nil {
		return nil, err
	}
	var old *Entry
	if value != "" && value != entry.Name {
		path, err := loc.XpathWithEntryName(vn, value)
		if err != nil {
			return nil, err
		}

		old, err = s.Read(ctx, loc, value, "get")

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
	} else if old == nil {
		return nil, fmt.Errorf("previous object doesn't exist for update")
	}
	if !SpecMatches(entry, old) {
		path, err := loc.XpathWithEntryName(vn, entry.Name)
		if err != nil {
			return nil, err
		}

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

	if len(updates.Operations) != 0 {
		if _, _, _, err = s.client.MultiConfig(ctx, updates, false, nil); err != nil {
			return nil, err
		}
	}
	return s.Read(ctx, loc, entry.Name, "get")
}

// Delete deletes the given item.
func (s *Service) Delete(ctx context.Context, loc Location, name ...string) error {
	return s.delete(ctx, loc, name)
}
func (s *Service) delete(ctx context.Context, loc Location, values []string) error {
	for _, value := range values {
		if value == "" {
			return errors.NameNotSpecifiedError
		}
	}

	vn := s.client.Versioning()
	var err error
	deletes := xmlapi.NewMultiConfig(len(values))
	for _, value := range values {
		var path []string
		path, err = loc.XpathWithEntryName(vn, value)
		if err != nil {
			return err
		}
		deletes.Add(&xmlapi.Config{
			Action: "delete",
			Xpath:  util.AsXpath(path),
			Target: s.client.GetTarget(),
		})
	}

	_, _, _, err = s.client.MultiConfig(ctx, deletes, false, nil)

	return err
}

// List returns a list of objects using the given action ("get" or "show").
// Params filter and quote are for client side filtering.
func (s *Service) List(ctx context.Context, loc Location, action, filter, quote string) ([]*Entry, error) {
	return s.list(ctx, loc, action, filter, quote, false)
}

// ListFromConfig returns a list of objects at the given location.
// Requires that client.LoadPanosConfig() has been invoked.
// Params filter and quote are for client side filtering.
func (s *Service) ListFromConfig(ctx context.Context, loc Location, filter, quote string) ([]*Entry, error) {
	return s.list(ctx, loc, "", filter, quote, true)
}

func (s *Service) list(ctx context.Context, loc Location, action, filter, quote string, usePanosConfig bool) ([]*Entry, error) {
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

	path, err := loc.XpathWithEntryName(vn, "")
	if err != nil {
		return nil, err
	}

	if usePanosConfig {
		if _, err = s.client.ReadFromConfig(ctx, path, false, normalizer); err != nil {
			return nil, err
		}
	} else {
		cmd := &xmlapi.Config{
			Action: action,
			Xpath:  util.AsXpath(path),
			Target: s.client.GetTarget(),
		}

		if _, _, err = s.client.Communicate(ctx, cmd, true, normalizer); err != nil {
			if err.Error() == "No such node" && action == "show" {
				return nil, nil
			}
			return nil, err
		}
	}

	listing, err := normalizer.Normalize()
	if err != nil || logic == nil {
		return listing, err
	}

	filtered := make([]*Entry, 0, len(listing))
	for _, x := range listing {
		ok, err := logic.Matches(x)
		if err != nil {
			return nil, err
		}
		if ok {
			filtered = append(filtered, x)
		}
	}

	return filtered, nil
}
