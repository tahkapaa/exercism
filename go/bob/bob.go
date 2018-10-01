package bob

import (
	"strings"
)

// Hey returns Bob's answer to your remark
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case containsLetters(remark) && strings.ToUpper(remark) == remark && strings.HasSuffix(remark, "?"):
		return "Calm down, I know what I'm doing!"
	case respondSure(remark):
		return "Sure."
	case containsLetters(remark) && strings.ToUpper(remark) == remark:
		return "Whoa, chill out!"
	case remark == "":
		return "Fine. Be that way!"
	}
	return "Whatever."
}

func respondSure(remark string) bool {
	if !containsLetters(remark) && strings.HasSuffix(remark, "?") {
		return true
	}
	return !(strings.ToUpper(remark) == remark) && strings.HasSuffix(remark, "?")
}

func containsLetters(s string) bool {
	for _, r := range s {
		if r > 'A' && r < 'Ö' || r > 'a' && r < 'ö' {
			return true
		}
	}
	return false
}
