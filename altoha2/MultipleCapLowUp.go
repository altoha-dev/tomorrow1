package main

import (
	"fmt"
	"strconv"
	"strings"
)

func MultipleCapLowUp(words []string) {
	for i := 0; i < len(words); i++ {
		up := strings.HasPrefix(words[i], "(up,")
		low := strings.HasPrefix(words[i], "(low,")
		cap := strings.HasPrefix(words[i], "(cap,")
		if (up || low || cap) && i+1 < len(words) {
			num := strings.Trim(words[i+1], ")")
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Error n")
				continue
			}
			for j := 1; j <= n; j++ {
				safe := i - j

				if up && safe >= 0 {
					words[i-j] = strings.ToUpper(words[i-j])
				}

				if low && safe >= 0 {
					words[i-j] = strings.ToLower(words[i-j])
				}

				if cap && safe >= 0 {
					words[i-j] = strings.Title(words[i-j])
				}

			}
			words[i] = ""
			words[i+1] = ""

		}

	}
}
