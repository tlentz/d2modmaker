package util

// Absf32 returns |x| (Absolute value of x)
func Absf32(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}

// AbsInt returns |x| (Absolute value of x)
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
