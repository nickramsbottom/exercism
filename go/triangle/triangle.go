// Package triangle has functionality to work with triangles
package triangle

import "math"

// Kind is the kind of triangle
type Kind string

// Identifiers used by the test program.
const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides determines the kind of triangle from side length
func KindFromSides(a, b, c float64) Kind {

	if math.IsNaN(a+b+c) || a+b <= c || a+c <= b || b+c <= a {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}
