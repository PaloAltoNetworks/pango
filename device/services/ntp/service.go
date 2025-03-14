package ntp

import (
	"context"
	"fmt"

	"github.com/PaloAltoNetworks/pango/errors"
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
func (s *Service) Create(ctx context.Context, loc Location, config *Config) (*Config, error) {

	vn := s.client.Versioning()

	specifier, _, err := Versioning(vn)
	if err != nil {
		return nil, err
	}
	path, err := loc.Xpath(vn)
	if err != nil {
		return nil, err
	}
	createSpec, err := specifier(config)
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
	return s.Read(ctx, loc, "get")
}

// Read returns the given config object, using the specified action ("get" or "show").
func (s *Service) Read(ctx context.Context, loc Location, action string) (*Config, error) {
	return s.read(ctx, loc, action, false)
}

// ReadFromConfig returns the given config object from the loaded XML config.
// Requires that client.LoadPanosConfig() has been invoked.
func (s *Service) ReadFromConfig(ctx context.Context, loc Location) (*Config, error) {
	return s.read(ctx, loc, "", true)
}

func (s *Service) read(ctx context.Context, loc Location, action string, usePanosConfig bool) (*Config, error) {
	vn := s.client.Versioning()
	_, normalizer, err := Versioning(vn)
	if err != nil {
		return nil, err
	}
	path, err := loc.Xpath(vn)
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

func (s *Service) Update(ctx context.Context, loc Location, entry *Config) (*Config, error) {

	vn := s.client.Versioning()
	updates := xmlapi.NewMultiConfig(2)
	specifier, _, err := Versioning(vn)
	if err != nil {
		return nil, err
	}
	var old *Config
	old, err = s.Read(ctx, loc, "get")
	if err != nil {
		return nil, err
	} else if old == nil {
		return nil, fmt.Errorf("previous object doesn't exist for update")
	}
	if !SpecMatches(entry, old) {
		path, err := loc.Xpath(vn)
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
	return s.Read(ctx, loc, "get")
}

// Delete deletes the given item.
func (s *Service) Delete(ctx context.Context, loc Location, config *Config) error {
	return s.delete(ctx, loc, config)
}
func (s *Service) delete(ctx context.Context, loc Location, config *Config) error {

	vn := s.client.Versioning()
	path, err := loc.Xpath(vn)
	if err != nil {
		return err
	}
	deleteSuffixes := []string{}

	for _, suffix := range deleteSuffixes {
		cmd := &xmlapi.Config{
			Action: "delete",
			Xpath:  util.AsXpath(append(path, suffix)),
			Target: s.client.GetTarget(),
		}

		_, _, err = s.client.Communicate(ctx, cmd, false, nil)

		if err != nil {
			return err
		}
	}
	return nil
}
