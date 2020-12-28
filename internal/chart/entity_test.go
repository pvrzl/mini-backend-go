package chart

import (
	"testing"
)

func TestValidateInsert(t *testing.T) {
	chart := new(Chart)
	// test with invalid value
	err := chart.ValidateInsert()
	if err == nil {
		t.Error("validate insert should return error when called with invalid value")
	}

	// test with valid value
	chart.Name = "blabla"
	err = chart.ValidateInsert()
	if err != nil {
		t.Errorf("validate insert should not return error when called with valid value, instead got %s", err)
	}
}
