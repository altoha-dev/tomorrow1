package main

import (
	"fmt"
)

func main() {
	s := []int{5, 4, 3, 2, 1, 0}
	SortIntegerTable(s)
	fmt.Println(s)
}

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for y := i + 1; y < len(table); y++ {
			if table[i] > table[y] {
				table[i], table[y] = table[y], table[i]
			}
		}
	}
}

// [0 1 2 3 4 5]
