// Package space implements operations using many planets as reference.
package space

import "math"

type Planet string

// orbit stores the obital period in seconds of each planet.
var orbit = map[Planet]float64{
	"Earth":   31557600,
	"Mercury": 31557600 * 0.2408467,
	"Venus":   31557600 * 0.61519726,
	"Mars":    31557600 * 1.8808158,
	"Jupiter": 31557600 * 11.862615,
	"Saturn":  31557600 * 29.447498,
	"Uranus":  31557600 * 84.016846,
	"Neptune": 31557600 * 164.79132,
}

// Age returns the age in years using the p planet as reference.
func Age(t float64, p Planet) float64 {
	if v, ok := orbit[p]; ok {
		return t / v
	}
	return math.NaN()
}
