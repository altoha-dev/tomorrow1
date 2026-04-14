package main

import (
	"fmt"
)

func UltimateDivMod(a *int, b *int) {
	temp := *a / *b
	temp1 := *a % *b
	*a = temp

	*b = temp1
}

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
