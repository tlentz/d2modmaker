package util

// Round32 a float32 to the nearest integer
func Round32(f float32) int {
	if f > 0 {
		return int(f + 0.5)
	}
	return int(f - 0.5)
}

// Round64 a float64 to the nearest integer
func Round64(f float64) int {
	if f > 0 {
		return int(f + 0.5)
	}
	return int(f - 0.5)
}
