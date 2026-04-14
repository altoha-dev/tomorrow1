package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}

func BasicAtoi2(s string) int {
	res := ""
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '1' && s[i] <= '9' {
			res += string((s[i]))
			num, _ = strconv.Atoi(res)
		}
		if string(s[i]) == " " {
			return 0
		}
	}
	return num
}

// $ go run .
// 12345
// 12345
// 0
// 0
// $
