package space

//Planet type simple string representing planet
type Planet string

//Orbital periods in Earth years
const (
	Earth   = 31557600
	Mercury = 0.2408467
	Venus   = 0.61519726
	Mars    = 1.8808158
	Jupiter = 11.862615
	Saturn  = 29.447498
	Uranus  = 84.016846
	Neptune = 164.79132
)

//Age takes seconds as float 64 and planet name as planet type and returns as float64
// how old in years would be in respective planet. If unknow planet name is given function returns 0.0
func Age(seconds float64, planet Planet) float64 {

	var earthYears float64 = seconds / Earth

	switch planet {
	case "Earth":
		return earthYears
	case "Mercury":
		return earthYears / Mercury
	case "Venus":
		return earthYears / Venus
	case "Mars":
		return earthYears / Mars
	case "Jupiter":
		return earthYears / Jupiter
	case "Saturn":
		return earthYears / Saturn
	case "Uranus":
		return earthYears / Uranus
	case "Neptune":
		return earthYears / Neptune
	default:
		return 0.0
	}

}
