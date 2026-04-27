package main

import "strings"

func HandleArticles(words []string) {
	for i := 0; i < len(words)-1; i++ {
		word := strings.ToLower(words[i])
		if word == "a" {
			nextWord := strings.ToLower(words[i+1])
			if len(nextWord) > 0 {
				firstChar := nextWord[0]
				if strings.ContainsAny(string(firstChar), "aeiouh") {
					if words[i] == "A" {
						words[i] = "An"
					} else {
						words[i] = "an"
					}
				}
			}
		}
	}
}

