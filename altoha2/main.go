package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	textBytes, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("НЕ МОГУ ПРОЧИТАТЬ ФАЙЛ:", err)
		return
	}

	content := HandleApostrophes(string(textBytes))

	words := strings.Fields(content)

	HexBinCapLowUp(words)
	MultipleCapLowUp(words)

	tempText := strings.Join(words, " ")
	spacedText := PrepareText(tempText)

	finalWords := strings.Fields(spacedText)
	FixPunctuations(finalWords)
	HandleArticles(finalWords)

	result := strings.Join(finalWords, " ")
	finalResult := strings.Join(strings.Fields(result), " ")

	os.WriteFile(output, []byte(finalResult), 0o644)
}
