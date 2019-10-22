//Package triangle provides a function to check type of a given triangle
package triangle

import "math"

// Kind triangle type
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if (a <= 0 || b <= 0 || c <= 0) || (a+b < c || a+c < b || b+c < a) {
		k = NaT
	} else if math.IsNaN(a+b+c) || math.IsInf(a+b+c, 0) {
		k = NaT
	} else if a == b && a == c {
		k = Equ
	} else if (a == b) || (b == c) || (a == c) {
		k = Iso
	} else {
		k = Sca
	}
	return k
}
