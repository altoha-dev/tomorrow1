package main

import "strings"

func FixPunctuations(words []string) {
	for i := 0; i < len(words); i++ {
		symbols := strings.ContainsAny(words[i], ".,!?:;")
		if symbols {
			words[i-1] = words[i-1] + words[i]
			words[i] = ""
		}
	}
}
