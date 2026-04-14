package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("+++++++++++++++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {
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
		if string(s[i]) == "-" {
			res += "-"
		}
		if string(s[i]) == "+" && string(s[i+1]) == "+" {
			return 0
		}

	}
	return num
}

// 12345
// 12345
// 0
// 0
// 1234
// -1234
// 0
// 0
