/*
Package triangle determines if a triangle is equilateral, isosceles, or scalene.

An _equilateral_ triangle has all three sides the same length.

An _isosceles_ triangle has at least two sides the same length. (It is sometimes
specified as having exactly two sides the same length, but for the purposes of
this exercise we'll say at least two.)

A _scalene_ triangle has all sides of different lengths.
*/
package triangle

import (
	"math"
)

// Kind contains the type of the triangle
type Kind int

const (
	// NaT is not a triangle
	NaT = -1
	// Equ is equilateral triangle
	Equ = 3
	// Iso is an isosceles triangle
	Iso = 1
	// Sca is a scalene triangle
	Sca = 0
)

// KindFromSides determines the type of the given triangle
func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	if a+b < c || b+c < a || c+a < b {
		return NaT
	}
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	var equalSides Kind
	if a == b {
		equalSides++
	}
	if b == c {
		equalSides++
	}
	if c == a {
		equalSides++
	}

	return equalSides
}
