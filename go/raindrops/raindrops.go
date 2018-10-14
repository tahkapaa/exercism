/*
Package raindrops converts a number to a string,
the contents of which depend on the number's factors.

- If the number has 3 as a factor, output 'Pling'.
- If the number has 5 as a factor, output 'Plang'.
- If the number has 7 as a factor, output 'Plong'.
- If the number does not have 3, 5, or 7 as a factor,
  just pass the number's digits straight through.

Examples

- 28's factors are 1, 2, 4, **7**, 14, 28.
  - In raindrop-speak, this would be a simple "Plong".
- 30's factors are 1, 2, **3**, **5**, 6, 10, 15, 30.
  - In raindrop-speak, this would be a "PlingPlang".
- 34 has four factors: 1, 2, 17, and 34.
  - In raindrop-speak, this would be "34".
*/
package raindrops

import (
	"strconv"
)

// Constants for raindrop-speak
const (
	K3String = "Pling"
	K5String = "Plang"
	K7String = "Plong"
)

// Convert converts a number to a raindrop string
func Convert(n int) string {
	var s string
	if n%3 == 0 {
		s += K3String
	}
	if n%5 == 0 {
		s += K5String
	}
	if n%7 == 0 {
		s += K7String
	}

	if len(s) == 0 {
		return strconv.Itoa(n)
	}

	return s
}
