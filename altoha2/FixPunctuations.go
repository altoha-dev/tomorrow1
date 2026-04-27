package main

import "strings"

func FixPunctuations(words []string) {
	for i := 0; i < len(words); i++ {
		if words[i] == "" {
			continue
		}

		if strings.ContainsAny(words[i], ".,!?:;") {
			if i > 0 {

				target := i - 1
				for target > 0 && words[target] == "" {
					target--
				}
				words[target] = words[target] + words[i]
				words[i] = ""
			}
		}
	}
}

func PrepareText(text string) string {
	symbols := []string{".", ",", "!", "?", ":", ";"}
	for _, s := range symbols {
		text = strings.ReplaceAll(text, s, " "+s+" ")
	}
	text = strings.Join(strings.Fields(text), " ")
	for i := 0; i < 3; i++ {
		text = strings.ReplaceAll(text, ". . .", "...")
		text = strings.ReplaceAll(text, "! ?", "!?")
	}

	return text
}
