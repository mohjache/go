package space

// Age takes a string planet and float seconds and calculates
// age in that planets years relative to an earth year
func Age(seconds float64, planet string) float64 {
	//earth year in seconds
	const earthYear float64 = 31557600
	relativeAgeBasedOnEarthYears := seconds / earthYear

	switch planet {
	case "Mercury":
		relativeAgeBasedOnEarthYears = relativeAgeBasedOnEarthYears / 0.2408467
	case "Venus":
		relativeAgeBasedOnEarthYears = relativeAgeBasedOnEarthYears / 0.61519726
	case "Mars":
		relativeAgeBasedOnEarthYears = relativeAgeBasedOnEarthYears / 1.8808158
	case "Jupiter":
		relativeAgeBasedOnEarthYears = relativeAgeBasedOnEarthYears / 11.862615
	case "Saturn":
		relativeAgeBasedOnEarthYears = relativeAgeBasedOnEarthYears / 29.447498
	case "Uranus":
		relativeAgeBasedOnEarthYears = relativeAgeBasedOnEarthYears / 84.016846
	case "Neptune":
		relativeAgeBasedOnEarthYears = relativeAgeBasedOnEarthYears / 164.79132
	}

	return relativeAgeBasedOnEarthYears
}
