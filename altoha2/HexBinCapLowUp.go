package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HexBinCapLowUp(words []string) {
	for i := 0; i < len(words); i++ {
		if words[i] == "(hex)" {

			hex, _ := strconv.ParseInt(words[i-1], 16, 64)

			words[i-1] = fmt.Sprint(hex)
			words[i] = ""
		}
		if words[i] == "(bin)" {

			bin, _ := strconv.ParseInt(words[i-1], 2, 64)

			words[i-1] = fmt.Sprint(bin)
			words[i] = ""
		}
		if words[i] == ("(up)") {
			up := strings.ToUpper(words[i-1])

			words[i-1] = up
			words[i] = ""
		}
		if words[i] == ("(low)") {
			low := strings.ToLower(words[i-1])

			words[i-1] = low
			words[i] = ""
		}
		if words[i] == ("(cap)") {
			cap := strings.Title(words[i-1])

			words[i-1] = cap
			words[i] = ""
		}

	}
}
