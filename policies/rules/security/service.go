package security

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/PaloAltoNetworks/pango/audit"
	"github.com/PaloAltoNetworks/pango/errors"
	"github.com/PaloAltoNetworks/pango/filtering"
	"github.com/PaloAltoNetworks/pango/movement"
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
	return s.read(ctx, loc, name, action, true, false)
}

// ReadById returns the given config object with specified ID, using the specified action ("get" or "show").
func (s *Service) ReadById(ctx context.Context, loc Location, uuid, action string) (*Entry, error) {
	return s.read(ctx, loc, uuid, action, false, false)
}

// ReadFromConfig returns the given config object from the loaded XML config.
// Requires that client.LoadPanosConfig() has been invoked.
func (s *Service) ReadFromConfig(ctx context.Context, loc Location, name string) (*Entry, error) {
	return s.read(ctx, loc, name, "", true, true)
}

// ReadFromConfigById returns the given config object with specified ID from the loaded XML config.
// Requires that client.LoadPanosConfig() has been invoked.
func (s *Service) ReadFromConfigById(ctx context.Context, loc Location, uuid string) (*Entry, error) {
	return s.read(ctx, loc, uuid, "", false, true)
}

func (s *Service) read(ctx context.Context, loc Location, value, action string, byName, usePanosConfig bool) (*Entry, error) {
	if byName && value == "" {
		return nil, errors.NameNotSpecifiedError
	}
	if !byName && value == "" {
		return nil, errors.UuidNotSpecifiedError
	}
	vn := s.client.Versioning()
	_, normalizer, err := Versioning(vn)
	if err != nil {
		return nil, err
	}
	var path []string
	if byName {
		path, err = loc.XpathWithEntryName(vn, value)
	} else {
		path, err = loc.XpathWithUuid(vn, value)
	}
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
	return s.update(ctx, loc, entry, name, true)
}

// UpdateById updates the given config object, then returns the result.
func (s *Service) UpdateById(ctx context.Context, loc Location, entry *Entry, uuid string) (*Entry, error) {
	return s.update(ctx, loc, entry, uuid, false)
}
func (s *Service) update(ctx context.Context, loc Location, entry *Entry, value string, byName bool) (*Entry, error) {
	if byName && entry.Name == "" {
		return nil, errors.NameNotSpecifiedError
	}
	if !byName && value == "" {
		return nil, errors.UuidNotSpecifiedError
	}

	vn := s.client.Versioning()
	updates := xmlapi.NewMultiConfig(2)
	specifier, _, err := Versioning(vn)
	if err != nil {
		return nil, err
	}
	var old *Entry
	if byName {
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
	} else {
		old, err = s.ReadById(ctx, loc, value, "get")
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
	if byName {
		return s.Read(ctx, loc, entry.Name, "get")
	} else {
		return s.ReadById(ctx, loc, value, "get")
	}
}

// Delete deletes the given item.
func (s *Service) Delete(ctx context.Context, loc Location, name ...string) error {
	return s.delete(ctx, loc, name, true)
}

// DeleteById deletes the given item with specified ID.
func (s *Service) DeleteById(ctx context.Context, loc Location, uuid ...string) error {
	return s.delete(ctx, loc, uuid, false)
}
func (s *Service) delete(ctx context.Context, loc Location, values []string, byName bool) error {
	for _, value := range values {
		if byName && value == "" {
			return errors.NameNotSpecifiedError
		}
		if !byName && value == "" {
			return errors.UuidNotSpecifiedError
		}
	}

	vn := s.client.Versioning()
	var err error
	deletes := xmlapi.NewMultiConfig(len(values))
	for _, value := range values {
		var path []string
		if byName {
			path, err = loc.XpathWithEntryName(vn, value)
		} else {
			path, err = loc.XpathWithUuid(vn, value)
		}
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

// MoveGroup arranges the given rules in the order specified.
// Any rule with a UUID specified is ignored.
// Only the rule names are considered for the purposes of the rule placement.
func (s *Service) MoveGroup(ctx context.Context, loc Location, position movement.Position, entries []*Entry, batchSize int) error {
	if len(entries) == 0 {
		return nil
	}

	existing, err := s.List(ctx, loc, "get", "", "")
	if err != nil {
		return err
	} else if len(existing) == 0 {
		return fmt.Errorf("no rules present")
	}

	movements, err := movement.MoveGroup(position, entries, existing)
	if err != nil {
		return err
	}

	updates := xmlapi.NewChunkedMultiConfig(len(movements), batchSize)

	for _, elt := range movements {
		path, err := loc.XpathWithEntryName(s.client.Versioning(), elt.Movable.EntryName())
		if err != nil {
			return err
		}

		switch elt.Where {
		case movement.ActionWhereFirst, movement.ActionWhereLast:
			updates.Add(&xmlapi.Config{
				Action:      "move",
				Xpath:       util.AsXpath(path),
				Where:       string(elt.Where),
				Destination: string(elt.Where),
				Target:      s.client.GetTarget(),
			})
		case movement.ActionWhereBefore, movement.ActionWhereAfter:
			updates.Add(&xmlapi.Config{
				Action:      "move",
				Xpath:       util.AsXpath(path),
				Where:       string(elt.Where),
				Destination: elt.Destination.EntryName(),
				Target:      s.client.GetTarget(),
			})
		}

	}

	if len(updates.Operations) > 0 {
		_, _, _, err = s.client.MultiConfig(ctx, updates, false, nil)
		return err
	}

	return nil
}

// HITCOUNT returns the hit count for the given rule.
func (s *Service) HitCount(ctx context.Context, loc Location, rules ...string) ([]util.HitCount, error) {
	switch {
	case loc.Vsys != nil:
		cmd := &xmlapi.Op{
			Command: util.NewHitCountRequest(RuleType, loc.Vsys.Vsys, rules),
			Target:  s.client.GetTarget(),
		}
		var resp util.HitCountResponse

		if _, _, err := s.client.Communicate(ctx, cmd, false, &resp); err != nil {
			return nil, err
		}

		return resp.Results, nil
	}

	return nil, fmt.Errorf("unsupported location")
}

// SetAuditComment sets the given audit comment for the given rule.
func (s *Service) SetAuditComment(ctx context.Context, loc Location, name, comment string) error {
	if name == "" {
		return errors.NameNotSpecifiedError
	}

	vn := s.client.Versioning()

	path, err := loc.XpathWithEntryName(vn, name)
	if err != nil {
		return err
	}

	cmd := &xmlapi.Op{
		Command: audit.SetComment{
			Xpath:   util.AsXpath(path),
			Comment: comment,
		},
		Target: s.client.GetTarget(),
	}

	_, _, err = s.client.Communicate(ctx, cmd, false, nil)
	return err
}

// CurrentAuditComment gets any current uncommitted audit comment for the given rule.
func (s *Service) CurrentAuditComment(ctx context.Context, loc Location, name string) (string, error) {
	if name == "" {
		return "", errors.NameNotSpecifiedError
	}

	vn := s.client.Versioning()

	path, err := loc.XpathWithEntryName(vn, name)
	if err != nil {
		return "", err
	}

	cmd := &xmlapi.Op{
		Command: audit.GetComment{
			Xpath: util.AsXpath(path),
		},
		Target: s.client.GetTarget(),
	}

	var resp audit.UncommittedComment
	if _, _, err = s.client.Communicate(ctx, cmd, false, &resp); err != nil {
		return "", err
	}

	return resp.Comment, nil
}

// AuditCommentHistory returns a chunk of historical audit comment logs.
func (s *Service) AuditCommentHistory(ctx context.Context, loc Location, name, direction string, nlogs, skip int) ([]audit.Comment, error) {
	if name == "" {
		return nil, errors.NameNotSpecifiedError
	}

	var err error
	var base, vsysDg string
	switch {
	case loc.Vsys != nil:
		vsysDg = loc.Vsys.Vsys
		base = "rulebase"
	case loc.Shared != nil:
		vsysDg = "shared"
		base = loc.Shared.Rulebase
	case loc.DeviceGroup != nil:
		vsysDg = loc.DeviceGroup.DeviceGroup
		base = loc.DeviceGroup.Rulebase
	}

	if vsysDg == "" || base == "" {
		return nil, fmt.Errorf("unsupported location")
	}

	query := strings.Join([]string{
		"(subtype eq audit-comment)",
		fmt.Sprintf("(path contains '\\'%s\\'')", name),
		fmt.Sprintf("(path contains '%s')", RuleType),
		fmt.Sprintf("(path contains %s)", base),
		fmt.Sprintf("(path contains '\\'%s\\'')", vsysDg),
	}, " and ")
	extras := url.Values{}
	extras.Set("uniq", "yes")

	cmd := &xmlapi.Log{
		LogType:   "config",
		Query:     query,
		Direction: direction,
		Nlogs:     nlogs,
		Skip:      skip,
		Extras:    extras,
	}

	var job util.JobResponse
	if _, _, err = s.client.Communicate(ctx, cmd, false, &job); err != nil {
		return nil, err
	}

	var resp audit.CommentHistory
	if _, err = s.client.WaitForLogs(ctx, job.Id, 1*time.Second, &resp); err != nil {
		return nil, err
	}

	if len(resp.Comments) != 0 {
		if clock, err := s.client.Clock(ctx); err == nil {
			for i := range resp.Comments {
				resp.Comments[i].SetTime(clock)
			}
		}
	}

	return resp.Comments, nil
}
