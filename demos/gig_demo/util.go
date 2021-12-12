package main

import "math"

// getSineF32 generates and returns a slice with float32 samples of sine function
// of the specified amplitude, offset and normalized frequency (frequency/sample rate)
func genSineF32(amplitude, offset, nf float64, count int) []float32 {

	const period = 2 * math.Pi
	phase := 0.0
	delta := period * nf
	out := make([]float32, count)
	for i := 0; i < count; i++ {
		out[i] = float32(amplitude*math.Sin(phase) + offset)
		phase += delta
		for phase > period {
			phase -= period
		}
	}
	return out
}

// getSineF64 generates and returns a slice with float64 samples of sine function
// of the specified amplitude, offset and normalized frequency (frequency/sample rate)
func genSineF64(amplitude, offset, nf float64, count int) []float64 {

	const period = 2 * math.Pi
	phase := 0.0
	delta := period * nf
	out := make([]float64, count)
	for i := 0; i < count; i++ {
		out[i] = amplitude*math.Sin(phase) + offset
		phase += delta
		for phase > period {
			phase -= period
		}
	}
	return out
}

// getSeqF32 generates and returns a slice fo a sequence of float32 values
// with the specified normalized frequency
func genSeqF32(nf float64, count int) []float32 {

	delta := 2 * math.Pi * nf
	out := make([]float32, count)
	for i := 0; i < count; i++ {
		out[i] = float32(delta * float64(i))
	}
	return out
}
