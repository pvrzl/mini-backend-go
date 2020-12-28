package utils

import "strconv"

// StringInSlice is a func to check wether a string is in slice or nots
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// StringToIntWithDefault convert string to int with default
func StringToIntWithDefault(s string, dflt int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return dflt
	}
	return i
}
