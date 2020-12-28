package users

import (
	"errors"
	"lion/pkg/utils"
	"lion/pkg/validator"

	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrInvalidType occured when user input invalid type
	ErrInvalidType = errors.New("invalid type")

	stringValidator = validator.NewStringValidator()
	allowedUpdate   = []string{"name"}
)

// User entity
type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name"`
}

// ValidateInsert list of insert validation rule
func (u *User) ValidateInsert() error {
	return validator.MergeError(
		validator.ValidateString(
			"email",
			u.Email,
			stringValidator.Required,
			stringValidator.Email,
		),
		validator.ValidateString(
			"password",
			u.Password,
			stringValidator.Required,
			stringValidator.MinLength(5),
		),
		validator.ValidateString(
			"name",
			u.Name,
			stringValidator.Required,
		),
	)
}

// ValidateAuth list of auth validation rule
func (u *User) ValidateAuth() error {
	return validator.MergeError(
		validator.ValidateString(
			"email",
			u.Email,
			stringValidator.Required,
			stringValidator.Email,
		),
		validator.ValidateString(
			"password",
			u.Password,
			stringValidator.Required,
			stringValidator.MinLength(5),
		),
	)
}

// SanitizeUpdate remove all not allowed key
func (u *User) SanitizeUpdate(updateData genericJSON) {
	for k := range updateData {
		if !utils.StringInSlice(k, allowedUpdate) {
			delete(updateData, k)
		}
	}
}

// EncryptPassword encrypt plain password
func (u *User) EncryptPassword() error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPass)
	return nil
}

// ComparePassword compare stored encrypted password with plain password
func (u *User) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}
