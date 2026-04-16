package main

import "fmt"

func main() {
	slice := []string{"altoha", "krasava", "!", "1"}

	res := ""

	for i := 0; i < len(slice); i++ {
		res += (slice[i])
	}
	fmt.Println(res)
}
