package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("NO")
		return
	}
	inText := os.Args[1]
	outText := os.Args[2]

	text, err := os.ReadFile(inText)
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	result := SplitText(string(text))

	os.WriteFile(outText, []byte(result), 0o644)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	fmt.Println(result)
}
