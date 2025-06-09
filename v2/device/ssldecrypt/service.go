package ssldecrypt

import (
	"context"
	"fmt"

	"github.com/PaloAltoNetworks/pango/v2/errors"
	"github.com/PaloAltoNetworks/pango/v2/util"
	"github.com/PaloAltoNetworks/pango/v2/xmlapi"
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
	path, err := loc.XpathWithComponents(vn)
	if err != nil {
		return nil, err
	}
	err = s.CreateWithXpath(ctx, util.AsXpath(path[:len(path)-1]), config)
	if err != nil {
		return nil, err
	}

	return s.ReadWithXpath(ctx, util.AsXpath(path), "get")
}

func (s *Service) CreateWithXpath(ctx context.Context, xpath string, config *Config) error {
	vn := s.client.Versioning()
	specifier, _, err := Versioning(vn)
	if err != nil {
		return err
	}
	createSpec, err := specifier(config)
	if err != nil {
		return err
	}

	cmd := &xmlapi.Config{
		Action:  "set",
		Xpath:   xpath,
		Element: createSpec,
		Target:  s.client.GetTarget(),
	}

	if _, _, err = s.client.Communicate(ctx, cmd, false, nil); err != nil {
		return err
	}

	return nil
}

// Read returns the given config object, using the specified action ("get" or "show").
func (s *Service) Read(ctx context.Context, loc Location, action string) (*Config, error) {
	return s.read(ctx, loc, action)
}

func (s *Service) ReadWithXpath(ctx context.Context, xpath string, action string) (*Config, error) {
	vn := s.client.Versioning()
	_, normalizer, err := Versioning(vn)
	if err != nil {
		return nil, err
	}

	cmd := &xmlapi.Config{
		Action: action,
		Xpath:  xpath,
		Target: s.client.GetTarget(),
	}

	if _, _, err = s.client.Communicate(ctx, cmd, true, normalizer); err != nil {
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

	return list[0], nil
}

func (s *Service) read(ctx context.Context, loc Location, action string) (*Config, error) {
	vn := s.client.Versioning()
	path, err := loc.XpathWithComponents(vn)
	if err != nil {
		return nil, err
	}

	return s.ReadWithXpath(ctx, util.AsXpath(path), action)
}

func (s *Service) Update(ctx context.Context, loc Location, entry *Config, name string) (*Config, error) {
	xpath, err := loc.XpathWithComponents(s.client.Versioning())
	if err != nil {
		return nil, err
	}

	err = s.UpdateWithXpath(ctx, util.AsXpath(xpath), entry)
	if err != nil {
		return nil, err
	}

	return s.ReadWithXpath(ctx, util.AsXpath(xpath), "get")
}
func (s *Service) UpdateWithXpath(ctx context.Context, xpath string, entry *Config) error {
	vn := s.client.Versioning()
	updates := xmlapi.NewMultiConfig(2)
	specifier, _, err := Versioning(vn)
	if err != nil {
		return err
	}

	var old *Config
	old, err = s.ReadWithXpath(ctx, xpath, "get")
	if err != nil {
		return err
	}

	if SpecMatches(entry, old) {
		return nil
	}

	updateSpec, err := specifier(entry)
	if err != nil {
		return err
	}

	updates.Add(&xmlapi.Config{
		Action:  "edit",
		Xpath:   util.AsXpath(xpath),
		Element: updateSpec,
		Target:  s.client.GetTarget(),
	})

	if len(updates.Operations) != 0 {
		if _, _, _, err = s.client.MultiConfig(ctx, updates, false, nil); err != nil {
			return err
		}
	}

	return nil
}

// Delete deletes the given item.
func (s *Service) Delete(ctx context.Context, loc Location, config *Config) error {
	return s.delete(ctx, loc, config)
}
func (s *Service) delete(ctx context.Context, loc Location, config *Config) error {

	vn := s.client.Versioning()
	path, err := loc.XpathWithComponents(vn)
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
