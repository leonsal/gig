package gig

func ClampFloat32(v, min, max float32) float32 {

	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
