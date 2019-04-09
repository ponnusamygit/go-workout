package space

// Age will return age on differnt planets
func Age(seconds float64, planet string) float64 {
	return seconds / planetYear(planet)
}

func planetYear(planet string) float64 {
	eathYear := 31556952.0
	if planet == "Mercury" {
		// - Mercury: orbital period 0.2408467 Earth years
		return eathYear * 0.2408467
	} else if planet == "Venus" {
		// - Venus: orbital period 0.61519726 Earth years
		return eathYear * 0.61519726
	} else if planet == "Mars" {
		// - Mars: orbital period 1.8808158 Earth years
		return eathYear * 1.8808158
	} else if planet == "Jupiter" {
		// - Jupiter: orbital period 11.862615 Earth years
		return eathYear * 11.862615
	} else if planet == "Saturn" {
		// - Saturn: orbital period 29.447498 Earth years
		return eathYear * 29.447498
	} else if planet == "Uranus" {
		// - Uranus: orbital period 84.016846 Earth years
		return eathYear * 84.016846
	} else if planet == "Neptune" {
		// - Neptune: orbital period 164.79132 Earth years
		return eathYear * 164.79132
	} else {
		// - Earth: orbital period 365.25 Earth days, or 31557600 seconds
		return eathYear
	}
}
