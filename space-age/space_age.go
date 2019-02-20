package space

// Age takes a string planet and seconds and does some arithmetic
// to calculate age based on that planets concept of a year relative to an earth year
func Age(seconds float64, planet string) float64 {
	//earth year in seconds
	const earthYear float64 = 31557600
	relativeAgeBasedOnEarthYears := seconds / earthYear

	switch planet {
	case "Mercury":
		relativeAgeBasedOnEarthYears = relativeAgeBasedOnEarthYears / 0.2408467
	}

	return relativeAgeBasedOnEarthYears
}
