package space

// Planet an enum of the planets in the solar system
type Planet string

var secondsInEarthYear = 31557600.0

var relativeOrbitalPeriods = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age determines the age of someone in orbital periods of a planet
// given their absolute age in seconds
func Age(ageSeconds float64, planet Planet) float64 {
	return ageSeconds / planet.orbitalPeriod()
}

func (planet Planet) orbitalPeriod() float64 {
	return relativeOrbitalPeriods[planet] * secondsInEarthYear
}
