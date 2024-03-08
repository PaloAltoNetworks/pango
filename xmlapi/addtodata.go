package xmlapi

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

type elementer interface {
	Element() any
}

func addToData(key string, i any, attemptMarshal bool, data *url.Values) error {
	if i == nil {
		return nil
	}

	val, err := asString(i, attemptMarshal)
	if err != nil {
		return err
	}

	data.Set(key, val)
	return nil
}

func asString(i any, attemptMarshal bool) (string, error) {
	if a, ok := i.(fmt.Stringer); ok {
		return a.String(), nil
	}

	if b, ok := i.(elementer); ok {
		i = b.Element()
	}

	switch val := i.(type) {
	case nil:
		return "", fmt.Errorf("nil encountered")
	case string:
		return val, nil
	default:
		if !attemptMarshal {
			return "", fmt.Errorf("value must be string or fmt.Stringer")
		}

		rb, err := xml.Marshal(val)
		if err != nil {
			return "", err
		}
		return string(rb), nil
	}
}
