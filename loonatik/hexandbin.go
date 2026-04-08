package main

import (
	"strings"
	"unicode"
)

func SplitText(s string) string {
	runes := []rune(s)
	var re1 strings.Builder

	for i := 0; i < len(runes); i++ {
		if i > 0 {
			prev := runes[i-1]
			curr := runes[i]

			if (unicode.IsDigit(prev) && unicode.IsLetter(curr)) ||
				(unicode.IsLetter(prev) && unicode.IsDigit(curr)) {
				re1.WriteRune(' ')
			}
		}
		re1.WriteRune(runes[i])
	}
	return re1.String()
}
