package certificate

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
	path, err := loc.XpathWithComponents(vn, util.AsEntryXpath(entry.Name))
	if err != nil {
		return nil, err
	}
	err = s.CreateWithXpath(ctx, util.AsXpath(path[:len(path)-1]), entry)
	if err != nil {
		return nil, err
	}

	return s.ReadWithXpath(ctx, util.AsXpath(path), "get")
}

func (s *Service) CreateWithXpath(ctx context.Context, xpath string, entry *Entry) error {
	return errors.UnsupportedMethodError
}

// Read returns the given config object, using the specified action ("get" or "show").
func (s *Service) Read(ctx context.Context, loc Location, name, action string) (*Entry, error) {
	return s.read(ctx, loc, name, action)
}

func (s *Service) ReadWithXpath(ctx context.Context, xpath string, action string) (*Entry, error) {
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

func (s *Service) read(ctx context.Context, loc Location, value, action string) (*Entry, error) {
	if value == "" {
		return nil, errors.NameNotSpecifiedError
	}
	vn := s.client.Versioning()
	var path []string
	var err error
	path, err = loc.XpathWithComponents(vn, value)
	if err != nil {
		return nil, err
	}

	return s.ReadWithXpath(ctx, util.AsXpath(path), action)
}

func (s *Service) Update(ctx context.Context, loc Location, entry *Entry, name string) (*Entry, error) {
	if entry.Name == "" {
		return nil, errors.NameNotSpecifiedError
	}

	xpath, err := loc.XpathWithComponents(s.client.Versioning(), entry.Name)
	if err != nil {
		return nil, err
	}

	err = s.UpdateWithXpath(ctx, util.AsXpath(xpath), entry, name)
	if err != nil {
		return nil, err
	}

	return s.ReadWithXpath(ctx, util.AsXpath(xpath), "get")
}
func (s *Service) UpdateWithXpath(ctx context.Context, xpath string, entry *Entry, name string) error {
	return errors.UnsupportedMethodError
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
		path, err = loc.XpathWithComponents(vn, util.AsEntryXpath(value))
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
	return s.list(ctx, loc, action, filter, quote)
}

func (s *Service) list(ctx context.Context, loc Location, action, filter, quote string) ([]*Entry, error) {
	xpath, err := loc.XpathWithComponents(s.client.Versioning(), util.AsEntryXpath(""))
	if err != nil {
		return nil, err
	}

	return s.ListWithXpath(ctx, util.AsXpath(xpath), action, filter, quote)
}

func (s *Service) ListWithXpath(ctx context.Context, xpath string, action, filter, quote string) ([]*Entry, error) {
	var logic *filtering.Group
	if filter != "" {
		var err error
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

	cmd := &xmlapi.Config{
		Action: action,
		Xpath:  util.AsXpath(xpath),
		Target: s.client.GetTarget(),
	}

	if _, _, err = s.client.Communicate(ctx, cmd, true, normalizer); err != nil {
		if err.Error() == "No such node" && action == "show" {
			return nil, nil
		}
		return nil, err
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

func (s *Service) filterEntriesByLocation(location Location, entries []*Entry) []*Entry {
	filter := location.LocationFilter()
	if filter == nil {
		return entries
	}

	getLocAttribute := func(entry *Entry) *string {
		for _, elt := range entry.GetMiscAttributes() {
			if elt.Name.Local == "loc" {
				return &elt.Value
			}
		}
		return nil
	}

	var filtered []*Entry
	for _, elt := range entries {
		location := getLocAttribute(elt)
		if location == nil || *location == *filter {
			filtered = append(filtered, elt)
		}
	}

	return filtered
}
