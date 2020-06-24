package util

// MaxInt Returns the max between 2 ints
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxFloat Returns the max between 2 float64
func MaxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// MinInt Returns the min between 2 ints
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MinFloat Returns the min between 2 float64
func MinFloat(a, b float64) float64 {
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
