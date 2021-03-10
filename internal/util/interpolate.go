package util

// Interpolate Calculate interpolation to a single int given a starting range (lmin, lmax) and a target range (smin,smax)
// Calculates avg(pin & pmax) and returns single value
// TODO: Change to return a range.
func Interpolate(pmin int, pmax int, lmin int, lmax int, smin int, smax int) int {
	var avg float32 = 0.0
	if pmin+pmax != 0 {
		avg = float32((pmin + pmax)) / 2.0
	}
	if lmin == lmax {
		return smax
	}
	return Round32((float32(avg)-float32(lmin))*(1.0/(float32(lmax-lmin)))*float32((smax-smin))) + smin
}

// Interpolatef32 Interpolate a value from a given starting range (lmin, lmax) and a target range (smin,smax)
func Interpolatef32(p float32, lmin int, lmax int, smin int, smax int) float32 {
	if lmin == lmax {
		return float32(smax)
	}
	//return (float32(avg)-float32(lmin))*(1.0/(float32(lmax-lmin)))*float32((smax-smin)) + float32(smin)
	return (p-float32(lmin))*(1.0/(float32(lmax-lmin)))*float32((smax-smin)) + float32(smin)
}
