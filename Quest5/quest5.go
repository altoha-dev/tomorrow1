package main

import "fmt"

func main() {
	num := 0

	f := &num

	num = *f

	fmt.Println(f)

	fmt.Println(num)

	num1 := 5

	*f = num1
}
