package Converter

import (
	"errors"
	"math"
)

func Radps2rpm(radps float64) float64 {
	return radps * 60 / (2 * math.Pi)
}

func Rpm2radps(rpm float64) float64 {
	return rpm * (2 * math.Pi) / 60
}

func Degps2Radps(degreesPerSecod float64) float64 {
	return degreesPerSecod * math.Pi / 180
}

func MapBetweenRanges(x, inMin, inMax, outMin, outMax float64) (float64, error) {
	if x < inMin {
		return 0, errors.New("value < inMin")
	}
	return (x-inMin)*(outMax-outMin)/(inMax-inMin) + outMin, nil
}
