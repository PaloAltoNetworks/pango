package ethernet

import (
	"context"
	"fmt"
	"sync"

	"github.com/PaloAltoNetworks/pango/errors"
	"github.com/PaloAltoNetworks/pango/filtering"
	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/xmlapi"
)

var (
	importsMutexMap     = make(map[string]*sync.Mutex)
	importsMutexMapLock = sync.Mutex{}
)

func (s *Service) getImportMutex(xpath string) *sync.Mutex {
	importsMutexMapLock.Lock()
	defer importsMutexMapLock.Unlock()

	var importMutex *sync.Mutex
	var ok bool
	importMutex, ok = importsMutexMap[xpath]
	if !ok {
		importMutex = &sync.Mutex{}
		importsMutexMap[xpath] = importMutex
	}

	return importMutex
}

type Service struct {
	client util.PangoClient
}

func NewService(client util.PangoClient) *Service {
	return &Service{
		client: client,
	}
}

// Create adds new item, then returns the result.
func (s *Service) Create(ctx context.Context, loc Location, importLocations []ImportLocation, entry *Entry) (*Entry, error) {
	if entry.Name == "" {
		return nil, errors.NameNotSpecifiedError
	}

	vn := s.client.Versioning()
	path, err := loc.XpathWithEntryName(vn, entry.Name)
	if err != nil {
		return nil, err
	}
	err = s.CreateWithXpath(ctx, util.AsXpath(path[:len(path)-1]), entry)
	if err != nil {
		return nil, err
	}
	err = s.ImportToLocations(ctx, loc, importLocations, entry.Name)
	if err != nil {
		return nil, err
	}

	return s.ReadWithXpath(ctx, util.AsXpath(path), "get")
}

func (s *Service) CreateWithXpath(ctx context.Context, xpath string, entry *Entry) error {
	vn := s.client.Versioning()
	specifier, _, err := Versioning(vn)
	if err != nil {
		return err
	}
	createSpec, err := specifier(entry)
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

func (s *Service) ImportToLocations(ctx context.Context, loc Location, importLocations []ImportLocation, entryName string) error {
	vn := s.client.Versioning()

	importToLocation := func(il ImportLocation) error {
		xpath, err := il.XpathForLocation(vn, loc)
		if err != nil {
			return err
		}

		mutex := s.getImportMutex(util.AsXpath(xpath))
		mutex.Lock()
		defer mutex.Unlock()

		cmd := &xmlapi.Config{
			Action: "get",
			Xpath:  util.AsXpath(xpath),
		}

		bytes, _, err := s.client.Communicate(ctx, cmd, false, nil)
		if err != nil && !errors.IsObjectNotFound(err) {
			return err
		}

		existing, err := il.UnmarshalPangoXML(bytes)
		if err != nil {
			return err
		}

		for _, elt := range existing {
			if elt == entryName {
				return nil
			}
		}

		existing = append(existing, entryName)

		element, err := il.MarshalPangoXML(existing)
		if err != nil {
			return err
		}

		cmd = &xmlapi.Config{
			Action:  "set",
			Xpath:   util.AsXpath(xpath[:len(xpath)-1]),
			Element: element,
		}

		_, _, err = s.client.Communicate(ctx, cmd, false, nil)
		if err != nil {
			return err
		}

		return err
	}

	for _, elt := range importLocations {
		err := importToLocation(elt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) UnimportFromLocations(ctx context.Context, loc Location, importLocations []ImportLocation, values []string) error {
	vn := s.client.Versioning()
	valuesByName := make(map[string]bool)
	for _, elt := range values {
		valuesByName[elt] = true
	}

	unimportFromLocation := func(il ImportLocation) error {
		xpath, err := il.XpathForLocation(vn, loc)
		if err != nil {
			return err
		}

		mutex := s.getImportMutex(util.AsXpath(xpath))
		mutex.Lock()
		defer mutex.Unlock()

		cmd := &xmlapi.Config{
			Action: "get",
			Xpath:  util.AsXpath(xpath),
		}

		bytes, _, err := s.client.Communicate(ctx, cmd, false, nil)
		if err != nil && !errors.IsObjectNotFound(err) {
			return err
		}

		existing, err := il.UnmarshalPangoXML(bytes)
		if err != nil {
			return err
		}

		var filtered []string
		for _, elt := range existing {
			if _, found := valuesByName[elt]; !found {
				filtered = append(filtered, elt)
			}
		}

		element, err := il.MarshalPangoXML(filtered)
		if err != nil {
			return err
		}

		cmd = &xmlapi.Config{
			Action:  "edit",
			Xpath:   util.AsXpath(xpath),
			Element: element,
		}

		_, _, err = s.client.Communicate(ctx, cmd, false, nil)
		if err != nil {
			return err
		}

		return err
	}

	for _, elt := range importLocations {
		err := unimportFromLocation(elt)
		if err != nil {
			return err
		}
	}

	return nil
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
	path, err = loc.XpathWithEntryName(vn, value)
	if err != nil {
		return nil, err
	}

	return s.ReadWithXpath(ctx, util.AsXpath(path), action)
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
func (s *Service) Delete(ctx context.Context, loc Location, importLocations []ImportLocation, name ...string) error {
	return s.delete(ctx, loc, importLocations, name)
}
func (s *Service) delete(ctx context.Context, loc Location, importLocations []ImportLocation, values []string) error {
	for _, value := range values {
		if value == "" {
			return errors.NameNotSpecifiedError
		}
	}

	vn := s.client.Versioning()
	var err error
	deletes := xmlapi.NewMultiConfig(len(values))
	err = s.UnimportFromLocations(ctx, loc, importLocations, values)
	if err != nil {
		return err
	}
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
