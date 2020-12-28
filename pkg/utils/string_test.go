package utils

import "testing"

func TestStringInSlice(t *testing.T) {
	if !StringInSlice("a", []string{"a", "b", "c", "d"}) {
		t.Error("should return true since list has a")
	}

	if StringInSlice("a", []string{"b", "c", "d"}) {
		t.Error("should return false since list doesnt has a")
	}
}
