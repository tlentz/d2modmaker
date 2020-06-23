package util

// Max Returns the max between 2 ints
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min Returns the min between 2 ints
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ContainsString checks if a slice contains a given string
func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
