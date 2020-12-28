package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
)

type (
	// Errors is key value error type
	Errors map[string]error
)

var (
	// ErrRequired occured if required variable does not receive anything during validation
	ErrRequired = errors.New("is required")
	// ErrEmailInvalid occured if the string is not a valid email address
	ErrEmailInvalid = errors.New("invalid email address")
)

// MergeError combine every errors into key value error type
func MergeError(errs ...error) error {
	m := make(Errors)

	for _, err := range errs {
		if err != nil {
			val, ok := err.(Errors)
			if ok {
				for k, v := range val {
					m[k] = v
				}
			}

		}

	}

	if len(m) == 0 {
		return nil
	}

	return m
}

func (es Errors) Error() string {
	if len(es) == 0 {
		return ""
	}

	keys := []string{}
	for key := range es {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var s strings.Builder
	for i, key := range keys {
		if i > 0 {
			s.WriteString("; ")
		}
		if errs, ok := es[key].(Errors); ok {
			fmt.Fprintf(&s, "%v: (%v)", key, errs)
		} else {
			fmt.Fprintf(&s, "%v: %v", key, es[key].Error())
		}
	}
	s.WriteString(".")
	return s.String()
}

// MarshalJSON convert errors into a json
func (es Errors) MarshalJSON() ([]byte, error) {
	errs := map[string]interface{}{}
	for key, err := range es {
		if ms, ok := err.(json.Marshaler); ok {
			errs[key] = ms
		} else {
			errs[key] = err.Error()
		}
	}
	return json.Marshal(errs)
}
