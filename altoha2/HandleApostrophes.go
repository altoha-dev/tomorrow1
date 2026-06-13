package main

func HandleApostrophes(text string) string {
	runes := []rune(text)
	var result []rune
	inApostrophes := false

	for i := 0; i < len(runes); i++ {
		if runes[i] == '\'' {
			if !inApostrophes {

				inApostrophes = true
				result = append(result, runes[i])

				if i+1 < len(runes) && runes[i+1] == ' ' {
					i++
				}
			} else {

				inApostrophes = false

				if len(result) > 0 && result[len(result)-1] == ' ' {
					result = result[:len(result)-1]
				}
				result = append(result, runes[i])
			}
		} else {
			result = append(result, runes[i])
		}
	}
	return string(result)
}
