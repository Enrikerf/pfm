package Math

func MapBetweenRanges(x, inMin, inMax, outMin, outMax float64) float64 {
	if x < inMin {
		panic("map value below inMin")
	}
	return (x-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}
