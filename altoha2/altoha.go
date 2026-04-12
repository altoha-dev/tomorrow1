package main

import (
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	text, err := os.ReadFile(input)
	if err != nil {
		return
	}

	words := strings.Fields(string(text))

	HexBinCapLowUp(words)

	MultipleCapLowUp(words)

	result := strings.Join(words, " ")
	FixPunctuations(words)
	os.WriteFile(output, []byte(result), 0o644)
}
