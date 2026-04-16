package main

import "fmt"

func main() {
	text := "substring"
	// res := ""
	fmt.Println(len(text))
	for k, v := range text {
		if k >= 2 && k <= 5 {
			fmt.Println("pito i ", k, "значение", string(v))
		}
	}
	// fmt.Println(res)
}

// s u b s t r i n g
// 0 1 2 3 4 5 6 7 8
// bstr
