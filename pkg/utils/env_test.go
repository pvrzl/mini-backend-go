package utils

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	const (
		DEFAULT = "DEFAULT"
		UPDATED = "UPDATED_ENV"
	)

	s := GetEnv("TEST_ENV", DEFAULT)
	if s != DEFAULT {
		t.Errorf("get env should return it's default value, instead got %s ", s)
	}

	os.Setenv("TEST_ENV", UPDATED)
	s = GetEnv("TEST_ENV", DEFAULT)
	if s != UPDATED {
		t.Errorf("get env should return new value, instead got %s ", s)
	}
}
