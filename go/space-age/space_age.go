package space

const EarthSecPerYear = 365.25 * 24 * 60 * 60

func Age(seconds float64, planet string) float64 {
	m := make(map[string]float64)
	m["Earth"] = EarthSecPerYear
	m["Mercury"] = 0.2408467 * EarthSecPerYear
	m["Venus"] = 0.61519726 * EarthSecPerYear
	m["Mars"] = 1.8808158 * EarthSecPerYear
	m["Jupiter"] = 11.862615 * EarthSecPerYear
	m["Saturn"] = 29.447498 * EarthSecPerYear
	m["Uranus"] = 84.016846 * EarthSecPerYear
	m["Neptune"] = 164.79132 * EarthSecPerYear
	return seconds / m[planet]
}
