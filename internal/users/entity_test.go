package users

import (
	"testing"
)

func TestValidateInsert(t *testing.T) {
	user := new(User)
	// test with invalid value
	err := user.ValidateInsert()
	if err == nil {
		t.Error("validate insert should return error when called with invalid value")
	}

	// test with valid value
	user.Email = "test@test.com"
	user.Password = "blabla"
	user.Name = "blabla"
	err = user.ValidateInsert()
	if err != nil {
		t.Errorf("validate insert should not return error when called with valid value, instead got %s", err)
	}
}

func TestPassword(t *testing.T) {
	user := new(User)
	plainPassword := "xxxx"
	user.Password = plainPassword
	err := user.EncryptPassword()
	if err != nil {
		t.Error("should not return error when encrypting password with correct value")
	}

	if user.Password == plainPassword {
		t.Error("password should be encrypted right now, so the value must not be same with plain password")
	}

	err = user.ComparePassword("test")
	if err == nil {
		t.Error("should return error when comparing encrypting password with incorrect plain password")
	}

	err = user.ComparePassword(plainPassword)
	if err != nil {
		t.Error("should not return error when comparing encrypting password with correct plain password")
	}
}
