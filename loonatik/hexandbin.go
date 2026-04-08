package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func SplitText(s string) string {
	runes := []rune(s)
	var re1 strings.Builder

	for i := 0; i < len(runes); i++ {
		curr := runes[i]
		if unicode.IsPunct(curr) {
			re1.WriteRune(' ')
			continue
		}
		if i > 0 {
			prev := runes[i-1]

			if (unicode.IsDigit(prev) && unicode.IsLetter(curr)) ||
				(unicode.IsLetter(prev) && unicode.IsDigit(curr)) {
				re1.WriteRune(' ')
			}
		}
		re1.WriteRune(runes[i])
	}
	return strings.Join(strings.Fields(re1.String()), " ")
}

func HexBin(text string) string {
	words := strings.Fields(SplitText(text))

	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" && i > 0 {
			// (hex)
			res, err := strconv.ParseInt(words[i-1], 16, 64)
			// 1E
			if err == nil {

				words[i-1] = fmt.Sprint(res) // 1E ->30 перезапись
				words[i] = ""
			}

		}
		if words[i] == "(bin)" && i > 0 {
			// (bin)
			res, err := strconv.ParseInt(words[i-1], 2, 64)
			// (10)
			if err == nil {

				words[i-1] = fmt.Sprint(res) // 10 ->2
				words[i] = ""
				// (bin del)
			}

		}

	}

	return strings.Join(words, " ")
}
