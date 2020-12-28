package utils

// StringInSlice is a func to check wether a string is in slice or nots
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
