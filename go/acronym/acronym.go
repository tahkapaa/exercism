package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate converts a long name like Portable Network Graphics to its acronym (PNG).
func Abbreviate(s string) string {
	words := strings.FieldsFunc(s, split)
	abb := make([]rune, len(words))

	for i, word := range words {
		r := []rune(word)
		abb[i] = unicode.ToUpper(r[0])
	}
	return string(abb)
}

func split(r rune) bool {
	return r == ' ' || r == '-'
}
