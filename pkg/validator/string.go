package validator

import (
	"fmt"
	"regexp"
)

type stringValidator struct {
	RequiredError        error
	MinLengthError       func(int) error
	EmailValidationError error
}

type stringValidatorType func(string) error

// NewStringValidator create a new stringValidator
func NewStringValidator() stringValidator {
	return stringValidator{
		RequiredError: ErrRequired,
		MinLengthError: func(l int) error {
			return fmt.Errorf("min length is %d", l)
		},
		EmailValidationError: ErrEmailInvalid,
	}
}

// ValidateString validate every registered validator
func ValidateString(key string, value string, validators ...stringValidatorType) error {
	errors := make(Errors)
	for _, validator := range validators {
		if err := validator(value); err != nil {
			errors[key] = err
			return errors
		}
	}

	return nil
}

func (cfg stringValidator) Required(s string) error {
	if s == "" {
		return cfg.RequiredError
	}

	return nil
}

func (cfg stringValidator) MinLength(l int) func(string) error {
	return func(s string) error {
		if s == "" {
			return nil
		}
		if len(s) < l {
			return cfg.MinLengthError(l)
		}

		return nil
	}
}

func (cfg stringValidator) Email(s string) error {
	if s == "" {
		return nil
	}
	match, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, s)
	if !match {
		return cfg.EmailValidationError
	}

	return nil
}
