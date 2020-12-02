/*
Package space implements a routine
 calculating, given an age in seconds,
 how old someone would be on misc. planets.
*/
package space

// Planet represents a planet as a string.
type Planet string
const(
		mercury Planet = "Mercury"
		venus   Planet = "Venus"
		earth   Planet = "Earth"
		mars    Planet = "Mars"
		jupiter Planet = "Jupiter"
		saturn  Planet = "Saturn"
		uranus  Planet = "Uranus"
		neptune Planet = "Neptune"
	)

// Orbital period in Earth years of misc. planets.
// const map
var opieyMap = func() map[Planet]float64 {
		return map[Planet]float64{
				"Mercury":   0.2408467,
				"Venus":     0.61519726,
				"Earth":     1.0,
				"Mars":      1.8808158,
				"Jupiter":  11.862615,
				"Saturn":   29.447498,
				"Uranus":   84.016846,
				"Neptune": 164.79132,
			}
	}
func opiey(planet Planet) float64 {
	return opieyMap()[planet]
}

// One Earth year in seconds.
const eyis float64 = 365.25 * 24 * 60 * 60;

// Age returns,
//  given the age of someone in seconds, and a planet,
//  how old this someone would be in years on this planet.
func Age(seconds float64, planet Planet) float64 {
	return seconds / eyis / opiey(planet);
}
