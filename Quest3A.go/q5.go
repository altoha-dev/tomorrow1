package main

import "github.com/01-edu/z01"

func main() {
	PrintStr("Hello World!")
}

func PrintStr(s string) {
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
	z01.PrintRune('\n')
}

// go run . | cat -e
// Hello World!
