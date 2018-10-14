/*
Package space given an age in seconds, calculate how old someone would be on:

   - Earth: orbital period 365.25 Earth days, or 31557600 seconds
   - Mercury: orbital period 0.2408467 Earth years
   - Venus: orbital period 0.61519726 Earth years
   - Mars: orbital period 1.8808158 Earth years
   - Jupiter: orbital period 11.862615 Earth years
   - Saturn: orbital period 29.447498 Earth years
   - Uranus: orbital period 84.016846 Earth years
   - Neptune: orbital period 164.79132 Earth years

So if you were told someone were 1,000,000,000 seconds old, you should
be able to say that they're 31.69 Earth-years old.
*/
package space

// Planet defines the name of the planet
type Planet string

const secondsInYear = 31557600

var orbitalPeriods = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns how old someone of age given in seconds, would be in the given planet.
func Age(seconds float64, p Planet) float64 {
	return seconds / (secondsInYear * orbitalPeriods[p])
}
