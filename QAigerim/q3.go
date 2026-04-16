package main

import "fmt"

func main() {
	Slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(Slice[1 : len(Slice)-1])
}
