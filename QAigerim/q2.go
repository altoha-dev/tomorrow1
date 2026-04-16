package main

import "fmt"

func main() {
	Slice := []int{1, 2, 3}

	Slice = append(Slice, 4, 5, 6)

	fmt.Println(Slice)
}
