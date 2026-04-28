package file_type

import (
	"fmt"
	"strings"

	"github.com/PaloAltoNetworks/pango/errors"
	"github.com/PaloAltoNetworks/pango/version"
)

type Location struct {
	Predefined *PredefinedLocation `json:"predefined"`
}

type PredefinedLocation struct {
}

func NewPredefinedLocation() *Location {
	return &Location{Predefined: &PredefinedLocation{},
	}
}

func (o Location) IsValid() error {
	count := 0

	switch {
	case o.Predefined != nil:
		count++
	}

	if count == 0 {
		return fmt.Errorf("no path specified")
	}

	if count > 1 {
		return fmt.Errorf("multiple paths specified: only one should be specified")
	}

	return nil
}

func (o Location) LocationFilter() *string {

	return nil
}

func (o Location) XpathPrefix(vn version.Number) ([]string, error) {

	var ans []string

	switch {
	case o.Predefined != nil:
		ans = []string{
			"config",
			"predefined",
			"dlp-file-property",
		}
	default:
		return nil, errors.NoLocationSpecifiedError
	}

	return ans, nil
}

func (o Location) XpathWithComponents(vn version.Number, components ...string) ([]string, error) {
	if len(components) != 1 {
		return nil, fmt.Errorf("invalid number of arguments for XpathWithComponents() call")
	}

	{
		component := components[0]
		if component != "entry" {
			if !strings.HasPrefix(component, "entry[@name=\"]") && !strings.HasPrefix(component, "entry[@name='") {
				return nil, errors.NewInvalidXpathComponentError(fmt.Sprintf("Name must be formatted as entry: %s", component))
			}

			if !strings.HasSuffix(component, "\"]") && !strings.HasSuffix(component, "']") {
				return nil, errors.NewInvalidXpathComponentError(fmt.Sprintf("Name must be formatted as entry: %s", component))
			}
		}
	}

	ans, err := o.XpathPrefix(vn)
	if err != nil {
		return nil, err
	}

	ans = append(ans, "file-type")
	ans = append(ans, components[0])

	return ans, nil
}
