package main

import (
	"fmt"
)

func main() {
	l := StrLen("Hello World!")

	fmt.Println(l)
}

func StrLen(s string) int {
	count := 0

	for i := 0; i < len(s); i++ {
		count++
	}

	return count
}

// 12
