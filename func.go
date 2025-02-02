package main

import "math"

func FtoC(f float64) float64 {
	c := (f - 32) * 5 / 9
	return roundFloat(c, 1)
}

func InToHpa(in float64) float64 {
	hpa := in / 0.029529980164712
	return roundFloat(hpa, 1)
}

func MphToMps(mph float64) float64 {
	mps := mph * 0.44704
	return roundFloat(mps, 1)
}

// Source: https://gosamples.dev/round-float//
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
