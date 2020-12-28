package validator

import (
	"errors"
	"log"
	"testing"
)

func TestCustomErrorAndRequired(t *testing.T) {
	customRequiredError := errors.New("you should not pass")
	sv := NewStringValidator()
	// check if validator has default error value
	if sv.RequiredError != ErrRequired {
		t.Errorf("new created string validator should has default value of error, instead got %s ", sv.RequiredError)
	}

	// check if custom error is applied to validator
	sv.RequiredError = customRequiredError
	requiredData := ""
	if err := sv.Required(requiredData); err != customRequiredError {
		t.Errorf("should return custom error, instead got %s ", err)
	}

	requiredData = "hola"
	if err := sv.Required(requiredData); err != nil {
		t.Errorf("should not return error since the required data is filled, instead got error %s ", err)
	}
}

func TestEmailValidation(t *testing.T) {
	sv := NewStringValidator()
	// test with invalid email
	email := "email"
	if err := sv.Email(email); err != sv.EmailValidationError {
		t.Errorf("should return error since the email has invalid value, instead got %s", err)
	}

	// test with valid value
	email = "email@gmail.com"
	if err := sv.Email(email); err != nil {
		t.Errorf("should not return error since the email has valid value, instead got error %s", err)
	}
}

func TestMinLength(t *testing.T) {
	sv := NewStringValidator()

	// test with invalid value
	data := "1"
	if err := sv.MinLength(2)(data).Error(); err != sv.MinLengthError(2).Error() {
		t.Errorf("should return %s since the data has invalid value, instead got %s", sv.MinLengthError(2), err)
	}

	// test with valid value
	data = "test"
	if err := sv.MinLength(2)(data); err != nil {
		t.Errorf("should not return error since the data has valid value, instead got %s", err)
	}
}

func TestValidateStringFN(t *testing.T) {
	key := "name"
	value := ""
	sv := NewStringValidator()
	// test with error value
	err := ValidateString(
		key,
		value,
		sv.Required,
	)

	errors, ok := err.(Errors)
	if !ok {
		t.Errorf("error should return with Errors type")
	}

	val, ok := errors[key]
	if !ok {
		t.Errorf("errors type should have a %s as a key", key)
	}

	if val != sv.RequiredError {
		t.Errorf("val should return requiredError, instead got %s", val)
	}

	key = "address"
	value = "a"
	err2 := ValidateString(
		key,
		value,
		sv.MinLength(3),
	)

	merged := MergeError(err, err2)
	if len(merged.(Errors)) != 2 {
		t.Errorf("merged shoud have 2 errors now, instead got %d", len(merged.(Errors)))
	}

	log.Println(merged.Error())
}
